package utils

import (
	"server/common/constant"
	"strings"

	typeJob "volcano.sh/apis/pkg/apis/batch/v1alpha1"

	"k8s.io/client-go/tools/cache"
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
