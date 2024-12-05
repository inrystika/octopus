package common

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/util/rand"
	"path"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	commapi "server/common/api/v1"
	"strings"
	typeJob "volcano.sh/apis/pkg/apis/batch/v1alpha1"

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

func GetVolumes(mounts Mounts, octopusVolume string) ([]v1.Volume, []v1.VolumeMount) {
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

		if m.Octopus != nil {
			volumeMounts = append(volumeMounts, v1.VolumeMount{
				Name:      octopusVolume,
				MountPath: m.ContainerPath,
				SubPath:   path.Join(m.Octopus.Bucket, m.Octopus.Object),
				ReadOnly:  m.ReadOnly,
			})
		}
	}

	return volumes, volumeMounts
}

func GetCacheName() string {
	return "cache" + rand.String(10)

}

func GetTmpHomePath(userId string) string {
	return "/tmp/" + userId
}

func GetUserHomeVFName(userId string) string {
	return fmt.Sprintf("%s-userhome", userId)
}

func GetUserHomeVirtualPath() string {
	return fmt.Sprintf("/userhome")
}

func GetUserHomeMappedPath(userId string) string {
	return fmt.Sprintf("/minio/%s/%s", userId, USERHOME)
}

func GetExtraHomeVFName(idx int, userId string) string {
	return fmt.Sprintf("%s-extrahome%v", userId, idx)
}

func GetExtraHomeVirtualPath(idx int) string {
	return fmt.Sprintf("/extrahome%v", idx)
}

func GetExtraHomeMappedPath(idx int, userId string) string {
	return fmt.Sprintf("/storage%v/%s", idx, userId)
}

func AssignExtraHome(job *typeJob.Job) {
	volumes := make([]v1.Volume, 0)
	volumeMounts := make([]v1.VolumeMount, 0)
	for idx, _ := range conf.Storages {
		name := fmt.Sprintf("extrahome%v", idx)
		volumes = append(volumes, v1.Volume{
			Name: name,
			VolumeSource: v1.VolumeSource{
				PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
					ClaimName: GetStoragePersistentVolumeChaim(GetExtraHomeVFName(idx, job.Namespace)),
				},
			},
		})

		volumeMounts = append(volumeMounts, v1.VolumeMount{
			Name:      name,
			MountPath: GetExtraHomeVirtualPath(idx),
			SubPath:   job.Namespace,
			ReadOnly:  false,
		})
	}

	for i, task := range job.Spec.Tasks {
		job.Spec.Tasks[i].Template.Spec.Volumes = append(task.Template.Spec.Volumes, volumes...)
		job.Spec.Tasks[i].Template.Spec.Containers[0].VolumeMounts = append(task.Template.Spec.Containers[0].VolumeMounts, volumeMounts...)
	}
}

func BuildUserEndpoint(endpoint string) string {

	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	if !strings.HasSuffix(endpoint, "/") {
		endpoint = endpoint + "/"
	}

	return "/userendpoint" + endpoint
}

func GetCompany(
	ctx context.Context,
	resources *api.ResourceList,
	resourceSpec *api.ResourceSpec) (string, error) {

	companyResource := []string{"nvidia", "huawei", "cambricon", "enflame", "iluvatar", "metax-tech", "hygon"}
	for _, v := range companyResource {
		for _, r := range resources.Resources {
			for k, _ := range resourceSpec.ResourceQuantity {
				if r.Name == k {
					if strings.Contains(r.ResourceRef, v) || strings.Contains(r.Name, v) {
						return v, nil
					}
				}
			}
		}
	}
	return "", nil
}
