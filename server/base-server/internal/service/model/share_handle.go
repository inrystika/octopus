package model

import (
	"context"
	"fmt"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"
	"server/common/utils"
	"time"

	"server/common/log"

	"github.com/jinzhu/copier"
)

type ModelShareHandle interface {
	// 分享模型版本到公共模型
	ShareModelVersionHandle(ctx context.Context, req *api.ShareModelVersionRequest) (*api.ShareModelVersionReply, error)
	// 取消分享模型版本到公共模型
	CloseShareModelVersionHandle(ctx context.Context, req *api.CloseShareModelVersionRequest) (*api.CloseShareModelVersionReply, error)
	// 取消分享模型到公共模型
	CloseShareModelHandle(ctx context.Context, req *api.CloseShareModelRequest) (*api.CloseShareModelReply, error)
	// 取消分享模型版本到所有公共模型
	AllCloseShareModelVersionHandle(ctx context.Context, req *api.AllCloseShareModelVersionRequest) (*api.AllCloseShareModelVersionReply, error)
	// 取消分享模型到所有公共模型
	AllCloseShareModelHandle(ctx context.Context, req *api.AllCloseShareModelRequest) (*api.AllCloseShareModelReply, error)

	// 下载模型版本
	DownloadModelVersionHandle(ctx context.Context, req *api.DownloadModelVersionRequest) (*api.DownloadModelVersionReply, error)
	// 预览模型版本
	ListModelVersionFileHandle(ctx context.Context, req *api.ListModelVersionFileRequest) (*api.ListModelVersionFileReply, error)
}

type modelShareHandle struct {
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewModelShareHandle(conf *conf.Bootstrap, logger log.Logger, data *data.Data) ModelShareHandle {
	return &modelShareHandle{
		conf: conf,
		log:  log.NewHelper("ModelShareHandle", logger),
		data: data,
	}
}

// 分享模型版本到公共模型
func (h *modelShareHandle) ShareModelVersionHandle(ctx context.Context, req *api.ShareModelVersionRequest) (*api.ShareModelVersionReply, error) {
	modelDao := h.data.ModelDao
	modelId := req.ModelId
	version := req.Version

	myModel, err := modelDao.GetModel(ctx, modelId)
	if err != nil {
		return nil, err
	}
	if myModel.SpaceId != req.SpaceId || myModel.UserId != req.UserId || myModel.IsPrefab {
		err := errors.Errorf(nil, errors.ErrorModelNoPermission)
		return nil, err
	}

	_, err = modelDao.GetModelVersion(ctx, modelId, version)
	if err != nil {
		return nil, err
	}

	for _, shareSpaceId := range req.ShareSpaceIdList {
		// 查可见模型信息
		modelAccess := &model.ModelAccess{}
		_, modelAccesses, err := modelDao.ListModelAccess(ctx, &model.ModelAccessList{
			SpaceIds: []string{shareSpaceId},
			ModelIds: []string{modelId},
		})
		if err != nil {
			return nil, err
		}
		if len(modelAccesses) == 0 {
			// 插入可见模型信息
			modelAccess, err = modelDao.CreateModelAccess(ctx, &model.ModelAccess{
				Id:      utils.GetUUIDWithoutSeparator(),
				ModelId: modelId,
				SpaceId: shareSpaceId,
			})
			if err != nil {
				return nil, err
			}
		} else {
			modelAccess = modelAccesses[0]
		}

		modelAccessId := modelAccess.Id

		// 查看可见模型版本信息
		_, err = modelDao.GetModelVersionAccess(ctx, modelAccessId, version)
		if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
			return nil, err
		}
		if err == nil {
			continue
		}

		// 插入可见模型版本信息
		_, err = modelDao.CreateModelVersionAccess(ctx, &model.ModelVersionAccess{
			Id:            utils.GetUUIDWithoutSeparator(),
			ModelAccessId: modelAccessId,
			ModelVersion:  version,
			ModelId:       modelId,
		})
		if err != nil {
			return nil, err
		}

		// 更新可见模型信息
		if modelAccess.LatestModelVersion < version {
			modelAccess.LatestModelVersion = version
			err = modelDao.UpdateModelAccess(ctx, modelAccess)
			if err != nil {
				return nil, err
			}
		}
	}

	return &api.ShareModelVersionReply{
		SharedAt: time.Now().Unix(),
	}, nil
}

