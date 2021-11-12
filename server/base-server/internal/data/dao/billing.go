package dao

import (
	"context"
	stderrors "errors"
	"fmt"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"
	"server/common/log"
	"server/common/transaction"
	"server/common/utils"
	"time"

	"gorm.io/gorm"
)

type BillingDao interface {
	Transaction(ctx context.Context, fc func(ctx context.Context) error) error
	GetBillingOwner(ctx context.Context, key *model.BillingOwnerKey) (*model.BillingOwner, error)
	CreateBillingOwner(ctx context.Context, owner *model.BillingOwner) error
	UpdateBillingOwnerSelective(ctx context.Context, key *model.BillingOwnerKey, record *model.BillingOwner) error
	ListBillingOwner(ctx context.Context, query *model.BillingOwnerQuery) ([]*model.BillingOwner, int64, error)

	GetBillingPayRecord(ctx context.Context, key *model.BillingPayRecordKey) (*model.BillingPayRecord, error)
	CreateBillingPayRecord(ctx context.Context, record *model.BillingPayRecord) error
	UpdateBillingPayRecordSelective(ctx context.Context, key *model.BillingPayRecordKey, record *model.BillingPayRecord) error
	ListBillingPayRecord(ctx context.Context, query *model.BillingPayRecordQuery) ([]*model.BillingPayRecord, int64, error)

	CreateBillingRechargeRecord(ctx context.Context, record *model.BillingRechargeRecord) error
	ListBillingRechargeRecord(ctx context.Context, query *model.BillingRechargeRecordQuery) ([]*model.BillingRechargeRecord, int64, error)
}

type billingDao struct {
	log *log.Helper
	db  transaction.GetDB
}

func NewBillingDao(db *gorm.DB, logger log.Logger) BillingDao {
	return &billingDao{
		log: log.NewHelper("BillingDao", logger),
		db: func(ctx context.Context) *gorm.DB {
			return transaction.GetDBFromCtx(ctx, db)
		},
	}
}

func (d *billingDao) Transaction(ctx context.Context, fc func(ctx context.Context) error) error {
	return transaction.Transaction(ctx, d.db(ctx), fc)
}

func (d *billingDao) GetBillingOwner(ctx context.Context, key *model.BillingOwnerKey) (*model.BillingOwner, error) {
	db := d.db(ctx)
	owner := &model.BillingOwner{}
	res := db.First(owner, "owner_id = ? and owner_type = ? ", key.OwnerId, key.OwnerType)

	if res.Error != nil {
		if stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFindEmpty)
		} else {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFirstFailed)
		}
	}
	return owner, nil
}

func (d *billingDao) CreateBillingOwner(ctx context.Context, owner *model.BillingOwner) error {
	db := d.db(ctx)
	res := db.Create(owner)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *billingDao) UpdateBillingOwnerSelective(ctx context.Context, key *model.BillingOwnerKey, owner *model.BillingOwner) error {
	db := d.db(ctx)
	res := db.Where("owner_id = ? and owner_type = ? ", key.OwnerId, key.OwnerType).Updates(owner)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *billingDao) ListBillingOwner(ctx context.Context, query *model.BillingOwnerQuery) ([]*model.BillingOwner, int64, error) {
	db := d.db(ctx)
	records := make([]*model.BillingOwner, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)

	if query.OwnerId != "" {
		querySql += " and owner_id = ? "
		params = append(params, query.OwnerId)
	}

	if query.OwnerType != 0 {
		querySql += " and owner_type = ? "
		params = append(params, query.OwnerType)
	}

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Model(&model.BillingOwner{}).Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if query.PageIndex != 0 {
		db = db.Limit(query.PageSize).Offset((query.PageIndex - 1) * query.PageSize)
	}

	sortBy := "created_at"
	orderBy := "desc"
	if query.SortBy != "" {
		sortBy = utils.CamelToSnake(query.SortBy)
	}

	if query.OrderBy != "" {
		orderBy = query.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", sortBy, orderBy))

	res = db.Find(&records)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return records, totalSize, nil
}

func (d *billingDao) GetBillingPayRecord(ctx context.Context, key *model.BillingPayRecordKey) (*model.BillingPayRecord, error) {
	db := d.db(ctx)
	nb := &model.BillingPayRecord{}
	res := db.First(nb, "owner_id = ? and owner_type = ? and biz_id = ? and biz_type = ? ", key.OwnerId, key.OwnerType, key.BizId, key.BizType)

	if res.Error != nil {
		if stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFindEmpty)
		} else {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFirstFailed)
		}
	}
	return nb, nil
}

