package service

import (
	"context"
	"server/common/log"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"
)

type ModelDeployService struct {
	api.UnimplementedModelDeployServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewModelDeployService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.ModelDeployServiceServer {
	return &ModelDeployService{
		conf: conf,
		log:  log.NewHelper("ModelDeployService", logger),
		data: data,
	}
}

//创建模型服务
func (s *ModelDeployService) DeployModel(ctx context.Context, req *api.DepRequest) (*api.DepReply, error) {
	return nil,nil
}

// 停止模型服务
func (s *ModelDeployService) StopDepModel(ctx context.Context, req *api.StopDepRequest) (*api.StopDepReply, error) {
	return nil, nil
}

//删除模型服务
func (s *ModelDeployService) DeleteDepModel(ctx context.Context, req *api.DeleteDepRequest) (*api.DeleteDepReply, error) {
	return nil, nil
}

// 获取模型服务详情
func (s *ModelDeployService) GetModelDepInfo(ctx context.Context, req *api.DepInfoRequest) (*api.DepInfoReply, error) {
	return nil, nil
}

// 模型服务列表
func (s *ModelDeployService) ListDepModel(ctx context.Context, req *api.DepListRequest) (*api.DepListReply, error) {
	return nil, nil
}

// 模型服务事件列表
func (s *ModelDeployService) ListDepEvent(ctx context.Context, req *api.DepEventListRequest) (*api.DepEventListReply, error) {
	return nil, nil
}