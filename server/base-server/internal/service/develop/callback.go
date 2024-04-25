package develop

import (
	"context"
	"encoding/json"
	"fmt"
	"server/base-server/internal/data/dao/model"
	"server/common/constant"
	"server/common/utils"
	"strings"

	jobUtil "server/base-server/internal/common"

	typeJob "volcano.sh/apis/pkg/apis/batch/v1alpha1"
)

func (s *developService) onJobAdd(obj interface{}) {
}

func (s *developService) onJobDelete(obj interface{}) {
	job := utils.ConvertObjToOtjob(obj)
	if job == nil {
		return
	}
	if job.Annotations == nil {
		return
	}
	jobType, found := job.Annotations[constant.JOB_TYPE]
	if !found || jobType != constant.NotebookJob {
		return
	}

	nbJob, err := s.data.DevelopDao.GetNotebookJob(context.TODO(), job.Name)
	if err != nil {
		s.log.Error(context.TODO(), "GetNotebookJob err when onJobDelete: "+job.Name, err)
		return
	}
	nb, err := s.data.DevelopDao.GetNotebook(context.TODO(), nbJob.NotebookId)
	if err != nil {
		s.log.Error(context.TODO(), "GetNotebook err when onJobDelete: "+job.Name, err)
		return
	}

	detail := jobUtil.GetStopDetail(nbJob.Detail)
	detailBuf, err := json.Marshal(detail)
	if err != nil {
		s.log.Error(context.TODO(), "Marshal err when onJobDelete:"+job.Name, err)
	}
	newJob := &model.NotebookJob{
		Id:     job.Name,
		Detail: string(detailBuf),
	}
	if !utils.IsCompletedState(nbJob.Status) {
		newJob.Status = constant.STOPPED
		err = s.data.DevelopDao.UpdateNotebookSelective(context.TODO(), &model.Notebook{
			Id:     nbJob.NotebookId,
			Status: constant.STOPPED,
		})
		if err != nil {
			s.log.Error(context.TODO(), "UpdateNotebookSelective err when onJobDelete:"+job.Name, err)
		}
		s.sendEmail(nb.UserId, fmt.Sprintf("Notebook %s %s", nb.Name, newJob.Status))
	}
	err = s.data.DevelopDao.UpdateNotebookJobSelective(context.TODO(), newJob)
	if err != nil {
		s.log.Error(context.TODO(), "UpdateNotebookJobSelective err when onJobDelete:"+job.Name, err)
	}
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

	jobCopy := newjob.DeepCopy()
	s.updatedJob <- jobCopy
}
