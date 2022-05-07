package service

import (
	"context"
	innterapi "server/base-server/api/v1"
	"server/common/constant"
	commctx "server/common/context"
	"server/common/errors"
	"server/common/log"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"
)

type UserService struct {
	api.UnimplementedUserServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewUserService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.UserServer {
	return &UserService{
		conf: conf,
		log:  log.NewHelper("UserService", logger),
		data: data,
	}
}

func (s *UserService) GetUserInfo(ctx context.Context, req *api.GetUserInfoRequest) (*api.GetUserInfoReply, error) {
	userId := commctx.UserIdFromContext(ctx)
	if userId == "" {
		return nil, errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	reply, err := s.data.UserClient.FindUser(ctx, &innterapi.FindUserRequest{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}
	if reply.User == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserAccountNotExisted)
	}

	return &api.GetUserInfoReply{
		User: &api.UserItem{
			Id:        reply.User.Id,
			CreatedAt: reply.User.CreatedAt,
			UpdatedAt: reply.User.UpdatedAt,
			FullName:  reply.User.FullName,
			Email:     reply.User.Email,
			Phone:     reply.User.Phone,
			Gender:    int32(reply.User.Gender),
			Status:    int32(reply.User.Status),
		},
	}, nil
}

func (s *UserService) ListUserWorkspaces(ctx context.Context, req *api.ListUserWorkspacesRequest) (*api.ListUserWorkspacesReply, error) {
	userId := commctx.UserIdFromContext(ctx)
	if userId == "" {
		return nil, errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	if userId != req.UserId {
		return nil, errors.Errorf(nil, errors.ErrorUserIdNotRight)
	}

	result, err := s.data.WorkspaceClient.ListUserWorkspaces(ctx, &innterapi.ListUserWorkspacesRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	workspaces := make([]*api.WorkspaceItem, len(result.Workspaces))
	for idx, w := range result.Workspaces {
		workspaces[idx] = &api.WorkspaceItem{
			Id:        w.Id,
			Name:      w.Name,
			CreatedAt: w.CreatedAt,
			UpdatedAt: w.UpdatedAt,
		}
	}

	// add default workspace
	workspaces = append(workspaces, &api.WorkspaceItem{
		Id:   constant.SYSTEM_WORKSPACE_DEFAULT,
		Name: "",
	})

	return &api.ListUserWorkspacesReply{
		Workspaces: workspaces,
	}, nil
}

func (s *UserService) GetUserConfig(ctx context.Context, req *api.GetUserConfigRequest) (*api.GetUserConfigReply, error) {
	userId, _ := commctx.UserIdAndSpaceIdFromContext(ctx)
	reply, err := s.data.UserClient.GetUserConfig(ctx, &innterapi.GetUserConfigRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return &api.GetUserConfigReply{Config: reply.Config}, nil
}

func (s *UserService) UpdateUserFtpAccount(ctx context.Context, req *api.UpdateUserFtpAccountRequest) (*api.UpdateUserFtpAccountReply, error) {
	userId := commctx.UserIdFromContext(ctx)

	_, err := s.data.UserClient.UpdateUserFtpAccount(ctx, &innterapi.UpdateUserFtpAccountRequest{
		FtpPassword: req.FtpPassword,
		FtpUserName: req.FtpUserName,
		UserId:      userId,
	})

	if err != nil {
		return nil, err
	}

	return &api.UpdateUserFtpAccountReply{}, nil
}
