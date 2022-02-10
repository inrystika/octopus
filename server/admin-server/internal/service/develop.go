package service

import (
	"context"
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	innerapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"
	"server/common/utils"

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

func (s *DevelopService) StopNotebook(ctx context.Context, req *api.StopNotebookRequest) (*api.StopNotebookReply, error) {
	reply, err := s.data.DevelopClient.StopNotebook(ctx, &innerapi.StopNotebookRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &api.StopNotebookReply{Id: reply.Id}, nil
}

func (s *DevelopService) ListNotebook(ctx context.Context, req *api.ListNotebookRequest) (*api.ListNotebookReply, error) {
	innerReq := &innerapi.ListNotebookRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.DevelopClient.ListNotebook(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListNotebookReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	err = s.assignValue(ctx, reply.Notebooks)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *DevelopService) assignValue(ctx context.Context, notebooks []*api.Notebook) error {
	if len(notebooks) > 0 {
		userIdMap := map[string]interface{}{}
		spaceIdMap := map[string]interface{}{}
		for _, i := range notebooks {
			userIdMap[i.UserId] = true
			spaceIdMap[i.WorkspaceId] = true
		}

		users, err := s.data.UserClient.ListUserInCond(ctx, &innerapi.ListUserInCondRequest{Ids: utils.MapKeyToSlice(userIdMap)})
		if err != nil {
			return err
		}
		userMap := map[string]*innerapi.UserItem{}
		for _, i := range users.Users {
			userMap[i.Id] = i
		}

		spaces, err := s.data.WorkspaceClient.ListWorkspaceInCond(ctx, &innerapi.ListWorkspaceInCondRequest{
			Ids: utils.MapKeyToSlice(spaceIdMap),
		})
		if err != nil {
			return err
		}
		spaceMap := map[string]*innerapi.WorkspaceItem{}
		for _, i := range spaces.Workspaces {
			spaceMap[i.Id] = i
		}

		for _, i := range notebooks {
			if v, ok := userMap[i.UserId]; ok {
				i.UserName = v.FullName
				i.Email = v.Email
			}

			if v, ok := spaceMap[i.WorkspaceId]; ok {
				i.WorkspaceName = v.Name
			}
		}
	}

	return nil
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

func (s *DevelopService) ListNotebookEventRecord(ctx context.Context, req *api.ListNotebookEventRecordRequest) (*api.ListNotebookEventRecordReply, error) {
	innerReq := &innerapi.ListNotebookEventRecordRequest{}
	err := copier.Copy(innerReq, req)
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
