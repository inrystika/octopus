package algorithm_dao

import (
	"context"
	"fmt"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"
	"time"

	"gorm.io/gorm"
)

// 算法列表查询
func (d *algorithmDao) ListAlgorithm(ctx context.Context, req *model.AlgorithmList) (int64, []*model.Algorithm, error) {
	db := d.db.Model(&model.Algorithm{})
	algorithms := make([]*model.Algorithm, 0)

	params := make([]interface{}, 0)

	querySql := "1 = 1"
	querySql += " and is_prefab = ? "
	params = append(params, req.IsPrefab)

	if req.SpaceId != "" {
		querySql += " and space_id = ? "
		params = append(params, req.SpaceId)
	}
	if req.UserId != "" {
		querySql += " and user_id = ? "
		params = append(params, req.UserId)
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

	// 总数计算
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
	if req.UserIdOrder {
		orderSql := fmt.Sprintf("user_id %s", req.UserIdSort)
		db = db.Order(orderSql)
	}
	if req.CreatedAtOrder {
		orderSql := fmt.Sprintf("created_at %s", req.CreatedAtSort)
		db = db.Order(orderSql)
	}

	if req.SortBy != "" {
		db = req.Order(db)
	}

	db = db.Find(&algorithms)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFirstFailed)
		d.log.Errorw(ctx, err)
		return 0, nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		d.log.Errorw(ctx, err)
		return 0, nil, err
	}

	d.log.Infof(ctx, "successfully ListAlgorithm, totalSize=%d", totalSize)
	return totalSize, algorithms, nil
}

// 查询算法
func (d *algorithmDao) QueryAlgorithm(ctx context.Context, req *model.AlgorithmQuery) (*model.Algorithm, error) {
	db := d.db.Model(&model.Algorithm{})
	oneAlgorithm := model.Algorithm{}

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := " algorithm_id = ? "
	params = append(params, req.AlgorithmId)

	// 	查询
	db = db.Where(querySql, params...).First(&oneAlgorithm)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFirstFailed)
		//d.log.Errorw(ctx, err)
		return nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		//d.log.Errorw(ctx, err)
		return nil, err
	}

	d.log.Infof(ctx, "successfully QueryAlgorithm, algorithmId=%s", req.AlgorithmId)
	return &oneAlgorithm, nil
}

// 查询算法
func (d *algorithmDao) QueryAlgorithmByInfo(ctx context.Context, req *model.AlgorithmQueryByInfo) (*model.Algorithm, error) {
	db := d.db.Model(&model.Algorithm{})
	oneAlgorithm := model.Algorithm{}

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"

	querySql += " and is_prefab = ? "
	params = append(params, req.IsPrefab)

	if req.UserId != "" {
		querySql += " and user_id = ? "
		params = append(params, req.UserId)
	}

	if req.SpaceId != "" {
		querySql += " and space_id = ? "
		params = append(params, req.SpaceId)
	}

	if req.AlgorithmName != "" {
		querySql += " and algorithm_name = ? "
		params = append(params, req.AlgorithmName)
	}

	// 	查询
	db = db.Where(querySql, params...).First(&oneAlgorithm)
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

	d.log.Infof(ctx, "successfully QueryAlgorithm, algorithmName=%s", req.AlgorithmName)
	return &oneAlgorithm, nil
}

// 批量查询算法
func (d *algorithmDao) BatchQueryAlgorithm(ctx context.Context, req *model.AlgorithmBatchQuery) ([]*model.Algorithm, error) {
	db := d.db.Model(&model.Algorithm{})
	algorithms := make([]*model.Algorithm, 0)

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := " algorithm_id = ? "
	for _, query := range req.List {
		params = append(params, query.AlgorithmId)
	}

	// 	查询
	db = db.Where(querySql, params...).Find(&algorithms)
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

	d.log.Infof(ctx, "successfully BatchQueryAlgorithm, algorithmId=%#v", req.List)
	return algorithms, nil
}

// 添加算法
func (d *algorithmDao) AddAlgorithm(ctx context.Context, req *model.Algorithm) (*model.Algorithm, error) {
	db := d.db.Model(&model.Algorithm{})
	db = db.Create(req)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBCreateFailed)
		d.log.Errorw(ctx, err)
		return nil, err
	}

	d.log.Infof(ctx, "successfully AddAlgorithm, algorithmId=%s|spaceId=%s|userId=%s", req.AlgorithmId, req.SpaceId, req.UserId)
	return req, nil
}

// 删除算法
func (d *algorithmDao) DeleteAlgorithm(ctx context.Context, req *model.AlgorithmDelete) error {
	db := d.db.Model(&model.Algorithm{})

	// where 语句拼接
	params := make([]interface{}, 0)
	deleteSql := " algorithm_id = ? "
	params = append(params, req.AlgorithmId)

	// 删除
	db = db.Where(deleteSql, params).Delete(&model.Algorithm{})
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBUpdateFailed)
		d.log.Errorw(ctx, err)
		return err
	}

	d.log.Infof(ctx, "successfully DeleteAlgorithm, algorithmId=%s", req.AlgorithmId)
	return nil
}

// 修改算法
func (d *algorithmDao) UpdateAlgorithm(ctx context.Context, req *model.Algorithm) error {
	db := d.db.Table(model.Algorithm{}.TableName())

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

	d.log.Infof(ctx, "successfully UpdateAlgorithm, id=%s", req.AlgorithmId)
	return nil
}

