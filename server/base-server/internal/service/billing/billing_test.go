package billing

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

func newBillingClient() (api.BillingServiceClient, error) {
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
	client := api.NewBillingServiceClient(conn)
	return client, nil
}

func TestBillingService_CreateBillingOwner(t *testing.T) {
	client, err := newBillingClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	owner, err := client.CreateBillingOwner(ctx, &api.CreateBillingOwnerRequest{
		OwnerId:   "fc077cd7a7c7424ea7ec48e178150e85",
		OwnerType: api.BillingOwnerType_BOT_USER,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(owner)
}

func TestBillingService_GetBillingOwner(t *testing.T) {
	client, err := newBillingClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	owner, err := client.GetBillingOwner(ctx, &api.GetBillingOwnerRequest{
		OwnerId:   "fc077cd7a7c7424ea7ec48e178150e85",
		OwnerType: api.BillingOwnerType_BOT_USER,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(owner)
}

func TestBillingService_ListBillingOwner(t *testing.T) {
	client, err := newBillingClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	owner, err := client.ListBillingOwner(ctx, &api.ListBillingOwnerRequest{
		PageIndex: 1,
		PageSize:  10,
		OwnerId:   "fc077cd7a7c7424ea7ec48e178150e85",
		OwnerType: api.BillingOwnerType_BOT_USER,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(owner)
}

func TestBillingService_Pay(t *testing.T) {
	client, err := newBillingClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	owner, err := client.Pay(ctx, &api.PayRequest{
		OwnerId:   "fc077cd7a7c7424ea7ec48e178150e85",
		OwnerType: api.BillingOwnerType_BOT_USER,
		Amount:    3,
		BizType:   api.BillingBizType_BBT_NOTEBOOK,
		BizId:     "bizId2",
		Title:     "title1",
		StartedAt: time.Date(2021, 6, 16, 1, 0, 0, 0, time.UTC).Unix(),
		EndedAt:   time.Date(2021, 6, 16, 3, 0, 0, 0, time.UTC).Unix(),
		Status:    api.BillingPayRecordStatus_BPRS_PAYING,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(owner)
}

func TestHourService_ListBillingPayRecord(t *testing.T) {
	client, err := newBillingClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	owner, err := client.ListBillingPayRecord(ctx, &api.ListBillingPayRecordRequest{
		PageIndex: 1,
		PageSize:  10,
		OwnerId:   "fc077cd7a7c7424ea7ec48e178150e85",
		OwnerType: api.BillingOwnerType_BOT_USER,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(owner)
}

func TestHourService_Recharge(t *testing.T) {
	client, err := newBillingClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	owner, err := client.Recharge(ctx, &api.RechargeRequest{
		OwnerId:   "fc077cd7a7c7424ea7ec48e178150e85",
		OwnerType: api.BillingOwnerType_BOT_USER,
		Amount:    5,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(owner)
}

func TestHourService_ListBillingRechargeRecord(t *testing.T) {
	client, err := newBillingClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	owner, err := client.ListBillingRechargeRecord(ctx, &api.ListBillingRechargeRecordRequest{
		PageIndex: 1,
		PageSize:  10,
		OwnerId:   "fc077cd7a7c7424ea7ec48e178150e85",
		OwnerType: api.BillingOwnerType_BOT_USER,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(owner)
}
