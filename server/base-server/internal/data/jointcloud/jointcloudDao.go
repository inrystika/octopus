package jointcloud

import (
	"context"
	stderrors "errors"
	"fmt"
	"server/common/errors"
	"server/common/utils"
	"time"

	"gorm.io/gorm"
)

type JointcloudDao interface {
	//网关层生成任务信息
	CreateTrainJob(ctx context.Context, trainJob *TrainJob) error
	//网关层查询任务信息
	GetTrainJob(ctx context.Context, id string) (*TrainJob, error)
	//网关层查询训练任务名称是否重复
	GetTrainJobByName(ctx context.Context, jobName, userId, workspaceId string) (*TrainJob, error)
	//网关层查询任务列表
	GetTrainJobList(ctx context.Context, query *TrainJobListQuery) ([]*TrainJob, int64, error)
	//网关层更新user对任务的操作记录
	UpdateTrainJobOperation(jobId string, operation string) error
}

type jointcloudDao struct {
	db *gorm.DB
}

func NewJointcloudDao(db *gorm.DB) JointcloudDao {
	return &jointcloudDao{
		db: db,
	}
}

func (d *jointcloudDao) CreateTrainJob(ctx context.Context, trainJob *TrainJob) error {
	res := d.db.Create(trainJob)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *jointcloudDao) UpdateTrainJobOperation(jobId string, operation string) error {
	var trainJob TrainJob
	trainJob.Operation = operation
	if err := d.db.Model(trainJob).Where("id = ? ", jobId).Updates(trainJob).Error; err != nil {
		return errors.Errorf(err, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *jointcloudDao) GetTrainJob(ctx context.Context, id string) (*TrainJob, error) {
	trainJob := &TrainJob{}
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

func (d *jointcloudDao) GetTrainJobByName(ctx context.Context, jobName, userId, workspaceId string) (*TrainJob, error) {
	trainJob := &TrainJob{}
	db := d.db.Where("1=1 and task_name = ? and user_id = ? and workspace_id = ? and deleted_at = 0 ", jobName, userId, workspaceId).Find(&trainJob)
	var totalSize int64
	res := db.Model(&TrainJob{}).Count(&totalSize)
	if res.Error != nil {
		return trainJob, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}
	if totalSize != 0 {
		return trainJob, errors.Errorf(nil, errors.ErrorJobUniqueIndexConflict)
	}
	return nil, nil
}

func (d *jointcloudDao) GetTrainJobList(ctx context.Context, query *TrainJobListQuery) ([]*TrainJob, int64, error) {
	db := d.db
	trainJobs := make([]*TrainJob, 0)

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
		querySql += " and task_name like ?"
		params = append(params, "%"+query.SearchKey+"%")
	}

	if query.WorkspaceId != "" {
		querySql += " and workspace_id = ? "
		params = append(params, query.WorkspaceId)
	}

	if query.UserId != "" {
		querySql += " and user_id = ? "
		params = append(params, query.UserId)
	}

	if len(query.Ids) != 0 {
		querySql += " and id in ? "
		params = append(params, query.Ids)
	}

	db = db.Where(querySql, params...)

	var totalSize int64
	res := db.Model(&TrainJob{}).Count(&totalSize)
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

	res = db.Find(&trainJobs)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return trainJobs, totalSize, nil

}
