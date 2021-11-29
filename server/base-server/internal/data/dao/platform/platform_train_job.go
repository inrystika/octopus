package platform

import (
	"context"
	stderrors "errors"
	"fmt"
	m "server/base-server/internal/data/dao/model"
	model "server/base-server/internal/data/dao/model/platform"
	"server/common/errors"
	"server/common/utils"
	"time"

	"server/common/log"

	"gorm.io/gorm"
)

type PlatformTrainJobDao interface {
	//网关层生成任务信息
	CreateTrainJob(ctx context.Context, trainJob *model.PlatformTrainJob) error
	//网关层查询任务信息
	GetTrainJob(ctx context.Context, id string) (*model.PlatformTrainJob, error)
	//网关层查询训练任务名称是否重复
	GetTrainJobByName(ctx context.Context, jobName string, platformId string) (*model.PlatformTrainJob, error)
	//网关层查询任务列表
	GetTrainJobList(ctx context.Context, query *model.PlatformTrainJobListQuery) ([]*model.PlatformTrainJob, int64, error)
	//网关层更新user对任务的操作记录
	UpdateTrainJobOperation(jobId string, operation string) error
	//网关层更新来自taskset的任务状态信息
	UpdateTrainJob(ctx context.Context, trainJob *model.PlatformTrainJob) error
	//网关层删除任务（软删除）
	DeleteTrainJob(ctx context.Context, id string) error
	//网关层获取任务统计信息
	TrainJobStastics(ctx context.Context, query *model.TrainJobStastics) (*model.TrainJobStasticsReply, error)
}

type platformTrainJobDao struct {
	log *log.Helper
	db  *gorm.DB
}

func NewPlatformTrainJobDao(db *gorm.DB, logger log.Logger) PlatformTrainJobDao {
	return &platformTrainJobDao{
		log: log.NewHelper("PlatformTrainJobDao", logger),
		db:  db,
	}
}