// 算法版本列表查询
func (d *algorithmDao) ListAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersionList) (int64, []*model.AlgorithmVersion, error) {
	db := d.db.Model(&model.AlgorithmVersion{})
	algorithmVersions := make([]*model.AlgorithmVersion, 0)

	// 查询where语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	querySql += " and algorithm_id = ? "
	params = append(params, req.AlgorithmId)

	if req.FileStatus != 0 {
		querySql += " and file_status = ? "
		params = append(params, req.FileStatus)
	}

	db = db.Where(querySql, params...)

	// 总数计算
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
	if req.VersionOrder {
		orderSql := fmt.Sprintf("version %s", req.VersionSort)
		db = db.Order(orderSql)
	}
	if req.CreatedAtOrder {
		orderSql := fmt.Sprintf("created_at %s", req.CreatedAtSort)
		db = db.Order(orderSql)
	}

	db = db.Find(&algorithmVersions)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFirstFailed)
		d.log.Errorw(ctx, err)
		return 0, nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		d.log.Errorw(ctx, err)
		return 0, nil, err
	}

	d.log.Infof(ctx, "successfully ListAlgorithmVersion, totalSize=%d", totalSize)
	return totalSize, algorithmVersions, nil
}

// 查询算法版本
func (d *algorithmDao) QueryAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersionQuery) (*model.AlgorithmVersion, error) {
	db := d.db.Model(&model.AlgorithmVersion{})
	algorithmVersion := model.AlgorithmVersion{}

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	querySql += " and algorithm_id = ? "
	params = append(params, req.AlgorithmId)
	querySql += " and version = ? "
	params = append(params, req.Version)

	// 	查询
	db = db.Where(querySql, params...).First(&algorithmVersion)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFirstFailed)
		//d.log.Errorw(ctx, err)
		return nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		//d.log.Errorw(ctx, err)
		return nil, err
	}

	d.log.Infof(ctx, "successfully QueryAlgorithmVersion, algorithmId=%s|version=%s", req.AlgorithmId, req.Version)
	return &algorithmVersion, nil
}

// 批量查询算法版本
func (d *algorithmDao) BatchQueryAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersionBatchQuery) ([]*model.AlgorithmVersion, error) {
	db := d.db.Model(&model.AlgorithmVersion{})
	algorithmVersions := make([]*model.AlgorithmVersion, 0)

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	querySql += " and (algprithm_id, version) in ? "
	for _, query := range req.AlgorithmVersionList {
		oneParams := fmt.Sprintf("(%s, %s)", query.AlgorithmId, query.Version)
		params = append(params, oneParams)
	}

	// 	查询
	db = db.Where(querySql, params...).Find(&algorithmVersions)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindFailed)
		//d.log.Errorw(ctx, err)
		return nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		//d.log.Errorw(ctx, err)
		return nil, err
	}

	d.log.Infof(ctx, "successfully BatchQueryAlgorithmVersion, algorithmId,version=%#v", req.AlgorithmVersionList)
	return algorithmVersions, nil
}

// 添加算法版本
func (d *algorithmDao) AddAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersion) (*model.AlgorithmVersion, error) {
	db := d.db.Model(&model.AlgorithmVersion{})

	db = db.Create(req)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBCreateFailed)
		d.log.Errorw(ctx, err)
		return nil, err
	}

	d.log.Infof(ctx, "successfully AddAlgorithmVersion, id=%s|algorithmId=%s|version=%s", req.Id, req.AlgorithmId, req.Version)
	return req, nil
}

// 删除算法版本
func (d *algorithmDao) DeleteAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersionDelete) error {
	db := d.db.Model(&model.AlgorithmVersion{})

	// 删除where语句拼接
	deleteParams := make([]interface{}, 0)
	deleteSql := "1 = 1"
	deleteSql += " and algorithm_id = ? "
	deleteParams = append(deleteParams, req.AlgorithmId)
	deleteSql += " and version = ? "
	deleteParams = append(deleteParams, req.AlgorithmVersion)
	// 删除
	db = db.Where(deleteSql, deleteParams...).Delete(&model.AlgorithmVersion{})

	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBUpdateFailed)
		d.log.Errorw(ctx, err)
		return err
	}
	d.log.Infof(ctx, "successfully DeleteAlgorithmVersion, algorithmId=%s|version=%s", req.AlgorithmId, req.AlgorithmVersion)
	return nil
}

// 批量删除算法版本
func (d *algorithmDao) BatchDeleteAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersionBatchDelete) error {
	db := d.db.Model(&model.AlgorithmVersion{})

	// 删除where语句拼接
	deleteParams := make([]interface{}, 0)
	deleteSql := "1 = 1"
	if req.AlgorithmId != "" {
		deleteSql += " and algorithm_id = ? "
		deleteParams = append(deleteParams, req.AlgorithmId)
	}
	if len(req.AlgorithmVersion) > 0 {
		deleteSql += " and version in ? "
		deleteParams = append(deleteParams, req.AlgorithmVersion)
	}

	// 删除
	db = db.Where(deleteSql, deleteParams...).Delete(&model.AlgorithmVersion{})

	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBUpdateFailed)
		d.log.Errorw(ctx, err)
		return err
	}

	d.log.Infof(ctx, "successfully BatchDeleteAlgorithmVersion, algorithmId=%s|version=%#v", req.AlgorithmId, req.AlgorithmVersion)
	return nil
}

// 修改算法版本
func (d *algorithmDao) UpdateAlgorithmVersion(ctx context.Context, req *model.AlgorithmVersion) error {
	db := d.db.Table(model.AlgorithmVersion{}.TableName())

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

	d.log.Infof(ctx, "successfully UpdateAlgorithmVersion, id=%s|algorithmId=%s|version=%s", req.Id, req.AlgorithmId, req.Version)
	return nil
}
