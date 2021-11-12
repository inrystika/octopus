package dao

import (
	"context"
	stderrors "errors"
	"fmt"
	"server/base-server/internal/data/dao/model"
	"server/common/errors"
	"server/common/transaction"
	"server/common/utils"
	"time"

	"server/common/log"

	"gorm.io/gorm"
)

type DatasetDao interface {
	Transaction(ctx context.Context, fc func(ctx context.Context) error) error

	AddDatasetType(ctx context.Context, datasetType *model.DatasetType) error
	ListDatasetType(ctx context.Context, query *model.DatasetTypeQuery) ([]*model.DatasetType, int64, error)
	GetDatasetType(ctx context.Context, id string) (*model.DatasetType, error)
	QueryDatasetType(ctx context.Context, typeDesc string) (*model.DatasetType, error)
	DeleteDatasetType(ctx context.Context, id string) error
	UpdateDatasetType(ctx context.Context, datasetType *model.DatasetType) error

	CreateDataset(ctx context.Context, dataset *model.Dataset) error
	ListDataset(ctx context.Context, query *model.DatasetQuery) ([]*model.Dataset, int64, error)
	ListCommDataset(ctx context.Context, query *model.CommDatasetQuery) ([]*model.Dataset, int64, error)
	GetDataset(ctx context.Context, id string) (*model.Dataset, error)
	DeleteDataset(ctx context.Context, id string) error
	UpdateDatasetSelective(ctx context.Context, dataset *model.Dataset) error

	CreateDatasetVersion(ctx context.Context, version *model.DatasetVersion) error
	ListDatasetVersion(ctx context.Context, query *model.DatasetVersionQuery) ([]*model.DatasetVersion, int64, error)
	ListCommDatasetVersion(ctx context.Context, query *model.CommDatasetVersionQuery) ([]*model.DatasetVersion, int64, error)
	GetDatasetVersion(ctx context.Context, datasetId string, version string) (*model.DatasetVersion, error)
	DeleteDatasetVersion(ctx context.Context, delete *model.DatasetVersionDelete) error
	UpdateDatasetVersionSelective(ctx context.Context, version *model.DatasetVersion) error
	ListDatasetVersionLatestVersion(ctx context.Context, datasetIds []string) (map[string]int64, error)

	ListDatasetAccess(ctx context.Context, query *model.DatasetAccessQuery) ([]*model.DatasetAccess, error)
	CreateDatasetAccess(ctx context.Context, access *model.DatasetAccess) error
	DeleteDatasetAccess(ctx context.Context, delete *model.DatasetAccessDelete) error
	UpdateDatasetAccessSelective(ctx context.Context, access *model.DatasetAccess) error

	ListDatasetVersionAccess(ctx context.Context, query *model.DatasetVersionAccessQuery) ([]*model.DatasetVersionAccess, error)
	CreateDatasetVersionAccess(ctx context.Context, versionAccess *model.DatasetVersionAccess) error
	DeleteDatasetVersionAccess(ctx context.Context, delete *model.DatasetVersionAccessDelete) error
	ListDatasetVersionAccessLatestVersion(ctx context.Context, ids []model.DatasetAccessId) (map[model.DatasetAccessId]int64, error)
}

type datasetDao struct {
	log *log.Helper
	db  transaction.GetDB
}

func NewDatasetDao(db *gorm.DB, logger log.Logger) DatasetDao {
	return &datasetDao{
		log: log.NewHelper("DatasetDao", logger),
		db: func(ctx context.Context) *gorm.DB {
			return transaction.GetDBFromCtx(ctx, db)
		},
	}
}

func (d *datasetDao) Transaction(ctx context.Context, fc func(ctx context.Context) error) error {
	return transaction.Transaction(ctx, d.db(ctx), fc)
}

func (d *datasetDao) AddDatasetType(ctx context.Context, datasetType *model.DatasetType) error {
	res := d.db(ctx).Create(datasetType)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *datasetDao) ListDatasetType(ctx context.Context, query *model.DatasetTypeQuery) ([]*model.DatasetType, int64, error) {
	db := d.db(ctx).Model(&model.DatasetType{})
	datasetTypes := make([]*model.DatasetType, 0)

	var totalSize int64
	res := db.Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if query.PageIndex != 0 {
		db = db.Limit(query.PageSize).Offset((query.PageIndex - 1) * query.PageSize)
	}

	res = db.Find(&datasetTypes)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return datasetTypes, totalSize, nil
}

