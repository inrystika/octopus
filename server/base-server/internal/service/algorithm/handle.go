package algorithm

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
	"strings"
	"time"

	"server/common/log"
)

const (
	CREATE_MIN_TIME_INTERVAL int64 = 3 // 创建的最短时间间隔，用于防止重复创建
)

const (
	FILESTATUS_INIT       uint8 = 1
	FILESTATUS_UPLOGADING uint8 = 2
	FILESTATUS_FINISH     uint8 = 3
	FILESTATUS_FAILED     uint8 = 4
)

type AlgorithmHandle interface {
	// 查询预置算法列表
	ListPreAlgorithmHandle(ctx context.Context, req *api.ListPreAlgorithmRequest) (*api.ListPreAlgorithmReply, error)
	// 查询我的算法列表
	ListMyAlgorithmHandle(ctx context.Context, req *api.ListMyAlgorithmRequest) (*api.ListMyAlgorithmReply, error)
	// 查询公共算法列表
	ListCommAlgorithmHandle(ctx context.Context, req *api.ListCommAlgorithmRequest) (*api.ListCommAlgorithmReply, error)
	// 查询算法版本列表
	ListAlgorithmVersionHandle(ctx context.Context, req *api.ListAlgorithmVersionRequest) (*api.ListAlgorithmVersionReply, error)
	// 查询公共算法版本列表
	ListCommAlgorithmVersionHandle(ctx context.Context, req *api.ListCommAlgorithmVersionRequest) (*api.ListCommAlgorithmVersionReply, error)
	// 查询所有用户算法列表
	ListAllUserAlgorithmHandle(ctx context.Context, req *api.ListAllUserAlgorithmRequest) (*api.ListAllUserAlgorithmReply, error)
	// 查询算法版本详情
	QueryAlgorithmVersionHandle(ctx context.Context, req *api.QueryAlgorithmVersionRequest) (*api.QueryAlgorithmVersionReply, error)
	// 批量查询算法
	BatchQueryAlgorithmHandle(ctx context.Context, req *api.BatchQueryAlgorithmRequest) (*api.BatchQueryAlgorithmReply, error)

	// 分享算法版本到公共算法
	ShareAlgorithmVersionHandle(ctx context.Context, req *api.ShareAlgorithmVersionRequest) (*api.ShareAlgorithmVersionReply, error)
	// 取消分享算法版本到公共算法
	CloseShareAlgorithmVersionHandle(ctx context.Context, req *api.CloseShareAlgorithmVersionRequest) (*api.CloseShareAlgorithmVersionReply, error)
	// 取消分享算法到公共算法
	CloseShareAlgorithmHandle(ctx context.Context, req *api.CloseShareAlgorithmRequest) (*api.CloseShareAlgorithmReply, error)
	// 取消分享算法版本到所有公共算法
	AllCloseShareAlgorithmVersionHandle(ctx context.Context, req *api.AllCloseShareAlgorithmVersionRequest) (*api.AllCloseShareAlgorithmVersionReply, error)
	// 取消分享算法到所有公共算法
	AllCloseShareAlgorithmHandle(ctx context.Context, req *api.AllCloseShareAlgorithmRequest) (*api.AllCloseShareAlgorithmReply, error)

	// 新增算法
	AddAlgorithmHandle(ctx context.Context, req *api.AddAlgorithmRequest) (*api.AddAlgorithmReply, error)
	// 上传算法文件
	UploadAlgorithmHandle(ctx context.Context, req *api.UploadAlgorithmRequest) (*api.UploadAlgorithmReply, error)
	// 上传算法确认
	ConfirmUploadAlgorithmHandle(ctx context.Context, req *api.ConfirmUploadAlgorithmRequest) (*api.ConfirmUploadAlgorithmReply, error)
	// 修改算法
	UpdateAlgorithmHandle(ctx context.Context, req *api.UpdateAlgorithmRequest) (*api.UpdateAlgorithmReply, error)

	// 复制算法版本
	CopyAlgorithmVersionHandle(ctx context.Context, req *api.CopyAlgorithmVersionRequest) (*api.CopyAlgorithmVersionReply, error)

	// 新增我的算法版本
	AddMyAlgorithmVersionHandle(ctx context.Context, req *api.AddMyAlgorithmVersionRequest) (*api.AddMyAlgorithmVersionReply, error)
	// 删除我的算法版本
	DeleteMyAlgorithmVersionHandle(ctx context.Context, req *api.DeleteMyAlgorithmVersionRequest) (*api.DeleteMyAlgorithmVersionReply, error)
	// 删除我的算法
	DeleteMyAlgorithmHandle(ctx context.Context, req *api.DeleteMyAlgorithmRequest) (*api.DeleteMyAlgorithmReply, error)
	// 新增预置算法版本
	AddPreAlgorithmVersionHandle(ctx context.Context, req *api.AddPreAlgorithmVersionRequest) (*api.AddPreAlgorithmVersionReply, error)
	// 删除预置算法版本
	DeletePreAlgorithmVersionHandle(ctx context.Context, req *api.DeletePreAlgorithmVersionRequest) (*api.DeletePreAlgorithmVersionReply, error)
	// 删除预置算法
	DeletePreAlgorithmHandle(ctx context.Context, req *api.DeletePreAlgorithmRequest) (*api.DeletePreAlgorithmReply, error)
	// 修改算法版本
	UpdateAlgorithmVersionHandle(ctx context.Context, req *api.UpdateAlgorithmVersionRequest) (*api.UpdateAlgorithmVersionReply, error)

	// 压缩算法版本包
	DownloadAlgorithmVersionCompressHandle(ctx context.Context, req *api.DownloadAlgorithmVersionCompressRequest) (*api.DownloadAlgorithmVersionCompressReply, error)
	// 下载算法版本
	DownloadAlgorithmVersionHandle(ctx context.Context, req *api.DownloadAlgorithmVersionRequest) (*api.DownloadAlgorithmVersionReply, error)
}

type algorithmHandle struct {
	conf         *conf.Bootstrap
	log          *log.Helper
	data         *data.Data
	lableService api.LableServiceServer
}

func NewAlgorithmHandle(conf *conf.Bootstrap, logger log.Logger, data *data.Data, lableService api.LableServiceServer) AlgorithmHandle {
	return &algorithmHandle{
		conf:         conf,
		log:          log.NewHelper("AlgorithmHandle", logger),
		data:         data,
		lableService: lableService,
	}
}

// 查询预置算法列表
func (h *algorithmHandle) ListPreAlgorithmHandle(ctx context.Context, req *api.ListPreAlgorithmRequest) (*api.ListPreAlgorithmReply, error) {
	algorithmDao := h.data.AlgorithmDao
	totalSize, algorithmList, err := algorithmDao.ListAlgorithm(ctx, &model.AlgorithmList{
		IsPrefab:         true,
		CreatedAtOrder:   true,
		CreatedAtSort:    model.DESC,
		PageIndex:        int(req.PageIndex),
		PageSize:         int(req.PageSize),
		SearchKey:        req.SearchKey,
		NameLike:         req.NameLike,
		SortBy:           req.SortBy,
		OrderBy:          req.OrderBy,
		AlgorithmVersion: req.AlgorithmVersion,
		FileStatus:       int(req.FileStatus),
		CreatedAtGte:     req.CreatedAtGte,
		CreatedAtLt:      req.CreatedAtLt,
	})
	if err != nil {
		return nil, err
	}

	algorithms := make([]*api.AlgorithmDetail, 0)
	for _, m := range algorithmList {
		algorithmVersion, err := algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{
			AlgorithmId: m.AlgorithmId,
			Version:     m.LatestVersion,
		})
		if err != nil {
			continue
		}
		var detail *api.AlgorithmDetail = h.transferAlgorithmDetail(ctx, m, algorithmVersion)
		detail.AlgorithmDescript = m.AlgorithmDescript
		algorithms = append(algorithms, detail)
	}

	return &api.ListPreAlgorithmReply{
		TotalSize:  totalSize,
		Algorithms: algorithms,
	}, nil
}

