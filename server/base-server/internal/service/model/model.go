package model

import (
	"context"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"

	"server/common/log"
)

type ModelService struct {
	api.UnimplementedModelServer
	conf         *conf.Bootstrap
	log          *log.Helper
	data         *data.Data
	queryHandle  ModelQueryHandle
	addHandle    ModelAddHandle
	deleteHandle ModelDeleteHandle
	shareHandle  ModelShareHandle
}

func NewModelService(conf *conf.Bootstrap, logger log.Logger, data *data.Data, algorithmService api.AlgorithmServer) api.ModelServer {
	return &ModelService{
		conf:         conf,
		log:          log.NewHelper("ModelService", logger),
		data:         data,
		queryHandle:  NewModelQueryHandle(conf, logger, data),
		addHandle:    NewModelAddHandle(conf, logger, data, algorithmService),
		deleteHandle: NewModelDeleteHandle(conf, logger, data),
		shareHandle:  NewModelShareHandle(conf, logger, data),
	}
}

// 查询预置模型列表
func (s *ModelService) ListPreModel(ctx context.Context, req *api.ListPreModelRequest) (*api.ListPreModelReply, error) {
	return s.queryHandle.ListPreModelHandle(ctx, req)
}

// 查询我的模型列表
func (s *ModelService) ListMyModel(ctx context.Context, req *api.ListMyModelRequest) (*api.ListMyModelReply, error) {
	return s.queryHandle.ListMyModelHandle(ctx, req)
}

// 查询公共模型列表
func (s *ModelService) ListCommModel(ctx context.Context, req *api.ListCommModelRequest) (*api.ListCommModelReply, error) {
	return s.queryHandle.ListCommModelHandle(ctx, req)
}

// 查询所有用户模型列表
func (s *ModelService) ListAllUserModel(ctx context.Context, req *api.ListAllUserModelRequest) (*api.ListAllUserModelReply, error) {
	return s.queryHandle.ListAllUserModelHandle(ctx, req)
}

// 查询模型版本列表
func (s *ModelService) ListModelVersion(ctx context.Context, req *api.ListModelVersionRequest) (*api.ListModelVersionReply, error) {
	return s.queryHandle.ListModelVersionHandle(ctx, req)
}

// 查询公共模型版本列表
func (s *ModelService) ListCommModelVersion(ctx context.Context, req *api.ListCommModelVersionRequest) (*api.ListCommModelVersionReply, error) {
	return s.queryHandle.ListCommModelVersionHandle(ctx, req)
}

// 查询模型详情
func (s *ModelService) QueryModel(ctx context.Context, req *api.QueryModelRequest) (*api.QueryModelReply, error) {
	return s.queryHandle.QueryModelHandle(ctx, req)
}

// 查询模型版本详情
func (s *ModelService) QueryModelVersion(ctx context.Context, req *api.QueryModelVersionRequest) (*api.QueryModelVersionReply, error) {
	return s.queryHandle.QueryModelVersionHandle(ctx, req)
}

// 分享模型到公共模型
func (s *ModelService) ShareModelVersion(ctx context.Context, req *api.ShareModelVersionRequest) (*api.ShareModelVersionReply, error) {
	return s.shareHandle.ShareModelVersionHandle(ctx, req)
}

// 取消分享模型版本到公共模型
func (s *ModelService) CloseShareModelVersion(ctx context.Context, req *api.CloseShareModelVersionRequest) (*api.CloseShareModelVersionReply, error) {
	return s.shareHandle.CloseShareModelVersionHandle(ctx, req)
}

// 取消分享模型到公共模型
func (s *ModelService) CloseShareModel(ctx context.Context, req *api.CloseShareModelRequest) (*api.CloseShareModelReply, error) {
	return s.shareHandle.CloseShareModelHandle(ctx, req)
}

// 取消分享模型版本到所有公共模型
func (s *ModelService) AllCloseShareModelVersion(ctx context.Context, req *api.AllCloseShareModelVersionRequest) (*api.AllCloseShareModelVersionReply, error) {
	return s.shareHandle.AllCloseShareModelVersionHandle(ctx, req)
}

// 取消分享模型到所有公共模型
func (s *ModelService) AllCloseShareModel(ctx context.Context, req *api.AllCloseShareModelRequest) (*api.AllCloseShareModelReply, error) {
	return s.shareHandle.AllCloseShareModelHandle(ctx, req)
}

// 新增我的模型
func (s *ModelService) AddMyModel(ctx context.Context, req *api.AddMyModelRequest) (*api.AddMyModelReply, error) {
	return s.addHandle.AddMyModelHandle(ctx, req)
}

// 新增预置模型
func (s *ModelService) AddPreModel(ctx context.Context, req *api.AddPreModelRequest) (*api.AddPreModelReply, error) {
	return s.addHandle.AddPreModelHandle(ctx, req)
}

// 新增预置模型版本
func (s *ModelService) AddPreModelVersion(ctx context.Context, req *api.AddPreModelVersionRequest) (*api.AddPreModelVersionReply, error) {
	return s.addHandle.AddPreModelVersionHandle(ctx, req)
}

// 上传预置模型版本
func (s *ModelService) UploadPreModelVersion(ctx context.Context, req *api.UploadPreModelVersionRequest) (*api.UploadPreModelVersionReply, error) {
	return s.addHandle.UploadPreModelVersionHandle(ctx, req)
}

// 上传预置模型版本确认
func (s *ModelService) ConfirmUploadPreModelVersion(ctx context.Context, req *api.ConfirmUploadPreModelVersionRequest) (*api.ConfirmUploadPreModelVersionReply, error) {
	return s.addHandle.ConfirmUploadPreModelVersionHandle(ctx, req)
}

// 删除我的模型版本
func (s *ModelService) DeleteMyModelVersion(ctx context.Context, req *api.DeleteMyModelVersionRequest) (*api.DeleteMyModelVersionReply, error) {
	return s.deleteHandle.DeleteMyModelVersionHandle(ctx, req)
}

// 删除我的模型
func (s *ModelService) DeleteMyModel(ctx context.Context, req *api.DeleteMyModelRequest) (*api.DeleteMyModelReply, error) {
	return s.deleteHandle.DeleteMyModelHandle(ctx, req)
}

// 删除我的模型版本
func (s *ModelService) DeletePreModelVersion(ctx context.Context, req *api.DeletePreModelVersionRequest) (*api.DeletePreModelVersionReply, error) {
	return s.deleteHandle.DeletePreModelVersionHandle(ctx, req)
}

// 删除我的模型
func (s *ModelService) DeletePreModel(ctx context.Context, req *api.DeletePreModelRequest) (*api.DeletePreModelReply, error) {
	return s.deleteHandle.DeletePreModelHandle(ctx, req)
}

// 下载模型版本
func (s *ModelService) DownloadModelVersion(ctx context.Context, req *api.DownloadModelVersionRequest) (*api.DownloadModelVersionReply, error) {
	return s.shareHandle.DownloadModelVersionHandle(ctx, req)
}

// 预览模型版本
func (s *ModelService) ListModelVersionFile(ctx context.Context, req *api.ListModelVersionFileRequest) (*api.ListModelVersionFileReply, error) {
	return s.shareHandle.ListModelVersionFileHandle(ctx, req)
}
