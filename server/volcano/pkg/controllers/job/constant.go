/*
Copyright 2020 The Volcano Authors.

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

package job

// Reasons for pod events.
const (
	// FailedCreatePodReason is added in an event and in a replica set condition
	// when a pod for a replica set is failed to be created.
	FailedCreatePodReason = "FailedCreate"
	// SuccessfulCreatePodReason is added in an event when a pod for a replica set
	// is successfully created.
	SuccessfulCreatePodReason = "SuccessfulCreate"
	// FailedDeletePodReason is added in an event and in a replica set condition
	// when a pod for a replica set is failed to be deleted.
	FailedDeletePodReason = "FailedDelete"
	// SuccessfulDeletePodReason is added in an event when a pod for a replica set
	// is successfully deleted.
	SuccessfulDeletePodReason = "SuccessfulDelete"

	//EnvNameNamespace define which namespace the taskset is running
	EnvNameNamespace = "TASKSET_NAMESPACE"
	// EnvNameTaskSetName define the key of TaskSet name which will be injected in  the pod
	EnvNameTaskSetName = "TASKSET_NAME"
	// EnvNameTaskRoleName define the key of TaskRole name which will be injected in the pod
	EnvNameTaskRoleName = "TASKROLE_NAME"
	// EnvNameTaskRoleReplicaIndex define the key of TaskRole replica index
	EnvNameTaskRoleReplicaIndex = "TASKROLE_REPLICA_INDEX"
)
