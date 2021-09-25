package dao

import (
	"context"
	stderrors "errors"
	"fmt"
	"server/base-server/internal/data/dao/model"
	"server/base-server/internal/data/influxdb"
	"server/common/errors"
	"server/common/utils"
	"time"

	"server/common/log"

	"gorm.io/gorm"
)

type TrainJobDao interface {
	//网关层生成任务信息
	CreateTrainJob(ctx context.Context, trainJob *model.TrainJob) error
	//网关层查询任务信息
	GetTrainJob(ctx context.Context, id string) (*model.TrainJob, error)
	//网关层查询训练任务名称是否重复
	GetTrainJobByName(ctx context.Context, jobName string, userId string, workspaceId string) (*model.TrainJob, error)
	//网关层查询任务列表
	GetTrainJobList(ctx context.Context, query *model.TrainJobListQuery) ([]*model.TrainJob, int64, error)
	//网关层更新user对任务的操作记录
	UpdateTrainJobOperation(jobId string, operation string) error
	//网关层更新来自taskset的任务状态信息
	UpdateTrainJob(ctx context.Context, trainJob *model.TrainJob) error
	//网关层删除任务（软删除）
	DeleteTrainJob(ctx context.Context, id string) error

	//网关层创建模板信息
	CreateTrainJobTemplate(ctx context.Context, trainJobTemplate *model.TrainJobTemplate) error
	//网关层查询单个模板信息
	GetTrainJobTemplate(ctx context.Context, id string) (*model.TrainJobTemplate, error)
	//网关层查询训练模板名称是否重复
	GetTrainJobTemplateByName(ctx context.Context, jobName string, userId string, workspaceId string) (*model.TrainJobTemplate, error)
	//网关层查询模板列表
	GetTrainJobTemplateList(ctx context.Context, query *model.TrainJobTemPlateListQuery) ([]*model.TrainJobTemplate, int64, error)
	//网关层编辑任务模板
	UpdateTrainJobTemplate(ctx context.Context, trainJobTemplate *model.TrainJobTemplate) error
	//网关层删除任务模板（软删除）
	DeleteTrainJobTemplate(userId string, ids []string) error
	//获取训练任务事件
	GetTrainJobEvents(jobId string) error
}

type trainJobDao struct {
	log      *log.Helper
	db       *gorm.DB
	influxdb influxdb.Influxdb
}

func NewTrainJobDao(db *gorm.DB, influxdb influxdb.Influxdb, logger log.Logger) TrainJobDao {
	return &trainJobDao{
		log:      log.NewHelper("TrainJobDao", logger),
		db:       db,
		influxdb: influxdb,
	}
}

func (d *trainJobDao) GetTrainJob(ctx context.Context, id string) (*model.TrainJob, error) {
	trainJob := &model.TrainJob{}
	res := d.db.First(trainJob, "id = ?", id)
	if res.Error != nil {
		if stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFindEmpty)
		} else {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFirstFailed)
		}
	}
	return trainJob, nil
}

