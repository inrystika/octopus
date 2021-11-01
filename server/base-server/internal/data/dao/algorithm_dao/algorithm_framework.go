package algorithm_dao

import (
	"context"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"

	stderrors "errors"

	"gorm.io/gorm"
)

// 新增算法类型
func (d *algorithmDao) AddAlgorithmFramework(ctx context.Context, req *model.AlgorithmFramework) error {
	db := d.db.Model(&model.AlgorithmFramework{})
	db = db.Create(req)
	if db.Error != nil {
		return errors.Errorf(db.Error, errors.ErrorDBCreateFailed)
	}

	return nil
}

// 查询算法类型列表
func (d *algorithmDao) ListAlgorithmFramework(ctx context.Context, req *model.AlgorithmFrameworkQuery) ([]*model.AlgorithmFramework, int64, error) {
	db := d.db.Model(&model.AlgorithmFramework{})
	algorithmFrameworks := make([]*model.AlgorithmFramework, 0)

	var totalSize int64
	res := db.Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if req.PageIndex != 0 {
		db = db.Limit(req.PageSize).Offset((req.PageIndex - 1) * req.PageSize)
	}

	res = db.Find(&algorithmFrameworks)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return algorithmFrameworks, totalSize, nil
}

// 查询单个算法类型
func (d *algorithmDao) GetAlgorithmFramework(ctx context.Context, id string) (*model.AlgorithmFramework, error) {
	db := d.db

	nb := &model.AlgorithmFramework{}
	res := db.First(nb, "id = ?", id)

	if res.Error != nil {
		if stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFindEmpty)
		} else {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFirstFailed)
		}
	}
	return nb, nil
}

func (d *algorithmDao) QueryAlgorithmFramework(ctx context.Context, FrameworkDesc string) (*model.AlgorithmFramework, error) {
	db := d.db

	nb := &model.AlgorithmFramework{}
	res := db.First(nb, "desc = ?", FrameworkDesc)

	if res.Error != nil {
		if stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFindEmpty)
		} else {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFirstFailed)
		}
	}
	return nb, nil
}

// 删除算法类型
func (d *algorithmDao) DeleteAlgorithmFramework(ctx context.Context, id string) error {
	db := d.db

	if id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	res := db.Where("id = ? ", id).Delete(&model.AlgorithmFramework{})
	if res.Error != nil {
		return errors.Errorf(nil, errors.ErrorDBDeleteFailed)
	}

	return nil
}

// 修改算法类型描述
func (d *algorithmDao) UpdateAlgorithmFramework(ctx context.Context, req *model.AlgorithmFramework) error {
	db := d.db
	if req.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := db.Where("id = ?", req.Id).Updates(req)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}
