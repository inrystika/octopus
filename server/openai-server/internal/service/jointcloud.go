package service

import (
	"context"
	innerapi "server/base-server/api/v1"
	commctx "server/common/context"
	"server/common/errors"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"

	"github.com/jinzhu/copier"
)

type JointCloudService struct {
	api.UnimplementedJointCloudServiceServer
	conf *conf.Bootstrap
	data *data.Data
}

func NewJointCloudService(conf *conf.Bootstrap, data *data.Data) api.JointCloudServiceServer {
	s := &JointCloudService{
		conf: conf,
		data: data,
	}

	return s
}

func (s *JointCloudService) ListJointCloudDataset(ctx context.Context, req *api.ListJointCloudDatasetRequest) (*api.ListJointCloudDatasetReply, error) {
	innerReq := &innerapi.ListJointCloudDatasetRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.JointCloudClient.ListJointCloudDataset(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListJointCloudDatasetReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *JointCloudService) ListJointCloudDatasetVersion(ctx context.Context, req *api.ListJointCloudDatasetVersionRequest) (*api.ListJointCloudDatasetVersionReply, error) {
	innerReq := &innerapi.ListJointCloudDatasetVersionRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.JointCloudClient.ListJointCloudDatasetVersion(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListJointCloudDatasetVersionReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *JointCloudService) ListJointCloudFramework(ctx context.Context, req *api.ListJointCloudFrameworkRequest) (*api.ListJointCloudFrameworkReply, error) {
	innerReq := &innerapi.ListJointCloudFrameworkRequest{}
	innerReply, err := s.data.JointCloudClient.ListJointCloudFramework(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListJointCloudFrameworkReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *JointCloudService) ListJointCloudFrameworkVersion(ctx context.Context, req *api.ListJointCloudFrameworkVersionRequest) (*api.ListJointCloudFrameworkVersionReply, error) {
	innerReq := &innerapi.ListJointCloudFrameworkVersionRequest{Key: req.Key}
	innerReply, err := s.data.JointCloudClient.ListJointCloudFrameworkVersion(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListJointCloudFrameworkVersionReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *JointCloudService) ListJointCloudInterpreter(ctx context.Context, req *api.ListJointCloudInterpreterRequest) (*api.ListJointCloudInterpreterReply, error) {
	innerReq := &innerapi.ListJointCloudInterpreterRequest{}
	innerReply, err := s.data.JointCloudClient.ListJointCloudInterpreter(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListJointCloudInterpreterReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
func (s *JointCloudService) ListJointCloudInterpreterVersion(ctx context.Context, req *api.ListJointCloudInterpreterVersionRequest) (*api.ListJointCloudInterpreterVersionReply, error) {
	innerReq := &innerapi.ListJointCloudInterpreterVersionRequest{Key: req.Key}
	innerReply, err := s.data.JointCloudClient.ListJointCloudInterpreterVersion(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListJointCloudInterpreterVersionReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

//创建训练任务
func (s *JointCloudService) JointCloudTrainJob(ctx context.Context, req *api.JointCloudTrainJobRequest) (*api.JointCloudTrainJobReply, error) {
	userId, spaceId := commctx.UserIdAndSpaceIdFromContext(ctx)

	innerReq := &innerapi.JointCloudTrainJobRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.UserId = userId
	innerReq.WorkspaceId = spaceId

	innerReply, err := s.data.JointCloudClient.TrainJob(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.JointCloudTrainJobReply{JobId: innerReply.JobId}, nil
}

func (s *JointCloudService) ListJointCloudJob(ctx context.Context, req *api.ListJointCloudJobRequest) (*api.ListJointCloudJobReply, error) {
	innerReq := &innerapi.ListJointCloudJobRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	userId, spaceId := commctx.UserIdAndSpaceIdFromContext(ctx)

	innerReq.UserId = userId
	innerReq.SpaceId = spaceId

	innerReply, err := s.data.JointCloudClient.ListJointCloudJob(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListJointCloudJobReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// 停止训练任务
func (s *JointCloudService) StopJob(ctx context.Context, req *api.JointCloudStopJobRequest) (*api.JointCloudStopJobReply, error) {
	//trainJob, err := s.data.JointCloudClient.GetTrainJobInfo(ctx, &innerapi.TrainJobInfoRequest{Id: req.Id})
	//if err != nil {
	//	return nil, err
	//}
	//if trainJob.TrainJob.UserId != session.UserId {
	//	return nil, errors.Errorf(nil, errors.ErrorNotAuthorized)
	//}
	innerReq := &innerapi.JointCloudStopJobRequest{
		Id:        req.Id,
		Operation: "user stop job",
	}
	reply, err := s.data.JointCloudClient.StopJob(ctx, innerReq)
	if err != nil {
		return nil, err
	}
	return &api.JointCloudStopJobReply{StoppedAt: reply.StoppedAt}, nil
}
