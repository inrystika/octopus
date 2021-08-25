package service

import (
	"context"
	"encoding/json"
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	innerapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"

	"github.com/golang/protobuf/ptypes/empty"
)

type ResourcePoolService struct {
	api.UnimplementedResourcePoolServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewResourcePoolService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.ResourcePoolServiceServer {

	return &ResourcePoolService{
		conf: conf,
		log:  log.NewHelper("ResourcePoolService", logger),
		data: data,
	}
}

func (rsps *ResourcePoolService) ListResourcePool(ctx context.Context, req *empty.Empty) (*api.ResourcePoolList, error) {
	reply, err := rsps.data.ResourcePoolClient.ListResourcePool(ctx, &empty.Empty{})

	if err != nil {
		return nil, err
	}

	replyBytes, err := json.Marshal(reply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorListResourcePool)
	}

	apiReply := &api.ResourcePoolList{}
	err = json.Unmarshal(replyBytes, apiReply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorListResourcePool)
	}

	return apiReply, nil
}

func (rsps *ResourcePoolService) CreateResourcePool(ctx context.Context, req *api.CreateResourcePoolRequest) (*api.ResourcePoolReply, error) {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorCreateResourcePool)
	}

	innerApiReq := &innerapi.CreateResourcePoolRequest{}
	err = json.Unmarshal(reqBytes, innerApiReq)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorCreateResourcePool)
	}

	reply, err := rsps.data.ResourcePoolClient.CreateResourcePool(ctx, innerApiReq)

	if err != nil {
		return nil, err
	}

	replyBytes, err := json.Marshal(reply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorCreateResourcePool)
	}

	apiReply := &api.ResourcePoolReply{}
	err = json.Unmarshal(replyBytes, apiReply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorCreateResourcePool)
	}

	return apiReply, nil
}

func (rsps *ResourcePoolService) UpdateResourcePool(ctx context.Context, req *api.UpdateResourcePoolRequest) (*api.ResourcePoolReply, error) {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorCreateResourcePool)
	}

	innerApiReq := &innerapi.UpdateResourcePoolRequest{}
	err = json.Unmarshal(reqBytes, innerApiReq)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorCreateResourcePool)
	}

	reply, err := rsps.data.ResourcePoolClient.UpdateResourcePool(ctx, innerApiReq)

	if err != nil {
		return nil, err
	}

	replyBytes, err := json.Marshal(reply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorCreateResourcePool)
	}

	apiReply := &api.ResourcePoolReply{}
	err = json.Unmarshal(replyBytes, apiReply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorCreateResourcePool)
	}

	return apiReply, nil
}

func (rsps *ResourcePoolService) DeleteResourcePool(ctx context.Context, req *api.DeleteResourcePoolRequest) (*api.DeleteResourcePoolReply, error) {

	reply, err := rsps.data.ResourcePoolClient.DeleteResourcePool(ctx, &innerapi.DeleteResourcePoolRequest{Id: req.Id})

	if err != nil {
		return nil, err
	}

	replyBytes, err := json.Marshal(reply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorDeleteResourcePool)
	}

	deleteReply := &api.DeleteResourcePoolReply{}
	err = json.Unmarshal(replyBytes, deleteReply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorDeleteResourcePool)
	}

	return deleteReply, nil
}
