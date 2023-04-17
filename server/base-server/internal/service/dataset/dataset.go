package dataset

import (
	"context"
	"fmt"
	"path/filepath"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"server/common/constant"
	commctx "server/common/context"
	"server/common/errors"
	"server/common/graceful"
	"server/common/log"
	"server/common/utils"
	"time"

	fluidv1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
	Common "github.com/fluid-cloudnative/fluid/pkg/common"
	"github.com/jinzhu/copier"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	statusForUpload = []int{int(api.DatasetVersionStatus_DVS_Init), int(api.DatasetVersionStatus_DVS_UnzipFailed)}
)

type datasetService struct {
	api.UnimplementedDatasetServiceServer
	conf         *conf.Bootstrap
	log          *log.Helper
	data         *data.Data
	lableService api.LableServiceServer
	tieredStore  fluidv1.TieredStore
}

func NewDatasetService(conf *conf.Bootstrap, logger log.Logger, data *data.Data, lableService api.LableServiceServer) api.DatasetServiceServer {
	log := log.NewHelper("DatasetService", logger)

	s := &datasetService{
		conf:         conf,
		log:          log,
		data:         data,
		lableService: lableService,
	}

	return s
}

func (s *datasetService) CreateDataset(ctx context.Context, req *api.CreateDatasetRequest) (*api.CreateDatasetReply, error) {
	v := common.VersionStrBuild(1)

	datasetId := utils.GetUUIDWithoutSeparator()
	dataset := &model.Dataset{}
	err := copier.CopyWithOption(dataset, req, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	dataset.Id = datasetId

	_, size, err := s.data.DatasetDao.ListDataset(ctx, &model.DatasetQuery{
		UserId:     req.UserId,
		SpaceId:    req.SpaceId,
		Name:       req.Name,
		SourceType: int(req.SourceType),
	})
	if err != nil {
		return nil, err
	}
	if size > 0 {
		return nil, errors.Errorf(nil, errors.ErrorDatasetRepeat)
	}

	err = s.data.DatasetDao.CreateDataset(ctx, dataset)
	if err != nil {
		return nil, err
	}

	toPath := s.getPath(dataset, v)
	version := &model.DatasetVersion{
		DatasetId:  datasetId,
		Version:    v,
		VersionInt: 1,
		Desc:       req.Desc,
		Status:     int(api.DatasetVersionStatus_DVS_Init),
		Path:       toPath,
	}
	err = s.data.DatasetDao.CreateDatasetVersion(ctx, version)
	if err != nil {
		return nil, err
	}

	// 检查数据类型id
	if req.TypeId != "" {
		datasetType, err := s.lableService.GetLable(ctx, &api.GetLableRequest{Id: req.TypeId})
		if err != nil {
			return nil, err
		}
		// 新增数据类型引用
		_, _ = s.lableService.IncreaseLableReferTimes(ctx, &api.IncreaseLableReferTimesRequest{Id: datasetType.Lable.Id})
	}
	// 检查数据用途id
	for _, id := range req.ApplyIds {
		datasetApply, err := s.lableService.GetLable(ctx, &api.GetLableRequest{Id: id})
		if err != nil {
			return nil, err
		}
		// 新增数据用途引用
		_, _ = s.lableService.IncreaseLableReferTimes(ctx, &api.IncreaseLableReferTimesRequest{Id: datasetApply.Lable.Id})
	}

	return &api.CreateDatasetReply{
		Id:      datasetId,
		Version: version.Version,
	}, nil
}

func (s *datasetService) ListDataset(ctx context.Context, req *api.ListDatasetRequest) (*api.ListDatasetReply, error) {
	query := &model.DatasetQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	datasetsTbl, totalSize, err := s.data.DatasetDao.ListDataset(ctx, query)
	if err != nil {
		return nil, err
	}

	ids := make([]string, 0)
	for _, d := range datasetsTbl {
		ids = append(ids, d.Id)
	}
	idsV, err := s.data.DatasetDao.ListDatasetVersionLatestVersion(ctx, ids)

	if err != nil {
		return nil, err
	}

	datasets := make([]*api.Dataset, 0)
	for _, n := range datasetsTbl {
		dataset := &api.Dataset{}
		err := copier.Copy(dataset, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		dataset.CreatedAt = n.CreatedAt.Unix()
		dataset.UpdatedAt = n.UpdatedAt.Unix()
		dataset.LatestVersion = common.VersionStrBuild(idsV[n.Id])
		datasetType, err := s.lableService.GetLable(ctx, &api.GetLableRequest{Id: n.TypeId})
		if err != nil {
			dataset.TypeDesc = ""
		} else {
			dataset.TypeDesc = datasetType.Lable.LableDesc
		}

		if len(n.ApplyIds) > 0 {
			datasetApply, err := s.lableService.ListLable(ctx, &api.ListLableRequest{PageIndex: 1, PageSize: int64(len(n.ApplyIds)), Ids: n.ApplyIds})

			if err != nil {
				return nil, err
			}
			for _, a := range datasetApply.Lables {
				dataset.Applies = append(dataset.Applies, &api.Dataset_Apply{
					Id:   a.Id,
					Desc: a.LableDesc,
				})
			}
		}

		datasets = append(datasets, dataset)
	}
	return &api.ListDatasetReply{
		TotalSize: totalSize,
		Datasets:  datasets,
	}, nil
}

func (s *datasetService) ListCommDataset(ctx context.Context, req *api.ListCommDatasetRequest) (*api.ListCommDatasetReply, error) {
	query := &model.CommDatasetQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	datasetsTbl, totalSize, err := s.data.DatasetDao.ListCommDataset(ctx, query)
	if err != nil {
		return nil, err
	}

	shareIds := make([]model.DatasetAccessId, 0)
	shareIdsV := map[model.DatasetAccessId]int64{}
	for _, d := range datasetsTbl {
		shareIds = append(shareIds, model.DatasetAccessId{
			DatasetId: d.Id,
			SpaceId:   req.ShareSpaceId,
		})
	}

	if len(shareIds) > 0 {
		shareIdsV, err = s.data.DatasetDao.ListDatasetVersionAccessLatestVersion(ctx, shareIds)
		if err != nil {
			return nil, err
		}
	}

	datasets := make([]*api.Dataset, 0)
	for _, n := range datasetsTbl {
		dataset := &api.Dataset{}
		err := copier.Copy(dataset, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		dataset.CreatedAt = n.CreatedAt.Unix()
		dataset.UpdatedAt = n.UpdatedAt.Unix()
		if req.ShareSpaceId != "" {
			dataset.LatestVersion = common.VersionStrBuild(shareIdsV[model.DatasetAccessId{
				DatasetId: n.Id,
				SpaceId:   req.ShareSpaceId,
			}])
		}

		datasetType, err := s.lableService.GetLable(ctx, &api.GetLableRequest{Id: n.TypeId})
		if err != nil {
			dataset.TypeDesc = ""
		} else {
			dataset.TypeDesc = datasetType.Lable.LableDesc
		}

		if len(n.ApplyIds) > 0 {
			datasetApply, err := s.lableService.ListLable(ctx, &api.ListLableRequest{PageIndex: 1, PageSize: int64(len(n.ApplyIds)), Ids: n.ApplyIds})
			if err != nil {
				return nil, err
			}
			for _, a := range datasetApply.Lables {
				dataset.Applies = append(dataset.Applies, &api.Dataset_Apply{
					Id:   a.Id,
					Desc: a.LableDesc,
				})
			}
		}

		datasets = append(datasets, dataset)
	}

	return &api.ListCommDatasetReply{
		TotalSize: totalSize,
		Datasets:  datasets,
	}, nil
}

func (s *datasetService) GetDataset(ctx context.Context, req *api.GetDatasetRequest) (*api.GetDatasetReply, error) {
	listDatasetReply, err := s.ListDataset(ctx, &api.ListDatasetRequest{Ids: []string{req.Id}})
	if err != nil {
		return nil, err
	}

	if len(listDatasetReply.Datasets) == 0 {
		return nil, errors.Errorf(nil, errors.ErrorDBFindEmpty)
	}

	return &api.GetDatasetReply{
		Dataset: listDatasetReply.Datasets[0],
	}, nil
}

func (s *datasetService) CreateDatasetVersion(ctx context.Context, req *api.CreateDatasetVersionRequest) (*api.CreateDatasetVersionReply, error) {
	dataset, err := s.data.DatasetDao.GetDataset(ctx, req.DatasetId)
	if err != nil {
		return nil, err
	}

	vInt, err := s.data.DatasetDao.ListDatasetVersionLatestVersion(ctx, []string{req.DatasetId})
	if err != nil {
		return nil, err
	}
	newVInt := vInt[req.DatasetId] + 1
	newV := common.VersionStrBuild(newVInt)

	toPath := s.getPath(dataset, newV)

	version := &model.DatasetVersion{
		DatasetId:  dataset.Id,
		Version:    newV,
		VersionInt: newVInt,
		Desc:       req.Desc,
		Status:     int(api.DatasetVersionStatus_DVS_Init),
		Path:       toPath,
	}
	err = s.data.DatasetDao.CreateDatasetVersion(ctx, version)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorDBCreateFailed)
	}

	return &api.CreateDatasetVersionReply{
		DatasetId: req.DatasetId,
		Version:   newV,
	}, nil
}

func (s *datasetService) ConfirmUploadDatasetVersion(ctx context.Context, req *api.ConfirmUploadDatasetVersionRequest) (*api.ConfirmUploadDatasetVersionReply, error) {
	dataset, err := s.data.DatasetDao.GetDataset(ctx, req.DatasetId)

	if err != nil {
		return nil, err
	}

	version, err := s.data.DatasetDao.GetDatasetVersion(ctx, req.DatasetId, req.Version)

	if err != nil {
		return nil, err
	}

	if !utils.IntInSlice(version.Status, statusForUpload) {
		return nil, errors.Errorf(nil, errors.ErrorDatasetStatusForbidden)
	}

	fromBucket, fromObject := getTempMinioPath(dataset, req.Version, req.FileName)
	isExist, err := s.data.Minio.ObjectExist(fromBucket, fromObject)

	if err != nil {
		return nil, err
	}
	if !isExist {
		err := errors.Errorf(nil, errors.ErrorDatasetFileNotFound)
		return nil, err
	}

	toBucket, toObject := getMinioPath(dataset, req.Version)
	fromPath := fmt.Sprintf("%s/%s/%s", s.conf.Data.Minio.Base.MountPath, fromBucket, fromObject)
	toPath := fmt.Sprintf("%s/%s/%s", s.conf.Data.Minio.Base.MountPath, toBucket, toObject)
	graceful.AddOne()
	go utils.HandlePanic(ctx, func(i ...interface{}) {
		defer graceful.Done()
		// 删除数据集压缩包临时文件
		defer func() {
			go utils.HandlePanicBG(func(i ...interface{}) {
				s.data.Redis.SAddMinioRemovingObject(fromBucket + "-" + fromObject)
				defer s.data.Redis.SRemMinioRemovingObject(fromBucket + "-" + fromObject)
				s.data.Minio.RemoveObject(fromBucket, fromObject)
			})()
		}()
		ctx := i[0].(context.Context)
		err := s.data.DatasetDao.UpdateDatasetVersionSelective(ctx, &model.DatasetVersion{
			DatasetId: req.DatasetId,
			Version:   req.Version,
			Status:    int(api.DatasetVersionStatus_DVS_Unzipping),
		})
		if err != nil {
			s.log.Errorw(ctx, err)
		}

		err = utils.Unzip(fromPath, toPath)
		if err != nil {
			s.log.Errorw(ctx, err)

			err := s.data.DatasetDao.UpdateDatasetVersionSelective(ctx, &model.DatasetVersion{
				DatasetId: req.DatasetId,
				Version:   req.Version,
				Status:    int(api.DatasetVersionStatus_DVS_UnzipFailed),
			})
			if err != nil {
				s.log.Errorw(ctx, err)
			}
			return
		}
		err = s.data.DatasetDao.UpdateDatasetVersionSelective(ctx, &model.DatasetVersion{
			DatasetId:    req.DatasetId,
			Version:      req.Version,
			Status:       int(api.DatasetVersionStatus_DVS_Unzipped),
			OriginalPath: fmt.Sprintf("%s/%s", fromBucket, fromObject),
		})
		if err != nil {
			s.log.Errorw(ctx, err)
		}
	})(commctx.WithoutCancel(ctx)) // http请求结束后ctx会被cancel 这里创建一个不会取消的ctx并传值

	return &api.ConfirmUploadDatasetVersionReply{UpdatedAt: time.Now().Unix()}, nil
}

func (s *datasetService) ListDatasetVersion(ctx context.Context, req *api.ListDatasetVersionRequest) (*api.ListDatasetVersionReply, error) {
	query := &model.DatasetVersionQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	versionsTbl, totalSize, err := s.data.DatasetDao.ListDatasetVersion(ctx, query)
	if err != nil {
		return nil, err
	}
	versions := make([]*api.DatasetVersion, 0)
	for _, n := range versionsTbl {
		version := &api.DatasetVersion{}
		err := copier.CopyWithOption(version, n, copier.Option{DeepCopy: true})
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		version.CreatedAt = n.CreatedAt.Unix()
		version.UpdatedAt = n.UpdatedAt.Unix()
		versions = append(versions, version)
	}
	return &api.ListDatasetVersionReply{
		TotalSize: totalSize,
		Versions:  versions,
	}, nil
}

func (s *datasetService) ListCommDatasetVersion(ctx context.Context, req *api.ListCommDatasetVersionRequest) (*api.ListCommDatasetVersionReply, error) {
	query := &model.CommDatasetVersionQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	versionsTbl, totalSize, err := s.data.DatasetDao.ListCommDatasetVersion(ctx, query)
	if err != nil {
		return nil, err
	}

	versions := make([]*api.DatasetVersion, 0)
	for _, n := range versionsTbl {
		version := &api.DatasetVersion{}
		err := copier.Copy(version, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		version.CreatedAt = n.CreatedAt.Unix()
		version.UpdatedAt = n.UpdatedAt.Unix()
		if n.Cache == nil {
			version.Cache = nil
		} else {
			version.Cache.Quota = n.Cache.Quota
		}
		versions = append(versions, version)
	}

	return &api.ListCommDatasetVersionReply{
		TotalSize: totalSize,
		Versions:  versions,
	}, nil
}

func (s *datasetService) ShareDatasetVersion(ctx context.Context, req *api.ShareDatasetVersionRequest) (*api.ShareDatasetVersionReply, error) {
	vInt, err := common.VersionStrParse(req.Version)
	if err != nil {
		return nil, err
	}

	_, err = s.data.DatasetDao.GetDatasetVersion(ctx, req.DatasetId, req.Version)
	if err != nil {
		return nil, err
	}

	accesses, err := s.data.DatasetDao.ListDatasetAccess(ctx, &model.DatasetAccessQuery{DatasetId: req.DatasetId})
	if err != nil {
		return nil, err
	}

	if len(accesses) == 0 {
		err := s.data.DatasetDao.CreateDatasetAccess(ctx, &model.DatasetAccess{
			DatasetId: req.DatasetId,
			SpaceId:   req.ShareSpaceId,
		})
		if err != nil {
			return nil, err
		}
	}

	versionAccesses, err := s.data.DatasetDao.ListDatasetVersionAccess(ctx, &model.DatasetVersionAccessQuery{
		DatasetId: req.DatasetId,
		Version:   req.Version,
		SpaceId:   req.ShareSpaceId,
	})
	if err != nil {
		return nil, err
	}

	if len(versionAccesses) > 0 {
		return nil, errors.Errorf(nil, errors.ErrorDatasetAlreadyShared)
	}

	err = s.data.DatasetDao.CreateDatasetVersionAccess(ctx, &model.DatasetVersionAccess{
		DatasetId:  req.DatasetId,
		Version:    req.Version,
		VersionInt: vInt,
		SpaceId:    req.ShareSpaceId,
	})
	if err != nil {
		return nil, err
	}

	return &api.ShareDatasetVersionReply{SharedAt: time.Now().Unix()}, nil
}

func (s *datasetService) CloseShareDatasetVersion(ctx context.Context, req *api.CloseShareDatasetVersionRequest) (*api.CloseShareDatasetVersionReply, error) {
	err := s.data.DatasetDao.DeleteDatasetVersionAccess(ctx, &model.DatasetVersionAccessDelete{
		DatasetId: req.DatasetId,
		Version:   req.Version,
		SpaceId:   req.ShareSpaceId,
	})
	if err != nil {
		return nil, err
	}

	accesses, err := s.data.DatasetDao.ListDatasetVersionAccess(ctx, &model.DatasetVersionAccessQuery{
		DatasetId: req.DatasetId,
		SpaceId:   req.ShareSpaceId,
	})
	if err != nil {
		return nil, err
	}

	if len(accesses) == 0 {
		err := s.data.DatasetDao.DeleteDatasetAccess(ctx, &model.DatasetAccessDelete{
			DatasetId: req.DatasetId,
			SpaceId:   req.ShareSpaceId,
		})
		if err != nil {
			return nil, err
		}
	}

	return &api.CloseShareDatasetVersionReply{ClosedAt: time.Now().Unix()}, nil
}

func (s *datasetService) DeleteDatasetVersion(ctx context.Context, req *api.DeleteDatasetVersionRequest) (*api.DeleteDatasetVersionReply, error) {

	dataset, err := s.data.DatasetDao.GetDataset(ctx, req.DatasetId)
	if err != nil {
		return nil, err
	}

	version, err := s.data.DatasetDao.GetDatasetVersion(ctx, req.DatasetId, req.Version)
	if err != nil {
		return nil, err
	}

	err = s.data.DatasetDao.DeleteDatasetVersionAccess(ctx, &model.DatasetVersionAccessDelete{
		DatasetId: req.DatasetId,
		Version:   req.Version,
	})
	if err != nil {
		return nil, err
	}

	err = s.data.DatasetDao.DeleteDatasetVersion(ctx, &model.DatasetVersionDelete{
		DatasetId: req.DatasetId,
		Version:   req.Version,
	})
	if err != nil {
		return nil, err
	}

	versionAccesses, err := s.data.DatasetDao.ListDatasetVersionAccess(ctx, &model.DatasetVersionAccessQuery{
		DatasetId: req.DatasetId,
	})
	if err != nil {
		return nil, err
	}

	if len(versionAccesses) == 0 {
		err := s.data.DatasetDao.DeleteDatasetAccess(ctx, &model.DatasetAccessDelete{
			DatasetId: req.DatasetId,
		})
		if err != nil {
			return nil, err
		}
	}

	_, totalSize, err := s.data.DatasetDao.ListDatasetVersion(ctx, &model.DatasetVersionQuery{
		DatasetId: req.DatasetId,
	})
	if err != nil {
		return nil, err
	}

	if totalSize == 0 {
		err = s.data.DatasetDao.DeleteDataset(ctx, req.DatasetId)
		if err != nil {
			return nil, err
		}
	}
	// 删除数据集版本Minio存储
	go utils.HandlePanicBG(func(i ...interface{}) {
		bucketName, objectName := getMinioPath(dataset, version.Version)
		s.data.Redis.SAddMinioRemovingObject(bucketName + "-" + objectName)
		defer s.data.Redis.SRemMinioRemovingObject(bucketName + "-" + objectName)
		s.data.Minio.RemoveObject(bucketName, objectName)
	})()
	return &api.DeleteDatasetVersionReply{DeletedAt: time.Now().Unix()}, nil
}

func (s *datasetService) DeleteDataset(ctx context.Context, req *api.DeleteDatasetRequest) (*api.DeleteDatasetReply, error) {
	dataset, err := s.data.DatasetDao.GetDataset(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	err = s.data.DatasetDao.DeleteDatasetVersionAccess(ctx, &model.DatasetVersionAccessDelete{
		DatasetId: req.Id,
	})
	if err != nil {
		return nil, err
	}

	err = s.data.DatasetDao.DeleteDatasetVersion(ctx, &model.DatasetVersionDelete{
		DatasetId: req.Id,
	})
	if err != nil {
		return nil, err
	}

	err = s.data.DatasetDao.DeleteDatasetAccess(ctx, &model.DatasetAccessDelete{
		DatasetId: req.Id,
	})
	if err != nil {
		return nil, err
	}

	err = s.data.DatasetDao.DeleteDataset(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	// 删除数据集Minio存储
	go utils.HandlePanicBG(func(i ...interface{}) {
		bucket, object := getMinioPathObject(dataset)
		s.data.Redis.SAddMinioRemovingObject(bucket + "-" + object)
		defer s.data.Redis.SRemMinioRemovingObject(bucket + "-" + object)
		s.data.Minio.RemoveObject(bucket, object)
	})()

	// 减小数据类型引用
	_, _ = s.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: dataset.TypeId})
	for _, id := range dataset.ApplyIds {
		// 减小数据用途引用
		_, _ = s.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: id})
	}

	return &api.DeleteDatasetReply{DeletedAt: time.Now().Unix()}, nil
}

