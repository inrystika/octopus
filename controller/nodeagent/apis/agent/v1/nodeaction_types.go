/*
Copyright 2021.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NodeActionSpec defines the desired state of NodeAction
type NodeActionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// NodeName is where the action will execute
	NodeName string `json:"nodeName,omitempty"`
	Actions  Action `json:"actions"`
}

type Action struct {
	Docker *DockerAction `json:"docker,omitempty"`
}

type DockerAction struct {
	Commit        *DockerCommitCommand `json:"commit,omitempty"`
	CommitAndPush *DockerCommitCommand `json:"commitAndPush,omitempty"`
}

type DockerCommitCommand struct {
	Container  string   `json:"container,omitempty"`
	Repository string   `json:"repository,omitempty"`
	Tag        string   `json:"tag,omitempty"`
	Author     string   `json:"author,omitempty"`
	Message    string   `json:"message,omitempty"`
	Changes    []string `json:"changes"`
}

type CommandStatus struct {
	Name   string        `json:"name"`
	Result CommandResult `json:"result"`
	Reason string        `json:"reason"`
}

// NodeActionStatus defines the observed state of NodeAction
type NodeActionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	State   ActionState      `json:"state"`
	Actions []*CommandStatus `json:"actions,omitempty"`
}

//+genclient
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// NodeAction is the Schema for the nodeactions API
type NodeAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeActionSpec   `json:"spec,omitempty"`
	Status NodeActionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NodeActionList contains a list of NodeAction
type NodeActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeAction `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodeAction{}, &NodeActionList{})
}
