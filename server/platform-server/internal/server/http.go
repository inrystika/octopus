package server

import (
	"context"
	nethttp "net/http"
	commctx "server/common/context"
	"server/common/errors"
	comHttp "server/common/http"
	"server/common/middleware/logging"
	"server/common/middleware/validate"
	"server/third-server/internal/conf"
	"server/third-server/internal/service"
	"strings"

	oserver "github.com/go-oauth2/oauth2/v4/server"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

const (
	authHeader = "Authorization"
	authType   = "Bearer"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Bootstrap, service *service.Service) *http.Server {
	var opts = []http.ServerOption{}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}

	osrv := newOauthServer(c, service)

	handleOptions := comHttp.NewHandleOptions()
	_ = []http.HandleOption{
		http.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				tracing.Server(),
				logging.Server(),
				authMiddleware([]string{"/v1/oauth/token"}, osrv),
				validate.Server(),
			),
		),
		http.RequestDecoder(handleOptions.DecodeRequest),
		http.ResponseEncoder(handleOptions.EncodeResponse),
		http.ErrorEncoder(handleOptions.EncodeError),
	}

	srv := http.NewServer(opts...)
	srv.HandleFunc("/v1/oauth/token", func(w nethttp.ResponseWriter, r *nethttp.Request) {
		token(w, r, osrv)
	})
	return srv
}

func authMiddleware(noAuthUris []string, server *oserver.Server) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var request *nethttp.Request
			if info, ok := http.FromServerContext(ctx); ok {
				request = info.Request
			} else {
				return handler(ctx, req)
			}

			needAuth := true
			for _, i := range noAuthUris {
				if strings.Contains(request.RequestURI, i) {
					needAuth = false
					break
				}
			}

			if needAuth {
				authorization := request.Header.Get(authHeader)
				if authorization == "" {
					return nil, errors.Errorf(nil, errors.ErrorNotAuthorized)
				}

				if strings.Index(authorization, authType) != 0 || authorization[0:len(authType)] != authType {
					return nil, errors.Errorf(nil, errors.ErrorTokenInvalid)
				}

				tokenInfo, err := server.Manager.LoadAccessToken(ctx, authorization[len(authType)+1:])
				if err != nil {
					return nil, err
				}
				ctx = commctx.PlatformIdToContext(ctx, tokenInfo.GetClientID())
			}

			return handler(ctx, req)
		}
	}
}
