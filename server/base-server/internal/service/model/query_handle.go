package model

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"strings"

	"server/common/errors"

	"server/common/log"

	"github.com/jinzhu/copier"
)

type ModelQueryHandle interface {
	// 查询预置模型列表
	ListPreModelHandle(ctx context.Context, req *api.ListPreModelRequest) (*api.ListPreModelReply, error)
	// 查询我的模型列表
	ListMyModelHandle(ctx context.Context, req *api.ListMyModelRequest) (*api.ListMyModelReply, error)
	// 查询公共模型列表
	ListCommModelHandle(ctx context.Context, req *api.ListCommModelRequest) (*api.ListCommModelReply, error)
	// 查询模型版本列表
	ListModelVersionHandle(ctx context.Context, req *api.ListModelVersionRequest) (*api.ListModelVersionReply, error)
	// 查询所有用户模型列表
	ListAllUserModelHandle(ctx context.Context, req *api.ListAllUserModelRequest) (*api.ListAllUserModelReply, error)
	// 查询公共模型版本列表
	ListCommModelVersionHandle(ctx context.Context, req *api.ListCommModelVersionRequest) (*api.ListCommModelVersionReply, error)
	// 查询模型详情
	QueryModelHandle(ctx context.Context, req *api.QueryModelRequest) (*api.QueryModelReply, error)
	// 查询模型版本详情
	QueryModelVersionHandle(ctx context.Context, req *api.QueryModelVersionRequest) (*api.QueryModelVersionReply, error)
}

