package dao

import (
	"context"
	"fmt"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"
	"time"

	"server/common/log"

	"gorm.io/gorm"
)

type ModelDao interface {
	// 模型列表查询
	ListModel(ctx context.Context, req *model.ModelList) (int64, []*model.Model, error)
	// 查询模型
	GetModel(ctx context.Context, modelId string) (*model.Model, error)
	// 添加模型
	CreateModel(ctx context.Context, req *model.Model) (*model.Model, error)
	// 删除模型
	DeleteModel(ctx context.Context, modelId string) error
	// 修改模型
	UpdateModel(ctx context.Context, req *model.Model) error

	// 模型版本列表查询
	ListModelVersion(ctx context.Context, req *model.ModelVersionList) (int64, []*model.ModelVersion, error)
	// 查询模型版本
	GetModelVersion(ctx context.Context, modelId string, version string) (*model.ModelVersion, error)
	// 添加模型版本
	CreateModelVersion(ctx context.Context, req *model.ModelVersion) (*model.ModelVersion, error)
	// 删除模型版本
	DeleteModelVersion(ctx context.Context, modelId string, version string) error
	// 修改模型版本
	UpdateModelVersion(ctx context.Context, req *model.ModelVersion) error

	// 模型可见列表查询
	ListModelAccess(ctx context.Context, req *model.ModelAccessList) (int64, []*model.ModelAccess, error)
	// 查询可见模型
	GetModelAccess(ctx context.Context, modelAccessId string) (*model.ModelAccess, error)
	// 添加可见模型
	CreateModelAccess(ctx context.Context, req *model.ModelAccess) (*model.ModelAccess, error)
	// 删除可见模型
	DeleteModelAccess(ctx context.Context, modelAccessId string) error
	// 修改可见模型
	UpdateModelAccess(ctx context.Context, req *model.ModelAccess) error

	// 模型可见版本列表查询
	ListModelVersionAccess(ctx context.Context, req *model.ModelVersionAccessList) (int64, []*model.ModelVersionAccess, error)
	// 查询可见模型版本
	GetModelVersionAccess(ctx context.Context, modelAccessId string, modelVersion string) (*model.ModelVersionAccess, error)
	// 添加可见模型版本
	CreateModelVersionAccess(ctx context.Context, req *model.ModelVersionAccess) (*model.ModelVersionAccess, error)
	// 删除可见模型版本
	DeleteModelVersionAccess(ctx context.Context, modelAccessId string, modelVersion string) error
	// 修改可见模型版本
	UpdateModelVersionAccess(ctx context.Context, req *model.ModelVersionAccess) error
}

type modelDao struct {
	log *log.Helper
	db  *gorm.DB
}

func NewModelDao(db *gorm.DB, logger log.Logger) ModelDao {
	return &modelDao{
		log: log.NewHelper("ModelDao", logger),
		db:  db,
	}
}

// 查询模型列表
func (d *modelDao) ListModel(ctx context.Context, req *model.ModelList) (int64, []*model.Model, error) {
	db := d.db.Model(&model.Model{})
	models := make([]*model.Model, 0)

	// 查询where语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	if req.SpaceId != "" {
		querySql += " and space_id = ? "
		params = append(params, req.SpaceId)
	}

	if req.UserId != "" {
		querySql += " and user_id = ? "
		params = append(params, req.UserId)
	}

	if req.AlgorithmId != "" {
		querySql += " and algorithm_id = ? "
		params = append(params, req.AlgorithmId)
	}

	if req.AlgorithmVersion != "" {
		querySql += " and algorithm_version = ? "
		params = append(params, req.AlgorithmVersion)
	}

	if len(req.Ids) != 0 {
		querySql += " and id in ? "
		params = append(params, req.Ids)
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
		querySql += " and (model_name like ? "
		params = append(params, "%"+req.SearchKey+"%")
		querySql += " or model_descript like ? )"
		params = append(params, "%"+req.SearchKey+"%")
	}

	if req.FrameWorkId != "" {
		querySql += " and framework_id = ? "
		params = append(params, req.FrameWorkId)
	}

	querySql += " and is_prefab = ? "
	params = append(params, req.IsPrefab)

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Count(&totalSize)
	if res.Error != nil {
		err := errors.Errorf(res.Error, errors.ErrorDBCountFailed)
		return 0, nil, err
	}

	if req.PageIndex > 0 {
		// limit语句拼接
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

	db = db.Find(&models)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBFindFailed)
		return 0, nil, err
	}

	d.log.Infof(ctx, "successfully ListModel, totalSize=%d", totalSize)
	return totalSize, models, nil
}

