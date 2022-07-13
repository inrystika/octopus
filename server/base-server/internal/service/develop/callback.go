package develop

import (
	"context"
	"server/base-server/internal/common"
	"server/base-server/internal/data/dao/model"
	commapi "server/common/api/v1"
	"server/common/constant"
	"server/common/utils"
	"strings"
	"time"
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

	oldState := utils.MapPhaseToState(typeJob.JobPhase(oldjob.Status.State.Phase))
	newState := utils.MapPhaseToState(typeJob.JobPhase(newjob.Status.State.Phase))

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
		NotebookJobId: req.Id,
		Status:        newState,
	}

	nbJobUp := &model.NotebookJob{
		Id:     req.Id,
		Status: newState,
	}

	now := time.Now()
	record := &model.NotebookEventRecord{
		Time:       &now,
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

	if utils.IsRunningOrCompletedState(req.CurrentState) {
		err = s.data.DevelopDao.CreateNotebookEventRecord(ctx, record)
		if err != nil { // 插入事件记录出错只打印
			s.log.Error(ctx, "create notebook event record error:", err)
		}
	}

	return common.PipeLineCallbackOK
}
