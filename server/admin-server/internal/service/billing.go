package service

import (
	"context"
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	innerapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"
	"server/common/utils/collections/set"

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
		userIds := make([]string, 0)
		for _, i := range billingUsers {
			userIds = append(userIds, i.UserId)
		}

		userMap, err := s.listUserInCond(ctx, set.NewStrings(userIds...).Values())
		if err != nil {
			return err
		}

		for _, i := range billingUsers {

			if v, ok := userMap[i.UserId]; ok {
				i.UserName = v.FullName
			}
		}
	}

	return nil
}

func (s *billingService) listUserInCond(ctx context.Context, ids []string) (map[string]*innerapi.UserItem, error) {
	userMap := map[string]*innerapi.UserItem{}
	users, err := s.data.UserClient.ListUserInCond(ctx, &innerapi.ListUserInCondRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	for _, i := range users.Users {
		userMap[i.Id] = i
	}
	return userMap, nil
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

func (s *billingService) ListUserPayRecord(ctx context.Context, req *api.ListUserPayRecordRequest) (*api.ListUserPayRecordReply, error) {
	innerReq := &innerapi.ListBillingPayRecordRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.OwnerId = req.UserId
	innerReq.OwnerType = innerapi.BillingOwnerType_BOT_USER

	innerReply, err := s.data.BillingClient.ListBillingPayRecord(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListUserPayRecordReply{
		TotalSize: innerReply.TotalSize,
	}
	for _, i := range innerReply.Records {
		r := &api.UserPayRecord{}
		err = copier.Copy(r, i)
		if err != nil {
			return nil, err
		}
		r.UserId = i.OwnerId
		reply.Records = append(reply.Records, r)
	}

	err = s.assignUserPayRecord(ctx, reply.Records)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
func (s *billingService) ListUserRechargeRecord(ctx context.Context, req *api.ListUserRechargeRecordRequest) (*api.ListUserRechargeRecordReply, error) {
	innerReq := &innerapi.ListBillingRechargeRecordRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.OwnerId = req.UserId
	innerReq.OwnerType = innerapi.BillingOwnerType_BOT_USER

	innerReply, err := s.data.BillingClient.ListBillingRechargeRecord(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListUserRechargeRecordReply{
		TotalSize: innerReply.TotalSize,
	}
	for _, i := range innerReply.Records {
		r := &api.UserRechargeRecord{}
		err = copier.Copy(r, i)
		if err != nil {
			return nil, err
		}
		r.UserId = i.OwnerId
		reply.Records = append(reply.Records, r)
	}

	err = s.assignUserRechargeRecord(ctx, reply.Records)
	if err != nil {
		return nil, err
	}

	return reply, nil
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
		spaceIds := make([]string, 0)
		for _, i := range billingSpaces {
			spaceIds = append(spaceIds, i.SpaceId)
		}

		spaceMap, err := s.listWorkspaceInCond(ctx, set.NewStrings(spaceIds...).Values())
		if err != nil {
			return err
		}

		for _, i := range billingSpaces {
			if v, ok := spaceMap[i.SpaceId]; ok {
				i.SpaceName = v.Name
			}
		}
	}

	return nil
}

func (s *billingService) listWorkspaceInCond(ctx context.Context, ids []string) (map[string]*innerapi.WorkspaceItem, error) {
	spaceMap := map[string]*innerapi.WorkspaceItem{}
	spaces, err := s.data.WorkspaceClient.ListWorkspaceInCond(ctx, &innerapi.ListWorkspaceInCondRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	for _, i := range spaces.Workspaces {
		spaceMap[i.Id] = i
	}

	return spaceMap, nil
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

func (s *billingService) ListSpacePayRecord(ctx context.Context, req *api.ListSpacePayRecordRequest) (*api.ListSpacePayRecordReply, error) {
	innerReq := &innerapi.ListBillingPayRecordRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.OwnerId = req.SpaceId
	innerReq.OwnerType = innerapi.BillingOwnerType_BOT_SPACE

	innerReply, err := s.data.BillingClient.ListBillingPayRecord(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListSpacePayRecordReply{
		TotalSize: innerReply.TotalSize,
	}
	for _, i := range innerReply.Records {
		r := &api.SpacePayRecord{}
		err = copier.Copy(r, i)
		if err != nil {
			return nil, err
		}
		r.SpaceId = i.OwnerId
		reply.Records = append(reply.Records, r)
	}

	err = s.assignSpacePayRecord(ctx, reply.Records)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
func (s *billingService) ListSpaceRechargeRecord(ctx context.Context, req *api.ListSpaceRechargeRecordRequest) (*api.ListSpaceRechargeRecordReply, error) {
	innerReq := &innerapi.ListBillingRechargeRecordRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.OwnerId = req.SpaceId
	innerReq.OwnerType = innerapi.BillingOwnerType_BOT_SPACE

	innerReply, err := s.data.BillingClient.ListBillingRechargeRecord(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListSpaceRechargeRecordReply{
		TotalSize: innerReply.TotalSize,
	}
	for _, i := range innerReply.Records {
		r := &api.SpaceRechargeRecord{}
		err = copier.Copy(r, i)
		if err != nil {
			return nil, err
		}
		r.SpaceId = i.OwnerId
		reply.Records = append(reply.Records, r)
	}

	err = s.assignSpaceRechargeRecord(ctx, reply.Records)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *billingService) assignUserPayRecord(ctx context.Context, records []*api.UserPayRecord) error {
	if len(records) > 0 {
		userIds := make([]string, 0)
		for _, i := range records {
			userIds = append(userIds, i.UserId)
		}

		userMap, err := s.listUserInCond(ctx, set.NewStrings(userIds...).Values())
		if err != nil {
			return err
		}

		for _, i := range records {

			if v, ok := userMap[i.UserId]; ok {
				i.UserName = v.FullName
			}
		}
	}

	return nil
}

func (s *billingService) assignUserRechargeRecord(ctx context.Context, records []*api.UserRechargeRecord) error {
	if len(records) > 0 {
		userIds := make([]string, 0)
		for _, i := range records {
			userIds = append(userIds, i.UserId)
		}

		userMap, err := s.listUserInCond(ctx, set.NewStrings(userIds...).Values())
		if err != nil {
			return err
		}

		for _, i := range records {

			if v, ok := userMap[i.UserId]; ok {
				i.UserName = v.FullName
			}
		}
	}

	return nil
}

func (s *billingService) assignSpacePayRecord(ctx context.Context, records []*api.SpacePayRecord) error {
	if len(records) > 0 {
		spaceIds := make([]string, 0)
		for _, i := range records {
			spaceIds = append(spaceIds, i.SpaceId)
		}

		spaceMap, err := s.listWorkspaceInCond(ctx, set.NewStrings(spaceIds...).Values())
		if err != nil {
			return err
		}

		for _, i := range records {
			if v, ok := spaceMap[i.SpaceId]; ok {
				i.SpaceName = v.Name
			}
		}
	}

	return nil
}

func (s *billingService) assignSpaceRechargeRecord(ctx context.Context, records []*api.SpaceRechargeRecord) error {
	if len(records) > 0 {
		spaceIds := make([]string, 0)
		for _, i := range records {
			spaceIds = append(spaceIds, i.SpaceId)
		}

		spaceMap, err := s.listWorkspaceInCond(ctx, set.NewStrings(spaceIds...).Values())
		if err != nil {
			return err
		}

		for _, i := range records {
			if v, ok := spaceMap[i.SpaceId]; ok {
				i.SpaceName = v.Name
			}
		}
	}

	return nil
}
