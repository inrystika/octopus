package server

import (
	"context"
	"fmt"
	nethttp "net/http"
	"net/http/httputil"
	"net/url"
	"server/common/constant"
	"server/common/constant/userconfig"
	commctx "server/common/context"
	"server/common/errors"
	comHttp "server/common/http"
	"server/common/log"
	"server/common/middleware/jwt"
	"server/common/middleware/logging"
	"server/common/middleware/validate"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/service"
	"strings"

	"github.com/gorilla/mux"

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

	noAuthUris := []string{"/v1/authmanage/token", "/v1/authmanage/registerandbind", "/v1/systemmanage/webconfig"}
	var jwtOpts = []jwt.Option{}
	jwtOpts = append(jwtOpts, func(options *jwt.Options) {
		options.Secret = c.Http.JwtSecrect
		options.NoAuthUris = noAuthUris
	})

	options := []http.HandleOption{
		http.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				tracing.Server(),
				logging.Server(),
				jwt.Server(jwtOpts...),
				handleHeader(),
				checkJointCloudPerm(service),
				validate.Server(),
			),
		),
		http.RequestDecoder(comHttp.DecodeRequest),
		http.ResponseEncoder(comHttp.EncodeResponse),
		http.ErrorEncoder(comHttp.EncodeError),
	}

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/v1/oauth2", NewOauthHandler(c, context.Background(), service))
	srv.HandlePrefix("/v1/usermanage", api.NewUserHandler(service.UserService, options...))
	srv.HandlePrefix("/v1/authmanage", api.NewAuthHandler(service.AuthService, options...))
	srv.HandlePrefix("/v1/algorithmmanage", api.NewAlgorithmHandler(service.AlgorithmService, options...))
	srv.HandlePrefix("/v1/developmanage", api.NewDevelopHandler(service.DevelopService, options...))
	srv.HandleFunc("/v1/trainmanage/trainjob/{id}/task/{taskId}/replica/{replicaIdx}/log", func(w nethttp.ResponseWriter, r *nethttp.Request) {
		proto := "http"
		if strings.Contains(r.Proto, "HTTPS") {
			proto = "https"
		}
		url, _ := url.Parse(fmt.Sprintf("%s://%s/log/user/trainjob/%s/%s/%s/index.log", proto, r.Host, mux.Vars(r)["id"], mux.Vars(r)["taskId"], mux.Vars(r)["replicaIdx"]))
		log.Info(context.TODO(), url)
		proxy := httputil.ReverseProxy{
			Director: func(request *nethttp.Request) {
				request.URL = url
			},
		}

		proxy.ServeHTTP(w, r)
	})
	srv.HandlePrefix("/v1/trainmanage", api.NewTrainJobServiceHandler(service.TrainJobService, options...))
	srv.HandlePrefix("/v1/modelmanage", api.NewModelHandler(service.ModelService, options...))
	srv.HandlePrefix("/v1/datasetmanage", api.NewDatasetServiceHandler(service.DatasetService, options...))
	srv.HandlePrefix("/v1/resourcemanage/resourcespec", api.NewResourceSpecServiceHandler(service.ResourceSpecService, options...))
	srv.HandlePrefix("/v1/imagemanage", api.NewImageServiceHandler(service.ImageService, options...))
	srv.HandlePrefix("/v1/billingmanage", api.NewBillingServiceHandler(service.BillingService, options...))
	srv.HandlePrefix("/v1/jointcloudmanage", api.NewJointCloudServiceHandler(service.JointCloudService, options...))
	srv.HandlePrefix("/v1/systemmanage", api.NewSystemServiceHandler(service.SystemService, options...))
	srv.HandlePrefix("/v1/deploymanage", api.NewModelDeployServiceHandler(service.ModelDeployService, options...))
	srv.HandlePrefix("/v1/platformstatisticsmanage", api.NewPlatformStatisticsHandler(service.PlatformStatisticsService, options...))
	return srv
}

func checkJointCloudPerm(service *service.Service) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var request *nethttp.Request
			if info, ok := http.FromServerContext(ctx); ok {
				request = info.Request
			} else {
				return handler(ctx, req)
			}

			if strings.Contains(request.RequestURI, "/v1/jointcloudmanage") {
				config, err := service.UserService.GetUserConfig(ctx, &api.GetUserConfigRequest{})
				if err != nil {
					return nil, err
				}
				if config.Config == nil || !strings.EqualFold(config.Config[userconfig.JointCloudPermission], userconfig.JointCloudPermissionYes) {
					return nil, errors.Errorf(nil, errors.ErrorJointCloudNoPermission)
				}
			}

			return handler(ctx, req)
		}
	}
}

func handleHeader() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var request *nethttp.Request
			if info, ok := http.FromServerContext(ctx); ok {
				request = info.Request
			} else {
				return handler(ctx, req)
			}

			spaceId := request.Header.Get("Octopus-Space-Id")
			if spaceId == "" {
				spaceId = constant.SYSTEM_WORKSPACE_DEFAULT
			}

			ctx = commctx.SpaceIdToContext(ctx, spaceId)
			return handler(ctx, req)
		}
	}
}
