package dao

import (
	"context"
	"gorm.io/gorm"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"
)

type userEndpointDao struct {
	db *gorm.DB
}

type UserEndpointDao interface {
	CreateUserEndpoint(ctx context.Context, userEndpoint *model.UserEndpoint) error
	IsUserEndpointExist(ctx context.Context, endpoint string) (bool, error)
	DeleteUserEndpoint(ctx context.Context, endpoint string) error
}

func NewUserEndpointDao(db *gorm.DB) UserEndpointDao {
	return &userEndpointDao{
		db: db,
	}
}

func (d *userEndpointDao) CreateUserEndpoint(ctx context.Context, userEndpoint *model.UserEndpoint) error {
	res := d.db.Create(userEndpoint)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *userEndpointDao) IsUserEndpointExist(ctx context.Context, endpoint string) (bool, error) {
	var count int64
	res := d.db.Where("endpoint = ?", endpoint).Count(&count)

	if res.Error != nil {
		return false, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	return count > 0, nil
}

func (d *userEndpointDao) DeleteUserEndpoint(ctx context.Context, endpoint string) error {
	if endpoint == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	res := d.db.Where("endpoint = ?", endpoint).Delete(&model.UserEndpoint{})
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBDeleteFailed)
	}

	return nil
}
