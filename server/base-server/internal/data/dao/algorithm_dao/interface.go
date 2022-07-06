package algorithm_dao

import (
	"context"
	"server/base-server/internal/data/dao/model"

	"server/common/log"

	"gorm.io/gorm"
)

type AlgorithmDao interface {
	// 算法列表查询
	ListAlgorithm(ctx context.Context, req *model.AlgorithmList) (int64, []*model.Algorithm, error)
	// 查询算法
	QueryAlgorithm(ctx context.Context, req *model.AlgorithmQuery) (*model.Algorithm, error)
	// 根据算法信息查询算法
	QueryAlgorithmByInfo(ctx context.Context, req *model.AlgorithmQueryByInfo) (*model.Algorithm, error)
	// 批量查询算法
	BatchQueryAlgorithm(ctx context.Context, req *model.AlgorithmBatchQuery) ([]*model.Algorithm, error)
	// 添加算法
	AddAlgorithm(ctx context.Context, req *model.Algorithm) (*model.Algorithm, error)
	// 删除算法
	DeleteAlgorithm(ctx context.Context, req *model.AlgorithmDelete) error
	// 修改算法
	UpdateAlgorithm(ctx context.Context, req *model.Algorithm) error

	// 算法版本列表查询
	ListAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersionList) (int64, []*model.AlgorithmVersion, error)
	// 查询算法版本
	QueryAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersionQuery) (*model.AlgorithmVersion, error)
	// 批量查询算法版本
	BatchQueryAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersionBatchQuery) ([]*model.AlgorithmVersion, error)
	// 添加算法版本
	AddAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersion) (*model.AlgorithmVersion, error)
	// 删除算法版本
	DeleteAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersionDelete) error
	// 批量删除算法版本
	BatchDeleteAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersionBatchDelete) error
	// 修改算法版本
	UpdateAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersion) error

	// 公共算法列表查询
	ListAlgorithmAccess(ctx context.Context, req *model.AlgorithmAccessList) (int64, []*model.AlgorithmAccess, error)
	// 查询公共算法
	QueryAlgorithmAccess(ctx context.Context, req *model.AlgorithmAccessQuery) (*model.AlgorithmAccess, error)
	// 批量查询公共算法
	BatchQueryAlgorithmAccess(ctx context.Context, req *model.AlgorithmAccessBatchQuery) ([]*model.AlgorithmAccess, error)
	// 批量查询公共算法
	BatchQueryAlgorithmAccessById(ctx context.Context, req *model.AlgorithmAccessBatchQueryById) ([]*model.AlgorithmAccess, error)
	// 添加公共算法
	AddAlgorithmAccess(ctx context.Context, req *model.AlgorithmAccess) (*model.AlgorithmAccess, error)
	// 删除公共算法
	DeleteAlgorithmAccess(ctx context.Context, req *model.AlgorithmAccessDelete) error
	// 修改公共算法
	UpdateAlgorithmAccess(ctx context.Context, req *model.AlgorithmAccess) error

	// 公共算法版本列表查询
	ListAlgorithmAccessVersion(ctx context.Context, req *model.AlgorithmAccessVersionList) (int64, []*model.AlgorithmAccessVersion, error)
	// 查询公共算法版本
	QueryAlgorithmAccessVersion(ctx context.Context, req *model.AlgorithmAccessVersionQuery) (*model.AlgorithmAccessVersion, error)
	// 批量查询公共算法版本
	BatchQueryAlgorithmAccessVersion(ctx context.Context, req *model.AlgorithmAccessVersionBatchQuery) ([]*model.AlgorithmAccessVersion, error)
	// 添加公共算法版本
	AddAlgorithmAccessVersion(ctx context.Context, req *model.AlgorithmAccessVersion) (*model.AlgorithmAccessVersion, error)
	// 修改公共算法版本
	UpdateAlgorithmAccessVersion(ctx context.Context, req *model.AlgorithmAccessVersion) error
	// 删除公共算法版本
	DeleteAlgorithmAccessVersion(ctx context.Context, req *model.AlgorithmAccessVersionDelete) error
	// 批量删除公共算法所有版本
	BatchDeleteAlgorithmAccessVersion(ctx context.Context, req *model.AlgorithmAccessVersionBatchDelete) error
}

type algorithmDao struct {
	log *log.Helper
	db  *gorm.DB
}

func NewAlgorithmDao(db *gorm.DB, logger log.Logger) *algorithmDao {
	return &algorithmDao{
		log: log.NewHelper("AlgorithmDao", logger),
		db:  db,
	}
}
