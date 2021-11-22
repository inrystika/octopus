package dao

import (
	"context"
	"errors"
	"server/base-server/internal/data/dao/model"
	"server/common/log"
	"server/common/utils/collections/set"

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
	result := db.Where(&model.User{
		Id:    condition.Id,
		Email: condition.Email,
		Phone: condition.Phone,
	}).First(&user)

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

	u := model.User{
		Id:       user.Id,
		FullName: user.FullName,
		Gender:   user.Gender,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
		Status:   user.Status,
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
		FullName: user.FullName,
		Email:    user.Email,
		Phone:    user.Phone,
		Gender:   user.Gender,
		Password: user.Password,
		Status:   user.Status,
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
	if userId == "" {
		return commerrors.Errorf(nil, commerrors.ErrorInvalidRequestParameter)
	}

	configDb, err := d.GetConfig(ctx, userId)
	if err != nil {
		return err
	}

	deleteKeys := make([]string, 0)
	for k, _ := range configDb {
		_, ok := config[k]
		if !ok {
			deleteKeys = append(deleteKeys, k)
		}
	}

	if len(deleteKeys) > 0 {
		res := db.Where("user_id = ? and key in ?", userId, deleteKeys).Delete(&model.UserConfig{})
		if res.Error != nil {
			return commerrors.Errorf(res.Error, commerrors.ErrorDBDeleteFailed)
		}
	}

	insertKeys := make([]string, 0)
	updateKeys := make([]string, 0)
	for k, v := range config {
		vdb, ok := configDb[k]
		if ok && v != vdb {
			updateKeys = append(updateKeys, k)
			continue
		}
		if !ok && v != "" {
			insertKeys = append(insertKeys, k)
			continue
		}
	}

	if len(insertKeys) > 0 {
		items := make([]*model.UserConfig, 0)
		for _, k := range insertKeys {
			items = append(items, &model.UserConfig{
				UserId: userId,
				Key:    k,
				Value:  config[k],
			})
		}
		res := db.Create(&items)
		if res.Error != nil {
			return commerrors.Errorf(res.Error, commerrors.ErrorDBCreateFailed)
		}
	}

	if len(updateKeys) > 0 {
		for _, k := range updateKeys {
			res := db.Model(&model.UserConfig{}).Where("user_id = ? and `key` = ?", userId, k).Limit(1).Update("value", config[k])
			if res.Error != nil {
				return commerrors.Errorf(res.Error, commerrors.ErrorDBUpdateFailed)
			}
		}
	}

	return nil
}

func (d *userDao) GetConfig(ctx context.Context, userId string) (map[string]string, error) {
	db := d.db
	userConfigs := make([]*model.UserConfig, 0)

	res := db.Where("user_id = ?", userId).Find(&userConfigs)
	if res.Error != nil {
		return nil, commerrors.Errorf(res.Error, commerrors.ErrorDBFindFailed)
	}

	r := map[string]string{}
	for _, i := range userConfigs {
		r[i.Key] = i.Value
	}

	return r, nil
}
