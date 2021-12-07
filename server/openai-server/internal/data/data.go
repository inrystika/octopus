package data

import (
	"context"
	api "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"
	"server/common/middleware/ctxcopy"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data/session"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type Data struct {
	log                *log.Helper
	AlgorithmClient    api.AlgorithmServiceClient
	UserClient         api.UserServiceClient
	TrainJobClient     api.TrainJobServiceClient
	DevelopClient      api.DevelopClient
	ModelClient        api.ModelServiceClient
	WorkspaceClient    api.WorkspaceServiceClient
	SessionClient      session.SessionClient
	ImageClient        api.ImageServiceClient
	DatasetClient      api.DatasetServiceClient
	ResourceSpecClient api.ResourceSpecServiceClient
	ResourcePoolClient api.ResourcePoolServiceClient
	BillingClient      api.BillingServiceClient
	LableClient        api.LableServiceClient
	JointCloudClient   api.JointCloudServiceClient
	ModelDeployClient  api.ModelDeployServiceClient
}

func NewData(confData *conf.Data, logger log.Logger) (*Data, error) {
	log := log.NewHelper("Data", logger)
	baseServerRequestTimeout, err := time.ParseDuration(confData.BaseServerRequestTimeout)
	if err != nil {
		return nil, err
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
		UserClient:         api.NewUserServiceClient(conn),
		DevelopClient:      api.NewDevelopClient(conn),
		TrainJobClient:     api.NewTrainJobServiceClient(conn),
		ModelClient:        api.NewModelServiceClient(conn),
		WorkspaceClient:    api.NewWorkspaceServiceClient(conn),
		SessionClient:      session.NewSessionClient(confData, logger),
		ImageClient:        api.NewImageServiceClient(conn),
		DatasetClient:      api.NewDatasetServiceClient(conn),
		ResourceSpecClient: api.NewResourceSpecServiceClient(conn),
		ResourcePoolClient: api.NewResourcePoolServiceClient(conn),
		BillingClient:      api.NewBillingServiceClient(conn),
		LableClient:        api.NewLableServiceClient(conn),
		JointCloudClient:   api.NewJointCloudServiceClient(conn),
		ModelDeployClient:  api.NewModelDeployServiceClient(conn),
	}, nil
}
