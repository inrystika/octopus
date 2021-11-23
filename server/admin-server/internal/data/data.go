package data

import (
	"context"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data/session"
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
	log                    *log.Helper
	AlgorithmClient        api.AlgorithmClient
	AdminUserClient        api.AdminUserClient
	UserClient             api.UserClient
	DevelopClient          api.DevelopClient
	WorkspaceClient        api.WorkspaceClient
	ModelClient            api.ModelClient
	TrainJobClient         api.TrainJobServiceClient
	ImageClient            api.ImageClient
	NodeClient             api.NodeServiceClient
	ResourceClient         api.ResourceServiceClient
	ResourceSpecClient     api.ResourceSpecServiceClient
	ResourcePoolClient     api.ResourcePoolServiceClient
	DatasetClient          api.DatasetServiceClient
	SessionClient          session.SessionClient
	BillingClient          api.BillingServiceClient
	PlatformClient         api.PlatformServiceClient
	JointCloudClient       api.JointCloudServiceClient
	PlatformTrainJobClient api.PlatformTrainJobServiceClient
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
		log:                    log,
		AlgorithmClient:        api.NewAlgorithmClient(conn),
		AdminUserClient:        api.NewAdminUserClient(conn),
		UserClient:             api.NewUserClient(conn),
		DevelopClient:          api.NewDevelopClient(conn),
		WorkspaceClient:        api.NewWorkspaceClient(conn),
		ModelClient:            api.NewModelClient(conn),
		TrainJobClient:         api.NewTrainJobServiceClient(conn),
		ImageClient:            api.NewImageClient(conn),
		NodeClient:             api.NewNodeServiceClient(conn),
		ResourceClient:         api.NewResourceServiceClient(conn),
		ResourceSpecClient:     api.NewResourceSpecServiceClient(conn),
		ResourcePoolClient:     api.NewResourcePoolServiceClient(conn),
		DatasetClient:          api.NewDatasetServiceClient(conn),
		SessionClient:          session.NewSessionClient(confData, logger),
		BillingClient:          api.NewBillingServiceClient(conn),
		PlatformClient:         api.NewPlatformServiceClient(conn),
		JointCloudClient:       api.NewJointCloudServiceClient(conn),
		PlatformTrainJobClient: api.NewPlatformTrainJobServiceClient(conn),
	}, nil
}
