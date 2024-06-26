package dao

import (
	"context"
	"encoding/json"
	stderrors "errors"
	"fmt"
	"server/base-server/internal/data/dao/model"
	"server/base-server/internal/data/influxdb"
	v1 "server/common/api/v1"
	"server/common/errors"
	"server/common/utils"
	"time"

	"server/common/log"

	"gorm.io/gorm"
)

type DevelopDao interface {
	CreateNotebook(ctx context.Context, notebook *model.Notebook) error
	GetNotebook(ctx context.Context, id string) (*model.Notebook, error)
	UpdateNotebookSelective(ctx context.Context, notebook *model.Notebook) error
	UpdateNotebookSelectiveByJobId(ctx context.Context, notebook *model.Notebook) error
	UpdateNotebookByJobIdOnNotCompleted(ctx context.Context, notebook *model.Notebook) error
	DeleteNotebook(ctx context.Context, id string) error
	ListNotebook(ctx context.Context, query *model.NotebookQuery) ([]*model.Notebook, int64, error)

	CreateNotebookJob(ctx context.Context, notebookJob *model.NotebookJob) error
	GetNotebookJob(ctx context.Context, id string) (*model.NotebookJob, error)
	UpdateNotebookJobSelective(ctx context.Context, notebookJob *model.NotebookJob) error
	UpdateNotebookJobOnNotCompleted(ctx context.Context, notebookJob *model.NotebookJob) error
	DeleteNotebookJobByNbId(ctx context.Context, notebookId string) error
	ListNotebookJob(ctx context.Context, query *model.NotebookJobQuery) ([]*model.NotebookJob, error)
	//获取Notebook事件
	GetNotebookEvents(notebookEventQuery *model.NotebookEventQuery) ([]*model.NotebookEvent, int64, error)

	CreateNotebookEventRecord(ctx context.Context, r *model.NotebookEventRecord) error
	ListNotebookEventRecord(ctx context.Context, query *model.NotebookEventRecordQuery) ([]*model.NotebookEventRecord, int64, error)
}

type developDao struct {
	log      *log.Helper
	db       *gorm.DB
	influxdb influxdb.Influxdb
}

func NewDevelopDao(db *gorm.DB, influxdb influxdb.Influxdb, logger log.Logger) DevelopDao {
	return &developDao{
		log:      log.NewHelper("DevelopDao", logger),
		db:       db,
		influxdb: influxdb,
	}
}

func (d *developDao) CreateNotebook(ctx context.Context, notebook *model.Notebook) error {
	res := d.db.Create(notebook)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *developDao) GetNotebook(ctx context.Context, id string) (*model.Notebook, error) {
	nb := &model.Notebook{}
	res := d.db.First(nb, "id = ?", id)

	if res.Error != nil {
		if stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFindEmpty)
		} else {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFirstFailed)
		}
	}
	return nb, nil
}

func (d *developDao) UpdateNotebookSelective(ctx context.Context, notebook *model.Notebook) error {
	if notebook.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := d.db.Where("id = ?", notebook.Id).Updates(notebook)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *developDao) UpdateNotebookSelectiveByJobId(ctx context.Context, notebook *model.Notebook) error {
	if notebook.NotebookJobId == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := d.db.Where("notebook_job_id = ?", notebook.NotebookJobId).Updates(notebook)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *developDao) UpdateNotebookByJobIdOnNotCompleted(ctx context.Context, notebook *model.Notebook) error {
	if notebook.NotebookJobId == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := d.db.Where("notebook_job_id = ? and status not in ?", notebook.NotebookJobId, utils.CompletedStates).Updates(notebook)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *developDao) DeleteNotebook(ctx context.Context, id string) error {
	if id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	res := d.db.Where("id = ?", id).Delete(&model.Notebook{})
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBDeleteFailed)
	}

	return nil
}

func (d *developDao) ListNotebook(ctx context.Context, query *model.NotebookQuery) ([]*model.Notebook, int64, error) {
	notebooks := make([]*model.Notebook, 0)

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

	if query.WorkspaceId != "" {
		querySql += " and workspace_id = ? "
		params = append(params, query.WorkspaceId)
	}

	if len(query.Ids) != 0 {
		querySql += " and id in ? "
		params = append(params, query.Ids)
	}

	if query.Name != "" {
		querySql += " and name = ? "
		params = append(params, query.Name)
	}

	db := d.db.Where(querySql, params...)

	var totalSize int64
	res := db.Model(&model.Notebook{}).Count(&totalSize)
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

	res = db.Find(&notebooks)
	if res.Error != nil {
		return nil, 0, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return notebooks, totalSize, nil
}

func (d *developDao) CreateNotebookJob(ctx context.Context, notebookJob *model.NotebookJob) error {
	res := d.db.Create(notebookJob)
	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBCreateFailed)
	}
	return nil
}

func (d *developDao) GetNotebookJob(ctx context.Context, id string) (*model.NotebookJob, error) {
	nb := &model.NotebookJob{}
	res := d.db.First(nb, "id = ?", id)

	if res.Error != nil {
		if stderrors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFindEmpty)
		} else {
			return nil, errors.Errorf(res.Error, errors.ErrorDBFirstFailed)
		}
	}
	return nb, nil
}

func (d *developDao) UpdateNotebookJobSelective(ctx context.Context, notebookJob *model.NotebookJob) error {
	if notebookJob.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := d.db.Where("id = ?", notebookJob.Id).Updates(notebookJob)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *developDao) UpdateNotebookJobOnNotCompleted(ctx context.Context, notebookJob *model.NotebookJob) error {
	if notebookJob.Id == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}
	res := d.db.Where("id = ? and status not in ?", notebookJob.Id, utils.CompletedStates).Updates(notebookJob)

	if res.Error != nil {
		return errors.Errorf(res.Error, errors.ErrorDBUpdateFailed)
	}
	return nil
}

