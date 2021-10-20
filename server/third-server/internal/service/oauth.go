package service

import (
	"context"
	innerapi "server/base-server/api/v1"
	"server/common/errors"
	"server/third-server/internal/conf"
	"server/third-server/internal/data"

	"github.com/go-oauth2/oauth2/v4/models"

	"github.com/go-oauth2/oauth2/v4"
)

type OauthService interface {
	oauth2.ClientStore
}

type oauthService struct {
	conf *conf.Bootstrap
	data *data.Data
}

func NewOauthService(conf *conf.Bootstrap, data *data.Data) OauthService {
	return &oauthService{
		conf: conf,
		data: data,
	}
}

func (s *oauthService) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	platformReply, err := s.data.PlatformClient.BatchGetPlatform(ctx, &innerapi.BatchGetPlatformRequest{Ids: []string{id}})
	if err != nil {
		return nil, err
	}

	if len(platformReply.Platforms) <= 0 {
		return nil, errors.Errorf(err, errors.ErrorDBFindEmpty)
	}

	platform := platformReply.Platforms[0]
	return &models.Client{
		ID:     platform.Id,
		Secret: platform.ClientSecret,
		Domain: "",
		UserID: "",
	}, nil
}
