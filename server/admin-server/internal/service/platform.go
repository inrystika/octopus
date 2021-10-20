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