// 取消分享模型版本到公共模型
func (h *modelShareHandle) CloseShareModelVersionHandle(ctx context.Context, req *api.CloseShareModelVersionRequest) (*api.CloseShareModelVersionReply, error) {
	modelDao := h.data.ModelDao
	modelId := req.ModelId
	version := req.Version

	myModel, err := modelDao.GetModel(ctx, modelId)
	if err != nil {
		return nil, err
	}
	if myModel.SpaceId != req.SpaceId || myModel.UserId != req.UserId || myModel.IsPrefab {
		err := errors.Errorf(nil, errors.ErrorModelNoPermission)
		return nil, err
	}

	_, err = modelDao.GetModelVersion(ctx, modelId, version)
	if err != nil {
		return nil, err
	}

	for _, shareSpaceId := range req.ShareSpaceIdList {
		// 查可见模型信息
		modelAccess := &model.ModelAccess{}
		_, modelAccesses, err := modelDao.ListModelAccess(ctx, &model.ModelAccessList{
			SpaceIds: []string{shareSpaceId},
			ModelIds: []string{modelId},
		})
		if err != nil {
			return nil, err
		}
		if len(modelAccesses) == 0 {
			continue
		}

		modelAccess = modelAccesses[0]
		modelAccessId := modelAccess.Id

		// 删除可见模型版本信息
		err = modelDao.DeleteModelVersionAccess(ctx, modelAccessId, version)
		if err != nil {
			return nil, err
		}

		maxModelVersion, err := h.findModelVersionAccessMaxId(ctx, modelAccessId)
		if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
			return nil, err
		}
		if err != nil && errors.IsError(errors.ErrorDBFindEmpty, err) {
			// 最后一个版本，那就都删了
			err := modelDao.DeleteModelAccess(ctx, modelAccessId)
			if err != nil {
				return nil, err
			}
		} else {
			modelAccess.LatestModelVersion = maxModelVersion
			err = modelDao.UpdateModelAccess(ctx, modelAccess)
			if err != nil {
				return nil, err
			}
		}
	}

	return &api.CloseShareModelVersionReply{
		CloseSharedAt: time.Now().Unix(),
	}, nil
}

// 取消分享模型到公共模型
func (h *modelShareHandle) CloseShareModelHandle(ctx context.Context, req *api.CloseShareModelRequest) (*api.CloseShareModelReply, error) {
	modelDao := h.data.ModelDao
	modelId := req.ModelId

	myModel, err := modelDao.GetModel(ctx, modelId)
	if err != nil {
		return nil, err
	}
	if myModel.SpaceId != req.SpaceId || myModel.UserId != req.UserId || myModel.IsPrefab {
		err := errors.Errorf(nil, errors.ErrorModelNoPermission)
		return nil, err
	}

	for _, shareSpaceId := range req.ShareSpaceIdList {
		// 查可见模型信息
		modelAccess := &model.ModelAccess{}
		_, modelAccesses, err := modelDao.ListModelAccess(ctx, &model.ModelAccessList{
			SpaceIds: []string{shareSpaceId},
			ModelIds: []string{modelId},
		})
		if err != nil {
			return nil, err
		}
		if len(modelAccesses) == 0 {
			continue
		}

		modelAccess = modelAccesses[0]
		modelAccessId := modelAccess.Id

		err = modelDao.DeleteModelAccess(ctx, modelAccessId)
		if err != nil {
			return nil, err
		}
	}

	return &api.CloseShareModelReply{
		CloseSharedAt: time.Now().Unix(),
	}, nil
}

