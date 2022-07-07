package apis

import (
	typeJob "server/apis/pkg/apis/batch/v1alpha1"

	v1 "k8s.io/api/core/v1"
	"volcano.sh/volcano/pkg/controllers/apis"
)

//JobInfo struct.
type JobInfo struct {
	apis.JobInfo

	Job *typeJob.Job
}

//Clone function clones the k8s pod values to the JobInfo struct.
func (ji *JobInfo) Clone() *JobInfo {
	job := &JobInfo{}
	job.Namespace = ji.Namespace
	job.Name = ji.Name
	job.Job = ji.Job
	job.Pods = make(map[string]map[string]*v1.Pod)

	for key, pods := range ji.Pods {
		job.Pods[key] = make(map[string]*v1.Pod)
		for pn, pod := range pods {
			job.Pods[key][pn] = pod
		}
	}

	return job
}

//SetJob sets the volcano jobs values to the JobInfo struct.
func (ji *JobInfo) SetJob(job *typeJob.Job) {
	ji.Name = job.Name
	ji.Namespace = job.Namespace
	ji.Job = job
}
