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
)

type AlgorithmService struct {
	api.UnimplementedAlgorithmServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewAlgorithmService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.AlgorithmServer {
	return &AlgorithmService{
		conf: conf,
		log:  log.NewHelper("AlgorithmService", logger),
		data: data,
	}
}

// 查询预置算法列表
func (s *AlgorithmService) ListPreAlgorithm(ctx context.Context, req *api.ListPreAlgorithmRequest) (*api.ListPreAlgorithmReply, error) {
	reply, err := s.data.AlgorithmClient.ListPreAlgorithm(ctx, &innterapi.ListPreAlgorithmRequest{
		PageIndex:        req.PageIndex,
		PageSize:         req.PageSize,
		SortBy:           req.SortBy,
		OrderBy:          req.OrderBy,
		AlgorithmVersion: req.AlgorithmVersion,
		SearchKey:        req.SearchKey,
		CreatedAtGte:     req.CreatedAtGte,
		CreatedAtLt:      req.CreatedAtLt,
	})
	if err != nil {
		return nil, err
	}

	algorithms := make([]*api.AlgorithmDetail, 0)
	for _, algorithm := range reply.Algorithms {
		algorithms = append(algorithms, s.algorithmTransfer(ctx, algorithm))
	}

	return &api.ListPreAlgorithmReply{
		TotalSize:  reply.TotalSize,
		Algorithms: algorithms,
	}, nil
}

// 查询我的算法列表
func (s *AlgorithmService) ListMyAlgorithm(ctx context.Context, req *api.ListMyAlgorithmRequest) (*api.ListMyAlgorithmReply, error) {

	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.AlgorithmClient.ListMyAlgorithm(ctx, &innterapi.ListMyAlgorithmRequest{
		SpaceId:          spaceId,
		UserId:           userId,
		PageIndex:        req.PageIndex,
		PageSize:         req.PageSize,
		SortBy:           req.SortBy,
		OrderBy:          req.OrderBy,
		AlgorithmVersion: req.AlgorithmVersion,
		SearchKey:        req.SearchKey,
		CreatedAtGte:     req.CreatedAtGte,
		CreatedAtLt:      req.CreatedAtLt,
	})
	if err != nil {
		return nil, err
	}

	algorithms := make([]*api.AlgorithmDetail, 0)
	for _, algorithm := range reply.Algorithms {
		algorithms = append(algorithms, s.algorithmTransfer(ctx, algorithm))
	}

	return &api.ListMyAlgorithmReply{
		TotalSize:  reply.TotalSize,
		Algorithms: algorithms,
	}, nil
}

// 查询公共算法列表
func (s *AlgorithmService) ListCommAlgorithm(ctx context.Context, req *api.ListCommAlgorithmRequest) (*api.ListCommAlgorithmReply, error) {
	_, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.AlgorithmClient.ListCommAlgorithm(ctx, &innterapi.ListCommAlgorithmRequest{
		SpaceId:          spaceId,
		PageIndex:        req.PageIndex,
		PageSize:         req.PageSize,
		SortBy:           req.SortBy,
		OrderBy:          req.OrderBy,
		AlgorithmVersion: req.AlgorithmVersion,
		SearchKey:        req.SearchKey,
		CreatedAtGte:     req.CreatedAtGte,
		CreatedAtLt:      req.CreatedAtLt,
	})
	if err != nil {
		return nil, err
	}

	algorithms := make([]*api.AlgorithmDetail, 0)
	for _, algorithm := range reply.Algorithms {
		algorithms = append(algorithms, s.algorithmTransfer(ctx, algorithm))
	}

	return &api.ListCommAlgorithmReply{
		TotalSize:  reply.TotalSize,
		Algorithms: algorithms,
	}, nil
}

// 查询算法版本列表
func (s *AlgorithmService) ListAlgorithmVersion(ctx context.Context, req *api.ListAlgorithmVersionRequest) (*api.ListAlgorithmVersionReply, error) {

	_, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.AlgorithmClient.ListAlgorithmVersion(ctx, &innterapi.ListAlgorithmVersionRequest{
		AlgorithmId: req.AlgorithmId,
		PageIndex:   req.PageIndex,
		PageSize:    req.PageSize,
		FileStatus:  req.FileStatus,
	})

	if err != nil {
		return nil, err
	}

	commReply, err := s.data.AlgorithmClient.ListCommAlgorithmVersion(ctx, &innterapi.ListCommAlgorithmVersionRequest{
		SpaceId:     spaceId,
		AlgorithmId: req.AlgorithmId,
		PageIndex:   req.PageIndex,
		PageSize:    req.PageSize,
	})
	noneShared := false
	if err != nil || commReply == nil {
		noneShared = true
	}
	algorithms := make([]*api.MyAlgorithmDetail, 0)
	for _, algorithm := range reply.Algorithms {
		if noneShared {
			algorithms = append(algorithms, s.myAlgorithmTransfer(ctx, algorithm, false))
		} else {
			isShared := false
			for _, commAlgorithm := range commReply.Algorithms {
				if algorithm.AlgorithmVersion == commAlgorithm.AlgorithmVersion {
					isShared = true
					break
				}
			}
			algorithms = append(algorithms, s.myAlgorithmTransfer(ctx, algorithm, isShared))
		}
	}

	return &api.ListAlgorithmVersionReply{
		TotalSize:  reply.TotalSize,
		Algorithms: algorithms,
	}, nil
}

// 新增我的算法版本
func (s *AlgorithmService) AddMyAlgorithmVersion(ctx context.Context, req *api.AddMyAlgorithmVersionRequest) (*api.AddMyAlgorithmVersionReply, error) {
	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.AlgorithmClient.AddMyAlgorithmVersion(ctx, &innterapi.AddMyAlgorithmVersionRequest{
		SpaceId:           spaceId,
		UserId:            userId,
		AlgorithmId:       req.AlgorithmId,
		OriVersion:        req.OriVersion,
		AlgorithmDescript: req.AlgorithmDescript,
	})

	if err != nil {
		return nil, err
	}

	return &api.AddMyAlgorithmVersionReply{
		AlgorithmId: reply.AlgorithmId,
		Version:     reply.Version,
		CreatedAt:   reply.CreatedAt,
	}, nil
}

// 查询公共算法版本列表
func (s *AlgorithmService) ListCommAlgorithmVersion(ctx context.Context, req *api.ListCommAlgorithmVersionRequest) (*api.ListCommAlgorithmVersionReply, error) {

	_, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.AlgorithmClient.ListCommAlgorithmVersion(ctx, &innterapi.ListCommAlgorithmVersionRequest{
		SpaceId:     spaceId,
		AlgorithmId: req.AlgorithmId,
		PageIndex:   req.PageIndex,
		PageSize:    req.PageSize,
		FileStatus:  req.FileStatus,
	})
	if err != nil {
		return nil, err
	}

	algorithms := make([]*api.AlgorithmDetail, 0)
	for _, algorithm := range reply.Algorithms {
		if algorithm.SpaceId != spaceId {
			err := errors.Errorf(nil, errors.ErrorFindAlgorithmAuthWrong)
			s.log.Errorw(ctx, err)
			return nil, err
		}

		algorithms = append(algorithms, s.algorithmTransfer(ctx, algorithm))
	}

	return &api.ListCommAlgorithmVersionReply{
		TotalSize:  reply.TotalSize,
		Algorithms: algorithms,
	}, nil
}

// 分享算法版本到公共算法
func (s *AlgorithmService) ShareAlgorithmVersion(ctx context.Context, req *api.ShareAlgorithmVersionRequest) (*api.ShareAlgorithmVersionReply, error) {

	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	ShareSpaceIdList := []string{}
	ShareSpaceIdList = append(ShareSpaceIdList, spaceId)

	reply, err := s.data.AlgorithmClient.ShareAlgorithmVersion(ctx, &innterapi.ShareAlgorithmVersionRequest{
		SpaceId:          spaceId,
		UserId:           userId,
		AlgorithmId:      req.AlgorithmId,
		Version:          req.Version,
		ShareSpaceIdList: ShareSpaceIdList,
	})
	if err != nil {
		return nil, err
	}

	return &api.ShareAlgorithmVersionReply{
		SharedAt: reply.SharedAt,
	}, nil
}

// 取消分享算法版本到公共算法
func (s *AlgorithmService) CloseShareAlgorithmVersion(ctx context.Context, req *api.CloseShareAlgorithmVersionRequest) (*api.CloseShareAlgorithmVersionReply, error) {

	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	ShareSpaceIdList := []string{}
	ShareSpaceIdList = append(ShareSpaceIdList, spaceId)

	reply, err := s.data.AlgorithmClient.CloseShareAlgorithmVersion(ctx, &innterapi.CloseShareAlgorithmVersionRequest{
		SpaceId:          spaceId,
		UserId:           userId,
		AlgorithmId:      req.AlgorithmId,
		Version:          req.Version,
		ShareSpaceIdList: ShareSpaceIdList,
	})
	if err != nil {
		return nil, err
	}

	return &api.CloseShareAlgorithmVersionReply{
		CloseSharedAt: reply.CloseSharedAt,
	}, nil
}

// 删除我的算法版本
func (s *AlgorithmService) DeleteMyAlgorithmVersion(ctx context.Context, req *api.DeleteMyAlgorithmVersionRequest) (*api.DeleteMyAlgorithmVersionReply, error) {
	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.AlgorithmClient.DeleteMyAlgorithmVersion(ctx, &innterapi.DeleteMyAlgorithmVersionRequest{
		SpaceId:     spaceId,
		UserId:      userId,
		AlgorithmId: req.AlgorithmId,
		Version:     req.Version,
	})
	if err != nil {
		return nil, err
	}

	return &api.DeleteMyAlgorithmVersionReply{
		DeletedAt: reply.DeletedAt,
	}, nil
}

// 删除我的算法
func (s *AlgorithmService) DeleteMyAlgorithm(ctx context.Context, req *api.DeleteMyAlgorithmRequest) (*api.DeleteMyAlgorithmReply, error) {

	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.AlgorithmClient.DeleteMyAlgorithm(ctx, &innterapi.DeleteMyAlgorithmRequest{
		SpaceId:     spaceId,
		UserId:      userId,
		AlgorithmId: req.AlgorithmId,
	})
	if err != nil {
		return nil, err
	}

	return &api.DeleteMyAlgorithmReply{
		DeletedAt: reply.DeletedAt,
	}, nil
}

// 压缩算法包
func (s *AlgorithmService) DownloadAlgorithmVersionCompress(ctx context.Context, req *api.DownloadAlgorithmVersionCompressRequest) (
	*api.DownloadAlgorithmVersionCompressReply, error) {

	reply, err := s.data.AlgorithmClient.DownloadAlgorithmVersionCompress(ctx, &innterapi.DownloadAlgorithmVersionCompressRequest{
		AlgorithmId: req.AlgorithmId,
		Version:     req.Version,
	})
	if err != nil {
		return nil, err
	}

	return &api.DownloadAlgorithmVersionCompressReply{
		CompressAt: reply.CompressAt,
	}, nil
}

// 复制算法
func (s *AlgorithmService) CopyAlgorithmVersion(ctx context.Context, req *api.CopyAlgorithmVersionRequest) (*api.CopyAlgorithmVersionReply, error) {

	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.AlgorithmClient.CopyAlgorithmVersion(ctx, &innterapi.CopyAlgorithmVersionRequest{
		SpaceId:           spaceId,
		UserId:            userId,
		AlgorithmId:       req.AlgorithmId,
		ModelName:         req.ModelName,
		Version:           req.Version,
		NewAlgorithmName:  req.NewAlgorithmName,
		AlgorithmDescript: req.AlgorithmDescript,
	})
	if err != nil {
		return nil, err
	}

	return &api.CopyAlgorithmVersionReply{
		NewAlgorithmId: reply.NewAlgorithmId,
		NewVersion:     reply.NewVersion,
		CreatedAt:      reply.CreatedAt,
	}, nil
}

// 下载算法
func (s *AlgorithmService) DownloadAlgorithmVersion(ctx context.Context, req *api.DownloadAlgorithmVersionRequest) (*api.DownloadAlgorithmVersionReply, error) {
	reply, err := s.data.AlgorithmClient.DownloadAlgorithmVersion(ctx, &innterapi.DownloadAlgorithmVersionRequest{
		AlgorithmId: req.AlgorithmId,
		Version:     req.Version,
		CompressAt:  req.CompressAt,
		Domain:      req.Domain,
	})
	if err != nil {
		return nil, err
	}

	return &api.DownloadAlgorithmVersionReply{
		DownloadUrl: reply.DownloadUrl,
	}, nil
}

// 新增我的算法
func (s *AlgorithmService) AddMyAlgorithm(ctx context.Context, req *api.AddMyAlgorithmRequest) (*api.AddMyAlgorithmReply, error) {
	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := s.data.AlgorithmClient.AddAlgorithm(ctx, &innterapi.AddAlgorithmRequest{
		SpaceId:           spaceId,
		UserId:            userId,
		IsPrefab:          false,
		IsEmpty:           req.IsEmpty,
		AlgorithmName:     req.AlgorithmName,
		ModelName:         req.ModelName,
		AlgorithmDescript: req.AlgorithmDescript,
	})
	if err != nil {
		return nil, err
	}
	return &api.AddMyAlgorithmReply{
		AlgorithmId: reply.AlgorithmId,
		Version:     reply.Version,
		CreatedAt:   reply.CreatedAt,
	}, nil
}

// 上传算法
func (s *AlgorithmService) UploadAlgorithm(ctx context.Context, req *api.UploadAlgorithmRequest) (*api.UploadAlgorithmReply, error) {
	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.AlgorithmClient.UploadAlgorithm(ctx, &innterapi.UploadAlgorithmRequest{
		SpaceId:     spaceId,
		UserId:      userId,
		AlgorithmId: req.AlgorithmId,
		Version:     req.Version,
		FileName:    req.FileName,
		Domain:      req.Domain,
	})
	if err != nil {
		return nil, err
	}

	return &api.UploadAlgorithmReply{
		UploadUrl: reply.UploadUrl,
	}, nil
}

// 上传算法确认
func (s *AlgorithmService) ConfirmUploadAlgorithm(ctx context.Context, req *api.ConfirmUploadAlgorithmRequest) (*api.ConfirmUploadAlgorithmReply, error) {

	userId, spaceId, err := s.getUserIdAndSpaceId(ctx)
	if err != nil {
		return nil, err
	}

	reply, err := s.data.AlgorithmClient.ConfirmUploadAlgorithm(ctx, &innterapi.ConfirmUploadAlgorithmRequest{
		SpaceId:     spaceId,
		UserId:      userId,
		AlgorithmId: req.AlgorithmId,
		FileName:    req.FileName,
		Version:     req.Version,
	})

	if err != nil {
		return nil, err
	}

	return &api.ConfirmUploadAlgorithmReply{
		UpdatedAt: reply.UpdatedAt,
	}, nil
}

// 查询算法版本
func (s *AlgorithmService) QueryAlgorithmVersion(ctx context.Context, req *api.QueryAlgorithmVersionRequest) (*api.QueryAlgorithmVersionReply, error) {

	reply, err := s.data.AlgorithmClient.QueryAlgorithmVersion(ctx, &innterapi.QueryAlgorithmVersionRequest{
		AlgorithmId: req.AlgorithmId,
		Version:     req.Version,
	})

	if err != nil {
		return nil, err
	}

	return &api.QueryAlgorithmVersionReply{
		Algorithm:       s.algorithmTransfer(ctx, reply.Algorithm),
		VersionAccesses: s.transferAlgorithmVersionAccess(reply.VersionAccesses),
	}, nil
}

func (s *AlgorithmService) BatchQueryAlgorithm(ctx context.Context, req *api.BatchQueryAlgorithmRequest) (*api.BatchQueryAlgorithmReply, error) {

	reply, err := s.data.AlgorithmClient.BatchQueryAlgorithm(ctx, &innterapi.BatchQueryAlgorithmRequest{
		AlgorithmId: req.AlgorithmId,
	})

	if err != nil {
		return nil, err
	}

	algorithms := []*api.AlgorithmInfo{}
	for _, algorithm := range reply.Algorithms {
		newAlgorithm := s.algorithmInfoTransfer(algorithm)
		algorithms = append(algorithms, newAlgorithm)
	}

	return &api.BatchQueryAlgorithmReply{
		TotalSize:  reply.TotalSize,
		Algorithms: algorithms,
	}, nil
}

func (s *AlgorithmService) algorithmInfoTransfer(algorithm *innterapi.AlgorithmInfo) *api.AlgorithmInfo {

	return &api.AlgorithmInfo{
		AlgorithmId:   algorithm.AlgorithmId,
		AlgorithmName: algorithm.AlgorithmName,
		CreatedAt:     algorithm.CreatedAt,
	}
}

func (s *AlgorithmService) algorithmTransfer(ctx context.Context, algorithm *innterapi.AlgorithmDetail) *api.AlgorithmDetail {

	var ret *api.AlgorithmDetail = &api.AlgorithmDetail{
		AlgorithmId:       algorithm.AlgorithmId,
		AlgorithmVersion:  algorithm.AlgorithmVersion,
		AlgorithmName:     algorithm.AlgorithmName,
		ModelName:         algorithm.ModelName,
		SpaceId:           algorithm.SpaceId,
		UserId:            algorithm.UserId,
		AlgorithmDescript: algorithm.AlgorithmDescript,
		LatestCompressed:  algorithm.LatestCompressed,
		FileStatus:        algorithm.FileStatus,
		IsPrefab:          algorithm.IsPrefab,
		CreatedAt:         algorithm.CreatedAt,
	}

	if algorithm.UserId != "" {
		userReply, err := s.data.UserClient.FindUser(ctx, &innterapi.FindUserRequest{Id: algorithm.UserId})
		if err != nil || userReply.User == nil {
			ret.UserName = ""
		} else {
			ret.UserName = userReply.User.FullName
		}
	}

	if algorithm.SpaceId != "" {
		spaceReply, err := s.data.WorkspaceClient.FindWorkspace(ctx, &innterapi.FindWorkspaceRequest{Id: algorithm.SpaceId})
		if err != nil || spaceReply.Workspace == nil {
			ret.SpaceName = ""
		} else {
			ret.SpaceName = spaceReply.Workspace.Name
		}
	}

	return ret
}

func (h *AlgorithmService) transferAlgorithmVersionAccess(versionAccesses []*innterapi.AlgorithmVersionAccess) []*api.AlgorithmVersionAccess {

	algorithmVersionAccess := []*api.AlgorithmVersionAccess{}

	for _, as := range versionAccesses {
		av := &api.AlgorithmVersionAccess{}
		av.AlgorithmId = as.AlgorithmId
		av.SpaceId = as.SpaceId
		av.Version = as.Version
		algorithmVersionAccess = append(algorithmVersionAccess, av)
	}
	return algorithmVersionAccess
}

func (s *AlgorithmService) myAlgorithmTransfer(ctx context.Context, algorithm *innterapi.AlgorithmDetail, isShared bool) *api.MyAlgorithmDetail {
	algorithmDetail := s.algorithmTransfer(ctx, algorithm)

	return &api.MyAlgorithmDetail{
		AlgorithmDetail: algorithmDetail,
		IsShared:        isShared,
	}
}

func (s *AlgorithmService) getUserIdAndSpaceId(ctx context.Context) (string, string, error) {
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
