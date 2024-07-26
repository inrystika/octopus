package conf

import (
	v1 "k8s.io/api/core/v1"
)

var Storages []*StorageExtender

type StorageExtender struct {
	StorageType struct {
		v1.PersistentVolumeSource
	} `json:"storageType"`
	Requests string `json:"requests,omitempty"`
}
