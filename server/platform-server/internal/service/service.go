package service

import (
	api "server/platform-server/api/v1"
	"server/platform-server/internal/conf"
	"server/platform-server/internal/data"
)

type Service struct {
	Data            *data.Data
	OauthService    OauthService
	TrainJobService api.TrainJobServiceServer
}

func NewService(conf *conf.Bootstrap, data *data.Data) *Service {
	service := &Service{
		OauthService:    NewOauthService(conf, data),
		TrainJobService: NewTrainJobService(conf, data),
	}
	return service
}
