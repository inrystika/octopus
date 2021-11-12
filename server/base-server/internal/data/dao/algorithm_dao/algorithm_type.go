package algorithm_dao

import (
	"context"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"

	stderrors "errors"

	"gorm.io/gorm"
)

// 新增算法类型
func (d *algorithmDao) AddAlgorithmType(ctx context.Context, req *model.AlgorithmType) error {
	db := d.db.Model(&model.AlgorithmType{})
	db = db.Create(req)
	if db.Error != nil {
		return errors.Errorf(db.Error, errors.ErrorDBCreateFailed)
	}

	return nil
}

// 查询算法类型列表
func (d *algorithmDao) ListAlgorithmType(ctx context.Context, req *model.AlgorithmTypeQuery) ([]*model.AlgorithmType, int64, error) {
	db := d.db.Model(&model.AlgorithmType{})
	algorithmTypes := make([]*model.AlgorithmType, 0)

	var totalSize int64
	res := db.Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if req.PageIndex != 0 {
		db = db.Limit(req.PageSize).Offset((req.PageIndex - 1) * req.PageSize)
	}

	res = db.Find(&algorithmTypes)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return algorithmTypes, totalSize, nil
}

// 查询单个算法类型
func (d *algorithmDao) GetAlgorithmType(ctx context.Context, id string) (*model.AlgorithmType, error) {
	db := d.db

	nb := &model.AlgorithmType{}
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

func (d *algorithmDao) QueryAlgorithmType(ctx context.Context, typeDesc string) (*model.AlgorithmType, error) {
	db := d.db

	nb := &model.AlgorithmType{}
	res := db.First(nb, "`desc` = ?", typeDesc)

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
func (d *algorithmDao) DeleteAlgorithmType(ctx context.Context, id string) error {
	db := d.db

	if id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	res := db.Where("id = ? ", id).Delete(&model.AlgorithmType{})
	if res.Error != nil {
		return errors.Errorf(nil, errors.ErrorDBDeleteFailed)
	}

	return nil
}

// 修改算法类型描述
func (d *algorithmDao) UpdateAlgorithmType(ctx context.Context, req *model.AlgorithmType) error {
	db := d.db
	if req.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := db.Model(&model.AlgorithmType{}).Where("id = ? ", req.Id).Updates(map[string]interface{}{
		"desc":        req.Desc,
		"refer_times": req.ReferTimes,
	})

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}
