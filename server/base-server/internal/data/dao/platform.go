package dao

import (
	"context"
	"fmt"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"
	"server/common/transaction"
	"server/common/utils"
	"time"

	"gorm.io/gorm"
)

type PlatformDao interface {
	CreatePlatform(ctx context.Context, platform *model.Platform) error
	ListPlatform(ctx context.Context, query *model.PlatformQuery) ([]*model.Platform, int64, error)
	// 空值字段不更新
	UpdatePlatformById(ctx context.Context, platform *model.Platform) error
}

type platformDao struct {
	db transaction.GetDB
}

func NewPlatformDao(db *gorm.DB) PlatformDao {
	return &platformDao{
		db: func(ctx context.Context) *gorm.DB {
			return transaction.GetDBFromCtx(ctx, db)
		},
	}
}

func (d *platformDao) CreatePlatform(ctx context.Context, platform *model.Platform) error {
	db := d.db(ctx)
	res := db.Create(platform)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *platformDao) ListPlatform(ctx context.Context, query *model.PlatformQuery) ([]*model.Platform, int64, error) {
	db := d.db(ctx)
	platforms := make([]*model.Platform, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if query.CreatedAtGte != 0 {
		querySql += " and created_at >= ? "
		params = append(params, time.Unix(query.CreatedAtGte, 0))
	}

	if query.CreatedAtLt != 0 {
		querySql += " and created_at < ? "
		params = append(params, time.Unix(query.CreatedAtLt, 0))
	}

	if query.SearchKey != "" {
		querySql += " and name like ?"
		params = append(params, "%"+query.SearchKey+"%")
	}

	if len(query.Ids) != 0 {
		querySql += " and id in ? "
		params = append(params, query.Ids)
	}

	if query.Name != "" {
		querySql += " and name = ? "
		params = append(params, query.Name)
	}

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Model(&model.Platform{}).Count(&totalSize)
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

	res = db.Find(&platforms)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return platforms, totalSize, nil
}

func (d *platformDao) UpdatePlatformById(ctx context.Context, platform *model.Platform) error {
	db := d.db(ctx)
	if platform.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := db.Where("id = ?", platform.Id).Updates(platform)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}
