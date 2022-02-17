package server

import (
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/service"
	comHttp "server/common/http"
	"server/common/middleware/jwt"
	"server/common/middleware/logging"
	"server/common/middleware/validate"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

const (
	AUTH_TOKEN_URL = "/v1/authmanage/token"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, service *service.Service) *http.Server {
	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	var jwtOpts = []jwt.Option{}
	jwtOpts = append(jwtOpts, func(options *jwt.Options) {
		options.Secret = c.Http.JwtSecrect
		options.NoAuthUris = []string{AUTH_TOKEN_URL}
	})

	handleOptions := comHttp.NewHandleOptions()
	options := []http.HandleOption{
		http.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				tracing.Server(),
				logging.Server(),
				jwt.Server(jwtOpts...),
				validate.Server(),
			),
		),
		http.RequestDecoder(handleOptions.DecodeRequest),
		http.ResponseEncoder(handleOptions.EncodeResponse),
		http.ErrorEncoder(handleOptions.EncodeError),
	}

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/v1/authmanage", api.NewAuthHandler(service.AuthService, options...))
	srv.HandlePrefix("/v1/algorithmmanage", api.NewAlgorithmHandler(service.AlgorithmService, options...))
	srv.HandlePrefix("/v1/usermanage/user", api.NewUserHandler(service.UserService, options...))
	srv.HandlePrefix("/v1/usermanage/workspace", api.NewWorkspaceHandler(service.WorkspaceService, options...))
	srv.HandlePrefix("/v1/developmanage", api.NewDevelopHandler(service.DevelopService, options...))
	srv.HandlePrefix("/v1/modelmanage", api.NewModelHandler(service.ModelService, options...))
	srv.HandlePrefix("/v1/trainmanage", api.NewTrainJobServiceHandler(service.TrainJobService, options...))
	srv.HandlePrefix("/v1/resourcemanage/node", api.NewNodeServiceHandler(service.NodeService, options...))
	srv.HandlePrefix("/v1/resourcemanage/resourcespec", api.NewResourceSpecServiceHandler(service.ResourceSpecService, options...))
	srv.HandlePrefix("/v1/resourcemanage/resourcepool", api.NewResourcePoolServiceHandler(service.ResourcePoolService, options...))
	srv.HandlePrefix("/v1/resourcemanage/resource", api.NewResourceServiceHandler(service.ResourceService, options...))
	srv.HandlePrefix("/v1/datasetmanage", api.NewDatasetServiceHandler(service.DatasetService, options...))
	srv.HandlePrefix("/v1/imagemanage", api.NewImageServiceHandler(service.ImageService, options...))
	srv.HandlePrefix("/v1/billingmanage", api.NewBillingServiceHandler(service.BillingService, options...))
	srv.HandlePrefix("/v1/platformmanage/platform", api.NewPlatformServiceHandler(service.PlatformService, options...))
	srv.HandlePrefix("/v1/jointcloudmanage", api.NewJointCloudServiceHandler(service.JointCloudService, options...))
	srv.HandlePrefix("/v1/deploymanage", api.NewModelDeployServiceHandler(service.ModelDeployService, options...))
	return srv
}
