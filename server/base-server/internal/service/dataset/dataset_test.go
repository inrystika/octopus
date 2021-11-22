package dataset

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

func newDatasetClient() (api.DatasetServiceClient, error) {
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
	client := api.NewDatasetServiceClient(conn)
	return client, nil
}

func TestDatasetService_CreateDateset(t *testing.T) {
	client, err := newDatasetClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.CreateDataset(ctx, &api.CreateDatasetRequest{
		SpaceId:    "spaceid1",
		UserId:     "userid1",
		SourceType: api.DatasetSourceType_DST_USER,
		Name:       "name1",
		TypeId:     "type1",
		Desc:       "desc1",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestDatasetService_FinishUploadDataset(t *testing.T) {
	client, err := newDatasetClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.ConfirmUploadDatasetVersion(ctx, &api.ConfirmUploadDatasetVersionRequest{
		DatasetId: "1bd1e52fca094787a33a28a105175e8d",
		Version:   "V1",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestDatasetService_CreateDatasetVersion(t *testing.T) {
	client, err := newDatasetClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.CreateDatasetVersion(ctx, &api.CreateDatasetVersionRequest{
		DatasetId: "e251e606b2024db190179b1fe264e445",
		Desc:      "desc2",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestDatasetService_ShareDatasetVersion(t *testing.T) {
	client, err := newDatasetClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.ShareDatasetVersion(ctx, &api.ShareDatasetVersionRequest{
		DatasetId:    "05f91483dac749d294633ab5e9f5ab84",
		Version:      "V1",
		ShareSpaceId: "spaceid1",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestDatasetService_CloseShareDatasetVersion(t *testing.T) {
	client, err := newDatasetClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.CloseShareDatasetVersion(ctx, &api.CloseShareDatasetVersionRequest{
		DatasetId:    "e251e606b2024db190179b1fe264e445",
		Version:      "V1",
		ShareSpaceId: "spaceid1",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestDatasetService_DeleteDatasetVersion(t *testing.T) {
	client, err := newDatasetClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.DeleteDatasetVersion(ctx, &api.DeleteDatasetVersionRequest{
		DatasetId: "e251e606b2024db190179b1fe264e445",
		Version:   "V2",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestDatasetService_DeleteDataset(t *testing.T) {
	client, err := newDatasetClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.DeleteDataset(ctx, &api.DeleteDatasetRequest{
		Id: "e251e606b2024db190179b1fe264e445",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestDatasetService_ListDataset(t *testing.T) {
	client, err := newDatasetClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.ListDataset(ctx, &api.ListDatasetRequest{
		PageIndex:    0,
		PageSize:     0,
		SortBy:       "",
		OrderBy:      "",
		CreatedAtGte: 0,
		CreatedAtLt:  0,
		SearchKey:    "",
		UserId:       "userid1",
		SpaceId:      "",
		SourceType:   0,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestDatasetService_GetDatasetUploadUrl(t *testing.T) {
	client, err := newDatasetClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.UploadDatasetVersion(ctx, &api.UploadDatasetVersionRequest{
		DatasetId: "e532cc46c849463abf1daa9c93db3c20",
		Version:   "V1",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestDatasetService_GetDataset(t *testing.T) {
	client, err := newDatasetClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.GetDataset(ctx, &api.GetDatasetRequest{
		Id: "05f91483dac749d294633ab5e9f5ab84",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestDatasetService_GetDatasetVersion(t *testing.T) {
	client, err := newDatasetClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.GetDatasetVersion(ctx, &api.GetDatasetVersionRequest{
		DatasetId: "05f91483dac749d294633ab5e9f5ab84",
		Version:   "V1",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestDatasetService_ListDatasetVersionFile(t *testing.T) {
	client, err := newDatasetClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.ListDatasetVersionFile(ctx, &api.ListDatasetVersionFileRequest{
		DatasetId: "05f91483dac749d294633ab5e9f5ab84",
		Version:   "V1",
		Path:      "/",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}