func (d *developDao) DeleteNotebookJobByNbId(ctx context.Context, nbId string) error {
	if nbId == "" {
		return errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	res := d.db.Where("notebook_id = ?", nbId).Delete(&model.NotebookJob{})
	if res.Error != nil {
		return errors.Errorf(nil, errors.ErrorDBDeleteFailed)
	}

	return nil
}

func (d *developDao) ListNotebookJob(ctx context.Context, query *model.NotebookJobQuery) ([]*model.NotebookJob, error) {
	notebookJobs := make([]*model.NotebookJob, 0)

	querySql := "1 = 1"
	params := make([]interface{}, 0)
	if query.StartedAtLt != 0 {
		querySql += " and started_at < ? "
		params = append(params, time.Unix(query.StartedAtLt, 0))
	}

	if query.Status != "" {
		querySql += " and status = ? "
		params = append(params, query.Status)
	}

	if len(query.StatusList) != 0 {
		querySql += " and status in ? "
		params = append(params, query.StatusList)
	}

	if query.PayStatus != 0 {
		querySql += " and pay_status = ? "
		params = append(params, query.PayStatus)
	}

	if len(query.Ids) != 0 {
		querySql += " and id in ? "
		params = append(params, query.Ids)
	}

	db := d.db.Where(querySql, params...)

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

	res := db.Find(&notebookJobs)
	if res.Error != nil {
		return nil, errors.Errorf(res.Error, errors.ErrorDBFindFailed)
	}

	return notebookJobs, nil
}

func (d *developDao) GetNotebookEvents(notebookEventQuery *model.NotebookEventQuery) ([]*model.NotebookEvent, int64, error) {

	keyName := "object_name"
	keyReason := "reason"
	keyMessage := "message"

	PageIndex := notebookEventQuery.PageIndex
	PageSize := notebookEventQuery.PageSize
	TaskIndex := notebookEventQuery.TaskIndex
	ReplicaIndex := notebookEventQuery.ReplicaIndex
	events := make([]*model.NotebookEvent, 0)

	objectName := fmt.Sprintf("%s-task%d-%d", notebookEventQuery.Id, TaskIndex-1, ReplicaIndex-1)

	countQuery := fmt.Sprintf("SELECT COUNT(%s) FROM octopus..events where object_name = '%s'", keyMessage, objectName)
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

	query := fmt.Sprintf("select %s, %s, %s from octopus..events where object_name = '%s' and kind = 'Pod' LIMIT %d OFFSET %d",
		keyName, keyReason, keyMessage, objectName, PageSize, (PageIndex-1)*PageSize)
	res, err = d.influxdb.Query(query)

	if err != nil {
		return events, 0, errors.Errorf(err, errors.ErroInfluxdbFindFailed)
	}

	for _, row := range res[0].Series[0].Values {

		event := &model.NotebookEvent{}
		event.Timestamp = row[0].(string)
		event.Name = row[1].(string)
		event.Reason = row[2].(string)
		event.Message = row[3].(string)
		events = append(events, event)
	}

	return events, totalSize, nil
}

func (d *developDao) CreateNotebookEventRecord(ctx context.Context, r *model.NotebookEventRecord) error {
	err := d.influxdb.Write("notebook_event_record", r.Time,
		map[string]string{"notebook_id": r.NotebookId},
		map[string]interface{}{"type": int(r.Type), "remark": r.Remark})
	if err != nil {
		return err
	}

	return nil
}

func (d *developDao) ListNotebookEventRecord(ctx context.Context, query *model.NotebookEventRecordQuery) ([]*model.NotebookEventRecord, int64, error) {
	records := make([]*model.NotebookEventRecord, 0)

	countQuery := fmt.Sprintf("select count(remark) from notebook_event_record where notebook_id = '%s'", query.NotebookId)
	res, err := d.influxdb.Query(countQuery)
	if err != nil {
		return nil, 0, errors.Errorf(err, errors.ErroInfluxdbFindFailed)
	}

	if len(res) == 0 || len(res[0].Series) == 0 || len(res[0].Series[0].Values) == 0 || len(res[0].Series[0].Values[0]) < 2 {
		return records, 0, nil
	}

	totalSize, err := res[0].Series[0].Values[0][1].(json.Number).Int64()
	if err != nil {
		return nil, 0, errors.Errorf(err, errors.ErroInfluxdbFindFailed)
	}

	q := fmt.Sprintf("select notebook_id, type, remark from notebook_event_record where notebook_id = '%s' order by time desc limit %d offset %d",
		query.NotebookId, query.PageSize, (query.PageIndex-1)*query.PageSize)
	res, err = d.influxdb.Query(q)
	if err != nil {
		return nil, 0, errors.Errorf(err, errors.ErroInfluxdbFindFailed)
	}

	for _, row := range res[0].Series[0].Values {
		event := &model.NotebookEventRecord{}
		ts := row[0].(string)
		time, err := time.Parse(time.RFC3339, ts)
		if err != nil {
			return nil, 0, errors.Errorf(err, errors.ErrorParseDurationFailed)
		}
		event.Time = time
		event.NotebookId = row[1].(string)
		eventType, err := row[2].(json.Number).Int64()
		if err != nil {
			return nil, 0, errors.Errorf(err, errors.ErrorParseDurationFailed)
		}
		event.Type = v1.NotebookEventRecordType(eventType)
		event.Remark = row[3].(string)
		records = append(records, event)
	}

	return records, totalSize, nil
}