// 查询模型
func (d *modelDao) GetModel(ctx context.Context, modelId string) (*model.Model, error) {
	db := d.db.Model(&model.Model{})
	oneModel := model.Model{}

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	querySql += " and id = ? "
	params = append(params, modelId)

	// 	查询
	db = db.Where(querySql, params...).First(&oneModel)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFirstFailed)
		return nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		return nil, err
	}

	d.log.Infof(ctx, "successfully GetModel, modelId=%s", modelId)
	return &oneModel, nil
}

// 添加模型
func (d *modelDao) CreateModel(ctx context.Context, req *model.Model) (*model.Model, error) {
	db := d.db.Model(&model.Model{})

	db = db.Create(req)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBCreateFailed)
		return nil, err
	}

	d.log.Infof(ctx, "successfully CreateModel, modeId=%s|spaceId=%s|userId=%s", req.Id, req.SpaceId, req.UserId)
	return req, nil
}

// 删除模型
func (d *modelDao) DeleteModel(ctx context.Context, modelId string) error {
	db := d.db.Model(&model.Model{})

	modelInt, err := d.GetModel(ctx, modelId)
	if err != nil {
		return err
	}

	// 删除
	db = db.Select("ModelVersions").Delete(modelInt)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBDeleteFailed)
		return err
	}

	d.log.Infof(ctx, "successfully DeleteModel, modelId=%s", modelId)
	return nil
}

// 修改模型
func (d *modelDao) UpdateModel(ctx context.Context, req *model.Model) error {
	db := d.db.Table(model.Model{}.TableName())

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
		return err
	}

	d.log.Infof(ctx, "successfully UpdateModel, id=%s", req.Id)
	return nil
}

// 查询模型版本列表
func (d *modelDao) ListModelVersion(ctx context.Context, req *model.ModelVersionList) (int64, []*model.ModelVersion, error) {
	db := d.db.Model(&model.ModelVersion{})
	modelVersions := make([]*model.ModelVersion, 0)

	// 查询where语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	if req.ModelId != "" {
		querySql += " and model_id = ? "
		params = append(params, req.ModelId)
	}

	if len(req.Versions) != 0 {
		querySql += " and version in ? "
		params = append(params, req.Versions)
	}

	if len(params) == 0 {
		return 0, nil, errors.Errorf(nil, errors.ErrorDBSelectParamsEmpty)
	}

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Count(&totalSize)
	if res.Error != nil {
		err := errors.Errorf(res.Error, errors.ErrorDBCountFailed)
		return 0, nil, err
	}

	if req.PageIndex > 0 {
		// limit语句拼接
		db = db.Limit(req.PageSize).
			Offset((req.PageIndex - 1) * req.PageSize)
	}

	// orderby语句拼接
	// left(version ,1) DESC , CAST(MID(version,2) AS UNSIGNED)
	if req.VersionOrder {
		orderSql := fmt.Sprintf("left(version ,1) %s", req.VersionSort)
		db = db.Order(orderSql)
		orderSql = fmt.Sprintf("CAST(MID(version,2) AS UNSIGNED) %s", req.VersionSort)
		db = db.Order(orderSql)
	}
	if req.CreatedAtOrder {
		orderSql := fmt.Sprintf("created_at %s", req.CreatedAtSort)
		db = db.Order(orderSql)
	}

	db = db.Find(&modelVersions)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBFindFailed)
		return 0, nil, err
	}

	d.log.Infof(ctx, "successfully ListModelVersion, totalSize=%d", totalSize)
	return totalSize, modelVersions, nil
}

