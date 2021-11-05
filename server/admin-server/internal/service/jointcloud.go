package service

import (
	"context"
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	"server/common/errors"

	innerapi "server/base-server/api/v1"

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

func (s *JointCloudService) ListJointCloudJob(ctx context.Context, req *api.ListJointCloudJobRequest) (*api.ListJointCloudJobReply, error) {
	innerReq := &innerapi.ListJointCloudJobRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

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
