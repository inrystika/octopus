package develop

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/data/dao/model"
	"server/base-server/internal/data/pipeline"
	"server/common/constant"
	"server/common/leaderleaselock"
	"server/common/utils"
	"server/common/utils/collections/set"
	"time"

	"gonum.org/v1/gonum/floats"
	"k8s.io/apimachinery/pkg/util/wait"
)

const (
	taskPageSize = 100
	leaseLock    = "notebookleaselock"
)

func (s *developService) getOwner(notebook *model.Notebook) (string, api.BillingOwnerType) {
	var ownerId string
	var ownerType api.BillingOwnerType

	if notebook.WorkspaceId == constant.SYSTEM_WORKSPACE_DEFAULT {
		ownerId = notebook.UserId
		ownerType = api.BillingOwnerType_BOT_USER
	} else {
		ownerId = notebook.WorkspaceId
		ownerType = api.BillingOwnerType_BOT_SPACE
	}

	return ownerId, ownerType
}

func (s *developService) startNotebookTask() {
	ctx := context.Background()
	k8sns := utils.GetEnvOrDefault("K8S_NAMESPACE", "default")
	lock := leaseLock
	if s.conf.App.IsDev {
		lock = lock + "-dev"
	}
	rdlock := leaderleaselock.NewLeaderLeaselock(k8sns, lock, s.data.Cluster.GetClusterConfig())
	rdlock.RunOrRetryLeaderElection(ctx, func(ctx context.Context) {
		s.log.Infof(ctx, "acquire %v", lock)
		go func() {
			wait.Until(func() {
				utils.HandlePanic(ctx, func(i ...interface{}) {
					for pageIndex := 1; ; pageIndex++ {
						notebookJobs, err := s.data.DevelopDao.ListNotebookJob(ctx, &model.NotebookJobQuery{
							PageIndex:   pageIndex,
							PageSize:    taskPageSize,
							StartedAtLt: time.Now().Add(-time.Second * time.Duration(s.conf.Service.Develop.AutoStopIntervalSec)).Unix(),
							StatusList:  pipeline.NonCompletedStates(),
						})

						if err != nil {
							s.log.Errorf(ctx, "ListNotebookJob err: %s", err)
							break
						}

						if len(notebookJobs) == 0 {
							break
						}

						for _, j := range notebookJobs {
							_, err := s.StopNotebook(ctx, &api.StopNotebookRequest{Id: j.NotebookId})
							if err != nil {
								s.log.Errorf(ctx, "StopNotebook err: %s", err)
								continue
							}
						}
					}
				})()
			}, 1*time.Minute, ctx.Done())
		}()

		go func() {
			BillingPeriodSec := int64(1800) //默认值
			if s.conf.Service.BillingPeriodSec > 0 {
				BillingPeriodSec = s.conf.Service.BillingPeriodSec
			}
			wait.Until(func() {
				utils.HandlePanic(ctx, func(i ...interface{}) {
					for pageIndex := 1; ; pageIndex++ {
						notebookJobs, err := s.data.DevelopDao.ListNotebookJob(ctx, &model.NotebookJobQuery{
							PageIndex: pageIndex,
							PageSize:  taskPageSize,
							PayStatus: api.BillingPayRecordStatus_BPRS_PAYING,
						})
						if err != nil {
							s.log.Errorf(ctx, "ListNotebookJob err: %s", err)
							break
						}

						if len(notebookJobs) == 0 {
							break
						}

						notebookIds := make([]string, 0)
						jobIds := make([]string, 0)
						for _, j := range notebookJobs {
							notebookIds = append(notebookIds, j.NotebookId)
							jobIds = append(jobIds, j.Id)
						}
						notebookIds = set.NewStrings(notebookIds...).Values()
						jobIds = set.NewStrings(jobIds...).Values()

						notebooks, _, err := s.data.DevelopDao.ListNotebook(ctx, &model.NotebookQuery{Ids: notebookIds})
						if err != nil {
							s.log.Errorf(ctx, "ListNotebook err: %s", err)
							continue
						}
						notebookMap := map[string]*model.Notebook{}
						for _, n := range notebooks {
							notebookMap[n.Id] = n
						}

						details, err := s.data.Pipeline.BatchGetJobDetail(ctx, jobIds)
						if err != nil {
							s.log.Errorf(ctx, "BatchGetJobDetail err: %s", err)
							continue
						}
						detailMap := map[string]*pipeline.JobStatusDetail{}
						for _, d := range details.Details {
							detailMap[d.Job.ID] = d
						}

						for _, j := range notebookJobs {
							utils.HandlePanic(ctx, func(i ...interface{}) {
								if j.StartedAt != nil {
									notebook := notebookMap[j.NotebookId]
									ownerId, ownerType := s.getOwner(notebook)

									var payEndAt int64
									var payStatus api.BillingPayRecordStatus
									if pipeline.IsCompletedState(j.Status) {
										payEndAt = j.StoppedAt.Unix()
										payStatus = api.BillingPayRecordStatus_BPRS_PAY_COMPLETED
									} else {
										if detailMap[j.Id].Job.FinishedAt != nil { //非用户stop且pod停止后未通知到notebook的情况
											payEndAt = detailMap[j.Id].Job.FinishedAt.Unix()
											payStatus = api.BillingPayRecordStatus_BPRS_PAY_COMPLETED
										} else {
											payEndAt = time.Now().Unix()
											payStatus = api.BillingPayRecordStatus_BPRS_PAYING
										}
									}

									payAmount := floats.Round(float64(payEndAt-j.StartedAt.Unix())*float64(j.ResourceSpecPrice)/3600.0, common.BillingPrecision)
									if payAmount <= j.PayAmount && payStatus != api.BillingPayRecordStatus_BPRS_PAY_COMPLETED {
										s.log.Infof(ctx, "amount less than amount in db and is not completed status")
										return
									}

									extraInfo := make(map[string]string)
									if ownerType == api.BillingOwnerType_BOT_SPACE {
										extraInfo = common.GetExtraInfo(notebook.UserId)
									}
									_, err := s.billingService.Pay(ctx, &api.PayRequest{
										OwnerId:   ownerId,
										OwnerType: ownerType,
										Amount:    payAmount,
										BizType:   api.BillingBizType_BBT_NOTEBOOK,
										BizId:     j.Id,
										Title:     notebook.Name,
										StartedAt: j.StartedAt.Unix(),
										EndedAt:   payEndAt,
										Status:    payStatus,
										ExtraInfo: extraInfo,
									})
									if err != nil {
										s.log.Errorf(ctx, "Pay err: %s", err)
										return
									}

									endAt := time.Unix(payEndAt, 0)
									err = s.data.DevelopDao.UpdateNotebookJobSelective(ctx, &model.NotebookJob{
										Id:           j.Id,
										PayStatus:    payStatus,
										PayStartedAt: j.StartedAt,
										PayEndedAt:   &endAt,
										PayAmount:    payAmount,
									})
									if err != nil {
										s.log.Errorf(ctx, "UpdateNotebookJobSelective err: %s", err)
										return
									}
								} else {
									if pipeline.IsCompletedState(j.Status) { //还没start就停止的情况
										err = s.data.DevelopDao.UpdateNotebookJobSelective(ctx, &model.NotebookJob{
											Id:        j.Id,
											PayStatus: api.BillingPayRecordStatus_BPRS_PAY_COMPLETED,
										})
										if err != nil {
											s.log.Errorf(ctx, "UpdateNotebookJobSelective err: %s", err)
											return
										}
									}
								}
							})()
						}
					}
				})()
			}, time.Duration(BillingPeriodSec)*time.Second, ctx.Done())
		}()
	})

}
