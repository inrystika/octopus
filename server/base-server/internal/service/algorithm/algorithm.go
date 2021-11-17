package algorithm

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"

	"server/common/log"
)

type AlgorithmService struct {
	api.UnimplementedAlgorithmServer
	conf   *conf.Bootstrap
	log    *log.Helper
	data   *data.Data
	handle AlgorithmHandle
}

func NewAlgorithmService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.AlgorithmServer {
	return &AlgorithmService{
		conf:   conf,
		log:    log.NewHelper("AlgorithmService", logger),
		data:   data,
		handle: NewAlgorithmHandle(conf, logger, data),
	}
}

// 新增算法类型
func (s *AlgorithmService) AddAlgorithmType(ctx context.Context, req *api.AddAlgorithmTypeRequest) (*api.AddAlgorithmTypeReply, error) {
	return s.handle.AddAlgorithmType(ctx, req)
}

// 查询算法类型列表
func (s *AlgorithmService) ListAlgorithmType(ctx context.Context, req *api.ListAlgorithmTypeRequest) (*api.ListAlgorithmTypeReply, error) {
	return s.handle.ListAlgorithmType(ctx, req)
}

// 查询单个算法类型
func (s *AlgorithmService) GetAlgorithmType(ctx context.Context, req *api.GetAlgorithmTypeRequest) (*api.GetAlgorithmTypeReply, error) {
	return s.handle.GetAlgorithmType(ctx, req)
}

// 删除算法类型
func (s *AlgorithmService) DeleteAlgorithmType(ctx context.Context, req *api.DeleteAlgorithmTypeRequest) (*api.DeleteAlgorithmTypeReply, error) {
	return s.handle.DeleteAlgorithmType(ctx, req)
}

// 修改算法类型描述
func (s *AlgorithmService) UpdateAlgorithmType(ctx context.Context, req *api.UpdateAlgorithmTypeRequest) (*api.UpdateAlgorithmTypeReply, error) {
	return s.handle.UpdateAlgorithmType(ctx, req)
}

// 新增算法框架
func (s *AlgorithmService) AddAlgorithmFramework(ctx context.Context, req *api.AddAlgorithmFrameworkRequest) (*api.AddAlgorithmFrameworkReply, error) {
	return s.handle.AddAlgorithmFramework(ctx, req)
}

// 查询算法框架列表
func (s *AlgorithmService) ListAlgorithmFramework(ctx context.Context, req *api.ListAlgorithmFrameworkRequest) (*api.ListAlgorithmFrameworkReply, error) {
	return s.handle.ListAlgorithmFramework(ctx, req)
}

// 查询单个算法框架
func (s *AlgorithmService) GetAlgorithmFramework(ctx context.Context, req *api.GetAlgorithmFrameworkRequest) (*api.GetAlgorithmFrameworkReply, error) {
	return s.handle.GetAlgorithmFramework(ctx, req)
}

// 删除算法框架
func (s *AlgorithmService) DeleteAlgorithmFramework(ctx context.Context, req *api.DeleteAlgorithmFrameworkRequest) (*api.DeleteAlgorithmFrameworkReply, error) {
	return s.handle.DeleteAlgorithmFramework(ctx, req)
}

// 修改算法框架描述
func (s *AlgorithmService) UpdateAlgorithmFramework(ctx context.Context, req *api.UpdateAlgorithmFrameworkRequest) (*api.UpdateAlgorithmFrameworkReply, error) {
	return s.handle.UpdateAlgorithmFramework(ctx, req)
}

// 查询预置算法列表
func (s *AlgorithmService) ListPreAlgorithm(ctx context.Context, req *api.ListPreAlgorithmRequest) (*api.ListPreAlgorithmReply, error) {
	return s.handle.ListPreAlgorithmHandle(ctx, req)
}

