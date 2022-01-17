package service

import (
	"context"
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	innerapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"
	"server/common/utils"

	"github.com/jinzhu/copier"
)

type ModelDeployService struct {
	api.UnimplementedModelDeployServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewModelDeployService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.ModelDeployServiceServer {
	return &ModelDeployService{
		conf: conf,
		log:  log.NewHelper("ModelDeployService", logger),
		data: data,
	}
}

// 获取模型服务详情
func (s *ModelDeployService) GetModelDepInfo(ctx context.Context, req *api.DepInfoRequest) (*api.DepInfoReply, error) {

	innerDepInfo, err := s.data.ModelDeployClient.GetModelDepInfo(ctx, &innerapi.DepInfoRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	depInfo := &api.DepInfoReply{}
	err = copier.Copy(depInfo, innerDepInfo)
	if err != nil {
		return nil, err
	}

	return depInfo, nil
}

// 模型服务列表
func (s *ModelDeployService) ListDepModel(ctx context.Context, req *api.DepListRequest) (*api.DepListReply, error) {
	innerReq := &innerapi.DepListRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.ModelDeployClient.ListDepModel(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.DepListReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	if reply.DepInfos == nil {
		reply := &api.DepListReply{
			TotalSize: 0,
			DepInfos:  nil,
		}
		return reply, nil
	}

	if reply.DepInfos == nil {
		reply := &api.DepListReply{
			TotalSize: 0,
			DepInfos:  nil,
		}
		return reply, nil
	} else {
		err = s.assignValue(ctx, reply.DepInfos)
		if err != nil {
			return nil, err
		}
		return reply, nil
	}

	return reply, nil
}

// 模型服务事件列表
func (s *ModelDeployService) ListDepEvent(ctx context.Context, req *api.DepEventListRequest) (*api.DepEventListReply, error) {
	innerReq := &innerapi.DepEventListRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.ModelDeployClient.ListDepEvent(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.DepEventListReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *ModelDeployService) assignValue(ctx context.Context, depInfos []*api.DepInfo) error {
	userIdMap := map[string]interface{}{}
	spaceIdMap := map[string]interface{}{}
	for _, i := range depInfos {
		userIdMap[i.UserId] = true
	}

	users, err := s.data.UserClient.ListUserInCond(ctx, &innerapi.ListUserInCondRequest{Ids: utils.MapKeyToSlice(userIdMap)})
	if err != nil {
		return err
	}
	userMap := map[string]*innerapi.UserItem{}
	for _, i := range users.Users {
		userMap[i.Id] = i
	}

	spaces, err := s.data.WorkspaceClient.ListWorkspaceInCond(ctx, &innerapi.ListWorkspaceInCondRequest{
		Ids: utils.MapKeyToSlice(spaceIdMap),
	})
	if err != nil {
		return err
	}
	spaceMap := map[string]*innerapi.WorkspaceItem{}
	for _, i := range spaces.Workspaces {
		spaceMap[i.Id] = i
	}

	for _, i := range depInfos {

		if v, ok := userMap[i.UserId]; ok {
			i.UserName = v.FullName
		}

		if v, ok := spaceMap[i.WorkspaceId]; ok {
			i.WorkspaceName = v.Name
		}
	}

	return nil
}
