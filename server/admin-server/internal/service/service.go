package service

import (
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	"server/common/log"
)

type Service struct {
	AlgorithmService    api.AlgorithmServer
	AuthService         api.AuthServer
	UserService         api.UserServer
	WorkspaceService    api.WorkspaceServer
	DevelopService      api.DevelopServer
	ModelService        api.ModelServer
	TrainJobService     api.TrainJobServiceServer
	ImageService        api.ImageServiceServer
	NodeService         api.NodeServiceServer
	ResourceService     api.ResourceServiceServer
	ResourceSpecService api.ResourceSpecServiceServer
	ResourcePoolService api.ResourcePoolServiceServer
	DatasetService      api.DatasetServiceServer
	BillingService      api.BillingServiceServer
}

func NewService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) *Service {
	service := &Service{}
	service.AlgorithmService = NewAlgorithmService(conf, logger, data)
	service.AuthService = NewAuthService(conf, logger, data)
	service.UserService = NewUserService(conf, logger, data)
	service.WorkspaceService = NewWorkspaceService(conf, logger, data)
	service.DevelopService = NewDevelopService(conf, logger, data)
	service.ModelService = NewModelService(conf, logger, data)
	service.TrainJobService = NewTrainJobService(conf, logger, data)
	service.ImageService = NewImageService(conf, logger, data)
	service.NodeService = NewNodeService(conf, logger, data)
	service.ResourceService = NewResourceService(conf, logger, data)
	service.ResourceSpecService = NewResourceSpecService(conf, logger, data)
	service.ResourcePoolService = NewResourcePoolService(conf, logger, data)
	service.DatasetService = NewDatasetService(conf, logger, data)
	service.BillingService = NewBillingService(conf, logger, data)

	return service
}