func (s *datasetService) UpdateDataset(ctx context.Context, req *api.UpdateDatasetRequest) (*api.UpdateDatasetReply, error) {
	dataset, err := s.data.DatasetDao.GetDataset(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if dataset.SpaceId != req.SpaceId || dataset.UserId != req.UserId || dataset.SourceType != int(req.SourceType) {
		return nil, errors.Errorf(nil, errors.ErrorDatasetNoPermission)
	}

	// 减小数据类型引用
	_, _ = s.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: dataset.TypeId})
	for _, id := range dataset.ApplyIds {
		// 减小数据用途引用
		_, _ = s.lableService.ReduceLableReferTimes(ctx, &api.ReduceLableReferTimesRequest{Id: id})
	}

	dataset.TypeId = req.TypeId
	dataset.ApplyIds = req.ApplyIds
	dataset.Desc = req.Desc
	err = s.data.DatasetDao.UpdateDatasetSelective(ctx, dataset)
	if err != nil {
		return nil, err
	}

	// 增加数据类型引用
	_, _ = s.lableService.IncreaseLableReferTimes(ctx, &api.IncreaseLableReferTimesRequest{Id: dataset.TypeId})
	// 增加数据用途引用
	for _, id := range dataset.ApplyIds {
		// 新增数据用途引用
		_, _ = s.lableService.IncreaseLableReferTimes(ctx, &api.IncreaseLableReferTimesRequest{Id: id})
	}

	return &api.UpdateDatasetReply{UpdatedAt: time.Now().Unix()}, nil
}

