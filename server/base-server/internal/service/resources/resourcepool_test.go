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

func newResourcePoolClient() (api.ResourcePoolServiceClient, error) {
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
	client := api.NewResourcePoolServiceClient(conn)
	return client, nil
}

func TestGetDefaultResourcePool(t *testing.T) {
	client, err := newResourcePoolClient()
	if err != nil {
		panic(err)
	}

	pool, err := client.GetDefaultResourcePool(context.Background(), &empty.Empty{})

	if err != nil {
		panic(err)
	} else {
		encoder := &protojson.MarshalOptions{
			EmitUnpopulated: true,
		}

		bytes, err := encoder.Marshal(pool)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
	}

	assert.Assert(t, err == nil)
}

func TestGetResourcePool(t *testing.T) {
	client, err := newResourcePoolClient()
	if err != nil {
		panic(err)
	}

	pool, err := client.GetResourcePool(context.Background(), &api.GetResourcePoolRequest{Id: "a"})

	if err != nil {
		panic(err)
	} else {
		encoder := &protojson.MarshalOptions{
			EmitUnpopulated: true,
		}

		bytes, err := encoder.Marshal(pool)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
	}

	assert.Assert(t, err == nil)
}

func TestListResourcePool(t *testing.T) {
	client, err := newResourcePoolClient()
	if err != nil {
		panic(err)
	}

	list, err := client.ListResourcePool(context.Background(), &empty.Empty{})

	if err != nil {
		panic(err)
	}

	encoder := &protojson.MarshalOptions{
		EmitUnpopulated: true,
	}

	for _, item := range list.ResourcePools {
		bytes, err := encoder.Marshal(item)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
	}

	assert.Assert(t, err == nil)
}

func TestCreateResourcePool(t *testing.T) {
	client, err := newResourcePoolClient()
	if err != nil {
		panic(err)
	}

	reply, err := client.CreateResourcePool(context.Background(), &api.CreateResourcePoolRequest{
		Name:         "test",
		Desc:         "test",
		BindingNodes: []string{"p100-1"},
		MapResourceSpecIdList: map[string]*api.ResourceSpecIdList{
			"debug": {
				ResourceSpecIds: []string{
					"21814953-d5c1-4562-986e-a4e0c802d26a",
				},
			},
			"train": {
				ResourceSpecIds: []string{
					"21814953-d5c1-4562-986e-a4e0c802d26a",
				},
			},
			"deploy": {
				ResourceSpecIds: []string{
					"21814953-d5c1-4562-986e-a4e0c802d26a",
				},
			},
		},
	})

	if err != nil {
		panic(err)
	}

	encoder := &protojson.MarshalOptions{
		EmitUnpopulated: true,
	}

	bytes, err := encoder.Marshal(reply)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	assert.Assert(t, err == nil)
}

func TestUpdateResourcePoolWithDiffNodes(t *testing.T) {
	client, err := newResourcePoolClient()
	if err != nil {
		panic(err)
	}

	reply, err := client.UpdateResourcePool(context.Background(), &api.UpdateResourcePoolRequest{
		Id:           "test",
		Desc:         "test2",
		BindingNodes: []string{"amax1"},
		MapResourceSpecIdList: map[string]*api.ResourceSpecIdList{
			"debug": {
				ResourceSpecIds: []string{
					"1", "2", "3",
				},
			},
			"train": {
				ResourceSpecIds: []string{
					"1", "2", "3",
				},
			},
			"deploy": {
				ResourceSpecIds: []string{
					"1", "2", "3",
				},
			},
		},
	})

	if err != nil {
		panic(err)
	}

	encoder := &protojson.MarshalOptions{
		EmitUnpopulated: true,
	}

	bytes, err := encoder.Marshal(reply)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	assert.Assert(t, err == nil)
}

func TestDeleteResourcePool(t *testing.T) {
	client, err := newResourcePoolClient()
	if err != nil {
		panic(err)
	}

	reply, err := client.DeleteResourcePool(context.Background(), &api.DeleteResourcePoolRequest{
		Id: "test",
	})

	if err != nil {
		panic(err)
	}

	encoder := &protojson.MarshalOptions{
		EmitUnpopulated: true,
	}

	bytes, err := encoder.Marshal(reply)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

	assert.Assert(t, err == nil)
}
