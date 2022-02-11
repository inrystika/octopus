package service

import (
	"context"
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	innterapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"
	"server/common/utils/collections/set"

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

// 查询用户模型列表
func (s *ModelService) ListUserModel(ctx context.Context, req *api.ListUserModelRequest) (*api.ListUserModelReply, error) {
	reply, err := s.data.ModelClient.ListAllUserModel(ctx, &innterapi.ListAllUserModelRequest{
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
		SearchKey: req.SearchKey,
	})
	if err != nil {
		return nil, err
	}

	userIds := []string{}
	for _, i := range reply.Models {
		userIds = append(userIds, i.UserId)
	}

	models := make([]*api.ModelDetail, 0)
	for _, m := range reply.Models {
		model, err := s.modelTransfer(ctx, m)
		if err != nil {
			return nil, err
		}

		models = append(models, model)
	}

	if len(userIds) > 0 {
		userIds = set.NewStrings(userIds...).Values()
		userReply, err := s.data.UserClient.ListUserInCond(ctx, &innterapi.ListUserInCondRequest{Ids: userIds})
		if err != nil {
			return nil, err
		}

		emailMap := make(map[string]string)
		for _, u := range userReply.Users {
			emailMap[u.Id] = u.Email
		}

		for _, model := range models {
			model.UserEmail = emailMap[model.UserId]
		}
	}

	return &api.ListUserModelReply{
		TotalSize: reply.TotalSize,
		Models:    models,
	}, nil
}

// 查询模型版本列表
func (s *ModelService) ListModelVersion(ctx context.Context, req *api.ListModelVersionRequest) (*api.ListModelVersionReply, error) {
	reply, err := s.data.ModelClient.ListModelVersion(ctx, &innterapi.ListModelVersionRequest{
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

	return &api.ListModelVersionReply{
		TotalSize:     reply.TotalSize,
		ModelVersions: modelVersions,
	}, nil
}

// 新增预置模型版本
func (s *ModelService) AddPreModelVersion(ctx context.Context, req *api.AddPreModelVersionRequest) (*api.AddPreModelVersionReply, error) {
	reply, err := s.data.ModelClient.AddPreModelVersion(ctx, &innterapi.AddPreModelVersionRequest{
		ModelId:  req.ModelId,
		Descript: req.Descript,
	})
	if err != nil {
		return nil, err
	}

	apiReply := &api.AddPreModelVersionReply{}

	err = copier.Copy(apiReply, reply)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorStructCopy)
		s.log.Errorw(ctx, err)
		return nil, err
	}

	return apiReply, nil
}

// 新增预置模型
func (s *ModelService) AddPreModel(ctx context.Context, req *api.AddPreModelRequest) (*api.AddPreModelReply, error) {
	reply, err := s.data.ModelClient.AddPreModel(ctx, &innterapi.AddPreModelRequest{
		ModelDescript:    req.ModelDescript,
		AlgorithmId:      req.AlgorithmId,
		AlgorithmVersion: req.AlgorithmVersion,
	})
	if err != nil {
		return nil, err
	}

	apiReply := &api.AddPreModelReply{}

	err = copier.Copy(apiReply, reply)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorStructCopy)
		s.log.Errorw(ctx, err)
		return nil, err
	}

	return apiReply, nil
}

// 上传预置模型版本
func (s *ModelService) UploadPreModelVersion(ctx context.Context, req *api.UploadPreModelVersionRequest) (*api.UploadPreModelVersionReply, error) {
	reply, err := s.data.ModelClient.UploadPreModelVersion(ctx, &innterapi.UploadPreModelVersionRequest{
		ModelId:  req.ModelId,
		Version:  req.Version,
		FileName: req.FileName,
		Domain:   req.Domain,
	})
	if err != nil {
		return nil, err
	}

	apiReply := &api.UploadPreModelVersionReply{}

	err = copier.Copy(apiReply, reply)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorStructCopy)
		s.log.Errorw(ctx, err)
		return nil, err
	}

	return apiReply, nil
}

// 上传预置模型版本确认
func (s *ModelService) ConfirmUploadPreModelVersion(ctx context.Context, req *api.ConfirmUploadPreModelVersionRequest) (*api.ConfirmUploadPreModelVersionReply, error) {
	reply, err := s.data.ModelClient.ConfirmUploadPreModelVersion(ctx, &innterapi.ConfirmUploadPreModelVersionRequest{
		ModelId:  req.ModelId,
		Version:  req.Version,
		FileName: req.FileName,
	})
	if err != nil {
		return nil, err
	}

	apiReply := &api.ConfirmUploadPreModelVersionReply{}

	err = copier.Copy(apiReply, reply)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorStructCopy)
		s.log.Errorw(ctx, err)
		return nil, err
	}

	return apiReply, nil
}

// 删除预置模型版本
func (s *ModelService) DeletePreModelVersion(ctx context.Context, req *api.DeletePreModelVersionRequest) (*api.DeletePreModelVersionReply, error) {
	reply, err := s.data.ModelClient.DeletePreModelVersion(ctx, &innterapi.DeletePreModelVersionRequest{
		ModelId: req.ModelId,
		Version: req.Version,
	})
	if err != nil {
		return nil, err
	}

	return &api.DeletePreModelVersionReply{
		DeletedAt: reply.DeletedAt,
	}, nil
}

// 删除预置模型
func (s *ModelService) DeletePreModel(ctx context.Context, req *api.DeletePreModelRequest) (*api.DeletePreModelReply, error) {
	reply, err := s.data.ModelClient.DeletePreModel(ctx, &innterapi.DeletePreModelRequest{
		ModelId: req.ModelId,
	})
	if err != nil {
		return nil, err
	}

	return &api.DeletePreModelReply{
		DeletedAt: reply.DeletedAt,
	}, nil
}

// 下载模型
func (s *ModelService) DownloadModelVersion(ctx context.Context, req *api.DownloadModelVersionRequest) (*api.DownloadModelVersionReply, error) {
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
