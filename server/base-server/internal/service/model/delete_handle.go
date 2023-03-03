package model

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"
	"server/common/utils"
	"time"

	"server/common/log"
)

type ModelDeleteHandle interface {
	// 删除我的模型版本
	DeleteMyModelVersionHandle(ctx context.Context, req *api.DeleteMyModelVersionRequest) (*api.DeleteMyModelVersionReply, error)
	// 删除我的模型
	DeleteMyModelHandle(ctx context.Context, req *api.DeleteMyModelRequest) (*api.DeleteMyModelReply, error)
	// 删除预置模型版本
	DeletePreModelVersionHandle(ctx context.Context, req *api.DeletePreModelVersionRequest) (*api.DeletePreModelVersionReply, error)
	// 删除预置模型
	DeletePreModelHandle(ctx context.Context, req *api.DeletePreModelRequest) (*api.DeletePreModelReply, error)
}

type modelDeleteHandle struct {
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewModelDeleteHandle(conf *conf.Bootstrap, logger log.Logger, data *data.Data) ModelDeleteHandle {
	return &modelDeleteHandle{
		conf: conf,
		log:  log.NewHelper("ModelDeleteHandle", logger),
		data: data,
	}
}

// 删除我的模型版本
func (h *modelDeleteHandle) DeleteMyModelVersionHandle(ctx context.Context, req *api.DeleteMyModelVersionRequest) (*api.DeleteMyModelVersionReply, error) {
	modelDao := h.data.ModelDao
	spaceId := req.SpaceId
	userId := req.UserId
	modelId := req.ModelId
	version := req.Version

	// 取消分享
	modelShareHandle := &modelShareHandle{conf: h.conf, log: h.log, data: h.data}
	_, err := modelShareHandle.AllCloseShareModelVersionHandle(ctx, &api.AllCloseShareModelVersionRequest{
		SpaceId: spaceId,
		UserId:  userId,
		ModelId: modelId,
		Version: version,
	})
	if err != nil {
		return nil, err
	}

	// 删除版本信息
	err = modelDao.DeleteModelVersion(ctx, modelId, version)
	if err != nil {
		return nil, err
	}

	maxVersion, err := h.findModelVersionMaxId(ctx, modelId)
	if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
		return nil, err
	}
	if err != nil && errors.IsError(errors.ErrorDBFindEmpty, err) {
		// 最后一个版本，那就都删了
		err := modelDao.DeleteModel(ctx, modelId)
		if err != nil {
			return nil, err
		}
	} else {
		// 更新最新版本号
		modelInt, err := modelDao.GetModel(ctx, modelId)
		if err != nil {
			return nil, err
		}

		modelInt.LatestVersion = maxVersion
		err = modelDao.UpdateModel(ctx, modelInt)
		if err != nil {
			return nil, err
		}
	}
	// 删除模型版本Minio存储
	go utils.HandlePanic(ctx, func(i ...interface{}) {
		bucketName := common.GetMinioBucket()
		objectName := common.GetMinioModelObject(spaceId, userId, modelId, version)
		h.data.Redis.SAddMinioRemovingObject(ctx, bucketName+"-"+objectName)
		defer h.data.Redis.SRemMinioRemovingObject(ctx, bucketName+"-"+objectName)
		h.data.Minio.RemoveObject(bucketName, objectName)
	})()
	return &api.DeleteMyModelVersionReply{
		DeletedAt: time.Now().Unix(),
	}, nil
}