// 取消分享模型版本到所有公共模型
func (h *modelShareHandle) AllCloseShareModelVersionHandle(ctx context.Context, req *api.AllCloseShareModelVersionRequest) (*api.AllCloseShareModelVersionReply, error) {
	modelDao := h.data.ModelDao
	spaceId := req.SpaceId
	userId := req.UserId
	modelId := req.ModelId
	version := req.Version

	myModel, err := modelDao.GetModel(ctx, modelId)
	if err != nil {
		return nil, err
	}
	if myModel.SpaceId != req.SpaceId || myModel.UserId != req.UserId || myModel.IsPrefab {
		err := errors.Errorf(nil, errors.ErrorModelNoPermission)
		return nil, err
	}

	_, err = modelDao.GetModelVersion(ctx, modelId, version)
	if err != nil {
		return nil, err
	}

	// 查该版本有多少spaceId可见
	_, modelVersionAccessList, err := modelDao.ListModelVersionAccess(ctx, &model.ModelVersionAccessList{
		ModelId:       modelId,
		ModelVersions: []string{version},
	})
	if err != nil {
		return nil, err
	}
	if len(modelVersionAccessList) == 0 {
		// 没有分享的，直接返回
		return &api.AllCloseShareModelVersionReply{
			CloseSharedAt: time.Now().Unix(),
		}, nil
	}

	modelAccessIdList := make([]string, 0)
	for _, m := range modelVersionAccessList {
		modelAccessIdList = append(modelAccessIdList, m.ModelAccessId)
	}
	_, modelAccessList, err := modelDao.ListModelAccess(ctx, &model.ModelAccessList{
		Ids: modelAccessIdList,
	})
	if err != nil {
		return nil, err
	}
	if len(modelAccessList) == 0 {
		// 没有分享的，直接返回
		return &api.AllCloseShareModelVersionReply{
			CloseSharedAt: time.Now().Unix(),
		}, nil
	}

	shareSpaceIdList := make([]string, 0)
	for _, sp := range modelAccessList {
		shareSpaceIdList = append(shareSpaceIdList, sp.SpaceId)
	}
	_, err = h.CloseShareModelVersionHandle(ctx, &api.CloseShareModelVersionRequest{
		SpaceId:          spaceId,
		UserId:           userId,
		ModelId:          modelId,
		Version:          version,
		ShareSpaceIdList: shareSpaceIdList,
	})
	if err != nil {
		return nil, err
	}

	return &api.AllCloseShareModelVersionReply{
		CloseSharedAt: time.Now().Unix(),
	}, nil
}

// 取消分享模型到所有公共模型
func (h *modelShareHandle) AllCloseShareModelHandle(ctx context.Context, req *api.AllCloseShareModelRequest) (*api.AllCloseShareModelReply, error) {
	modelDao := h.data.ModelDao
	spaceId := req.SpaceId
	userId := req.UserId
	modelId := req.ModelId

	myModel, err := modelDao.GetModel(ctx, modelId)
	if err != nil {
		return nil, err
	}
	if myModel.SpaceId != req.SpaceId || myModel.UserId != req.UserId || myModel.IsPrefab {
		err := errors.Errorf(nil, errors.ErrorModelNoPermission)
		return nil, err
	}

	// 查该模型有多少spaceId可见
	_, accessModels, err := modelDao.ListModelAccess(ctx, &model.ModelAccessList{
		ModelIds: []string{modelId},
	})
	if err != nil {
		return nil, err
	}
	if len(accessModels) == 0 {
		// 没有分享的，直接返回
		return &api.AllCloseShareModelReply{
			CloseSharedAt: time.Now().Unix(),
		}, nil
	}

	shareSpaceIdList := make([]string, 0)
	for _, m := range accessModels {
		shareSpaceIdList = append(shareSpaceIdList, m.SpaceId)
	}

	_, err = h.CloseShareModelHandle(ctx, &api.CloseShareModelRequest{
		SpaceId:          spaceId,
		UserId:           userId,
		ModelId:          modelId,
		ShareSpaceIdList: shareSpaceIdList,
	})
	if err != nil {
		return nil, err
	}

	return &api.AllCloseShareModelReply{
		CloseSharedAt: time.Now().Unix(),
	}, nil
}

