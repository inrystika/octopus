package service

import (
	"context"
	pb "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	innterapi "server/base-server/api/v1"
	"server/common/log"
)

type WorkspaceService struct {
	pb.UnimplementedWorkspaceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewWorkspaceService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) pb.WorkspaceServer {
	return &WorkspaceService{
		conf: conf,
		log:  log.NewHelper("WorkspaceService", logger),
		data: data,
	}
}

func (s *WorkspaceService) CreateWorkspace(ctx context.Context, req *pb.CreateWorkspaceRequest) (*pb.CreateWorkspaceReply, error) {
	_, err := s.data.ResourcePoolClient.GetResourcePool(ctx, &innterapi.GetResourcePoolRequest{Id: req.ResourcePoolId})
	if err != nil {
		return nil, err
	}

	result, err := s.data.WorkspaceClient.CreateWorkspace(ctx, &innterapi.CreateWorkspaceRequest{
		Name:           req.Name,
		UserIds:        req.UserIds,
		ResourcePoolId: req.ResourcePoolId,
	})
	if err != nil {
		return nil, err
	}

	_, err = s.data.BillingClient.CreateBillingOwner(ctx, &innterapi.CreateBillingOwnerRequest{
		OwnerId:   result.Id,
		OwnerType: innterapi.BillingOwnerType_BOT_SPACE,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateWorkspaceReply{
		Id: result.Id,
	}, nil
}

func (s *WorkspaceService) UpdateWorkspace(ctx context.Context, req *pb.UpdateWorkspaceRequest) (*pb.UpdateWorkspaceReply, error) {
	_, err := s.data.ResourcePoolClient.GetResourcePool(ctx, &innterapi.GetResourcePoolRequest{Id: req.ResourcePoolId})
	if err != nil {
		return nil, err
	}

	reply, err := s.data.WorkspaceClient.UpdateWorkspace(ctx, &innterapi.UpdateWorkspaceRequest{
		WorkspaceId:    req.WorkspaceId,
		Name:           req.Name,
		UserIds:        req.UserIds,
		ResourcePoolId: req.ResourcePoolId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateWorkspaceReply{UpdatedAt: reply.UpdatedAt}, nil
}

func (s *WorkspaceService) DeleteWorkspace(ctx context.Context, req *pb.DeleteWorkspaceRequest) (*pb.DeleteWorkspaceReply, error) {
	reply, err := s.data.WorkspaceClient.DeleteWorkspace(ctx, &innterapi.DeleteWorkspaceRequest{
		WorkspaceId: req.WorkspaceId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteWorkspaceReply{DeletedAt: reply.DeletedAt}, nil
}

func (s *WorkspaceService) ListWorkspace(ctx context.Context, req *pb.ListWorkspaceRequest) (*pb.ListWorkspaceReply, error) {
	result, err := s.data.WorkspaceClient.ListWorkspace(ctx, &innterapi.ListWorkspaceRequest{
		Name:      req.Name,
		SearchKey: req.SearchKey,
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		SortBy:    req.SortBy,
		OrderBy:   req.OrderBy,
	})
	if err != nil {
		return nil, err
	}

	workspaces := []*pb.WorkspaceItem{}
	for _, w := range result.Workspaces {
		workspaces = append(workspaces, &pb.WorkspaceItem{
			Id:             w.Id,
			Name:           w.Name,
			ResourcePoolId: w.ResourcePoolId,
			CreatedAt:      w.CreatedAt,
			UpdatedAt:      w.UpdatedAt,
		})
	}

	return &pb.ListWorkspaceReply{
		TotalSize:  result.TotalSize,
		Workspaces: workspaces,
	}, nil
}

func (s *WorkspaceService) ListWorkspaces(ctx context.Context, req *pb.ListWorkspacesRequest) (*pb.ListWorkspacesReply, error) {
	result, err := s.data.WorkspaceClient.ListWorkspaces(ctx, &innterapi.ListWorkspacesRequest{})
	if err != nil {
		return nil, err
	}

	workspaces := []*pb.WorkspaceItem{}
	for _, w := range result.Workspaces {
		workspaces = append(workspaces, &pb.WorkspaceItem{
			Id:             w.Id,
			Name:           w.Name,
			ResourcePoolId: w.ResourcePoolId,
		})
	}

	return &pb.ListWorkspacesReply{
		Workspaces: workspaces,
	}, nil
}

func (s *WorkspaceService) GetWorkspace(ctx context.Context, req *pb.GetWorkspaceRequest) (*pb.GetWorkspaceReply, error) {
	result, err := s.data.WorkspaceClient.GetWorkspace(ctx, &innterapi.GetWorkspaceRequest{
		WorkspaceId: req.WorkspaceId,
	})
	if err != nil {
		return nil, err
	}
	if result.Workspace == nil {
		return &pb.GetWorkspaceReply{}, nil
	}

	workspace := &pb.WorkspaceItem{
		Id:             result.Workspace.Id,
		Name:           result.Workspace.Name,
		ResourcePoolId: result.Workspace.ResourcePoolId,
		CreatedAt:      result.Workspace.CreatedAt,
		UpdatedAt:      result.Workspace.UpdatedAt,
	}

	users := []*pb.WorkspaceUser{}
	for _, u := range result.Users {
		users = append(users, &pb.WorkspaceUser{
			Id:       u.Id,
			FullName: u.FullName,
			Email:    u.Email,
		})
	}
	return &pb.GetWorkspaceReply{
		Workspace: workspace,
		Users:     users,
	}, nil
}