// 删除我的模型
func (h *modelDeleteHandle) DeleteMyModelHandle(ctx context.Context, req *api.DeleteMyModelRequest) (*api.DeleteMyModelReply, error) {
	modelDao := h.data.ModelDao
	spaceId := req.SpaceId
	userId := req.UserId
	modelId := req.ModelId

	// 取消分享
	modelShareHandle := &modelShareHandle{conf: h.conf, log: h.log, data: h.data}
	_, err := modelShareHandle.AllCloseShareModelHandle(ctx, &api.AllCloseShareModelRequest{
		SpaceId: spaceId,
		UserId:  userId,
		ModelId: modelId,
	})
	if err != nil {
		return nil, err
	}

	// 删除模型信息
	err = modelDao.DeleteModel(ctx, modelId)
	if err != nil {
		return nil, err
	}

	// 删除模型版本Minio存储
	go utils.HandlePanic(ctx, func(i ...interface{}) {
		bucketName := common.GetMinioBucket()
		objectName := common.GetMinioModelPathObject(spaceId, userId, modelId)
		h.data.Redis.SAddMinioRemovingObject(ctx, bucketName+"-"+objectName)
		defer h.data.Redis.SRemMinioRemovingObject(ctx, bucketName+"-"+objectName)
		h.data.Minio.RemoveObject(bucketName, objectName)
	})()
	return &api.DeleteMyModelReply{
		DeletedAt: time.Now().Unix(),
	}, nil
}

// 删除预置模型版本
func (h *modelDeleteHandle) DeletePreModelVersionHandle(ctx context.Context, req *api.DeletePreModelVersionRequest) (*api.DeletePreModelVersionReply, error) {
	modelDao := h.data.ModelDao
	modelId := req.ModelId
	version := req.Version

	preModel, err := modelDao.GetModel(ctx, req.ModelId)
	if err != nil {
		return nil, err
	}
	if !preModel.IsPrefab {
		err := errors.Errorf(nil, errors.ErrorModelNoPermission)
		return nil, err
	}

	// 删除版本信息
	err = modelDao.DeleteModelVersion(ctx, modelId, version)
	if err != nil {
		return nil, err
	}

	maxVersion, err := h.findModelVersionMaxId(ctx, modelId)
	if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
		return nil, err
	}
	if err != nil && errors.IsError(errors.ErrorDBFindEmpty, err) {
		// 最后一个版本，那就都删了
		err := modelDao.DeleteModel(ctx, modelId)
		if err != nil {
			return nil, err
		}
	} else {
		// 查模型信息
		modelInt, err := modelDao.GetModel(ctx, modelId)
		if err != nil {
			return nil, err
		}

		modelInt.LatestVersion = maxVersion
		err = modelDao.UpdateModel(ctx, modelInt)
		if err != nil {
			return nil, err
		}
	}

	return &api.DeletePreModelVersionReply{
		DeletedAt: time.Now().Unix(),
	}, nil
}

// 删除预置模型
func (h *modelDeleteHandle) DeletePreModelHandle(ctx context.Context, req *api.DeletePreModelRequest) (*api.DeletePreModelReply, error) {
	modelDao := h.data.ModelDao
	modelId := req.ModelId

	preModel, err := modelDao.GetModel(ctx, req.ModelId)
	if err != nil {
		return nil, err
	}
	if !preModel.IsPrefab {
		err := errors.Errorf(nil, errors.ErrorModelNoPermission)
		return nil, err
	}

	// 删除模型信息
	err = modelDao.DeleteModel(ctx, modelId)
	if err != nil {
		return nil, err
	}

	return &api.DeletePreModelReply{
		DeletedAt: time.Now().Unix(),
	}, nil
}

func (h *modelDeleteHandle) findModelVersionMaxId(ctx context.Context, modelId string) (string, error) {
	modelDao := h.data.ModelDao

	_, modelVersionList, err := modelDao.ListModelVersion(ctx, &model.ModelVersionList{
		ModelId:      modelId,
		VersionOrder: true,
		VersionSort:  model.DESC,
		PageIndex:    1,
		PageSize:     1,
	})
	if err != nil {
		return "", err
	}
	if len(modelVersionList) == 0 {
		err = errors.Errorf(nil, errors.ErrorDBFindEmpty)
		return "", err
	}

	return modelVersionList[0].Version, nil
}
