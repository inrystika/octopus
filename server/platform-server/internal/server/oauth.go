package server

import (
	"context"
	nethttp "net/http"
	"server/common/log"
	"server/platform-server/internal/conf"
	"server/platform-server/internal/service"
	"time"

	"github.com/go-oauth2/oauth2/v4/generates"

	"github.com/go-oauth2/oauth2/v4/errors"

	"github.com/go-oauth2/oauth2/v4/manage"
	oserver "github.com/go-oauth2/oauth2/v4/server"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
)

func newOauthServer(c *conf.Bootstrap, service *service.Service) *oserver.Server {
	manager := manage.NewDefaultManager()
	manager.SetClientTokenCfg(&manage.Config{AccessTokenExp: time.Second * time.Duration(c.Service.TokenExpirationSec)})

	// token store
	manager.MustTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr:     c.Data.Redis.Addr,
		DB:       0,
		Username: c.Data.Redis.Username,
		Password: c.Data.Redis.Password,
	}), nil)

	manager.MapAccessGenerate(generates.NewAccessGenerate())

	manager.MapClientStorage(service.OauthService)

	srv := oserver.NewServer(oserver.NewConfig(), manager)
	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Info(context.TODO(), "Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Info(context.TODO(), "Response Error:", re.Error.Error())
	})
	srv.SetClientInfoHandler(oserver.ClientFormHandler)
	return srv
}

func token(w nethttp.ResponseWriter, r *nethttp.Request, osrv *oserver.Server) {
	err := osrv.HandleTokenRequest(w, r)
	if err != nil {
		nethttp.Error(w, err.Error(), nethttp.StatusInternalServerError)
	}
}