type modelQueryHandle struct {
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewModelQueryHandle(conf *conf.Bootstrap, logger log.Logger, data *data.Data) ModelQueryHandle {
	return &modelQueryHandle{
		conf: conf,
		log:  log.NewHelper("ModelQueryHandle", logger),
		data: data,
	}
}

// 查询预置模型列表
func (h *modelQueryHandle) ListPreModelHandle(ctx context.Context, req *api.ListPreModelRequest) (*api.ListPreModelReply, error) {
	modelDao := h.data.ModelDao
	totalSize, modelList, err := modelDao.ListModel(ctx, &model.ModelList{
		IsPrefab:       true,
		CreatedAtOrder: true,
		CreatedAtSort:  model.DESC,
		PageIndex:      int(req.PageIndex),
		PageSize:       int(req.PageSize),
		SearchKey:      req.SearchKey,
		FrameWorkId:    req.FrameWorkId,
		CreatedAtGte:   req.CreatedAtGte,
		CreatedAtLt:    req.CreatedAtLt,
	})
	if err != nil {
		return nil, err
	}

	models := make([]*api.ModelDetail, 0)
	for _, m := range modelList {
		modelDetail, err := h.modelTransfer(m)
		if err != nil {
			return nil, err
		}

		models = append(models, modelDetail)
	}

	return &api.ListPreModelReply{
		TotalSize: totalSize,
		Models:    models,
	}, nil
}

// 查询我的模型列表
func (h *modelQueryHandle) ListMyModelHandle(ctx context.Context, req *api.ListMyModelRequest) (*api.ListMyModelReply, error) {
	modelDao := h.data.ModelDao
	totalSize, modelList, err := modelDao.ListModel(ctx, &model.ModelList{
		IsPrefab:       false,
		SpaceId:        req.SpaceId,
		UserId:         req.UserId,
		CreatedAtOrder: true,
		CreatedAtSort:  model.DESC,
		PageIndex:      int(req.PageIndex),
		PageSize:       int(req.PageSize),
		SearchKey:      req.SearchKey,
		FrameWorkId:    req.FrameWorkId,
		CreatedAtGte:   req.CreatedAtGte,
		CreatedAtLt:    req.CreatedAtLt,
	})
	if err != nil {
		return nil, err
	}

	models := make([]*api.ModelDetail, 0)
	for _, m := range modelList {
		modelDetail, err := h.modelTransfer(m)
		if err != nil {
			return nil, err
		}

		models = append(models, modelDetail)
	}

	return &api.ListMyModelReply{
		TotalSize: totalSize,
		Models:    models,
	}, nil
}

// 查询公共模型列表
func (h *modelQueryHandle) ListCommModelHandle(ctx context.Context, req *api.ListCommModelRequest) (*api.ListCommModelReply, error) {
	modelDao := h.data.ModelDao
	pageIndex := 0
	pageSize := 0

	if req.SearchKey == "" {
		pageIndex = int(req.PageIndex)
		pageSize = int(req.PageSize)
	}

	totalSize, modelAccessList, err := modelDao.ListModelAccess(ctx, &model.ModelAccessList{
		SpaceIds:       []string{req.SpaceId},
		CreatedAtOrder: true,
		CreatedAtSort:  model.DESC,
		PageIndex:      pageIndex,
		PageSize:       pageSize,
		CreatedAtGte:   req.CreatedAtGte,
		CreatedAtLt:    req.CreatedAtLt,
		FrameWorkId:    req.FrameWorkId,
	})
	if err != nil {
		return nil, err
	}

	models := make([]*api.ModelDetail, 0)
	count := 0
	for _, mc := range modelAccessList {
		m, err := modelDao.GetModel(ctx, mc.ModelId)
		if err != nil {
			continue
		}

		m.LatestVersion = mc.LatestModelVersion
		modelDetail, err := h.modelTransfer(m)
		if err != nil {
			continue
		}

		// 模糊搜索
		if req.SearchKey != "" {
			if !strings.Contains(m.ModelName, req.SearchKey) &&
				!strings.Contains(m.ModelDescript, req.SearchKey) {
				continue
			}
			if count >= int(req.PageIndex-1)*int(req.PageSize) &&
				count < int(req.PageIndex)*int(req.PageSize) {
				models = append(models, modelDetail)
			}
			count++
		} else {
			models = append(models, modelDetail)
		}
	}

	if count != 0 {
		totalSize = int64(count)
	}

	return &api.ListCommModelReply{
		TotalSize: totalSize,
		Models:    models,
	}, nil
}

// 查询所有用户模型列表
func (h *modelQueryHandle) ListAllUserModelHandle(ctx context.Context, req *api.ListAllUserModelRequest) (*api.ListAllUserModelReply, error) {
	modelDao := h.data.ModelDao
	totalSize, modelList, err := modelDao.ListModel(ctx, &model.ModelList{
		IsPrefab:       false,
		CreatedAtOrder: true,
		SpaceIdOrder:   true,
		SpaceIdSort:    model.DESC,
		UserIdOrder:    true,
		UserIdSort:     model.DESC,
		CreatedAtSort:  model.DESC,
		PageIndex:      int(req.PageIndex),
		PageSize:       int(req.PageSize),
		SearchKey:      req.SearchKey,
		UserId:         req.UserId,
		SpaceId:        req.SpaceId,
		CreatedAtGte:   req.CreatedAtGte,
		CreatedAtLt:    req.CreatedAtLt,
	})
	if err != nil {
		return nil, err
	}

	models := make([]*api.ModelDetail, 0)
	for _, m := range modelList {
		modelDetail, err := h.modelTransfer(m)
		if err != nil {
			return nil, err
		}

		models = append(models, modelDetail)
	}

	return &api.ListAllUserModelReply{
		TotalSize: totalSize,
		Models:    models,
	}, nil
}

// 查询模型版本列表
func (h *modelQueryHandle) ListModelVersionHandle(ctx context.Context, req *api.ListModelVersionRequest) (*api.ListModelVersionReply, error) {
	modelDao := h.data.ModelDao

	_, err := modelDao.GetModel(ctx, req.ModelId)
	if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
		return nil, err
	}
	if errors.IsError(errors.ErrorDBFindEmpty, err) {
		return &api.ListModelVersionReply{
			TotalSize:     0,
			ModelVersions: nil,
		}, nil
	}

	totalSize, modelVersionList, err := modelDao.ListModelVersion(ctx, &model.ModelVersionList{
		ModelId:      req.ModelId,
		VersionOrder: true,
		VersionSort:  model.DESC,
		PageIndex:    int(req.PageIndex),
		PageSize:     int(req.PageSize),
	})
	if err != nil {
		return nil, err
	}

	modelVersions := make([]*api.VersionDetail, 0)
	for _, mv := range modelVersionList {
		versionDetail, err := h.modelVersionTransfer(mv)
		if err != nil {
			return nil, err
		}

		modelVersions = append(modelVersions, versionDetail)
	}

	return &api.ListModelVersionReply{
		TotalSize:     totalSize,
		ModelVersions: modelVersions,
	}, nil
}

