package main

import (
	"context"
	"flag"
	"fmt"
	"server/common/log"
	"server/third-server/internal/conf"
	"server/third-server/internal/data"
	"server/third-server/internal/server"
	"server/third-server/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/transport/http"
	"gopkg.in/yaml.v2"
)

var (
	// Name is the name of the compiled software.
	// go build -ldflags "-X main.Name=xyz"
	Name string
	// Version is the version of the compiled software.
	// go build -ldflags "-X main.Version=x.y.z"
	v       bool
	Version string
	// Built is the build time of the compiled software.
	// go build -ldflags "-X main.Built=2021-06-02 17:29:36"
	Built string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.BoolVar(&v, "v", false, "software version, eg: -v")
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	if v {
		fmt.Printf("Version: %s\nBuilt: %s\n", Version, Built)
		return
	}
	if flagconf == "" {
		fmt.Printf("Miss param: -conf, use -h to get help\n")
		return
	}

	conf, c, err := initConf()
	if err != nil {
		panic(err)
	}

	l := log.ConvertFromString(conf.App.LogLevel)
	log.DefaultLogger.ResetLevel(l)
	err = c.Watch("app.logLevel", func(k string, v config.Value) {
		ls, _ := v.String()
		l := log.ConvertFromString(ls)
		log.DefaultLogger.ResetLevel(l)
	})
	if err != nil {
		log.Infof(context.TODO(), "watch app.logLevel err:%v", err)
	}

	app, err := initApp(context.Background(), conf)
	if err != nil {
		panic(err)
	}

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newApp(ctx context.Context, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.Context(ctx),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(log.DefaultLogger),
		kratos.Server(
			hs,
		),
	)
}

// initApp init kratos application.
func initApp(ctx context.Context, bc *conf.Bootstrap) (*kratos.App, error) {
	data, err := data.NewData(bc.Data)
	if err != nil {
		return nil, err
	}

	service := service.NewService(bc, data)
	httpServer := server.NewHTTPServer(bc.Server, service)

	app := newApp(ctx, httpServer)
	return app, nil
}

func initConf() (*conf.Bootstrap, config.Config, error) {
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	if err := c.Load(); err != nil {
		return nil, nil, err
	}

	var conf conf.Bootstrap
	if err := c.Scan(&conf); err != nil {
		return nil, nil, err
	}

	//// init Secrect
	//conf.Server.Http.JwtSecrect = jwt.FormatToken(constant.SYSTEM_TYPE_AI, conf.Server.Http.JwtSecrect)
	//err := c.Watch("server.http.jwtSecrect", func(s string, value config.Value) {
	//	val, err := value.String()
	//	if err != nil {
	//		fmt.Printf("watch to load config %s error", "server.http.jwtSecrect")
	//		return
	//	}
	//	conf.Server.Http.JwtSecrect = jwt.FormatToken(constant.SYSTEM_TYPE_AI, val)
	//})
	//if err != nil {
	//	log.Infof(context.TODO(), "watch server.http.jwtSecrect err:%v", err)
	//}

	return &conf, c, nil
}