func (s *datasetService) UploadDatasetVersion(ctx context.Context, req *api.UploadDatasetVersionRequest) (*api.UploadDatasetVersionReply, error) {
	dataset, err := s.data.DatasetDao.GetDataset(ctx, req.DatasetId)

	if err != nil {
		return nil, err
	}

	version, err := s.data.DatasetDao.GetDatasetVersion(ctx, req.DatasetId, req.Version)

	if err != nil {
		return nil, err
	}

	if !utils.IntInSlice(version.Status, statusForUpload) {
		return nil, errors.Errorf(nil, errors.ErrorDatasetStatusForbidden)
	}
	uploadUrl, err := s.getUploadUrl(dataset, req.Version, req.FileName, req.Domain)

	if err != nil {
		return nil, err
	}
	return &api.UploadDatasetVersionReply{
		UploadUrl: uploadUrl,
	}, nil
}

func (s *datasetService) GetDatasetVersion(ctx context.Context, req *api.GetDatasetVersionRequest) (*api.GetDatasetVersionReply, error) {
	reply := &api.GetDatasetVersionReply{}

	ids := []*api.DatasetVersionId{{
		DatasetId: req.DatasetId,
		Version:   req.Version,
	}}
	listDatasetReply, err := s.ListDatasetVersion(ctx, &api.ListDatasetVersionRequest{Ids: ids})
	if err != nil {
		return nil, err
	}
	if len(listDatasetReply.Versions) == 0 {
		return nil, errors.Errorf(nil, errors.ErrorDBFindEmpty)
	}
	reply.Version = listDatasetReply.Versions[0]

	getDatasetReply, err := s.ListDataset(ctx, &api.ListDatasetRequest{Ids: []string{req.DatasetId}})
	if err != nil {
		return nil, err
	}
	reply.Dataset = getDatasetReply.Datasets[0]

	versionAccesses, err := s.data.DatasetDao.ListDatasetVersionAccess(ctx, &model.DatasetVersionAccessQuery{DatasetId: req.DatasetId, Version: req.Version})
	if err != nil {
		return nil, err
	}

	err = copier.Copy(&reply.VersionAccesses, versionAccesses)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	return reply, nil
}

