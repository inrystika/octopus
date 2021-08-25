package module

import (
	typeJob "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
)

type JobEvent struct {
	JobID     string
	EventName string
	Namespace string
	Header    []byte
	TaskSet   *typeJob.Job
}

func (j *JobEvent) DeepCopy() *JobEvent {
	je := &JobEvent{
		JobID:     j.JobID,
		EventName: j.EventName,
		Namespace: j.Namespace,
	}
	if nil != j.Header {

		je.Header = make([]byte, len(j.Header))
		copy(je.Header, j.Header)
	}
	if nil != j.TaskSet {
		je.TaskSet = j.TaskSet.DeepCopy()
	}
	return je
}
