package service

import (
	"context"
	"encoding/base64"
	innterapi "server/base-server/api/v1"
	commctx "server/common/context"
	"server/common/errors"
	"server/common/jwt"
	"server/common/log"
	ss "server/common/session"
	"server/common/utils"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"
	"time"
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

	reply, err := s.data.UserClient.FindUser(ctx, &innterapi.FindUserRequest{
		Email: req.Username,
	})
	if err != nil {
		return nil, err
	}
	if reply.User == nil {
		return nil, errors.Errorf(nil, errors.ErrorAuthenticationFailed)
	}
	if reply.User.Status != innterapi.UserStatus_ACTIVITY {
		return nil, errors.Errorf(nil, errors.ErrorAuthenticationForbidden)
	}
	if !utils.ValidatePassword(reply.User.Password, req.Password) {
		return nil, errors.Errorf(err, errors.ErrorAuthenticationFailed)
	}
	//绑定第三方账号
	if req.Bind != nil {
		userId, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(req.Bind.UserId, "%3D", "="))
		if err != nil {
			return nil, err
		}
		reqBind := &innterapi.Bind{
			Platform: req.Bind.Platform,
			UserId:   string(userId),
			UserName: req.Bind.UserName,
		}
		rep, err := s.data.UserClient.FindUser(ctx, &innterapi.FindUserRequest{
			Bind: reqBind,
		})
		if err != nil {
			return nil, err
		}
		if rep.User != nil {
			return nil, errors.Errorf(nil, errors.ErrorUserAccountBinded)
		}
		bindInfo := make([]*innterapi.Bind, 0)
		bindInfo = append(bindInfo, reqBind)
		if reply.User.Bind != nil {
			for _, a := range reply.User.Bind {
				bindInfo = append(bindInfo, a)
			}
		}
		_, err0 := s.data.UserClient.UpdateUser(ctx, &innterapi.UpdateUserRequest{
			Id:   reply.User.Id,
			Bind: bindInfo,
		})
		if err0 != nil {
			return nil, err0
		}
	} //完成绑定

	token, err := jwt.CreateToken(reply.User.Id, s.conf.Server.Http.JwtSecrect, time.Second*time.Duration(s.conf.Service.TokenExpirationSec))
	if err != nil {
		return nil, err
	}
	tokenClaim, err := jwt.ParseToken(token, s.conf.Server.Http.JwtSecrect)
	if err != nil {
		return nil, err
	}
	// create user online session
	if err = s.data.SessionClient.Create(ctx, &ss.Session{
		Id:         reply.User.Id,
		UserId:     reply.User.Id,
		Status:     int32(reply.User.Status),
		Attributes: make(map[string]string),
		CreatedAt:  tokenClaim.CreatedAt,
	}); err != nil {
		return nil, err
	}

	return &api.GetTokenReply{
		Token:      token,
		Expiration: 0,
	}, nil
}

func (s *AuthService) DeleteToken(ctx context.Context, req *api.DeleteTokenRequest) (*api.DeleteTokenReply, error) {
	userId := commctx.UserIdFromContext(ctx)
	if err := s.data.SessionClient.Delete(ctx, userId); err != nil {
		return nil, err
	}
	return &api.DeleteTokenReply{}, nil
}

func (s *AuthService) RegisterAndBind(ctx context.Context, req *api.RegisterRequest) (*api.RegisterReply, error) {
	if req.Bind == nil {
		return nil, errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	//判断用户名是否已存在以及第三方账号是否已绑定
	userId, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(req.Bind.UserId, "%3D", "="))
	if err != nil {
		return nil, err
	}
	reqBind := &innterapi.Bind{
		Platform: req.Bind.Platform,
		UserId:   string(userId),
		UserName: req.Bind.UserName,
	}
	//注册并绑定
	newUser, err := s.data.UserClient.AddUser(ctx, &innterapi.AddUserRequest{
		Email:    req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Gender:   innterapi.GenderType(req.Gender),
		Bind:     reqBind,
	})
	if err != nil {
		return nil, err
	}
	//生成token
	token, err := jwt.CreateToken(newUser.User.Id, s.conf.Server.Http.JwtSecrect, time.Second*time.Duration(s.conf.Service.TokenExpirationSec))
	if err != nil {
		return nil, err
	}
	tokenClaim, err := jwt.ParseToken(token, s.conf.Server.Http.JwtSecrect)
	if err != nil {
		return nil, err
	}
	// create user online session
	if err = s.data.SessionClient.Create(ctx, &ss.Session{
		Id:         newUser.User.Id,
		UserId:     newUser.User.Id,
		Status:     int32(newUser.User.Status),
		Attributes: make(map[string]string),
		CreatedAt:  tokenClaim.CreatedAt,
	}); err != nil {
		return nil, err
	}

	return &api.RegisterReply{
		Token:      token,
		Expiration: 0,
		UserId:     newUser.User.Id,
	}, nil
}

func (s *AuthService) GetTokenByBind(ctx context.Context, req *api.GetTokenRequest) (*api.GetTokenReply, error) {
	if req.Bind == nil {
		return nil, errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	reqBind := &innterapi.Bind{
		Platform: req.Bind.Platform,
		UserId:   req.Bind.UserId,
		UserName: req.Bind.UserName,
	}
	reply, err := s.data.UserClient.FindUser(ctx, &innterapi.FindUserRequest{
		Bind: reqBind,
	})
	if err != nil {
		return nil, err
	}
	//已绑定返回token,未绑定返回空
	if reply.User != nil {
		token, err := jwt.CreateToken(reply.User.Id, s.conf.Server.Http.JwtSecrect, time.Second*time.Duration(s.conf.Service.TokenExpirationSec))
		if err != nil {
			return nil, err
		}
		tokenClaim, err := jwt.ParseToken(token, s.conf.Server.Http.JwtSecrect)
		if err != nil {
			return nil, err
		}
		// create user online session
		if err = s.data.SessionClient.Create(ctx, &ss.Session{
			Id:         reply.User.Id,
			UserId:     reply.User.Id,
			Status:     int32(reply.User.Status),
			Attributes: make(map[string]string),
			CreatedAt:  tokenClaim.CreatedAt,
		}); err != nil {
			return nil, err
		}

		return &api.GetTokenReply{
			Token:      token,
			Expiration: 0,
		}, nil
	} else {
		return &api.GetTokenReply{
			Token:      "",
			Expiration: 0,
		}, nil
	}
}
