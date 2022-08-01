package utils

import (
	"server/common/constant"
	"strings"

	typeJob "volcano.sh/apis/pkg/apis/batch/v1alpha1"

	"k8s.io/client-go/tools/cache"
	jsoniter "github.com/json-iterator/go"
	"math"
)

func MapPhaseToState(phase typeJob.JobPhase) string {

	if phase == typeJob.Pending || phase == typeJob.Restarting {
		return constant.PENDING
	}

	if phase == typeJob.Running {
		return constant.RUNNING
	}

	if phase == typeJob.Failed || phase == typeJob.Aborting || phase == typeJob.Aborted {
		return constant.FAILED
	}
	if phase == typeJob.Completed {
		return constant.SUCCEEDED
	}
	return constant.UNKNOWN
}

func IsCompletedState(state string) bool {
	if strings.EqualFold(state, constant.STOPPED) ||
		strings.EqualFold(state, constant.SUCCEEDED) ||
		strings.EqualFold(state, constant.FAILED) {
		return true
	}
	return false
}

func IsRunningOrCompletedState(state string) bool {
	if strings.EqualFold(state, constant.STOPPED) ||
		strings.EqualFold(state, constant.SUCCEEDED) ||
		strings.EqualFold(state, constant.FAILED) ||
		strings.EqualFold(state, constant.RUNNING) {
		return true
	}
	return false
}

func JobRunningState(state string) bool {
	return strings.EqualFold(state, constant.RUNNING)
}

func NonCompletedStates() []string {
	return []string{constant.PREPARING, constant.PENDING, constant.RUNNING}
}

func ConvertJobState(job *typeJob.Job) string {
	if nil == job {
		return constant.UNKNOWN
	}

	state := typeJob.JobPhase(job.Status.State.Phase)

	switch state {
	case typeJob.Pending, typeJob.Restarting:
		{
			return constant.PENDING
		}
	case typeJob.Running:
		{
			return constant.RUNNING
		}
	case typeJob.Failed, typeJob.Aborting, typeJob.Aborted:
		{
			return constant.FAILED
		}
	case typeJob.Completed:
		{
			return constant.SUCCEEDED
		}
	}

	return constant.UNKNOWN
}

func ConvertTaskRoleState(taskroleStatus *typeJob.TaskRoleStatus) string {

	if nil == taskroleStatus {
		return constant.UNKNOWN
	}

	state := taskroleStatus.Phase

	switch state {
	case string(typeJob.Pending):
		{
			return constant.PENDING
		}
	case string(typeJob.Running):
		{
			return constant.RUNNING
		}
	case string(typeJob.Failed):
		{
			return constant.FAILED
		}
	case string(typeJob.Succeeded):
		{
			return constant.SUCCEEDED
		}
	case string(typeJob.Completed):
		{
			return constant.SUCCEEDED
		}
	}

	return constant.UNKNOWN
}

func ConvertTaskRoleReplicaState(replica *typeJob.ReplicaStatus) string {
	if nil == replica {
		return constant.PENDING
	}
	switch replica.Phase {
	case string(typeJob.Pending):
		{
			return constant.PENDING
		}
	case string(typeJob.Running):
		{
			return constant.RUNNING
		}
	case string(typeJob.Succeeded):
		{
			return constant.SUCCEEDED
		}
	case string(typeJob.Failed):
		{
			return constant.FAILED
		}
	}

	if nil == replica.TerminatedInfo {
		return constant.FAILED
	}
	if 0 == replica.TerminatedInfo.ExitCode {
		return constant.SUCCEEDED
	}

	return constant.UNKNOWN
}

func ConvertObjToOtjob(obj interface{}) *typeJob.Job {

	job, ok := obj.(*typeJob.Job)
	if ok {
		return job
	}
	deletedFinalStateUnknown, ok := obj.(cache.DeletedFinalStateUnknown)
	if !ok {
		return nil
	}

	job, ok = deletedFinalStateUnknown.Obj.(*typeJob.Job)
	if !ok {
		return nil
	}
	return job
}

func CurrentLegalStates(state string) []string {

	if constant.PREPARING == state {
		return []string{
			constant.PREPARING,
			constant.UNKNOWN,
		}
	}

	if constant.PENDING == state {
		return []string{
			constant.SUSPENDED,
			constant.PREPARING,
			constant.PENDING,
			constant.UNKNOWN,
		}
	}

	if constant.RUNNING == state {
		return []string{
			constant.SUSPENDED,
			constant.PENDING,
			constant.RUNNING,
			constant.UNKNOWN,
		}
	}

	if constant.SUSPENDED == state ||
	constant.FAILED == state ||
	constant.STOPPED == state ||
	constant.SUCCEEDED == state {
		return []string{
			constant.SUSPENDED,
			constant.PREPARING,
			constant.RUNNING,
			constant.PENDING,
			constant.UNKNOWN,
		}
	}

	return []string{
		constant.UNKNOWN,
	}
}

