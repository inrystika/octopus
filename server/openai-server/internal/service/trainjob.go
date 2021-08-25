package service

import (
	"context"
	innerapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"
	"server/common/session"
	"server/common/utils"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"

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

//创建训练任务
func (s *TrainJobService) TrainJob(ctx context.Context, req *api.TrainJobRequest) (*api.TrainJobReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.TrainJobRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.UserId = session.UserId
	innerReq.WorkspaceId = session.GetWorkspace()

	innerReply, err := s.data.TrainJobClient.TrainJob(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.TrainJobReply{JobId: innerReply.JobId}, nil
}

// 停止训练任务
func (s *TrainJobService) StopJob(ctx context.Context, req *api.StopJobRequest) (*api.StopJobReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}
	//查询任务是否存在及用户是否一直
	trainJob, err := s.data.TrainJobClient.GetTrainJobInfo(ctx, &innerapi.TrainJobInfoRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	if trainJob.TrainJob.UserId != session.UserId {
		return nil, errors.Errorf(nil, errors.ErrorNotAuthorized)
	}
	innerReq := &innerapi.StopJobRequest{
		Id:        req.Id,
		Operation: "user stop job",
	}
	reply, err := s.data.TrainJobClient.StopJob(ctx, innerReq)
	if err != nil {
		return nil, err
	}
	return &api.StopJobReply{StoppedAt: reply.StoppedAt}, nil
}

//删除训练任务
func (s *TrainJobService) DeleteJob(ctx context.Context, req *api.DeleteJobRequest) (*api.DeleteJobReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	err := s.checkPermission(ctx, req.JobIds, session.UserId)
	if err != nil {
		return nil, err
	}

	innerReq := &innerapi.DeleteJobRequest{UserId: session.UserId, JobIds: req.JobIds}
	reply, err := s.data.TrainJobClient.DeleteJob(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.DeleteJobReply{DeletedAt: reply.DeletedAt}, nil
}

// 获取训练任务详情
func (s *TrainJobService) GetJobInfo(ctx context.Context, req *api.TrainJobInfoRequest) (*api.TrainJobInfoReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerTrainJobInfo, err := s.data.TrainJobClient.GetTrainJobInfo(ctx, &innerapi.TrainJobInfoRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	if innerTrainJobInfo.TrainJob.UserId != session.UserId {
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
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.TrainJobListRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.UserId = session.UserId
	innerReq.WorkspaceId = session.GetWorkspace()

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
		err = s.assignValueToJob(ctx, reply.TrainJobs)
		if err != nil {
			return nil, err
		}
		return reply, nil
	}
}

// 创建训练任务模板
func (s *TrainJobService) CreateJobTemplate(ctx context.Context, req *api.TrainJobTemplateRequest) (*api.TrainJobTemplateReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.TrainJobTemplateRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, err
	}
	innerReq.UserId = session.UserId
	innerReq.WorkspaceId = session.GetWorkspace()

	innerReply, err := s.data.TrainJobClient.CreateJobTemplate(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.TrainJobTemplateReply{
		TemplateId: innerReply.TemplateId,
	}, nil
}

//获取任务模板信息
func (s *TrainJobService) GetJobTemplate(ctx context.Context, req *api.GetJobTemplateRequest) (*api.GetJobTemplateReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReply, err := s.data.TrainJobClient.GetJobTemplate(ctx, &innerapi.GetJobTemplateRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	if innerReply.JobTemplate.UserId != session.UserId {
		return nil, errors.Errorf(nil, errors.ErrorNotAuthorized)
	}

	reply := &api.GetJobTemplateReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	err = s.assignValueToTemplate(ctx, []*api.TrainJobTemplate{reply.JobTemplate})
	if err != nil {
		return nil, err
	}

	return reply, nil
}

//更新任务模板
func (s *TrainJobService) UpdateJobTemplate(ctx context.Context, req *api.TrainJobTemplate) (*api.TrainJobTemplateReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerJobTemplate, err := s.data.TrainJobClient.GetJobTemplate(ctx, &innerapi.GetJobTemplateRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}
	if innerJobTemplate.JobTemplate.UserId != session.UserId {
		return nil, errors.Errorf(nil, errors.ErrorNotAuthorized)
	}

	innerReq := &innerapi.TrainJobTemplateRequest{}
	err = copier.Copy(innerReq, req)
	if err != nil {
		return nil, err
	}
	innerReq.UserId = session.UserId
	innerReq.WorkspaceId = innerJobTemplate.JobTemplate.WorkspaceId

	innerReply, err := s.data.TrainJobClient.UpdateJobTemplate(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.TrainJobTemplateReply{
		TemplateId: innerReply.TemplateId,
	}, nil
}

// 删除任务模板
func (s *TrainJobService) DeleteTemplate(ctx context.Context, req *api.DeleteJobTemplateRequest) (*api.DeleteJobTemplateReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.DeleteJobTemplateRequest{
		UserId:      session.UserId,
		TemplateIds: req.TemplateIds,
	}

	reply, err := s.data.TrainJobClient.DeleteJobTemplate(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.DeleteJobTemplateReply{DeletedAt: reply.DeletedAt}, nil

}

// 任务模板列表
func (s *TrainJobService) TrainJobTemplateList(ctx context.Context, req *api.TrainJobTemplateListRequest) (*api.TrainJobTemplateListReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.TrainJobTemplateListRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.UserId = session.UserId
	innerReq.WorkspaceId = session.GetWorkspace()

	innerReply, err := s.data.TrainJobClient.ListJobTemplate(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.TrainJobTemplateListReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	err = s.assignValueToTemplate(ctx, reply.JobTemplates)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *TrainJobService) assignValueToJob(ctx context.Context, trainJobs []*api.TrainJob) error {
	if len(trainJobs) > 0 {
		algorithmIdMap := map[string]interface{}{}
		imageIdMap := map[string]interface{}{}
		datasetIdMap := map[string]interface{}{}
		for _, i := range trainJobs {
			algorithmIdMap[i.AlgorithmId] = true
			imageIdMap[i.ImageId] = true
			datasetIdMap[i.DataSetId] = true
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
	}

	return nil
}

func (s *TrainJobService) checkPermission(ctx context.Context, jobIds []string, userId string) error {
	for _, jobId := range jobIds {
		reply, err := s.data.TrainJobClient.GetTrainJobInfo(ctx, &innerapi.TrainJobInfoRequest{Id: jobId})
		if err != nil {
			return err
		}

		if reply.TrainJob.UserId != userId {
			return errors.Errorf(nil, errors.ErrorNotAuthorized)
		}
	}
	return nil
}

func (s *TrainJobService) assignValueToTemplate(ctx context.Context, templates []*api.TrainJobTemplate) error {
	if len(templates) > 0 {
		algorithmIdMap := map[string]interface{}{}
		imageIdMap := map[string]interface{}{}
		datasetIdMap := map[string]interface{}{}
		for _, i := range templates {
			algorithmIdMap[i.AlgorithmId] = true
			imageIdMap[i.ImageId] = true
			datasetIdMap[i.DataSetId] = true
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

		for _, i := range templates {
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
	}

	return nil
}
