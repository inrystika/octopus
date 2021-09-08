package algorithm_dao

import (
	"context"
	"fmt"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"
	"time"

	"gorm.io/gorm"
)

// 公共算法列表查询
func (d *algorithmDao) ListAlgorithmAccess(ctx context.Context, req *model.AlgorithmAccessList) (int64, []*model.AlgorithmAccess, error) {
	db := d.db.Model(&model.AlgorithmAccess{})
	accessAlgorithms := make([]*model.AlgorithmAccess, 0)

	params := make([]interface{}, 0)
	querySql := "1 = 1"
	if req.SpaceId != "" {
		querySql += " and space_id = ? "
		params = append(params, req.SpaceId)
	}
	if req.AlgorithmId != "" {
		querySql += " and algorithm_id = ? "
		params = append(params, req.AlgorithmId)
	}

	if req.CreatedAtGte != 0 {
		querySql += " and created_at >= ? "
		params = append(params, time.Unix(req.CreatedAtGte, 0))
	}

	if req.CreatedAtLt != 0 {
		querySql += " and created_at < ? "
		params = append(params, time.Unix(req.CreatedAtLt, 0))
	}

	// 模糊搜索
	if req.SearchKey != "" {
		querySql += " and (algorithm_name like ? "
		params = append(params, "%"+req.SearchKey+"%")
		querySql += " or algorithm_descript like ? )"
		params = append(params, "%"+req.SearchKey+"%")
	}

	if req.NameLike != "" {
		querySql += " and algorithm_name like ? "
		params = append(params, "%"+req.NameLike+"%")
	}

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Count(&totalSize)

	if res.Error != nil {
		err := errors.Errorf(res.Error, errors.ErrorDBCountFailed)
		d.log.Errorw(ctx, err)
		return 0, nil, err
	}

	// limit语句拼接
	if req.PageSize > 0 && req.PageIndex > 0 {
		db = db.Limit(req.PageSize).
			Offset((req.PageIndex - 1) * req.PageSize)
	}

	// orderby语句拼接
	if req.SpaceIdOrder {
		orderSql := fmt.Sprintf("space_id %s", req.SpaceIdSort)
		db = db.Order(orderSql)
	}
	if req.AlgorithmIdOrder {
		orderSql := fmt.Sprintf("algorithm_id %s", req.AlgorithmIdSort)
		db = db.Order(orderSql)
	}
	if req.CreatedAtOrder {
		orderSql := fmt.Sprintf("created_at %s", req.CreatedAtSort)
		db = db.Order(orderSql)
	}

	if req.SortBy != "" {
		db = req.Order(db)
	}

	db = db.Find(&accessAlgorithms)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindFailed)
		d.log.Errorw(ctx, err)
		return 0, nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		d.log.Errorw(ctx, err)
		return 0, nil, err
	}

	d.log.Infof(ctx, "successfully ListAlgorithmAccess, totalSize=%d", totalSize)
	return totalSize, accessAlgorithms, nil
}

// 查询公共算法
func (d *algorithmDao) QueryAlgorithmAccess(ctx context.Context, req *model.AlgorithmAccessQuery) (*model.AlgorithmAccess, error) {
	db := d.db.Model(&model.AlgorithmAccess{})
	oneAlgorithmAccess := model.AlgorithmAccess{}

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	querySql += " and space_id = ? "
	params = append(params, req.SpaceId)
	querySql += " and algorithm_id = ? "
	params = append(params, req.AlgorithmId)

	// 	查询
	db = db.Where(querySql, params...).First(&oneAlgorithmAccess)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFirstFailed)
		d.log.Errorw(ctx, err)
		return nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		d.log.Errorw(ctx, err)
		return nil, err
	}

	d.log.Infof(ctx, "successfully QueryAlgorithmAccess, spaceId=%s|algorithmId=%s", req.SpaceId, req.AlgorithmId)
	return &oneAlgorithmAccess, nil
}

