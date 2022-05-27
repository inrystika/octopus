package service

import (
	"context"
	innerapi "server/base-server/api/v1"
	"server/common/errors"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/data"

	"github.com/jinzhu/copier"
)

type platformStatisticsService struct {
	api.UnimplementedPlatformStatisticsServer
	data *data.Data
}

func NewPlatformStatisticsService(data *data.Data) api.PlatformStatisticsServer {
	return &platformStatisticsService{
		data: data,
	}
}

func (s *platformStatisticsService) Summary(ctx context.Context, req *api.PlatformStatSummaryRequest) (*api.PlatformStatSummaryReply, error) {
	innerReply, err := s.data.PlatformStatisticsClient.Summary(ctx, &innerapi.PlatformStatSummaryRequest{})
	if err != nil {
		return nil, err
	}

	reply := &api.PlatformStatSummaryReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	return reply, nil
}
