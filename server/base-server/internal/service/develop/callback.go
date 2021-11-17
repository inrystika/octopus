package develop

import (
	"context"
	"server/base-server/internal/common"
	"server/base-server/internal/data/dao/model"
	"server/base-server/internal/data/pipeline"
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

	if strings.EqualFold(req.CurrentState, pipeline.RUNNING) {
		nbJobUp.StartedAt = &req.CurrentTime
	} else if strings.EqualFold(req.CurrentState, pipeline.FAILED) ||
		strings.EqualFold(req.CurrentState, pipeline.SUCCEEDED) ||
		strings.EqualFold(req.CurrentState, pipeline.STOPPED) {
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
	}

	err = s.data.DevelopDao.UpdateNotebookSelectiveByJobId(ctx, nbUp)
	if err != nil {
		return common.PipeLineCallbackRE
	}

	err = s.data.DevelopDao.UpdateNotebookJobSelective(ctx, nbJobUp)
	if err != nil {
		return common.PipeLineCallbackRE
	}

	return common.PipeLineCallbackOK
}
