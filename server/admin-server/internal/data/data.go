package data

import (
	"context"
	"server/admin-server/internal/conf"
	api "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"
	"server/common/middleware/ctxcopy"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type Data struct {
	log                *log.Helper
	AlgorithmClient    api.AlgorithmServiceClient
	AdminUserClient    api.AdminUserClient
	UserClient         api.UserServiceClient
	DevelopClient      api.DevelopClient
	WorkspaceClient    api.WorkspaceServiceClient
	ModelClient        api.ModelServiceClient
	TrainJobClient     api.TrainJobServiceClient
	ImageClient        api.ImageServiceClient
	NodeClient         api.NodeServiceClient
	ResourceClient     api.ResourceServiceClient
	ResourceSpecClient api.ResourceSpecServiceClient
	ResourcePoolClient api.ResourcePoolServiceClient
	DatasetClient      api.DatasetServiceClient
	BillingClient      api.BillingServiceClient
	LableClient        api.LableServiceClient
	ModelDeployClient  api.ModelDeployServiceClient
}

func NewData(confData *conf.Data, logger log.Logger) (*Data, error) {
	log := log.NewHelper("data", logger)
	baseServerRequestTimeout, err := time.ParseDuration(confData.BaseServerRequestTimeout)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorParseDurationFailed)
	}

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(confData.BaseServerAddr),
		grpc.WithTimeout(baseServerRequestTimeout),
		grpc.WithMiddleware(
			middleware.Chain(
				status.Client(status.WithHandler(errors.ErrorDecode)),
				recovery.Recovery(),
				ctxcopy.Client(),
			),
		),
	)
	if err != nil {
		return nil, err
	}

	return &Data{
		log:                log,
		AlgorithmClient:    api.NewAlgorithmServiceClient(conn),
		AdminUserClient:    api.NewAdminUserClient(conn),
		UserClient:         api.NewUserServiceClient(conn),
		DevelopClient:      api.NewDevelopClient(conn),
		WorkspaceClient:    api.NewWorkspaceServiceClient(conn),
		ModelClient:        api.NewModelServiceClient(conn),
		TrainJobClient:     api.NewTrainJobServiceClient(conn),
		ImageClient:        api.NewImageServiceClient(conn),
		NodeClient:         api.NewNodeServiceClient(conn),
		ResourceClient:     api.NewResourceServiceClient(conn),
		ResourceSpecClient: api.NewResourceSpecServiceClient(conn),
		ResourcePoolClient: api.NewResourcePoolServiceClient(conn),
		DatasetClient:      api.NewDatasetServiceClient(conn),
		BillingClient:      api.NewBillingServiceClient(conn),
		LableClient:        api.NewLableServiceClient(conn),
		ModelDeployClient:  api.NewModelDeployServiceClient(conn),
	}, nil
}
