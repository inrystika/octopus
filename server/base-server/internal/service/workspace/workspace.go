package workspace

import (
	"context"
	api "server/base-server/api/v1"
	cm "server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"server/base-server/internal/data/registry"
	"server/common/constant"
	"server/common/errors"
	"server/common/snowflake"
	"server/common/utils/collections/set"
	"strconv"
	"time"

	"server/common/log"
)

type WorkspaceService struct {
	api.UnimplementedWorkspaceServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewWorkspaceService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.WorkspaceServiceServer {
	service := &WorkspaceService{
		conf: conf,
		log:  log.NewHelper("WorkspaceService", logger),
		data: data,
	}

	// create default project
	if err := data.Registry.CreateProject(&registry.ProjectReq{ProjectName: constant.SYSTEM_WORKSPACE_DEFAULT, Public: true}); err != nil {
		if !errors.IsError(errors.ErrorHarborProjectExists, err) {
			panic(err)
		}
	}
	// create global project for pre image
	if err := data.Registry.CreateProject(&registry.ProjectReq{ProjectName: cm.PREAB_FOLDER, Public: true}); err != nil {
		if !errors.IsError(errors.ErrorHarborProjectExists, err) {
			panic(err)
		}
	}
	return service
}

func (s *WorkspaceService) FindWorkspace(ctx context.Context, req *api.FindWorkspaceRequest) (*api.FindWorkspaceReply, error) {
	workspaceModel, err := s.data.WorkspaceDao.Find(ctx, &model.WorkspaceQuery{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if workspaceModel == nil {
		return &api.FindWorkspaceReply{}, nil
	}

	return &api.FindWorkspaceReply{
		Workspace: &api.WorkspaceItem{
			Id:             workspaceModel.Id,
			Name:           workspaceModel.Name,
			ResourcePoolId: workspaceModel.RPoolId,
			CreatedAt:      workspaceModel.CreatedAt.Unix(),
			UpdatedAt:      workspaceModel.UpdatedAt.Unix(),
		},
	}, nil
}

func (s *WorkspaceService) ListWorkspace(ctx context.Context, req *api.ListWorkspaceRequest) (*api.ListWorkspaceReply, error) {
	workspaceModels, err := s.data.WorkspaceDao.List(ctx, &model.WorkspaceList{
		SearchKey: req.SearchKey,
		SortBy:    req.SortBy,
		OrderBy:   req.OrderBy,
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		Name:      req.Name,
	})
	if err != nil {
		return nil, err
	}

	workspaceCount, err := s.data.WorkspaceDao.Count(ctx, &model.WorkspaceList{
		SearchKey: req.SearchKey,
		Name:      req.Name,
	})
	if err != nil {
		return nil, err
	}

	workspaces := make([]*api.WorkspaceItem, 0)
	for _, ws := range workspaceModels {
		item := &api.WorkspaceItem{
			Id:             ws.Id,
			Name:           ws.Name,
			ResourcePoolId: ws.RPoolId,
			CreatedAt:      ws.CreatedAt.Unix(),
			UpdatedAt:      ws.UpdatedAt.Unix(),
		}
		workspaces = append(workspaces, item)
	}

	return &api.ListWorkspaceReply{
		TotalSize:  workspaceCount,
		Workspaces: workspaces,
	}, nil
}

func (s *WorkspaceService) ListWorkspaces(ctx context.Context, req *api.ListWorkspacesRequest) (*api.ListWorkspacesReply, error) {
	workspaceModels, err := s.data.WorkspaceDao.List(ctx, &model.WorkspaceList{})
	if err != nil {
		return nil, err
	}

	workspaces := make([]*api.WorkspaceItem, 0)
	for _, ws := range workspaceModels {
		item := &api.WorkspaceItem{
			Id:             ws.Id,
			Name:           ws.Name,
			ResourcePoolId: ws.RPoolId,
			CreatedAt:      ws.CreatedAt.Unix(),
			UpdatedAt:      ws.UpdatedAt.Unix(),
		}
		workspaces = append(workspaces, item)
	}

	return &api.ListWorkspacesReply{
		Workspaces: workspaces,
	}, nil
}

func (s *WorkspaceService) GetWorkspace(ctx context.Context, req *api.GetWorkspaceRequest) (*api.GetWorkspaceReply, error) {
	workspaceModel, err := s.data.WorkspaceDao.Find(ctx, &model.WorkspaceQuery{
		Id: req.WorkspaceId,
	})
	if err != nil {
		return nil, err
	}
	if workspaceModel == nil {
		return &api.GetWorkspaceReply{
			Workspace: nil,
		}, nil
	}

	// load user info
	users, err := s.data.WorkspaceDao.ListWorkspaceUser(ctx, &model.WorkspaceUserList{WorkspaceId: req.WorkspaceId})
	if err != nil {
		return nil, err
	}

	workspaceUsers := []*api.WorkspaceUser{}
	for _, user := range users {
		workspaceUsers = append(workspaceUsers, &api.WorkspaceUser{
			Id:       user.Id,
			FullName: user.FullName,
			Email:    user.Email,
		})
	}

	reply := &api.GetWorkspaceReply{
		Workspace: &api.WorkspaceItem{
			Id:             workspaceModel.Id,
			Name:           workspaceModel.Name,
			ResourcePoolId: workspaceModel.RPoolId,
			CreatedAt:      workspaceModel.CreatedAt.Unix(),
			UpdatedAt:      workspaceModel.UpdatedAt.Unix(),
		},
		Users: workspaceUsers,
	}

	return reply, nil
}

func (s *WorkspaceService) CreateWorkspace(ctx context.Context, req *api.CreateWorkspaceRequest) (*api.CreateWorkspaceReply, error) {
	wq := model.WorkspaceList{
		RPoolId: req.ResourcePoolId,
	}

	existed, err := s.data.WorkspaceDao.Count(ctx, &wq)
	if err != nil {
		return nil, err
	}
	if existed > 0 {
		return nil, errors.Errorf(nil, errors.ErrorWorkSpaceResourcePoolBound)
	}
	existed, err = s.data.WorkspaceDao.Count(ctx, &model.WorkspaceList{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	if existed > 0 {
		return nil, errors.Errorf(nil, errors.ErrorWorkSpaceExisted)
	}

	wa := model.WorkspaceAdd{
		Id:      strconv.FormatUint(snowflake.NextUID(), 10),
		Name:    req.Name,
		RPoolId: req.ResourcePoolId,
	}
	workspace, err := s.data.WorkspaceDao.Add(ctx, &wa)
	if err != nil {
		return nil, err
	}

	// add users to workspace
	if len(req.UserIds) > 0 {
		users, err := s.data.UserDao.ListIn(ctx, &model.UserListIn{Ids: req.UserIds})
		if err != nil {
			return nil, err
		}
		userIds := []string{}
		for _, u := range users {
			userIds = append(userIds, u.Id)
		}
		if len(userIds) > 0 {
			batchUsersAppend := model.WorkspaceUserBatchAdd{UserIds: userIds}
			batchUsersAppend.WorkspaceId = wa.Id
			if err = s.data.WorkspaceDao.AddWorkspaceUsers(ctx, &batchUsersAppend); err != nil {
				return nil, err
			}
		}
	}

	// create registry project, like: harbor workspace.Id
	if err = s.data.Registry.CreateProject(&registry.ProjectReq{ProjectName: workspace.Id, Public: true}); err != nil {
		return nil, err
	}

	reply := &api.CreateWorkspaceReply{
		Id: workspace.Id,
	}

	return reply, nil
}

func (s *WorkspaceService) UpdateWorkspace(ctx context.Context, req *api.UpdateWorkspaceRequest) (*api.UpdateWorkspaceReply, error) {
	target, err := s.data.WorkspaceDao.Find(ctx, &model.WorkspaceQuery{Id: req.WorkspaceId})
	if err != nil {
		return nil, err
	}
	if target == nil {
		return nil, errors.Errorf(nil, errors.ErrorWorkSpaceNotExist)
	}

	if target.Name != req.Name || target.RPoolId != req.ResourcePoolId {
		if target, err = s.data.WorkspaceDao.Update(ctx, &model.WorkspaceUpdate{Id: req.WorkspaceId, Name: req.Name, RPoolId: req.ResourcePoolId}); err != nil {
			return nil, err
		}
	}

	// sync workspace users
	users, err := s.data.WorkspaceDao.ListWorkspaceUser(ctx, &model.WorkspaceUserList{WorkspaceId: target.Id})
	if err != nil {
		return nil, err
	}

	destUserIds := []string{}
	for _, user := range users {
		destUserIds = append(destUserIds, user.Id)
	}
	targetUserIdsSet := set.NewStrings(destUserIds...)
	sourceUserIdsSet := set.NewStrings(req.UserIds...)
	userIdsAdd := sourceUserIdsSet.Difference(targetUserIdsSet).Values()
	userIdsDel := targetUserIdsSet.Difference(sourceUserIdsSet).Values()
	if len(userIdsAdd) > 0 {
		users, err := s.data.UserDao.ListIn(ctx, &model.UserListIn{Ids: userIdsAdd})
		if err != nil {
			return nil, err
		}
		userIds := []string{}
		for _, u := range users {
			userIds = append(userIds, u.Id)
		}
		if err = s.data.WorkspaceDao.AddWorkspaceUsers(ctx, &model.WorkspaceUserBatchAdd{WorkspaceId: target.Id, UserIds: userIds}); err != nil {
			return nil, err
		}
	}
	if len(userIdsDel) > 0 {
		if err = s.data.WorkspaceDao.DelWorkspaceUsers(ctx, &model.WorkspaceUserBatchDel{WorkspaceId: target.Id, UserIds: userIdsDel}); err != nil {
			return nil, err
		}
	}
	return &api.UpdateWorkspaceReply{UpdatedAt: time.Now().Unix()}, nil
}

func (s *WorkspaceService) DeleteWorkspace(ctx context.Context, req *api.DeleteWorkspaceRequest) (*api.DeleteWorkspaceReply, error) {
	if w, err := s.data.WorkspaceDao.Delete(ctx, &model.WorkspaceDelete{Id: req.WorkspaceId}); err != nil {
		return nil, err
	} else {
		if w == nil {
			return nil, errors.Errorf(nil, errors.ErrorWorkSpaceNotExist)
		}
		if err := s.data.WorkspaceDao.DelWorkspaceUsers(ctx, &model.WorkspaceUserBatchDel{WorkspaceId: req.WorkspaceId}); err != nil {
			return nil, err
		}
		return &api.DeleteWorkspaceReply{DeletedAt: time.Now().Unix()}, nil
	}
}

func (s *WorkspaceService) ListUserWorkspaces(ctx context.Context, req *api.ListUserWorkspacesRequest) (*api.ListUserWorkspacesReply, error) {
	workspaces, err := s.data.WorkspaceDao.ListUserWorkspace(ctx, &model.UserWorkspaceList{UserId: req.UserId})
	if err != nil {
		return nil, err
	}

	workspaceItems := []*api.WorkspaceItem{}
	for _, wi := range workspaces {
		workspaceItems = append(workspaceItems, &api.WorkspaceItem{
			Id:             wi.Id,
			Name:           wi.Name,
			ResourcePoolId: wi.RPoolId,
			CreatedAt:      wi.CreatedAt.Unix(),
			UpdatedAt:      wi.UpdatedAt.Unix(),
		})
	}
	return &api.ListUserWorkspacesReply{
		Workspaces: workspaceItems,
	}, nil
}

func (s *WorkspaceService) ListWorkspaceInCond(ctx context.Context, req *api.ListWorkspaceInCondRequest) (*api.ListWorkspaceInCondReply, error) {
	workspaces, err := s.data.WorkspaceDao.ListIn(ctx, &model.WorkspaceListIn{Ids: req.Ids})
	if err != nil {
		return nil, err
	}

	workspaceItems := make([]*api.WorkspaceItem, len(workspaces))
	for idx, workspace := range workspaces {
		item := &api.WorkspaceItem{
			Id:             workspace.Id,
			Name:           workspace.Name,
			ResourcePoolId: workspace.RPoolId,
			CreatedAt:      workspace.CreatedAt.Unix(),
			UpdatedAt:      workspace.UpdatedAt.Unix(),
		}
		workspaceItems[idx] = item
	}
	return &api.ListWorkspaceInCondReply{
		Workspaces: workspaceItems,
	}, nil
}
