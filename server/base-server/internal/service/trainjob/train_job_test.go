package trainjob

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

func newTrainJobClient() (api.TrainJobServiceClient, error) {
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
	client := api.NewTrainJobServiceClient(conn)
	return client, nil
}

func TestCreateTrainJob(t *testing.T) {
	client, err := newTrainJobClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	baseServerConfig := make([]*api.Config, 0)
	parameters := make([]*api.Parameter, 0)
	parameter := &api.Parameter{
	}
	parameters = append(parameters, parameter)
	reqConfig := &api.Config{
		Command:        "python /code/npu-test/npu-test.py",
		Parameters:     parameters,
		ResourceSpecId: "853765f501bb4f69be6fe72d3b36917d",
		TaskNumber: 1,
		MinFailedTaskCount: 1,
		MinSucceededTaskCount: 1,
	}
	baseServerConfig = append(baseServerConfig, reqConfig)

	reply, err := client.TrainJob(ctx, &api.TrainJobRequest{
		UserId:           "cfb1f7a8cb0a4eb6a5a20987765dbf23",
		WorkspaceId:      "default-workspace",
		Name:             "npu-test-1",
		Desc:             "this is a test",
		AlgorithmId:      "2bbe052081074fe5bc52da486559cd3d",
		AlgorithmVersion: "V1",
		ImageId:          "6baf34d8836b425186a56f5179ec088b",
		DataSetId:        "9b402ecb53364336b99f5423fe5efb75",
		DataSetVersion:   "V1",
		IsDistributed:    false,
		ResourcePool:     "common-pool",
		Config:           baseServerConfig,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestTrainJobInfo(t *testing.T) {
	client, err := newTrainJobClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.GetTrainJobInfo(ctx, &api.TrainJobInfoRequest{Id: "j21e6fac76aa41dabd8ffce92aa43768"})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)

}

func TestStopJob(t *testing.T) {

	client, err := newTrainJobClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.StopJob(ctx, &api.StopJobRequest{Id: "6d1c03de-63ea-4997-8546-2961134835ed"})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestDeleteJob(t *testing.T) {

	client, err := newTrainJobClient()
	if err != nil {
		panic(err)
	}

	deleteJobs := make([]string, 0)
	jobId := "4054129b-088f-4490-818e-b8308cca6beb"
	deleteJobs = append(deleteJobs, jobId)

	ctx := context.Background()
	reply, err := client.DeleteJob(ctx, &api.DeleteJobRequest{JobIds: deleteJobs})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestTrainJobList(t *testing.T) {
	client, err := newTrainJobClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.TrainJobList(ctx, &api.TrainJobListRequest{SearchKey: "train-job"})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)

}

func TestTrainJobTemplate(t *testing.T) {

	client, err := newTrainJobClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	baseServerConfig := make([]*api.Config, 0)
	parameters := make([]*api.Parameter, 0)
	parameter := &api.Parameter{
		Key:   "key-1",
		Value: "value-1",
	}
	parameters = append(parameters, parameter)
	reqConfig := &api.Config{
		Command:        "sleep 100",
		Parameters:     parameters,
		ResourceSpecId: "resourceSpecId1",
	}
	baseServerConfig = append(baseServerConfig, reqConfig)

	reply, err := client.CreateJobTemplate(ctx, &api.TrainJobTemplateRequest{
		UserId:           "ddbe4b31-cc13-416f-aa80-97495abb80c2",
		WorkspaceId:      "workspace_id_1",
		Name:             "train-job-test-002",
		Desc:             "this is a test",
		AlgorithmId:      "algorithmId",
		AlgorithmVersion: "algorithmVersion",
		ImageId:          "imageId",
		DataSetId:        "dataSetId",
		DataSetVersion:   "0.0.1",
		IsDistributed:    false,
		Config:           baseServerConfig,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)

}

func TestGetJobTemplate(t *testing.T) {
	client, err := newTrainJobClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.GetJobTemplate(ctx, &api.GetJobTemplateRequest{Id: "fbcacf904e1840f1836cb406c39f0166"})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestUpdateTrainJobTemplate(t *testing.T) {

	client, err := newTrainJobClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	baseServerConfig := make([]*api.Config, 0)
	parameters := make([]*api.Parameter, 0)
	parameter := &api.Parameter{
		Key:   "key-1",
		Value: "value-1",
	}
	parameters = append(parameters, parameter)
	reqConfig := &api.Config{
		Command:        "sleep 100",
		Parameters:     parameters,
		ResourceSpecId: "resourceSpecId1",
	}
	baseServerConfig = append(baseServerConfig, reqConfig)

	reply, err := client.UpdateJobTemplate(ctx, &api.TrainJobTemplateRequest{
		Id:               "fbcacf904e1840f1836cb406c39f0166",
		WorkspaceId:      "workspace_id_1",
		Name:             "temp-job-test-update-1",
		Desc:             "this is a test",
		AlgorithmId:      "algorithmId",
		AlgorithmVersion: "algorithmVersion",
		ImageId:          "imageId",
		DataSetId:        "dataSetId",
		DataSetVersion:   "0.0.2",
		IsDistributed:    false,
		Config:           baseServerConfig,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)

}

func TestDeleteTrainJobTemplate(t *testing.T) {

	client, err := newTrainJobClient()
	if err != nil {
		panic(err)
	}

	templateIds := make([]string, 0)
	templateId := "fbcacf904e1840f1836cb406c39f0166"
	templateIds = append(templateIds, templateId)

	ctx := context.Background()
	reply, err := client.DeleteJobTemplate(ctx, &api.DeleteJobTemplateRequest{
		UserId:      "ddbe4b31-cc13-416f-aa80-97495abb80c2",
		TemplateIds: templateIds})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}

func TestTrainJobTemplateList(t *testing.T) {
	client, err := newTrainJobClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	reply, err := client.ListJobTemplate(ctx, &api.TrainJobTemplateListRequest{SearchKey: "temp"})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}
