package service

import (
	"context"
	innterapi "server/base-server/api/v1"
	commctx "server/common/context"
	"server/common/errors"
	"server/common/log"
	ss "server/common/session"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"

	"github.com/jinzhu/copier"
)

type ModelService struct {
	api.UnimplementedModelServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewModelService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.ModelServer {
	return &ModelService{
		conf: conf,
		log:  log.NewHelper("ModelService", logger),
		data: data,
	}
}

// 查询预置模型列表
func (s *ModelService) ListPreModel(ctx context.Context, req *api.ListPreModelRequest) (*api.ListPreModelReply, error) {
	reply, err := s.data.ModelClient.ListPreModel(ctx, &innterapi.ListPreModelRequest{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		SearchKey: req.SearchKey,
	})
	if err != nil {
		return nil, err
	}

	models := make([]*api.ModelDetail, 0)
	for _, m := range reply.Models {
		model, err := s.modelTransfer(ctx, m)
		if err != nil {
			return nil, err
		}

		models = append(models, model)
	}

	return &api.ListPreModelReply{
		TotalSize: reply.TotalSize,
		Models:    models,
	}, nil
}

// 查询我的模型列表
func (s *ModelService) ListMyModel(ctx context.Context, req *api.ListMyModelRequest) (*api.ListMyModelReply, error) {
	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.ModelClient.ListMyModel(ctx, &innterapi.ListMyModelRequest{
		SpaceId:   spaceId,
		UserId:    userId,
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		SearchKey: req.SearchKey,
	})
	if err != nil {
		return nil, err
	}

	models := make([]*api.ModelDetail, 0)
	for _, m := range reply.Models {
		model, err := s.modelTransfer(ctx, m)
		if err != nil {
			return nil, err
		}

		models = append(models, model)
	}

	return &api.ListMyModelReply{
		TotalSize: reply.TotalSize,
		Models:    models,
	}, nil
}

// 查询公共模型列表
func (s *ModelService) ListCommModel(ctx context.Context, req *api.ListCommModelRequest) (*api.ListCommModelReply, error) {
	_, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.ModelClient.ListCommModel(ctx, &innterapi.ListCommModelRequest{
		SpaceId:   spaceId,
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		SearchKey: req.SearchKey,
	})
	if err != nil {
		return nil, err
	}

	models := make([]*api.ModelDetail, 0)
	for _, m := range reply.Models {
		model, err := s.modelTransfer(ctx, m)
		if err != nil {
			return nil, err
		}

		models = append(models, model)
	}

	return &api.ListCommModelReply{
		TotalSize: reply.TotalSize,
		Models:    models,
	}, nil
}

// 查询模型版本列表
func (s *ModelService) ListModelVersion(ctx context.Context, req *api.ListModelVersionRequest) (*api.ListModelVersionReply, error) {
	_, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.ModelClient.ListModelVersion(ctx, &innterapi.ListModelVersionRequest{
		ModelId:   req.ModelId,
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	commReply, err := s.data.ModelClient.ListCommModelVersion(ctx, &innterapi.ListCommModelVersionRequest{
		SpaceId: spaceId,
		ModelId: req.ModelId,
	})
	if err != nil {
		return nil, err
	}

	modelVersions := make([]*api.MyVersionDetail, 0)
	for _, mv := range reply.ModelVersions {
		isShared := false
		for _, cmv := range commReply.ModelVersions {
			if mv.Version == cmv.Version {
				isShared = true
				break
			}
		}

		modelVersion, err := s.modelVersionTransfer(mv)
		if err != nil {
			return nil, err
		}

		modelVersions = append(modelVersions, &api.MyVersionDetail{
			IsShared:      isShared,
			VersionDetail: modelVersion,
		})
	}

	return &api.ListModelVersionReply{
		TotalSize:     reply.TotalSize,
		ModelVersions: modelVersions,
	}, nil
}

// 查询公共模型版本列表
func (s *ModelService) ListCommModelVersion(ctx context.Context, req *api.ListCommModelVersionRequest) (*api.ListCommModelVersionReply, error) {
	_, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.ModelClient.ListCommModelVersion(ctx, &innterapi.ListCommModelVersionRequest{
		SpaceId:   spaceId,
		ModelId:   req.ModelId,
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	modelVersions := make([]*api.VersionDetail, 0)
	for _, mv := range reply.ModelVersions {
		modelVersion, err := s.modelVersionTransfer(mv)
		if err != nil {
			return nil, err
		}

		modelVersions = append(modelVersions, modelVersion)
	}

	return &api.ListCommModelVersionReply{
		TotalSize:     reply.TotalSize,
		ModelVersions: modelVersions,
	}, nil
}

// 分享模型版本到公共模型
func (s *ModelService) ShareModelVersion(ctx context.Context, req *api.ShareModelVersionRequest) (*api.ShareModelVersionReply, error) {
	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.ModelClient.ShareModelVersion(ctx, &innterapi.ShareModelVersionRequest{
		SpaceId:          spaceId,
		UserId:           userId,
		ModelId:          req.ModelId,
		Version:          req.Version,
		ShareSpaceIdList: []string{spaceId},
	})
	if err != nil {
		return nil, err
	}

	return &api.ShareModelVersionReply{
		SharedAt: reply.SharedAt,
	}, nil
}

// 取消分享模型版本到公共模型
func (s *ModelService) CloseShareModelVersion(ctx context.Context, req *api.CloseShareModelVersionRequest) (*api.CloseShareModelVersionReply, error) {
	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.ModelClient.CloseShareModelVersion(ctx, &innterapi.CloseShareModelVersionRequest{
		SpaceId:          spaceId,
		UserId:           userId,
		ModelId:          req.ModelId,
		Version:          req.Version,
		ShareSpaceIdList: []string{spaceId},
	})
	if err != nil {
		return nil, err
	}

	return &api.CloseShareModelVersionReply{
		CloseSharedAt: reply.CloseSharedAt,
	}, nil
}

// 删除我的模型版本
func (s *ModelService) DeleteMyModelVersion(ctx context.Context, req *api.DeleteMyModelVersionRequest) (*api.DeleteMyModelVersionReply, error) {
	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.ModelClient.DeleteMyModelVersion(ctx, &innterapi.DeleteMyModelVersionRequest{
		SpaceId: spaceId,
		UserId:  userId,
		ModelId: req.ModelId,
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}

	return &api.DeleteMyModelVersionReply{
		DeletedAt: reply.DeletedAt,
	}, nil
}

// 删除我的模型
func (s *ModelService) DeleteMyModel(ctx context.Context, req *api.DeleteMyModelRequest) (*api.DeleteMyModelReply, error) {
	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.ModelClient.DeleteMyModel(ctx, &innterapi.DeleteMyModelRequest{
		SpaceId: spaceId,
		UserId:  userId,
		ModelId: req.ModelId,
	})
	if err != nil {
		return nil, err
	}

	return &api.DeleteMyModelReply{
		DeletedAt: reply.DeletedAt,
	}, nil
}

// 下载模型
func (s *ModelService) DownloadModelVersion(ctx context.Context, req *api.DownloadModelVersionRequest) (*api.DownloadModelVersionReply, error) {
	isCanView, err := s.viewCtrl(ctx, req.ModelId, req.Version)
	if err != nil {
		return nil, err
	}
	if !isCanView {
		return nil, errors.Errorf(nil, errors.ErrorModelNoPermission)
	}

	reply, err := s.data.ModelClient.DownloadModelVersion(ctx, &innterapi.DownloadModelVersionRequest{
		ModelId: req.ModelId,
		Version: req.Version,
		Domain:  req.Domain,
	})
	if err != nil {
		return nil, err
	}

	return &api.DownloadModelVersionReply{
		DownloadUrl: reply.DownloadUrl,
	}, nil
}

// 预览模型版本
func (s *ModelService) ListModelVersionFile(ctx context.Context, req *api.ListModelVersionFileRequest) (*api.ListModelVersionFileReply, error) {
	isCanView, err := s.viewCtrl(ctx, req.ModelId, req.Version)
	if err != nil {
		return nil, err
	}
	if !isCanView {
		return nil, errors.Errorf(nil, errors.ErrorModelNoPermission)
	}

	reply, err := s.data.ModelClient.ListModelVersionFile(ctx, &innterapi.ListModelVersionFileRequest{
		ModelId: req.ModelId,
		Version: req.Version,
		Prefix:  req.Prefix,
	})
	if err != nil {
		return nil, err
	}

	modelInfoList := make([]*api.ModelInfo, 0)
	for _, m := range reply.ModelInfoList {
		modelInfo := &api.ModelInfo{}

		err = copier.Copy(modelInfo, m)
		if err != nil {
			err = errors.Errorf(err, errors.ErrorStructCopy)
			s.log.Errorw(ctx, err)
			return nil, err
		}

		modelInfoList = append(modelInfoList, modelInfo)
	}

	return &api.ListModelVersionFileReply{
		ModelInfoList: modelInfoList,
	}, nil
}

func (s *ModelService) modelTransfer(ctx context.Context, model *innterapi.ModelDetail) (*api.ModelDetail, error) {
	modelDetail := &api.ModelDetail{}

	err := copier.Copy(modelDetail, model)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorStructCopy)
		s.log.Errorw(ctx, err)
		return nil, err
	}

	if model.UserId != "" {
		userReply, err := s.data.UserClient.FindUser(ctx, &innterapi.FindUserRequest{Id: model.UserId})
		if err != nil || userReply.User == nil {
			modelDetail.UserName = ""
		} else {
			modelDetail.UserName = userReply.User.FullName
		}
	}
	if model.SpaceId != "" {
		spaceReply, err := s.data.WorkspaceClient.FindWorkspace(ctx, &innterapi.FindWorkspaceRequest{Id: model.SpaceId})
		if err != nil || spaceReply.Workspace == nil {
			modelDetail.SpaceName = ""
		} else {
			modelDetail.SpaceName = spaceReply.Workspace.Name
		}
	}
	algorithmReply, err := s.data.AlgorithmClient.QueryAlgorithmVersion(ctx, &innterapi.QueryAlgorithmVersionRequest{
		AlgorithmId: model.AlgorithmId,
		Version:     model.AlgorithmVersion,
	})
	if err != nil || algorithmReply.Algorithm == nil {
		modelDetail.AlgorithmName = ""
	} else {
		modelDetail.AlgorithmName = algorithmReply.Algorithm.AlgorithmName
	}

	return modelDetail, nil
}

func (s *ModelService) modelVersionTransfer(modelVersion *innterapi.VersionDetail) (*api.VersionDetail, error) {
	modelVersionDetail := &api.VersionDetail{}

	err := copier.Copy(modelVersionDetail, modelVersion)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorStructCopy)
		s.log.Errorw(context.Background(), err)
		return nil, err
	}

	return modelVersionDetail, nil
}

func (s *ModelService) getUserIdAndSpaceId(ctx context.Context) (string, string, error) {
	userId := commctx.UserIdFromContext(ctx)
	if userId == "" {
		err := errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
		s.log.Errorw(ctx, err)
		return "", "", err
	}

	session := ss.SessionFromContext(ctx)
	if session == nil {
		err := errors.Errorf(nil, errors.ErrorUserNoAuthSession)
		s.log.Errorw(ctx, err)
		return "", "", err
	}

	return userId, session.GetWorkspace(), nil
}

func (s *ModelService) viewCtrl(ctx context.Context, modelId string, version string) (bool, error) {
	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return false, err
	}

	reply, err := s.data.ModelClient.QueryModelVersion(ctx, &innterapi.QueryModelVersionRequest{
		ModelId: modelId,
		Version: version,
	})
	if err != nil {
		return false, err
	}
	if reply.Model == nil || reply.ModelVersion == nil {
		return false, nil
	}

	// 我的模型可以操作
	if reply.Model.SpaceId == spaceId && reply.Model.UserId == userId {
		return true, nil
	}
	// 预置模型可以操作
	if reply.Model.IsPrefab {
		return true, nil
	}
	// 分享的模型可以操作
	commReply, err := s.data.ModelClient.ListCommModelVersion(ctx, &innterapi.ListCommModelVersionRequest{
		SpaceId: spaceId,
		ModelId: modelId,
	})
	if err != nil {
		return false, err
	}
	for _, mv := range commReply.ModelVersions {
		if mv.Version == version {
			return true, nil
		}
	}

	return false, nil
}
