package modeldeploy

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/common/log"
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

	return s, nil
}

// 部署模型服务
func (s *modelDeployService) DeployModel(ctx context.Context, req *api.DepRequest) (*api.DepReply, error) {
   return nil,nil
}

//停止模型服务
func (s *modelDeployService) StopDepModel(ctx context.Context, req *api.StopDepRequest) (*api.StopDepReply, error) {
	return nil,nil
}

//删除模型服务
func (s *modelDeployService) DeleteDepModel(ctx context.Context, req *api.DeleteDepRequest) (*api.DeleteDepReply, error) {
	return nil,nil
}

//获取模型服务详情
func (s *modelDeployService) GetModelDepInfo(ctx context.Context, req *api.DepInfoRequest) (*api.DepInfoReply, error) {
	return nil,nil
}

//获取模型服务列表
func (s *modelDeployService) ListDepModel(ctx context.Context, req *api.DepListRequest) (*api.DepListReply, error) {
	return nil,nil
}

//获取模型事件
func (s *modelDeployService) ListDepEvent(ctx context.Context, req *api.DepEventListRequest) (*api.DepEventListReply, error) {
	return nil,nil
}

