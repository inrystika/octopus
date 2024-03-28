package service

import (
	"context"
	innerapi "server/base-server/api/v1"
	commctx "server/common/context"
	"server/common/errors"
	"server/common/log"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"

	"github.com/jinzhu/copier"
)

type DevelopService struct {
	api.UnimplementedDevelopServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewDevelopService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.DevelopServer {
	return &DevelopService{
		conf: conf,
		log:  log.NewHelper("DevelopService", logger),
		data: data,
	}
}

func (s *DevelopService) CreateNotebook(ctx context.Context, req *api.CreateNotebookRequest) (*api.CreateNotebookReply, error) {
	userId, spaceId := commctx.UserIdAndSpaceIdFromContext(ctx)

	innerReq := &innerapi.CreateNotebookRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, err
	}
	innerReq.UserId = userId
	innerReq.WorkspaceId = spaceId

	innerReply, err := s.data.DevelopClient.CreateNotebook(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.CreateNotebookReply{Id: innerReply.Id}, nil
}

func (s *DevelopService) checkPermission(ctx context.Context, notebookId string, userId string) error {
	reply, err := s.data.DevelopClient.GetNotebook(ctx, &innerapi.GetNotebookRequest{Id: notebookId})
	if err != nil {
		return err
	}

	if reply.Notebook.UserId != userId {
		return errors.Errorf(nil, errors.ErrorNotAuthorized)
	}
	return nil
}

func (s *DevelopService) StartNotebook(ctx context.Context, req *api.StartNotebookRequest) (*api.StartNotebookReply, error) {
	userId, _ := commctx.UserIdAndSpaceIdFromContext(ctx)

	err := s.checkPermission(ctx, req.Id, userId)
	if err != nil {
		return nil, err
	}

	innerReq := &innerapi.StartNotebookRequest{}
	err = copier.Copy(innerReq, req)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.DevelopClient.StartNotebook(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.StartNotebookReply{Id: reply.Id}, nil
}

func (s *DevelopService) StopNotebook(ctx context.Context, req *api.StopNotebookRequest) (*api.StopNotebookReply, error) {
	userId, _ := commctx.UserIdAndSpaceIdFromContext(ctx)

	err := s.checkPermission(ctx, req.Id, userId)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.DevelopClient.StopNotebook(ctx, &innerapi.StopNotebookRequest{Id: req.Id, Operation: "user stop job"})
	if err != nil {
		return nil, err
	}

	return &api.StopNotebookReply{Id: reply.Id}, nil
}

func (s *DevelopService) DeleteNotebook(ctx context.Context, req *api.DeleteNotebookRequest) (*api.DeleteNotebookReply, error) {
	userId, _ := commctx.UserIdAndSpaceIdFromContext(ctx)

	err := s.checkPermission(ctx, req.Id, userId)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.DevelopClient.DeleteNotebook(ctx, &innerapi.DeleteNotebookRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &api.DeleteNotebookReply{Id: reply.Id}, nil
}

func (s *DevelopService) ListNotebook(ctx context.Context, req *api.ListNotebookRequest) (*api.ListNotebookReply, error) {
	userId, spaceId := commctx.UserIdAndSpaceIdFromContext(ctx)

	innerReq := &innerapi.ListNotebookRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.UserId = userId
	innerReq.WorkspaceId = spaceId

	innerReply, err := s.data.DevelopClient.ListNotebook(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListNotebookReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *DevelopService) QueryNotebook(ctx context.Context, req *api.QueryNotebookRequest) (*api.QueryNotebookReply, error) {
	innerReq := &innerapi.GetNotebookRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.DevelopClient.GetNotebook(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.QueryNotebookReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Notebook事件列表
func (s *DevelopService) GetNotebookEventList(ctx context.Context, req *api.NotebookEventListRequest) (*api.NotebookEventListReply, error) {
	innerReq := &innerapi.NotebookEventListRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.DevelopClient.GetNotebookEventList(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.NotebookEventListReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// 保存notebook
func (s *DevelopService) SaveNotebook(ctx context.Context, req *api.SaveNotebookRequest) (*api.SaveNotebookReply, error) {
	sReq := &innerapi.SaveNotebookRequest{
		NotebookId:       req.NotebookId,
		TaskName:         req.TaskName,
		ImageName:        req.ImageName,
		ImageVersion:     req.ImageVersion,
		LayerDescription: req.LayerDescription,
	}
	sReply, err := s.data.DevelopClient.SaveNotebook(ctx, sReq)
	if err != nil {
		return nil, err
	}
	return &api.SaveNotebookReply{ImageId: sReply.ImageId}, nil
}

func (s *DevelopService) ListNotebookEventRecord(ctx context.Context, req *api.ListNotebookEventRecordRequest) (*api.ListNotebookEventRecordReply, error) {
	userId, _ := commctx.UserIdAndSpaceIdFromContext(ctx)

	err := s.checkPermission(ctx, req.NotebookId, userId)
	if err != nil {
		return nil, err
	}

	innerReq := &innerapi.ListNotebookEventRecordRequest{}
	err = copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.DevelopClient.ListNotebookEventRecord(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListNotebookEventRecordReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *DevelopService) GetNotebookMetric(ctx context.Context, req *api.GetNotebookMetricRequest) (*api.GetNotebookMetricReply, error) {
	innerReq := &innerapi.GetNotebookMetricRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.DevelopClient.GetNotebookMetric(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.GetNotebookMetricReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
