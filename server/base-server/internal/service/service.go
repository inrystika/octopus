package service

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/service/algorithm"
	"server/base-server/internal/service/billing"
	"server/base-server/internal/service/dataset"
	"server/base-server/internal/service/develop"
	"server/base-server/internal/service/image"
	"server/base-server/internal/service/jointcloud"
	"server/base-server/internal/service/model"
	"server/base-server/internal/service/platform"
	"server/base-server/internal/service/resources"
	"server/base-server/internal/service/trainjob"
	"server/base-server/internal/service/user"
	"server/base-server/internal/service/workspace"

	"server/common/log"
)

type Service struct {
	AlgorithmService        api.AlgorithmServer
	UserService             api.UserServer
	AdminUserService        api.AdminUserServer
	ModelService            api.ModelServer
	ResourceService         api.ResourceServiceServer
	ResourceSpecService     api.ResourceSpecServiceServer
	ResourcePoolService     api.ResourcePoolServiceServer
	NodeService             api.NodeServiceServer
	DevelopService          develop.DevelopService
	TrainJobService         trainjob.TrainJobService
	WorkspaceService        api.WorkspaceServer
	DatasetService          api.DatasetServiceServer
	ImageService            api.ImageServer
	BillingService          api.BillingServiceServer
	PlatformService         api.PlatformServiceServer
	PlatformTrainJobService api.PlatformTrainJobServiceServer
	JointCloudService       api.JointCloudServiceServer
}

func NewService(ctx context.Context, conf *conf.Bootstrap, logger log.Logger, data *data.Data) (*Service, error) {
	var err error
	service := &Service{}

	service.AlgorithmService = algorithm.NewAlgorithmService(conf, logger, data)
	service.UserService = user.NewUserService(conf, logger, data)
	service.AdminUserService = user.NewAdminUserService(conf, logger, data)

	service.ResourceService = resources.NewResourceService(ctx, conf, logger, data)
	service.ResourceSpecService = resources.NewResourceSpecService(conf, logger, data)
	service.ResourcePoolService = resources.NewResourcePoolService(conf, logger, data)
	service.NodeService = resources.NewNodeService(conf, logger, data)
	service.ModelService = model.NewModelService(conf, logger, data, service.AlgorithmService)
	service.WorkspaceService = workspace.NewWorkspaceService(conf, logger, data)
	service.DatasetService = dataset.NewDatasetService(conf, logger, data)
	service.ImageService = image.NewImageService(conf, logger, data)
	service.BillingService = billing.NewBillingService(conf, logger, data)
	service.DevelopService, err = develop.NewDevelopService(conf, logger, data,
		service.WorkspaceService, service.AlgorithmService, service.ImageService, service.DatasetService,
		service.ResourceSpecService, service.ResourceService, service.ResourcePoolService, service.BillingService)
	if err != nil {
		return nil, err
	}
	service.TrainJobService, err = trainjob.NewTrainJobService(conf, logger, data,
		service.WorkspaceService, service.AlgorithmService, service.ImageService, service.DatasetService,
		service.ModelService, service.ResourceSpecService, service.ResourceService, service.ResourcePoolService, service.BillingService)
	if err != nil {
		return nil, err
	}
	service.PlatformService = platform.NewPlatformService(conf, data)
	service.PlatformTrainJobService, err = platform.NewPlatformTrainJobService(conf, logger, data,
		service.ResourceService, service.PlatformService)
	if err != nil {
		return nil, err
	}
	service.JointCloudService = jointcloud.NewJointCloudService(conf, data)

	return service, nil
}