// 批量查询公共算法
func (d *algorithmDao) BatchQueryAlgorithmAccess(ctx context.Context, req *model.AlgorithmAccessBatchQuery) ([]*model.AlgorithmAccess, error) {
	db := d.db.Model(&model.AlgorithmAccess{})
	algorithmAccessList := make([]*model.AlgorithmAccess, 0)

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	querySql += " and (space_id, algorithm_id) in ? "
	for _, query := range req.List {
		oneParams := fmt.Sprintf("(%s, %s)", query.SpaceId, query.AlgorithmId)
		params = append(params, oneParams)
	}

	// 	查询
	db = db.Where(querySql, params...).Find(&algorithmAccessList)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindFailed)
		d.log.Errorw(ctx, err)
		return nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		d.log.Errorw(ctx, err)
		return nil, err
	}

	d.log.Infof(ctx, "successfully BatchQueryAlgorithmAccess, spaceId,algorithm=%#v", req.List)
	return algorithmAccessList, nil
}

// 批量查询公共算法
func (d *algorithmDao) BatchQueryAlgorithmAccessById(ctx context.Context, req *model.AlgorithmAccessBatchQueryById) ([]*model.AlgorithmAccess, error) {
	db := d.db.Model(&model.AlgorithmAccess{})
	algorithmAccessList := make([]*model.AlgorithmAccess, 0)

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	querySql += " and id in ? "
	params = append(params, req.AlgorithmAccessIdList)

	// 	查询
	db = db.Where(querySql, params...).Find(&algorithmAccessList)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindFailed)
		d.log.Errorw(ctx, err)
		return nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		d.log.Errorw(ctx, err)
		return nil, err
	}

	d.log.Infof(ctx, "successfully BatchQueryAlgorithmAccessById, id=%#v", req.AlgorithmAccessIdList)
	return algorithmAccessList, nil
}

// 添加公共算法
func (d *algorithmDao) AddAlgorithmAccess(ctx context.Context, req *model.AlgorithmAccess) (*model.AlgorithmAccess, error) {
	db := d.db.Model(&model.AlgorithmAccess{})

	db = db.Create(req)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBCreateFailed)
		d.log.Errorw(ctx, err)
		return nil, err
	}

	d.log.Infof(ctx, "successfully AddAlgorithmAccess, algorithmId=%s|spaceId=%s", req.AlgorithmId, req.SpaceId)
	return req, nil
}

// 删除公共算法
func (d *algorithmDao) DeleteAlgorithmAccess(ctx context.Context, req *model.AlgorithmAccessDelete) error {
	db := d.db.Model(&model.AlgorithmAccess{})

	deleteParams := make([]interface{}, 0)
	deleteSql := "1 = 1"
	deleteSql += " and space_id = ? "
	deleteParams = append(deleteParams, req.SpaceId)
	deleteSql += " and algorithm_id = ? "
	deleteParams = append(deleteParams, req.AlgorithmId)

	// 删除
	db = db.Where(deleteSql, deleteParams...).Delete(&model.AlgorithmAccess{})
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBUpdateFailed)
		d.log.Errorw(ctx, err)
		return err
	}

	d.log.Infof(ctx, "successfully DeleteAlgorithmAccess, spaceId=%s|algorithmId=%s", req.SpaceId, req.AlgorithmId)
	return nil
}

// 修改公共算法
func (d *algorithmDao) UpdateAlgorithmAccess(ctx context.Context, req *model.AlgorithmAccess) error {
	db := d.db.Table(model.AlgorithmAccess{}.TableName())

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	querySql += " and data_version = ? "
	params = append(params, req.DataVersion)
	// 乐观锁自增
	req.DataVersion += 1

	db = db.Where(querySql, params).Updates(req)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBUpdateFailed)
		d.log.Errorw(ctx, err)
		return err
	}

	d.log.Infof(ctx, "successfully UpdateAlgorithmAccess, id=%s|algorithmlId=%s", req.Id, req.AlgorithmId)
	return nil
}

