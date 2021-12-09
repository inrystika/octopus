package modeldeploy

import (
	"context"
	"gonum.org/v1/gonum/floats"
	"k8s.io/apimachinery/pkg/util/wait"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/data/dao/model"
	"server/base-server/internal/data/pipeline"
	"server/common/constant"
	"server/common/leaderleaselock"
	"server/common/utils"
	"strings"
	"time"
)

const (
	taskPageSize = 100
	leaseLock    = "modeldeployleaselock"
)

func (s *modelDeployService) getOwner(modelDeploy *model.ModelDeploy) (string, api.BillingOwnerType) {
	var ownerId string
	var ownerType api.BillingOwnerType

	if modelDeploy.WorkspaceId == constant.SYSTEM_WORKSPACE_DEFAULT {
		ownerId = modelDeploy.UserId
		ownerType = api.BillingOwnerType_BOT_USER
	} else {
		ownerId = modelDeploy.WorkspaceId
		ownerType = api.BillingOwnerType_BOT_SPACE
	}

	return ownerId, ownerType
}

func (s *modelDeployService) calAmount(startAt int64, endedAt int64, price float64) float64 {
	return float64(endedAt-startAt) * price / 3600.0
}


func (s *modelDeployService) modelServiceBilling(ctx context.Context) {

	k8sns := utils.GetEnvOrDefault("K8S_NAMESPACE", "default")
	rdlock := leaderleaselock.NewLeaderLeaselock(k8sns, leaseLock, s.data.Cluster.GetClusterConfig())
	rdlock.RunOrRetryLeaderElection(ctx, func(ctx context.Context) {
		s.log.Infof(ctx, "model deploy service billing service acquire %v", leaseLock)

		go func() {
			BillingPeriodSec := int64(1800) //默认值
			if s.conf.Service.BillingPeriodSec > 0 {
				BillingPeriodSec = s.conf.Service.BillingPeriodSec
			}
			s.log.Infof(ctx, "model deploy service billing service billing period time is %v", BillingPeriodSec)
			wait.Until(func() {
				utils.HandlePanic(ctx, func(i ...interface{}) {
					s.log.Info(ctx, "start model deploy -billing cron service.....")
					for pageIndex := 1; ; pageIndex++ {
						seldonServices, _, err := s.data.ModelDeployDao.GetModelDeployServiceList(ctx, &model.ModelDeployListQuery{
							PageIndex: pageIndex,
							PageSize:  taskPageSize,
							PayStatus: api.BillingPayRecordStatus_BPRS_PAYING,
						})
						if err != nil {
							s.log.Errorf(ctx, "List Model Deploy  Service err: %s", err)
							break
						}

						if len(seldonServices) == 0 {
							s.log.Info(ctx, "There is no model deploy service to pay!")
							break
						}
						//系统升级，或者taskset重装时，会导致任务后续状态丢失。
						//这些任务可能没有启动时间，但状态却是终止的，这些任务不计费,设置计费状态为完成。
						for _, j := range seldonServices {
							if j.StartedAt == nil && pipeline.IsCompletedState(j.Status) {
								j.PayStatus = api.BillingPayRecordStatus_BPRS_PAY_COMPLETED
								err = s.data.ModelDeployDao.UpdateModelDeployService(ctx, j)
								if err != nil {
									s.log.Errorf(ctx, "update ineffective job to completed err: %s", err)
									break
								}
							}
						}

						//更新无效任务状态后再查询、计费
						seldonServices, _, err = s.data.ModelDeployDao.GetModelDeployServiceList(ctx, &model.ModelDeployListQuery{
							PageIndex: pageIndex,
							PageSize:  taskPageSize,
							PayStatus: api.BillingPayRecordStatus_BPRS_PAYING,
						})

						if err != nil {
							s.log.Errorf(ctx, "List TrainJob err: %s", err)
							break
						}

						if len(seldonServices) == 0 {
							s.log.Info(ctx, "There is no trainJob to pay!")
							break
						}

						for _, j := range seldonServices {
							//判断任务是否已经启动。如果没有启动时间，则说明未启动，不计费。
							if j.StartedAt == nil {
								//s.log.Info(ctx, "job "+j.Id+"no need to pay, because it is not started!")
								continue
							}

							payAmount := 0.0
							//job已经启动，则以job的启动时间作为计费起始点。
							payStartAt := j.StartedAt.Unix()
							s.log.Info(ctx, "model deploy bill service try to calculate job pay amount, jobId is: "+j.Id)
							now := time.Now().Unix()
							var payEndAt int64
							var payStatus api.BillingPayRecordStatus
                            if j.Status == STATE_AVAILABLE{
								//仍在running，则取当前系统时间，作为该周期计费终止点。
								payEndAt = now
								payStatus = api.BillingPayRecordStatus_BPRS_PAYING
							}else if strings.EqualFold(j.Status, STATE_FAILED) || strings.EqualFold(j.Status,STATE_STOPPED) {
								payEndAt = j.CompletedAt.Unix()
								payStatus = api.BillingPayRecordStatus_BPRS_PAY_COMPLETED
							}
                            //推理资源单价
                            resourcePrice := j.ResSpecPrice
							payAmount += s.calAmount(payStartAt, payEndAt, resourcePrice)
							ownerId, ownerType := s.getOwner(j)

							payAmount = floats.Round(payAmount, common.BillingPrecision)

							_, err := s.billingService.Pay(ctx, &api.PayRequest{
								OwnerId:   ownerId,
								OwnerType: ownerType,
								Amount:    payAmount,
								BizType:   api.BillingBizType_BBT_ModelDeploy,
								BizId:     j.Id,
								Title:     j.Name,
								StartedAt: payStartAt,
								EndedAt:   payEndAt,
								Status:    payStatus,
							})
							if err != nil {
								s.log.Errorf(ctx, "Pay err: %s", err)
								continue
							}

							startAt := time.Unix(payStartAt, 0)
							endAt := time.Unix(payEndAt, 0)
							err = s.data.ModelDeployDao.UpdateModelDeployService(ctx, &model.ModelDeploy{
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
						}
					}
				})()
			}, time.Duration(BillingPeriodSec)*time.Second, ctx.Done())
		}()
	})
}