package platformstatistics

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/data"
	"server/common/errors"

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
	summaryTbl, err := s.data.PlatformStatisticsDao.Summary()
	if err != nil {
		return nil, err
	}
	reply := &api.PlatformStatSummaryReply{}
	err = copier.Copy(reply, summaryTbl)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	return reply, nil
}
