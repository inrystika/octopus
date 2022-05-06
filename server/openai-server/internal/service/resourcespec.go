package service

import (
	"context"
	innerapi "server/base-server/api/v1"
	"server/common/constant"
	commctx "server/common/context"
	"server/common/errors"
	"server/common/log"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"

	"github.com/golang/protobuf/ptypes/empty"
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

func (rsvc *ResourceSpecService) ListResourceSpec(ctx context.Context, req *empty.Empty) (*api.ListResourceSpecReply, error) {

	_, workSpaceId, isDefaultSpace := rsvc.getUserIdAndSpaceId(ctx)

	rq := &innerapi.GetResourcePoolReply{}
	if isDefaultSpace {
		_, err := rsvc.data.ResourcePoolClient.GetDefaultResourcePool(ctx, &empty.Empty{})

		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorListResourceSpec)
		}
	} else {
		wsReply, err := rsvc.data.WorkspaceClient.FindWorkspace(ctx, &innerapi.FindWorkspaceRequest{
			Id: workSpaceId,
		})

		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorListResourceSpec)
		}

		resourcePoolId := wsReply.Workspace.ResourcePoolId

		rq, err = rsvc.data.ResourcePoolClient.GetResourcePool(ctx, &innerapi.GetResourcePoolRequest{
			Id: resourcePoolId,
		})

		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorListResourceSpec)
		}
	}

	allResourceSpecList, err := rsvc.data.ResourceSpecClient.ListResourceSpec(ctx, &innerapi.ListResourceSpecRequest{})
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorListResourceSpec)
	}

	allResourceSpecMap := make(map[string]*api.ResourceSpec)
	for _, resourceSpec := range allResourceSpecList.ResourceSpecs {
		allResourceSpecMap[resourceSpec.Id] = &api.ResourceSpec{
			Id:               resourceSpec.Id,
			Name:             resourceSpec.Name,
			Price:            resourceSpec.Price,
			ResourceQuantity: resourceSpec.ResourceQuantity,
		}
	}

	mapResourceSpecIdList := make(map[string]*api.ResourceSpecList)

	for taskType, resourceSpecIdList := range rq.ResourcePool.MapResourceSpecIdList {

		resourceSpecs := make([]*api.ResourceSpec, 0)

		for _, resourceSpecId := range resourceSpecIdList.ResourceSpecIds {
			if resSpec, ok := allResourceSpecMap[resourceSpecId]; ok {
				resourceSpecs = append(resourceSpecs, resSpec)
			}
		}

		mapResourceSpecIdList[taskType] = &api.ResourceSpecList{
			ResourceSpecs: resourceSpecs,
		}
	}

	return &api.ListResourceSpecReply{
		MapResourceSpecIdList: mapResourceSpecIdList,
	}, nil
}

func (rsvc *ResourceSpecService) getUserIdAndSpaceId(ctx context.Context) (string, string, bool) {
	userId, spaceId := commctx.UserIdAndSpaceIdFromContext(ctx)

	return userId, spaceId, spaceId == constant.SYSTEM_WORKSPACE_DEFAULT
}