func (s *datasetService) ListDatasetVersionFile(ctx context.Context, req *api.ListDatasetVersionFileRequest) (*api.ListDatasetVersionFileReply, error) {
	reply := &api.ListDatasetVersionFileReply{}

	dataset, err := s.data.DatasetDao.GetDataset(ctx, req.DatasetId)
	if err != nil {
		return nil, err
	}

	version, err := s.data.DatasetDao.GetDatasetVersion(ctx, req.DatasetId, req.Version)
	if err != nil {
		return nil, err
	}

	bucketName, objectName := getMinioPath(dataset, version.Version)
	objectPrefix := objectName + "/"
	objects, err := s.data.Minio.ListObjects(bucketName, objectPrefix+req.Path, false)
	if err != nil {
		return nil, err
	}
	for _, i := range objects {
		file := &api.ListDatasetVersionFileReply_File{
			Name: filepath.Base(i.Name),
			Path: i.Name[len(objectPrefix):],
		}
		if file.Path[len(file.Path)-1:] == "/" {
			file.Type = "directory"
		} else {
			file.Type = "file"
		}

		reply.Files = append(reply.Files, file)
	}

	return reply, nil
}

func (s *datasetService) UpdateDatasetVersion(ctx context.Context, req *api.UpdateDatasetVersionRequest) (*api.UpdateDatasetVersionReply, error) {
	dataset, err := s.data.DatasetDao.GetDataset(ctx, req.DatasetId)
	if err != nil {
		return nil, err
	}
	if dataset.SpaceId != req.SpaceId || dataset.UserId != req.UserId || dataset.SourceType != int(req.SourceType) {
		return nil, errors.Errorf(nil, errors.ErrorDatasetNoPermission)
	}

	version, err := s.data.DatasetDao.GetDatasetVersion(ctx, req.DatasetId, req.Version)
	if err != nil {
		return nil, err
	}

	version.Desc = req.Desc
	err = s.data.DatasetDao.UpdateDatasetVersionSelective(ctx, version)
	if err != nil {
		return nil, err
	}

	return &api.UpdateDatasetVersionReply{UpdatedAt: time.Now().Unix()}, nil
}

