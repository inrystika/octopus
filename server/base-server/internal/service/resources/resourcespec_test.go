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
	"google.golang.org/protobuf/encoding/protojson"
	"gotest.tools/assert"
)

func newResourceSpecClient() (api.ResourceSpecServiceClient, error) {
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
	client := api.NewResourceSpecServiceClient(conn)
	return client, nil
}

func TestListResourceSpec(t *testing.T) {
	client, err := newResourceSpecClient()
	if err != nil {
		panic(err)
	}

	list, err := client.ListResourceSpec(context.Background(), &api.ListResourceSpecRequest{})
	if err != nil {
		panic(err)
	}

	encoder := &protojson.MarshalOptions{
		EmitUnpopulated: true,
	}

	for _, item := range list.ResourceSpecs {
		bytes, err := encoder.Marshal(item)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
	}

	assert.Assert(t, err == nil)
}

func TestCreateResourceSpec(t *testing.T) {
	client, err := newResourceSpecClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.CreateResourceSpec(ctx, &api.CreateResourceSpecRequest{
		Name:  "8cpu-2p100-100MB",
		Price: 1,
		ResourceQuantity: map[string]string{
			"cpu":         "8",
			"nvidia-p100": "2",
			"memory":      "1000Ki",
		},
	})

	fmt.Println(reply, err)

	assert.Assert(t, err == nil)
}

//shm > 1/2 memory is a error
func TestCreateResourceSpecWithInvalidShm(t *testing.T) {
	client, err := newResourceSpecClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	//shm > 1/2 memory is a error
	reply, err := client.CreateResourceSpec(ctx, &api.CreateResourceSpecRequest{
		Name:  "8cpu-mem100M-shm60M",
		Price: 1,
		ResourceQuantity: map[string]string{
			"cpu":    "8",
			"memory": "1000Ki",
			"shm":    "600Ki",
		},
	})

	fmt.Println(reply, err)

	assert.Assert(t, err == nil)
}

func TestCreateResourceSpecWithValidShm(t *testing.T) {
	client, err := newResourceSpecClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	//shm > 1/2 memory is a error
	reply, err := client.CreateResourceSpec(ctx, &api.CreateResourceSpecRequest{
		Name:  "8cpu-mem100M-shm50M",
		Price: 1,
		ResourceQuantity: map[string]string{
			"cpu":    "8",
			"memory": "1000Ki",
			"shm":    "500Ki",
		},
	})

	fmt.Println(reply, err)

	assert.Assert(t, err == nil)
}

func TestCreateInValidResourceSpecWithShmWithoutMem(t *testing.T) {
	client, err := newResourceSpecClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	//shm > 1/2 memory is a error
	reply, err := client.CreateResourceSpec(ctx, &api.CreateResourceSpecRequest{
		Name:  "8cpu-shm50K",
		Price: 1,
		ResourceQuantity: map[string]string{
			"cpu": "8",
			"shm": "500Ki",
		},
	})

	fmt.Println(reply, err)

	assert.Assert(t, err == nil)
}

func TestDeleteResourceSpec(t *testing.T) {
	client, err := newResourceSpecClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.DeleteResourceSpec(ctx, &api.DeleteResourceSpecRequest{
		Id: "fef3d6fe-06ab-4e1f-9bbd-9874b70410bb",
	})

	fmt.Println(reply, err)

	assert.Assert(t, err == nil)
}
