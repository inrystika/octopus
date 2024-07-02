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
	CreateUserEndpoints(ctx context.Context, userEndpoints []*model.UserEndpoint) error
	IsUserEndpointsNotExist(ctx context.Context, endpoints []string) (bool, error)
	DeleteUserEndpoints(ctx context.Context, endpoints []string) error
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

func (d *userEndpointDao) CreateUserEndpoints(ctx context.Context, userEndpoints []*model.UserEndpoint) error {
	res := d.db.CreateInBatches(userEndpoints, 100)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *userEndpointDao) IsUserEndpointsNotExist(ctx context.Context, endpoints []string) (bool, error) {
	var count int64
	res := d.db.Where("endpoint in ?", endpoints).Model(&model.UserEndpoint{}).Count(&count)

	if res.Error != nil {
		return false, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	return count == 0, nil
}

func (d *userEndpointDao) DeleteUserEndpoints(ctx context.Context, endpoints []string) error {
	if len(endpoints) == 0 {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	res := d.db.Where("endpoint in ?", endpoints).Delete(&model.UserEndpoint{})
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBDeleteFailed)
	}

	return nil
}
