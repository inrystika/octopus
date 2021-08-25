package service

import (
	"context"
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	innterapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/jwt"
	"server/common/utils"
	"time"

	"server/common/log"
)

type AuthService struct {
	api.UnimplementedAuthServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewAuthService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.AuthServer {
	return &AuthService{
		conf: conf,
		log:  log.NewHelper("AuthService", logger),
		data: data,
	}
}

func (s *AuthService) GetToken(ctx context.Context, req *api.GetTokenRequest) (*api.GetTokenReply, error) {
	if req.Username == "" {
		return nil, errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	adminUser, err := s.data.AdminUserClient.FindAdminUserByUsername(ctx, &innterapi.AdminUsernameRequest{
		UserName: req.Username,
	})
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorAuthenticationFailed)
	}

	if !utils.ValidatePassword(adminUser.Password, req.Password) {
		return nil, errors.Errorf(err, errors.ErrorAuthenticationFailed)
	}

	token, err := jwt.CreateToken(adminUser.Id, s.conf.Server.Http.JwtSecrect, time.Second*time.Duration(s.conf.Service.TokenExpirationSec))
	if err != nil {
		return nil, err
	}

	return &api.GetTokenReply{
		Token:      token,
		Expiration: 0,
	}, nil
}
