package develop

import (
	"context"
	"encoding/json"
	"server/base-server/internal/data/dao/model"
	commapi "server/common/api/v1"
	"server/common/constant"
	"server/common/utils"
	"strings"
	"time"

	typeJob "volcano.sh/apis/pkg/apis/batch/v1alpha1"
)

func (s *developService) onJobAdd(obj interface{}) {
}

func (s *developService) onJobDelete(obj interface{}) {

}

func (s *developService) onJobUpdate(old, obj interface{}) {

	oldjob := utils.ConvertObjToOtjob(old)
	newjob := utils.ConvertObjToOtjob(obj)
	if oldjob == nil || newjob == nil {
		return
	}

	if newjob.Annotations == nil {
		return
	}
	jobType, found := newjob.Annotations[constant.JOB_TYPE]
	if !found || jobType != constant.NotebookJob {
		return
	}

	oldState := utils.MapPhaseToState(typeJob.JobPhase(oldjob.Status.State.Phase))
	newState := utils.MapPhaseToState(typeJob.JobPhase(newjob.Status.State.Phase))

	if strings.EqualFold(constant.UNKNOWN, newState) {
		return
	}

	if newState == string(typeJob.Pending) && nil != oldjob {
		if oldState == string(typeJob.Running) {
			return
		}
	}
	ctx := context.TODO()
	nbJob, err := s.data.DevelopDao.GetNotebookJob(ctx, newjob.Name)
	if err != nil {
		s.log.Error(ctx, "GetTrainJob err when onJobUpdate:"+newjob.Name, err)
		return
	}

	if utils.IsCompletedState(nbJob.Status) || strings.EqualFold(nbJob.Status, newState) {
		return
	}

	nb, err := s.data.DevelopDao.GetNotebook(ctx, nbJob.NotebookId)
	if err != nil {
		s.log.Error(ctx, "GetNotebook err when onJobUpdate:"+newjob.Name, err)
		return
	}

	nbUp := &model.Notebook{
		NotebookJobId: newjob.Name,
		Status:        newState,
	}

	nbJobUp := &model.NotebookJob{
		Id:     newjob.Name,
		Status: newState,
	}

	status := utils.Format(newjob.Name, "notebook", newjob.Namespace, "", "", newjob)
	if nil != status {
		buf, err := json.Marshal(status)
		if err != nil {
			s.log.Error(context.TODO(), "UpdateNotebook err when onJobUpdate: "+newjob.Name, err)
		}
		nbJobUp.Detail = string(buf)
	}

	now := time.Now()
	record := &model.NotebookEventRecord{
		Time:       now,
		NotebookId: nb.Id,
	}

	if strings.EqualFold(newState, constant.RUNNING) {
		nbJobUp.StartedAt = &now
		record.Type = commapi.NotebookEventRecordType_RUN
	} else if utils.IsCompletedState(newState) {
		nbJobUp.StoppedAt = &now
		nbJobUp.Status = constant.STOPPED //转为stopped
		nbUp.Status = constant.STOPPED    //转为stopped

		err = s.deleteIngress(ctx, nb, nbJob)
		if err != nil {
			s.log.Error(ctx, "deleteIngress err when onJobUpdate:"+newjob.Name, err)
		}

		err = s.deleteService(ctx, nb, nbJob)
		if err != nil {
			s.log.Error(ctx, "deleteService err when onJobUpdate:"+newjob.Name, err)
		}
		record.Type = commapi.NotebookEventRecordType_STOP
	}

	err = s.data.DevelopDao.UpdateNotebookSelectiveByJobId(ctx, nbUp)
	if err != nil {
		s.log.Error(ctx, "UpdateNotebookSelectiveByJobId err when onJobUpdate:"+newjob.Name, err)
	}

	err = s.data.DevelopDao.UpdateNotebookJobSelective(ctx, nbJobUp)
	if err != nil {
		s.log.Error(ctx, "UpdateNotebookJobSelective err when onJobUpdate:"+newjob.Name, err)
	}

	if utils.IsRunningOrCompletedState(newState) {
		err = s.data.DevelopDao.CreateNotebookEventRecord(ctx, record)
		if err != nil { // 插入事件记录出错只打印
			s.log.Error(ctx, "create notebook event record error:", err)
		}
	}
}
