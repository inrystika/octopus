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
)

type ModelAddHandle interface {
	// 新增我的模型
	AddMyModelHandle(ctx context.Context, req *api.AddMyModelRequest) (*api.AddMyModelReply, error)
	// 新增预置模型
	AddPreModelHandle(ctx context.Context, req *api.AddPreModelRequest) (*api.AddPreModelReply, error)
	// 新增预置模型版本
	AddPreModelVersionHandle(ctx context.Context, req *api.AddPreModelVersionRequest) (*api.AddPreModelVersionReply, error)
	// 上传预置模型版本
	UploadPreModelVersionHandle(ctx context.Context, req *api.UploadPreModelVersionRequest) (*api.UploadPreModelVersionReply, error)
	// 上传预置模型版本确认
	ConfirmUploadPreModelVersionHandle(ctx context.Context, req *api.ConfirmUploadPreModelVersionRequest) (*api.ConfirmUploadPreModelVersionReply, error)
}

type modelAddHandle struct {
	conf             *conf.Bootstrap
	log              *log.Helper
	data             *data.Data
	algorithmService api.AlgorithmServiceServer
}

func NewModelAddHandle(conf *conf.Bootstrap, logger log.Logger, data *data.Data, algorithmService api.AlgorithmServiceServer) ModelAddHandle {
	return &modelAddHandle{
		conf:             conf,
		log:              log.NewHelper("ModelAddHandle", logger),
		data:             data,
		algorithmService: algorithmService,
	}
}

const (
	CREATE_MIN_TIME_INTERVAL int64 = 3 // 创建的最短时间间隔，用于防止重复创建
)

// 新增我的模型
func (h *modelAddHandle) AddMyModelHandle(ctx context.Context, req *api.AddMyModelRequest) (*api.AddMyModelReply, error) {
	modelDao := h.data.ModelDao

	_, myModels, err := modelDao.ListModel(ctx, &model.ModelList{
		SpaceId:          req.SpaceId,
		UserId:           req.UserId,
		AlgorithmId:      req.AlgorithmId,
		AlgorithmVersion: req.AlgorithmVersion,
		IsPrefab:         false,
	})
	if err != nil {
		return nil, err
	}
	if len(myModels) == 0 {
		// 新增模型
		return h.addNewMyModel(ctx, req)
	}

	// 新增模型版本
	return h.addNewMyModelVersion(ctx, req, myModels[0])
}