// 查询模型版本
func (d *modelDao) GetModelVersion(ctx context.Context, modelId string, version string) (*model.ModelVersion, error) {
	db := d.db.Model(&model.ModelVersion{})
	modelVersion := model.ModelVersion{}

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	querySql += " and model_id = ? "
	params = append(params, modelId)
	querySql += " and version = ? "
	params = append(params, version)

	// 	查询
	db = db.Where(querySql, params...).First(&modelVersion)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFirstFailed)
		return nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		return nil, err
	}

	d.log.Infof(ctx, "successfully GetModelVersion, modelId=%s|version=%s", modelId, version)
	return &modelVersion, nil
}

// 添加模型版本
func (d *modelDao) CreateModelVersion(ctx context.Context, req *model.ModelVersion) (*model.ModelVersion, error) {
	db := d.db.Model(&model.ModelVersion{})

	db = db.Create(req)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBCreateFailed)
		return nil, err
	}

	d.log.Infof(ctx, "successfully CreateModelVersion, id=%s|modeId=%s|version=%s", req.Id, req.ModelId, req.Version)
	return req, nil
}

// 删除模型版本
func (d *modelDao) DeleteModelVersion(ctx context.Context, modelId string, version string) error {
	db := d.db.Model(&model.ModelVersion{})

	// 删除where语句拼接
	params := make([]interface{}, 0)
	deleteSql := "1 = 1"
	deleteSql += " and model_id = ? "
	params = append(params, modelId)
	deleteSql += " and version = ? "
	params = append(params, version)

	// 删除
	db = db.Where(deleteSql, params...).Delete(&model.ModelVersion{})

	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBDeleteFailed)
		return err
	}

	d.log.Infof(ctx, "successfully DeleteModelVersion, modelId=%s|version=%s", modelId, version)
	return nil
}

// 修改模型版本
func (d *modelDao) UpdateModelVersion(ctx context.Context, req *model.ModelVersion) error {
	db := d.db.Table(model.ModelVersion{}.TableName())

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
		return err
	}

	d.log.Infof(ctx, "successfully UpdateModelVersion, id=%s|modeId=%s|version=%s", req.Id, req.ModelId, req.Version)
	return nil
}

// 查询模型可见列表
func (d *modelDao) ListModelAccess(ctx context.Context, req *model.ModelAccessList) (int64, []*model.ModelAccess, error) {
	db := d.db.Model(&model.ModelAccess{})
	accessModels := make([]*model.ModelAccess, 0)

	// 查询where语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	if len(req.SpaceIds) != 0 {
		querySql += " and space_id in ? "
		params = append(params, req.SpaceIds)
	}

	if len(req.ModelIds) != 0 {
		querySql += " and model_id in ? "
		params = append(params, req.ModelIds)
	}

	if len(req.Ids) != 0 {
		querySql += " and id in ? "
		params = append(params, req.Ids)
	}

	if req.FrameWorkId != "" {
		joinSql := " Inner JOIN (select id as mid,framework_id from model)mm on mm.mid = model_access.model_id "
		querySql += " and framework_id = ?"
		params = append(params, req.FrameWorkId)
		db = db.Joins(joinSql).Where(querySql, params...)
	}

	if len(params) == 0 {
		return 0, nil, errors.Errorf(nil, errors.ErrorDBSelectParamsEmpty)
	}

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Count(&totalSize)
	if res.Error != nil {
		err := errors.Errorf(res.Error, errors.ErrorDBCountFailed)
		return 0, nil, err
	}

	if req.PageIndex > 0 {
		// limit语句拼接
		db = db.Limit(req.PageSize).
			Offset((req.PageIndex - 1) * req.PageSize)
	}

	// orderby语句拼接
	if req.SpaceIdOrder {
		orderSql := fmt.Sprintf("space_id %s", req.SpaceIdSort)
		db = db.Order(orderSql)
	}
	if req.ModelIdOrder {
		orderSql := fmt.Sprintf("model_id %s", req.ModelIdSort)
		db = db.Order(orderSql)
	}
	if req.CreatedAtOrder {
		orderSql := fmt.Sprintf("created_at %s", req.CreatedAtSort)
		db = db.Order(orderSql)
	}

	db = db.Find(&accessModels)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBFindFailed)
		return 0, nil, err
	}

	d.log.Infof(ctx, "successfully ListModelAccess, totalSize=%d", totalSize)
	return totalSize, accessModels, nil
}

