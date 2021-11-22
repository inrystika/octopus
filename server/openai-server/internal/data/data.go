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
	AlgorithmClient    api.AlgorithmClient
	UserClient         api.UserClient
	TrainJobClient     api.TrainJobServiceClient
	DevelopClient      api.DevelopClient
	ModelClient        api.ModelClient
	WorkspaceClient    api.WorkspaceClient
	SessionClient      session.SessionClient
	ImageClient        api.ImageClient
	DatasetClient      api.DatasetServiceClient
	ResourceSpecClient api.ResourceSpecServiceClient
	ResourcePoolClient api.ResourcePoolServiceClient
	BillingClient      api.BillingServiceClient
	JointCloudClient   api.JointCloudServiceClient
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
		AlgorithmClient:    api.NewAlgorithmClient(conn),
		UserClient:         api.NewUserClient(conn),
		DevelopClient:      api.NewDevelopClient(conn),
		TrainJobClient:     api.NewTrainJobServiceClient(conn),
		ModelClient:        api.NewModelClient(conn),
		WorkspaceClient:    api.NewWorkspaceClient(conn),
		SessionClient:      session.NewSessionClient(confData, logger),
		ImageClient:        api.NewImageClient(conn),
		DatasetClient:      api.NewDatasetServiceClient(conn),
		ResourceSpecClient: api.NewResourceSpecServiceClient(conn),
		ResourcePoolClient: api.NewResourcePoolServiceClient(conn),
		BillingClient:      api.NewBillingServiceClient(conn),
		JointCloudClient:   api.NewJointCloudServiceClient(conn),
	}, nil
}