// 查询我的算法列表
func (h *algorithmHandle) ListMyAlgorithmHandle(ctx context.Context, req *api.ListMyAlgorithmRequest) (*api.ListMyAlgorithmReply, error) {
	algorithmDao := h.data.AlgorithmDao
	totalSize, algorithmList, err := algorithmDao.ListAlgorithm(ctx, &model.AlgorithmList{
		IsPrefab:         false,
		SpaceId:          req.SpaceId,
		UserId:           req.UserId,
		CreatedAtOrder:   true,
		CreatedAtSort:    model.DESC,
		PageIndex:        int(req.PageIndex),
		PageSize:         int(req.PageSize),
		SearchKey:        req.SearchKey,
		NameLike:         req.NameLike,
		SortBy:           req.SortBy,
		OrderBy:          req.OrderBy,
		AlgorithmVersion: req.AlgorithmVersion,
		FileStatus:       int(req.FileStatus),
		CreatedAtGte:     req.CreatedAtGte,
		CreatedAtLt:      req.CreatedAtLt,
	})
	if err != nil {
		return nil, err
	}

	algorithms := make([]*api.AlgorithmDetail, 0)
	for _, m := range algorithmList {
		algorithmVersion, err := algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{
			AlgorithmId: m.AlgorithmId,
			Version:     m.LatestVersion,
		})
		if err != nil {
			continue
		}

		var detail *api.AlgorithmDetail = h.transferAlgorithmDetail(ctx, m, algorithmVersion)
		detail.AlgorithmDescript = m.AlgorithmDescript
		algorithms = append(algorithms, detail)
	}

	return &api.ListMyAlgorithmReply{
		TotalSize:  totalSize,
		Algorithms: algorithms,
	}, nil
}

// 查询公共算法列表
func (h *algorithmHandle) ListCommAlgorithmHandle(ctx context.Context, req *api.ListCommAlgorithmRequest) (*api.ListCommAlgorithmReply, error) {
	algorithmDao := h.data.AlgorithmDao
	pageIndex := 0
	pageSize := 0

	if req.SearchKey == "" {
		pageIndex = int(req.PageIndex)
		pageSize = int(req.PageSize)
	}

	totalSize, algorithmAccessList, err := algorithmDao.ListAlgorithmAccess(ctx, &model.AlgorithmAccessList{
		SpaceId:          req.SpaceId,
		CreatedAtOrder:   true,
		CreatedAtSort:    model.DESC,
		PageIndex:        pageIndex,
		PageSize:         pageSize,
		SearchKey:        req.SearchKey,
		NameLike:         req.NameLike,
		SortBy:           req.SortBy,
		OrderBy:          req.OrderBy,
		AlgorithmVersion: req.AlgorithmVersion,
		FileStatus:       int(req.FileStatus),
		CreatedAtGte:     req.CreatedAtGte,
		CreatedAtLt:      req.CreatedAtLt,
	})
	if err != nil {
		return nil, err
	}

	algorithms := make([]*api.AlgorithmDetail, 0)
	count := 0
	for _, mc := range algorithmAccessList {
		m, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
			AlgorithmId: mc.AlgorithmId,
		})
		if err != nil {
			continue
		}

		algorithmVersion, err := algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{
			AlgorithmId: mc.AlgorithmId,
			Version:     mc.LatestAlgorithmVersion,
		})
		if err != nil {
			continue
		}

		var detail *api.AlgorithmDetail = h.transferAlgorithmDetail(ctx, m, algorithmVersion)
		detail.AlgorithmDescript = m.AlgorithmDescript

		if err != nil {
			continue
		}

		// 模糊搜索
		if req.SearchKey != "" {
			if !strings.Contains(m.AlgorithmName, req.SearchKey) &&
				!strings.Contains(algorithmVersion.AlgorithmDescript, req.SearchKey) {
				continue
			}
			if count >= int(req.PageIndex-1)*int(req.PageSize) &&
				count < int(req.PageIndex)*int(req.PageSize) {
				algorithms = append(algorithms, detail)
			}
			count++
		} else {
			algorithms = append(algorithms, detail)
		}
	}

	return &api.ListCommAlgorithmReply{
		TotalSize:  totalSize,
		Algorithms: algorithms,
	}, nil
}

// 查询算法版本列表
func (h *algorithmHandle) ListAlgorithmVersionHandle(ctx context.Context, req *api.ListAlgorithmVersionRequest) (*api.ListAlgorithmVersionReply, error) {
	algorithmDao := h.data.AlgorithmDao

	m, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
		AlgorithmId: req.AlgorithmId,
	})
	if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
		return nil, err
	}

	if errors.IsError(errors.ErrorDBFindEmpty, err) {
		return &api.ListAlgorithmVersionReply{
			TotalSize:  0,
			Algorithms: nil,
		}, nil
	}

	totalSize, algorithmVersionList, err := algorithmDao.ListAlgorithmVersion(ctx, &model.AlgorithmVersionList{
		AlgorithmId:  req.AlgorithmId,
		VersionOrder: true,
		VersionSort:  model.DESC,
		PageIndex:    int(req.PageIndex),
		PageSize:     int(req.PageSize),
		FileStatus:   int(req.FileStatus),
	})
	if err != nil {
		return nil, err
	}

	algorithmVersions := make([]*api.AlgorithmDetail, 0)
	for _, mv := range algorithmVersionList {
		algorithmVersions = append(algorithmVersions, h.transferAlgorithmDetail(ctx, m, mv))
	}

	return &api.ListAlgorithmVersionReply{
		TotalSize:  totalSize,
		Algorithms: algorithmVersions,
	}, nil
}

// 查询公共算法版本列表
func (h *algorithmHandle) ListCommAlgorithmVersionHandle(ctx context.Context, req *api.ListCommAlgorithmVersionRequest) (*api.ListCommAlgorithmVersionReply, error) {
	algorithmDao := h.data.AlgorithmDao

	algorithmAccess, err := algorithmDao.QueryAlgorithmAccess(ctx, &model.AlgorithmAccessQuery{
		SpaceId:     req.SpaceId,
		AlgorithmId: req.AlgorithmId,
	})
	if err != nil {
		return nil, err
	}

	totalSize, algorithmVersionAccessList, err := algorithmDao.ListAlgorithmAccessVersion(ctx, &model.AlgorithmAccessVersionList{
		AlgorithmAccessId:     algorithmAccess.Id,
		AlgorithmVersionOrder: true,
		AlgorithmVersionSort:  model.DESC,
		PageIndex:             int(req.PageIndex),
		PageSize:              int(req.PageSize),
		FileStatus:            int(req.FileStatus),
	})

	if err != nil {
		return nil, err
	}
	if len(algorithmVersionAccessList) == 0 {
		return &api.ListCommAlgorithmVersionReply{
			TotalSize:  totalSize,
			Algorithms: nil,
		}, nil
	}

	algorithms := make([]*api.AlgorithmDetail, 0)
	for _, mc := range algorithmVersionAccessList {
		m, err := h.QueryAlgorithmVersionHandle(ctx, &api.QueryAlgorithmVersionRequest{
			AlgorithmId: mc.AlgorithmId,
			Version:     mc.AlgorithmVersion,
		})
		if err != nil {
			continue
		}
		algorithms = append(algorithms, m.Algorithm)
	}

	return &api.ListCommAlgorithmVersionReply{
		TotalSize:  totalSize,
		Algorithms: algorithms,
	}, nil
}

// 查询所有用户算法列表
func (h *algorithmHandle) ListAllUserAlgorithmHandle(ctx context.Context, req *api.ListAllUserAlgorithmRequest) (*api.ListAllUserAlgorithmReply, error) {
	algorithmDao := h.data.AlgorithmDao
	totalSize, algorithmList, err := algorithmDao.ListAlgorithm(ctx, &model.AlgorithmList{
		IsPrefab:         false,
		CreatedAtOrder:   true,
		SpaceIdOrder:     true,
		SpaceIdSort:      model.DESC,
		UserIdOrder:      true,
		UserIdSort:       model.DESC,
		CreatedAtSort:    model.DESC,
		PageIndex:        int(req.PageIndex),
		PageSize:         int(req.PageSize),
		SearchKey:        req.SearchKey,
		NameLike:         req.NameLike,
		SortBy:           req.SortBy,
		OrderBy:          req.OrderBy,
		AlgorithmVersion: req.AlgorithmVersion,
		FileStatus:       int(req.FileStatus),
		CreatedAtGte:     req.CreatedAtGte,
		CreatedAtLt:      req.CreatedAtLt,
		UserId:           req.UserId,
		SpaceId:          req.SpaceId,
	})
	if err != nil {
		return nil, err
	}

	algorithms := make([]*api.AlgorithmDetail, 0)
	for _, m := range algorithmList {
		algorithmVersion, err := algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{
			AlgorithmId: m.AlgorithmId,
			Version:     m.LatestVersion,
		})
		if err != nil {
			continue
		}

		var detail *api.AlgorithmDetail = h.transferAlgorithmDetail(ctx, m, algorithmVersion)
		detail.AlgorithmDescript = m.AlgorithmDescript
		algorithms = append(algorithms, detail)
	}

	return &api.ListAllUserAlgorithmReply{
		TotalSize:  totalSize,
		Algorithms: algorithms,
	}, nil
}

