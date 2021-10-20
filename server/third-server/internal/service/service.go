package service

import (
	"server/third-server/internal/conf"
	"server/third-server/internal/data"
)

type Service struct {
	Data         *data.Data
	OauthService OauthService
}

func NewService(conf *conf.Bootstrap, data *data.Data) *Service {
	service := &Service{
		OauthService: NewOauthService(conf, data),
	}
	return service
}
