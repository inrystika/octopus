package server

import (
	comHttp "server/common/http"
	"server/common/middleware/logging"
	"server/common/middleware/validate"
	"server/third-server/internal/conf"
	"server/third-server/internal/service"

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

	handleOptions := comHttp.NewHandleOptions()
	_ = []http.HandleOption{
		http.Middleware(
			middleware.Chain(
				recovery.Recovery(),
				tracing.Server(),
				logging.Server(),
				validate.Server(),
			),
		),
		http.RequestDecoder(handleOptions.DecodeRequest),
		http.ResponseEncoder(handleOptions.EncodeResponse),
		http.ErrorEncoder(handleOptions.EncodeError),
	}

	srv := http.NewServer(opts...)
	return srv
}