// 查询算法版本详情
func (h *algorithmHandle) QueryAlgorithmVersionHandle(ctx context.Context, req *api.QueryAlgorithmVersionRequest) (*api.QueryAlgorithmVersionReply, error) {
	algorithmDao := h.data.AlgorithmDao

	m, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
		AlgorithmId: req.AlgorithmId,
	})
	if err != nil {
		return nil, err
	}

	mv, err := algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{
		AlgorithmId: req.AlgorithmId,
		Version:     req.Version,
	})
	if err != nil {
		return nil, err
	}

	_, accessAlgorithmVersions, err := algorithmDao.ListAlgorithmAccessVersion(ctx, &model.AlgorithmAccessVersionList{
		AlgorithmId:      req.AlgorithmId,
		AlgorithmVersion: req.Version,
	})
	if err != nil {
		return nil, err
	}

	return &api.QueryAlgorithmVersionReply{
		Algorithm:       h.transferAlgorithmDetail(ctx, m, mv),
		VersionAccesses: h.transferAlgorithmVersionAccess(accessAlgorithmVersions),
	}, nil
}

//  批量查询算法
func (h *algorithmHandle) BatchQueryAlgorithmHandle(ctx context.Context, req *api.BatchQueryAlgorithmRequest) (*api.BatchQueryAlgorithmReply, error) {
	algorithmDao := h.data.AlgorithmDao

	algorithmIdList := make([]*model.AlgorithmQuery, 0)
	for _, id := range req.AlgorithmId {
		algorithmIdList = append(algorithmIdList, &model.AlgorithmQuery{
			AlgorithmId: id,
		})
	}

	algorithms, err := algorithmDao.BatchQueryAlgorithm(ctx, &model.AlgorithmBatchQuery{
		List: algorithmIdList,
	})
	if err != nil {
		return nil, err
	}

	algorithmInfos := []*api.AlgorithmInfo{}

	for _, algorithm := range algorithms {
		info := &api.AlgorithmInfo{
			AlgorithmId:   algorithm.AlgorithmId,
			AlgorithmName: algorithm.AlgorithmName,
			CreatedAt:     algorithm.CreatedAt.Unix(),
		}
		algorithmInfos = append(algorithmInfos, info)
	}

	return &api.BatchQueryAlgorithmReply{
		Algorithms: algorithmInfos,
	}, nil
}

func (h *algorithmHandle) transferAlgorithmDetail(ctx context.Context, algorithm *model.Algorithm, version *model.AlgorithmVersion) *api.AlgorithmDetail {
	algorithmVersion := &api.AlgorithmDetail{}

	algorithmVersion.AlgorithmId = algorithm.AlgorithmId
	algorithmVersion.SpaceId = algorithm.SpaceId
	algorithmVersion.UserId = algorithm.UserId
	algorithmVersion.IsPrefab = algorithm.IsPrefab
	algorithmVersion.AlgorithmName = algorithm.AlgorithmName
	algorithmVersion.ModelName = algorithm.ModelName
	algorithmVersion.FileStatus = int64(version.FileStatus)
	algorithmVersion.LatestCompressed = version.LatestCompressed
	algorithmVersion.AlgorithmVersion = version.Version
	algorithmVersion.AlgorithmDescript = version.AlgorithmDescript
	algorithmVersion.CreatedAt = algorithm.CreatedAt.Unix()
	algorithmVersion.ApplyId = algorithm.ApplyId
	algorithmVersion.FrameworkId = algorithm.FrameworkId

	algorithmApply, err := h.lableService.GetLable(ctx, &api.GetLableRequest{Id: algorithmVersion.ApplyId})
	if err != nil {
		algorithmVersion.ApplyName = ""
	} else {
		algorithmVersion.ApplyName = algorithmApply.Lable.LableDesc
	}

	algorithmFramework, err := h.lableService.GetLable(ctx, &api.GetLableRequest{Id: algorithmVersion.FrameworkId})
	if err != nil {
		algorithmVersion.FrameworkName = ""
	} else {
		algorithmVersion.FrameworkName = algorithmFramework.Lable.LableDesc
	}

	bucketName := ""
	objectName := ""
	if algorithm.IsPrefab {
		bucketName = common.GetMinioBucket()
		objectName = common.GetMinioPreCodeObject(version.AlgorithmId, version.Version)
	} else {
		bucketName = common.GetMinioBucket()
		objectName = common.GetMinioCodeObject(algorithm.SpaceId, algorithm.UserId, version.AlgorithmId, version.Version)
	}
	algorithmVersion.Path = fmt.Sprintf("%s/%s", bucketName, objectName)
	return algorithmVersion
}

func (h *algorithmHandle) transferAlgorithmVersionAccess(algorithmAccessVersions []*model.AlgorithmAccessVersion) []*api.AlgorithmVersionAccess {

	algorithmVersionAccess := []*api.AlgorithmVersionAccess{}

	for _, as := range algorithmAccessVersions {
		av := &api.AlgorithmVersionAccess{}
		av.AlgorithmId = as.AlgorithmId
		av.SpaceId = as.SpaceId
		av.Version = as.AlgorithmVersion
		algorithmVersionAccess = append(algorithmVersionAccess, av)
	}
	return algorithmVersionAccess
}

// 分享算法版本到公共算法
func (h *algorithmHandle) ShareAlgorithmVersionHandle(ctx context.Context, req *api.ShareAlgorithmVersionRequest) (*api.ShareAlgorithmVersionReply, error) {
	algorithmDao := h.data.AlgorithmDao
	algorithmId := req.AlgorithmId
	version := req.Version

	for _, spaceId := range req.ShareSpaceIdList {

		algorithm, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
			AlgorithmId: algorithmId,
		})
		if err != nil {
			return nil, err
		}
		// 查可见算法信息
		algorithmAccess, err := algorithmDao.QueryAlgorithmAccess(ctx, &model.AlgorithmAccessQuery{
			SpaceId:     spaceId,
			AlgorithmId: algorithmId,
		})
		if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
			continue
		}
		if errors.IsError(errors.ErrorDBFindEmpty, err) {
			// 插入可见算法信息
			algorithmAccess, err = algorithmDao.AddAlgorithmAccess(ctx, &model.AlgorithmAccess{
				Id:                utils.GetUUIDWithoutSeparator(),
				AlgorithmId:       algorithmId,
				AlgorithmName:     algorithm.AlgorithmName,
				AlgorithmDescript: algorithm.AlgorithmDescript,
				SpaceId:           spaceId,
			})
			if err != nil {
				continue
			}
			//更新algorithm_access表的创建时间与源算法的创建时间保持一致。
			algorithmAccess.CreatedAt = algorithm.CreatedAt
			err = algorithmDao.UpdateAlgorithmAccess(ctx, algorithmAccess)
			if err != nil {
				continue
			}
		}

		algorithmAccessId := algorithmAccess.Id

		asv, err := algorithmDao.QueryAlgorithmAccessVersion(ctx, &model.AlgorithmAccessVersionQuery{
			AlgorithmAccessId: algorithmAccessId,
			AlgorithmVersion:  version,
		})

		if asv != nil && err == nil {
			err = errors.Errorf(nil, errors.ErrorAlgorithmAccessVersionExisted)
			return nil, err
		}

		// 插入可见算法版本信息
		algorithmAccessVersion, err := algorithmDao.AddAlgorithmAccessVersion(ctx, &model.AlgorithmAccessVersion{
			Id:                utils.GetUUIDWithoutSeparator(),
			AlgorithmAccessId: algorithmAccessId,
			AlgorithmVersion:  version,
			AlgorithmId:       algorithmId,
			AlgorithmName:     algorithm.AlgorithmName,
			SpaceId:           spaceId,
		})
		if err != nil {
			continue
		}

		//更新algorithm_access_version表的创建时间与源算法的创建时间保持一致。
		algorithmAccessVersion.CreatedAt = algorithm.CreatedAt
		err = algorithmDao.UpdateAlgorithmAccessVersion(ctx, algorithmAccessVersion)
		if err != nil {
			continue
		}

		// 更新可见算法信息
		if algorithmAccess.LatestAlgorithmVersion < version {
			algorithmAccess.LatestAlgorithmVersion = version
			err = algorithmDao.UpdateAlgorithmAccess(ctx, algorithmAccess)
			if err != nil {
				return nil, err
			}
		}
	}

	return &api.ShareAlgorithmVersionReply{
		SharedAt: time.Now().Unix(),
	}, nil
}

