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

	"k8s.io/apimachinery/pkg/api/resource"

	"google.golang.org/protobuf/types/known/emptypb"

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

	nodes, err := rsps.data.NodeClient.ListNode(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorListResourcePool)
	}
	for _, p := range apiReply.ResourcePools {
		capacity := make(map[string]*resource.Quantity)
		allocated := make(map[string]*resource.Quantity)
		for _, nn := range p.BindingNodes {
			for _, n := range nodes.Nodes {
				if nn == n.Name {
					for k, v := range n.Capacity {
						q, err := resource.ParseQuantity(v)
						if err != nil {
							log.Errorf(ctx, "parse %s error", v)
							continue
						}
						_, exist := capacity[k]
						if !exist {
							capacity[k] = &q
						} else {
							capacity[k].Add(q)
						}
					}

					for k, v := range n.Allocated {
						q, err := resource.ParseQuantity(v)
						if err != nil {
							log.Errorf(ctx, "parse %s error", v)
							continue
						}
						_, exist := allocated[k]
						if !exist {
							allocated[k] = &q
						} else {
							allocated[k].Add(q)
						}
					}
				}
			}
		}
		p.ResourceCapacity = make(map[string]string)
		p.ResourceAllocated = make(map[string]string)
		for k, v := range capacity {
			p.ResourceCapacity[k] = v.String()
		}
		for k, v := range allocated {
			p.ResourceAllocated[k] = v.String()
		}
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
