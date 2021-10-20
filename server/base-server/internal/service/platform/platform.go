package platform

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"
	"server/common/utils"

	"github.com/jinzhu/copier"
)

type platformService struct {
	api.UnimplementedPlatformServiceServer
	conf *conf.Bootstrap
	data *data.Data
}

func NewPlatformService(conf *conf.Bootstrap, data *data.Data) api.PlatformServiceServer {
	s := &platformService{
		conf: conf,
		data: data,
	}

	return s
}

func (s *platformService) ListPlatform(ctx context.Context, req *api.ListPlatformRequest) (*api.ListPlatformReply, error) {
	query := &model.PlatformQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	platforms, totalSize, err := s.listPlatform(ctx, query)
	if err != nil {
		return nil, err
	}

	return &api.ListPlatformReply{
		TotalSize: totalSize,
		Platforms: platforms,
	}, nil
}

func (s *platformService) listPlatform(ctx context.Context, query *model.PlatformQuery) ([]*api.Platform, int64, error) {
	platformsTbl, totalSize, err := s.data.PlatformDao.ListPlatform(ctx, query)
	if err != nil {
		return nil, 0, err
	}
	platforms := make([]*api.Platform, 0)
	for _, n := range platformsTbl {
		platform := &api.Platform{}
		err := copier.Copy(platform, n)
		if err != nil {
			return nil, 0, errors.Errorf(err, errors.ErrorStructCopy)
		}
		platform.CreatedAt = n.CreatedAt.Unix()
		platform.UpdatedAt = n.UpdatedAt.Unix()
		platforms = append(platforms, platform)
	}

	return platforms, totalSize, nil
}

func (s *platformService) BatchGetPlatform(ctx context.Context, req *api.BatchGetPlatformRequest) (*api.BatchGetPlatformReply, error) {
	platforms, _, err := s.listPlatform(ctx, &model.PlatformQuery{Ids: req.Ids})
	if err != nil {
		return nil, err
	}

	return &api.BatchGetPlatformReply{Platforms: platforms}, nil
}

func (s *platformService) CreatePlatform(ctx context.Context, req *api.CreatePlatformRequest) (*api.CreatePlatformReply, error) {
	platform := &model.Platform{}
	err := copier.Copy(platform, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	platform.Id = utils.GetUUIDWithoutSeparator()
	platform.ClientSecret = utils.GetUUIDWithoutSeparator()

	_, size, err := s.data.PlatformDao.ListPlatform(ctx, &model.PlatformQuery{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	if size > 0 {
		return nil, errors.Errorf(nil, errors.ErrorPlatformNameRepeat)
	}

	err = s.data.PlatformDao.CreatePlatform(ctx, platform)
	if err != nil {
		return nil, err
	}

	return &api.CreatePlatformReply{Id: platform.Id}, nil
}
func (s *platformService) UpdatePlatform(ctx context.Context, req *api.UpdatePlatformRequest) (*api.UpdatePlatformReply, error) {
	platform := &model.Platform{}
	err := copier.Copy(platform, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	err = s.data.PlatformDao.UpdatePlatformById(ctx, platform)
	if err != nil {
		return nil, err
	}

	return &api.UpdatePlatformReply{}, nil
}