// 查询我的算法列表
func (s *AlgorithmService) ListMyAlgorithm(ctx context.Context, req *api.ListMyAlgorithmRequest) (*api.ListMyAlgorithmReply, error) {
	return s.handle.ListMyAlgorithmHandle(ctx, req)
}

// 查询公共算法列表
func (s *AlgorithmService) ListCommAlgorithm(ctx context.Context, req *api.ListCommAlgorithmRequest) (*api.ListCommAlgorithmReply, error) {
	return s.handle.ListCommAlgorithmHandle(ctx, req)
}

// 批量查询算法
func (s *AlgorithmService) BatchQueryAlgorithm(ctx context.Context, req *api.BatchQueryAlgorithmRequest) (*api.BatchQueryAlgorithmReply, error) {
	return s.handle.BatchQueryAlgorithmHandle(ctx, req)
}

// 查询算法版本列表
func (s *AlgorithmService) ListAlgorithmVersion(ctx context.Context, req *api.ListAlgorithmVersionRequest) (*api.ListAlgorithmVersionReply, error) {
	return s.handle.ListAlgorithmVersionHandle(ctx, req)
}

// 查询公共算法版本列表
func (s *AlgorithmService) ListCommAlgorithmVersion(ctx context.Context, req *api.ListCommAlgorithmVersionRequest) (*api.ListCommAlgorithmVersionReply, error) {
	return s.handle.ListCommAlgorithmVersionHandle(ctx, req)
}

// 查询所有用户算法列表
func (s *AlgorithmService) ListAllUserAlgorithm(ctx context.Context, req *api.ListAllUserAlgorithmRequest) (*api.ListAllUserAlgorithmReply, error) {
	return s.handle.ListAllUserAlgorithmHandle(ctx, req)
}

// 查询算法版本详情
func (s *AlgorithmService) QueryAlgorithmVersion(ctx context.Context, req *api.QueryAlgorithmVersionRequest) (*api.QueryAlgorithmVersionReply, error) {
	return s.handle.QueryAlgorithmVersionHandle(ctx, req)
}

// 分享算法到公共算法
func (s *AlgorithmService) ShareAlgorithmVersion(ctx context.Context, req *api.ShareAlgorithmVersionRequest) (*api.ShareAlgorithmVersionReply, error) {
	return s.handle.ShareAlgorithmVersionHandle(ctx, req)
}

// 取消分享算法版本到公共算法
func (s *AlgorithmService) CloseShareAlgorithmVersion(ctx context.Context, req *api.CloseShareAlgorithmVersionRequest) (*api.CloseShareAlgorithmVersionReply, error) {
	return s.handle.CloseShareAlgorithmVersionHandle(ctx, req)
}

// 取消分享算法到公共算法
func (s *AlgorithmService) CloseShareAlgorithm(ctx context.Context, req *api.CloseShareAlgorithmRequest) (*api.CloseShareAlgorithmReply, error) {
	return s.handle.CloseShareAlgorithmHandle(ctx, req)
}

// 取消分享算法版本到所有公共算法
func (s *AlgorithmService) AllCloseShareAlgorithmVersion(ctx context.Context, req *api.AllCloseShareAlgorithmVersionRequest) (*api.AllCloseShareAlgorithmVersionReply, error) {
	return s.handle.AllCloseShareAlgorithmVersionHandle(ctx, req)
}

// 取消分享算法到所有公共算法
func (s *AlgorithmService) AllCloseShareAlgorithm(ctx context.Context, req *api.AllCloseShareAlgorithmRequest) (*api.AllCloseShareAlgorithmReply, error) {
	return s.handle.AllCloseShareAlgorithmHandle(ctx, req)
}

// 新增算法
func (s *AlgorithmService) AddAlgorithm(ctx context.Context, req *api.AddAlgorithmRequest) (*api.AddAlgorithmReply, error) {
	return s.handle.AddAlgorithmHandle(ctx, req)
}

