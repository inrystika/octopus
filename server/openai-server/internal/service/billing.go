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

func (s *billingService) GetBillingUser(ctx context.Context, req *api.GetBillingUserRequest) (*api.GetBillingUserReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	owner, err := s.data.BillingClient.GetBillingOwner(ctx, &innerapi.GetBillingOwnerRequest{
		OwnerId:   session.UserId,
		OwnerType: innerapi.BillingOwnerType_BOT_USER,
	})
	if err != nil {
		return nil, err
	}

	return &api.GetBillingUserReply{BillingUser: &api.BillingUser{
		CreatedAt: owner.BillingOwner.CreatedAt,
		UpdatedAt: owner.BillingOwner.UpdatedAt,
		Amount:    owner.BillingOwner.Amount,
	}}, nil
}
func (s *billingService) ListUserPayRecord(ctx context.Context, req *api.ListUserPayRecordRequest) (*api.ListUserPayRecordReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.ListBillingPayRecordRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.OwnerId = session.UserId
	innerReq.OwnerType = innerapi.BillingOwnerType_BOT_USER

	innerReply, err := s.data.BillingClient.ListBillingPayRecord(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListUserPayRecordReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
func (s *billingService) ListUserRechargeRecord(ctx context.Context, req *api.ListUserRechargeRecordRequest) (*api.ListUserRechargeRecordReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.ListBillingRechargeRecordRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.OwnerId = session.UserId
	innerReq.OwnerType = innerapi.BillingOwnerType_BOT_USER

	innerReply, err := s.data.BillingClient.ListBillingRechargeRecord(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListUserRechargeRecordReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
func (s *billingService) GetBillingSpace(ctx context.Context, req *api.GetBillingSpaceRequest) (*api.GetBillingSpaceReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	owner, err := s.data.BillingClient.GetBillingOwner(ctx, &innerapi.GetBillingOwnerRequest{
		OwnerId:   session.GetWorkspace(),
		OwnerType: innerapi.BillingOwnerType_BOT_SPACE,
	})
	if err != nil {
		return nil, err
	}

	return &api.GetBillingSpaceReply{BillingSpace: &api.BillingSpace{
		CreatedAt: owner.BillingOwner.CreatedAt,
		UpdatedAt: owner.BillingOwner.UpdatedAt,
		Amount:    owner.BillingOwner.Amount,
	}}, nil
}
func (s *billingService) ListSpacePayRecord(ctx context.Context, req *api.ListSpacePayRecordRequest) (*api.ListSpacePayRecordReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.ListBillingPayRecordRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.OwnerId = session.GetWorkspace()
	innerReq.OwnerType = innerapi.BillingOwnerType_BOT_SPACE

	innerReply, err := s.data.BillingClient.ListBillingPayRecord(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListSpacePayRecordReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
func (s *billingService) ListSpaceRechargeRecord(ctx context.Context, req *api.ListSpaceRechargeRecordRequest) (*api.ListSpaceRechargeRecordReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.ListBillingRechargeRecordRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.OwnerId = session.GetWorkspace()
	innerReq.OwnerType = innerapi.BillingOwnerType_BOT_SPACE

	innerReply, err := s.data.BillingClient.ListBillingRechargeRecord(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListSpaceRechargeRecordReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
