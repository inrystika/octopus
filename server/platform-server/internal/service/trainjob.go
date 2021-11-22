package service

import (
	"context"
	innerapi "server/base-server/api/v1"
	commctx "server/common/context"
	"server/common/errors"
	api "server/platform-server/api/v1"
	"server/platform-server/internal/conf"
	"server/platform-server/internal/data"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
)

type TrainJobService struct {
	api.UnimplementedTrainJobServiceServer
	conf *conf.Bootstrap
	data *data.Data
}

func NewTrainJobService(conf *conf.Bootstrap, data *data.Data) api.TrainJobServiceServer {
	return &TrainJobService{
		conf: conf,
		data: data,
	}
}

//创建训练任务
func (s *TrainJobService) TrainJob(ctx context.Context, req *api.TrainJobRequest) (*api.TrainJobReply, error) {
	platformId, err := s.getPlatformId(ctx)
	if err != nil {
		return nil, err
	}
	resourcePool, err := s.getResourcePool(ctx, platformId)
	if err != nil {
		return nil, err
	}
	innerReq := &innerapi.PlatformTrainJobRequest{}
	err = copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.PlatformId = platformId
	innerReq.ResourcePool = resourcePool
	innerReply, err := s.data.PlatformTrainJobClient.TrainJob(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.TrainJobReply{JobId: innerReply.JobId}, nil
}

// 停止训练任务
func (s *TrainJobService) StopJob(ctx context.Context, req *api.StopJobRequest) (*api.StopJobReply, error) {
	platformId, err := s.getPlatformId(ctx)
	if err != nil {
		return nil, err
	}
	trainJob, err := s.data.PlatformTrainJobClient.GetTrainJobInfo(ctx, &innerapi.PlatformTrainJobInfoRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	if trainJob.TrainJob.PlatformId != platformId {
		return nil, errors.Errorf(nil, errors.ErrorNotAuthorized)
	}
	innerReq := &innerapi.PlatformStopJobRequest{
		Id:         req.Id,
		PlatformId: platformId,
		Operation:  "user stop job",
	}
	reply, err := s.data.PlatformTrainJobClient.StopJob(ctx, innerReq)
	if err != nil {
		return nil, err
	}
	return &api.StopJobReply{StoppedAt: reply.StoppedAt}, nil
}

// 获取训练任务详情
func (s *TrainJobService) GetJobInfo(ctx context.Context, req *api.TrainJobInfoRequest) (*api.TrainJobInfoReply, error) {
	platformId, err := s.getPlatformId(ctx)
	if err != nil {
		return nil, err
	}

	innerTrainJobInfo, err := s.data.PlatformTrainJobClient.GetTrainJobInfo(ctx, &innerapi.PlatformTrainJobInfoRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	if innerTrainJobInfo.TrainJob.PlatformId != platformId {
		return nil, errors.Errorf(nil, errors.ErrorNotAuthorized)
	}

	reply := &api.TrainJobInfoReply{}
	err = copier.Copy(reply, innerTrainJobInfo)
	if err != nil {
		return nil, err
	}

	err = s.assignValueToJob(ctx, []*api.TrainJob{reply.TrainJob})
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// 训练任务列表
func (s *TrainJobService) TrainJobList(ctx context.Context, req *api.TrainJobListRequest) (*api.TrainJobListReply, error) {
	platformId, err := s.getPlatformId(ctx)
	if err != nil {
		return nil, err
	}

	innerReq := &innerapi.PlatformTrainJobListRequest{}
	err = copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.PlatformId = platformId

	innerReply, err := s.data.PlatformTrainJobClient.TrainJobList(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.TrainJobListReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	if reply.TrainJobs == nil {
		reply := &api.TrainJobListReply{
			TotalSize: 0,
			TrainJobs: nil,
		}
		return reply, nil
	} else {
		err = s.assignValueToJob(ctx, reply.TrainJobs)
		if err != nil {
			return nil, err
		}
		return reply, nil
	}
}

//获取训练任务统计信息
func (s *TrainJobService) TrainJobStastics(ctx context.Context, req *api.TrainJobStasticsRequest) (*api.TrainJobStasticsReply, error) {
	innerReply, err := s.data.PlatformTrainJobClient.TrainJobStastics(ctx, &innerapi.TrainJobStasticsRequest{
		CreatedAtGte: req.CreatedAtGte,
		CreatedAtLt:  req.CreatedAtLt,
	})
	if err != nil {
		return nil, err
	}

	reply := &api.TrainJobStasticsReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

//获取集群资源信息
func (s *TrainJobService) PlatformResources(ctx context.Context, req *empty.Empty) (*api.PlatformResourcesReply, error) {
	platformId, err := s.getPlatformId(ctx)
	if err != nil {
		return nil, err
	}
	resourcePool, err := s.getResourcePool(ctx, platformId)
	if err != nil {
		return nil, err
	}
	innerReply, err := s.data.PlatformTrainJobClient.PlatformResources(ctx, &innerapi.PlatformResourcesRequest{ResourcePool: resourcePool})
	if err != nil {
		return nil, err
	}
	reply := &api.PlatformResourcesReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *TrainJobService) assignValueToJob(ctx context.Context, trainJobs []*api.TrainJob) error {

	return nil
}

func (s *TrainJobService) getPlatformId(ctx context.Context) (string, error) {
	platformId := commctx.PlatformIdFromContext(ctx)
	if platformId == "" {
		err := errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
		return "", err
	}

	return platformId, nil
}

func (s *TrainJobService) getResourcePool(ctx context.Context, platformId string) (string, error) {
	ids := []string{platformId}
	reply, err := s.data.PlatformClient.BatchGetPlatform(ctx, &innerapi.BatchGetPlatformRequest{Ids: ids})
	if err != nil {
		return "", err
	}
	if len(reply.Platforms) == 0 {
		return "", errors.Errorf(nil, errors.ErrorPlatformBatchGetPlatform)
	}
	resourcePool := reply.Platforms[0].ResourcePool
	return resourcePool, nil
}
