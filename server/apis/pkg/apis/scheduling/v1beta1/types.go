package v1beta1

import (
	"volcano.sh/apis/pkg/apis/scheduling/v1beta1"
)

// Queue is a queue of PodGroup.
type Queue struct {
	v1beta1.Queue

	Spec QueueSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// QueueSpec represents the template of Queue.
type QueueSpec struct {
	v1beta1.QueueSpec

	NodeSelector map[string]string `json:"nodeSelector,omitempty" protobuf:"bytes,4,opt,name=nodeSelector"`
}
