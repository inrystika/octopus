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

	updateNb := &model.Notebook{
		NotebookJobId: req.Id,
		Status:        req.CurrentState,
	}

	updateNbJob := &model.NotebookJob{
		Id:     req.Id,
		Status: req.CurrentState,
	}
	if strings.EqualFold(req.CurrentState, pipeline.RUNNING) {
		updateNbJob.StartedAt = &req.CurrentTime
	} else if strings.EqualFold(req.CurrentState, pipeline.FAILED) ||
		strings.EqualFold(req.CurrentState, pipeline.SUCCEEDED) ||
		strings.EqualFold(req.CurrentState, pipeline.STOPPED) {
		updateNbJob.StoppedAt = &req.CurrentTime
		updateNbJob.Status = pipeline.STOPPED //转为stopped
		updateNb.Status = pipeline.STOPPED    //转为stopped

		err = s.data.Cluster.DeleteIngress(ctx, req.UserID, req.Id)
		if err != nil {
			return common.PipeLineCallbackRE
		}

		err = s.data.Cluster.DeleteService(ctx, req.UserID, req.Id)
		if err != nil {
			return common.PipeLineCallbackRE
		}
	}

	err = s.data.DevelopDao.UpdateNotebookSelectiveByJobId(ctx, updateNb)
	if err != nil {
		return common.PipeLineCallbackRE
	}

	err = s.data.DevelopDao.UpdateNotebookJobSelective(ctx, updateNbJob)
	if err != nil {
		return common.PipeLineCallbackRE
	}

	return common.PipeLineCallbackOK
}
