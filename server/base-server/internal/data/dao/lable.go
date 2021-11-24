package dao

import (
	"context"
	stderrors "errors"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"
	"server/common/transaction"

	"server/common/log"

	"gorm.io/gorm"
)

type LableDao interface {
	AddLable(ctx context.Context, lable *model.Lable) error
	ListLable(ctx context.Context, query *model.LableListQuery) ([]*model.Lable, int64, error)
	GetLable(ctx context.Context, id string) (*model.Lable, error)
	QueryLable(ctx context.Context, query *model.LableQuery) (*model.Lable, error)
	DeleteLable(ctx context.Context, id string) error
	UpdateLable(ctx context.Context, lable *model.Lable) error
}

type lableDao struct {
	log *log.Helper
	db  transaction.GetDB
}

func NewLableDao(db *gorm.DB, logger log.Logger) LableDao {
	return &lableDao{
		log: log.NewHelper("LableDao", logger),
		db: func(ctx context.Context) *gorm.DB {
			return transaction.GetDBFromCtx(ctx, db)
		},
	}
}

func (d *lableDao) AddLable(ctx context.Context, lable *model.Lable) error {
	res := d.db(ctx).Create(lable)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *lableDao) ListLable(ctx context.Context, query *model.LableListQuery) ([]*model.Lable, int64, error) {
	lables := make([]*model.Lable, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if query.RelegationType != 0 {
		querySql += " and relegation_type = ? "
		params = append(params, query.RelegationType)
	}

	if query.SourceType != 0 {
		querySql += " and source_type = ? "
		params = append(params, query.SourceType)
	}

	if query.LableType != 0 {
		querySql += " and lable_type = ?"
		params = append(params, query.LableType)
	}

	db := d.db(ctx).Model(&model.Lable{}).Where(querySql, params...)

	var totalSize int64
	res := db.Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if query.PageIndex != 0 {
		db = db.Limit(query.PageSize).Offset((query.PageIndex - 1) * query.PageSize)
	}

	res = db.Find(&lables)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return lables, totalSize, nil
}

func (d *lableDao) GetLable(ctx context.Context, id string) (*model.Lable, error) {
	db := d.db(ctx)

	nb := &model.Lable{}
	res := db.First(nb, "id = ?", id)

	if res.Error != nil {
		if stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFindEmpty)
		} else {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFirstFailed)
		}
	}
	return nb, nil
}

func (d *lableDao) QueryLable(ctx context.Context, query *model.LableQuery) (*model.Lable, error) {
	db := d.db(ctx)

	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if query.RelegationType != 0 {
		querySql += " and relegation_type = ? "
		params = append(params, query.RelegationType)
	}

	if query.SourceType != 0 {
		querySql += " and source_type = ? "
		params = append(params, query.SourceType)
	}

	if query.LableType != 0 {
		querySql += " and lable_type = ?"
		params = append(params, query.LableType)
	}

	if query.LableDesc != "" {
		querySql += " and lable_desc = ?"
		params = append(params, query.LableDesc)
	}

	nb := &model.Lable{}
	res := db.Where(querySql, params...).First(nb)

	if res.Error != nil {
		if stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFindEmpty)
		} else {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFirstFailed)
		}
	}
	return nb, nil
}

func (d *lableDao) DeleteLable(ctx context.Context, id string) error {
	db := d.db(ctx)

	if id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	res := db.Where("id = ? ", id).Delete(&model.Lable{})
	if res.Error != nil {
		return errors.Errorf(nil, errors.ErrorDBDeleteFailed)
	}

	return nil
}

func (d *lableDao) UpdateLable(ctx context.Context, lable *model.Lable) error {
	db := d.db(ctx)
	if lable.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := db.Model(&model.Lable{}).Where("id = ? ", lable.Id).Updates(map[string]interface{}{
		"relegation_type": lable.RelegationType,
		"source_type":     lable.SourceType,
		"Lable_type":      lable.LableType,
		"lable_desc":      lable.LableDesc,
		"refer_times":     lable.ReferTimes,
	})

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}
