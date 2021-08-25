package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	nethttp "net/http"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/base-server/internal/service"
	"server/common/middleware/logging"

	"server/common/log"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, service *service.Service) *http.Server {
	var opts = []http.ServerOption{}

	http.Middleware(
		middleware.Chain(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(),
		),
	)

	//http.WithTimeout(time.Minute *2)

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)

	//实现pipeline回调 request和response编码是pipeline要求的格式 先单独写
	srv.HandleFunc("/v1/developmanage/pipelinecallback", func(w nethttp.ResponseWriter, r *nethttp.Request) {
		pipelineCallback(w, r, service.DevelopService)
	})

	srv.HandleFunc("/v1/trainmanage/pipelinecallback", func(w nethttp.ResponseWriter, r *nethttp.Request) {
		pipelineCallback(w, r, service.TrainJobService)
	})
	return srv
}

func pipelineCallback(w nethttp.ResponseWriter, r *nethttp.Request, callback common.PipelineCallback) {
	ctx := context.TODO()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, err := w.Write([]byte(common.PipeLineCallbackRE))
		if err != nil {
			log.Error(ctx, err)
		}
	}

	p := &common.PipelineCallbackReq{}
	err = json.Unmarshal(data, p)
	if err != nil {
		_, err := w.Write([]byte(common.PipeLineCallbackRE))
		if err != nil {
			log.Error(ctx, err)
		}
	}

	res := callback.PipelineCallback(context.Background(), p)
	_, err = w.Write([]byte(res))
	if err != nil {
		log.Error(ctx, err)
	}
}
