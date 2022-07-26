/*
Copyright 2019 The Volcano Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package helpers

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "k8s.io/api/core/v1"

	batch "volcano.sh/apis/pkg/apis/batch/v1alpha1"
	"volcano.sh/volcano/pkg/controllers/apis"
	"volcano.sh/volcano/pkg/scheduler/api"
)

const (
	// PodNameFmt pod name format
	PodNameFmt = "%s-%s-%d"
	// persistentVolumeClaimFmt represents persistent volume claim name format
	persistentVolumeClaimFmt = "%s-pvc-%s"
)

// GetTaskIndex returns task Index.
func GetTaskIndex(pod *v1.Pod) string {
	num := strings.Split(pod.Name, "-")
	if len(num) >= 3 {
		return num[len(num)-1]
	}

	return ""
}

// GetTaskName returns task Name.
func GetTaskName(pod *v1.Pod) string {
	num := strings.Split(pod.Name, "-")
	if len(num) >= 3 {
		return num[len(num)-2]
	}

	return ""
}

// GetPodIndexUnderTask returns task Index.
func GetPodIndexUnderTask(pod *v1.Pod) string {
	num := strings.Split(pod.Name, "-")
	if len(num) >= 3 {
		return num[len(num)-1]
	}

	return ""
}

// CompareTask by pod index
func CompareTask(lv, rv *api.TaskInfo) bool {
	lStr := GetPodIndexUnderTask(lv.Pod)
	rStr := GetPodIndexUnderTask(rv.Pod)
	lIndex, lErr := strconv.Atoi(lStr)
	rIndex, rErr := strconv.Atoi(rStr)
	if lErr != nil || rErr != nil || lIndex == rIndex {
		return lv.Pod.CreationTimestamp.Before(&rv.Pod.CreationTimestamp)
	}
	if lIndex > rIndex {
		return false
	}
	return true
}

// GetTaskKey returns task key/name
func GetTaskKey(pod *v1.Pod) string {
	if pod.Annotations == nil || pod.Annotations[batch.TaskSpecKey] == "" {
		return batch.DefaultTaskSpec
	}
	return pod.Annotations[batch.TaskSpecKey]
}

// GetTaskSpec returns task spec
func GetTaskSpec(job *batch.Job, taskName string) (batch.TaskSpec, bool) {
	for _, ts := range job.Spec.Tasks {
		if ts.Name == taskName {
			return ts, true
		}
	}
	return batch.TaskSpec{}, false
}

// MakeDomainName creates task domain name
func MakeDomainName(ts batch.TaskSpec, job *batch.Job, index int) string {
	hostName := ts.Template.Spec.Hostname
	subdomain := ts.Template.Spec.Subdomain
	if len(hostName) == 0 {
		hostName = MakePodName(job.Name, ts.Name, index)
	}
	if len(subdomain) == 0 {
		subdomain = job.Name
	}
	return hostName + "." + subdomain
}

// MakePodName creates pod name.
func MakePodName(jobName string, taskName string, index int) string {
	return fmt.Sprintf(PodNameFmt, jobName, taskName, index)
}

// GenRandomStr generate random str with specified length l.
func GenRandomStr(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// GenPVCName generates pvc name with job name.
func GenPVCName(jobName string) string {
	return fmt.Sprintf(persistentVolumeClaimFmt, jobName, GenRandomStr(12))
}

// GetJobKeyByReq gets the key for the job request.
func GetJobKeyByReq(req *apis.Request) string {
	return fmt.Sprintf("%s/%s", req.Namespace, req.JobName)
}

// GetTasklndexUnderJob return index of the task in the job.
func GetTasklndexUnderJob(taskName string, job *batch.Job) int {
	for index, task := range job.Spec.Tasks {
		if task.Name == taskName {
			return index
		}
	}
	return -1
}

// GetPodsNameUnderTask return names of all pods in the task.
func GetPodsNameUnderTask(taskName string, job *batch.Job) []string {
	var res []string
	for _, task := range job.Spec.Tasks {
		if task.Name == taskName {
			for index := 0; index < int(task.Replicas); index++ {
				res = append(res, MakePodName(job.Name, taskName, index))
			}
			break
		}
	}
	return res
}

func UpdateReplicaStatus(record *batch.ReplicaStatus, pod *v1.Pod) {

	if nil == pod || nil == record {
		return
	}
	if record.Phase != string(pod.Status.Phase) {
		record.Phase = string(pod.Status.Phase)
		record.TransitionTime = metav1.Now()
		record.PhaseMessage = pod.Status.Message
	}

	if nil == record.PodUID {
		record.PodUID = &pod.UID
	}

	if pod.Name != record.PodName {
		record.PodName = pod.Name
	}

	if pod.Status.Reason != record.PodReason {
		record.PodReason = pod.Status.Reason
	}

	if nil == record.StartAt && pod.Status.StartTime != nil {
		record.StartAt = &metav1.Time{Time: (*pod.Status.StartTime).Time}
	}

	if pod.Status.PodIP != record.PodIP && "" != pod.Status.PodIP {
		record.PodIP = pod.Status.PodIP
	}
	if pod.Status.HostIP != record.PodHostIP && "" != pod.Status.HostIP {
		record.PodHostIP = pod.Status.HostIP
	}

	if record.ContainerName == "" && len(pod.Status.ContainerStatuses) > 0 {
		record.ContainerName = pod.Status.ContainerStatuses[0].Name
	}

	if record.ContainerID == "" && len(pod.Status.ContainerStatuses) > 0 {
		record.ContainerID = pod.Status.ContainerStatuses[0].ContainerID
	}

	if record.TerminatedInfo == nil && len(pod.Status.ContainerStatuses) > 0 &&
		nil != pod.Status.ContainerStatuses[0].State.Terminated {

		record.FinishAt = &metav1.Time{Time: pod.Status.ContainerStatuses[0].State.Terminated.FinishedAt.Time}
		//can't use container's startTime as replica's startTime ,because container maybe restarts many times

		if record.TerminatedInfo == nil {
			record.TerminatedInfo = &batch.ContainerTerminatedInfo{}
		}

		terminated := pod.Status.ContainerStatuses[0].State.Terminated //pointer
		terminatedInfo := record.TerminatedInfo                        //pointer

		terminatedInfo.ExitCode = terminated.ExitCode
		terminatedInfo.ExitMessage = terminated.Message
		terminatedInfo.Signal = terminated.Signal
		terminatedInfo.Reason = terminated.Reason
	}
}

func UpdateTaskRoleStatus(status *batch.TaskRoleStatus) {
	runningCnt := 0
	succeedCnt := 0
	failedCnt := 0
	for j := 0; j < len(status.ReplicaStatuses); j++ {
		phase := status.ReplicaStatuses[j].Phase
		if phase == string(batch.Running) {
			runningCnt += 1
		} else if phase == string(batch.Succeeded) {
			succeedCnt += 1
		} else if phase == string(batch.Completed) {
			succeedCnt += 1
		} else if phase == string(batch.Failed) {
			failedCnt += 1
		}
	}
	if runningCnt == len(status.ReplicaStatuses) {
		status.Phase = string(batch.Running)
	} else if succeedCnt == len(status.ReplicaStatuses) {
		status.Phase = string(batch.Completed)
	} else if failedCnt > 0 {
		status.Phase = string(batch.Failed)
	}
}

// ResetJobStatus reset the status of a job.
func ResetJobStatus(status *batch.JobStatus) {
	for i := 0; i < len(status.TaskRoleStatus); i++ {
		for j := 0; j < len(status.TaskRoleStatus[i].ReplicaStatuses); j++ {
			status.TaskRoleStatus[i].ReplicaStatuses[j].PodUID = nil
			status.TaskRoleStatus[i].ReplicaStatuses[j].TerminatedInfo = nil
			status.TaskRoleStatus[i].ReplicaStatuses[j].PodName = ""
			status.TaskRoleStatus[i].ReplicaStatuses[j].PodIP = ""
			status.TaskRoleStatus[i].ReplicaStatuses[j].PodHostIP = ""
			status.TaskRoleStatus[i].ReplicaStatuses[j].ContainerName = ""
			status.TaskRoleStatus[i].ReplicaStatuses[j].ContainerID = ""
			status.TaskRoleStatus[i].ReplicaStatuses[j].StartAt = nil
			status.TaskRoleStatus[i].ReplicaStatuses[j].FinishAt = nil
		}
	}
}