// 查询公共模型版本列表
func (h *modelQueryHandle) ListCommModelVersionHandle(ctx context.Context, req *api.ListCommModelVersionRequest) (*api.ListCommModelVersionReply, error) {
	modelDao := h.data.ModelDao

	_, modelAccesses, err := modelDao.ListModelAccess(ctx, &model.ModelAccessList{
		SpaceIds: []string{req.SpaceId},
		ModelIds: []string{req.ModelId},
	})
	if err != nil {
		return nil, err
	}
	if len(modelAccesses) == 0 {
		return &api.ListCommModelVersionReply{
			TotalSize:     0,
			ModelVersions: nil,
		}, nil
	}

	totalSize, modelVersionAccessList, err := modelDao.ListModelVersionAccess(ctx, &model.ModelVersionAccessList{
		ModelAccessId:     modelAccesses[0].Id,
		ModelVersionOrder: true,
		ModelVersionSort:  model.DESC,
		PageIndex:         int(req.PageIndex),
		PageSize:          int(req.PageSize),
	})
	if err != nil {
		return nil, err
	}

	modelVersions := make([]*api.VersionDetail, 0)
	for _, mc := range modelVersionAccessList {
		m, err := h.QueryModelVersionHandle(ctx, &api.QueryModelVersionRequest{
			ModelId: mc.ModelId,
			Version: mc.ModelVersion,
		})
		if err != nil {
			continue
		}

		modelVersions = append(modelVersions, m.ModelVersion)
	}

	return &api.ListCommModelVersionReply{
		TotalSize:     totalSize,
		ModelVersions: modelVersions,
	}, nil
}

// 查询模型详情
func (h *modelQueryHandle) QueryModelHandle(ctx context.Context, req *api.QueryModelRequest) (*api.QueryModelReply, error) {
	modelDao := h.data.ModelDao

	m, err := modelDao.GetModel(ctx, req.ModelId)
	if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
		return nil, err
	}
	if errors.IsError(errors.ErrorDBFindEmpty, err) {
		return &api.QueryModelReply{
			Model: nil,
		}, nil
	}

	modelDetail, err := h.modelTransfer(m)
	if err != nil {
		return nil, err
	}

	return &api.QueryModelReply{
		Model: modelDetail,
	}, nil
}

// 查询模型版本详情
func (h *modelQueryHandle) QueryModelVersionHandle(ctx context.Context, req *api.QueryModelVersionRequest) (*api.QueryModelVersionReply, error) {
	modelDao := h.data.ModelDao

	m, err := modelDao.GetModel(ctx, req.ModelId)
	if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
		return nil, err
	}
	if errors.IsError(errors.ErrorDBFindEmpty, err) {
		return &api.QueryModelVersionReply{
			ModelVersion: nil,
		}, nil
	}

	mv, err := modelDao.GetModelVersion(ctx, req.ModelId, req.Version)
	if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
		return nil, err
	}
	if errors.IsError(errors.ErrorDBFindEmpty, err) {
		return &api.QueryModelVersionReply{
			ModelVersion: nil,
		}, nil
	}

	modelDetail, err := h.modelTransfer(m)
	if err != nil {
		return nil, err
	}
	versionDetail, err := h.modelVersionTransfer(mv)
	if err != nil {
		return nil, err
	}

	return &api.QueryModelVersionReply{
		Model:        modelDetail,
		ModelVersion: versionDetail,
	}, nil
}

func (h *modelQueryHandle) modelTransfer(model *model.Model) (*api.ModelDetail, error) {
	modelDetail := &api.ModelDetail{}

	err := copier.Copy(modelDetail, model)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorStructCopy)
		return nil, err
	}
	modelDetail.ModelId = model.Id
	modelDetail.CreatedAt = model.CreatedAt.Unix()
	return modelDetail, nil
}

func (h *modelQueryHandle) modelVersionTransfer(modelVersion *model.ModelVersion) (*api.VersionDetail, error) {
	modelVersionDetail := &api.VersionDetail{}

	err := copier.Copy(modelVersionDetail, modelVersion)
	if err != nil {
		err = errors.Errorf(err, errors.ErrorStructCopy)
		return nil, err
	}
	modelVersionDetail.CreatedAt = modelVersion.CreatedAt.Unix()

	return modelVersionDetail, nil
}