func (d *datasetDao) GetDatasetType(ctx context.Context, id string) (*model.DatasetType, error) {
	db := d.db(ctx)

	nb := &model.DatasetType{}
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

func (d *datasetDao) QueryDatasetType(ctx context.Context, typeDesc string) (*model.DatasetType, error) {
	db := d.db(ctx)

	nb := &model.DatasetType{}
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

func (d *datasetDao) DeleteDatasetType(ctx context.Context, id string) error {
	db := d.db(ctx)

	if id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	res := db.Where("id = ? ", id).Delete(&model.DatasetType{})
	if res.Error != nil {
		return errors.Errorf(nil, errors.ErrorDBDeleteFailed)
	}

	return nil
}

func (d *datasetDao) UpdateDatasetType(ctx context.Context, datasetType *model.DatasetType) error {
	db := d.db(ctx)
	if datasetType.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := db.Model(&model.DatasetType{}).Updates(map[string]interface{}{
		"id":          datasetType.Id,
		"desc":        datasetType.Desc,
		"refer_times": datasetType.ReferTimes,
	})

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *datasetDao) CreateDataset(ctx context.Context, dataset *model.Dataset) error {
	res := d.db(ctx).Create(dataset)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *datasetDao) ListDataset(ctx context.Context, query *model.DatasetQuery) ([]*model.Dataset, int64, error) {
	db := d.db(ctx)
	datasets := make([]*model.Dataset, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if query.CreatedAtGte != 0 {
		querySql += " and created_at >= ? "
		params = append(params, time.Unix(query.CreatedAtGte, 0))
	}

	if query.CreatedAtLt != 0 {
		querySql += " and created_at < ? "
		params = append(params, time.Unix(query.CreatedAtLt, 0))
	}

	if query.SearchKey != "" {
		querySql += " and name like ?"
		params = append(params, "%"+query.SearchKey+"%")
	}

	if query.NameLike != "" {
		querySql += " and name like ?"
		params = append(params, "%"+query.NameLike+"%")
	}

	if query.UserId != "" {
		querySql += " and user_id = ? "
		params = append(params, query.UserId)
	}

	if query.SpaceId != "" {
		querySql += " and space_id = ? "
		params = append(params, query.SpaceId)
	}

	if query.SourceType != 0 {
		querySql += " and source_type = ? "
		params = append(params, query.SourceType)
	}

	if len(query.Ids) != 0 {
		querySql += " and id in ? "
		params = append(params, query.Ids)
	}

	if query.Name != "" {
		querySql += " and name = ? "
		params = append(params, query.Name)
	}

	db = db.Model(&model.Dataset{}).Where(querySql, params...)

	var totalSize int64
	res := db.Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if query.PageIndex != 0 {
		db = db.Limit(query.PageSize).Offset((query.PageIndex - 1) * query.PageSize)
	}

	sortBy := "created_at"
	orderBy := "desc"
	if query.SortBy != "" {
		sortBy = utils.CamelToSnake(query.SortBy)
	}

	if query.OrderBy != "" {
		orderBy = query.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", sortBy, orderBy))

	res = db.Find(&datasets)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return datasets, totalSize, nil
}

func (d *datasetDao) ListCommDataset(ctx context.Context, query *model.CommDatasetQuery) ([]*model.Dataset, int64, error) {
	db := d.db(ctx)
	datasets := make([]*model.Dataset, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if query.CreatedAtGte != 0 {
		querySql += " and dataset.created_at >= ? "
		params = append(params, time.Unix(query.CreatedAtGte, 0))
	}

	if query.CreatedAtLt != 0 {
		querySql += " and dataset.created_at < ? "
		params = append(params, time.Unix(query.CreatedAtLt, 0))
	}

	if query.SearchKey != "" {
		querySql += " and dataset.name like ?"
		params = append(params, "%"+query.SearchKey+"%")
	}

	if query.NameLike != "" {
		querySql += " and dataset.name like ?"
		params = append(params, "%"+query.NameLike+"%")
	}

	if query.UserId != "" {
		querySql += " and dataset.user_id = ? "
		params = append(params, query.UserId)
	}

	if query.SpaceId != "" {
		querySql += " and dataset.space_id = ? "
		params = append(params, query.SpaceId)
	}

	if query.SourceType != 0 {
		querySql += " and dataset.source_type = ? "
		params = append(params, query.SourceType)
	}

	if len(query.Ids) != 0 {
		querySql += " and dataset.id in ? "
		params = append(params, query.Ids)
	}

	if query.ShareSpaceId != "" {
		querySql += " and da.space_id = ? "
		params = append(params, query.ShareSpaceId)

		querySql += " and da.deleted_at = 0 "
	}

	db = db.Joins("join dataset_access as da on da.dataset_id = dataset.id")

	db = db.Model(&model.Dataset{}).Where(querySql, params...)

	var totalSize int64
	res := db.Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if query.PageIndex != 0 {
		db = db.Limit(query.PageSize).Offset((query.PageIndex - 1) * query.PageSize)
	}

	sortBy := "dataset.created_at"
	orderBy := "desc"
	if query.SortBy != "" {
		sortBy = "dataset." + utils.CamelToSnake(query.SortBy)
	}

	if query.OrderBy != "" {
		orderBy = query.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", sortBy, orderBy))

	res = db.Find(&datasets)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return datasets, totalSize, nil
}

func (d *datasetDao) GetDataset(ctx context.Context, id string) (*model.Dataset, error) {
	db := d.db(ctx)

	nb := &model.Dataset{}
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

func (d *datasetDao) DeleteDataset(ctx context.Context, id string) error {
	db := d.db(ctx)

	if id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	res := db.Where("id = ? ", id).Delete(&model.Dataset{})
	if res.Error != nil {
		return errors.Errorf(nil, errors.ErrorDBDeleteFailed)
	}

	return nil
}

func (d *datasetDao) UpdateDatasetSelective(ctx context.Context, dataset *model.Dataset) error {
	db := d.db(ctx)
	if dataset.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := db.Where("id = ?", dataset.Id).Updates(dataset)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *datasetDao) CreateDatasetVersion(ctx context.Context, version *model.DatasetVersion) error {
	res := d.db(ctx).Create(version)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *datasetDao) ListDatasetVersion(ctx context.Context, query *model.DatasetVersionQuery) ([]*model.DatasetVersion, int64, error) {
	db := d.db(ctx)
	versions := make([]*model.DatasetVersion, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if query.DatasetId != "" {
		querySql += " and dataset_id = ? "
		params = append(params, query.DatasetId)
	}

	if len(query.Ids) > 0 {
		inSql := ""
		for i, j := range query.Ids {
			if i != 0 {
				inSql += ","
			}
			inSql += fmt.Sprintf("('%s', '%s')", j.DatasetId, j.Version)
		}

		querySql += fmt.Sprintf(" and (dataset_id, version) in (%s) ", inSql)
	}

	if query.Status != 0 {
		querySql += " and status = ? "
		params = append(params, query.Status)
	}

	db = db.Model(&model.DatasetVersion{}).
		Where(querySql, params...)

	var totalSize int64
	res := db.Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if query.PageIndex != 0 {
		db = db.Limit(query.PageSize).Offset((query.PageIndex - 1) * query.PageSize)
	}

	sortBy := "created_at"
	orderBy := "desc"
	if query.SortBy != "" {
		sortBy = utils.CamelToSnake(query.SortBy)
	}

	if query.OrderBy != "" {
		orderBy = query.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", sortBy, orderBy))

	res = db.Find(&versions)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return versions, totalSize, nil
}

func (d *datasetDao) ListCommDatasetVersion(ctx context.Context, query *model.CommDatasetVersionQuery) ([]*model.DatasetVersion, int64, error) {
	db := d.db(ctx)
	versions := make([]*model.DatasetVersion, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if query.DatasetId != "" {
		querySql += " and dataset_version.dataset_id = ? "
		params = append(params, query.DatasetId)
	}

	if len(query.Ids) > 0 {
		inSql := ""
		for i, j := range query.Ids {
			if i != 0 {
				inSql += ","
			}
			inSql += fmt.Sprintf("('%s', '%s')", j.DatasetId, j.Version)
		}

		querySql += fmt.Sprintf(" and (dataset_version.dataset_id, dataset_version.version) in (%s) ", inSql)
	}

	if query.ShareSpaceId != "" {
		querySql += " and dva.space_id = ? "
		params = append(params, query.ShareSpaceId)
		querySql += " and dva.deleted_at = 0"
	}

	if query.Status != 0 {
		querySql += " and dataset_version.status = ? "
		params = append(params, query.Status)
	}

	db = db.Joins("join dataset_version_access as dva " +
		"on dva.dataset_id = dataset_version.dataset_id and dva.version = dataset_version.version")

	db = db.Model(&model.DatasetVersion{}).
		Where(querySql, params...)

	var totalSize int64
	res := db.Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if query.PageIndex != 0 {
		db = db.Limit(query.PageSize).Offset((query.PageIndex - 1) * query.PageSize)
	}

	sortBy := "dataset_version.created_at"
	orderBy := "desc"
	if query.SortBy != "" {
		sortBy = "dataset_version." + utils.CamelToSnake(query.SortBy)
	}

	if query.OrderBy != "" {
		orderBy = query.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", sortBy, orderBy))

	res = db.Find(&versions)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return versions, totalSize, nil
}

func (d *datasetDao) GetDatasetVersion(ctx context.Context, datasetId string, version string) (*model.DatasetVersion, error) {
	db := d.db(ctx)
	nb := &model.DatasetVersion{}
	res := db.First(nb, "dataset_id = ? and version = ?", datasetId, version)

	if res.Error != nil {
		if stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFindEmpty)
		} else {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFirstFailed)
		}
	}
	return nb, nil
}

func (d *datasetDao) DeleteDatasetVersion(ctx context.Context, delete *model.DatasetVersionDelete) error {
	if (*delete == model.DatasetVersionDelete{}) {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	db := d.db(ctx)
	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if delete.DatasetId != "" {
		querySql += " and dataset_id = ? "
		params = append(params, delete.DatasetId)
	}

	if delete.Version != "" {
		querySql += " and version = ? "
		params = append(params, delete.Version)
	}

	if len(params) == 0 {
		return errors.Errorf(nil, errors.ErrorDBSelectParamsEmpty)
	}

	res := db.Where(querySql, params...).Delete(&model.DatasetVersion{})
	if res.Error != nil {
		return errors.Errorf(nil, errors.ErrorDBDeleteFailed)
	}

	return nil
}

func (d *datasetDao) UpdateDatasetVersionSelective(ctx context.Context, version *model.DatasetVersion) error {
	if version.DatasetId == "" || version.Version == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	db := d.db(ctx)
	res := db.Where("dataset_id = ? and version = ?", version.DatasetId, version.Version).Updates(version)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *datasetDao) ListDatasetVersionLatestVersion(ctx context.Context, datasetIds []string) (map[string]int64, error) {
	db := d.db(ctx)
	versions := make([]*model.DatasetVersion, 0)

	res := db.Table(model.DatasetVersion{}.TableName()).
		Select("dataset_id, max(version_int) as version_int").
		Group("dataset_id").Where("dataset_id in ?", datasetIds).Find(&versions)
	if res.Error != nil {
		return nil, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	m := map[string]int64{}
	for _, i := range versions {
		m[i.DatasetId] = i.VersionInt
	}

	return m, nil
}

func (d *datasetDao) ListDatasetAccess(ctx context.Context, query *model.DatasetAccessQuery) ([]*model.DatasetAccess, error) {
	db := d.db(ctx)
	accesses := make([]*model.DatasetAccess, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if query.DatasetId != "" {
		querySql += " and dataset_id = ? "
		params = append(params, query.DatasetId)
	}

	db = db.Model(&model.DatasetAccess{}).Where(querySql, params...)

	res := db.Find(&accesses)
	if res.Error != nil {
		return nil, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return accesses, nil
}

func (d *datasetDao) CreateDatasetAccess(ctx context.Context, version *model.DatasetAccess) error {
	db := d.db(ctx)
	res := db.Create(version)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *datasetDao) DeleteDatasetAccess(ctx context.Context, delete *model.DatasetAccessDelete) error {
	if *delete == (model.DatasetAccessDelete{}) {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	db := d.db(ctx)
	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if delete.DatasetId != "" {
		querySql += " and dataset_id = ? "
		params = append(params, delete.DatasetId)
	}

	if delete.SpaceId != "" {
		querySql += " and space_id = ? "
		params = append(params, delete.SpaceId)
	}

	if len(params) == 0 {
		return errors.Errorf(nil, errors.ErrorDBSelectParamsEmpty)
	}
	res := db.Where(querySql, params...).Delete(&model.DatasetAccess{})
	if res.Error != nil {
		return errors.Errorf(nil, errors.ErrorDBDeleteFailed)
	}

	return nil
}

func (d *datasetDao) UpdateDatasetAccessSelective(ctx context.Context, access *model.DatasetAccess) error {
	if access.DatasetId == "" || access.SpaceId == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	db := d.db(ctx)
	res := db.Where("dataset_id = ? and space_id = ?", access.DatasetId, access.SpaceId).Updates(access)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *datasetDao) ListDatasetVersionAccess(ctx context.Context, query *model.DatasetVersionAccessQuery) ([]*model.DatasetVersionAccess, error) {
	db := d.db(ctx)
	accesses := make([]*model.DatasetVersionAccess, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if query.DatasetId != "" {
		querySql += " and dataset_id = ? "
		params = append(params, query.DatasetId)
	}

	if query.Version != "" {
		querySql += " and version = ? "
		params = append(params, query.Version)
	}

	if query.SpaceId != "" {
		querySql += " and space_id = ? "
		params = append(params, query.SpaceId)
	}

	db = db.Model(&model.DatasetVersionAccess{}).Where(querySql, params...)

	res := db.Find(&accesses)
	if res.Error != nil {
		return nil, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return accesses, nil
}

func (d *datasetDao) CreateDatasetVersionAccess(ctx context.Context, versionAccess *model.DatasetVersionAccess) error {
	db := d.db(ctx)
	res := db.Create(versionAccess)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *datasetDao) DeleteDatasetVersionAccess(ctx context.Context, delete *model.DatasetVersionAccessDelete) error {
	if *delete == (model.DatasetVersionAccessDelete{}) {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	db := d.db(ctx)
	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if delete.DatasetId != "" {
		querySql += " and dataset_id = ? "
		params = append(params, delete.DatasetId)
	}

	if delete.Version != "" {
		querySql += " and version = ? "
		params = append(params, delete.Version)
	}

	if delete.SpaceId != "" {
		querySql += " and space_id = ? "
		params = append(params, delete.SpaceId)
	}

	if len(params) == 0 {
		return errors.Errorf(nil, errors.ErrorDBSelectParamsEmpty)
	}
	res := db.Where(querySql, params...).Delete(&model.DatasetVersionAccess{})
	if res.Error != nil {
		return errors.Errorf(nil, errors.ErrorDBDeleteFailed)
	}

	return nil
}

func (d *datasetDao) ListDatasetVersionAccessLatestVersion(ctx context.Context, ids []model.DatasetAccessId) (map[model.DatasetAccessId]int64, error) {
	db := d.db(ctx)
	versions := make([]*model.DatasetVersionAccess, 0)

	inSql := ""
	for i, j := range ids {
		if i != 0 {
			inSql += ","
		}
		inSql += fmt.Sprintf("('%s', '%s')", j.DatasetId, j.SpaceId)
	}

	res := db.Raw(fmt.Sprintf("select dataset_id, space_id, max(version_int) as version_int from dataset_version_access "+
		"where (dataset_id, space_id) in (%s) and deleted_at = 0 group by dataset_id, space_id", inSql)).Scan(&versions)
	if res.Error != nil {
		return nil, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	m := map[model.DatasetAccessId]int64{}
	for _, i := range versions {
		m[model.DatasetAccessId{
			DatasetId: i.DatasetId,
			SpaceId:   i.SpaceId,
		}] = i.VersionInt
	}

	return m, nil
}
