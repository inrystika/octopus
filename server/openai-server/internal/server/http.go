package server

import (
	"context"
	innterapi "server/base-server/api/v1"
	comCtx "server/common/context"
	"server/common/errors"
	comHttp "server/common/http"
	"server/common/middleware/jwt"
	"server/common/middleware/logging"
	"server/common/middleware/session"
	"server/common/middleware/validate"
	ss "server/common/session"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/service"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
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
		options.NoAuthUris = []string{"/v1/authmanage/token"}
	})

	var sessionOpts = []session.Option{}
	sessionOpts = append(sessionOpts, func(options *session.Options) {
		options.Store = service.Data.SessionClient
		options.NoAuthUris = []string{"/v1/authmanage/token"}
		options.CheckSession = func(ctx context.Context, s *ss.Session) error {
			if s.Status != int32(innterapi.UserStatus_ACTIVITY) {
				return errors.Errorf(nil, errors.ErrorAuthenticationForbidden)
			}

			if comCtx.CreatedAtFromContext(ctx) != s.CreatedAt {
				return errors.Errorf(nil, errors.ErrorTokenRenew)
			}
			return nil
		}
	})

	handleOptions := comHttp.NewHandleOptions()
	options := []http.HandleOption{
		http.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				tracing.Server(),
				logging.Server(),
				jwt.Server(jwtOpts...),
				session.Server(sessionOpts...),
				validate.Server(),
			),
		),
		http.RequestDecoder(handleOptions.DecodeRequest),
		http.ResponseEncoder(handleOptions.EncodeResponse),
		http.ErrorEncoder(handleOptions.EncodeError),
	}

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/v1/usermanage", api.NewUserHandler(service.UserService, options...))
	srv.HandlePrefix("/v1/authmanage", api.NewAuthHandler(service.AuthService, options...))
	srv.HandlePrefix("/v1/algorithmmanage", api.NewAlgorithmHandler(service.AlgorithmService, options...))
	srv.HandlePrefix("/v1/developmanage", api.NewDevelopHandler(service.DevelopService, options...))
	srv.HandlePrefix("/v1/trainmanage", api.NewTrainJobServiceHandler(service.TrainJobService, options...))
	srv.HandlePrefix("/v1/modelmanage", api.NewModelHandler(service.ModelService, options...))
	srv.HandlePrefix("/v1/datasetmanage", api.NewDatasetServiceHandler(service.DatasetService, options...))
	srv.HandlePrefix("/v1/resourcemanage/resourcespec", api.NewResourceSpecServiceHandler(service.ResourceSpecService, options...))
	srv.HandlePrefix("/v1/imagemanage", api.NewImageServiceHandler(service.ImageService, options...))
	srv.HandlePrefix("/v1/billingmanage", api.NewBillingServiceHandler(service.BillingService, options...))
	srv.HandlePrefix("/v1/deploymanage", api.NewModelDeployServiceHandler(service.ModelDeployService, options...))
	return srv
}
