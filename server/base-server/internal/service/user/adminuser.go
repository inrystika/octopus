package user

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"server/common/utils"

	"server/common/log"

	"github.com/jinzhu/copier"
	"gopkg.in/errgo.v2/errors"
)

type AdminUserService struct {
	api.UnimplementedAdminUserServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewAdminUserService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.AdminUserServer {
	svcLog := log.NewHelper("AlgorithmService", logger)
	service := &AdminUserService{
		conf: conf,
		log:  svcLog,
		data: data,
	}
	_, err := service.AddDefaultAdminUser(context.TODO())
	if err != nil {
		svcLog.Warnf(context.TODO(), "add default admin user failed, cause by:　%v", err)
	}
	return service
}

func (s *AdminUserService) ListUser(ctx context.Context, req *api.ListAdminUserRequest) (*api.ListAdminUserReply, error) {
	usersTbl, err := s.data.AdminUserDao.List(ctx, model.AdminUserQuery{
		PageIndex: int(req.PageIndex),
		PageSize:  int(req.PageSize),
		Username:  req.UserName,
	})
	if err != nil {
		return nil, err
	}

	users := make([]*api.AdminUserItem, 0)
	for _, a := range usersTbl {
		user := &api.AdminUserItem{}
		_ = copier.Copy(user, a)
		users = append(users, user)
	}

	return &api.ListAdminUserReply{
		Users: users,
	}, nil
}

func (s *AdminUserService) FindAdminUserByUsername(ctx context.Context, req *api.AdminUsernameRequest) (*api.AdminUserItem, error) {
	adminUser, err := s.data.AdminUserDao.Find(ctx, model.AdminUser{
		Username: req.UserName,
	})
	if err != nil {
		return nil, err
	}

	return &api.AdminUserItem{
		Id:       adminUser.Id,
		Username: adminUser.Username,
		Email:    adminUser.Email,
		Phone:    adminUser.Phone,
		Password: adminUser.Password,
	}, nil
}

func (s *AdminUserService) AddDefaultAdminUser(ctx context.Context) (*api.AdminUserItem, error) {
	if s.conf.Administrator.Username == "" || s.conf.Administrator.Password == "" {
		return nil, errors.New("Please configure administrator information.")
	}

	password, err := utils.EncryptPassword(s.conf.Administrator.Password)
	if err != nil {
		return nil, err
	}

	var defaultAdminUser = model.AdminUser{
		Id:       utils.GetUUIDWithoutSeparator(),
		Username: s.conf.Administrator.Username,
		Password: password,
		Email:    s.conf.Administrator.Email,
		Phone:    s.conf.Administrator.Phone,
	}
	result, err := s.data.AdminUserDao.Find(ctx, model.AdminUser{
		Username: defaultAdminUser.Username + "aaa",
	})
	if err != nil {
		return nil, err
	}
	if result != nil {
		return nil, nil
	}

	err = s.data.AdminUserDao.Add(ctx, &defaultAdminUser)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
