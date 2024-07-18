package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"server/base-server/internal/data"
	"server/base-server/internal/server"
	"server/common/errors"
	"server/common/graceful"
	"server/common/utils"
	"strings"
	"time"

	"google.golang.org/grpc/reflection"

	"server/base-server/internal/conf"
	"server/base-server/internal/service"

	"server/common/log"

	"server/common/third_party/kratos/config"
	"server/common/third_party/kratos/config/file"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"gopkg.in/yaml.v2"

	"context"

	"github.com/go-kratos/kratos/v2/transport/http"
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
	flag.StringVar(&flagconf, "conf", "", "config path, eg: -conf config.yaml")
}

// marshalJson error, when values contains map[interface{}]interface{}. 临时修改代码后放入third_party, 后续升级kratos解决
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
	log.DefaultGormLogger.LogMode(log.ConvertToGorm(l))
	err = c.Watch("app.logLevel", func(k string, v config.Value) {
		ls, _ := v.String()
		l := log.ConvertFromString(ls)
		log.DefaultLogger.ResetLevel(l)
		log.DefaultGormLogger.LogMode(log.ConvertToGorm(l))
	})
	if err != nil {
		log.Infof(context.TODO(), "watch app.logLevel err:%v", err)
	}

	app, close, err := initApp(context.Background(), conf, log.DefaultLogger)
	if err != nil {
		panic(err)
	}
	defer close()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}

	// 协程优雅退出
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		graceful.Shutdown(ctx)
	}()
}

func newApp(ctx context.Context, logger log.Logger, hs *http.Server, gs *grpc.Server) *kratos.App {
	return kratos.New(
		kratos.Context(ctx),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}

// initApp init kratos application.
func initApp(ctx context.Context, bc *conf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	data, close, err := data.NewData(bc, logger)
	if err != nil {
		return nil, nil, err
	}
	service, err := service.NewService(ctx, bc, logger, data)
	if err != nil {
		return nil, nil, err
	}

	grpcServer := server.NewGRPCServer(bc.Server, service)
	reflection.Register(grpcServer.Server)
	httpServer := server.NewHTTPServer(bc.Server, service)
	app := newApp(ctx, logger, httpServer, grpcServer)

	// 服务初始化启动时，重试Minio对象删除任务
	go utils.HandlePanicBG(func(i ...interface{}) {
		initMinioRemovingObjectTask(data)
	})()

	return app, close, nil
}

func initMinioRemovingObjectTask(data *data.Data) error {
	objects, err := data.Redis.SMembersMinioRemovingObject()
	if err != nil {
		return err
	}
	for _, object := range objects {
		bucketName := object[:strings.Index(object, "-")]
		objectName := object[strings.Index(object, "-")+1:]
		go func(objectItem string, bucketNameItem string, objectNameItem string) {
			success, _ := data.Minio.RemoveObject(bucketNameItem, objectNameItem)
			if success {
				data.Redis.SRemMinioRemovingObject(objectItem)
			}
		}(object, bucketName, objectName)
	}
	return nil
}

func initStorageConf(c config.Config) ([]byte, error) {
	m := make(map[string]interface{})
	value := c.Value("module.storage.source")
	err := value.Scan(&m)
	if err != nil {
		return nil, err
	}

	storageConf, err := json.Marshal(m)
	if err != nil {
		return nil, errors.Errorf(nil, errors.ErrorJsonMarshal)
	}

	return storageConf, nil
}

func initStoragesConf(c config.Config) ([]byte, error) {
	var m []map[string]interface{}
	value := c.Value("module.storages")
	err := value.Scan(&m)
	if err != nil {
		return nil, err
	}

	storageConf, err := json.Marshal(m)
	if err != nil {
		return nil, errors.Errorf(nil, errors.ErrorJsonMarshal)
	}

	return storageConf, nil
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
	if Name != "" {
		conf.App.Name = Name
	} else {
		Name = conf.App.Name
	}
	if Version != "" {
		conf.App.Version = Version
	} else {
		Version = conf.App.Version
	}

	// json Marshal []byte
	storageConf, err := initStorageConf(c)
	if err != nil {
		return nil, nil, err
	}
	conf.Storage = storageConf

	storagesBytes, err := initStoragesConf(c)
	if err != nil {
		return nil, nil, err
	}
	conf.Storages = storagesBytes

	return &conf, c, nil
}
