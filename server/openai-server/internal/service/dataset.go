package service

import (
	"context"
	innerapi "server/base-server/api/v1"
	commctx "server/common/context"
	"server/common/errors"
	"server/common/log"
	"server/common/session"
	ss "server/common/session"
	"server/common/utils/collections/set"
	api "server/openai-server/api/v1"
	"server/openai-server/internal/conf"
	"server/openai-server/internal/data"

	"github.com/jinzhu/copier"
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

func (s *DatasetService) checkDatasetPerm(ctx context.Context, datasetId string, session *session.Session) error {
	reply, err := s.data.DatasetClient.GetDataset(ctx, &innerapi.GetDatasetRequest{Id: datasetId})
	if err != nil {
		return err
	}

	if reply.Dataset.UserId != session.UserId {
		return errors.Errorf(nil, errors.ErrorNotAuthorized)
	}
	return nil
}

func (s *DatasetService) checkVersionQueryPerm(ctx context.Context, datasetId string, version string, session *session.Session) error {
	reply, err := s.data.DatasetClient.GetDatasetVersion(ctx, &innerapi.GetDatasetVersionRequest{DatasetId: datasetId, Version: version})
	if err != nil {
		return err
	}
	if session.UserId != reply.Dataset.UserId && reply.Dataset.SourceType == innerapi.DatasetSourceType_DST_USER {
		hasPerm := false
		for _, i := range reply.VersionAccesses {
			if session.GetWorkspace() == i.SpaceId {
				hasPerm = true
			}
		}

		if !hasPerm {
			return errors.Errorf(err, errors.ErrorDatasetNoPermission)
		}
	}

	return nil
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

func (s *DatasetService) CreateDataset(ctx context.Context, req *api.CreateDatasetRequest) (*api.CreateDatasetReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.CreateDatasetRequest{
		SpaceId:    session.GetWorkspace(),
		UserId:     session.UserId,
		SourceType: innerapi.DatasetSourceType_DST_USER,
		Name:       req.Name,
		TypeId:     req.TypeId,
		ApplyId:    req.ApplyId,
		Desc:       req.Desc,
	}

	reply, err := s.data.DatasetClient.CreateDataset(ctx, innerReq)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return &api.CreateDatasetReply{
		Id:      reply.Id,
		Version: reply.Version,
	}, nil
}

func (s *DatasetService) ListMyDataset(ctx context.Context, req *api.ListMyDatasetRequest) (*api.ListMyDatasetReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.ListDatasetRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.SourceType = innerapi.DatasetSourceType_DST_USER
	innerReq.UserId = session.UserId
	innerReq.SpaceId = session.GetWorkspace()

	innerReply, err := s.data.DatasetClient.ListDataset(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListMyDatasetReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (s *DatasetService) ListPreDataset(ctx context.Context, req *api.ListPreDatasetRequest) (*api.ListPreDatasetReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

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

func (s *DatasetService) ListCommDataset(ctx context.Context, req *api.ListCommDatasetRequest) (*api.ListCommDatasetReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	innerReq := &innerapi.ListCommDatasetRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.SourceType = innerapi.DatasetSourceType_DST_USER
	innerReq.ShareSpaceId = session.GetWorkspace()

	innerReply, err := s.data.DatasetClient.ListCommDataset(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListCommDatasetReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	userIds := make([]string, 0)
	for _, i := range innerReply.Datasets {
		userIds = append(userIds, i.UserId)
	}

	users, err := s.listUserInCond(ctx, set.NewStrings(userIds...).Values())
	for _, i := range reply.Datasets {
		for _, j := range innerReply.Datasets {
			if i.Id == j.Id {
				i.UserName = users[j.UserId].FullName
			}
		}
	}

	return reply, nil
}

func (s *DatasetService) listUserInCond(ctx context.Context, ids []string) (map[string]*innerapi.UserItem, error) {
	userMap := map[string]*innerapi.UserItem{}
	users, err := s.data.UserClient.ListUserInCond(ctx, &innerapi.ListUserInCondRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	for _, i := range users.Users {
		userMap[i.Id] = i
	}
	return userMap, nil
}

func (s *DatasetService) DeleteDataset(ctx context.Context, req *api.DeleteDatasetRequest) (*api.DeleteDatasetReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	err := s.checkDatasetPerm(ctx, req.Id, session)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.DatasetClient.DeleteDataset(ctx, &innerapi.DeleteDatasetRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &api.DeleteDatasetReply{DeletedAt: reply.DeletedAt}, nil
}

func (s *DatasetService) CreateDatasetVersion(ctx context.Context, req *api.CreateDatasetVersionRequest) (*api.CreateDatasetVersionReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	err := s.checkDatasetPerm(ctx, req.DatasetId, session)
	if err != nil {
		return nil, err
	}

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

func (s *DatasetService) ListDatasetVersion(ctx context.Context, req *api.ListDatasetVersionRequest) (*api.ListDatasetVersionReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	if req.Shared {
		return s.listCommDatasetVersion(ctx, session, req)
	} else {
		return s.listDatasetVersion(ctx, session, req)
	}
}

func (s *DatasetService) listDatasetVersion(ctx context.Context, session *session.Session, req *api.ListDatasetVersionRequest) (*api.ListDatasetVersionReply, error) {
	reply := &api.ListDatasetVersionReply{}

	innerReq := &innerapi.ListDatasetVersionRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	innerReply, err := s.data.DatasetClient.ListDatasetVersion(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	commReq := &innerapi.ListCommDatasetVersionRequest{}
	err = copier.Copy(commReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	commReq.ShareSpaceId = session.GetWorkspace()
	commReply, err := s.data.DatasetClient.ListCommDatasetVersion(ctx, commReq)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	for _, v := range reply.Versions {
		for _, c := range commReply.Versions {
			if v.DatasetId == c.DatasetId && v.Version == c.Version {
				v.Shared = true
				break
			}
		}
	}

	return reply, nil
}

func (s *DatasetService) listCommDatasetVersion(ctx context.Context, session *session.Session, req *api.ListDatasetVersionRequest) (*api.ListDatasetVersionReply, error) {
	reply := &api.ListDatasetVersionReply{}

	innerReq := &innerapi.ListCommDatasetVersionRequest{}
	err := copier.Copy(innerReq, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	innerReq.ShareSpaceId = session.GetWorkspace()

	innerReply, err := s.data.DatasetClient.ListCommDatasetVersion(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	for _, v := range reply.Versions {
		v.Shared = true
	}

	return reply, nil
}

func (s *DatasetService) DeleteDatasetVersion(ctx context.Context, req *api.DeleteDatasetVersionRequest) (*api.DeleteDatasetVersionReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	err := s.checkDatasetPerm(ctx, req.DatasetId, session)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.DatasetClient.DeleteDatasetVersion(ctx, &innerapi.DeleteDatasetVersionRequest{
		DatasetId: req.DatasetId,
		Version:   req.Version,
	})
	if err != nil {
		return nil, err
	}

	return &api.DeleteDatasetVersionReply{DeletedAt: reply.DeletedAt}, nil
}

func (s *DatasetService) ShareDatasetVersion(ctx context.Context, req *api.ShareDatasetVersionRequest) (*api.ShareDatasetVersionReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	err := s.checkDatasetPerm(ctx, req.DatasetId, session)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.DatasetClient.ShareDatasetVersion(ctx, &innerapi.ShareDatasetVersionRequest{
		DatasetId:    req.DatasetId,
		Version:      req.Version,
		ShareSpaceId: session.GetWorkspace(),
	})
	if err != nil {
		return nil, err
	}

	return &api.ShareDatasetVersionReply{SharedAt: reply.SharedAt}, nil
}

func (s *DatasetService) CloseShareDatasetVersion(ctx context.Context, req *api.CloseShareDatasetVersionRequest) (*api.CloseShareDatasetVersionReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	err := s.checkDatasetPerm(ctx, req.DatasetId, session)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.DatasetClient.CloseShareDatasetVersion(ctx, &innerapi.CloseShareDatasetVersionRequest{
		DatasetId:    req.DatasetId,
		Version:      req.Version,
		ShareSpaceId: session.GetWorkspace(),
	})
	if err != nil {
		return nil, err
	}

	return &api.CloseShareDatasetVersionReply{ClosedAt: reply.ClosedAt}, nil
}

func (s *DatasetService) ConfirmUploadDatasetVersion(ctx context.Context, req *api.ConfirmUploadDatasetVersionRequest) (*api.ConfirmUploadDatasetVersionReply, error) {
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	err := s.checkDatasetPerm(ctx, req.DatasetId, session)
	if err != nil {
		return nil, err
	}

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
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	err := s.checkDatasetPerm(ctx, req.DatasetId, session)
	if err != nil {
		return nil, err
	}

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
	session := session.SessionFromContext(ctx)
	if session == nil {
		return nil, errors.Errorf(nil, errors.ErrorUserNoAuthSession)
	}

	err := s.checkVersionQueryPerm(ctx, req.DatasetId, req.Version, session)
	if err != nil {
		return nil, err
	}

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

func (s *DatasetService) UpdateMyDataset(ctx context.Context, req *api.UpdateMyDatasetRequest) (*api.UpdateMyDatasetReply, error) {
	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.DatasetClient.UpdateDataset(ctx, &innerapi.UpdateDatasetRequest{
		SpaceId:    spaceId,
		UserId:     userId,
		Id:         req.DatasetId,
		SourceType: innerapi.DatasetSourceType_DST_USER,
		TypeId:     req.TypeId,
		ApplyId:    req.ApplyId,
		Desc:       req.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &api.UpdateMyDatasetReply{
		UpdatedAt: reply.UpdatedAt,
	}, nil
}

func (s *DatasetService) UpdateMyDatasetVersion(ctx context.Context, req *api.UpdateMyDatasetVersionRequest) (*api.UpdateMyDatasetVersionReply, error) {
	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.DatasetClient.UpdateDatasetVersion(ctx, &innerapi.UpdateDatasetVersionRequest{
		SpaceId:    spaceId,
		UserId:     userId,
		DatasetId:  req.DatasetId,
		Version:    req.Version,
		SourceType: innerapi.DatasetSourceType_DST_USER,
		Desc:       req.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &api.UpdateMyDatasetVersionReply{
		UpdatedAt: reply.UpdatedAt,
	}, nil
}

func (s *DatasetService) getUserIdAndSpaceId(ctx context.Context) (string, string, error) {
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
