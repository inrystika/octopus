package develop

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/status"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	api "server/base-server/api/v1"
	"server/common/errors"
	"testing"
)

func newDevelopClient() (api.DevelopClient, error) {
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
	client := api.NewDevelopClient(conn)
	return client, nil
}

func TestCreateNotebook(t *testing.T) {
	client, err := newDevelopClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.CreateNotebook(ctx, &api.CreateNotebookRequest{
		UserId:           "ddbe4b31-cc13-416f-aa80-97495abb80c2",
		WorkspaceId:      "default-workspace",
		Name:             "aaaa",
		Desc:             "bbbb",
		ImageId:          "imageId1",
		AlgorithmId:      "algorithmId1",
		AlgorithmVersion: "V1",
		ResourceSpecId:   "resourceSpecId1",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestStopNotebook(t *testing.T) {
	client, err := newDevelopClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.StopNotebook(ctx, &api.StopNotebookRequest{Id: "s45ae04f16f141e0b23e329b7acd1f9b"})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestStartNotebook(t *testing.T) {
	client, err := newDevelopClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.StartNotebook(ctx, &api.StartNotebookRequest{Id: "cf4a154ba68b4e3485ac65c91a024be7"})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestListNotebook(t *testing.T) {
	client, err := newDevelopClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.ListNotebook(ctx, &api.ListNotebookRequest{SearchKey: "aa"})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestGetNotebook(t *testing.T) {
	client, err := newDevelopClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.GetNotebook(ctx, &api.GetNotebookRequest{Id: "1d0b958fd4f047878ecb259b956b2947"})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestDeleteNotebook(t *testing.T) {
	client, err := newDevelopClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.DeleteNotebook(ctx, &api.DeleteNotebookRequest{Id: "1d0b958fd4f047878ecb259b956b2947"})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}
