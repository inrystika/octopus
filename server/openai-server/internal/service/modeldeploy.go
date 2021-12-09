package service

import (
	"context"
	innerapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"
	"server/common/session"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"

	"github.com/jinzhu/copier"
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
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}
	//查询任务是否存在及用户是否一致
	depInfo, err := s.data.ModelDeployClient.GetModelDepInfo(ctx, &innerapi.DepInfoRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	if depInfo.DepInfo.UserId != session.UserId {
		return nil, errors.Errorf(nil, errors.ErrorNotAuthorized)
	}
	innerReq := &innerapi.StopDepRequest{
		Id:        req.Id,
		Operation: "user stop job",
	}
	reply, err := s.data.ModelDeployClient.StopDepModel(ctx, innerReq)
	if err != nil {
		return nil, err
	}
	return &api.StopDepReply{StoppedAt: reply.StoppedAt}, nil
}

//删除模型服务
func (s *ModelDeployService) DeleteDepModel(ctx context.Context, req *api.DeleteDepRequest) (*api.DeleteDepReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	err := s.checkPermission(ctx, req.JobIds, session.UserId)
	if err != nil {
		return nil, err
	}

	innerReq := &innerapi.DeleteDepRequest{UserId: session.UserId, JobIds: req.JobIds}
	reply, err := s.data.ModelDeployClient.DeleteDepModel(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.DeleteDepReply{DeletedAt: reply.DeletedAt}, nil
}

func (s *ModelDeployService) checkPermission(ctx context.Context, serviceIds []string, userId string) error {
	for _, jobId := range serviceIds {
		reply, err := s.data.ModelDeployClient.GetModelDepInfo(ctx, &innerapi.DepInfoRequest{Id: jobId})
		if err != nil {
			return err
		}

		if reply.DepInfo.UserId != userId {
			return errors.Errorf(nil, errors.ErrorNotAuthorized)
		}
	}
	return nil
}

// 模型服务调用
func (s *ModelDeployService) ModelServiceInfer(ctx context.Context, req *api.ServiceRequest) (*api.ServiceReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.ServiceRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.UserId = session.UserId
	innerReq.WorkspaceId = session.GetWorkspace()

	innerReply, err := s.data.ModelDeployClient.ModelServiceInfer(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.ServiceReply{
		Response: innerReply.Response,
	}, nil
}

// 获取模型服务详情
func (s *ModelDeployService) GetModelDepInfo(ctx context.Context, req *api.DepInfoRequest) (*api.DepInfoReply, error) {

	innerDepInfo, err := s.data.ModelDeployClient.GetModelDepInfo(ctx, &innerapi.DepInfoRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	depInfo := &api.DepInfoReply{}
	err = copier.Copy(depInfo, innerDepInfo)
	if err != nil {
		return nil, err
	}

	return depInfo, nil
}

// 模型服务列表
func (s *ModelDeployService) ListDepModel(ctx context.Context, req *api.DepListRequest) (*api.DepListReply, error) {
	innerReq := &innerapi.DepListRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.ModelDeployClient.ListDepModel(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.DepListReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	if reply.DepInfos == nil {
		reply := &api.DepListReply{
			TotalSize: 0,
			DepInfos:  nil,
		}
		return reply, nil
	}

	return reply, nil
}

// 模型服务事件列表
func (s *ModelDeployService) ListDepEvent(ctx context.Context, req *api.DepEventListRequest) (*api.DepEventListReply, error) {
	innerReq := &innerapi.DepEventListRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.ModelDeployClient.ListDepEvent(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.DepEventListReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
