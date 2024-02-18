package develop

import (
	"context"
	"encoding/json"
	"fmt"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/data/dao/model"
	commapi "server/common/api/v1"
	"server/common/constant"
	"server/common/leaderleaselock"
	"server/common/utils"
	"server/common/utils/collections/set"
	"strings"
	"time"

	typeJob "volcano.sh/apis/pkg/apis/batch/v1alpha1"

	"k8s.io/client-go/tools/cache"

	nav1 "nodeagent/apis/agent/v1"

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
		lock = lock + "-dev11"
	}
	rdlock := leaderleaselock.NewLeaderLeaselock(k8sns, lock, s.data.Cluster.GetClusterConfig())
	rdlock.RunOrRetryLeaderElection(ctx, func(ctx context.Context) {
		s.log.Infof(ctx, "acquire %v", lock)
		go func() {
			wait.Until(func() {
				utils.HandlePanic(ctx, func(i ...interface{}) {
					for pageIndex := 1; ; pageIndex++ {
						notebookJobs, err := s.data.DevelopDao.ListNotebookJob(ctx, &model.NotebookJobQuery{
							PageIndex: pageIndex,
							PageSize:  taskPageSize,
							//StartedAtLt: time.Now().Add(-time.Second * time.Duration(s.conf.Service.Develop.AutoStopIntervalSec)).Unix(),
							StatusList: utils.NonCompletedStates(),
						})

						if err != nil {
							s.log.Errorf(ctx, "ListNotebookJob err: %s", err)
							break
						}

						if len(notebookJobs) == 0 {
							break
						}

						nbs := make(map[string]*model.Notebook)
						nbIds := make([]string, 0)
						for _, j := range notebookJobs {
							nbIds = append(nbIds, j.NotebookId)
						}

						notebooks, _, err := s.data.DevelopDao.ListNotebook(ctx, &model.NotebookQuery{Ids: nbIds})
						for _, n := range notebooks {
							nbs[n.Id] = n
						}

						for _, j := range notebookJobs {
							duration := int64(0)
							if nbs[j.NotebookId].AutoStopDuration == 0 {
								duration = s.conf.Service.Develop.AutoStopIntervalSec
							} else {
								duration = nbs[j.NotebookId].AutoStopDuration
							}

							if duration > 0 && j.StartedAt != nil && time.Now().Sub(*j.StartedAt).Seconds() >= float64(duration) {
								_, err := s.StopNotebook(ctx, &api.StopNotebookRequest{Id: j.NotebookId})
								if err != nil {
									s.log.Errorf(ctx, "StopNotebook err: %s", err)
									continue
								}
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
						nbJobs, err := s.data.DevelopDao.ListNotebookJob(ctx, &model.NotebookJobQuery{
							PageIndex: pageIndex,
							PageSize:  taskPageSize,
							PayStatus: api.BillingPayRecordStatus_BPRS_PAYING,
						})
						if err != nil {
							s.log.Errorf(ctx, "ListNotebookJob err: %s", err)
							break
						}

						if len(nbJobs) == 0 {
							break
						}

						nbIds := make([]string, 0)
						jobIds := make([]string, 0)
						for _, j := range nbJobs {
							nbIds = append(nbIds, j.NotebookId)
							jobIds = append(jobIds, j.Id)
						}
						nbIds = set.NewStrings(nbIds...).Values()
						jobIds = set.NewStrings(jobIds...).Values()

						nbs, _, err := s.data.DevelopDao.ListNotebook(ctx, &model.NotebookQuery{Ids: nbIds})
						if err != nil {
							s.log.Errorf(ctx, "ListNotebook err: %s", err)
							continue
						}
						nbMap := map[string]*model.Notebook{}
						for _, n := range nbs {
							nbMap[n.Id] = n
						}

						nbjNs := map[string]string{}
						for _, j := range nbJobs {
							nb := nbMap[j.NotebookId]
							nbjNs[j.Id] = nb.UserId
						}

						details := make([]*typeJob.JobStatusDetail, 0)
						for _, id := range jobIds {
							info, err := s.getJobDetail(ctx, nbjNs[id], id)
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

						for _, j := range nbJobs {
							utils.HandlePanic(ctx, func(i ...interface{}) {
								nb := nbMap[j.NotebookId]
								ownerId, ownerType := s.getOwner(nb)

								var payEndAt int64
								var payStatus api.BillingPayRecordStatus
								if utils.IsCompletedState(j.Status) {
									if j.StoppedAt == nil {
										payEndAt = time.Now().Unix()
									} else {
										payEndAt = j.StoppedAt.Unix()
									}
									payStatus = api.BillingPayRecordStatus_BPRS_PAY_COMPLETED
								} else {
									payEndAt = time.Now().Unix()
									payStatus = api.BillingPayRecordStatus_BPRS_PAYING
								}

								prices := make([]float64, 0)
								for i := 0; i < nb.TaskNumber; i++ {
									prices = append(prices, j.ResourceSpecPrice)
								}
								payAmount, startTime := common.CalculateAmount(ctx, detailMap[j.Id], prices)
								if payAmount > 0 {
									extraInfo := make(map[string]string)
									if ownerType == api.BillingOwnerType_BOT_SPACE {
										extraInfo = common.GetExtraInfo(nb.UserId)
									}
									_, err := s.billingService.Pay(ctx, &api.PayRequest{
										OwnerId:   ownerId,
										OwnerType: ownerType,
										Amount:    payAmount,
										BizType:   api.BillingBizType_BBT_NOTEBOOK,
										BizId:     j.Id,
										Title:     nb.Name,
										StartedAt: startTime.Unix(),
										EndedAt:   payEndAt,
										Status:    payStatus,
										ExtraInfo: extraInfo,
									})
									if err != nil {
										s.log.Errorf(ctx, "Pay err: %s", err)
										return
									}
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

								owner, err := s.billingService.GetBillingOwner(ctx, &api.GetBillingOwnerRequest{
									OwnerId:   ownerId,
									OwnerType: ownerType,
								})
								if err != nil {
									s.log.Errorf(ctx, "GetBillingOwner err: %s", err)
									return
								}
								if s.conf.Service.StopWhenArrears && owner.BillingOwner.Amount < 0 {
									_, err = s.StopNotebook(ctx, &api.StopNotebookRequest{Id: j.NotebookId, Operation: "system stop job due to arrears"})
									if err != nil {
										s.log.Errorf(ctx, "StopNotebook err: %s", err)
										return
									}
									s.log.Info(ctx, "StopNotebook due to arrears, jobId: %s", j.Id)
								}
							})()
						}
					}
				})()
			}, time.Duration(BillingPeriodSec)*time.Second, ctx.Done())
		}()

		// 任务状态更新处理逻辑
		go func() {
			utils.HandlePanic(ctx, func(i ...interface{}) {
				for {
					select {
					case job := <-s.updatedJob:
						nbJob, err := s.data.DevelopDao.GetNotebookJob(ctx, job.Name)
						if err != nil {
							s.log.Warn(ctx, "GetTrainJob err when onJobUpdate:"+job.Name, err)
							continue
						}

						state := utils.MapPhaseToState(typeJob.JobPhase(job.Status.State.Phase))

						if utils.IsCompletedState(nbJob.Status) {
							continue
						}

						nb, err := s.data.DevelopDao.GetNotebook(ctx, nbJob.NotebookId)
						if err != nil {
							s.log.Error(ctx, "GetNotebook err when onJobUpdate:"+job.Name, err)
							continue
						}

						nbUp := &model.Notebook{
							NotebookJobId: job.Name,
							Status:        state,
						}

						nbJobUp := &model.NotebookJob{
							Id:     job.Name,
							Status: state,
						}

						status := utils.Format(job.Name, "notebook", job.Namespace, "", "", job)
						if nil != status {
							buf, err := json.Marshal(status)
							if err != nil {
								s.log.Error(context.TODO(), "UpdateNotebook err when onJobUpdate: "+job.Name, err)
							}
							nbJobUp.Detail = string(buf)
						}

						now := time.Now()
						record := &model.NotebookEventRecord{
							Time:       now,
							NotebookId: nb.Id,
						}
						pendingToRunning := false
						if strings.EqualFold(state, constant.RUNNING) && strings.EqualFold(nbJob.Status, constant.PENDING) {
							pendingToRunning = true
							nbJobUp.StartedAt = &now
							record.Type = commapi.NotebookEventRecordType_RUN
						} else if utils.IsCompletedState(state) {
							nbJobUp.StoppedAt = &now
							nbJobUp.Status = constant.STOPPED //转为stopped
							nbUp.Status = constant.STOPPED    //转为stopped

							err = s.deleteIngress(ctx, nb, nbJob)
							if err != nil {
								s.log.Error(ctx, "deleteIngress err when onJobUpdate:"+job.Name, err)
							}

							err = s.deleteService(ctx, nb, nbJob)
							if err != nil {
								s.log.Error(ctx, "deleteService err when onJobUpdate:"+job.Name, err)
							}
							record.Type = commapi.NotebookEventRecordType_STOP
						}

						err = s.data.DevelopDao.UpdateNotebookSelectiveByJobId(ctx, nbUp)
						if err != nil {
							s.log.Error(ctx, "UpdateNotebookSelectiveByJobId err when onJobUpdate:"+job.Name, err)
						}

						err = s.data.DevelopDao.UpdateNotebookJobSelective(ctx, nbJobUp)
						if err != nil {
							s.log.Error(ctx, "UpdateNotebookJobSelective err when onJobUpdate:"+job.Name, err)
						}

						if utils.IsCompletedState(state) || pendingToRunning {
							err = s.data.DevelopDao.CreateNotebookEventRecord(ctx, record)
							if err != nil { // 插入事件记录出错只打印
								s.log.Error(ctx, "create notebook event record error:", err)
							}
						}

						if utils.IsCompletedState(state) {
							err = s.data.Cluster.DeleteJob(context.TODO(), job.Namespace, job.Name)
							if err != nil {
								s.log.Error(context.TODO(), "DeleteJob err when onJobUpdate: "+job.Name, err)
							}
						}
					case <-ctx.Done():
						return
					}
				}
			})()
		}()

		go func() {
			nodeActionInformer := s.data.Cluster.GetNodeActionInformer()
			nodeActionInformer.Informer().AddEventHandlerWithResyncPeriod(
				cache.FilteringResourceEventHandler{
					FilterFunc: func(obj interface{}) bool {
						na := obj.(*nav1.NodeAction)
						matchedLabels := false
						for lk, _ := range na.Labels {
							if lk == nodeActionLabelNotebookId {
								matchedLabels = true
							}
						}
						if !matchedLabels {
							return false
						}
						if na.Status.State != nav1.ActionCompletedState {
							return false
						}
						return true
					},
					Handler: cache.ResourceEventHandlerFuncs{
						AddFunc: func(obj interface{}) {
							na := obj.(*nav1.NodeAction)
							s.handleNodeActions(na)
						},
						UpdateFunc: func(oldObj, newObj interface{}) {
							na := newObj.(*nav1.NodeAction)
							s.handleNodeActions(na)
						},
					},
				},
				0,
			)
		}()
	})

}

func (s *developService) handleNodeActions(na *nav1.NodeAction) {
	notebookId := na.Labels[nodeActionLabelNotebookId]
	imageId := na.Labels[nodeActionLabelImageId]

	var remark string = "{\"state\":\"%s\",\"reason\":\"%s\",\"imageId\":\"%s\"}"
	actionStatus := na.Status.Actions
	var commandStatus *nav1.CommandStatus
	for _, s := range actionStatus {
		if s.Name == "docker.commitAndPush" {
			commandStatus = s
			break
		}
	}
	if commandStatus == nil {
		return
	}

	var imageStatus api.ImageStatus
	switch commandStatus.Result {
	case nav1.CommandFailedResult:
		imageStatus = api.ImageStatus_IMAGE_STATUS_MADE_FAILED
	case nav1.CommandSucceedResult:
		imageStatus = api.ImageStatus_IMAGE_STATUS_MADE
	default:
		imageStatus = api.ImageStatus_IMAGE_STATUS_MADE_FAILED
	}

	ctx := context.Background()
	_, err := s.imageService.UpdateImage(ctx, &api.UpdateImageRequest{
		ImageId:     imageId,
		ImageStatus: imageStatus,
	})
	if err != nil {
		s.log.Errorw(ctx, err)
	}

	if err := s.data.Cluster.DeleteNodeAction(ctx, na.Namespace, na.Name); err != nil {
		s.log.Error(ctx, err)
	}

	err = s.data.DevelopDao.CreateNotebookEventRecord(ctx, &model.NotebookEventRecord{
		Time:       time.Now(),
		NotebookId: notebookId,
		Type:       commapi.NotebookEventRecordType_SAVE,
		Remark:     fmt.Sprintf(remark, commandStatus.Result, commandStatus.Reason, imageId),
	})
	if err != nil { // 插入事件记录出错只打印
		s.log.Error(ctx, "save notebook event record error:", err)
	}
}