// 取消分享算法版本到公共算法
func (h *algorithmHandle) CloseShareAlgorithmVersionHandle(ctx context.Context, req *api.CloseShareAlgorithmVersionRequest) (*api.CloseShareAlgorithmVersionReply, error) {
	algorithmDao := h.data.AlgorithmDao
	algorithmId := req.AlgorithmId
	version := req.Version

	for _, spaceId := range req.ShareSpaceIdList {
		// 查可见算法信息
		algorithmAccess, err := algorithmDao.QueryAlgorithmAccess(ctx, &model.AlgorithmAccessQuery{
			SpaceId:     spaceId,
			AlgorithmId: algorithmId,
		})
		if err != nil {
			continue
		}

		algorithmAccessId := algorithmAccess.Id

		// 删除可见算法版本信息
		err = algorithmDao.DeleteAlgorithmAccessVersion(ctx, &model.AlgorithmAccessVersionDelete{
			AlgorithmAccessId: algorithmAccessId,
			AlgorithmVersion:  version,
		})
		if err != nil {
			continue
		}

		maxAlgorithmVersion, err := h.findAlgorithmVersionAccessMaxId(ctx, algorithmAccessId)
		if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
			continue
		}
		if err != nil && errors.IsError(errors.ErrorDBFindEmpty, err) {
			// 最后一个版本，那就都删了
			err := algorithmDao.DeleteAlgorithmAccess(ctx, &model.AlgorithmAccessDelete{
				SpaceId:     spaceId,
				AlgorithmId: algorithmId,
			})
			if err != nil {
				continue
			}
		} else {
			algorithmAccess.LatestAlgorithmVersion = maxAlgorithmVersion
			err = algorithmDao.UpdateAlgorithmAccess(ctx, algorithmAccess)
			if err != nil {
				return nil, err
			}
		}
	}

	return &api.CloseShareAlgorithmVersionReply{
		CloseSharedAt: time.Now().Unix(),
	}, nil
}

// 取消分享算法到公共算法
func (h *algorithmHandle) CloseShareAlgorithmHandle(ctx context.Context, req *api.CloseShareAlgorithmRequest) (*api.CloseShareAlgorithmReply, error) {
	algorithmDao := h.data.AlgorithmDao
	algorithmId := req.AlgorithmId

	for _, spaceId := range req.SpaceIdList {
		algorithmAccess, err := algorithmDao.QueryAlgorithmAccess(ctx, &model.AlgorithmAccessQuery{
			SpaceId:     spaceId,
			AlgorithmId: algorithmId,
		})
		if err != nil {
			continue
		}

		algorithmAccesslId := algorithmAccess.Id

		err = algorithmDao.BatchDeleteAlgorithmAccessVersion(ctx, &model.AlgorithmAccessVersionBatchDelete{
			AlgorithmAccessId: algorithmAccesslId,
		})
		if err != nil {
			continue
		}

		err = algorithmDao.DeleteAlgorithmAccess(ctx, &model.AlgorithmAccessDelete{
			SpaceId:     spaceId,
			AlgorithmId: algorithmId,
		})
		if err != nil {
			continue
		}
	}

	return &api.CloseShareAlgorithmReply{
		CloseSharedAt: time.Now().Unix(),
	}, nil
}

// 取消分享算法版本到所有公共算法
func (h *algorithmHandle) AllCloseShareAlgorithmVersionHandle(ctx context.Context, req *api.AllCloseShareAlgorithmVersionRequest) (*api.AllCloseShareAlgorithmVersionReply, error) {
	algorithmDao := h.data.AlgorithmDao
	algorithmId := req.AlgorithmId
	version := req.Version

	// 查该版本有多少spaceId可见
	_, algorithmVersionAccessList, err := algorithmDao.ListAlgorithmAccessVersion(ctx, &model.AlgorithmAccessVersionList{
		AlgorithmId:      algorithmId,
		AlgorithmVersion: version,
	})
	if err != nil {
		return nil, err
	}

	algorithmAccessIdList := make([]string, 0)
	for _, m := range algorithmVersionAccessList {
		algorithmAccessIdList = append(algorithmAccessIdList, m.AlgorithmAccessId)
	}
	algorithmAccessList, err := algorithmDao.BatchQueryAlgorithmAccessById(ctx, &model.AlgorithmAccessBatchQueryById{
		AlgorithmAccessIdList: algorithmAccessIdList,
	})
	if err != nil {
		return nil, err
	}

	spaceIdList := make([]string, 0)
	for _, sp := range algorithmAccessList {
		spaceIdList = append(algorithmAccessIdList, sp.SpaceId)
	}
	_, err = h.CloseShareAlgorithmVersionHandle(ctx, &api.CloseShareAlgorithmVersionRequest{
		AlgorithmId:      algorithmId,
		Version:          version,
		ShareSpaceIdList: spaceIdList,
	})
	if err != nil {
		return nil, err
	}

	return &api.AllCloseShareAlgorithmVersionReply{
		CloseSharedAt: time.Now().Unix(),
	}, nil
}

// 取消分享算法到所有公共算法
func (h *algorithmHandle) AllCloseShareAlgorithmHandle(ctx context.Context, req *api.AllCloseShareAlgorithmRequest) (*api.AllCloseShareAlgorithmReply, error) {
	algorithmDao := h.data.AlgorithmDao
	algorithmId := req.AlgorithmId

	// 查该版本有多少spaceId可见
	_, accessAlgorithms, err := algorithmDao.ListAlgorithmAccess(ctx, &model.AlgorithmAccessList{
		AlgorithmId: algorithmId,
	})
	if err != nil {
		return nil, err
	}

	spaceIdList := make([]string, 0)
	for _, m := range accessAlgorithms {
		spaceIdList = append(spaceIdList, m.SpaceId)
	}

	_, err = h.CloseShareAlgorithmHandle(ctx, &api.CloseShareAlgorithmRequest{
		SpaceIdList: spaceIdList,
		AlgorithmId: algorithmId,
	})
	if err != nil {
		return nil, err
	}

	return &api.AllCloseShareAlgorithmReply{
		CloseSharedAt: time.Now().Unix(),
	}, nil
}

// 新增算法
func (h *algorithmHandle) AddAlgorithmHandle(ctx context.Context, req *api.AddAlgorithmRequest) (*api.AddAlgorithmReply, error) {
	algorithmDao := h.data.AlgorithmDao

	algorithm, _ := algorithmDao.QueryAlgorithmByInfo(ctx, &model.AlgorithmQueryByInfo{
		AlgorithmName: req.AlgorithmName,
		UserId:        req.UserId,
		SpaceId:       req.SpaceId,
		IsPrefab:      req.IsPrefab,
	})

	if algorithm != nil {
		err := errors.Errorf(nil, errors.ErrorAlgorithmRepeat)
		return nil, err
	}

	algorithmId := utils.GetUUIDWithoutSeparator()
	algorithmVersionId := utils.GetUUIDWithoutSeparator()
	algorithmVersion := common.VersionStrBuild(1)
	fileStatus := FILESTATUS_INIT
	if req.IsEmpty {
		bucketName := common.GetMinioBucket()
		objectName := common.GetMinioCodeObject(req.SpaceId, req.UserId, algorithmId, algorithmVersion)
		path := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, bucketName, objectName)
		err := utils.CreateDir(path)
		if err != nil {
			return nil, err
		}
		fileStatus = FILESTATUS_FINISH
	}

	myAlgorithm, err := algorithmDao.AddAlgorithm(ctx, &model.Algorithm{
		AlgorithmId:       algorithmId,
		SpaceId:           req.SpaceId,
		UserId:            req.UserId,
		IsPrefab:          req.IsPrefab,
		AlgorithmName:     req.AlgorithmName,
		AlgorithmDescript: req.AlgorithmDescript,
		ModelName:         req.ModelName,
		LatestVersion:     algorithmVersion,
		ApplyId:           req.ApplyId,
		FrameworkId:       req.FrameworkId,
		AlgorithmVersions: []*model.AlgorithmVersion{
			{
				Id:                algorithmVersionId,
				Version:           algorithmVersion,
				AlgorithmDescript: req.AlgorithmDescript,
				FileStatus:        fileStatus,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// 检查数据类型id
	if req.ApplyId != "" {
		algorithmApply, err := h.lableService.GetLable(ctx, &api.GetLableRequest{Id: req.ApplyId})
		if err != nil {
			return nil, err
		}
		// 新增算法类型引用
		_, _ = h.lableService.IncreaseLableReferTimes(ctx, &api.IncreaseLableReferTimesRequest{Id: algorithmApply.Lable.Id})
	}
	// 检查框架id
	if req.FrameworkId != "" {
		algorithmFramework, err := h.lableService.GetLable(ctx, &api.GetLableRequest{Id: req.FrameworkId})
		if err != nil {
			return nil, err
		}
		// 新增算法框架引用
		_, _ = h.lableService.IncreaseLableReferTimes(ctx, &api.IncreaseLableReferTimesRequest{Id: algorithmFramework.Lable.Id})
	}

	return &api.AddAlgorithmReply{
		AlgorithmId: myAlgorithm.AlgorithmId,
		Version:     myAlgorithm.LatestVersion,
		CreatedAt:   myAlgorithm.CreatedAt.Unix(),
	}, nil
}

// 上传算法
func (h *algorithmHandle) UploadAlgorithmHandle(ctx context.Context, req *api.UploadAlgorithmRequest) (*api.UploadAlgorithmReply, error) {
	algorithmDao := h.data.AlgorithmDao

	algorithm, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
		AlgorithmId: req.AlgorithmId,
	})

	if err != nil {
		return nil, err
	}

	if !algorithm.IsPrefab {
		if algorithm.UserId != req.UserId {
			err := errors.Errorf(nil, errors.ErrorAlgorithmNotMy)
			return nil, err
		}
		if req.Version != common.VersionStrBuild(1) {
			err := errors.Errorf(nil, errors.ErrorAlgorithmVersionUploadAuth)
			return nil, err
		}
	}

	version := ""
	if algorithm.IsPrefab {
		version = algorithm.LatestVersion
	} else {
		version = common.VersionStrBuild(1)
	}

	algorithmVersion, err := algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{
		AlgorithmId: req.AlgorithmId,
		Version:     version,
	})
	if err != nil {
		return nil, err
	}
	if algorithmVersion.FileStatus != FILESTATUS_FINISH {
		bucektName := common.GetMinioBucket()
		objectName := common.GetMinioUploadCodeObject(algorithm.AlgorithmId, version, req.FileName)

		url, err := h.data.Minio.PresignedUploadObject(bucektName, objectName, req.Domain)
		if err != nil {
			return nil, err
		}

		return &api.UploadAlgorithmReply{
			UploadUrl: url.String(),
		}, nil
	}

	// 已经上传成功了，就不要再获取上传接口了
	err = errors.Errorf(nil, errors.ErrorAlgorithmVersionFileExisted)
	return nil, err
}