// 下载模型版本
func (h *modelShareHandle) DownloadModelVersionHandle(ctx context.Context, req *api.DownloadModelVersionRequest) (*api.DownloadModelVersionReply, error) {
	modelDao := h.data.ModelDao
	modelId := req.ModelId
	version := req.Version

	modelInt, err := modelDao.GetModel(ctx, modelId)
	if err != nil {
		return nil, err
	}

	modelVersion, err := modelDao.GetModelVersion(ctx, modelId, version)
	if err != nil {
		return nil, err
	}
	// 看下文件状态是否正常
	if modelVersion.FileStatus != model.FILESTATUS_FINISH {
		return nil, errors.Errorf(nil, errors.ErrorModelVersionFileNotFound)
	}

	fromBucketName := ""
	fromObjectName := ""
	toBucektName := ""
	toObjectName := ""
	if modelInt.IsPrefab {
		fromBucketName = common.GetMinioBucket()
		fromObjectName = common.GetMinioPreModelObject(modelVersion.ModelId, modelVersion.Version)
	} else {
		fromBucketName = common.GetMinioBucket()
		fromObjectName = common.GetMinioModelObject(modelInt.SpaceId, modelInt.UserId, modelVersion.ModelId, modelVersion.Version)
	}

	toBucektName = common.GetMinioBucket()
	toObjectName = common.GetMinioDownloadModelObject(modelVersion.ModelId, modelVersion.Version, fmt.Sprintf("%s-%s.zip", modelInt.ModelName, modelVersion.Version))

	fromPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, fromBucketName, fromObjectName)
	toPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, toBucektName, toObjectName)
	// 看下有没有缓存
	isExist, err := h.data.Minio.ObjectExist(toBucektName, toObjectName)
	if err != nil {
		return nil, err
	}
	if isExist {
		url, err := h.data.Minio.PresignedDownloadObject(toBucektName, toObjectName, req.Domain)
		if err != nil {
			return nil, err
		}

		return &api.DownloadModelVersionReply{
			DownloadUrl: url.String(),
		}, nil
	}

	// 这里只是做兜底，在模型上传时已经拷贝了一份，正常不会走到这里
	// 压缩拷贝
	h.log.Infof(ctx, "begin to Zip, from: %s, to:%s", fromPath, toPath)
	startT := time.Now()
	err = utils.Zip(fromPath, toPath)
	if err != nil {
		return nil, err
	}
	h.log.Infof(ctx, "Zip success, from: %s, to:%s, cost time: %d", fromPath, toPath, time.Since(startT))

	url, err := h.data.Minio.PresignedDownloadObject(toBucektName, toObjectName, req.Domain)
	if err != nil {
		return nil, err
	}

	return &api.DownloadModelVersionReply{
		DownloadUrl: url.String(),
	}, nil
}

// 预览模型版本
func (h *modelShareHandle) ListModelVersionFileHandle(ctx context.Context, req *api.ListModelVersionFileRequest) (*api.ListModelVersionFileReply, error) {
	modelDao := h.data.ModelDao
	modelId := req.ModelId
	version := req.Version
	prefix := req.Prefix

	modelInt, err := modelDao.GetModel(ctx, modelId)
	if err != nil {
		return nil, err
	}

	modelVersion, err := modelDao.GetModelVersion(ctx, modelId, version)
	if err != nil {
		return nil, err
	}
	// 看下文件状态是否正常
	if modelVersion.FileStatus != model.FILESTATUS_FINISH {
		return nil, errors.Errorf(nil, errors.ErrorModelVersionFileNotFound)
	}

	bucketName := ""
	objectPrefix := ""
	if modelInt.IsPrefab {
		bucketName = common.GetMinioBucket()
		objectPrefix = fmt.Sprintf("%s/%s", common.GetMinioPreModelObject(modelVersion.ModelId, modelVersion.Version), prefix)
	} else {
		bucketName = common.GetMinioBucket()
		objectPrefix = fmt.Sprintf("%s/%s", common.GetMinioModelObject(modelInt.SpaceId, modelInt.UserId, modelVersion.ModelId, modelVersion.Version), prefix)
	}

	objectInfoList, err := h.data.Minio.ListObjects(bucketName, objectPrefix, false)
	if err != nil {
		return nil, err
	}

	modelInfoList := make([]*api.ModelInfo, 0)
	for _, m := range objectInfoList {
		modelInfo := &api.ModelInfo{}

		m.Name = m.Name[len(objectPrefix):]
		err = copier.Copy(modelInfo, m)
		if err != nil {
			err = errors.Errorf(err, errors.ErrorStructCopy)
			return nil, err
		}

		modelInfoList = append(modelInfoList, modelInfo)
	}

	return &api.ListModelVersionFileReply{
		ModelInfoList: modelInfoList,
	}, nil
}

func (h *modelShareHandle) findModelVersionAccessMaxId(ctx context.Context, modelAccessId string) (string, error) {
	modelDao := h.data.ModelDao

	_, modelVersionAccessList, err := modelDao.ListModelVersionAccess(ctx, &model.ModelVersionAccessList{
		ModelAccessId:     modelAccessId,
		ModelVersionOrder: true,
		ModelVersionSort:  model.DESC,
		PageIndex:         1,
		PageSize:          1,
	})
	if err != nil {
		return "", err
	}
	if len(modelVersionAccessList) == 0 {
		err = errors.Errorf(nil, errors.ErrorDBFindEmpty)
		return "", err
	}

	return modelVersionAccessList[0].ModelVersion, nil
}
