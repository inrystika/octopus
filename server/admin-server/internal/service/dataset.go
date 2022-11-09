package service

import (
	"context"
	"github.com/jinzhu/copier"
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	innerapi "server/base-server/api/v1"
	"server/common/errors"
	"server/common/log"
	"server/common/utils"
)

type DatasetService struct {
	api.UnimplementedDatasetServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewDatasetService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.DatasetServiceServer {
	return &DatasetService{
		conf: conf,
		log:  log.NewHelper("DatasetService", logger),
		data: data,
	}
}

func (s *DatasetService) AddDatasetType(ctx context.Context, req *api.AddDatasetTypeRequest) (*api.AddDatasetTypeReply, error) {
	innerReq := &innerapi.AddLableRequest{
		RelegationType: int32(innerapi.Relegation_LABLE_RELEGATION_DATASET),
		LableType:      int32(innerapi.Type_LABLE_TYPE_DATASET_TYPE),
		LableDesc:      req.LableDesc,
	}

	innerReply, err := s.data.LableClient.AddLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.AddDatasetTypeReply{
		Lable: &api.DatasetLable{
			Id:         innerReply.Lable.Id,
			LableDesc:  innerReply.Lable.LableDesc,
			SourceType: innerReply.Lable.SourceType,
		},
	}, nil
}

func (s *DatasetService) ListDatasetType(ctx context.Context, req *api.ListDatasetTypeRequest) (*api.ListDatasetTypeReply, error) {
	innerReq := &innerapi.ListLableRequest{
		RelegationType: int32(innerapi.Relegation_LABLE_RELEGATION_DATASET),
		LableType:      int32(innerapi.Type_LABLE_TYPE_DATASET_TYPE),
		PageIndex:      req.PageIndex,
		PageSize:       req.PageSize,
	}

	innerReply, err := s.data.LableClient.ListLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListDatasetTypeReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *DatasetService) DeleteDatasetType(ctx context.Context, req *api.DeleteDatasetTypeRequest) (*api.DeleteDatasetTypeReply, error) {
	innerReq := &innerapi.DeleteLableRequest{
		Id: req.Id,
	}

	innerReply, err := s.data.LableClient.DeleteLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.DeleteDatasetTypeReply{
		DeletedAt: innerReply.DeletedAt,
	}, nil
}

func (s *DatasetService) UpdateDatasetType(ctx context.Context, req *api.UpdateDatasetTypeRequest) (*api.UpdateDatasetTypeReply, error) {
	innerReq := &innerapi.UpdateLableRequest{
		Id:        req.Id,
		LableDesc: req.LableDesc,
	}

	innerReply, err := s.data.LableClient.UpdateLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.UpdateDatasetTypeReply{
		UpdatedAt: innerReply.UpdatedAt,
	}, nil
}

func (s *DatasetService) AddDatasetApply(ctx context.Context, req *api.AddDatasetApplyRequest) (*api.AddDatasetApplyReply, error) {
	innerReq := &innerapi.AddLableRequest{
		RelegationType: int32(innerapi.Relegation_LABLE_RELEGATION_DATASET),
		LableType:      int32(innerapi.Type_LABLE_TYPE_DATASET_APPLY),
		LableDesc:      req.LableDesc,
	}

	innerReply, err := s.data.LableClient.AddLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.AddDatasetApplyReply{
		Lable: &api.DatasetLable{
			Id:         innerReply.Lable.Id,
			LableDesc:  innerReply.Lable.LableDesc,
			SourceType: innerReply.Lable.SourceType,
		},
	}, nil
}

func (s *DatasetService) ListDatasetApply(ctx context.Context, req *api.ListDatasetApplyRequest) (*api.ListDatasetApplyReply, error) {
	innerReq := &innerapi.ListLableRequest{
		RelegationType: int32(innerapi.Relegation_LABLE_RELEGATION_DATASET),
		LableType:      int32(innerapi.Type_LABLE_TYPE_DATASET_APPLY),
		PageIndex:      req.PageIndex,
		PageSize:       req.PageSize,
	}

	innerReply, err := s.data.LableClient.ListLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListDatasetApplyReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *DatasetService) DeleteDatasetApply(ctx context.Context, req *api.DeleteDatasetApplyRequest) (*api.DeleteDatasetApplyReply, error) {
	innerReq := &innerapi.DeleteLableRequest{
		Id: req.Id,
	}

	innerReply, err := s.data.LableClient.DeleteLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.DeleteDatasetApplyReply{
		DeletedAt: innerReply.DeletedAt,
	}, nil
}

func (s *DatasetService) UpdateDatasetApply(ctx context.Context, req *api.UpdateDatasetApplyRequest) (*api.UpdateDatasetApplyReply, error) {
	innerReq := &innerapi.UpdateLableRequest{
		Id:        req.Id,
		LableDesc: req.LableDesc,
	}

	innerReply, err := s.data.LableClient.UpdateLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.UpdateDatasetApplyReply{
		UpdatedAt: innerReply.UpdatedAt,
	}, nil
}

func (s *DatasetService) ListUserDataset(ctx context.Context, req *api.ListUserDatasetRequest) (*api.ListUserDatasetReply, error) {
	innerReq := &innerapi.ListDatasetRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.SourceType = innerapi.DatasetSourceType_DST_USER

	innerReply, err := s.data.DatasetClient.ListDataset(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListUserDatasetReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}
	err = s.assignValue(ctx, reply.Datasets)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *DatasetService) ListPreDataset(ctx context.Context, req *api.ListPreDatasetRequest) (*api.ListPreDatasetReply, error) {
	innerReq := &innerapi.ListDatasetRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.SourceType = innerapi.DatasetSourceType_DST_PRE

	innerReply, err := s.data.DatasetClient.ListDataset(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListPreDatasetReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *DatasetService) ListDatasetVersion(ctx context.Context, req *api.ListDatasetVersionRequest) (*api.ListDatasetVersionReply, error) {

	innerReq := &innerapi.ListDatasetVersionRequest{}
	err := copier.Copy(innerReq, req)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.DatasetClient.ListDatasetVersion(ctx, innerReq)

	if err != nil {
		return nil, err
	}

	reply := &api.ListDatasetVersionReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *DatasetService) CreateDataset(ctx context.Context, req *api.CreateDatasetRequest) (*api.CreateDatasetReply, error) {
	innerReq := &innerapi.CreateDatasetRequest{
		SourceType: innerapi.DatasetSourceType_DST_PRE,
		Name:       req.Name,
		TypeId:     req.TypeId,
		ApplyIds:   req.ApplyIds,
		Desc:       req.Desc,
	}

	reply, err := s.data.DatasetClient.CreateDataset(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.CreateDatasetReply{
		Id:      reply.Id,
		Version: reply.Version,
	}, nil
}

func (s *DatasetService) CreateDatasetVersion(ctx context.Context, req *api.CreateDatasetVersionRequest) (*api.CreateDatasetVersionReply, error) {
	reply, err := s.data.DatasetClient.CreateDatasetVersion(ctx, &innerapi.CreateDatasetVersionRequest{
		DatasetId: req.DatasetId,
		Desc:      req.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &api.CreateDatasetVersionReply{
		DatasetId: reply.DatasetId,
		Version:   reply.Version,
	}, nil
}

func (s *DatasetService) DeleteDataset(ctx context.Context, req *api.DeleteDatasetRequest) (*api.DeleteDatasetReply, error) {
	reply, err := s.data.DatasetClient.DeleteDataset(ctx, &innerapi.DeleteDatasetRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &api.DeleteDatasetReply{DeletedAt: reply.DeletedAt}, nil
}

func (s *DatasetService) DeleteDatasetVersion(ctx context.Context, req *api.DeleteDatasetVersionRequest) (*api.DeleteDatasetVersionReply, error) {
	reply, err := s.data.DatasetClient.DeleteDatasetVersion(ctx, &innerapi.DeleteDatasetVersionRequest{
		DatasetId: req.DatasetId,
		Version:   req.Version,
	})
	if err != nil {
		return nil, err
	}

	return &api.DeleteDatasetVersionReply{DeletedAt: reply.DeletedAt}, nil
}

func (s *DatasetService) ConfirmUploadDatasetVersion(ctx context.Context, req *api.ConfirmUploadDatasetVersionRequest) (*api.ConfirmUploadDatasetVersionReply, error) {
	reply, err := s.data.DatasetClient.ConfirmUploadDatasetVersion(ctx, &innerapi.ConfirmUploadDatasetVersionRequest{
		DatasetId: req.DatasetId,
		Version:   req.Version,
		FileName:  req.FileName,
	})
	if err != nil {
		return nil, err
	}

	return &api.ConfirmUploadDatasetVersionReply{UpdatedAt: reply.UpdatedAt}, nil
}

func (s *DatasetService) UploadDatasetVersion(ctx context.Context, req *api.UploadDatasetVersionRequest) (*api.UploadDatasetVersionReply, error) {
	reply, err := s.data.DatasetClient.UploadDatasetVersion(ctx, &innerapi.UploadDatasetVersionRequest{
		DatasetId: req.DatasetId,
		Version:   req.Version,
		FileName:  req.FileName,
		Domain:    req.Domain,
	})
	if err != nil {
		return nil, err
	}

	return &api.UploadDatasetVersionReply{
		UploadUrl: reply.UploadUrl,
	}, nil
}

func (s *DatasetService) ListDatasetVersionFile(ctx context.Context, req *api.ListDatasetVersionFileRequest) (*api.ListDatasetVersionFileReply, error) {
	innerReply, err := s.data.DatasetClient.ListDatasetVersionFile(ctx, &innerapi.ListDatasetVersionFileRequest{
		DatasetId: req.DatasetId,
		Version:   req.Version,
		Path:      req.Path,
	})
	if err != nil {
		return nil, err
	}

	reply := &api.ListDatasetVersionFileReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *DatasetService) UpdatePreDataset(ctx context.Context, req *api.UpdatePreDatasetRequest) (*api.UpdatePreDatasetReply, error) {
	reply, err := s.data.DatasetClient.UpdateDataset(ctx, &innerapi.UpdateDatasetRequest{
		SpaceId:    "",
		UserId:     "",
		Id:         req.DatasetId,
		SourceType: innerapi.DatasetSourceType_DST_PRE,
		TypeId:     req.TypeId,
		ApplyIds:   req.ApplyIds,
		Desc:       req.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &api.UpdatePreDatasetReply{
		UpdatedAt: reply.UpdatedAt,
	}, nil
}

func (s *DatasetService) UpdatePreDatasetVersion(ctx context.Context, req *api.UpdatePreDatasetVersionRequest) (*api.UpdatePreDatasetVersionReply, error) {
	reply, err := s.data.DatasetClient.UpdateDatasetVersion(ctx, &innerapi.UpdateDatasetVersionRequest{
		SpaceId:    "",
		UserId:     "",
		DatasetId:  req.DatasetId,
		Version:    req.Version,
		SourceType: innerapi.DatasetSourceType_DST_PRE,
		Desc:       req.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &api.UpdatePreDatasetVersionReply{
		UpdatedAt: reply.UpdatedAt,
	}, nil
}

func (s *DatasetService) assignValue(ctx context.Context, datasets []*api.Dataset) error {
	if len(datasets) > 0 {
		userIdMap := map[string]interface{}{}
		spaceIdMap := map[string]interface{}{}
		for _, i := range datasets {
			if i.UserId != "" {
				userIdMap[i.UserId] = true
			}

			if i.SpaceId != "" {
				spaceIdMap[i.SpaceId] = true
			}
		}

		userMap := map[string]*innerapi.UserItem{}
		if len(userIdMap) > 0 {
			users, err := s.data.UserClient.ListUserInCond(ctx, &innerapi.ListUserInCondRequest{
				Ids: utils.MapKeyToSlice(userIdMap),
			})
			if err != nil {
				return err
			}
			for _, i := range users.Users {
				userMap[i.Id] = i
			}
		}

		spaceMap := map[string]*innerapi.WorkspaceItem{}
		if len(spaceIdMap) > 0 {
			spaces, err := s.data.WorkspaceClient.ListWorkspaceInCond(ctx, &innerapi.ListWorkspaceInCondRequest{
				Ids: utils.MapKeyToSlice(spaceIdMap),
			})
			if err != nil {
				return err
			}
			for _, i := range spaces.Workspaces {
				spaceMap[i.Id] = i

			}

		}

		for _, i := range datasets {
			if v, ok := userMap[i.UserId]; ok {
				i.UserName = v.FullName
				i.UserEmail = v.Email
			}

			if v, ok := spaceMap[i.SpaceId]; ok {
				i.SpaceName = v.Name
			}
		}
	}

	return nil
}
func (s *DatasetService) CreateDatasetVersionCache(ctx context.Context, req *api.CacheRequest) (*api.CacheReply, error) {
	innerReq := &innerapi.CacheRequest{}
	err := copier.Copy(innerReq, req)
	_, err = s.data.DatasetClient.CreateCache(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.CacheReply{}, nil
}
func (s *DatasetService) DeleteDatasetVersionCache(ctx context.Context, req *api.DeleteCacheRequest) (*api.CacheReply, error) {
	innerReq := &innerapi.DeleteCacheRequest{}
	err := copier.Copy(innerReq, req)
	_, err = s.data.DatasetClient.DeleteCache(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.CacheReply{}, nil
}
