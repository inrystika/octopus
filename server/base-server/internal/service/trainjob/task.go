package trainjob

import (
	"context"
	"encoding/json"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/data/dao/model"
	"server/common/constant"
	"server/common/leaderleaselock"
	"server/common/utils"
	"server/common/utils/collections/set"
	"strings"
	"time"

	typeJob "volcano.sh/apis/pkg/apis/batch/v1alpha1"

	"gonum.org/v1/gonum/floats"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	taskPageSize = 100
	leaseLock    = "trainjobleaselock"
)

func (s *trainJobService) getOwner(trainJob *model.TrainJob) (string, api.BillingOwnerType) {
	var ownerId string
	var ownerType api.BillingOwnerType

	if trainJob.WorkspaceId == constant.SYSTEM_WORKSPACE_DEFAULT {
		ownerId = trainJob.UserId
		ownerType = api.BillingOwnerType_BOT_USER
	} else {
		ownerId = trainJob.WorkspaceId
		ownerType = api.BillingOwnerType_BOT_SPACE
	}

	return ownerId, ownerType
}

func (s *trainJobService) calAmount(startAt int64, endedAt int64, price float64) float64 {
	return float64(endedAt-startAt) * price / 3600.0
}

func (s *trainJobService) trainJobBilling(ctx context.Context) {

	k8sns := utils.GetEnvOrDefault("K8S_NAMESPACE", "default")
	rdlock := leaderleaselock.NewLeaderLeaselock(k8sns, leaseLock, s.data.Cluster.GetClusterConfig())
	rdlock.RunOrRetryLeaderElection(ctx, func(ctx context.Context) {
		s.log.Infof(ctx, "train job billing service acquire %v", leaseLock)

		go func() {
			BillingPeriodSec := int64(1800) //默认值
			if s.conf.Service.BillingPeriodSec > 0 {
				BillingPeriodSec = s.conf.Service.BillingPeriodSec
			}
			s.log.Infof(ctx, "train job billing service billing period time is %v", BillingPeriodSec)
			wait.Until(func() {
				utils.HandlePanic(ctx, func(i ...interface{}) {
					s.log.Info(ctx, "start train-job-billing cron service.....")
					for pageIndex := 1; ; pageIndex++ {
						trainJobs, _, err := s.data.TrainJobDao.GetTrainJobList(ctx, &model.TrainJobListQuery{
							PageIndex: pageIndex,
							PageSize:  taskPageSize,
							PayStatus: api.BillingPayRecordStatus_BPRS_PAYING,
						})
						if err != nil {
							s.log.Errorf(ctx, "List TrainJob err: %s", err)
							break
						}

						if len(trainJobs) == 0 {
							s.log.Info(ctx, "There is no trainJob to pay!")
							break
						}

						//这些任务可能没有启动时间，但状态却是终止的，这些任务不计费,设置计费状态为完成。
						for _, j := range trainJobs {
							if j.StartedAt == nil && utils.IsCompletedState(j.Status) {
								j.PayStatus = api.BillingPayRecordStatus_BPRS_PAY_COMPLETED
								err = s.data.TrainJobDao.UpdateTrainJob(ctx, j)
								if err != nil {
									s.log.Errorf(ctx, "update ineffective job to completed err: %s", err)
									break
								}
							}
						}

						//删除后再查询
						trainJobs, _, err = s.data.TrainJobDao.GetTrainJobList(ctx, &model.TrainJobListQuery{
							PageIndex: pageIndex,
							PageSize:  taskPageSize,
							PayStatus: api.BillingPayRecordStatus_BPRS_PAYING,
						})

						if err != nil {
							s.log.Errorf(ctx, "List TrainJob err: %s", err)
							break
						}

						if len(trainJobs) == 0 {
							s.log.Info(ctx, "There is no trainJob to pay!")
							break
						}

						//计费逻辑
						trainJobIds := make([]string, 0)
						for _, j := range trainJobs {
							trainJobIds = append(trainJobIds, j.Id)
						}
						trainJobIds = set.NewStrings(trainJobIds...).Values()

						trainJobMap := map[string]*model.TrainJob{}
						for _, job := range trainJobs {
							trainJobMap[job.Id] = job
						}

						jobNs := map[string]string{}
						for _, id := range trainJobIds {
							jobNs[id] = trainJobMap[id].UserId
						}

						details := make([]*typeJob.JobStatusDetail, 0)
						for _, id := range trainJobIds {
							info, err := s.getJobDetail(ctx, id)
							if err != nil {
								s.log.Errorf(ctx, "GetJob err: %s", err)
							} else {
								details = append(details, info)
							}
						}

						detailMap := map[string]*typeJob.JobStatusDetail{}
						for _, d := range details {
							detailMap[d.Job.ID] = d
						}

						for _, j := range trainJobs {
							//判断任务是否已经启动。如果没有启动时间，则说明未启动，不计费。
							if j.StartedAt == nil {
								//s.log.Info(ctx, "job "+j.Id+"no need to pay, because it is not started!")
								continue
							}
							payAmount := 0.0
							//job已经启动，则以job的启动时间作为每个task的启动时间，以此为计费起始点。
							payStartAt := j.StartedAt.Unix()
							s.log.Info(ctx, "train bill service try to calculate job pay amount, jobId is: "+j.Id)
							now := time.Now().Unix()
							specPriceMap := map[int]float64{}
							for _, p := range j.ResSpecPrice {
								specPriceMap[p.Task] = p.Price
							}

							trainJob := trainJobMap[j.Id]
							ownerId, ownerType := s.getOwner(trainJob)

							detail := detailMap[j.Id]
							for ti, t := range detail.Tasks {
								for _, r := range t.Replicas { //计算副本消费
									var endAt int64
									//查看副本任务是否终止，以便获取副本终止时间。
									if utils.IsCompletedState(r.State) {
										// 副本状态终止，但无终止时间。
										if r.FinishedAt == nil {
											//若job终止时间也缺失，系统级错误，结束时间 = 启动时间，不计入费用！
											if j.CompletedAt == nil {
												s.log.Errorf(ctx, j.Id+"'s replica finished-time is null && job finished time is also null!")
												s.log.Info(ctx, "Attention!!! System err, user no need to pay! job id is :"+j.Id)
												endAt = r.StartAt.Unix()
											} else {
												//若job终止时间存在, 则将其作为副本终止时间，完成计费。
												s.log.Warn(ctx, "replica finished-time is null! So instead to use job finished time!")
												endAt = j.CompletedAt.Unix()
											}
										} else {
											endAt = r.FinishedAt.Unix()
										}
									} else if strings.EqualFold(r.State, constant.RUNNING) {
										//副本仍在running，则取当前系统时间，作为该周期计费终止点。
										endAt = now
									}
									//计算副本用时，启动时间恒定，变化的只有终止时间。
									if endAt != 0 {
										payAmount += s.calAmount(payStartAt, endAt, specPriceMap[ti])
									}
								}
							}
							payAmount = floats.Round(payAmount, common.BillingPrecision)

							var payStatus api.BillingPayRecordStatus
							var payEndAt int64
							if utils.IsCompletedState(detail.Job.State) {
								payEndAt = detail.Job.FinishedAt.Unix()
								payStatus = api.BillingPayRecordStatus_BPRS_PAY_COMPLETED
							} else {
								payEndAt = now
								payStatus = api.BillingPayRecordStatus_BPRS_PAYING
							}

							if payAmount <= j.PayAmount && payStatus != api.BillingPayRecordStatus_BPRS_PAY_COMPLETED {
								continue
							}

							extraInfo := make(map[string]string)
							if ownerType == api.BillingOwnerType_BOT_SPACE {
								extraInfo = common.GetExtraInfo(j.UserId)
							}
							_, err := s.billingService.Pay(ctx, &api.PayRequest{
								OwnerId:   ownerId,
								OwnerType: ownerType,
								Amount:    payAmount,
								BizType:   api.BillingBizType_BBT_TRAIN_JOB,
								BizId:     j.Id,
								Title:     trainJob.Name,
								StartedAt: payStartAt,
								EndedAt:   payEndAt,
								Status:    payStatus,
								ExtraInfo: extraInfo,
							})
							if err != nil {
								s.log.Errorf(ctx, "Pay err: %s", err)
								continue
							}

							startAt := time.Unix(payStartAt, 0)
							endAt := time.Unix(payEndAt, 0)
							err = s.data.TrainJobDao.UpdateTrainJob(ctx, &model.TrainJob{
								Id:           j.Id,
								PayStatus:    payStatus,
								PayStartedAt: &startAt,
								PayEndedAt:   &endAt,
								PayAmount:    payAmount,
							})
							if err != nil {
								s.log.Errorf(ctx, "Update train job selective err: %s", err)
								continue
							}

							owner, err := s.billingService.GetBillingOwner(ctx, &api.GetBillingOwnerRequest{
								OwnerId:   ownerId,
								OwnerType: ownerType,
							})
							if err != nil {
								s.log.Errorf(ctx, "GetBillingOwner err: %s", err)
								continue
							}
							if s.conf.Service.StopWhenArrears && owner.BillingOwner.Amount < 0 {
								_, err = s.StopJob(ctx, &api.StopJobRequest{Id: j.Id, Operation: "system stop job due to arrears"})
								if err != nil {
									s.log.Errorf(ctx, "StopJob err: %s", err)
									continue
								}
								s.log.Infof(ctx, "StopJob due to arrears, jobId: %s", j.Id)
							}
						}
					}
				})()
			}, time.Duration(BillingPeriodSec)*time.Second, ctx.Done())
		}()
	})
}

