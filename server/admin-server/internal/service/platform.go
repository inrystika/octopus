package service

import (
	"context"
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	innerapi "server/base-server/api/v1"
	"server/common/errors"

	"github.com/jinzhu/copier"
)

type platformService struct {
	api.UnimplementedPlatformServiceServer
	conf *conf.Bootstrap
	data *data.Data
}

func NewPlatformService(conf *conf.Bootstrap, data *data.Data) api.PlatformServiceServer {
	s := &platformService{
		conf: conf,
		data: data,
	}

	return s
}

func (s *platformService) ListPlatform(ctx context.Context, req *api.ListPlatformRequest) (*api.ListPlatformReply, error) {
	innerReq := &innerapi.ListPlatformRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.PlatformClient.ListPlatform(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListPlatformReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *platformService) CreatePlatform(ctx context.Context, req *api.CreatePlatformRequest) (*api.CreatePlatformReply, error) {
	innerReq := &innerapi.CreatePlatformRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, err
	}

	innerReply, err := s.data.PlatformClient.CreatePlatform(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.CreatePlatformReply{Id: innerReply.Id}, nil
}
func (s *platformService) UpdatePlatform(ctx context.Context, req *api.UpdatePlatformRequest) (*api.UpdatePlatformReply, error) {
	innerReq := &innerapi.UpdatePlatformRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, err
	}

	_, err = s.data.PlatformClient.UpdatePlatform(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.UpdatePlatformReply{}, nil
}

func (s *platformService) ListPlatformConfigKey(ctx context.Context, req *api.ListPlatformConfigKeyRequest) (*api.ListPlatformConfigKeyReply, error) {
	innerReq := &innerapi.ListPlatformConfigKeyRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.PlatformClient.ListPlatformConfigKey(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListPlatformConfigKeyReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *platformService) ListPlatformStorageConfig(ctx context.Context, req *api.ListPlatformStorageConfigRequest) (*api.ListPlatformStorageConfigReply, error) {
	innerReq := &innerapi.ListPlatformStorageConfigRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.PlatformClient.ListPlatformStorageConfig(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListPlatformStorageConfigReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
func (s *platformService) CreatePlatformStorageConfig(ctx context.Context, req *api.CreatePlatformStorageConfigRequest) (*api.CreatePlatformStorageConfigReply, error) {
	innerReq := &innerapi.CreatePlatformStorageConfigRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, err
	}

	_, err = s.data.PlatformClient.CreatePlatformStorageConfig(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.CreatePlatformStorageConfigReply{}, nil
}

func (s *platformService) DeletePlatformStorageConfig(ctx context.Context, req *api.DeletePlatformStorageConfigRequest) (*api.DeletePlatformStorageConfigReply, error) {
	_, err := s.data.PlatformClient.DeletePlatformStorageConfig(ctx, &innerapi.DeletePlatformStorageConfigRequest{PlatformId: req.PlatformId, Name: req.Name})
	if err != nil {
		return nil, err
	}

	return &api.DeletePlatformStorageConfigReply{}, nil
}

func (s *platformService) GetPlatformConfig(ctx context.Context, req *api.GetPlatformConfigRequest) (*api.GetPlatformConfigReply, error) {
	reply, err := s.data.PlatformClient.GetPlatformConfig(ctx, &innerapi.GetPlatformConfigRequest{PlatformId: req.PlatformId})
	if err != nil {
		return nil, err
	}
	return &api.GetPlatformConfigReply{Config: reply.Config}, nil
}
func (s *platformService) UpdatePlatformConfig(ctx context.Context, req *api.UpdatePlatformConfigRequest) (*api.UpdatePlatformConfigReply, error) {
	_, err := s.data.PlatformClient.UpdatePlatformConfig(ctx, &innerapi.UpdatePlatformConfigRequest{PlatformId: req.PlatformId, Config: req.Config})
	if err != nil {
		return nil, err
	}
	return &api.UpdatePlatformConfigReply{}, nil
}

// 训练任务列表
func (s *platformService) PlatformTrainJobList(ctx context.Context, req *api.PlatformTrainJobListRequest) (*api.PlatformTrainJobListReply, error) {

	innerReq := &innerapi.PlatformTrainJobListRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.PlatformTrainJobClient.TrainJobList(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.PlatformTrainJobListReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	if reply.TrainJobs == nil {
		reply := &api.PlatformTrainJobListReply{
			TotalSize: 0,
			TrainJobs: nil,
		}
		return reply, nil
	}

	return reply, nil
}