// 上传算法确认
func (h *algorithmHandle) ConfirmUploadAlgorithmHandle(ctx context.Context, req *api.ConfirmUploadAlgorithmRequest) (*api.ConfirmUploadAlgorithmReply, error) {
	algorithmDao := h.data.AlgorithmDao

	myAlgorithm, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
		AlgorithmId: req.AlgorithmId,
	})
	if err != nil {
		return nil, err
	}

	if !myAlgorithm.IsPrefab {
		if myAlgorithm.UserId != req.UserId {
			err := errors.Errorf(nil, errors.ErrorAlgorithmNotMy)
			return nil, err
		}
	}

	myAlgorithmVersion, err := algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{
		AlgorithmId: req.AlgorithmId,
		Version:     req.Version,
	})
	if err != nil {
		return nil, err
	}

	bucketName := common.GetMinioBucket()
	objectName := common.GetMinioUploadCodeObject(myAlgorithm.AlgorithmId, myAlgorithmVersion.Version, req.FileName)

	fromPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, bucketName, objectName)

	// 看下文件在不在
	isExist, err := h.data.Minio.ObjectExist(bucketName, objectName)
	if err != nil {
		return nil, err
	}
	if !isExist {
		err := errors.Errorf(nil, errors.ErrorAlgorithmVersionFileNotFound)
		h.log.Errorw(ctx, err)
		return nil, err
	}

	myAlgorithmVersion.FileStatus = FILESTATUS_UPLOGADING
	err = algorithmDao.UpdateAlgorithmVersion(ctx, myAlgorithmVersion)
	if err != nil {
		return nil, err
	}
	toBucketName := ""
	toObjectName := ""

	if myAlgorithm.IsPrefab {
		toBucketName = common.GetMinioBucket()
		toObjectName = common.GetMinioPreCodeObject(myAlgorithmVersion.AlgorithmId, myAlgorithmVersion.Version)
	} else {
		toBucketName = common.GetMinioBucket()
		toObjectName = common.GetMinioCodeObject(myAlgorithm.SpaceId, myAlgorithm.UserId, myAlgorithmVersion.AlgorithmId, myAlgorithmVersion.Version)
	}
	toPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, toBucketName, toObjectName)
	// 解压
	go func() {
		// 删除算法压缩包临时文件
		defer func() {
			go utils.HandlePanicBG(func(i ...interface{}) {
				h.data.Redis.SAddMinioRemovingObject(bucketName + "-" + objectName)
				defer h.data.Redis.SRemMinioRemovingObject(bucketName + "-" + objectName)
				h.data.Minio.RemoveObject(bucketName, objectName)
			})()
		}()
		err := utils.Unzip(fromPath, toPath)
		if err != nil {
			myAlgorithmVersion.FileStatus = FILESTATUS_FAILED
		} else {
			myAlgorithmVersion.FileStatus = FILESTATUS_FINISH
		}

		err = algorithmDao.UpdateAlgorithmVersion(ctx, myAlgorithmVersion)
		if err != nil {
			return
		}
	}()

	return &api.ConfirmUploadAlgorithmReply{
		UpdatedAt: time.Now().Unix(),
	}, nil
}

// 修改算法
func (h *algorithmHandle) UpdateAlgorithmHandle(ctx context.Context, req *api.UpdateAlgorithmRequest) (*api.UpdateAlgorithmReply, error) {
	algorithmDao := h.data.AlgorithmDao
	algorithm, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{AlgorithmId: req.AlgorithmId})
	if err != nil {
		return nil, err
	}
	if algorithm.SpaceId != req.SpaceId || algorithm.UserId != algorithm.UserId || algorithm.IsPrefab != req.IsPrefab {
		return nil, errors.Errorf(nil, errors.ErrorFindAlgorithmAuthWrong)
	}

	// 减少算法类型引用
	_, _ = h.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: algorithm.ApplyId})
	// 减少算法框架引用
	_, _ = h.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: algorithm.FrameworkId})

	algorithm.ApplyId = req.ApplyId
	algorithm.FrameworkId = req.FrameworkId
	algorithm.AlgorithmDescript = req.AlgorithmDescript
	algorithm.ModelName = req.ModelName
	err = algorithmDao.UpdateAlgorithm(ctx, algorithm)
	if err != nil {
		return nil, err
	}

	// 增加算法类型引用
	_, _ = h.lableService.IncreaseLableReferTimes(ctx, &api.IncreaseLableReferTimesRequest{Id: algorithm.ApplyId})
	// 增加算法框架引用
	_, _ = h.lableService.IncreaseLableReferTimes(ctx, &api.IncreaseLableReferTimesRequest{Id: algorithm.FrameworkId})

	return &api.UpdateAlgorithmReply{UpdatedAt: time.Now().Unix()}, nil
}