// 公共算法版本列表查询
func (d *algorithmDao) ListAlgorithmAccessVersion(ctx context.Context, req *model.AlgorithmAccessVersionList) (int64, []*model.AlgorithmAccessVersion, error) {
	db := d.db.Model(&model.AlgorithmAccessVersion{})
	accessAlgorithmVersions := make([]*model.AlgorithmAccessVersion, 0)

	// 查询where语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	if req.AlgorithmAccessId != "" {
		querySql += " and algorithm_access_id = ? "
		params = append(params, req.AlgorithmAccessId)
	}
	if req.AlgorithmId != "" {
		querySql += " and algorithm_id = ? "
		params = append(params, req.AlgorithmId)
	}
	if req.AlgorithmVersion != "" {
		querySql += " and algorithm_version = ? "
		params = append(params, req.AlgorithmVersion)
	}
	if req.SpaceId != "" {
		querySql += " and space_id = ? "
		params = append(params, req.SpaceId)
	}

	if len(params) == 0 {
		return 0, nil, errors.Errorf(nil, errors.ErrorDBSelectParamsEmpty)
	}

	db = db.Where(querySql, params...)

	// 总数计算
	var totalSize int64
	res := db.Count(&totalSize)

	if res.Error != nil {
		err := errors.Errorf(res.Error, errors.ErrorDBCountFailed)
		return 0, nil, err
	}

	if req.PageSize > 0 && req.PageIndex > 0 {
		db = db.Limit(req.PageSize).
			Offset((req.PageIndex - 1) * req.PageSize)
	}

	if req.AlgorithmAccessIdOrder {
		orderSql := fmt.Sprintf("algorithm_access_id %s", req.AlgorithmAccessIdSort)
		db = db.Order(orderSql)
	}
	if req.AlgorithmIdOrder {
		orderSql := fmt.Sprintf("algorithm_id %s", req.AlgorithmIdSort)
		db = db.Order(orderSql)
	}
	if req.AlgorithmVersionOrder {
		orderSql := fmt.Sprintf("algorithm_version %s", req.AlgorithmVersionSort)
		db = db.Order(orderSql)
	}
	if req.CreatedAtOrder {
		orderSql := fmt.Sprintf("created_at %s", req.CreatedAtSort)
		db = db.Order(orderSql)
	}

	db = db.Find(&accessAlgorithmVersions)

	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindFailed)
		d.log.Errorw(ctx, err)
		return 0, nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		d.log.Errorw(ctx, err)
		return 0, nil, err
	}
	d.log.Infof(ctx, "successfully ListAlgorithmVersionAccess, totalSize=%d", totalSize)
	return totalSize, accessAlgorithmVersions, nil
}

// 查询公共算法版本
func (d *algorithmDao) QueryAlgorithmAccessVersion(ctx context.Context, req *model.AlgorithmAccessVersionQuery) (*model.AlgorithmAccessVersion, error) {
	db := d.db.Model(&model.AlgorithmAccessVersion{})
	oneAlgorithmAccessVersion := model.AlgorithmAccessVersion{}

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	querySql += " and algorithm_access_id = ? "
	params = append(params, req.AlgorithmAccessId)
	querySql += " and algorithm_version = ? "
	params = append(params, req.AlgorithmVersion)

	// 	查询
	db = db.Where(querySql, params...).First(&oneAlgorithmAccessVersion)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFirstFailed)
		d.log.Errorw(ctx, err)
		return nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		d.log.Errorw(ctx, err)
		return nil, err
	}

	d.log.Infof(ctx, "successfully QueryAlgorithmAccessVersion, AlgorithmAccessId=%s|AlgorithmVersion=%s", req.AlgorithmAccessId, req.AlgorithmVersion)
	return &oneAlgorithmAccessVersion, nil
}

