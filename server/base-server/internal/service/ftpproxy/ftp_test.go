package ftpproxy

import (
	"context"
	api "server/base-server/api/v1"
	"server/common/errors"
	"testing"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func newClient() (api.FtpProxyServiceClient, error) {
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
	client := api.NewFtpProxyServiceClient(conn)
	return client, nil
}

func TestCreateOrUpdateUser(t *testing.T) {
	client, err := newClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	_, err = client.CreateOrUpdateUser(ctx, &api.CreateOrUpdateUserRequest{
		Username:     "user1",
		Email:        "user1@pcl.ac.cn",
		Password:     "123456",
		HomeS3Bucket: "user1",
	})
	if err != nil {
		panic(err)
	}
}

func TestCreateVirtualFolder(t *testing.T) {
	client, err := newClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	_, err = client.CreateVirtualFolder(ctx, &api.CreateVirtualFolderRequest{
		Name:        "d1",
		Username:    "user1",
		VirtualPath: "/dataset/d1",
		S3Bucket:    "virtual",
		S3Object:    "d1/",
	})
	if err != nil {
		panic(err)
	}
}

func TestDeleteVirtualFolder(t *testing.T) {
	client, err := newClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	_, err = client.DeleteVirtualFolder(ctx, &api.DeleteVirtualFolderRequest{
		Name:     "d1",
		Username: "user2",
	})
	if err != nil {
		panic(err)
	}
}
