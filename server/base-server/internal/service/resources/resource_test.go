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

func newResourceClient() (api.ResourceServiceClient, error) {
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
	client := api.NewResourceServiceClient(conn)
	return client, nil
}

func TestListResource(t *testing.T) {
	client, err := newResourceClient()
	if err != nil {
		panic(err)
	}

	list, err := client.ListResource(context.Background(), &empty.Empty{})

	if err != nil {
		panic(err)
	}

	encoder := &protojson.MarshalOptions{
		EmitUnpopulated: true,
	}

	for _, item := range list.Resources {
		bytes, err := encoder.Marshal(item)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
	}

	assert.Assert(t, err == nil)
}

func TestCreateCustomizedResourceWithInvalidName(t *testing.T) {
	client, err := newResourceClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	_, err = client.CreateCustomizedResource(ctx, &api.CreateCustomizedResourceRequest{
		Name: "nvidia-p012345678901234567890123456789",
	})

	fmt.Println(err)

	assert.Assert(t, err != nil)
}

func TestCreateCustomizedResource(t *testing.T) {
	client, err := newResourceClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.CreateCustomizedResource(ctx, &api.CreateCustomizedResourceRequest{
		Name:         "nvidia-p012345678901234567890123456789",
		BindingNodes: []string{"p100-1"},
		ResourceRef:  "nvidia.com/gpu",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", reply)

	assert.Assert(t, err == nil)
}

func TestCreateCustomizedResourceWithDuplicateName(t *testing.T) {
	client, err := newResourceClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	_, err = client.CreateCustomizedResource(ctx, &api.CreateCustomizedResourceRequest{
		Name:        "nvidia-p100",
		ResourceRef: "nvidia.com/gpu",
	})

	if err != nil {
		fmt.Println(err)
	}

	assert.Assert(t, err != nil)
}

func TestUpdateSystemResource(t *testing.T) {
	client, err := newResourceClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	reply, err := client.UpdateResource(ctx, &api.UpdateResourceRequest{
		Id:           "21de36ef-cc9b-49d7-9ccd-f0252f8badc4",
		Desc:         "change",
		ResourceRef:  "",
		BindingNodes: []string{"p100-1"},
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}

	assert.Assert(t, err == nil)
}

func TestUpdateCustomizedResource(t *testing.T) {
	client, err := newResourceClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	reply, err := client.UpdateResource(ctx, &api.UpdateResourceRequest{
		Id:           "2519a268-51ba-49da-8d85-e16a5eaf3a22",
		Desc:         "change",
		ResourceRef:  "nvidia.com/gpu",
		BindingNodes: []string{"p100-1"},
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}

	assert.Assert(t, err == nil)
}

func TestUpdateCustomizedResourceWithNotExistNode(t *testing.T) {
	client, err := newResourceClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.UpdateResource(ctx, &api.UpdateResourceRequest{
		Id:           "2519a268-51ba-49da-8d85-e16a5eaf3a22",
		Desc:         "change",
		ResourceRef:  "nvidia.com/gpu",
		BindingNodes: []string{"xxx"},
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}

	assert.Assert(t, err != nil)
}

func TestUpdateCustomizedResourceWithExistNode(t *testing.T) {
	client, err := newResourceClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.UpdateResource(ctx, &api.UpdateResourceRequest{
		Id:           "xxx",
		Desc:         "change",
		ResourceRef:  "nvidia.com/gpu",
		BindingNodes: []string{"p100-1"},
	})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
	}

	assert.Assert(t, err != nil)
}

func TestCreateCustomizedResourceWithNotExistNode(t *testing.T) {
	client, err := newResourceClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	_, err = client.CreateCustomizedResource(ctx, &api.CreateCustomizedResourceRequest{
		Name:         "nvidia-n100",
		ResourceRef:  "nvidia.com/gpu",
		BindingNodes: []string{"xxx"},
	})

	if err != nil {
		fmt.Println(err)
	}

	assert.Assert(t, err != nil)
}

func TestDeleteCustomizedResource(t *testing.T) {
	client, err := newResourceClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.DeleteCustomizedResource(ctx, &api.DeleteCustomizedResourceRequest{
		Id: "ccb5b33e-b78f-4424-b7aa-a03419b13485",
	})

	if err == nil {
		fmt.Println(reply)
	} else {
		fmt.Println(err)
	}

	assert.Assert(t, err == nil)
}
