package server

import (
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/service"
	"server/common/errors"
	"server/common/middleware/ctxcopy"
	"server/common/middleware/logging"
	"server/common/middleware/validate"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, s *service.Service) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				ctxcopy.Server(),
				status.Server(status.WithHandler(errors.ErrorEncode)),
				tracing.Server(),
				logging.Server(),
				validate.Server(),
			),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}

	gs := grpc.NewServer(opts...)
	api.RegisterAlgorithmServiceServer(gs, s.AlgorithmService)
	api.RegisterUserServiceServer(gs, s.UserService)
	api.RegisterAdminUserServer(gs, s.AdminUserService)
	api.RegisterDevelopServer(gs, s.DevelopService)
	api.RegisterResourceServiceServer(gs, s.ResourceService)
	api.RegisterResourceSpecServiceServer(gs, s.ResourceSpecService)
	api.RegisterResourcePoolServiceServer(gs, s.ResourcePoolService)
	api.RegisterNodeServiceServer(gs, s.NodeService)
	api.RegisterModelServiceServer(gs, s.ModelService)
	api.RegisterTrainJobServiceServer(gs, s.TrainJobService)
	api.RegisterWorkspaceServiceServer(gs, s.WorkspaceService)
	api.RegisterDatasetServiceServer(gs, s.DatasetService)
	api.RegisterImageServiceServer(gs, s.ImageService)
	api.RegisterBillingServiceServer(gs, s.BillingService)

	return gs
}
