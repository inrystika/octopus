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

type ResourceService struct {
	api.UnimplementedResourceServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewResourceService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.ResourceServiceServer {

	return &ResourceService{
		conf: conf,
		log:  log.NewHelper("ResourceService", logger),
		data: data,
	}
}

func (rsvc *ResourceService) ListResource(ctx context.Context, req *empty.Empty) (*api.ResourceList, error) {
	reply, err := rsvc.data.ResourceClient.ListResource(ctx, &empty.Empty{})

	if err != nil {
		return nil, err
	}

	resourceListBytes, err := json.Marshal(reply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorListResource)
	}

	resourceList := &api.ResourceList{}
	err = json.Unmarshal(resourceListBytes, resourceList)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorListResource)
	}

	return resourceList, nil
}

func (rsvc *ResourceService) CreateCustomizedResource(ctx context.Context, req *api.CreateCustomizedResourceRequest) (*api.CreateCustomizedResourceReply, error) {
	reply, err := rsvc.data.ResourceClient.CreateCustomizedResource(ctx, &innerapi.CreateCustomizedResourceRequest{
		Name:         req.Name,
		Desc:         req.Desc,
		ResourceRef:  req.ResourceRef,
		BindingNodes: req.BindingNodes,
	})

	if err != nil {
		return nil, err
	}

	replyBytes, err := json.Marshal(reply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorCreateResource)
	}

	apiReply := &api.CreateCustomizedResourceReply{}
	err = json.Unmarshal(replyBytes, apiReply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorCreateResource)
	}

	return apiReply, nil
}

func (rsvc *ResourceService) UpdateResource(ctx context.Context, req *api.UpdateResourceRequest) (*api.UpdateResourceReply, error) {
	reply, err := rsvc.data.ResourceClient.UpdateResource(ctx, &innerapi.UpdateResourceRequest{
		Id:           req.Id,
		Desc:         req.Desc,
		ResourceRef:  req.ResourceRef,
		BindingNodes: req.BindingNodes,
	})

	if err != nil {
		return nil, err
	}

	replyBytes, err := json.Marshal(reply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorUpdateResource)
	}

	apiReply := &api.UpdateResourceReply{}
	err = json.Unmarshal(replyBytes, apiReply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorUpdateResource)
	}

	return apiReply, nil
}

func (rsvc *ResourceService) DeleteCustomizedResource(ctx context.Context, req *api.DeleteCustomizedResourceRequest) (*api.DeleteCustomizedResourceReply, error) {
	reply, err := rsvc.data.ResourceClient.DeleteCustomizedResource(ctx, &innerapi.DeleteCustomizedResourceRequest{Id: req.Id})

	if err != nil {
		return nil, err
	}

	replyBytes, err := json.Marshal(reply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorDeleteResource)
	}

	apiReply := &api.DeleteCustomizedResourceReply{}
	err = json.Unmarshal(replyBytes, apiReply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorDeleteResource)
	}

	return apiReply, nil
}
