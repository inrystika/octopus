package service

import (
	"server/common/log"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"
)

type Service struct {
	Data                *data.Data
	AlgorithmService    api.AlgorithmServer
	AuthService         api.AuthServer
	TrainJobService     api.TrainJobServiceServer
	UserService         api.UserServer
	DevelopService      api.DevelopServer
	ModelService        api.ModelServer
	ImageService        api.ImageServiceServer
	DatasetService      api.DatasetServiceServer
	ResourceSpecService api.ResourceSpecServiceServer
	BillingService      api.BillingServiceServer
	JointCloudService   api.JointCloudServiceServer
	SystemService       api.SystemServiceServer
	ModelDeployService  api.ModelDeployServiceServer
}

func NewService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) *Service {
	service := &Service{}
	service.AlgorithmService = NewAlgorithmService(conf, logger, data)
	service.AuthService = NewAuthService(conf, logger, data)
	service.TrainJobService = NewTrainJobService(conf, logger, data)
	service.UserService = NewUserService(conf, logger, data)
	service.DevelopService = NewDevelopService(conf, logger, data)
	service.ModelService = NewModelService(conf, logger, data)
	service.ImageService = NewImageService(conf, logger, data)
	service.Data = data
	service.DatasetService = NewDatasetService(conf, logger, data)
	service.ResourceSpecService = NewResourceSpecService(conf, logger, data)
	service.BillingService = NewBillingService(conf, logger, data)
	service.JointCloudService = NewJointCloudService(conf, data)
	service.SystemService = NewSystemService(conf, data)
	service.ModelDeployService =  NewModelDeployService(conf, logger, data)
	return service
}