func (d *platformTrainJobDao) GetTrainJob(ctx context.Context, id string) (*model.PlatformTrainJob, error) {
	trainJob := &model.PlatformTrainJob{}
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

func (d *platformTrainJobDao) GetTrainJobByName(ctx context.Context, jobName string, platformId string) (*model.PlatformTrainJob, error) {
	trainJob := &model.PlatformTrainJob{}
	db := d.db.Where("1=1 and name = ? and platform_id = ? and deleted_at = 0 ", jobName, platformId).Find(&trainJob)
	var totalSize int64
	res := db.Model(&model.PlatformTrainJob{}).Count(&totalSize)
	if res.Error != nil {
		return trainJob, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	if totalSize != 0 {
		return trainJob, errors.Errorf(nil, errors.ErrorJobUniqueIndexConflict)
	}
	return nil, nil
}

func (d *platformTrainJobDao) GetTrainJobList(ctx context.Context, query *model.PlatformTrainJobListQuery) ([]*model.PlatformTrainJob, int64, error) {
	db := d.db
	trainJobs := make([]*model.PlatformTrainJob, 0)

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

	if query.PlatformId != "" {
		querySql += " and platform_id = ? "
		params = append(params, query.PlatformId)
	}

	if len(query.Ids) != 0 {
		querySql += " and id in ? "
		params = append(params, query.Ids)
	}

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Model(&model.PlatformTrainJob{}).Count(&totalSize)
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

	res = db.Find(&trainJobs)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return trainJobs, totalSize, nil

}

func (d *platformTrainJobDao) CreateTrainJob(ctx context.Context, trainJob *model.PlatformTrainJob) error {
	res := d.db.Create(trainJob)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *platformTrainJobDao) UpdateTrainJobOperation(jobId string, operation string) error {
	var trainJob model.PlatformTrainJob
	trainJob.Operation = operation
	if err := d.db.Model(trainJob).Where("id = ? ", jobId).Updates(trainJob).Error; err != nil {
		return errors.Errorf(err, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *platformTrainJobDao) UpdateTrainJob(ctx context.Context, trainJob *model.PlatformTrainJob) error {
	if trainJob.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := d.db.Where("id = ?", trainJob.Id).Updates(trainJob)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *platformTrainJobDao) DeleteTrainJob(ctx context.Context, id string) error {
	if id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := d.db.Where("id = ? ", id).Delete(&model.PlatformTrainJob{})
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBDeleteFailed)
	}

	return nil
}

func (d *platformTrainJobDao) TrainJobStastics(ctx context.Context, query *model.TrainJobStastics) (*model.TrainJobStasticsReply, error) {

	db := d.db
	trainJobStasticsReply := &model.TrainJobStasticsReply{}
	var size, totalSize, succeededSize, failedSize, stoppedSize, runningSize, waitingSize int64 = 0, 0, 0, 0, 0, 0, 0

	baseQuerySql := "1 = 1"
	baseParams := make([]interface{}, 0)
	succeedParams := make([]interface{}, 0)
	failedParams := make([]interface{}, 0)
	stoppedParams := make([]interface{}, 0)
	runningParams := make([]interface{}, 0)
	waitingParams := make([]interface{}, 0)

	if query.CreatedAtGte != 0 {
		baseQuerySql += " and created_at >= ? "
		baseParams = append(baseParams, time.Unix(query.CreatedAtGte, 0))
	}

	if query.CreatedAtLt != 0 {
		baseQuerySql += " and created_at < ? "
		baseParams = append(baseParams, time.Unix(query.CreatedAtLt, 0))
	}

	succeededQuerySql := baseQuerySql + " and status = ? "
	succeedParams = append(baseParams, "succeeded")
	failedQuerySql := baseQuerySql + " and status = ? "
	failedParams = append(baseParams, "failed")
	stoppedQuerySql := baseQuerySql + " and status = ? "
	stoppedParams = append(baseParams, "stopped")
	runningQuerySql := baseQuerySql + " and status = ? "
	runningParams = append(baseParams, "running")
	waitingQuerySql := baseQuerySql + " and status = ? or status = ?"
	waitingParams = append(baseParams, "pending")
	waitingParams = append(waitingParams, "preparing")

	res := db.Where(baseQuerySql, baseParams...).Model(&m.TrainJob{}).Count(&size)
	if res.Error != nil {
		return trainJobStasticsReply, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	totalSize += size

	res = db.Where(succeededQuerySql, succeedParams...).Model(&m.TrainJob{}).Count(&size)
	if res.Error != nil {
		return trainJobStasticsReply, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	succeededSize += size

	res = db.Where(failedQuerySql, failedParams...).Model(&m.TrainJob{}).Count(&size)
	if res.Error != nil {
		return trainJobStasticsReply, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	failedSize += size

	res = db.Where(stoppedQuerySql, stoppedParams...).Model(&m.TrainJob{}).Count(&size)
	if res.Error != nil {
		return trainJobStasticsReply, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	stoppedSize += size

	res = db.Where(runningQuerySql, runningParams...).Model(&m.TrainJob{}).Count(&size)
	if res.Error != nil {
		return trainJobStasticsReply, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	runningSize += size

	res = db.Where(waitingQuerySql, waitingParams...).Model(&m.TrainJob{}).Count(&size)
	if res.Error != nil {
		return trainJobStasticsReply, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	waitingSize += size

	res = db.Where(baseQuerySql, baseParams...).Model(&model.PlatformTrainJob{}).Count(&size)
	if res.Error != nil {
		return trainJobStasticsReply, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	totalSize += size

	res = db.Where(succeededQuerySql, succeedParams...).Model(&model.PlatformTrainJob{}).Count(&size)
	if res.Error != nil {
		return trainJobStasticsReply, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	succeededSize += size

	res = db.Where(failedQuerySql, failedParams...).Model(&model.PlatformTrainJob{}).Count(&size)
	if res.Error != nil {
		return trainJobStasticsReply, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	failedSize += size

	res = db.Where(stoppedQuerySql, stoppedParams...).Model(&model.PlatformTrainJob{}).Count(&size)
	if res.Error != nil {
		return trainJobStasticsReply, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	stoppedSize += size

	res = db.Where(runningQuerySql, runningParams...).Model(&model.PlatformTrainJob{}).Count(&size)
	if res.Error != nil {
		return trainJobStasticsReply, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	runningSize += size

	res = db.Where(waitingQuerySql, waitingParams...).Model(&model.PlatformTrainJob{}).Count(&size)
	if res.Error != nil {
		return trainJobStasticsReply, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	waitingSize += size

	trainJobStasticsReply.TotalSize = totalSize
	trainJobStasticsReply.SucceededSize = succeededSize
	trainJobStasticsReply.FailedSize = failedSize
	trainJobStasticsReply.StoppedSize = stoppedSize
	trainJobStasticsReply.RunningSize = runningSize
	trainJobStasticsReply.WaitingSize = waitingSize

	return trainJobStasticsReply, nil
}