// 批量查询公共算法版本
func (d *algorithmDao) BatchQueryAlgorithmAccessVersion(ctx context.Context, req *model.AlgorithmAccessVersionBatchQuery) ([]*model.AlgorithmAccessVersion, error) {
	db := d.db.Model(&model.AlgorithmAccessVersion{})
	accessAlgorithmVersion := make([]*model.AlgorithmAccessVersion, 0)

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	querySql += " and (algorithm_access_id, algorithm_version) in ? "
	for _, query := range req.List {
		oneParams := fmt.Sprintf("(%s, %s)", query.AlgorithmAccessId, query.AlgorithmVersion)
		params = append(params, oneParams)
	}

	// 	查询
	db = db.Where(querySql, params...).Find(&accessAlgorithmVersion)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindFailed)
		d.log.Errorw(ctx, err)
		return nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		d.log.Errorw(ctx, err)
		return nil, err
	}

	d.log.Infof(ctx, "successfully BatchQueryAlgorithmAccessVersion, algorithmAccessId,algorithmVersion=%#v", req.List)
	return accessAlgorithmVersion, nil
}

// 添加公共算法版本
func (d *algorithmDao) AddAlgorithmAccessVersion(ctx context.Context, req *model.AlgorithmAccessVersion) (*model.AlgorithmAccessVersion, error) {
	db := d.db.Model(&model.AlgorithmAccessVersion{})

	db = db.Create(req)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBCreateFailed)
		d.log.Errorw(ctx, err)
		return nil, err
	}

	d.log.Infof(ctx, "successfully AddAlgorithmAccessVersion, id=%s|accessId=%s|algorithmId=%s|Versi=%s", req.Id, req.AlgorithmAccessId, req.AlgorithmId, req.AlgorithmVersion)
	return req, nil
}

// 批量删除公共算法
func (d *algorithmDao) BatchDeleteAlgorithmAccessVersion(ctx context.Context, req *model.AlgorithmAccessVersionBatchDelete) error {
	db := d.db.Model(&model.AlgorithmAccessVersion{})

	// 删除where语句拼接
	deleteParams := make([]interface{}, 0)
	deleteSql := "1 = 1"
	if req.AlgorithmAccessId != "" {
		deleteSql += " and algorithm_access_id = ? "
		deleteParams = append(deleteParams, req.AlgorithmAccessId)

		if len(req.AlgorithmVersion) > 0 {
			deleteSql += " and algorithm_version in ? "
			deleteParams = append(deleteParams, req.AlgorithmVersion)
		}
	}

	// 删除
	db = db.Where(deleteSql, deleteParams...).Delete(&model.AlgorithmAccessVersion{})
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBUpdateFailed)
		d.log.Errorw(ctx, err)
		return err
	}

	d.log.Infof(ctx, "successfully BatchDeleteAlgorithmAccessVersion, algorithmAccessId=%s|algorithmVersion=%#v", req.AlgorithmAccessId, req.AlgorithmVersion)
	return nil
}

// 删除公共算法版本
func (d *algorithmDao) DeleteAlgorithmAccessVersion(ctx context.Context, req *model.AlgorithmAccessVersionDelete) error {
	db := d.db.Model(&model.AlgorithmAccessVersion{})

	// 删除where语句拼接
	deleteParams := make([]interface{}, 0)
	deleteSql := "1 = 1"
	deleteSql += " and algorithm_access_id = ? "
	deleteParams = append(deleteParams, req.AlgorithmAccessId)
	deleteSql += " and algorithm_version = ? "
	deleteParams = append(deleteParams, req.AlgorithmVersion)

	// 删除
	db = db.Where(deleteSql, deleteParams...).Delete(&model.AlgorithmAccessVersion{})
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBUpdateFailed)
		d.log.Errorw(ctx, err)
		return err
	}

	d.log.Infof(ctx, "successfully DeleteAlgorithmAccessVersion, algorithmAccessId=%s|algorithmVersion=%s", req.AlgorithmAccessId, req.AlgorithmVersion)
	return nil
}