func getTempMinioPath(dataset *model.Dataset, version string, fileName string) (bucketName string, objectName string) {
	bucketName = common.GetMinioBucket()
	objectName = common.GetMinioUploadDataSetObject(dataset.Id, version, fileName)
	return
}

func getMinioPath(dataset *model.Dataset, version string) (bucketName string, objectName string) {
	if dataset.SourceType == int(api.DatasetSourceType_DST_PRE) {
		bucketName = common.GetMinioBucket()
		objectName = common.GetMinioPreDataSetObject(dataset.Id, version)
	} else {
		bucketName = common.GetMinioBucket()
		objectName = common.GetMinioDataSetObject(dataset.SpaceId, dataset.UserId, dataset.Id, version)
	}
	return
}

func getMinioPathObject(dataset *model.Dataset) (bucketName string, objectName string) {
	if dataset.SourceType == int(api.DatasetSourceType_DST_PRE) {
		bucketName = common.GetMinioBucket()
		objectName = common.GetMinioPreDataSetPathObject(dataset.Id)
	} else {
		bucketName = common.GetMinioBucket()
		objectName = common.GetMinioDataSetPathObject(dataset.SpaceId, dataset.UserId, dataset.Id)
	}
	return
}

func (s *datasetService) getUploadUrl(dataset *model.Dataset, version string, fileName string, domain string) (string, error) {
	bucketName, objectName := getTempMinioPath(dataset, version, fileName)
	url, err := s.data.Minio.PresignedUploadObject(bucketName, objectName, domain)
	if err != nil {
		return "", err
	}

	return url.String(), nil
}

