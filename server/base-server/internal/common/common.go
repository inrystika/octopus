package common

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	commapi "server/common/api/v1"

	v1 "k8s.io/api/core/v1"
)

type Mounts []*commapi.Mount

func (r Mounts) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Mounts) Scan(input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}

func GetVolumes(mounts Mounts) ([]v1.Volume, []v1.VolumeMount) {
	volumes := make([]v1.Volume, 0)
	volumeMounts := make([]v1.VolumeMount, 0)
	for i, m := range mounts {
		if m.Nfs != nil {
			name := fmt.Sprintf("mount%d", i)
			volumes = append(volumes, v1.Volume{
				Name: name,
				VolumeSource: v1.VolumeSource{
					NFS: &v1.NFSVolumeSource{
						Server:   m.Nfs.Server,
						Path:     m.Nfs.Path,
						ReadOnly: m.ReadOnly,
					}},
			})
			volumeMounts = append(volumeMounts, v1.VolumeMount{
				Name:      name,
				MountPath: m.ContainerPath,
				ReadOnly:  m.ReadOnly,
			})
		}
	}

	return volumes, volumeMounts
}
