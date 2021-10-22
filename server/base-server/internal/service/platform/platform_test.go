package platform

import (
	"context"
	"fmt"
	api "server/base-server/api/v1"
	"server/common/errors"
	"testing"

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
		Type:       "juice",
		Options:    nil,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}
