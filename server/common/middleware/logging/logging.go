package logging

import (
	"context"
	"fmt"
	"server/common/log"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// Option is HTTP logging option.
type Option func(*options)

type options struct {
	logger log.Logger
}

// WithLogger with middleware logger.
func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

// Server is an server logging middleware.
func Server(opts ...Option) middleware.Middleware {
	options := options{
		logger: log.DefaultLogger,
	}
	for _, o := range opts {
		o(&options)
	}
	log := log.NewHelper("middleware/logging", options.logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				path    string
				request string
			)

			if info, ok := http.FromServerContext(ctx); ok {
				path = info.Request.RequestURI
				request = info.Request.Form.Encode()
			} else if info, ok := grpc.FromServerContext(ctx); ok {
				path = info.FullMethod
				request = req.(fmt.Stringer).String()
			}
			reply, err := handler(ctx, req)
			if err != nil {
				log.Errorw(ctx,
					"interface", path,
					"request", request,
					"error", err.Error(),
				)
				return nil, err
			}
			log.Infow(ctx,
				"interface", path,
				"request", request,
			)
			log.Debugw(ctx,
				"interface", path,
				"reply", reply)
			return reply, nil
		}
	}
}

// Client is an client logging middleware.
func Client(opts ...Option) middleware.Middleware {
	options := options{
		logger: log.DefaultLogger,
	}
	for _, o := range opts {
		o(&options)
	}
	log := log.NewHelper("middleware/logging", options.logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			var (
				path    string
				request string
			)
			if info, ok := http.FromClientContext(ctx); ok {
				path = info.Request.RequestURI
				request = info.Request.Form.Encode()
			} else if info, ok := grpc.FromClientContext(ctx); ok {
				path = info.FullMethod
				request = req.(fmt.Stringer).String()
			}
			reply, err := handler(ctx, req)
			if err != nil {
				log.Errorw(ctx,
					"path", path,
					"request", request,
					"error", err.Error(),
				)
				return nil, err
			}
			log.Infow(ctx,
				"interface", path,
				"request", request,
			)
			log.Debugw(ctx,
				"interface", path,
				"reply", reply)
			return reply, nil
		}
	}
}