func (d *billingDao) CreateBillingPayRecord(ctx context.Context, record *model.BillingPayRecord) error {
	db := d.db(ctx)
	res := db.Create(record)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *billingDao) UpdateBillingPayRecordSelective(ctx context.Context, key *model.BillingPayRecordKey, record *model.BillingPayRecord) error {
	db := d.db(ctx)
	res := db.Where("owner_id = ? and owner_type = ? and biz_id = ? and biz_type = ? ", key.OwnerId, key.OwnerType, key.BizId, key.BizType).Updates(record)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *billingDao) ListBillingPayRecord(ctx context.Context, query *model.BillingPayRecordQuery) ([]*model.BillingPayRecord, int64, error) {
	db := d.db(ctx)
	records := make([]*model.BillingPayRecord, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)

	if query.SearchKey != "" {
		querySql += " and title like ?"
		params = append(params, "%"+query.SearchKey+"%")
	}

	if query.OwnerId != "" {
		querySql += " and owner_id = ? "
		params = append(params, query.OwnerId)
	}

	if query.OwnerType != 0 {
		querySql += " and owner_type = ? "
		params = append(params, query.OwnerType)
	}

	if query.StartedAtGte != 0 {
		querySql += " and started_at >= ? "
		params = append(params, time.Unix(query.StartedAtGte, 0))
	}

	if query.StartedAtLt != 0 {
		querySql += " and started_at < ? "
		params = append(params, time.Unix(query.StartedAtLt, 0))
	}

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Model(&model.BillingPayRecord{}).Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if query.PageIndex != 0 {
		db = db.Limit(query.PageSize).Offset((query.PageIndex - 1) * query.PageSize)
	}

	sortBy := "created_at"
	orderBy := "desc"
	if query.SortBy != "" {
		sortBy = utils.CamelToSnake(query.SortBy)
	}

	if query.OrderBy != "" {
		orderBy = query.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", sortBy, orderBy))

	res = db.Find(&records)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return records, totalSize, nil
}

func (d *billingDao) CreateBillingRechargeRecord(ctx context.Context, record *model.BillingRechargeRecord) error {
	db := d.db(ctx)
	res := db.Create(record)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *billingDao) ListBillingRechargeRecord(ctx context.Context, query *model.BillingRechargeRecordQuery) ([]*model.BillingRechargeRecord, int64, error) {
	db := d.db(ctx)
	records := make([]*model.BillingRechargeRecord, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)

	if query.OwnerId != "" {
		querySql += " and owner_id = ? "
		params = append(params, query.OwnerId)
	}

	if query.OwnerType != 0 {
		querySql += " and owner_type = ? "
		params = append(params, query.OwnerType)
	}

	if query.CreatedAtGte != 0 {
		querySql += " and created_at >= ? "
		params = append(params, time.Unix(query.CreatedAtGte, 0))
	}

	if query.CreatedAtLt != 0 {
		querySql += " and created_at < ? "
		params = append(params, time.Unix(query.CreatedAtLt, 0))
	}

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Model(&model.BillingRechargeRecord{}).Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if query.PageIndex != 0 {
		db = db.Limit(query.PageSize).Offset((query.PageIndex - 1) * query.PageSize)
	}

	sortBy := "created_at"
	orderBy := "desc"
	if query.SortBy != "" {
		sortBy = utils.CamelToSnake(query.SortBy)
	}

	if query.OrderBy != "" {
		orderBy = query.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", sortBy, orderBy))

	res = db.Find(&records)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return records, totalSize, nil
}
