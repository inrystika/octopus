package develop

import (
	"context"
	"fmt"
	"server/base-server/internal/common"
	"server/base-server/internal/data/dao/model"
	"server/base-server/internal/data/pipeline"
	commapi "server/common/api/v1"
	"strings"
)

func (s *developService) PipelineCallback(ctx context.Context, req *common.PipelineCallbackReq) string {
	nbJob, err := s.data.DevelopDao.GetNotebookJob(ctx, req.Id)
	if err != nil {
		return common.PipeLineCallbackRE
	}

	if pipeline.IsCompletedState(nbJob.Status) || strings.EqualFold(nbJob.Status, req.CurrentState) {
		return common.PipeLineCallbackOK
	}

	nb, err := s.data.DevelopDao.GetNotebook(ctx, nbJob.NotebookId)
	if err != nil {
		return common.PipeLineCallbackRE
	}

	nbUp := &model.Notebook{
		NotebookJobId: req.Id,
		Status:        req.CurrentState,
	}

	nbJobUp := &model.NotebookJob{
		Id:     req.Id,
		Status: req.CurrentState,
	}

	record := &model.NotebookEventRecord{
		Time:       req.CurrentTime,
		NotebookId: nb.Id,
	}

	if strings.EqualFold(req.CurrentState, pipeline.RUNNING) {
		nbJobUp.StartedAt = &req.CurrentTime
		record.Type = commapi.NotebookEventRecordType_RUN
		record.Title = fmt.Sprintf("%s run", nb.Name)
	} else if pipeline.IsCompletedState(req.CurrentState) {
		nbJobUp.StoppedAt = &req.CurrentTime
		nbJobUp.Status = pipeline.STOPPED //转为stopped
		nbUp.Status = pipeline.STOPPED    //转为stopped

		err = s.deleteIngress(ctx, nb, nbJob)
		if err != nil {
			return common.PipeLineCallbackRE
		}

		err = s.deleteService(ctx, nb, nbJob)
		if err != nil {
			return common.PipeLineCallbackRE
		}

		record.Type = commapi.NotebookEventRecordType_STOP
		record.Title = fmt.Sprintf("%s stop", nb.Name)
	}

	err = s.data.DevelopDao.UpdateNotebookSelectiveByJobId(ctx, nbUp)
	if err != nil {
		return common.PipeLineCallbackRE
	}

	err = s.data.DevelopDao.UpdateNotebookJobSelective(ctx, nbJobUp)
	if err != nil {
		return common.PipeLineCallbackRE
	}

	if pipeline.IsRunningOrCompletedState(req.CurrentState) {
		err = s.data.DevelopDao.CreateNotebookEventRecord(ctx, record)
		if err != nil { // 插入事件记录出错只打印
			s.log.Error(ctx, "create notebook event record error:", err)
		}
	}

	return common.PipeLineCallbackOK
}
