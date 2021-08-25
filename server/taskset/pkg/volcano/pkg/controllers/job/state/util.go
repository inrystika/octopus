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

package state

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	vcbatch "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
	bus "volcano.sh/volcano/pkg/apis/bus/v1alpha1"
)

// DefaultMaxRetry is the default number of retries.
const DefaultMaxRetry int32 = 0

//DefaultTaskRoleCompletedPolicies define the default role completion event policy
var DefaultTaskRoleCompletedPolicies = []vcbatch.LifecyclePolicy{
	{
		Event:  bus.PodFailedEvent,
		Action: bus.AbortJobAction,
	},
	{
		Event:  bus.TaskCompletedEvent,
		Action: bus.NoActionAction,
	},
}

// TotalTasks returns number of tasks in a given volcano job.
func TotalTasks(job *vcbatch.Job) int32 {
	var rep int32

	for _, task := range job.Spec.Tasks {
		rep += task.Replicas
	}

	return rep
}

//GetTaskSpec returns the task spec
func GetTaskSpec(job *vcbatch.Job, rolename string) *vcbatch.TaskSpec {
	var taskSpec *vcbatch.TaskSpec
	for i := 0; i < len(job.Spec.Tasks); i++ {
		if rolename == job.Spec.Tasks[i].Name {
			taskSpec = &job.Spec.Tasks[i]
			break
		}
	}
	return taskSpec
}

//GetTaskStatus returns the status of specific taskrole
func GetTaskStatus(job *vcbatch.Job, taskname string) *vcbatch.TaskRoleStatus {
	var taskStatus *vcbatch.TaskRoleStatus
	for i := 0; i < len(job.Status.TaskRoleStatus); i++ {
		if taskname == job.Status.TaskRoleStatus[i].Name {
			taskStatus = &job.Status.TaskRoleStatus[i]
			break
		}
	}
	return taskStatus
}

//GetTaskCompletionPolicy returns the completion policy of specific taskrole
func GetTaskCompletionPolicy(job *vcbatch.Job, taskname string) *vcbatch.CompletionPolicy {
	spec := GetTaskSpec(job, taskname)
	if nil == spec {
		return nil
	}
	return &spec.CompletionPolicy
}

func ShouldJobCompleted(job *vcbatch.Job, status *vcbatch.JobStatus) (should bool, succeeded bool) {

	should = false
	succeeded = false

	var completedCount int = 0
	var succeededCount int = 0

	for i := 0; i < len(status.TaskRoleStatus); i++ {

		status := &status.TaskRoleStatus[i]

		taskSpec := GetTaskSpec(job, status.Name)

		if nil == taskSpec {
			continue
		}

		if status.Phase != string(vcbatch.Completed) && status.Phase != string(vcbatch.Failed) && status.Phase != string(vcbatch.Succeeded) {
			continue
		}

		if status.Phase == string(vcbatch.Completed) || status.Phase == string(vcbatch.Succeeded) {
			succeededCount++
		}

		completedCount++

		var eventPolicies []vcbatch.LifecyclePolicy = taskSpec.Policies

		if len(eventPolicies) == 0 {
			eventPolicies = DefaultTaskRoleCompletedPolicies
		} else {
			var completedEvent, failedEvent bool = false, false
			for k := 0; k < len(eventPolicies); k++ {
				policy := &eventPolicies[k]
				if policy.Event == bus.PodFailedEvent {
					failedEvent = true
				}
				if policy.Event == bus.TaskCompletedEvent {
					completedEvent = true
				}
			}
			if false == completedEvent && failedEvent == false {
				eventPolicies = append(eventPolicies, vcbatch.LifecyclePolicy{
					Event:  bus.PodFailedEvent,
					Action: bus.AbortJobAction,
				})
			}
		}

		for k := 0; k < len(eventPolicies); k++ {
			policy := &eventPolicies[k]
			should, succeeded = ShouldJobCompletedByEventPolicy(policy, status)
			if true == should {
				break
			}
		}

		if should == true {
			break
		}
	}

	if false == should && completedCount == len(status.TaskRoleStatus) {
		if succeededCount == len(status.TaskRoleStatus) {
			return true, true
		}

		return true, false
	}

	return should, succeeded
}

//ShouldJobCompletedByEventPolicy determines if the Job should be completed by event policy
func ShouldJobCompletedByEventPolicy(policy *vcbatch.LifecyclePolicy,
	rolestatus *vcbatch.TaskRoleStatus) (completed bool, success bool) {

	if (rolestatus.State == string(vcbatch.Failed) && policy.Event == bus.PodFailedEvent) ||
		(rolestatus.State == string(vcbatch.Completed) && policy.Event == bus.RoleSucceededEvent) ||
		((rolestatus.State == string(vcbatch.Completed) || rolestatus.State == string(vcbatch.Failed)) && policy.Event == bus.TaskCompletedEvent) {

		if policy.Action == bus.AbortJobAction {
			return true, false
		}
		//if policy.Action == bus.TaskSetSucceededAction {
		//	return true, true
		//}
		if policy.Action == bus.CompleteJobAction {
			return true, true
		}

		if policy.Action == bus.NoActionAction {
			return false, false
		}
	}

	return false, false
}

func StopReplicas(record *vcbatch.TaskRoleStatus) {

	for i := 0; i < len(record.ReplicaStatuses); i++ {

		replica := &record.ReplicaStatuses[i]

		if replica.Phase == string(vcbatch.Running) || replica.Phase == string(vcbatch.Pending) || replica.Phase == string(vcbatch.Restarting) {
			replica.Phase = string(vcbatch.Completed)
			replica.PhaseMessage = "Stop replica"
			replica.TransitionTime = metav1.Now()
			replica.Stopped = true
		}

	}
}