// 查询可见模型
func (d *modelDao) GetModelAccess(ctx context.Context, modelAccessId string) (*model.ModelAccess, error) {
	db := d.db.Model(&model.ModelAccess{})
	oneModelAccess := model.ModelAccess{}

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	querySql += " and id = ? "
	params = append(params, modelAccessId)

	// 	查询
	db = db.Where(querySql, params...).First(&oneModelAccess)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFirstFailed)
		return nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		return nil, err
	}

	d.log.Infof(ctx, "successfully GetModelAccess, modelAccessId=%s", modelAccessId)
	return &oneModelAccess, nil
}

// 添加可见模型
func (d *modelDao) CreateModelAccess(ctx context.Context, req *model.ModelAccess) (*model.ModelAccess, error) {
	db := d.db.Model(&model.ModelAccess{})

	db = db.Create(req)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBCreateFailed)
		return nil, err
	}

	d.log.Infof(ctx, "successfully CreateModelAccess, modeId=%s|spaceId=%s", req.ModelId, req.SpaceId)
	return req, nil
}

// 删除可见模型
func (d *modelDao) DeleteModelAccess(ctx context.Context, modelAccessId string) error {
	db := d.db.Model(&model.ModelAccess{})

	modelAccessInt, err := d.GetModelAccess(ctx, modelAccessId)
	if err != nil {
		return err
	}

	// 删除
	db = db.Select("ModelVersionAccesss").Delete(modelAccessInt)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBDeleteFailed)
		return err
	}

	d.log.Infof(ctx, "successfully DeleteModelAccess, modelAccessId=%s", modelAccessId)
	return nil
}

// 修改模型
func (d *modelDao) UpdateModelAccess(ctx context.Context, req *model.ModelAccess) error {
	db := d.db.Table(model.ModelAccess{}.TableName())

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
		return err
	}

	d.log.Infof(ctx, "successfully UpdateModelAccess, id=%s|modelId=%s", req.Id, req.ModelId)
	return nil
}

// 查询模型可见版本列表
func (d *modelDao) ListModelVersionAccess(ctx context.Context, req *model.ModelVersionAccessList) (int64, []*model.ModelVersionAccess, error) {
	db := d.db.Model(&model.ModelVersionAccess{})
	accessModelVersions := make([]*model.ModelVersionAccess, 0)

	// 查询where语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	if req.ModelAccessId != "" {
		querySql += " and model_access_id = ? "
		params = append(params, req.ModelAccessId)
	}
	if req.ModelId != "" {
		querySql += " and model_id = ? "
		params = append(params, req.ModelId)
	}
	if len(req.ModelVersions) != 0 {
		querySql += " and model_version in ? "
		params = append(params, req.ModelVersions)
	}

	if len(params) == 0 {
		return 0, nil, errors.Errorf(nil, errors.ErrorDBSelectParamsEmpty)
	}

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Count(&totalSize)
	if res.Error != nil {
		err := errors.Errorf(res.Error, errors.ErrorDBCountFailed)
		return 0, nil, err
	}

	// 列表查询
	if req.PageIndex > 0 {
		// limit语句拼接
		db = db.Limit(req.PageSize).
			Offset((req.PageIndex - 1) * req.PageSize)
	}

	// orderby语句拼接
	if req.ModelAccessIdOrder {
		orderSql := fmt.Sprintf("model_access_id %s", req.ModelAccessIdSort)
		db = db.Order(orderSql)
	}
	if req.ModelIdOrder {
		orderSql := fmt.Sprintf("model_id %s", req.ModelIdSort)
		db = db.Order(orderSql)
	}
	if req.ModelVersionOrder {
		// left(model_version ,1) DESC , CAST(MID(model_version,2) AS UNSIGNED)
		orderSql := fmt.Sprintf("left(model_version ,1) %s", req.ModelVersionSort)
		db = db.Order(orderSql)
		orderSql = fmt.Sprintf("CAST(MID(model_version,2) AS UNSIGNED) %s", req.ModelVersionSort)
		db = db.Order(orderSql)
	}
	if req.CreatedAtOrder {
		orderSql := fmt.Sprintf("created_at %s", req.CreatedAtSort)
		db = db.Order(orderSql)
	}

	db = db.Find(&accessModelVersions)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBFindFailed)
		return 0, nil, err
	}

	d.log.Infof(ctx, "successfully ListModelVersionAccess, totalSize=%d", totalSize)
	return totalSize, accessModelVersions, nil
}