// 新增我的算法版本
func (h *algorithmHandle) AddMyAlgorithmVersionHandle(ctx context.Context, req *api.AddMyAlgorithmVersionRequest) (*api.AddMyAlgorithmVersionReply, error) {

	algorithmDao := h.data.AlgorithmDao

	myAlgorithm, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
		AlgorithmId: req.AlgorithmId,
	})

	if err != nil {
		return nil, err
	}

	if (time.Now().Unix() - myAlgorithm.UpdatedAt.Unix()) < CREATE_MIN_TIME_INTERVAL {
		err := errors.Errorf(nil, errors.ErrorAlgorithmVersionRepeat)
		h.log.Errorw(ctx, err)
		return nil, err
	}

	latestVersion, err := common.VersionStrParse(myAlgorithm.LatestVersion)
	if err != nil {
		return nil, err
	}

	oriAlgorithmVersion, err := algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{
		AlgorithmId: req.AlgorithmId,
		Version:     req.OriVersion,
	})

	if err != nil {
		return nil, err
	}

	if oriAlgorithmVersion.FileStatus != FILESTATUS_FINISH {
		err := errors.Errorf(nil, errors.ErrorAlgorithmVersionFileNotReady)
		return nil, err
	}

	myAlgorithmVersion, err := algorithmDao.AddAlgorithmVersion(ctx, &model.AlgorithmVersion{
		Id:                utils.GetUUIDWithoutSeparator(),
		AlgorithmId:       req.AlgorithmId,
		Version:           common.VersionStrBuild(latestVersion + 1),
		AlgorithmDescript: req.AlgorithmDescript,
		FileStatus:        FILESTATUS_UPLOGADING,
	})
	if err != nil {
		return nil, err
	}

	myAlgorithm.LatestVersion = myAlgorithmVersion.Version
	err = algorithmDao.UpdateAlgorithm(ctx, myAlgorithm)
	if err != nil {
		return nil, err
	}
	fromBucektName := ""
	fromObjectName := ""
	toBucektName := ""
	toObjectName := ""

	if myAlgorithm.IsPrefab {
		fromBucektName = common.GetMinioBucket()
		fromObjectName = common.GetMinioPreCodeObject(myAlgorithm.AlgorithmId, req.OriVersion)
		toBucektName = common.GetMinioBucket()
		toObjectName = common.GetMinioPreCodeObject(myAlgorithm.AlgorithmId, myAlgorithmVersion.Version)
	} else {
		fromBucektName = common.GetMinioBucket()
		fromObjectName = common.GetMinioCodeObject(myAlgorithm.SpaceId, myAlgorithm.UserId, myAlgorithm.AlgorithmId, req.OriVersion)
		toBucektName = common.GetMinioBucket()
		toObjectName = common.GetMinioCodeObject(myAlgorithm.SpaceId, myAlgorithm.UserId, myAlgorithm.AlgorithmId, myAlgorithmVersion.Version)
	}
	fromPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, fromBucektName, fromObjectName)
	toPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, toBucektName, toObjectName)

	// 拷贝
	go func() {
		h.log.Infof(ctx, "copy dir when AddMyAlgorithmVersion, from [%s] to [%s]", fromPath, toPath)
		err := utils.CopyDir(fromPath, toPath)
		if err != nil {
			myAlgorithmVersion.FileStatus = FILESTATUS_FAILED
			h.log.Errorw(ctx, err)
		} else {
			myAlgorithmVersion.FileStatus = FILESTATUS_FINISH
		}

		err = algorithmDao.UpdateAlgorithmVersion(ctx, myAlgorithmVersion)
		if err != nil {
			h.log.Errorw(ctx, err)
			return
		}
	}()

	return &api.AddMyAlgorithmVersionReply{
		AlgorithmId: myAlgorithmVersion.AlgorithmId,
		Version:     myAlgorithmVersion.Version,
		CreatedAt:   myAlgorithmVersion.CreatedAt.Unix(),
	}, nil
}

// 删除我的算法版本
func (h *algorithmHandle) DeleteMyAlgorithmVersionHandle(ctx context.Context, req *api.DeleteMyAlgorithmVersionRequest) (*api.DeleteMyAlgorithmVersionReply, error) {
	algorithmDao := h.data.AlgorithmDao
	algorithmId := req.AlgorithmId
	version := req.Version

	// 查算法信息
	algorithmInt, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
		AlgorithmId: algorithmId,
	})
	if err != nil {
		return nil, err
	}
	if algorithmInt.SpaceId != req.SpaceId || algorithmInt.UserId != req.UserId {
		err := errors.Errorf(nil, errors.ErrorAlgorithmNotMy)
		h.log.Errorw(ctx, err)
		return nil, err
	}

	_, err = algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{
		AlgorithmId: algorithmId,
		Version:     version,
	})

	if err != nil {
		return nil, err
	}

	// 取消分享
	_, err = h.AllCloseShareAlgorithmVersionHandle(ctx, &api.AllCloseShareAlgorithmVersionRequest{
		AlgorithmId: algorithmId,
		Version:     version,
	})
	if err != nil {
		return nil, err
	}

	// 删除版本信息
	err = algorithmDao.DeleteAlgorithmVersion(ctx, &model.AlgorithmVersionDelete{
		AlgorithmId:      algorithmId,
		AlgorithmVersion: version,
	})
	if err != nil {
		return nil, err
	}

	maxVersion, err := h.findAlgorithmVersionMaxId(ctx, algorithmId)

	if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
		return nil, err
	}
	if err != nil && errors.IsError(errors.ErrorDBFindEmpty, err) {
		// 最后一个版本，那就都删了
		err := algorithmDao.DeleteAlgorithm(ctx, &model.AlgorithmDelete{
			AlgorithmId: algorithmId,
		})
		if err != nil {
			return nil, err
		}
		// 减少算法类型引用
		_, _ = h.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: algorithmInt.ApplyId})
		// 减少算法框架引用
		_, _ = h.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: algorithmInt.FrameworkId})
	} else {
		algorithmInt.LatestVersion = maxVersion
		err = algorithmDao.UpdateAlgorithm(ctx, algorithmInt)
		if err != nil {
			return nil, err
		}
	}
	// 删除算法版本Minio存储
	go utils.HandlePanicBG(func(i ...interface{}) {
		bucketName := common.GetMinioBucket()
		objectName := common.GetMinioCodeObject(req.SpaceId, req.UserId, algorithmId, req.Version)
		h.data.Redis.SAddMinioRemovingObject(bucketName + "-" + objectName)
		defer h.data.Redis.SRemMinioRemovingObject(bucketName + "-" + objectName)
		h.data.Minio.RemoveObject(bucketName, objectName)
	})()
	go utils.HandlePanicBG(func(i ...interface{}) {
		bucketName := common.GetMinioBucket()
		objectName := common.GetMinioDownloadCodeVersionObject(algorithmId, req.Version)
		h.data.Redis.SAddMinioRemovingObject(bucketName + "-" + objectName)
		defer h.data.Redis.SRemMinioRemovingObject(bucketName + "-" + objectName)
		h.data.Minio.RemoveObject(bucketName, objectName)
	})()

	return &api.DeleteMyAlgorithmVersionReply{
		DeletedAt: time.Now().Unix(),
	}, nil
}

// 删除我的算法
func (h *algorithmHandle) DeleteMyAlgorithmHandle(ctx context.Context, req *api.DeleteMyAlgorithmRequest) (*api.DeleteMyAlgorithmReply, error) {
	algorithmDao := h.data.AlgorithmDao
	algorithmId := req.AlgorithmId

	// 查算法信息
	algorithmInt, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
		AlgorithmId: algorithmId,
	})
	if err != nil {
		return nil, err
	}
	if algorithmInt.SpaceId != req.SpaceId || algorithmInt.UserId != req.UserId {
		err := errors.Errorf(nil, errors.ErrorAlgorithmNotMy)
		h.log.Errorw(ctx, err)
		return nil, err
	}

	// 取消分享
	_, err = h.AllCloseShareAlgorithmHandle(ctx, &api.AllCloseShareAlgorithmRequest{
		AlgorithmId: algorithmId,
	})
	if err != nil {
		return nil, err
	}

	// 删除算法版本信息
	err = algorithmDao.BatchDeleteAlgorithmVersion(ctx, &model.AlgorithmVersionBatchDelete{
		AlgorithmId: algorithmId,
	})
	if err != nil {
		return nil, err
	}

	// 删除算法信息
	err = algorithmDao.DeleteAlgorithm(ctx, &model.AlgorithmDelete{
		AlgorithmId: algorithmId,
	})
	if err != nil {
		return nil, err
	}

	// 减少算法类型引用
	_, _ = h.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: algorithmInt.ApplyId})
	// 减少算法框架引用
	_, _ = h.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: algorithmInt.FrameworkId})

	// 删除算法版本Minio存储
	go utils.HandlePanicBG(func(i ...interface{}) {
		bucketName := common.GetMinioBucket()
		objectName := common.GetMinioCodePathObject(req.SpaceId, req.UserId, algorithmId)
		h.data.Redis.SAddMinioRemovingObject(bucketName + "-" + objectName)
		defer h.data.Redis.SRemMinioRemovingObject(bucketName + "-" + objectName)
		h.data.Minio.RemoveObject(bucketName, objectName)
	})()
	go utils.HandlePanicBG(func(i ...interface{}) {
		bucketName := common.GetMinioBucket()
		objectName := common.GetMinioDownloadCodePathObject(algorithmId)
		h.data.Redis.SAddMinioRemovingObject(bucketName + "-" + objectName)
		defer h.data.Redis.SRemMinioRemovingObject(bucketName + "-" + objectName)
		h.data.Minio.RemoveObject(bucketName, objectName)
	})()

	return &api.DeleteMyAlgorithmReply{
		DeletedAt: time.Now().Unix(),
	}, nil
}

