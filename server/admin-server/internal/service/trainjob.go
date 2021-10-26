package service

import (
	"context"
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	innerapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"
	"server/common/utils"

	"github.com/jinzhu/copier"
)

type TrainJobService struct {
	api.UnimplementedTrainJobServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewTrainJobService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.TrainJobServiceServer {
	return &TrainJobService{
		conf: conf,
		log:  log.NewHelper("TrainJobService", logger),
		data: data,
	}
}

func (s *TrainJobService) StopJob(ctx context.Context, req *api.StopJobRequest) (*api.StopJobReply, error) {
	reply, err := s.data.TrainJobClient.StopJob(ctx, &innerapi.StopJobRequest{Id: req.Id, Operation: "admin stop job"})
	if err != nil {
		return nil, err
	}

	return &api.StopJobReply{StoppedAt: reply.StoppedAt}, nil
}

func (s *TrainJobService) GetJobInfo(ctx context.Context, req *api.TrainJobInfoRequest) (*api.TrainJobInfoReply, error) {
	innerTrainJobInfo, err := s.data.TrainJobClient.GetTrainJobInfo(ctx, &innerapi.TrainJobInfoRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	trainJobInfo := &api.TrainJobInfoReply{}
	err = copier.Copy(trainJobInfo, innerTrainJobInfo)
	if err != nil {
		return nil, err
	}

	err = s.assignValue(ctx, []*api.TrainJob{trainJobInfo.TrainJob})
	if err != nil {
		return nil, err
	}

	return trainJobInfo, nil
}

func (s *TrainJobService) TrainJobList(ctx context.Context, req *api.TrainJobListRequest) (*api.TrainJobListReply, error) {
	innerReq := &innerapi.TrainJobListRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.TrainJobClient.TrainJobList(ctx, innerReq)
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
		err = s.assignValue(ctx, reply.TrainJobs)
		if err != nil {
			return nil, err
		}
		return reply, nil
	}
}

func (s *TrainJobService) assignValue(ctx context.Context, trainJobs []*api.TrainJob) error {
	userIdMap := map[string]interface{}{}
	spaceIdMap := map[string]interface{}{}
	algorithmIdMap := map[string]interface{}{}
	imageIdMap := map[string]interface{}{}
	datasetIdMap := map[string]interface{}{}
	for _, i := range trainJobs {
		userIdMap[i.UserId] = true
		algorithmIdMap[i.AlgorithmId] = true
		imageIdMap[i.ImageId] = true
		spaceIdMap[i.WorkspaceId] = true
		datasetIdMap[i.DataSetId] = true
	}

	users, err := s.data.UserClient.ListUserInCond(ctx, &innerapi.ListUserInCondRequest{Ids: utils.MapKeyToSlice(userIdMap)})
	if err != nil {
		return err
	}
	userMap := map[string]*innerapi.UserItem{}
	for _, i := range users.Users {
		userMap[i.Id] = i
	}

	spaces, err := s.data.WorkspaceClient.ListWorkspaceInCond(ctx, &innerapi.ListWorkspaceInCondRequest{
		Ids: utils.MapKeyToSlice(spaceIdMap),
	})
	if err != nil {
		return err
	}
	spaceMap := map[string]*innerapi.WorkspaceItem{}
	for _, i := range spaces.Workspaces {
		spaceMap[i.Id] = i
	}

	algorithms, err := s.data.AlgorithmClient.BatchQueryAlgorithm(ctx, &innerapi.BatchQueryAlgorithmRequest{AlgorithmId: utils.MapKeyToSlice(algorithmIdMap)})
	if err != nil {
		return err
	}
	algorithmMap := map[string]*innerapi.AlgorithmInfo{}
	for _, i := range algorithms.Algorithms {
		algorithmMap[i.AlgorithmId] = i
	}

	images, err := s.data.ImageClient.ListImageInCond(ctx, &innerapi.ListImageInCondRequest{Ids: utils.MapKeyToSlice(imageIdMap)})
	if err != nil {
		return err
	}
	imageMap := map[string]*innerapi.ImageDetail{}
	for _, i := range images.Images {
		imageMap[i.Id] = i
	}

	specs, err := s.data.ResourceSpecClient.ListResourceSpec(ctx, &innerapi.ListResourceSpecRequest{})
	if err != nil {
		return err
	}
	specMap := map[string]*innerapi.ResourceSpec{}
	for _, i := range specs.ResourceSpecs {
		specMap[i.Id] = i
	}

	datasets, err := s.data.DatasetClient.ListDataset(ctx, &innerapi.ListDatasetRequest{
		PageIndex: 1,
		PageSize:  int64(len(datasetIdMap)),
		Ids:       utils.MapKeyToSlice(datasetIdMap),
	})
	if err != nil {
		return err
	}
	datasetMap := map[string]*innerapi.Dataset{}
	for _, i := range datasets.Datasets {
		datasetMap[i.Id] = i
	}

	for _, i := range trainJobs {

		if v, ok := userMap[i.UserId]; ok {
			i.UserName = v.FullName
		}

		if v, ok := spaceMap[i.WorkspaceId]; ok {
			i.WorkspaceName = v.Name
		}

		if v, ok := algorithmMap[i.AlgorithmId]; ok {
			i.AlgorithmName = v.AlgorithmName
		}

		if v, ok := imageMap[i.ImageId]; ok {
			i.ImageName = v.ImageName
		}

		if v, ok := datasetMap[i.DataSetId]; ok {
			i.DataSetName = v.Name
		}
	}

	return nil
}

// 任务事件列表
func (s *TrainJobService) GetJobEventList(ctx context.Context, req *api.JobEventListRequest) (*api.JobEventListReply, error) {
	innerReq := &innerapi.JobEventListRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.TrainJobClient.GetJobEventList(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.JobEventListReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
