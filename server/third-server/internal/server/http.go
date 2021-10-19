package server

import (
	nethttp "net/http"
	comHttp "server/common/http"
	"server/common/middleware/logging"
	"server/common/middleware/validate"
	"server/third-server/internal/conf"
	"server/third-server/internal/server/middleware/auth"
	"server/third-server/internal/service"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
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

	osrv := newOauthServer(c)

	handleOptions := comHttp.NewHandleOptions()
	_ = []http.HandleOption{
		http.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				tracing.Server(),
				logging.Server(),
				auth.Server(func(options *auth.Options) {
					options.NoAuthUris = []string{"/v1/oauth/token"}
					options.Server = osrv
				}),
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
