package resources

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
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/encoding/protojson"
	"gotest.tools/assert"
)

func newNodeClient() (api.NodeServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		//grpc.WithEndpoint("127.0.0.1:9001"),
		grpc.WithEndpoint("dns:///127.0.0.1:9001"), //负载均衡
		grpc.WithTimeout(60*time.Second),
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
	client := api.NewNodeServiceClient(conn)
	return client, nil
}

func TestListNode(t *testing.T) {
	client, err := newNodeClient()
	if err != nil {
		panic(err)
	}

	list, err := client.ListNode(context.Background(), &empty.Empty{})

	if err != nil {
		panic(err)
	}

	encoder := &protojson.MarshalOptions{
		EmitUnpopulated: true,
	}

	for _, item := range list.Nodes {
		bytes, err := encoder.Marshal(item)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
	}

	assert.Assert(t, err == nil)
}
