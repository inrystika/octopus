package dao

import (
	"context"
	"errors"
	"server/base-server/internal/data/dao/model"
	"server/common/log"

	"gorm.io/gorm"
)

type AdminUserDao interface {
	List(ctx context.Context, query model.AdminUserList) ([]*model.AdminUser, error)
	Find(ctx context.Context, adminUser model.AdminUser) (*model.AdminUser, error)
	Add(ctx context.Context, adminUser *model.AdminUser) error
	Count(ctx context.Context, adminUser model.AdminUserQuery) (int64, error)
	Delete(ctx context.Context, adminUser model.AdminUserQuery) error
}

type adminUserDao struct {
	log *log.Helper
	db  *gorm.DB
}

func NewAdminUserDao(db *gorm.DB, logger log.Logger) AdminUserDao {
	return &adminUserDao{
		log: log.NewHelper("AdminUserDao", logger),
		db:  db,
	}
}

func (d *adminUserDao) Delete(ctx context.Context, condition model.AdminUserQuery) error {
	db := d.db.Unscoped()

	condition.Where(db)
	result := db.Delete(model.AdminUser{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (d *adminUserDao) Count(ctx context.Context, condition model.AdminUserQuery) (int64, error) {
	db := d.db
	var count int64
	condition.Where(db)

	db.Model(&model.AdminUser{}).Count(&count)
	return count, nil
}

func (d *adminUserDao) List(ctx context.Context, query model.AdminUserList) ([]*model.AdminUser, error) {
	db := d.db
	users := make([]*model.AdminUser, 0)
	if query.PageIndex != 0 {
		db = db.Limit(query.PageSize).Offset((query.PageIndex - 1) * query.PageSize)
	}

	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if query.Username != "" {
		querySql += " and name = ? "
		params = append(params, query.Username)
	}

	db = db.Where(querySql, params...)
	db.Find(&users)

	return users, nil
}

func (d *adminUserDao) Find(ctx context.Context, adminUser model.AdminUser) (*model.AdminUser, error) {
	db := d.db

	var user model.AdminUser
	result := db.Where(&adminUser).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (d *adminUserDao) Add(ctx context.Context, adminUser *model.AdminUser) error {
	db := d.db
	result := db.Create(adminUser)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
