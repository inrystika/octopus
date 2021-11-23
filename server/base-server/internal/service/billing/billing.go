package billing

import (
	"context"
	"fmt"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"
	"server/common/log"
	"server/common/utils"
	"time"

	"github.com/jinzhu/copier"
)

type billingService struct {
	api.UnimplementedBillingServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

const (
	changeAmountLockKey     = "billing:change-amount-lock:%s-%s"
	changeAmountLockTimeOut = 5 * time.Second
)

func NewBillingService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.BillingServiceServer {
	log := log.NewHelper("BillingService", logger)

	s := &billingService{
		conf: conf,
		log:  log,
		data: data,
	}

	return s
}

func (s *billingService) CreateBillingOwner(ctx context.Context, req *api.CreateBillingOwnerRequest) (*api.CreateBillingOwnerReply, error) {
	err := s.data.BillingDao.CreateBillingOwner(ctx, &model.BillingOwner{
		OwnerId:   req.OwnerId,
		OwnerType: req.OwnerType,
		Amount:    0,
	})
	if err != nil {
		return nil, err
	}
	return &api.CreateBillingOwnerReply{}, nil
}
func (s *billingService) GetBillingOwner(ctx context.Context, req *api.GetBillingOwnerRequest) (*api.GetBillingOwnerReply, error) {
	ownerTbl, err := s.data.BillingDao.GetBillingOwner(ctx, &model.BillingOwnerKey{
		OwnerId:   req.OwnerId,
		OwnerType: req.OwnerType,
	})
	if err != nil {
		return nil, err
	}
	owner := &api.BillingOwner{}
	err = copier.Copy(owner, ownerTbl)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	owner.CreatedAt = ownerTbl.CreatedAt.Unix()
	owner.UpdatedAt = ownerTbl.UpdatedAt.Unix()
	return &api.GetBillingOwnerReply{BillingOwner: owner}, nil
}
func (s *billingService) ListBillingOwner(ctx context.Context, req *api.ListBillingOwnerRequest) (*api.ListBillingOwnerReply, error) {
	query := &model.BillingOwnerQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	ownersTbl, totalSize, err := s.data.BillingDao.ListBillingOwner(ctx, query)
	if err != nil {
		return nil, err
	}

	owners := make([]*api.BillingOwner, 0)
	for _, n := range ownersTbl {
		owner := &api.BillingOwner{}
		err := copier.Copy(owner, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		owner.CreatedAt = n.CreatedAt.Unix()
		owner.UpdatedAt = n.UpdatedAt.Unix()
		owners = append(owners, owner)
	}

	return &api.ListBillingOwnerReply{
		TotalSize:     totalSize,
		BillingOwners: owners,
	}, nil
}

func (s *billingService) Pay(ctx context.Context, req *api.PayRequest) (*api.PayReply, error) {
	err := s.data.Redis.LockAndCall(ctx, fmt.Sprintf(changeAmountLockKey, req.OwnerType, req.OwnerId), changeAmountLockTimeOut, func() error {
		err := s.data.BillingDao.Transaction(ctx, func(ctx context.Context) error {
			ownerKey := &model.BillingOwnerKey{
				OwnerId:   req.OwnerId,
				OwnerType: req.OwnerType,
			}
			owner, err := s.data.BillingDao.GetBillingOwner(ctx, ownerKey)
			if err != nil {
				return err
			}

			changeAmount := 0.0
			recordTbl, err := s.data.BillingDao.GetBillingPayRecord(ctx, &model.BillingPayRecordKey{
				OwnerId:   req.OwnerId,
				OwnerType: req.OwnerType,
				BizId:     req.BizId,
				BizType:   req.BizType,
			})
			if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
				return err
			}

			startedAt := time.Unix(req.StartedAt, 0)
			endedAt := time.Unix(req.EndedAt, 0)
			if errors.IsError(errors.ErrorDBFindEmpty, err) {
				record := &model.BillingPayRecord{}
				err := copier.Copy(record, req)
				if err != nil {
					return err
				}
				record.StartedAt = &startedAt
				record.EndedAt = &endedAt
				record.Id = utils.GetUUIDWithoutSeparator()
				err = s.data.BillingDao.CreateBillingPayRecord(ctx, record)
				if err != nil {
					return err
				}
				changeAmount = req.Amount
			} else {
				if recordTbl.Status == api.BillingPayRecordStatus_BPRS_PAY_COMPLETED {
					return errors.Errorf(nil, errors.ErrorBillingStatusForbidden)
				}

				if req.StartedAt != recordTbl.StartedAt.Unix() {
					return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
				}

				if req.EndedAt <= recordTbl.EndedAt.Unix() && req.Status != api.BillingPayRecordStatus_BPRS_PAY_COMPLETED {
					return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
				}

				if req.Amount <= recordTbl.Amount && req.Status != api.BillingPayRecordStatus_BPRS_PAY_COMPLETED {
					return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
				}

				err := s.data.BillingDao.UpdateBillingPayRecordSelective(ctx, &model.BillingPayRecordKey{
					OwnerId:   req.OwnerId,
					OwnerType: req.OwnerType,
					BizId:     req.BizId,
					BizType:   req.BizType,
				}, &model.BillingPayRecord{
					EndedAt: &endedAt,
					Amount:  req.Amount,
					Status:  req.Status,
				})
				if err != nil {
					return err
				}
				changeAmount = req.Amount - recordTbl.Amount
			}

			err = s.data.BillingDao.UpdateBillingOwnerSelective(ctx, ownerKey, &model.BillingOwner{Amount: owner.Amount - changeAmount})
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return &api.PayReply{}, nil
}

func (s *billingService) Recharge(ctx context.Context, req *api.RechargeRequest) (*api.RechargeReply, error) {
	err := s.data.Redis.LockAndCall(ctx, fmt.Sprintf(changeAmountLockKey, req.OwnerType, req.OwnerId), changeAmountLockTimeOut, func() error {
		err := s.data.BillingDao.Transaction(ctx, func(ctx context.Context) error {
			ownerKey := &model.BillingOwnerKey{
				OwnerId:   req.OwnerId,
				OwnerType: req.OwnerType,
			}
			owner, err := s.data.BillingDao.GetBillingOwner(ctx, ownerKey)
			if err != nil {
				return err
			}

			err = s.data.BillingDao.CreateBillingRechargeRecord(ctx, &model.BillingRechargeRecord{
				Id:        utils.GetUUIDWithoutSeparator(),
				OwnerId:   req.OwnerId,
				OwnerType: req.OwnerType,
				Amount:    req.Amount,
				Title:     req.Title,
			})
			if err != nil {
				return err
			}

			err = s.data.BillingDao.UpdateBillingOwnerSelective(ctx, ownerKey, &model.BillingOwner{Amount: owner.Amount + req.Amount})
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &api.RechargeReply{}, nil
}

func (s *billingService) ListBillingPayRecord(ctx context.Context, req *api.ListBillingPayRecordRequest) (*api.ListBillingPayRecordReply, error) {
	query := &model.BillingPayRecordQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	recordsTbl, totalSize, err := s.data.BillingDao.ListBillingPayRecord(ctx, query)
	if err != nil {
		return nil, err
	}

	records := make([]*api.BillingPayRecord, 0)
	for _, n := range recordsTbl {
		record := &api.BillingPayRecord{}
		err := copier.Copy(record, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		record.CreatedAt = n.CreatedAt.Unix()
		record.UpdatedAt = n.UpdatedAt.Unix()
		record.StartedAt = n.StartedAt.Unix()
		record.EndedAt = n.EndedAt.Unix()
		records = append(records, record)
	}

	return &api.ListBillingPayRecordReply{
		TotalSize: totalSize,
		Records:   records,
	}, nil
}

func (s *billingService) ListBillingRechargeRecord(ctx context.Context, req *api.ListBillingRechargeRecordRequest) (*api.ListBillingRechargeRecordReply, error) {
	query := &model.BillingRechargeRecordQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	recordsTbl, totalSize, err := s.data.BillingDao.ListBillingRechargeRecord(ctx, query)
	if err != nil {
		return nil, err
	}

	records := make([]*api.BillingRechargeRecord, 0)
	for _, n := range recordsTbl {
		record := &api.BillingRechargeRecord{}
		err := copier.Copy(record, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		record.CreatedAt = n.CreatedAt.Unix()
		record.UpdatedAt = n.UpdatedAt.Unix()
		records = append(records, record)
	}

	return &api.ListBillingRechargeRecordReply{
		TotalSize: totalSize,
		Records:   records,
	}, nil
}