// 新增预置算法版本
func (h *algorithmHandle) AddPreAlgorithmVersionHandle(ctx context.Context, req *api.AddPreAlgorithmVersionRequest) (*api.AddPreAlgorithmVersionReply, error) {

	algorithmDao := h.data.AlgorithmDao

	preAlgorithm, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
		AlgorithmId: req.AlgorithmId,
	})

	if err != nil {
		return nil, err
	}

	if (time.Now().Unix() - preAlgorithm.UpdatedAt.Unix()) < CREATE_MIN_TIME_INTERVAL {
		err := errors.Errorf(nil, errors.ErrorAlgorithmVersionRepeat)
		h.log.Errorw(ctx, err)
		return nil, err
	}

	latestVersion, err := common.VersionStrParse(preAlgorithm.LatestVersion)
	if err != nil {
		return nil, err
	}

	preAlgorithmVersion, err := algorithmDao.AddAlgorithmVersion(ctx, &model.AlgorithmVersion{
		Id:                utils.GetUUIDWithoutSeparator(),
		AlgorithmId:       req.AlgorithmId,
		Version:           common.VersionStrBuild(latestVersion + 1),
		AlgorithmDescript: req.AlgorithmDescript,
		FileStatus:        FILESTATUS_INIT,
	})
	if err != nil {
		return nil, err
	}

	preAlgorithm.LatestVersion = preAlgorithmVersion.Version
	err = algorithmDao.UpdateAlgorithm(ctx, preAlgorithm)
	if err != nil {
		return nil, err
	}

	return &api.AddPreAlgorithmVersionReply{
		AlgorithmId: preAlgorithmVersion.AlgorithmId,
		Version:     preAlgorithmVersion.Version,
		CreatedAt:   preAlgorithmVersion.CreatedAt.Unix(),
	}, nil
}

// 删除预置算法版本
func (h *algorithmHandle) DeletePreAlgorithmVersionHandle(ctx context.Context, req *api.DeletePreAlgorithmVersionRequest) (*api.DeletePreAlgorithmVersionReply, error) {
	algorithmDao := h.data.AlgorithmDao
	algorithmId := req.AlgorithmId
	version := req.Version

	preAlgorithm, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
		AlgorithmId: req.AlgorithmId,
	})
	if err != nil {
		return nil, err
	}
	if !preAlgorithm.IsPrefab {
		err := errors.Errorf(nil, errors.ErrorAlgorithmNotMy)
		h.log.Errorw(ctx, err)
		return nil, err
	}

	_, err = algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{
		AlgorithmId: algorithmId,
		Version:     version,
	})

	if err != nil {
		return nil, err
	}

	// 删除版本信息
	err = algorithmDao.DeleteAlgorithmVersion(ctx, &model.AlgorithmVersionDelete{
		AlgorithmId:      algorithmId,
		AlgorithmVersion: version,
	})
	if err != nil {
		return nil, err
	}

	maxVersion, err := h.findAlgorithmVersionMaxId(ctx, algorithmId)
	if err != nil && !errors.IsError(errors.ErrorDBFindEmpty, err) {
		return nil, err
	}
	if err != nil && errors.IsError(errors.ErrorDBFindEmpty, err) {
		err := algorithmDao.DeleteAlgorithm(ctx, &model.AlgorithmDelete{
			AlgorithmId: algorithmId,
		})
		if err != nil {
			return nil, err
		}

		// 减少算法类型引用
		_, _ = h.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: preAlgorithm.ApplyId})
		// 减少算法框架引用
		_, _ = h.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: preAlgorithm.FrameworkId})

	} else {
		// 查算法信息
		algorithmInt, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
			AlgorithmId: algorithmId,
		})
		if err != nil {
			return nil, err
		}

		algorithmInt.LatestVersion = maxVersion
		err = algorithmDao.UpdateAlgorithm(ctx, algorithmInt)
		if err != nil {
			return nil, err
		}
	}

	// 删除预置算法版本Minio存储
	go utils.HandlePanicBG(func(i ...interface{}) {
		bucketName := common.GetMinioBucket()
		objectName := common.GetMinioPreCodeObject(algorithmId, version)
		h.data.Redis.SAddMinioRemovingObject(bucketName + "-" + objectName)
		defer h.data.Redis.SRemMinioRemovingObject(bucketName + "-" + objectName)
		h.data.Minio.RemoveObject(bucketName, objectName)
	})()
	go utils.HandlePanicBG(func(i ...interface{}) {
		bucketName := common.GetMinioBucket()
		objectName := common.GetMinioDownloadCodeVersionObject(algorithmId, version)
		h.data.Redis.SAddMinioRemovingObject(bucketName + "-" + objectName)
		defer h.data.Redis.SRemMinioRemovingObject(bucketName + "-" + objectName)
		h.data.Minio.RemoveObject(bucketName, objectName)
	})()

	return &api.DeletePreAlgorithmVersionReply{
		DeletedAt: time.Now().Unix(),
	}, nil
}

// 删除预置算法
func (h *algorithmHandle) DeletePreAlgorithmHandle(ctx context.Context, req *api.DeletePreAlgorithmRequest) (*api.DeletePreAlgorithmReply, error) {
	algorithmDao := h.data.AlgorithmDao
	algorithmId := req.AlgorithmId

	preAlgorithm, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
		AlgorithmId: req.AlgorithmId,
	})
	if err != nil {
		return nil, err
	}
	if !preAlgorithm.IsPrefab {
		err := errors.Errorf(nil, errors.ErrorAlgorithmNotMy)
		h.log.Errorw(ctx, err)
		return nil, err
	}

	// 删除算法信息
	err = algorithmDao.DeleteAlgorithm(ctx, &model.AlgorithmDelete{
		AlgorithmId: algorithmId,
	})
	if err != nil {
		return nil, err
	}

	// 减少算法类型引用
	_, _ = h.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: preAlgorithm.ApplyId})
	// 减少算法框架引用
	_, _ = h.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: preAlgorithm.FrameworkId})

	// 删除预置算法Minio存储
	go utils.HandlePanicBG(func(i ...interface{}) {
		bucketName := common.GetMinioBucket()
		objectName := common.GetMinioPreCodePathObject(algorithmId)
		h.data.Redis.SAddMinioRemovingObject(bucketName + "-" + objectName)
		defer h.data.Redis.SRemMinioRemovingObject(bucketName + "-" + objectName)
		h.data.Minio.RemoveObject(bucketName, objectName)
	})()
	go utils.HandlePanicBG(func(i ...interface{}) {
		bucketName := common.GetMinioBucket()
		objectName := common.GetMinioDownloadCodePathObject(algorithmId)
		h.data.Redis.SAddMinioRemovingObject(bucketName + "-" + objectName)
		defer h.data.Redis.SRemMinioRemovingObject(bucketName + "-" + objectName)
		h.data.Minio.RemoveObject(bucketName, objectName)
	})()

	return &api.DeletePreAlgorithmReply{
		DeletedAt: time.Now().Unix(),
	}, nil
}

// 修改算法版本
func (h *algorithmHandle) UpdateAlgorithmVersionHandle(ctx context.Context, req *api.UpdateAlgorithmVersionRequest) (*api.UpdateAlgorithmVersionReply, error) {
	algorithmDao := h.data.AlgorithmDao
	algorithm, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{AlgorithmId: req.AlgorithmId})
	if err != nil {
		return nil, err
	}
	if algorithm.SpaceId != req.SpaceId || algorithm.UserId != algorithm.UserId || algorithm.IsPrefab != req.IsPrefab {
		return nil, errors.Errorf(nil, errors.ErrorFindAlgorithmAuthWrong)
	}

	algorithmVersion, err := algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{AlgorithmId: req.AlgorithmId, Version: req.Version})
	if err != nil {
		return nil, err
	}

	algorithmVersion.AlgorithmDescript = req.AlgorithmDescript
	err = algorithmDao.UpdateAlgorithmVersion(ctx, algorithmVersion)
	if err != nil {
		return nil, err
	}

	return &api.UpdateAlgorithmVersionReply{UpdatedAt: time.Now().Unix()}, nil
}

// 压缩算法版本包
func (h *algorithmHandle) DownloadAlgorithmVersionCompressHandle(ctx context.Context,
	req *api.DownloadAlgorithmVersionCompressRequest) (*api.DownloadAlgorithmVersionCompressReply, error) {
	algorithmDao := h.data.AlgorithmDao
	myAlgorithm, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
		AlgorithmId: req.AlgorithmId,
	})
	if err != nil {
		return nil, err
	}
	mv, err := algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{
		AlgorithmId: req.AlgorithmId,
		Version:     req.Version,
	})
	if err != nil {
		return nil, err
	}

	compressAt := time.Now().Unix()

	fromBucektName := ""
	fromObjectName := ""
	toBucektName := ""
	toObjectName := ""

	if myAlgorithm.IsPrefab {
		fromBucektName = common.GetMinioBucket()
		fromObjectName = common.GetMinioPreCodeObject(myAlgorithm.AlgorithmId, req.Version)
	} else {
		fromBucektName = common.GetMinioBucket()
		fromObjectName = common.GetMinioCodeObject(myAlgorithm.SpaceId, myAlgorithm.UserId, myAlgorithm.AlgorithmId, req.Version)
	}
	toBucektName = common.GetMinioBucket()
	toObjectName = common.GetMinioDownloadCodeObject(myAlgorithm.AlgorithmId, mv.Version, fmt.Sprintf("%d/%s-%s.zip", compressAt, myAlgorithm.AlgorithmName, mv.Version))

	fromPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, fromBucektName, fromObjectName)
	toPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, toBucektName, toObjectName)

	go func() {
		err := utils.Zip(fromPath, toPath)
		if err != nil {
			h.log.Errorw(ctx, err)
			return
		}
		mv.LatestCompressed = compressAt
		err = algorithmDao.UpdateAlgorithmVersion(ctx, mv)
		if err != nil {
			h.log.Errorw(ctx, err)
		}
	}()
	return &api.DownloadAlgorithmVersionCompressReply{
		CompressAt: compressAt,
	}, nil
}

