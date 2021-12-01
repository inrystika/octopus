package dao

import (
	"context"
	"encoding/json"
	stderrors "errors"
	"fmt"
	"server/base-server/internal/data/dao/model"
	"server/base-server/internal/data/influxdb"
	"server/common/errors"
	"server/common/log"
	"server/common/utils"
	"time"

	"gorm.io/gorm"
)

type modelDeployDao struct {
	log      *log.Helper
	db       *gorm.DB
	influxdb influxdb.Influxdb
}

func NewModelDeployDao(db *gorm.DB, influxdb influxdb.Influxdb, logger log.Logger) ModelDeployDao {
	return &modelDeployDao{
		log:      log.NewHelper("modelDeployDao", logger),
		db:       db,
		influxdb: influxdb,
	}
}

type ModelDeployDao interface {
	//生成部署服务信息
	CreateModelDeployService(ctx context.Context, modelDeploy *model.ModelDeploy) error
	//查询部署服务信息
	GetModelDeployService(ctx context.Context, id string) (*model.ModelDeploy, error)
	//查询训练任务名称是否重复
	GetModelDeployServiceByName(ctx context.Context, jobName string, userId string, workspaceId string) (*model.ModelDeploy, error)
	//查询任务列表
	GetModelDeployServiceList(ctx context.Context, query *model.ModelDeployListQuery) ([]*model.ModelDeploy, int64, error)
	//更新部署服务信息
	UpdateModelDeployService(ctx context.Context, trainJob *model.ModelDeploy) error
	//删除部署服务信息（软删除）
	DeleteModelDeployService(ctx context.Context, id string) error
	//查询部署服务事件列表
	GetModelDeployEvents(deployEventQuery *model.DeployEventQuery) ([]*model.ModelDeployEvent, int64, error)
}

func (d *modelDeployDao) CreateModelDeployService(ctx context.Context, ModelDeploy *model.ModelDeploy) error {
	res := d.db.Create(ModelDeploy)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *modelDeployDao) GetModelDeployService(ctx context.Context, id string) (*model.ModelDeploy, error) {
	modelDeploy := &model.ModelDeploy{}
	res := d.db.First(modelDeploy, "id = ?", id)

	if res.Error != nil {
		if stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFindEmpty)
		} else {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFirstFailed)
		}
	}
	return modelDeploy, nil
}

func (d *modelDeployDao) GetModelDeployServiceByName(ctx context.Context, jobName string, userId string, workspaceId string) (*model.ModelDeploy, error) {
	modelDeploy := &model.ModelDeploy{}
	db := d.db.Where("name = ? and user_id = ? and workspace_id = ? and deleted_at = 0 ", jobName, userId, workspaceId).Find(&modelDeploy)
	var totalSize int64
	res := db.Model(&model.ModelDeploy{}).Count(&totalSize)
	if res.Error != nil {
		return modelDeploy, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	if totalSize != 0 {
		return modelDeploy, errors.Errorf(nil, errors.ErrorJobUniqueIndexConflict)
	}
	return nil, nil
}

func (d *modelDeployDao) GetModelDeployServiceList(ctx context.Context, query *model.ModelDeployListQuery) ([]*model.ModelDeploy, int64, error) {
	db := d.db
	modelDeployments := make([]*model.ModelDeploy, 0)

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

	if query.Status != "" {
		querySql += " and status = ? "
		params = append(params, query.Status)
	}

	if query.UserId != "" {
		querySql += " and user_id = ? "
		params = append(params, query.UserId)
	}

	if query.WorkspaceId != "" {
		querySql += " and workspace_id = ? "
		params = append(params, query.WorkspaceId)
	}

	if query.SearchKey != "" {
		querySql += " and name like ?"
		params = append(params, "%"+query.SearchKey+"%")
	}

	if len(query.Ids) != 0 {
		querySql += " and id in ? "
		params = append(params, query.Ids)
	}

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Model(&model.TrainJobTemplate{}).Count(&totalSize)
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
		sortBy = query.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", sortBy, orderBy))

	res = db.Find(&modelDeployments)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return modelDeployments, totalSize, nil
}

func (d *modelDeployDao) UpdateModelDeployService(ctx context.Context, modelDeploy *model.ModelDeploy) error {
	if modelDeploy.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	res := d.db.Where("id = ?", modelDeploy.Id).Updates(modelDeploy)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *modelDeployDao) DeleteModelDeployService(ctx context.Context, id string) error {
	if id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := d.db.Where("id = ? ", id).Delete(&model.ModelDeploy{})
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBDeleteFailed)
	}

	return nil
}

func (d *modelDeployDao) GetModelDeployEvents(deployEventQuery *model.DeployEventQuery) ([]*model.ModelDeployEvent, int64, error) {

	keyName := "object_name"
	keyReason := "reason"
	keyMessage := "message"

	PageIndex := deployEventQuery.PageIndex
	PageSize := deployEventQuery.PageSize
	events := make([]*model.ModelDeployEvent, 0)

	objectName := ""
	if !deployEventQuery.IsMain {
		objectName = fmt.Sprintf("%s-default-0-model", deployEventQuery.Id)
	} else {
		objectName = deployEventQuery.Id
	}

	countQuery := ""
	if !deployEventQuery.IsMain {
		countQuery = fmt.Sprintf("SELECT COUNT(%s) FROM octopus..events where object_name =~ /^%s/", keyMessage, objectName)
	} else {
		countQuery = fmt.Sprintf("SELECT COUNT(%s) FROM octopus..events where object_name = '%s'", keyMessage, objectName)
	}

	res, err := d.influxdb.Query(countQuery)

	if err != nil {
		return events, 0, errors.Errorf(err, errors.ErroInfluxdbFindFailed)
	}

	if len(res) == 0 || len(res[0].Series) == 0 || len(res[0].Series[0].Values) == 0 || len(res[0].Series[0].Values[0]) < 2 {
		return events, 0, errors.Errorf(err, errors.ErroInfluxdbFindFailed)
	}

	totalSize, err := res[0].Series[0].Values[0][1].(json.Number).Int64()
	if err != nil {
		return events, 0, errors.Errorf(err, errors.ErroInfluxdbFindFailed)
	}

	query := ""
	if !deployEventQuery.IsMain {
		query = fmt.Sprintf("select %s, %s, %s from octopus..events where object_name =~ /^%s/ LIMIT %d OFFSET %d",
			keyName, keyReason, keyMessage, objectName, PageSize, (PageIndex-1)*PageSize)

	} else {
		query = fmt.Sprintf("select %s, %s, %s from octopus..events where object_name = '%s' LIMIT %d OFFSET %d",
			keyName, keyReason, keyMessage, objectName, PageSize, (PageIndex-1)*PageSize)
	}

	res, err = d.influxdb.Query(query)

	if err != nil {
		return events, 0, errors.Errorf(err, errors.ErroInfluxdbFindFailed)
	}

	for _, row := range res[0].Series[0].Values {

		event := &model.ModelDeployEvent{}
		event.Timestamp = row[0].(string)
		event.Name = row[1].(string)
		event.Reason = row[2].(string)
		event.Message = row[3].(string)
		events = append(events, event)
	}

	return events, totalSize, nil
}