func (s *datasetService) getPath(dataset *model.Dataset, newV string) string {
	toBucket, toObject := getMinioPath(dataset, newV)
	toPath := fmt.Sprintf("%s/%s", toBucket, toObject)
	return toPath
}
func (s *datasetService) CreateCache(ctx context.Context, req *api.CacheRequest) (*api.CacheReply, error) {
	version, err := s.data.DatasetDao.GetDatasetVersion(ctx, req.DatasetId, req.Version)
	if err != nil {
		return nil, err
	}

	dataset, err := s.data.DatasetDao.GetDataset(ctx, req.DatasetId)
	if err != nil {
		return nil, err
	}

	if dataset.UserId == "" {
		return nil, errors.Errorf(nil, errors.OnlySupportCacheUserDataset)
	}
	namespace := dataset.UserId
	cacheName := common.GetCacheName()

	if version.Cache != nil {
		return nil, errors.Errorf(nil, errors.ErrorDatasetCacheExist)
	}

	quantity, err := resource.ParseQuantity(req.Cache.Quota)
	if err != nil {
		return nil, errors.Errorf(nil, errors.ErrorFormatParseFailed)
	}

	option := s.conf.Service.Dataset.Cache
	alluxioRuntime := fluidv1.AlluxioRuntime{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "data.fluid.io/v1alpha1",
			Kind:       "AlluxioRuntime",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cacheName,
			Namespace: namespace,
		},
		Spec: fluidv1.AlluxioRuntimeSpec{
			Replicas: 1,
			TieredStore: fluidv1.TieredStore{Levels: []fluidv1.Level{
				{MediumType: Common.MediumType(option.Mediumtype),
					Path:  fmt.Sprintf("%s", option.Path),
					Quota: &quantity,
					High:  "0.95",
					Low:   "0.7"},
			}},
		},
	}
	err = s.data.Cluster.CreateAlluxioRuntime(ctx, &alluxioRuntime)
	if err != nil {
		return nil, err
	}

	Dataset := fluidv1.Dataset{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "data.fluid.io/v1alpha1",
			Kind:       "Dataset",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cacheName,
			Namespace: namespace,
		},
		Spec: fluidv1.DatasetSpec{
			Mounts: []fluidv1.Mount{
				{
					MountPoint: fmt.Sprintf("s3://%s", version.Path),
					Name:       cacheName,
					Options: map[string]string{
						"aws.accessKeyId":             s.conf.Data.Minio.Base.AccessKeyID,
						"aws.secretKey":               s.conf.Data.Minio.Base.SecretAccessKey,
						"alluxio.underfs.s3.endpoint": fmt.Sprintf("%s%s", "http://", s.conf.Data.Minio.Base.EndPoint),
					},
				},
			},
			PlacementMode: "Shared",
			AccessModes:   []v1.PersistentVolumeAccessMode{v1.ReadWriteMany},
		},
	}
	err = s.data.Cluster.CreateFluidDataset(ctx, &Dataset)
	if err != nil {
		return nil, err
	}

	cache := &model.Cache{
		Quota: req.Cache.Quota,
		Name:  cacheName,
	}
	err = s.data.DatasetDao.UpdateDatasetVersionSelective(ctx, &model.DatasetVersion{
		DatasetId: req.DatasetId,
		Version:   req.Version,
		Cache:     cache,
	})

	return &api.CacheReply{}, nil

}
func (s *datasetService) DeleteCache(ctx context.Context, req *api.DeleteCacheRequest) (*api.CacheReply, error) {
	dataset, err := s.data.DatasetDao.GetDataset(ctx, req.DatasetId)
	if err != nil {
		return nil, err
	}

	version, err := s.data.DatasetDao.GetDatasetVersion(ctx, req.DatasetId, req.Version)
	if err != nil {
		return nil, err
	}

	if version.Cache == nil {
		return nil, errors.Errorf(nil, errors.ErrorDatasetCacheNotExist)
	}

	query := &model.TrainJobListQuery{Statuses: []string{constant.RUNNING, constant.PENDING}, UserId: dataset.UserId}
	jobList, _, err := s.data.TrainJobDao.GetTrainJobList(ctx, query)
	if err != nil {
		return nil, err
	}

	//判断是否正在被运行中或者等待中的任务使用
	for _, job := range jobList {
		if req.DatasetId == job.DataSetId && req.Version == job.DataSetVersion {
			return nil, errors.Errorf(nil, errors.ErrorDatasetisBeingUsing)
		}
	}

	namespace := dataset.UserId
	cacheName := version.Cache.Name

	err = s.data.Cluster.DeleteFluidDataset(ctx, namespace, cacheName)
	if err != nil {
		return nil, err
	}
	err = s.data.Cluster.DeleteAlluxioRuntime(ctx, namespace, cacheName)
	if err != nil {
		return nil, err
	}
	err = s.data.DatasetDao.UpdateDatasetVersionCache(ctx, &model.DatasetVersion{
		DatasetId: req.DatasetId,
		Version:   req.Version,
		Cache:     nil,
	})
	if err != nil {
		return nil, err
	}
	return &api.CacheReply{}, nil
}