func Format(jobName, jobKind, userID, cluster string, ExitDiagnostics string, job *typeJob.Job) *typeJob.JobStatusDetail {

	if nil == job {
		return nil
	}

	detail := &typeJob.JobStatusDetail{
		Version: "v1aplha1",
		Job: &typeJob.JobSummary{
			ID:     job.Name,
			Name:   jobName,
			Type:   jobKind,
			UserID: userID,
			State:  ConvertJobState(job),
		},
		Cluster: &typeJob.ClusterInfo{Identity: cluster},
		PlatformSpecificInfo: &typeJob.PlatformSpecificInfo{
			Platform:    "k8s",
			ApiVersion:  job.APIVersion,
			Namespace:   job.Namespace,
			InstanceUID: string(job.UID),
		},
	}

	if job.Status.StartAt != nil {
		detail.Job.StartAt = job.Status.StartAt
	}
	if job.Status.FinishAt != nil {
		detail.Job.FinishedAt = job.Status.FinishAt
	}

	detail.Job.TotalRetriedCount = uint(math.Min(float64(job.Status.RetryCount), float64(job.Spec.MaxRetry)))

	detail.Tasks = make([]*typeJob.TaskInfo, len(job.Spec.Tasks))

	detail.PlatformSpecificInfo.TaskRuntimeInfo = make([]*typeJob.TaskRuntimeInfo, len(job.Spec.Tasks))

	for i := 0; i < len(job.Spec.Tasks); i++ {
		role := &job.Spec.Tasks[i]
		task := &typeJob.TaskInfo{
			Name:          role.Name,
			ReplicaAmount: uint(role.Replicas),
			//MaxFailedTaskCount:    role.CompletionPolicy.MaxFailed,
			//MinSucceededTaskCount: role.CompletionPolicy.MinSucceeded,
		}

		container := role.Template.Spec.Containers[0]

		task.Image = container.Image
		task.Command = append([]string{}, container.Command...)
		buf, _ := jsoniter.Marshal(container.Resources.Limits)
		if nil != buf {
			task.Resource = string(buf)
		}

		detail.Tasks[i] = task

		runtimeInfo := &typeJob.TaskRuntimeInfo{
			Name:         role.Name,
			NodeSelector: role.Template.Spec.NodeSelector,
			VolumeMounts: role.Template.Spec.Volumes,
		}

		detail.PlatformSpecificInfo.TaskRuntimeInfo[i] = runtimeInfo
	}

	if ExitDiagnostics != "" {
		detail.Job.ExitDiagnostics = ExitDiagnostics
	}

	//if nil == job.Status {
	//	return detail
	//}

	for i := 0; i < len(detail.Tasks); i++ {

		var status *typeJob.TaskRoleStatus

		for j := 0; j < len(job.Status.TaskRoleStatus); j++ {
			if job.Status.TaskRoleStatus[j].Name == detail.Tasks[i].Name {
				status = &job.Status.TaskRoleStatus[j]
				break
			}
		}

		if nil == status {
			continue
		}

		detail.Tasks[i].State = ConvertTaskRoleState(status)

		if nil == detail.Tasks[i].Replicas {
			detail.Tasks[i].Replicas = make([]*typeJob.ReplicaInfo, len(status.ReplicaStatuses))
		}

		if nil == detail.PlatformSpecificInfo.TaskRuntimeInfo[i].Replicas {
			detail.PlatformSpecificInfo.TaskRuntimeInfo[i].Replicas = make([]*typeJob.TaskRuntimeReplicaInfo,
				len(status.ReplicaStatuses))
		}

		for k := 0; k < len(status.ReplicaStatuses); k++ {
			status := &status.ReplicaStatuses[k]
			replica := &typeJob.ReplicaInfo{
				Index:           status.Index,
				State:           ConvertTaskRoleReplicaState(status),
				RetriedCount:    status.TotalRetriedCount,
				StartAt:         status.StartAt,
				FinishedAt:      status.FinishAt,
				ContainerID:     status.ContainerID,
				ContainerHostIP: status.PodHostIP,
			}
			if nil != status.TerminatedInfo {
				replica.ExitCode = status.TerminatedInfo.ExitCode
				replica.ExitDiagnostics = status.TerminatedInfo.ExitMessage
				if replica.ExitCode != 0 {
					if detail.Job.State == constant.FAILED || detail.Job.State == constant.SUCCEEDED {
						detail.Job.ExitCode = replica.ExitCode
						if ExitDiagnostics == "" {
							detail.Job.ExitDiagnostics = replica.ExitDiagnostics
						}
					}
				}
			}
			detail.Tasks[i].Replicas[k] = replica

			rReplica := &typeJob.TaskRuntimeReplicaInfo{
				Index:     status.Index,
				PodIP:     status.PodIP,
				PodName:   status.PodName,
				PodHostIP: status.PodHostIP,
				PodReason: status.PodReason,
			}
			if nil != status.PodUID {
				rReplica.PodUID = string(*status.PodUID)
			}

			detail.PlatformSpecificInfo.TaskRuntimeInfo[i].Replicas[k] = rReplica
		}
	}

	if detail.Job.State == constant.SUCCEEDED {
		detail.Job.ExitCode = 0
		detail.Job.ExitDiagnostics = ""
		if ExitDiagnostics == "" {
			detail.Job.ExitDiagnostics = job.Status.State.Message
		}
	}

	return detail
}