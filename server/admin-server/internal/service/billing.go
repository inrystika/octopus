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

type billingService struct {
	api.UnimplementedBillingServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewBillingService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.BillingServiceServer {
	log := log.NewHelper("BillingService", logger)

	s := &billingService{
		conf: conf,
		log:  log,
		data: data,
	}

	return s
}

func (s *billingService) ListBillingUser(ctx context.Context, req *api.ListBillingUserRequest) (*api.ListBillingUserReply, error) {
	innerReq := &innerapi.ListBillingOwnerRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.OwnerId = req.UserId
	innerReq.OwnerType = innerapi.BillingOwnerType_BOT_USER

	innerReply, err := s.data.BillingClient.ListBillingOwner(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListBillingUserReply{
		TotalSize: innerReply.TotalSize,
	}
	for _, i := range innerReply.BillingOwners {
		reply.BillingUsers = append(reply.BillingUsers, &api.BillingUser{
			CreatedAt: i.CreatedAt,
			UpdatedAt: i.UpdatedAt,
			Amount:    i.Amount,
			UserId:    i.OwnerId,
		})
	}

	err = s.assignUser(ctx, reply.BillingUsers)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *billingService) assignUser(ctx context.Context, billingUsers []*api.BillingUser) error {
	if len(billingUsers) > 0 {
		userIdMap := map[string]interface{}{}
		for _, i := range billingUsers {
			if i.UserId != "" {
				userIdMap[i.UserId] = true
			}

		}

		userMap := map[string]*innerapi.UserItem{}
		if len(userIdMap) > 0 {
			users, err := s.data.UserClient.ListUserInCond(ctx, &innerapi.ListUserInCondRequest{
				Ids: utils.MapKeyToSlice(userIdMap),
			})
			if err != nil {
				return err
			}
			for _, i := range users.Users {
				userMap[i.Id] = i
			}
		}

		for _, i := range billingUsers {

			if v, ok := userMap[i.UserId]; ok {
				i.UserName = v.FullName
			}
		}
	}

	return nil
}

func (s *billingService) RechargeUser(ctx context.Context, req *api.RechargeUserRequest) (*api.RechargeUserReply, error) {
	_, err := s.data.BillingClient.Recharge(ctx, &innerapi.RechargeRequest{
		OwnerId:   req.UserId,
		OwnerType: innerapi.BillingOwnerType_BOT_USER,
		Amount:    req.Amount,
	})

	if err != nil {
		return nil, err
	}

	return &api.RechargeUserReply{}, nil
}
func (s *billingService) ListBillingSpace(ctx context.Context, req *api.ListBillingSpaceRequest) (*api.ListBillingSpaceReply, error) {
	innerReq := &innerapi.ListBillingOwnerRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.OwnerId = req.SpaceId
	innerReq.OwnerType = innerapi.BillingOwnerType_BOT_SPACE

	innerReply, err := s.data.BillingClient.ListBillingOwner(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListBillingSpaceReply{
		TotalSize: innerReply.TotalSize,
	}
	for _, i := range innerReply.BillingOwners {
		reply.BillingSpaces = append(reply.BillingSpaces, &api.BillingSpace{
			CreatedAt: i.CreatedAt,
			UpdatedAt: i.UpdatedAt,
			Amount:    i.Amount,
			SpaceId:   i.OwnerId,
		})
	}

	err = s.assignSpace(ctx, reply.BillingSpaces)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *billingService) assignSpace(ctx context.Context, billingSpaces []*api.BillingSpace) error {
	if len(billingSpaces) > 0 {
		userIdMap := map[string]interface{}{}
		spaceIdMap := map[string]interface{}{}
		for _, i := range billingSpaces {
			if i.SpaceId != "" {
				spaceIdMap[i.SpaceId] = true
			}
		}

		userMap := map[string]*innerapi.UserItem{}
		if len(userIdMap) > 0 {
			users, err := s.data.UserClient.ListUserInCond(ctx, &innerapi.ListUserInCondRequest{
				Ids: utils.MapKeyToSlice(userIdMap),
			})
			if err != nil {
				return err
			}
			for _, i := range users.Users {
				userMap[i.Id] = i
			}
		}

		spaceMap := map[string]*innerapi.WorkspaceItem{}
		if len(spaceIdMap) > 0 {
			spaces, err := s.data.WorkspaceClient.ListWorkspaceInCond(ctx, &innerapi.ListWorkspaceInCondRequest{
				Ids: utils.MapKeyToSlice(spaceIdMap),
			})
			if err != nil {
				return err
			}
			for _, i := range spaces.Workspaces {
				spaceMap[i.Id] = i
			}

		}

		for _, i := range billingSpaces {
			if v, ok := spaceMap[i.SpaceId]; ok {
				i.SpaceName = v.Name
			}
		}
	}

	return nil
}

func (s *billingService) RechargeSpace(ctx context.Context, req *api.RechargeSpaceRequest) (*api.RechargeSpaceReply, error) {
	_, err := s.data.BillingClient.Recharge(ctx, &innerapi.RechargeRequest{
		OwnerId:   req.SpaceId,
		OwnerType: innerapi.BillingOwnerType_BOT_SPACE,
		Amount:    req.Amount,
	})

	if err != nil {
		return nil, err
	}

	return &api.RechargeSpaceReply{}, nil
}