// 查询可见模型版本
func (d *modelDao) GetModelVersionAccess(ctx context.Context, modelAccessId string, modelVersion string) (*model.ModelVersionAccess, error) {
	db := d.db.Model(&model.ModelVersionAccess{})
	oneAccessModelVersion := model.ModelVersionAccess{}

	// where 语句拼接
	params := make([]interface{}, 0)
	querySql := "1 = 1"
	querySql += " and model_access_id = ? "
	params = append(params, modelAccessId)
	querySql += " and model_version = ? "
	params = append(params, modelVersion)

	// 	查询
	db = db.Where(querySql, params...).First(&oneAccessModelVersion)
	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFirstFailed)
		return nil, err
	}
	if db.Error == gorm.ErrRecordNotFound {
		err := errors.Errorf(db.Error, errors.ErrorDBFindEmpty)
		return nil, err
	}

	d.log.Infof(ctx, "successfully GetyModelVersionAccess, modelAccessId=%s|modelVersion=%s", modelAccessId, modelVersion)
	return &oneAccessModelVersion, nil
}

// 添加可见模型
func (d *modelDao) CreateModelVersionAccess(ctx context.Context, req *model.ModelVersionAccess) (*model.ModelVersionAccess, error) {
	db := d.db.Model(&model.ModelVersionAccess{})

	db = db.Create(req)
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBCreateFailed)
		return nil, err
	}

	d.log.Infof(ctx, "successfully CreateModelVersionAccess, id=%s|modelAccessId=%s|modelId=%s|modelVersion=%s", req.Id, req.ModelAccessId, req.ModelId, req.ModelVersion)
	return req, nil
}

// 删除可见模型
func (d *modelDao) DeleteModelVersionAccess(ctx context.Context, modelAccessId string, modelVersion string) error {
	db := d.db.Model(&model.ModelVersionAccess{})

	// 删除where语句拼接
	params := make([]interface{}, 0)
	deleteSql := "1 = 1"
	deleteSql += " and model_access_id = ? "
	params = append(params, modelAccessId)
	deleteSql += " and model_version = ? "
	params = append(params, modelVersion)

	// 删除
	db = db.Where(deleteSql, params...).Delete(&model.ModelVersionAccess{})
	if db.Error != nil {
		err := errors.Errorf(db.Error, errors.ErrorDBDeleteFailed)
		return err
	}

	d.log.Infof(ctx, "successfully DeleteModelVersionAccess, modelAccessId=%s|modelVersion=%s", modelAccessId, modelVersion)
	return nil
}

// 修改可见模型版本
func (d *modelDao) UpdateModelVersionAccess(ctx context.Context, req *model.ModelVersionAccess) error {
	db := d.db.Table(model.ModelVersionAccess{}.TableName())

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
		return err
	}

	d.log.Infof(ctx, "successfully UpdateModelVersionAccess, id=%s|modeId=%s|version=%s", req.Id, req.ModelId, req.ModelVersion)
	return nil
}
