package common

import (
	"encoding/json"
	"server/common/constant"
	"server/common/utils"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typeJob "volcano.sh/apis/pkg/apis/batch/v1alpha1"
)

func GetStopDetail(detailstr string) *typeJob.JobStatusDetail {
	detail := typeJob.JobStatusDetail{}
	json.Unmarshal([]byte(detailstr), &detail)

	if detail.Job != nil {
		if !utils.IsCompletedState(detail.Job.State) {
			detail.Job.State = constant.STOPPED
		}
		if detail.Job.FinishedAt == nil {
			detail.Job.FinishedAt = &metav1.Time{Time: time.Now()}
		}
		for _, role := range detail.Tasks {
			if !utils.IsCompletedState(role.State) {
				role.State = constant.STOPPED
			}
			if role.Replicas == nil {
				continue
			}
			for _, roleReplica := range role.Replicas {
				if !utils.IsCompletedState(roleReplica.State) {
					roleReplica.State = constant.STOPPED
				}
				if roleReplica.FinishedAt == nil {
					roleReplica.FinishedAt = &metav1.Time{Time: time.Now()}
				}
			}
		}
	}

	return &detail
}
