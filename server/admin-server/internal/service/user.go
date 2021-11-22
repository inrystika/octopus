package service

import (
	"context"
	pb "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	innerapi "server/base-server/api/v1"
	innterapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"
	"time"

	"github.com/jinzhu/copier"
)

type UserService struct {
	pb.UnimplementedUserServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewUserService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) pb.UserServer {
	return &UserService{
		conf: conf,
		log:  log.NewHelper("UserService", logger),
		data: data,
	}
}

func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	listUserReply, err := s.data.UserClient.ListUser(ctx, &innterapi.ListUserRequest{
		SortBy:    req.SortBy,
		OrderBy:   req.OrderBy,
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		SearchKey: req.SearchKey,
		FullName:  req.FullName,
		Email:     req.Email,
		Phone:     req.Phone,
		Status:    innterapi.UserStatus(req.Status),
	})

	if err != nil {
		return nil, err
	}

	users := make([]*pb.UserItem, len(listUserReply.Users))
	for idx, user := range listUserReply.Users {
		users[idx] = &pb.UserItem{
			Id:        user.Id,
			FullName:  user.FullName,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
			Phone:     user.Phone,
			Gender:    int32(user.Gender),
			Status:    int32(user.Status),
		}
	}

	return &pb.ListUserReply{
		Users:     users,
		TotalSize: listUserReply.TotalSize,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	reply, err := s.data.UserClient.FindUser(ctx, &innterapi.FindUserRequest{
		Id: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	user := reply.User
	if user == nil {
		return &pb.GetUserReply{}, nil
	}

	wreply, err := s.data.WorkspaceClient.ListUserWorkspaces(ctx, &innterapi.ListUserWorkspacesRequest{UserId: user.Id})
	if err != nil {
		return nil, err
	}
	workspaces := []*pb.WorkspaceItem{}
	for _, w := range wreply.Workspaces {
		workspaces = append(workspaces, &pb.WorkspaceItem{
			Id:   w.Id,
			Name: w.Name,
		})
	}

	return &pb.GetUserReply{
		User: &pb.UserItem{
			Id:        user.Id,
			FullName:  user.FullName,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
			Phone:     user.Phone,
			Gender:    int32(user.Gender),
			Status:    int32(user.Status),
		},
		Workspaces: workspaces,
	}, nil
}

func (s *UserService) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserReply, error) {
	addUserReply, err := s.data.UserClient.AddUser(ctx, &innterapi.AddUserRequest{
		Password: req.Password,
		FullName: req.FullName,
		Email:    req.Email,
		Phone:    req.Phone,
		Gender:   innterapi.GenderType(req.Gender),
	})

	if err != nil {
		return nil, err
	}

	user := addUserReply.User

	_, err = s.data.BillingClient.CreateBillingOwner(ctx, &innterapi.CreateBillingOwnerRequest{
		OwnerId:   user.Id,
		OwnerType: innterapi.BillingOwnerType_BOT_USER,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddUserReply{
		Id: user.Id,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	result, err := s.data.UserClient.UpdateUser(ctx, &innterapi.UpdateUserRequest{
		Id:       req.UserId,
		Password: req.User.Password,
		FullName: req.User.FullName,
		Phone:    req.User.Phone,
		Gender:   innterapi.GenderType(req.User.Gender),
	})

	if err != nil {
		return nil, err
	}

	// if updated password, reset session for user
	if req.User.Password != "" {
		userSession, err := s.data.SessionClient.Get(ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		if userSession != nil {
			if err = s.data.SessionClient.Delete(ctx, req.UserId); err != nil {
				return nil, err
			}
		}
	}

	return &pb.UpdateUserReply{
		User: &pb.UserItem{
			Id:        result.User.Id,
			FullName:  result.User.FullName,
			CreatedAt: result.User.CreatedAt,
			UpdatedAt: result.User.UpdatedAt,
			Email:     result.User.Email,
			Phone:     result.User.Phone,
			Gender:    int32(result.User.Gender),
			Status:    int32(result.User.Status),
		},
	}, nil
}

func (s *UserService) FreezeUser(ctx context.Context, req *pb.FreezeUserRequest) (*pb.FreezeUserReply, error) {
	_, err := s.data.UserClient.UpdateUser(ctx, &innterapi.UpdateUserRequest{
		Id:     req.UserId,
		Status: innterapi.UserStatus_FREEZE,
	})

	if err != nil {
		return nil, err
	}

	userSession, err := s.data.SessionClient.Get(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if userSession != nil {
		userSession.Status = int32(innterapi.UserStatus_FREEZE)
		if err = s.data.SessionClient.Update(ctx, userSession); err != nil {
			return nil, err
		}
	}
	return &pb.FreezeUserReply{FreezedAt: time.Now().Unix()}, nil
}

func (s *UserService) ThawUser(ctx context.Context, req *pb.ThawUserRequest) (*pb.ThawUserReply, error) {
	_, err := s.data.UserClient.UpdateUser(ctx, &innterapi.UpdateUserRequest{
		Id:     req.UserId,
		Status: innterapi.UserStatus_ACTIVITY,
	})

	if err != nil {
		return nil, err
	}

	userSession, err := s.data.SessionClient.Get(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if userSession != nil {
		userSession.Status = int32(innterapi.UserStatus_ACTIVITY)
		if err = s.data.SessionClient.Update(ctx, userSession); err != nil {
			return nil, err
		}
	}
	return &pb.ThawUserReply{ThawedAt: time.Now().Unix()}, nil
}

func (s *UserService) ListUserConfigKey(ctx context.Context, req *pb.ListUserConfigKeyRequest) (*pb.ListUserConfigKeyReply, error) {
	innerReq := &innerapi.ListUserConfigKeyRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.UserClient.ListUserConfigKey(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &pb.ListUserConfigKeyReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *UserService) GetUserConfig(ctx context.Context, req *pb.GetUserConfigRequest) (*pb.GetUserConfigReply, error) {
	reply, err := s.data.UserClient.GetUserConfig(ctx, &innerapi.GetUserConfigRequest{UserId: req.UserId})
	if err != nil {
		return nil, err
	}
	return &pb.GetUserConfigReply{Config: reply.Config}, nil
}

func (s *UserService) UpdateUserConfig(ctx context.Context, req *pb.UpdateUserConfigRequest) (*pb.UpdateUserConfigReply, error) {
	_, err := s.data.UserClient.UpdateUserConfig(ctx, &innerapi.UpdateUserConfigRequest{UserId: req.UserId, Config: req.Config})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserConfigReply{}, nil
}
