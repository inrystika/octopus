package taskset

import (
	"math"
	v1 "scheduler/pkg/pipeline/apis/http/v1"
	"scheduler/pkg/pipeline/constants/jobstate"

	jsoniter "github.com/json-iterator/go"
	typeJob "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
)

func convertTaskSetState(taskset *typeJob.Job) string {
	if nil == taskset {
		return jobstate.UNKNOWN
	}

	state := taskset.Status.State.Phase

	switch state {
	case typeJob.Pending, typeJob.Restarting:
		{
			return jobstate.PENDING
		}
	case typeJob.Running:
		{
			return jobstate.RUNNING
		}
	case typeJob.Failed, typeJob.Aborting, typeJob.Aborted:
		{
			return jobstate.FAILED
		}
	case typeJob.Completed:
		{
			return jobstate.SUCCEEDED
		}
	}

	return jobstate.UNKNOWN
}

func convertTaskRoleState(taskroleStatus *typeJob.TaskRoleStatus) string {

	if nil == taskroleStatus {
		return jobstate.UNKNOWN
	}

	state := taskroleStatus.Phase

	switch state {
	case string(typeJob.Pending):
		{
			return jobstate.PENDING
		}
	case string(typeJob.Running):
		{
			return jobstate.RUNNING
		}
	case string(typeJob.Failed):
		{
			return jobstate.FAILED
		}
	case string(typeJob.Succeeded):
		{
			return jobstate.SUCCEEDED
		}
	case string(typeJob.Completed):
		{
			return jobstate.SUCCEEDED
		}
	}

	return jobstate.UNKNOWN
}

func convertTaskRoleReplicaState(replica *typeJob.ReplicaStatus) string {
	if nil == replica {
		return jobstate.PENDING
	}
	switch replica.Phase {
	case string(typeJob.Pending):
		{
			return jobstate.PENDING
		}
	case string(typeJob.Running):
		{
			return jobstate.RUNNING
		}
	case string(typeJob.Succeeded):
		{
			return jobstate.SUCCEEDED
		}
	case string(typeJob.Failed):
		{
			return jobstate.FAILED
		}
	}
	/*
		if nil == replica.TerminatedInfo {
			return jobstate.FAILED
		}
		if 0 == replica.TerminatedInfo.ExitCode {
			return jobstate.SUCCEEDED
		}
	*/
	return jobstate.UNKNOWN
}

func Format(jobName, jobKind, userID, cluster string, ExitDiagnostics string, taskset *typeJob.Job) *v1.JobStatusDetail {

	if nil == taskset {
		return nil
	}

	detail := &v1.JobStatusDetail{
		Version: "v1",
		Job: &v1.JobSummary{
			ID:     taskset.Name,
			Name:   jobName,
			Type:   jobKind,
			UserID: userID,
			State:  convertTaskSetState(taskset),
		},
		Cluster: &v1.ClusterInfo{Identity: cluster},
		PlatformSpecificInfo: &v1.PlatformSpecificInfo{
			Platform:    "k8s",
			ApiVersion:  taskset.APIVersion,
			Namespace:   taskset.Namespace,
			InstanceUID: string(taskset.UID),
		},
	}

	if taskset.Status.StartAt != nil {
		detail.Job.StartAt = taskset.Status.StartAt
	}
	if taskset.Status.FinishAt != nil {
		detail.Job.FinishedAt = taskset.Status.FinishAt
	}

	detail.Job.TotalRetriedCount = uint(math.Min(float64(taskset.Status.RetryCount), float64(taskset.Spec.MaxRetry)))

	detail.Tasks = make([]*v1.TaskInfo, len(taskset.Spec.Tasks))

	detail.PlatformSpecificInfo.TaskRuntimeInfo = make([]*v1.TaskRuntimeInfo, len(taskset.Spec.Tasks))

	for i := 0; i < len(taskset.Spec.Tasks); i++ {
		role := &taskset.Spec.Tasks[i]
		task := &v1.TaskInfo{
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

		runtimeInfo := &v1.TaskRuntimeInfo{
			Name:         role.Name,
			NodeSelector: role.Template.Spec.NodeSelector,
			VolumeMounts: role.Template.Spec.Volumes,
		}

		detail.PlatformSpecificInfo.TaskRuntimeInfo[i] = runtimeInfo
	}

	if ExitDiagnostics != "" {
		detail.Job.ExitDiagnostics = ExitDiagnostics
	}

	//if nil == taskset.Status {
	//	return detail
	//}

	for i := 0; i < len(detail.Tasks); i++ {

		var status *typeJob.TaskRoleStatus

		for j := 0; j < len(taskset.Status.TaskRoleStatus); j++ {
			if taskset.Status.TaskRoleStatus[j].Name == detail.Tasks[i].Name {
				status = &taskset.Status.TaskRoleStatus[j]
				break
			}
		}

		if nil == status {
			continue
		}

		detail.Tasks[i].State = convertTaskRoleState(status)

		if nil == detail.Tasks[i].Replicas {
			detail.Tasks[i].Replicas = make([]*v1.ReplicaInfo, len(status.ReplicaStatuses))
		}

		if nil == detail.PlatformSpecificInfo.TaskRuntimeInfo[i].Replicas {
			detail.PlatformSpecificInfo.TaskRuntimeInfo[i].Replicas = make([]*v1.TaskRuntimeReplicaInfo,
				len(status.ReplicaStatuses))
		}

		for k := 0; k < len(status.ReplicaStatuses); k++ {
			status := &status.ReplicaStatuses[k]
			replica := &v1.ReplicaInfo{
				Index:           status.Index,
				State:           convertTaskRoleReplicaState(status),
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
					if detail.Job.State == jobstate.FAILED || detail.Job.State == jobstate.SUCCEEDED {
						detail.Job.ExitCode = replica.ExitCode
						if ExitDiagnostics == "" {
							detail.Job.ExitDiagnostics = replica.ExitDiagnostics
						}
					}
				}
			}
			detail.Tasks[i].Replicas[k] = replica

			rReplica := &v1.TaskRuntimeReplicaInfo{
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

	if detail.Job.State == jobstate.SUCCEEDED {
		detail.Job.ExitCode = 0
		detail.Job.ExitDiagnostics = ""
		if ExitDiagnostics == "" {
			detail.Job.ExitDiagnostics = taskset.Status.State.Message
		}
	}

	return detail
}