// 新增预置模型
func (h *modelAddHandle) AddPreModelHandle(ctx context.Context, req *api.AddPreModelRequest) (*api.AddPreModelReply, error) {
	modelDao := h.data.ModelDao

	_, preModels, err := modelDao.ListModel(ctx, &model.ModelList{
		AlgorithmId:      req.AlgorithmId,
		AlgorithmVersion: req.AlgorithmVersion,
		IsPrefab:         true,
	})
	if err != nil {
		return nil, err
	}
	if len(preModels) != 0 {
		err = errors.Errorf(nil, errors.ErrorModelRepeat)
		return nil, err
	}

	algorithmReply, err := h.algorithmService.QueryAlgorithmVersion(ctx, &api.QueryAlgorithmVersionRequest{
		AlgorithmId: req.AlgorithmId,
		Version:     req.AlgorithmVersion,
	})
	if err != nil {
		return nil, err
	}
	if algorithmReply.Algorithm == nil {
		err := errors.Errorf(nil, errors.ErrorModelCreateAlgorithmNotExisted)
		return nil, err
	}

	modelId := utils.GetUUIDWithoutSeparator()
	modelVersionId := utils.GetUUIDWithoutSeparator()
	modelVersion := common.VersionStrBuild(1)
	modelName := algorithmReply.Algorithm.ModelName
	modelDescript := req.ModelDescript
	preModel, err := modelDao.CreateModel(ctx, &model.Model{
		Id:               modelId,
		IsPrefab:         true,
		AlgorithmId:      req.AlgorithmId,
		AlgorithmVersion: req.AlgorithmVersion,
		ModelName:        modelName,
		ModelDescript:    modelDescript,
		LatestVersion:    modelVersion,
		ModelVersions: []*model.ModelVersion{
			{
				Id:         modelVersionId,
				Version:    modelVersion,
				Descript:   modelDescript,
				FileStatus: model.FILESTATUS_INIT,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return &api.AddPreModelReply{
		ModelId:   preModel.Id,
		Version:   preModel.LatestVersion,
		CreatedAt: preModel.CreatedAt.Unix(),
	}, nil
}

// 新增预置模型版本
func (h *modelAddHandle) AddPreModelVersionHandle(ctx context.Context, req *api.AddPreModelVersionRequest) (*api.AddPreModelVersionReply, error) {
	modelDao := h.data.ModelDao

	preModel, err := modelDao.GetModel(ctx, req.ModelId)
	if err != nil {
		return nil, err
	}
	if !preModel.IsPrefab {
		err := errors.Errorf(nil, errors.ErrorModelNoPermission)
		return nil, err
	}

	_, err = modelDao.GetModelVersion(ctx, preModel.Id, preModel.LatestVersion)
	if err != nil {
		return nil, err
	}

	if (time.Now().Unix() - preModel.UpdatedAt.Unix()) < CREATE_MIN_TIME_INTERVAL {
		err := errors.Errorf(nil, errors.ErrorModelVersionRepeat)
		return nil, err
	}

	latestVersion, err := common.VersionStrParse(preModel.LatestVersion)
	if err != nil {
		return nil, err
	}

	preModelVersion, err := modelDao.CreateModelVersion(ctx, &model.ModelVersion{
		Id:         utils.GetUUIDWithoutSeparator(),
		ModelId:    req.ModelId,
		Version:    common.VersionStrBuild(latestVersion + 1),
		Descript:   req.Descript,
		FileStatus: model.FILESTATUS_INIT,
	})
	if err != nil {
		return nil, err
	}

	preModel.LatestVersion = preModelVersion.Version
	err = modelDao.UpdateModel(ctx, preModel)
	if err != nil {
		return nil, err
	}

	return &api.AddPreModelVersionReply{
		ModelId:   preModelVersion.ModelId,
		Version:   preModelVersion.Version,
		CreatedAt: preModelVersion.CreatedAt.Unix(),
	}, nil
}

// 上传预置模型版本
func (h *modelAddHandle) UploadPreModelVersionHandle(ctx context.Context, req *api.UploadPreModelVersionRequest) (*api.UploadPreModelVersionReply, error) {
	modelDao := h.data.ModelDao

	preModel, err := modelDao.GetModel(ctx, req.ModelId)
	if err != nil {
		return nil, err
	}
	if !preModel.IsPrefab {
		err := errors.Errorf(nil, errors.ErrorModelNoPermission)
		return nil, err
	}

	preModelVersion, err := modelDao.GetModelVersion(ctx, req.ModelId, req.Version)
	if err != nil {
		return nil, err
	}
	if preModelVersion.FileStatus != model.FILESTATUS_FINISH {
		// 获取临时上传链接
		bucektName := common.GetMinioBucket()
		objectName := common.GetMinioUploadModelObject(preModelVersion.ModelId, preModelVersion.Version, req.FileName)
		url, err := h.data.Minio.PresignedUploadObject(bucektName, objectName, req.Domain)
		if err != nil {
			return nil, err
		}

		return &api.UploadPreModelVersionReply{
			UploadUrl: url.String(),
		}, nil
	}

	// 已经上传成功了，就不要再获取上传接口了
	err = errors.Errorf(nil, errors.ErrorModelVersionFileExisted)
	return nil, err
}

// 上传预置模型版本确认
func (h *modelAddHandle) ConfirmUploadPreModelVersionHandle(ctx context.Context, req *api.ConfirmUploadPreModelVersionRequest) (*api.ConfirmUploadPreModelVersionReply, error) {
	modelDao := h.data.ModelDao

	preModel, err := modelDao.GetModel(ctx, req.ModelId)
	if err != nil {
		return nil, err
	}
	if !preModel.IsPrefab {
		err := errors.Errorf(nil, errors.ErrorModelNoPermission)
		return nil, err
	}

	preModelVersion, err := modelDao.GetModelVersion(ctx, req.ModelId, req.Version)
	if err != nil {
		return nil, err
	}
	if preModelVersion.FileStatus == model.FILESTATUS_FINISH {
		h.log.Infof(ctx, "AddPreConfirmHandle file always upload")
		return &api.ConfirmUploadPreModelVersionReply{
			UpdatedAt: preModelVersion.UpdatedAt.Unix(),
		}, nil
	}

	// 看下文件在不在
	fromBucketName := common.GetMinioBucket()
	fromObjectName := common.GetMinioUploadModelObject(preModelVersion.ModelId, preModelVersion.Version, req.FileName)
	isExist, err := h.data.Minio.ObjectExist(fromBucketName, fromObjectName)
	if err != nil {
		return nil, err
	}
	if !isExist {
		err := errors.Errorf(nil, errors.ErrorModelVersionFileNotFound)
		return nil, err
	}

	preModelVersion.FileStatus = model.FILESTATUS_UPLOGADING
	err = modelDao.UpdateModelVersion(ctx, preModelVersion)
	if err != nil {
		return nil, err
	}

	fromPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, fromBucketName, fromObjectName)

	// 解压拷贝
	go func() {
		toBucketName := common.GetMinioBucket()
		toObjectName := common.GetMinioPreModelObject(preModelVersion.ModelId, preModelVersion.Version)
		toPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, toBucketName, toObjectName)
		h.log.Infof(ctx, "begin to Unzip, from: %s, to:%s", fromPath, toPath)
		startT := time.Now()
		err := utils.Unzip(fromPath, toPath)
		if err != nil {
			preModelVersion.FileStatus = model.FILESTATUS_FAILED
			h.log.Errorw(ctx, err)
		} else {
			preModelVersion.FileStatus = model.FILESTATUS_FINISH
			h.log.Infof(ctx, "Unzip success, from: %s, to:%s, cost time: %d", fromPath, toPath, time.Since(startT))
		}

		err = modelDao.UpdateModelVersion(ctx, preModelVersion)
		if err != nil {
			return
		}
	}()
	// 拷贝压缩包到特定位置，用于下载
	go func() {
		toBucketName := common.GetMinioBucket()
		toObjectName := common.GetMinioDownloadModelObject(preModelVersion.ModelId, preModelVersion.Version, fmt.Sprintf("%s-%s.zip", preModel.ModelName, preModelVersion.Version))
		toPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, toBucketName, toObjectName)
		h.log.Infof(ctx, "begin to CopyFile, from: %s, to:%s", fromPath, toPath)
		startT := time.Now()
		err := utils.CopyFile(fromPath, toPath)
		if err != nil {
			h.log.Errorw(ctx, err)
			return
		}

		h.log.Infof(ctx, "CopyFile success, from: %s, to:%s, cost time: %d", fromPath, toPath, time.Since(startT))
	}()

	return &api.ConfirmUploadPreModelVersionReply{
		UpdatedAt: time.Now().Unix(),
	}, nil
}

// 新增新的我的模型
func (h *modelAddHandle) addNewMyModel(ctx context.Context, req *api.AddMyModelRequest) (*api.AddMyModelReply, error) {
	modelDao := h.data.ModelDao

	algorithmReply, err := h.algorithmService.QueryAlgorithmVersion(ctx, &api.QueryAlgorithmVersionRequest{
		AlgorithmId: req.AlgorithmId,
		Version:     req.AlgorithmVersion,
	})
	if err != nil {
		return nil, err
	}
	if algorithmReply.Algorithm == nil {
		err := errors.Errorf(nil, errors.ErrorModelCreateAlgorithmNotExisted)
		return nil, err
	}

	modelId := utils.GetUUIDWithoutSeparator()
	modelVersionId := utils.GetUUIDWithoutSeparator()
	modelVersion := common.VersionStrBuild(1)
	modelName := algorithmReply.Algorithm.ModelName
	modelDescript := ""

	myModel, err := modelDao.CreateModel(ctx, &model.Model{
		Id:               modelId,
		SpaceId:          req.SpaceId,
		UserId:           req.UserId,
		IsPrefab:         false,
		AlgorithmId:      req.AlgorithmId,
		AlgorithmVersion: req.AlgorithmVersion,
		ModelName:        modelName,
		ModelDescript:    modelDescript,
		LatestVersion:    modelVersion,
		ModelVersions: []*model.ModelVersion{
			{
				Id:         modelVersionId,
				Version:    modelVersion,
				Descript:   modelDescript,
				FileStatus: model.FILESTATUS_UPLOGADING,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// 拷贝
	go func() {
		toBucketName := common.GetMinioBucket()
		toObjectName := common.GetMinioModelObject(myModel.SpaceId, myModel.UserId, myModel.Id, myModel.LatestVersion)
		toPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, toBucketName, toObjectName)
		h.log.Infof(ctx, "begin to CopyDir, fromPath:%s, toPath:%s", req.FilePath, toPath)
		err := utils.CopyDir(req.FilePath, toPath)
		if err != nil {
			myModel.ModelVersions[0].FileStatus = model.FILESTATUS_FAILED
			h.log.Errorw(ctx, err)
		} else {
			myModel.ModelVersions[0].FileStatus = model.FILESTATUS_FINISH
			h.log.Infof(ctx, "CopyDir Success, fromPath:%s, toPath:%s", req.FilePath, toPath)
		}

		err = modelDao.UpdateModelVersion(ctx, myModel.ModelVersions[0])
		if err != nil {
			h.log.Errorw(ctx, err)
			return
		}
	}()
	// 压缩拷贝
	go func() {
		toBucketName := common.GetMinioBucket()
		toObjectName := common.GetMinioDownloadModelObject(myModel.Id, myModel.LatestVersion, fmt.Sprintf("%s-%s.zip", myModel.ModelName, myModel.LatestVersion))
		toPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, toBucketName, toObjectName)
		h.log.Infof(ctx, "begin to Zip, from: %s, to:%s", req.FilePath, toPath)
		startT := time.Now()
		err = utils.Zip(req.FilePath, toPath)
		if err != nil {
			h.log.Errorw(ctx, err)
			return
		}

		h.log.Infof(ctx, "Zip success, from: %s, to:%s, cost time: %d", req.FilePath, toPath, time.Since(startT))
	}()

	return &api.AddMyModelReply{
		ModelId:   myModel.Id,
		Version:   myModel.LatestVersion,
		CreatedAt: myModel.CreatedAt.Unix(),
	}, nil
}

// 新增新的我的模型版本
func (h *modelAddHandle) addNewMyModelVersion(ctx context.Context, req *api.AddMyModelRequest, myModel *model.Model) (*api.AddMyModelReply, error) {
	modelDao := h.data.ModelDao

	myLatestModelVersion, err := modelDao.GetModelVersion(ctx, myModel.Id, myModel.LatestVersion)
	if err != nil {
		return nil, err
	}

	if myLatestModelVersion.FileStatus != model.FILESTATUS_FINISH {
		// 上个版本传失败了，那直接在这版本搞
		// 拷贝
		go func() {
			toBucketName := common.GetMinioBucket()
			toObjectName := common.GetMinioModelObject(myModel.SpaceId, myModel.UserId, myLatestModelVersion.ModelId, myLatestModelVersion.Version)
			toPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, toBucketName, toObjectName)

			h.log.Infof(ctx, "begin to CopyDir, fromPath:%s, toPath:%s", req.FilePath, toPath)
			err := utils.CopyDir(req.FilePath, toPath)
			if err != nil {
				myLatestModelVersion.FileStatus = model.FILESTATUS_FAILED
				h.log.Errorw(ctx, err)
			} else {
				myLatestModelVersion.FileStatus = model.FILESTATUS_FINISH
				h.log.Infof(ctx, "CopyDir Success, fromPath:%s, toPath:%s", req.FilePath, toPath)
			}

			err = modelDao.UpdateModelVersion(ctx, myLatestModelVersion)
			if err != nil {
				h.log.Errorw(ctx, err)
				return
			}
		}()
		// 压缩拷贝
		go func() {
			toBucketName := common.GetMinioBucket()
			toObjectName := common.GetMinioDownloadModelObject(myModel.Id, myModel.LatestVersion, fmt.Sprintf("%s-%s.zip", myModel.ModelName, myModel.LatestVersion))
			toPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, toBucketName, toObjectName)

			h.log.Infof(ctx, "begin to Zip, from: %s, to:%s", req.FilePath, toPath)
			startT := time.Now()
			err = utils.Zip(req.FilePath, toPath)
			if err != nil {
				h.log.Errorw(ctx, err)
				return
			}

			h.log.Infof(ctx, "Zip success, from: %s, to:%s, cost time: %d", req.FilePath, toPath, time.Since(startT))
		}()

		return &api.AddMyModelReply{
			ModelId:   myLatestModelVersion.ModelId,
			Version:   myLatestModelVersion.Version,
			CreatedAt: myLatestModelVersion.CreatedAt.Unix(),
		}, nil
	}

	if (time.Now().Unix() - myLatestModelVersion.UpdatedAt.Unix()) < CREATE_MIN_TIME_INTERVAL {
		err := errors.Errorf(nil, errors.ErrorModelVersionRepeat)
		return nil, err
	}

	latestVersion, err := common.VersionStrParse(myLatestModelVersion.Version)
	if err != nil {
		return nil, err
	}

	myModelVersion, err := modelDao.CreateModelVersion(ctx, &model.ModelVersion{
		Id:         utils.GetUUIDWithoutSeparator(),
		ModelId:    myModel.Id,
		Version:    common.VersionStrBuild(latestVersion + 1),
		Descript:   "",
		FileStatus: model.FILESTATUS_UPLOGADING,
	})
	if err != nil {
		return nil, err
	}

	myModel.LatestVersion = myModelVersion.Version
	err = modelDao.UpdateModel(ctx, myModel)
	if err != nil {
		return nil, err
	}

	// 拷贝
	go func() {
		toBucketName := common.GetMinioBucket()
		toObjectName := common.GetMinioModelObject(myModel.SpaceId, myModel.UserId, myModelVersion.ModelId, myModelVersion.Version)
		toPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, toBucketName, toObjectName)

		h.log.Infof(ctx, "begin to CopyDir, fromPath:%s, toPath:%s", req.FilePath, toPath)
		err := utils.CopyDir(req.FilePath, toPath)
		if err != nil {
			myModelVersion.FileStatus = model.FILESTATUS_FAILED
			h.log.Errorw(ctx, err)
		} else {
			myModelVersion.FileStatus = model.FILESTATUS_FINISH
			h.log.Infof(ctx, "CopyDir Success, fromPath:%s, toPath:%s", req.FilePath, toPath)
		}

		err = modelDao.UpdateModelVersion(ctx, myModelVersion)
		if err != nil {
			h.log.Errorw(ctx, err)
			return
		}
	}()
	// 压缩拷贝
	go func() {
		toBucketName := common.GetMinioBucket()
		toObjectName := common.GetMinioDownloadModelObject(myModel.Id, myModel.LatestVersion, fmt.Sprintf("%s-%s.zip", myModel.ModelName, myModel.LatestVersion))
		toPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, toBucketName, toObjectName)

		h.log.Infof(ctx, "begin to Zip, from: %s, to:%s", req.FilePath, toPath)
		startT := time.Now()
		err = utils.Zip(req.FilePath, toPath)
		if err != nil {
			h.log.Errorw(ctx, err)
			return
		}

		h.log.Infof(ctx, "Zip success, from: %s, to:%s, cost time: %d", req.FilePath, toPath, time.Since(startT))
	}()

	return &api.AddMyModelReply{
		ModelId:   myModelVersion.ModelId,
		Version:   myModelVersion.Version,
		CreatedAt: myModelVersion.CreatedAt.Unix(),
	}, nil
}