func (d *trainJobDao) GetTrainJobByName(ctx context.Context, jobName string, userId string, workspaceId string) (*model.TrainJob, error) {
	trainJob := &model.TrainJob{}
	db := d.db.Where("1=1 and name = ? and user_id = ? and workspace_id = ? and deleted_at = 0 ", jobName, userId, workspaceId).Find(&trainJob)
	var totalSize int64
	res := db.Model(&model.TrainJob{}).Count(&totalSize)
	if res.Error != nil {
		return trainJob, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	if totalSize != 0 {
		return trainJob, errors.Errorf(nil, errors.ErrorJobUniqueIndexConflict)
	}
	return nil, nil
}

func (d *trainJobDao) GetTrainJobList(ctx context.Context, query *model.TrainJobListQuery) ([]*model.TrainJob, int64, error) {
	db := d.db
	trainJobs := make([]*model.TrainJob, 0)

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

	if query.SearchKey != "" {
		querySql += " and name like ?"
		params = append(params, "%"+query.SearchKey+"%")
	}

	if query.UserId != "" {
		querySql += " and user_id = ? "
		params = append(params, query.UserId)
	}

	if query.PayStatus != 0 {
		querySql += " and pay_status = ? "
		params = append(params, query.PayStatus)
	}

	if query.WorkspaceId != "" {
		querySql += " and workspace_id = ? "
		params = append(params, query.WorkspaceId)
	}

	if len(query.Ids) != 0 {
		querySql += " and id in ? "
		params = append(params, query.Ids)
	}

	if query.UserNameLike != "" {
		joinSql := "INNER JOIN user ON train_job.user_id = user.id"
		querySql += " and user.full_name like ?"
		params = append(params, query.UserNameLike+"%")
		db = db.Joins(joinSql).Where(querySql, params...)
	} else {
		db = db.Where(querySql, params...)
	}

	var totalSize int64
	res := db.Model(&model.TrainJob{}).Count(&totalSize)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	if query.PageIndex != 0 {
		db = db.Limit(query.PageSize).Offset((query.PageIndex - 1) * query.PageSize)
	}

	sortBy := "train_job.created_at"
	orderBy := "desc"
	if query.SortBy != "" {
		sortBy = utils.CamelToSnake(query.SortBy)
	}

	if query.OrderBy != "" {
		sortBy = query.OrderBy
	}

	db = db.Order(fmt.Sprintf("%s %s", sortBy, orderBy))

	res = db.Select("train_job.*").Find(&trainJobs)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return trainJobs, totalSize, nil

}

func (d *trainJobDao) CreateTrainJob(ctx context.Context, trainJob *model.TrainJob) error {
	res := d.db.Create(trainJob)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *trainJobDao) UpdateTrainJobOperation(jobId string, operation string) error {
	var trainJob model.TrainJob
	trainJob.Operation = operation
	if err := d.db.Model(trainJob).Where("id = ? ", jobId).Updates(trainJob).Error; err != nil {
		return errors.Errorf(err, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *trainJobDao) UpdateTrainJob(ctx context.Context, trainJob *model.TrainJob) error {
	if trainJob.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := d.db.Where("id = ?", trainJob.Id).Updates(trainJob)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *trainJobDao) DeleteTrainJob(ctx context.Context, id string) error {
	if id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := d.db.Where("id = ? ", id).Delete(&model.TrainJob{})
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBDeleteFailed)
	}

	return nil
}

func (d *trainJobDao) CreateTrainJobTemplate(ctx context.Context, trainJobTemplate *model.TrainJobTemplate) error {
	res := d.db.Create(trainJobTemplate)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *trainJobDao) GetTrainJobTemplate(ctx context.Context, id string) (*model.TrainJobTemplate, error) {
	trainJobTemplate := &model.TrainJobTemplate{}
	res := d.db.First(trainJobTemplate, "id = ?", id)

	if res.Error != nil {
		if stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFindEmpty)
		} else {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFirstFailed)
		}
	}
	return trainJobTemplate, nil
}

func (d *trainJobDao) GetTrainJobTemplateByName(ctx context.Context, jobName string, userId string, workspaceId string) (*model.TrainJobTemplate, error) {
	trainJobTemplate := &model.TrainJobTemplate{}
	db := d.db.Where("name = ? and user_id = ? and workspace_id = ? and deleted_at = 0 ", jobName, userId, workspaceId).Find(&trainJobTemplate)
	var totalSize int64
	res := db.Model(&model.TrainJobTemplate{}).Count(&totalSize)
	if res.Error != nil {
		return trainJobTemplate, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	if totalSize != 0 {
		return trainJobTemplate, errors.Errorf(nil, errors.ErrorJobUniqueIndexConflict)
	}
	return nil, nil
}

func (d *trainJobDao) GetTrainJobTemplateList(ctx context.Context, query *model.TrainJobTemPlateListQuery) ([]*model.TrainJobTemplate, int64, error) {
	db := d.db
	trainJobTemplates := make([]*model.TrainJobTemplate, 0)

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

	res = db.Find(&trainJobTemplates)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return trainJobTemplates, totalSize, nil
}

func (d *trainJobDao) UpdateTrainJobTemplate(ctx context.Context, trainJobTemplate *model.TrainJobTemplate) error {
	if trainJobTemplate.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	res := d.db.Where("id = ?", trainJobTemplate.Id).Updates(trainJobTemplate)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *trainJobDao) DeleteTrainJobTemplate(userId string, ids []string) error {
	if userId == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	if len(ids) == 0 {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	res := d.db.Where("user_id = ? and id in ? ", userId, ids).Delete(&model.TrainJobTemplate{})
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBDeleteFailed)
	}
	return nil
}

func (d *trainJobDao) GetTrainJobEvents(jobId string) error {

	query := fmt.Sprintf("select object_name, reason, message from events where object_name =~ /^%s/ and kind = 'Pod'", jobId)

	res, err := d.influxdb.Query(query)

	if err != nil {
		return errors.Errorf(err, errors.ErroInfluxdbFindFailed)
	}

	fmt.Println(res)
	return nil
}
