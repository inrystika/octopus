package platform

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	model "server/base-server/internal/data/dao/model/platform"
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

	_, err = s.data.Cluster.CreateNamespace(ctx, platform.Id)
	if err != nil {
		return nil, err
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

func (s *platformService) ListPlatformConfigKey(ctx context.Context, req *api.ListPlatformConfigKeyRequest) (*api.ListPlatformConfigKeyReply, error) {
	return &api.ListPlatformConfigKeyReply{ConfigKeys: common.PlatformConfigKeys}, nil
}

func (s *platformService) ListPlatformStorageConfig(ctx context.Context, req *api.ListPlatformStorageConfigRequest) (*api.ListPlatformStorageConfigReply, error) {
	query := &model.PlatformStorageConfigQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	platformStorageConfigs, totalSize, err := s.listPlatformStorageConfig(ctx, query)
	if err != nil {
		return nil, err
	}

	return &api.ListPlatformStorageConfigReply{
		TotalSize:              totalSize,
		PlatformStorageConfigs: platformStorageConfigs,
	}, nil
}

func (s *platformService) listPlatformStorageConfig(ctx context.Context, query *model.PlatformStorageConfigQuery) ([]*api.PlatformStorageConfig, int64, error) {
	platformStorageConfigsTbl, totalSize, err := s.data.PlatformDao.ListPlatformStorageConfig(ctx, query)
	if err != nil {
		return nil, 0, err
	}
	platformStorageConfigs := make([]*api.PlatformStorageConfig, 0)
	for _, n := range platformStorageConfigsTbl {
		platformStorageConfig := &api.PlatformStorageConfig{}
		err := copier.CopyWithOption(platformStorageConfig, n, copier.Option{DeepCopy: true})
		if err != nil {
			return nil, 0, errors.Errorf(err, errors.ErrorStructCopy)
		}
		platformStorageConfig.CreatedAt = n.CreatedAt.Unix()
		platformStorageConfig.UpdatedAt = n.UpdatedAt.Unix()
		platformStorageConfigs = append(platformStorageConfigs, platformStorageConfig)
	}

	return platformStorageConfigs, totalSize, nil
}

func (s *platformService) CreatePlatformStorageConfig(ctx context.Context, req *api.CreatePlatformStorageConfigRequest) (*api.CreatePlatformStorageConfigReply, error) {
	platformStorageConfig := &model.PlatformStorageConfig{}
	err := copier.Copy(platformStorageConfig, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	_, size, err := s.data.PlatformDao.ListPlatformStorageConfig(ctx, &model.PlatformStorageConfigQuery{
		PlatformId: req.PlatformId,
		Name:       req.Name,
	})
	if err != nil {
		return nil, err
	}
	if size > 0 {
		return nil, errors.Errorf(nil, errors.ErrorPlatformStorageConfigNameRepeat)
	}

	err = s.data.PlatformDao.CreatePlatformStorageConfig(ctx, platformStorageConfig)
	if err != nil {
		return nil, err
	}

	return &api.CreatePlatformStorageConfigReply{}, nil
}

func (s *platformService) DeletePlatformStorageConfig(ctx context.Context, req *api.DeletePlatformStorageConfigRequest) (*api.DeletePlatformStorageConfigReply, error) {
	err := s.data.PlatformDao.DeletePlatformStorageConfig(ctx, req.PlatformId, req.Name)
	if err != nil {
		return nil, err
	}
	return &api.DeletePlatformStorageConfigReply{}, nil
}

func (s *platformService) GetPlatformStorageConfig(ctx context.Context, req *api.GetPlatformStorageConfigRequest) (*api.GetPlatformStorageConfigReply, error) {
	platformStorageConfigs, _, err := s.listPlatformStorageConfig(ctx, &model.PlatformStorageConfigQuery{PlatformId: req.PlatformId, Name: req.Name})
	if err != nil {
		return nil, err
	}
	if len(platformStorageConfigs) != 1 {
		return nil, errors.Errorf(nil, errors.ErrorDBFindEmpty)
	}

	return &api.GetPlatformStorageConfigReply{PlatformStorageConfig: platformStorageConfigs[0]}, nil
}

func (s *platformService) GetPlatformConfig(ctx context.Context, req *api.GetPlatformConfigRequest) (*api.GetPlatformConfigReply, error) {
	config, err := s.data.PlatformDao.GetPlatformConfig(ctx, req.PlatformId)
	if err != nil {
		return nil, err
	}

	return &api.GetPlatformConfigReply{
		Config: config,
	}, nil
}

func (s *platformService) UpdatePlatformConfig(ctx context.Context, req *api.UpdatePlatformConfigRequest) (*api.UpdatePlatformConfigReply, error) {
	for k, v := range req.Config {
		in := false
		for _, i := range common.PlatformConfigKeys {
			if k == i.Key {
				in = true
				err := i.ValidateValue(v)
				if err != nil {
					return nil, err
				}
				break
			}
		}

		if !in {
			return nil, errors.Errorf(nil, errors.ErrorPlatformConfigKeyNotExist)
		}
	}

	err := s.data.PlatformDao.UpdatePlatformConfig(ctx, req.PlatformId, req.Config)
	if err != nil {
		return nil, err
	}

	return &api.UpdatePlatformConfigReply{}, nil
}
