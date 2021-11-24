package common

import (
	"context"
	"server/base-server/internal/data/pipeline"
	"server/common/log"
	"strings"
	"time"
)

func CalculateAmount(ctx context.Context, job *pipeline.JobStatusDetail, prices []uint32) float64 {
	var rs float64
	for ti, t := range job.Tasks {
		for _, r := range t.Replicas {
			var startAt, finishedAt int64
			if pipeline.IsCompletedState(r.State) && r.FinishedAt != nil && r.StartAt != nil {
				startAt = r.StartAt.Unix()
				finishedAt = r.FinishedAt.Unix()
			} else if strings.EqualFold(r.State, pipeline.RUNNING) && r.StartAt != nil {
				startAt = r.StartAt.Unix()
				finishedAt = time.Now().Unix()
			} else {
				log.Infof(ctx, "calculate amount abnormal,jobId:%v", job.Job.ID)
			}
			rs += float64(finishedAt-startAt) * float64(prices[ti]) / 3600.0
		}
	}

	return rs
// 模型开发和训练生成extraInfo
func GetExtraInfo(userId string) map[string]string {
	return map[string]string{
		"userId": userId,
	}
}
