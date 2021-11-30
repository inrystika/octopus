package modeldeploy

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"server/base-server/internal/data/pipeline"
	"server/common/log"
	"server/common/utils"
	"time"

	"github.com/jinzhu/copier"
	seldonv2 "github.com/seldonio/seldon-core/operator/apis/machinelearning.seldon.io/v1alpha2"
)

const (
	STATE_AVAILABLE = "Available"
	STATE_CREATING  = "Creating"
	STATE_FAILED    = "Failed"
)

type modelDeployService struct {
	api.UnimplementedModelDeployServiceServer
	conf                *conf.Bootstrap
	log                 *log.Helper
	data                *data.Data
	modelService        api.ModelServer
	workspaceService    api.WorkspaceServer
	resourceSpecService api.ResourceSpecServiceServer
	resourceService     api.ResourceServiceServer
	resourcePoolService api.ResourcePoolServiceServer
	billingService      api.BillingServiceServer
}

type ModelDeployService interface {
	api.ModelDeployServiceServer
}

func NewModelDeployService(conf *conf.Bootstrap, logger log.Logger, data *data.Data,
	workspaceService api.WorkspaceServer, modelService api.ModelServer,
	resourceSpecService api.ResourceSpecServiceServer, resourceService api.ResourceServiceServer,
	resourcePoolService api.ResourcePoolServiceServer, billingService api.BillingServiceServer) (ModelDeployService, error) {
	log := log.NewHelper("ModelDeployService", logger)

	s := &modelDeployService{
		conf:                conf,
		log:                 log,
		data:                data,
		workspaceService:    workspaceService,
		modelService:        modelService,
		resourceSpecService: resourceSpecService,
		resourceService:     resourceService,
		resourcePoolService: resourcePoolService,
		billingService:      billingService,
	}

	//s.modelDepBilling(context.Background())
	ctx := context.Background()
	s.data.Cluster.RegisterDeploymentInformerCallback(ctx,
		func(obj interface{}) {},
		func(old, obj interface{}) {
			objSeldon, ok := obj.(*seldonv2.SeldonDeployment)
			if !ok {
				return
			}
			deployService, err := s.data.ModelDeployDao.GetModelDeployService(ctx, objSeldon.Name)
			if err != nil {
				return
			}
			newState := string(objSeldon.Status.State)
			if newState == deployService.Status {
				return
			}
			update := &model.ModelDeploy{
				Id:     objSeldon.Name,
				Status: newState,
			}
			now := time.Now()
			if newState == STATE_AVAILABLE {
				update.StartedAt = &now
			} else if newState == STATE_FAILED {
				update.CompletedAt = &now
			}
			s.data.ModelDeployDao.UpdateModelDeployService(ctx, update)
		},
		func(obj interface{}) {},
	)

	return s, nil
}

// 部署模型服务
func (s *modelDeployService) DeployModel(ctx context.Context, req *api.DepRequest) (*api.DepReply, error) {
	modelServiceId := utils.GetUUIDStartWithAlphabetic() //k8s service首字母不允许数字 为方便 uuid处理一下
	//check 任务是否重名，联合索引。同名且未软删除，则报错。
	_, err := s.data.ModelDeployDao.GetModelDeployServiceByName(ctx, req.Name, req.UserId, req.WorkspaceId)
	if err != nil {
		return nil, err
	}

	deployJob := &model.ModelDeploy{}
	err = copier.Copy(deployJob, req)
	if err != nil {
		return nil, err
	}
	deployJob.Id = modelServiceId
	deployJob.Status = pipeline.PREPARING
	//各类参数校验
	startJobInfo, err := s.checkParam(ctx, deployJob)
	if err != nil {
		return nil, err
	}
	//submit job
	closeFunc, err := s.submitDeployJob(ctx, deployJob, startJobInfo)
	defer func() { //如果出错 重要的资源需要删除
		if err != nil {
			_ = closeFunc(ctx)
		}
	}()
	if err != nil {
		return nil, err
	}
	//create recorde
	err = s.data.ModelDeployDao.CreateModelDeployService(ctx, deployJob)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

type deployInfo struct {
}

type closeFunc func(ctx context.Context) error

func (s *modelDeployService) submitDeployJob(ctx context.Context, modelDeploy *model.ModelDeploy, deployInfo *deployInfo) (closeFunc, error) {
	var err error
	closes := make([]closeFunc, 0)
	resFunc := func(ctx context.Context) error {
		var err2 error
		for i := len(closes) - 1; i >= 0; i-- {
			err1 := closes[i](ctx)
			if err1 != nil {
				err2 = err1
				s.log.Errorf(ctx, "err: %s", err1)
			}
		}

		return err2
	}

	defer func() {
		if err != nil {
			_ = resFunc(ctx)
		}
	}()
	return nil, nil
}

func (s *modelDeployService) checkParam(ctx context.Context, modelDeploy *model.ModelDeploy) (*deployInfo, error) {
	return nil, nil
}

//停止模型服务
func (s *modelDeployService) StopDepModel(ctx context.Context, req *api.StopDepRequest) (*api.StopDepReply, error) {
	return nil, nil
}

//删除模型服务
func (s *modelDeployService) DeleteDepModel(ctx context.Context, req *api.DeleteDepRequest) (*api.DeleteDepReply, error) {
	return nil, nil
}

//获取模型服务详情
func (s *modelDeployService) GetModelDepInfo(ctx context.Context, req *api.DepInfoRequest) (*api.DepInfoReply, error) {

	return nil, nil
}

//获取模型服务列表
func (s *modelDeployService) ListDepModel(ctx context.Context, req *api.DepListRequest) (*api.DepListReply, error) {
	return nil, nil
}

//获取模型事件
func (s *modelDeployService) ListDepEvent(ctx context.Context, req *api.DepEventListRequest) (*api.DepEventListReply, error) {
	return nil, nil
}
