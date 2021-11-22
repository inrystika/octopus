package platform

import (
	"context"
	"fmt"
	api "server/base-server/api/v1"
	"server/common/errors"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func newPlatformClient() (api.PlatformServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		//grpc.WithEndpoint("127.0.0.1:9001"),
		grpc.WithEndpoint("dns:///127.0.0.1:9001"), //负载均衡
		grpc.WithMiddleware(
			middleware.Chain(
				status.Client(status.WithHandler(errors.ErrorDecode)),
				recovery.Recovery(),
			),
		),
		grpc.WithTimeout(time.Minute),
	)
	if err != nil {
		return nil, err
	}
	client := api.NewPlatformServiceClient(conn)
	return client, nil
}

func TestPlatformService_ListPlatformConfigKey(t *testing.T) {
	client, err := newPlatformClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.ListPlatformConfigKey(ctx, &api.ListPlatformConfigKeyRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestPlatformService_CreatePlatformStorageConfig(t *testing.T) {
	client, err := newPlatformClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.CreatePlatformStorageConfig(ctx, &api.CreatePlatformStorageConfigRequest{
		PlatformId: "1f4a1345025f408a9ec7b62f0689c4fb",
		Name:       "tidu",
		Type:       "juicefs",
		Options: &api.StorageOptions{Juicefs: &api.StorageOptions_Juicefs{
			Name:    "jfs3",
			MetaUrl: "redis://:abcde@192.168.202.73:32388/1",
		}},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestPlatformService_GetPlatformStorageConfigByName(t *testing.T) {
	client, err := newPlatformClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.GetPlatformStorageConfigByName(ctx, &api.GetPlatformStorageConfigByNameRequest{
		PlatformId: "1f4a1345025f408a9ec7b62f0689c4fb",
		Name:       "tidu",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestPlatformService_UpdatePlatformConfig(t *testing.T) {
	client, err := newPlatformClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.UpdatePlatformConfig(ctx, &api.UpdatePlatformConfigRequest{
		PlatformId: "1f4a1345025f408a9ec7b62f0689c4fb",
		Config:     map[string]string{"k1": "v1", "k2": "v2"},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestPlatformService_GetPlatformConfig(t *testing.T) {
	client, err := newPlatformClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.GetPlatformConfig(ctx, &api.GetPlatformConfigRequest{
		PlatformId: "1f4a1345025f408a9ec7b62f0689c4fb",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}