// 上传算法文件
func (s *AlgorithmService) UploadAlgorithm(ctx context.Context, req *api.UploadAlgorithmRequest) (*api.UploadAlgorithmReply, error) {
	return s.handle.UploadAlgorithmHandle(ctx, req)
}

// 上传算法确认
func (s *AlgorithmService) ConfirmUploadAlgorithm(ctx context.Context, req *api.ConfirmUploadAlgorithmRequest) (*api.ConfirmUploadAlgorithmReply, error) {
	return s.handle.ConfirmUploadAlgorithmHandle(ctx, req)
}

// 修改算法
func (s *AlgorithmService) UpdateAlgorithm(ctx context.Context, req *api.UpdateAlgorithmRequest) (*api.UpdateAlgorithmReply, error) {
	return s.handle.UpdateAlgorithmHandle(ctx, req)
}

// 新增我的算法版本
func (s *AlgorithmService) AddMyAlgorithmVersion(ctx context.Context, req *api.AddMyAlgorithmVersionRequest) (*api.AddMyAlgorithmVersionReply, error) {
	return s.handle.AddMyAlgorithmVersionHandle(ctx, req)
}

// 删除我的算法版本
func (s *AlgorithmService) DeleteMyAlgorithmVersion(ctx context.Context, req *api.DeleteMyAlgorithmVersionRequest) (*api.DeleteMyAlgorithmVersionReply, error) {
	return s.handle.DeleteMyAlgorithmVersionHandle(ctx, req)
}

// 删除我的算法
func (s *AlgorithmService) DeleteMyAlgorithm(ctx context.Context, req *api.DeleteMyAlgorithmRequest) (*api.DeleteMyAlgorithmReply, error) {
	return s.handle.DeleteMyAlgorithmHandle(ctx, req)
}

// 新增预置算法版本
func (s *AlgorithmService) AddPreAlgorithmVersion(ctx context.Context, req *api.AddPreAlgorithmVersionRequest) (*api.AddPreAlgorithmVersionReply, error) {
	return s.handle.AddPreAlgorithmVersionHandle(ctx, req)
}

// 删除预置算法版本
func (s *AlgorithmService) DeletePreAlgorithmVersion(ctx context.Context, req *api.DeletePreAlgorithmVersionRequest) (*api.DeletePreAlgorithmVersionReply, error) {
	return s.handle.DeletePreAlgorithmVersionHandle(ctx, req)
}

// 删除预置算法
func (s *AlgorithmService) DeletePreAlgorithm(ctx context.Context, req *api.DeletePreAlgorithmRequest) (*api.DeletePreAlgorithmReply, error) {
	return s.handle.DeletePreAlgorithmHandle(ctx, req)
}

// 修改算法版本
func (s *AlgorithmService) UpdateAlgorithmVersion(ctx context.Context, req *api.UpdateAlgorithmVersionRequest) (*api.UpdateAlgorithmVersionReply, error) {
	return s.handle.UpdateAlgorithmVersionHandle(ctx, req)
}

// 压缩算法版本包
func (s *AlgorithmService) DownloadAlgorithmVersionCompress(ctx context.Context,
	req *api.DownloadAlgorithmVersionCompressRequest) (*api.DownloadAlgorithmVersionCompressReply, error) {
	return s.handle.DownloadAlgorithmVersionCompressHandle(ctx, req)
}

// 下载算法版本
func (s *AlgorithmService) DownloadAlgorithmVersion(ctx context.Context,
	req *api.DownloadAlgorithmVersionRequest) (*api.DownloadAlgorithmVersionReply, error) {
	return s.handle.DownloadAlgorithmVersionHandle(ctx, req)
}

// 复制算法版本
func (s *AlgorithmService) CopyAlgorithmVersion(ctx context.Context, req *api.CopyAlgorithmVersionRequest) (*api.CopyAlgorithmVersionReply, error) {
	return s.handle.CopyAlgorithmVersionHandle(ctx, req)
}