func (s *trainJobService) trainJobUpdateStaus(ctx context.Context) {
	// 任务状态更新处理逻辑
	go func() {
		utils.HandlePanic(ctx, func(i ...interface{}) {
			for {
				select {
				case job := <-s.updatedJob:
					trainJob, err := s.data.TrainJobDao.GetTrainJob(context.TODO(), job.Name)
					if err != nil {
						s.log.Warn(context.TODO(), "GetTrainJob err when onJobUpdate:"+job.Name, err)
						continue
					}
					state := utils.MapPhaseToState(typeJob.JobPhase(job.Status.State.Phase))
					if utils.IsCompletedState(trainJob.Status) {
						continue
					}

					update := &model.TrainJob{
						Id:     job.Name,
						Status: state,
					}

					now := time.Now()
					if strings.EqualFold(state, constant.RUNNING) && strings.EqualFold(trainJob.Status, constant.PENDING) {
						update.StartedAt = &now
					} else if utils.IsCompletedState(state) {
						update.CompletedAt = &now
					}

					status := utils.Format(job.Name, "trainJob", job.Namespace, "", "", job)
					if nil != status {
						buf, err := json.Marshal(status)
						if err != nil {
							s.log.Error(context.TODO(), "UpdateTrainJob err when onJobUpdate:"+job.Name, err)
						}
						update.Detail = string(buf)
					}

					err = s.data.TrainJobDao.UpdateTrainJob(context.TODO(), update)
					if err != nil {
						s.log.Error(context.TODO(), "UpdateTrainJob err when onJobUpdate:"+job.Name, err)
						continue
					}

					s.sendEmail(trainJob.UserId, trainJob.Name, trainJob.Status, update.Status)

					if utils.IsCompletedState(state) {
						err = s.addModel(context.TODO(), trainJob)
						if err != nil {
							s.log.Error(context.TODO(), err)
						}
						err = s.data.Cluster.DeleteJob(context.TODO(), job.Namespace, job.Name)
						if err != nil {
							s.log.Error(context.TODO(), "DeleteJob err when onJobUpdate:"+job.Name, err)
						}
					}
				case <-ctx.Done():
					return
				}
			}
		})()
	}()
}