// 下载算法版本
func (h *algorithmHandle) DownloadAlgorithmVersionHandle(ctx context.Context, req *api.DownloadAlgorithmVersionRequest) (*api.DownloadAlgorithmVersionReply, error) {
	algorithmDao := h.data.AlgorithmDao
	algorithmId := req.AlgorithmId
	version := req.Version

	algorithmInt, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
		AlgorithmId: algorithmId,
	})
	if err != nil {
		return nil, err
	}

	algorithmVersion, err := algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{
		AlgorithmId: algorithmId,
		Version:     version,
	})
	if err != nil {
		return nil, err
	}

	bucektName := ""
	objectName := ""

	bucektName = common.GetMinioBucket()
	objectName = common.GetMinioDownloadCodeObject(
		algorithmInt.AlgorithmId,
		algorithmVersion.Version, fmt.Sprintf("%d/%s-%s.zip", req.CompressAt, algorithmInt.AlgorithmName, algorithmVersion.Version))

	url, err := h.data.Minio.PresignedDownloadObject(bucektName, objectName, req.Domain)

	if err != nil {
		return nil, err
	}

	return &api.DownloadAlgorithmVersionReply{
		DownloadUrl: url.String(),
	}, nil
}

func (h *algorithmHandle) findAlgorithmVersionAccessMaxId(ctx context.Context, algorithmAccessId string) (string, error) {
	algorithmDao := h.data.AlgorithmDao

	_, algorithmVersionAccessList, err := algorithmDao.ListAlgorithmAccessVersion(ctx, &model.AlgorithmAccessVersionList{
		AlgorithmAccessId:     algorithmAccessId,
		AlgorithmVersionOrder: true,
		AlgorithmVersionSort:  model.DESC,
		PageIndex:             1,
		PageSize:              1,
	})
	if err != nil {
		return "", err
	}
	if len(algorithmVersionAccessList) == 0 {
		err = errors.Errorf(nil, errors.ErrorDBFindEmpty)
		return "", err
	}

	return algorithmVersionAccessList[0].AlgorithmVersion, nil
}

func (h *algorithmHandle) findAlgorithmVersionMaxId(ctx context.Context, algorithmId string) (string, error) {
	algorithmDao := h.data.AlgorithmDao

	_, algorithmVersionList, err := algorithmDao.ListAlgorithmVersion(ctx, &model.AlgorithmVersionList{
		AlgorithmId:  algorithmId,
		VersionOrder: true,
		VersionSort:  model.DESC,
		PageIndex:    1,
		PageSize:     1,
	})
	if err != nil {
		return "", err
	}
	if len(algorithmVersionList) == 0 {
		err = errors.Errorf(nil, errors.ErrorDBFindEmpty)
		return "", err
	} else if len(algorithmVersionList) != 1 {
		err = errors.Errorf(nil, errors.ErrorFindAlgorithmVersionMaxIdFailed)
		return "", err
	}
	return algorithmVersionList[0].Version, nil
}

// 复制算法版本
func (h *algorithmHandle) CopyAlgorithmVersionHandle(ctx context.Context, req *api.CopyAlgorithmVersionRequest) (*api.CopyAlgorithmVersionReply, error) {

	algorithmDao := h.data.AlgorithmDao

	algorithm, _ := algorithmDao.QueryAlgorithmByInfo(ctx, &model.AlgorithmQueryByInfo{
		AlgorithmName: req.NewAlgorithmName,
		UserId:        req.UserId,
		SpaceId:       req.SpaceId,
		IsPrefab:      false,
	})

	if algorithm != nil {
		err := errors.Errorf(nil, errors.ErrorAlgorithmRepeat)
		return nil, err
	}

	oriAlgorithm, err := algorithmDao.QueryAlgorithm(ctx, &model.AlgorithmQuery{
		AlgorithmId: req.AlgorithmId,
	})
	if err != nil {
		return nil, err
	}

	_, err = algorithmDao.QueryAlgorithmVersion(ctx, &model.AlgorithmVersionQuery{
		AlgorithmId: req.AlgorithmId,
		Version:     req.Version,
	})

	if err != nil {
		return nil, err
	}

	algorithmId := utils.GetUUIDWithoutSeparator()
	algorithmVersionId := utils.GetUUIDWithoutSeparator()
	algorithmVersion := common.VersionStrBuild(1)
	myAlgorithm, err := algorithmDao.AddAlgorithm(ctx, &model.Algorithm{
		AlgorithmId:       algorithmId,
		SpaceId:           req.SpaceId,
		UserId:            req.UserId,
		IsPrefab:          false,
		AlgorithmName:     req.NewAlgorithmName,
		AlgorithmDescript: req.AlgorithmDescript,
		ModelName:         req.ModelName,
		LatestVersion:     algorithmVersion,
		ApplyId:           oriAlgorithm.ApplyId,
		FrameworkId:       oriAlgorithm.FrameworkId,
		AlgorithmVersions: []*model.AlgorithmVersion{
			{
				Id:                algorithmVersionId,
				Version:           algorithmVersion,
				AlgorithmDescript: req.AlgorithmDescript,
				FileStatus:        FILESTATUS_UPLOGADING,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// 检查数据类型id
	if oriAlgorithm.ApplyId != "" {
		algorithmApply, err := h.lableService.GetLable(ctx, &api.GetLableRequest{Id: oriAlgorithm.ApplyId})
		if err != nil {
			return nil, err
		}
		// 新增算法类型引用
		_, _ = h.lableService.IncreaseLableReferTimes(ctx, &api.IncreaseLableReferTimesRequest{Id: algorithmApply.Lable.Id})
	}
	// 检查框架id
	if oriAlgorithm.FrameworkId != "" {
		algorithmFramework, err := h.lableService.GetLable(ctx, &api.GetLableRequest{Id: oriAlgorithm.FrameworkId})
		if err != nil {
			return nil, err
		}
		// 新增算法框架引用
		_, _ = h.lableService.IncreaseLableReferTimes(ctx, &api.IncreaseLableReferTimesRequest{Id: algorithmFramework.Lable.Id})
	}

	fromBucektName := ""
	fromObjectName := ""
	toBucektName := ""
	toObjectName := ""

	if oriAlgorithm.IsPrefab {
		fromBucektName = common.GetMinioBucket()
		fromObjectName = common.GetMinioPreCodeObject(oriAlgorithm.AlgorithmId, req.Version)
	} else {
		fromBucektName = common.GetMinioBucket()
		fromObjectName = common.GetMinioCodeObject(oriAlgorithm.SpaceId, oriAlgorithm.UserId, oriAlgorithm.AlgorithmId, req.Version)
	}

	toBucektName = common.GetMinioBucket()
	toObjectName = common.GetMinioCodeObject(myAlgorithm.SpaceId, myAlgorithm.UserId, myAlgorithm.AlgorithmId, algorithmVersion)

	fromPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, fromBucektName, fromObjectName)
	toPath := fmt.Sprintf("%s/%s/%s", h.conf.Data.Minio.Base.MountPath, toBucektName, toObjectName)

	// 拷贝
	go func() {
		h.log.Infof(ctx, "copy dir when CopyAlgorithmVersion, from [%s] to [%s]", fromPath, toPath)
		err := utils.CopyDir(fromPath, toPath)
		if err != nil {
			myAlgorithm.AlgorithmVersions[0].FileStatus = FILESTATUS_FAILED
			h.log.Errorw(ctx, err)
		} else {
			myAlgorithm.AlgorithmVersions[0].FileStatus = FILESTATUS_FINISH
		}

		err = algorithmDao.UpdateAlgorithmVersion(ctx, myAlgorithm.AlgorithmVersions[0])
		if err != nil {
			h.log.Errorw(ctx, err)
			return
		}
	}()

	return &api.CopyAlgorithmVersionReply{
		NewAlgorithmId: algorithmId,
		NewVersion:     algorithmVersion,
		CreatedAt:      myAlgorithm.AlgorithmVersions[0].CreatedAt.Unix(),
	}, nil
}
