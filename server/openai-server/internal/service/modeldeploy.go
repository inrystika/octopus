package service

import (
	"context"
	"github.com/jinzhu/copier"
	innerapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"
	"server/common/session"
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
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.DepRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.UserId = session.UserId
	innerReq.WorkspaceId = session.GetWorkspace()

	innerReply, err := s.data.ModelDeployClient.DeployModel(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.DepReply{
		ServiceId:  innerReply.ServiceId,
		ServiceUrl: innerReply.ServiceUrl,
		Message:    innerReply.Message,
	}, nil
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
