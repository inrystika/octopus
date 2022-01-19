package develop

import (
	"context"
	"fmt"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/data/dao/model"
	"server/base-server/internal/data/pipeline"
	commapi "server/common/api/v1"
	"server/common/constant"
	"server/common/leaderleaselock"
	"server/common/utils"
	"server/common/utils/collections/set"
	"time"

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

						details, err := s.data.Pipeline.BatchGetJobDetail(ctx, jobIds)
						if err != nil {
							s.log.Errorf(ctx, "BatchGetJobDetail err: %s", err)
							continue
						}
						detailMap := map[string]*pipeline.JobStatusDetail{}
						for _, d := range details.Details {
							detailMap[d.Job.ID] = d
						}

						for _, j := range nbJobs {
							utils.HandlePanic(ctx, func(i ...interface{}) {
								nb := nbMap[j.NotebookId]
								ownerId, ownerType := s.getOwner(nb)

								var payEndAt int64
								var payStatus api.BillingPayRecordStatus
								if pipeline.IsCompletedState(j.Status) {
									payEndAt = j.StoppedAt.Unix()
									payStatus = api.BillingPayRecordStatus_BPRS_PAY_COMPLETED
								} else {
									payEndAt = time.Now().Unix()
									payStatus = api.BillingPayRecordStatus_BPRS_PAYING
								}

								prices := make([]uint32, 0)
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
							})()
						}
					}
				})()
			}, time.Duration(BillingPeriodSec)*time.Second, ctx.Done())
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
	_, err := s.data.ImageDao.Update(ctx, &model.ImageUpdateCond{Id: imageId}, &model.ImageUpdate{
		Status: int32(imageStatus),
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
