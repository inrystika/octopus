package service

import (
	"context"
	api "server/admin-server/api/v1"
	"server/admin-server/internal/conf"
	"server/admin-server/internal/data"
	innerapi "server/base-server/api/v1"
	innterapi "server/base-server/api/v1"
	"server/common/log"

	"github.com/jinzhu/copier"
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

// 新增算法用途
func (s *AlgorithmService) AddAlgorithmApply(ctx context.Context, req *api.AddAlgorithmApplyRequest) (*api.AddAlgorithmApplyReply, error) {
	innerReq := &innerapi.AddLableRequest{
		RelegationType: int32(innerapi.Relegation_LABLE_RELEGATION_ALGORITHM),
		LableType:      int32(innerapi.Type_LABLE_TYPE_ALGORITHM_APPLY),
		LableDesc:      req.LableDesc,
	}

	innerReply, err := s.data.LableClient.AddLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.AddAlgorithmApplyReply{
		AlgorithmApply: &api.AlgorithmApply{
			Id:         innerReply.Lable.Id,
			LableDesc:  innerReply.Lable.LableDesc,
			SourceType: innerReply.Lable.SourceType,
		},
	}, nil
}

// 查询算法用途列表
func (s *AlgorithmService) ListAlgorithmApply(ctx context.Context, req *api.ListAlgorithmApplyRequest) (*api.ListAlgorithmApplyReply, error) {
	innerReq := &innerapi.ListLableRequest{
		RelegationType: int32(innerapi.Relegation_LABLE_RELEGATION_ALGORITHM),
		LableType:      int32(innerapi.Type_LABLE_TYPE_ALGORITHM_APPLY),
		PageIndex:      req.PageIndex,
		PageSize:       req.PageSize,
	}

	innerReply, err := s.data.LableClient.ListLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListAlgorithmApplyReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// 删除算法用途
func (s *AlgorithmService) DeleteAlgorithmApply(ctx context.Context, req *api.DeleteAlgorithmApplyRequest) (*api.DeleteAlgorithmApplyReply, error) {
	innerReq := &innerapi.DeleteLableRequest{
		Id: req.Id,
	}

	innerReply, err := s.data.LableClient.DeleteLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.DeleteAlgorithmApplyReply{
		DeletedAt: innerReply.DeletedAt,
	}, nil
}

// 修改算法用途描述
func (s *AlgorithmService) UpdateAlgorithmApply(ctx context.Context, req *api.UpdateAlgorithmApplyRequest) (*api.UpdateAlgorithmApplyReply, error) {
	innerReq := &innerapi.UpdateLableRequest{
		Id:        req.Id,
		LableDesc: req.LableDesc,
	}

	innerReply, err := s.data.LableClient.UpdateLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.UpdateAlgorithmApplyReply{
		UpdatedAt: innerReply.UpdatedAt,
	}, nil
}

// 新增算法框架
func (s *AlgorithmService) AddAlgorithmFramework(ctx context.Context, req *api.AddAlgorithmFrameworkRequest) (*api.AddAlgorithmFrameworkReply, error) {
	innerReq := &innerapi.AddLableRequest{
		RelegationType: int32(innerapi.Relegation_LABLE_RELEGATION_ALGORITHM),
		LableType:      int32(innerapi.Type_LABLE_TYPE_ALGORITHM_FRAMEWORK),
		LableDesc:      req.LableDesc,
	}

	innerReply, err := s.data.LableClient.AddLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.AddAlgorithmFrameworkReply{
		AlgorithmFramework: &api.AlgorithmFramework{
			Id:         innerReply.Lable.Id,
			LableDesc:  innerReply.Lable.LableDesc,
			SourceType: innerReply.Lable.SourceType,
		},
	}, nil
}

// 查询算法类型框架
func (s *AlgorithmService) ListAlgorithmFramework(ctx context.Context, req *api.ListAlgorithmFrameworkRequest) (*api.ListAlgorithmFrameworkReply, error) {
	innerReq := &innerapi.ListLableRequest{
		RelegationType: int32(innerapi.Relegation_LABLE_RELEGATION_ALGORITHM),
		LableType:      int32(innerapi.Type_LABLE_TYPE_ALGORITHM_FRAMEWORK),
		PageIndex:      req.PageIndex,
		PageSize:       req.PageSize,
	}

	innerReply, err := s.data.LableClient.ListLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	reply := &api.ListAlgorithmFrameworkReply{}
	err = copier.Copy(reply, innerReply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// 删除算法框架
func (s *AlgorithmService) DeleteAlgorithmFramework(ctx context.Context, req *api.DeleteAlgorithmFrameworkRequest) (*api.DeleteAlgorithmFrameworkReply, error) {
	innerReq := &innerapi.DeleteLableRequest{
		Id: req.Id,
	}

	innerReply, err := s.data.LableClient.DeleteLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.DeleteAlgorithmFrameworkReply{
		DeletedAt: innerReply.DeletedAt,
	}, nil
}

// 修改算法类型框架
func (s *AlgorithmService) UpdateAlgorithmFramework(ctx context.Context, req *api.UpdateAlgorithmFrameworkRequest) (*api.UpdateAlgorithmFrameworkReply, error) {
	innerReq := &innerapi.UpdateLableRequest{
		Id:        req.Id,
		LableDesc: req.LableDesc,
	}

	innerReply, err := s.data.LableClient.UpdateLable(ctx, innerReq)
	if err != nil {
		return nil, err
	}

	return &api.UpdateAlgorithmFrameworkReply{
		UpdatedAt: innerReply.UpdatedAt,
	}, nil
}

// 查询预置算法列表
func (s *AlgorithmService) ListPreAlgorithm(ctx context.Context, req *api.ListPreAlgorithmRequest) (*api.ListPreAlgorithmReply, error) {
	reply, err := s.data.AlgorithmClient.ListPreAlgorithm(ctx, &innterapi.ListPreAlgorithmRequest{
		PageIndex:        req.PageIndex,
		PageSize:         req.PageSize,
		SearchKey:        req.SearchKey,
		SortBy:           req.SortBy,
		OrderBy:          req.OrderBy,
		AlgorithmVersion: req.AlgorithmVersion,
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

// 查询所有用户算法列表
func (s *AlgorithmService) ListAllAlgorithm(ctx context.Context, req *api.ListAllAlgorithmRequest) (*api.ListAllAlgorithmReply, error) {

	reply, err := s.data.AlgorithmClient.ListAllUserAlgorithm(ctx, &innterapi.ListAllUserAlgorithmRequest{
		PageIndex:        req.PageIndex,
		PageSize:         req.PageSize,
		SearchKey:        req.SearchKey,
		SortBy:           req.SortBy,
		OrderBy:          req.OrderBy,
		AlgorithmVersion: req.AlgorithmVersion,
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

	return &api.ListAllAlgorithmReply{
		TotalSize:  reply.TotalSize,
		Algorithms: algorithms,
	}, nil
}

// 查询算法版本列表
func (s *AlgorithmService) ListAlgorithmVersion(ctx context.Context, req *api.ListAlgorithmVersionRequest) (*api.ListAlgorithmVersionReply, error) {

	reply, err := s.data.AlgorithmClient.ListAlgorithmVersion(ctx, &innterapi.ListAlgorithmVersionRequest{
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
		algorithmDetail := s.algorithmTransfer(ctx, algorithm)
		algorithms = append(algorithms, algorithmDetail)
	}

	return &api.ListAlgorithmVersionReply{
		TotalSize:  reply.TotalSize,
		Algorithms: algorithms,
	}, nil
}

// 新增预置算法版本
func (s *AlgorithmService) AddPreAlgorithmVersion(ctx context.Context, req *api.AddPreAlgorithmVersionRequest) (*api.AddPreAlgorithmVersionReply, error) {

	reply, err := s.data.AlgorithmClient.AddPreAlgorithmVersion(ctx, &innterapi.AddPreAlgorithmVersionRequest{
		SpaceId:           "",
		UserId:            "",
		AlgorithmId:       req.AlgorithmId,
		AlgorithmDescript: req.AlgorithmDescript,
	})

	if err != nil {
		return nil, err
	}

	return &api.AddPreAlgorithmVersionReply{
		AlgorithmId: reply.AlgorithmId,
		Version:     reply.Version,
		CreatedAt:   reply.CreatedAt,
	}, nil
}

// 删除预置算法
func (s *AlgorithmService) DeletePreAlgorithm(ctx context.Context, req *api.DeletePreAlgorithmRequest) (*api.DeletePreAlgorithmReply, error) {

	reply, err := s.data.AlgorithmClient.DeletePreAlgorithm(ctx, &innterapi.DeletePreAlgorithmRequest{
		AlgorithmId: req.AlgorithmId,
	})
	if err != nil {
		return nil, err
	}

	return &api.DeletePreAlgorithmReply{
		DeletedAt: reply.DeletedAt,
	}, nil
}

// 删除我的算法版本
func (s *AlgorithmService) DeletePreAlgorithmVersion(ctx context.Context, req *api.DeletePreAlgorithmVersionRequest) (*api.DeletePreAlgorithmVersionReply, error) {

	reply, err := s.data.AlgorithmClient.DeleteMyAlgorithmVersion(ctx, &innterapi.DeleteMyAlgorithmVersionRequest{
		SpaceId:     "",
		UserId:      "",
		AlgorithmId: req.AlgorithmId,
		Version:     req.Version,
	})
	if err != nil {
		return nil, err
	}

	return &api.DeletePreAlgorithmVersionReply{
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

// 新增预置算法
func (s *AlgorithmService) AddPreAlgorithm(ctx context.Context, req *api.AddAlgorithmRequest) (*api.AddAlgorithmReply, error) {

	reply, err := s.data.AlgorithmClient.AddAlgorithm(ctx, &innterapi.AddAlgorithmRequest{
		IsPrefab:          true,
		IsEmpty:           req.IsEmpty,
		AlgorithmName:     req.AlgorithmName,
		ModelName:         req.ModelName,
		AlgorithmDescript: req.AlgorithmDescript,
		ApplyId:           req.ApplyId,
		FrameworkId:       req.FrameworkId,
	})
	if err != nil {
		return nil, err
	}

	return &api.AddAlgorithmReply{
		AlgorithmId: reply.AlgorithmId,
		Version:     reply.Version,
		CreatedAt:   reply.CreatedAt,
	}, nil
}

// 上传算法
func (s *AlgorithmService) UploadPreAlgorithm(ctx context.Context, req *api.UploadPreAlgorithmRequest) (*api.UploadPreAlgorithmReply, error) {

	reply, err := s.data.AlgorithmClient.UploadAlgorithm(ctx, &innterapi.UploadAlgorithmRequest{
		SpaceId:     "",
		UserId:      "",
		AlgorithmId: req.AlgorithmId,
		Version:     req.Version,
		FileName:    req.FileName,
		Domain:      req.Domain,
	})
	if err != nil {
		return nil, err
	}

	return &api.UploadPreAlgorithmReply{
		UploadUrl: reply.UploadUrl,
	}, nil
}

// 上传预置算法确认
func (s *AlgorithmService) ConfirmUploadPreAlgorithm(ctx context.Context, req *api.ConfirmUploadPreAlgorithmRequest) (*api.ConfirmUploadPreAlgorithmReply, error) {

	reply, err := s.data.AlgorithmClient.ConfirmUploadAlgorithm(ctx, &innterapi.ConfirmUploadAlgorithmRequest{
		SpaceId:     "",
		UserId:      "",
		AlgorithmId: req.AlgorithmId,
		FileName:    req.FileName,
		Version:     req.Version,
	})

	if err != nil {
		return nil, err
	}

	return &api.ConfirmUploadPreAlgorithmReply{
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
		ApplyId:           algorithm.ApplyId,
		ApplyName:         algorithm.ApplyName,
		FrameworkId:       algorithm.FrameworkId,
		FrameworkName:     algorithm.FrameworkName,
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
