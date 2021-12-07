package service

import (
	"context"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"
)

type systemService struct {
	api.UnimplementedSystemServiceServer
	conf *conf.Bootstrap
	data *data.Data
}

func NewSystemService(conf *conf.Bootstrap, data *data.Data) api.SystemServiceServer {
	return &systemService{
		conf: conf,
		data: data,
	}
}

func (s *systemService) GetWebConfig(ctx context.Context, req *api.GetWebConfigRequest) (*api.GetWebConfigReply, error) {
	return &api.GetWebConfigReply{
		LogoAddr:     s.conf.Service.WebConfig.LogoAddr,
		ThemeColor:   s.conf.Service.WebConfig.ThemeColor,
		SystemNameEn: s.conf.Service.WebConfig.SystemNameEn,
		SystemNameZh: s.conf.Service.WebConfig.SystemNameZh,
		Organization: s.conf.Service.WebConfig.Organization,
	}, nil
}
