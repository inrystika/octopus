package dao

import (
	"context"
	"errors"
	stderrors "errors"
	"server/base-server/internal/data/dao/model"
	"server/common/log"
	"server/common/utils/collections/set"

	"gorm.io/gorm/clause"

	commerrors "server/common/errors"

	"gorm.io/gorm"
)

type UserDao interface {
	List(ctx context.Context, condition *model.UserList) ([]*model.User, error)
	Count(ctx context.Context, condition *model.UserList) (int64, error)
	Find(ctx context.Context, condition *model.UserQuery) (*model.User, error)
	Add(ctx context.Context, user *model.UserAdd) (*model.User, error)
	Update(ctx context.Context, condition *model.UserUpdateCond, user *model.UserUpdate) (*model.User, error)
	ListIn(ctx context.Context, condition *model.UserListIn) ([]*model.User, error)
	UpdateConfig(ctx context.Context, userId string, config map[string]string) error
	GetConfig(ctx context.Context, userId string) (map[string]string, error)
}

type userDao struct {
	log *log.Helper
	db  *gorm.DB
}

func NewUserDao(db *gorm.DB, logger log.Logger) UserDao {
	return &userDao{
		log: log.NewHelper("UserDao", logger),
		db:  db,
	}
}

func (d *userDao) List(ctx context.Context, condition *model.UserList) ([]*model.User, error) {
	db := d.db
	users := make([]*model.User, 0)

	db = condition.Pagination(db)
	db = condition.Order(db)
	db = condition.Where(db)
	db = condition.Or(db)
	db.Find(&users)

	return users, nil
}

func (d *userDao) Count(ctx context.Context, condition *model.UserList) (int64, error) {
	db := d.db
	var count int64

	db = condition.Where(db)
	db = condition.Or(db)

	db.Model(&model.User{}).Count(&count)
	return count, nil
}

func (d *userDao) Find(ctx context.Context, condition *model.UserQuery) (*model.User, error) {
	db := d.db

	var user model.User
	var result *gorm.DB
	if condition.Bind == nil {
		result = db.Where(&model.User{
			Id:    condition.Id,
			Email: condition.Email,
			Phone: condition.Phone,
		}).First(&user)
	} else {
		querySql := "1 = 1"
		params := make([]interface{}, 0)
		if condition.Email != "" {
			querySql += " and email = ? "
			params = append(params, condition.Email)
			if condition.Bind.UserId != "" {
				querySql += " or (JSON_CONTAINS(bind,JSON_OBJECT('platform', ?))"
				params = append(params, condition.Bind.Platform)
				querySql += " and JSON_CONTAINS(bind,JSON_OBJECT('userId', ?)))"
				params = append(params, condition.Bind.UserId)
			} else {
				querySql += " and JSON_CONTAINS(bind,JSON_OBJECT('platform', ?))"
				params = append(params, condition.Bind.Platform)
			}
		} else {
			querySql += " and JSON_CONTAINS(bind,JSON_OBJECT('platform', ?))"
			params = append(params, condition.Bind.Platform)
			querySql += " and JSON_CONTAINS(bind,JSON_OBJECT('userId', ?))"
			params = append(params, condition.Bind.UserId)
		}
		result = db.Where(querySql, params...).First(&user)
	}

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (d *userDao) Add(ctx context.Context, user *model.UserAdd) (*model.User, error) {
	db := d.db
	bindInfo := make([]*model.Bind, 0)
	if user.Bind != nil {
		bindInfo = append(bindInfo, user.Bind)
	}
	u := model.User{
		Id:       user.Id,
		FullName: user.FullName,
		Gender:   user.Gender,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
		Status:   user.Status,
		Bind:     bindInfo,
	}

	result := db.Create(&u)
	if result.Error != nil {
		return nil, result.Error
	}

	return &u, nil
}

func (d *userDao) Update(ctx context.Context, cond *model.UserUpdateCond, user *model.UserUpdate) (*model.User, error) {
	if cond.Id == "" {
		return nil, gorm.ErrPrimaryKeyRequired
	}

	condition := model.User{
		Id:    cond.Id,
		Email: cond.Email,
		Phone: cond.Phone,
	}

	result := d.db.Model(&condition).Updates(model.User{
		FullName:    user.FullName,
		Email:       user.Email,
		Phone:       user.Phone,
		Gender:      user.Gender,
		Password:    user.Password,
		Status:      user.Status,
		Bind:        user.Bind,
		FtpUserName: user.FtpUserName,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return d.Find(ctx, &model.UserQuery{
		Id:    cond.Id,
		Email: cond.Email,
		Phone: cond.Phone,
	})
}

func (d *userDao) ListIn(ctx context.Context, condition *model.UserListIn) ([]*model.User, error) {
	if len(condition.Ids) < 1 {
		return nil, gorm.ErrMissingWhereClause
	}
	idsSet := set.NewStrings(condition.Ids...)
	var users []*model.User
	result := d.db.Find(&users, idsSet.Values())
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (d *userDao) UpdateConfig(ctx context.Context, userId string, config map[string]string) error {
	db := d.db
	if userId == "" || len(config) == 0 {
		return commerrors.Errorf(nil, commerrors.ErrorInvalidRequestParameter)
	}

	configUp := make(map[string]string)
	for k, v := range config {
		if v != "" {
			configUp[k] = v
		}
	}

	c := &model.UserConfig{UserId: userId, Config: configUp}
	res := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(c)
	if res.Error != nil {
		return commerrors.Errorf(nil, commerrors.ErrorDBUpdateFailed)
	}

	return nil
}

func (d *userDao) GetConfig(ctx context.Context, userId string) (map[string]string, error) {
	db := d.db
	c := &model.UserConfig{}

	res := db.First(c, "user_id = ?", userId)
	if res.Error != nil && !stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, commerrors.Errorf(res.Error, commerrors.ErrorDBFindFailed)
	}

	return c.Config, nil
}
