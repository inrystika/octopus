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
)

type ResourceSpecService struct {
	api.UnimplementedResourceSpecServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewResourceSpecService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.ResourceSpecServiceServer {

	return &ResourceSpecService{
		conf: conf,
		log:  log.NewHelper("ResourceSpecService", logger),
		data: data,
	}
}

func (rsvc *ResourceSpecService) ListResourceSpec(ctx context.Context, req *api.ListResourceSpecRequest) (*api.ResourceSpecList, error) {
	reply, err := rsvc.data.ResourceSpecClient.ListResourceSpec(ctx, &innerapi.ListResourceSpecRequest{
		PageSize:  req.PageSize,
		PageIndex: req.PageIndex,
	})

	if err != nil {
		return nil, err
	}

	replyBytes, err := json.Marshal(reply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorListResourceSpec)
	}

	apiReply := &api.ResourceSpecList{}
	err = json.Unmarshal(replyBytes, apiReply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorListResourceSpec)
	}

	return apiReply, nil
}

func (rsvc *ResourceSpecService) CreateResourceSpec(ctx context.Context, req *api.CreateResourceSpecRequest) (*api.CreateResourceSpecReply, error) {

	reply, err := rsvc.data.ResourceSpecClient.CreateResourceSpec(ctx, &innerapi.CreateResourceSpecRequest{
		Name:             req.Name,
		Price:            req.Price,
		ResourceQuantity: req.ResourceQuantity,
	})

	if err != nil {
		return nil, err
	}

	replyBytes, err := json.Marshal(reply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorCreateResourceSpec)
	}

	apiReply := &api.CreateResourceSpecReply{}
	err = json.Unmarshal(replyBytes, apiReply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorCreateResourceSpec)
	}

	return apiReply, nil
}

func (rsvc *ResourceSpecService) DeleteResourceSpec(ctx context.Context, req *api.DeleteResourceSpecRequest) (*api.DeleteResourceSpecReply, error) {

	reply, err := rsvc.data.ResourceSpecClient.DeleteResourceSpec(ctx, &innerapi.DeleteResourceSpecRequest{
		Id: req.Id,
	})

	if err != nil {
		return nil, err
	}

	replyBytes, err := json.Marshal(reply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorDeleteResourceSpec)
	}

	apiReply := &api.DeleteResourceSpecReply{}
	err = json.Unmarshal(replyBytes, apiReply)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorDeleteResourceSpec)
	}

	return apiReply, nil
}
