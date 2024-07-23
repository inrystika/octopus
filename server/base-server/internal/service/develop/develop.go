package develop

import (
	"context"
	"fmt"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"server/base-server/internal/service/algorithm"
	"server/common/constant"
	"server/common/errors"
	"server/common/utils"
	"server/common/utils/collections/set"
	"strconv"
	"strings"
	"time"

	"k8s.io/utils/strings/slices"

	vcBus "volcano.sh/apis/pkg/apis/bus/v1alpha1"

	"server/common/log"

	commapi "server/common/api/v1"

	nav1 "nodeagent/apis/agent/v1"

	typeJob "volcano.sh/apis/pkg/apis/batch/v1alpha1"

	"encoding/json"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/tools/cache"
	vcBatch "volcano.sh/apis/pkg/apis/batch/v1alpha1"
)

type developService struct {
	api.UnimplementedDevelopServer
	conf                *conf.Bootstrap
	log                 *log.Helper
	data                *data.Data
	workspaceService    api.WorkspaceServiceServer
	algorithmService    api.AlgorithmServiceServer
	imageService        api.ImageServiceServer
	datasetService      api.DatasetServiceServer
	resourceSpecService api.ResourceSpecServiceServer
	resourceService     api.ResourceServiceServer
	resourcePoolService api.ResourcePoolServiceServer
	billingService      api.BillingServiceServer
	userService         api.UserServiceServer
	updatedJob          chan *vcBatch.Job
	modelService        api.ModelServiceServer
}

type DevelopService interface {
	api.DevelopServer
}

const (
	k8sTaskNamePrefix            = "task"
	servicePort                  = 8888
	shmResource                  = "shm"
	cpuResource                  = "cpu"
	memResource                  = "memory"
	nodeActionLabelNotebookId    = "nodebook.octopus.dev/id"
	nodeActionLabelImageId       = "image.octopus.dev/id"
	kubeAnnotationsProxyBodySize = "nginx.ingress.kubernetes.io/proxy-body-size"
	envNotebookBaseUrl           = "OCTOPUS_NOTEBOOK_BASE_URL"
	envNotebookPort              = "OCTOPUS_NOTEBOOK_PORT"
)

func buildCommand(nbDir string) string {
	c := `! [ -x "$(command -v jupyter)" ] && pip install jupyterlab==3 -i https://pypi.tuna.tsinghua.edu.cn/simple;jupyter lab --no-browser --ip=0.0.0.0 --allow-root --notebook-dir='%s' --port=$%s --LabApp.token='' --LabApp.allow_origin='*' --LabApp.base_url=$%s;`
	return fmt.Sprintf(c, nbDir, envNotebookPort, envNotebookBaseUrl)
}

func buildTaskName(idx int) string {
	return fmt.Sprintf("%s%d", k8sTaskNamePrefix, idx)
}

func buildServiceName(jobId string, idx int) string {
	return fmt.Sprintf("%s-%s", jobId, buildTaskName(idx))
}

func buildIngressName(jobId string, idx int) string {
	return fmt.Sprintf("%s-%s", jobId, buildTaskName(idx))
}

func buildNotebookUrl(id string, idx int) string {
	return fmt.Sprintf("/notebook_%s_%s", id, buildTaskName(idx))
}

func buildUserEndpoint(endpoint string) string {

	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	if !strings.HasSuffix(endpoint, "/") {
		endpoint = endpoint + "/"
	}

	return "/userendpoint" + endpoint
}

func NewDevelopService(conf *conf.Bootstrap, logger log.Logger, data *data.Data,
	workspaceService api.WorkspaceServiceServer, algorithmService api.AlgorithmServiceServer,
	imageService api.ImageServiceServer, datasetService api.DatasetServiceServer, resourceSpecService api.ResourceSpecServiceServer,
	resourceService api.ResourceServiceServer, resourcePoolService api.ResourcePoolServiceServer,
	billingService api.BillingServiceServer, userService api.UserServiceServer, modelService api.ModelServiceServer) (DevelopService, error) {

	log := log.NewHelper("DevelopService", logger)

	s := &developService{
		conf:                conf,
		log:                 log,
		data:                data,
		workspaceService:    workspaceService,
		algorithmService:    algorithmService,
		imageService:        imageService,
		datasetService:      datasetService,
		resourceSpecService: resourceSpecService,
		resourceService:     resourceService,
		resourcePoolService: resourcePoolService,
		billingService:      billingService,
		userService:         userService,
		modelService:        modelService,
		updatedJob:          make(chan *vcBatch.Job, 1000),
	}

	s.data.Cluster.RegisterJobEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    s.onJobAdd,
		UpdateFunc: s.onJobUpdate,
		DeleteFunc: s.onJobDelete,
	})

	s.startNotebookTask()
	return s, nil
}

type closeFunc func(ctx context.Context) error

func (s *developService) checkPermAndAssign(ctx context.Context, nb *model.Notebook, nbJob *model.NotebookJob) (*startJobInfo, error) {
	user, err := s.userService.FindUser(ctx, &api.FindUserRequest{Id: nb.UserId})
	if err != nil {
		return nil, err
	}
	queue := nb.ResourcePool
	if nb.WorkspaceId == constant.SYSTEM_WORKSPACE_DEFAULT {
		if !utils.StringSliceContainsValue(user.User.ResourcePools, queue) {
			return nil, errors.Errorf(nil, errors.ErrorNotebookResourcePoolForbidden)
		}
	} else {
		workspace, err := s.workspaceService.GetWorkspace(ctx, &api.GetWorkspaceRequest{WorkspaceId: nb.WorkspaceId})
		if err != nil {
			return nil, err
		}

		if queue != workspace.Workspace.ResourcePoolId {
			return nil, errors.Errorf(nil, errors.ErrorNotebookResourcePoolForbidden)
		}
	}

	imageAddr := ""
	if nb.ImageId != "" { //判空，允许通过API调用不传此参数
		image, err := s.imageService.FindImage(ctx, &api.FindImageRequest{ImageId: nb.ImageId})
		if err != nil {
			return nil, err
		}

		if image.Image == nil {
			return nil, errors.Errorf(nil, errors.ErrorImageNotExist)
		}

		if nb.UserId != image.Image.UserId && image.Image.IsPrefab == api.ImageIsPrefab_IMAGE_IS_PREFAB_NO {
			hasPerm := false
			for _, i := range image.Accesses {
				if nb.WorkspaceId == i.SpaceId {
					hasPerm = true
				}
			}

			if !hasPerm {
				return nil, errors.Errorf(err, errors.ErrorNotebookImageNoPermission)
			}
		}

		if image.Image.ImageStatus != api.ImageStatus_IMAGE_STATUS_MADE {
			return nil, errors.Errorf(nil, errors.ErrorNotebookImageStatusForbidden)
		}
		nb.ImageName = image.Image.ImageName
		nb.ImageVersion = image.Image.ImageVersion
		imageAddr = image.Image.ImageFullAddr
	} else if nb.ImageUrl != "" {
		imageAddr = nb.ImageUrl
	} else {
		return nil, errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
	}

	algorithmPath := ""
	if nb.AlgorithmId != "" { //判空，允许通过API调用不传此参数
		algorithmVersion, err := s.algorithmService.QueryAlgorithmVersion(ctx, &api.QueryAlgorithmVersionRequest{
			AlgorithmId: nb.AlgorithmId,
			Version:     nb.AlgorithmVersion,
		})
		if err != nil {
			return nil, err
		}

		if nb.UserId != algorithmVersion.Algorithm.UserId {
			return nil, errors.Errorf(err, errors.ErrorNotebookAlgorithmNoPermission)
		}

		if algorithmVersion.Algorithm.FileStatus != int64(algorithm.FILESTATUS_FINISH) {
			return nil, errors.Errorf(err, errors.ErrorNotebookAlgorithmStatusForbidden)
		}
		nb.AlgorithmName = algorithmVersion.Algorithm.AlgorithmName
		algorithmPath = algorithmVersion.Algorithm.Path
	}

	preTrainModelPath := ""
	if nb.PreTrainModelId != "" {
		modelVersion, err := s.modelService.QueryModelVersion(ctx, &api.QueryModelVersionRequest{
			ModelId: nb.PreTrainModelId,
			Version: nb.PreTrainModelVersion,
		})
		if err != nil {
			return nil, err
		}

		nb.PreTrainModelName = modelVersion.Model.ModelName
		preTrainModelPath = modelVersion.ModelVersion.Path
	}

	datasetPath := ""
	if nb.DatasetId != "" && nb.DatasetVersion != "" {
		datasetVersion, err := s.datasetService.GetDatasetVersion(ctx, &api.GetDatasetVersionRequest{DatasetId: nb.DatasetId, Version: nb.DatasetVersion})
		if err != nil {
			return nil, err
		}
		if nb.UserId != datasetVersion.Dataset.UserId && datasetVersion.Dataset.SourceType == api.DatasetSourceType_DST_USER {
			hasPerm := false
			for _, i := range datasetVersion.VersionAccesses {
				if nb.WorkspaceId == i.SpaceId {
					hasPerm = true
					break
				}
			}

			if !hasPerm {
				return nil, errors.Errorf(err, errors.ErrorNotebookDatasetNoPermission)
			}
		}

		if datasetVersion.Version.Status != int32(api.DatasetVersionStatus_DVS_Unzipped) {
			return nil, errors.Errorf(err, errors.ErrorNotebookDatasetStatusForbidden)
		}
		nb.DatasetName = datasetVersion.Dataset.Name
		datasetPath = datasetVersion.Version.Path
	}

	spec, err := s.resourceSpecService.GetResourceSpec(ctx, &api.GetResourceSpecRequest{Id: nb.ResourceSpecId})
	if err != nil {
		return nil, err
	}
	nbJob.ResourceSpecPrice = spec.ResourceSpec.Price
	if spec.ResourceSpec.Price == 0 {
		nbJob.PayStatus = api.BillingPayRecordStatus_BPRS_PAY_COMPLETED
	} else {
		nbJob.PayStatus = api.BillingPayRecordStatus_BPRS_PAYING
	}
	nb.ResourceSpecName = spec.ResourceSpec.Name

	if spec.ResourceSpec.Price > 0 {
		ownerId, ownerType := s.getOwner(nb)
		owner, err := s.billingService.GetBillingOwner(ctx, &api.GetBillingOwnerRequest{
			OwnerId:   ownerId,
			OwnerType: ownerType,
		})
		if err != nil {
			return nil, err
		}

		if owner.BillingOwner.Amount <= 0 {
			return nil, errors.Errorf(nil, errors.ErrorNotebookBalanceNotEnough)
		}
	}

	resources, err := s.resourceService.ListResource(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	k8sResources := map[v1.ResourceName]resource.Quantity{}
	nodeSelectors := map[string]string{}
	var shm *resource.Quantity = nil
	for _, r := range resources.Resources {
		for k, v := range spec.ResourceSpec.ResourceQuantity {
			quantity, err := resource.ParseQuantity(v)
			if err != nil {
				return nil, errors.Errorf(err, errors.ErrorNotebookParseResourceSpecFailed)
			}
			if r.Name == k {
				if r.ResourceRef == shmResource || r.Name == shmResource {
					shm = &quantity
					continue
				}
				if r.ResourceRef == "" {
					k8sResources[v1.ResourceName(r.Name)] = quantity
				} else {
					k8sResources[v1.ResourceName(r.ResourceRef)] = quantity
					nodeSelectors[r.BindingNodeLabelKey] = r.BindingNodeLabelValue
				}
			}
		}
	}

	for _, m := range nb.Mounts {
		if m.Octopus != nil {
			if !slices.Contains(user.User.Buckets, m.Octopus.Bucket) && m.Octopus.Bucket != nb.UserId {
				return nil, errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
			}
		}

		if m.Nfs != nil && (user.User.Permission == nil || !user.User.Permission.MountExternalStorage) {
			return nil, errors.Errorf(nil, errors.ErrorNotebookMountExternalForbidden)
		}
	}

	command := ""
	if nb.Command != "" {
		command = nb.Command
	} else {
		if nb.AlgorithmId != "" {
			command = buildCommand(s.conf.Service.DockerCodePath)
		} else {
			command = buildCommand("/")
		}
	}

	return &startJobInfo{
		queue:             queue,
		imageAddr:         imageAddr,
		algorithmPath:     algorithmPath,
		datasetPath:       datasetPath,
		preTrainModelPath: preTrainModelPath,
		resources:         k8sResources,
		nodeSelectors:     nodeSelectors,
		shm:               shm,
		command:           command,
	}, nil
}

func (s *developService) startJob(ctx context.Context, nb *model.Notebook, nbJob *model.NotebookJob, startJobInfo *startJobInfo) (closeFunc, error) {
	var err error
	closes := make([]closeFunc, 0)
	resFunc := func(ctx context.Context) error {
		var err2 error
		for i := len(closes) - 1; i >= 0; i-- {
			err1 := closes[i](ctx)
			if err1 != nil {
				err2 = err1
				s.log.Errorf(ctx, "err: %s", err1)
			}
		}

		return err2
	}

	defer func() {
		if err != nil {
			err1 := resFunc(ctx)
			if err1 != nil {
				s.log.Errorf(ctx, "err: %s", err1)
			}
		}
	}()

	err = s.submitJob(ctx, nb, nbJob, startJobInfo)
	if err != nil {
		return nil, err
	}

	closes = append(closes, func(ctx context.Context) error {
		err1 := s.data.Cluster.DeleteJob(ctx, nb.UserId, nbJob.Id)
		return err1
	})

	err = s.createService(ctx, nb, nbJob)
	if err != nil {
		return nil, err
	}
	closes = append(closes, func(ctx context.Context) error {
		return s.deleteService(ctx, nb, nbJob)
	})

	err = s.createIngress(ctx, nb, nbJob)
	if err != nil {
		return nil, err
	}
	closes = append(closes, func(ctx context.Context) error {
		return s.deleteIngress(ctx, nb, nbJob)
	})

	return resFunc, nil
}

type startJobInfo struct {
	queue             string
	imageAddr         string
	algorithmPath     string
	datasetPath       string
	preTrainModelPath string
	command           string
	resources         map[v1.ResourceName]resource.Quantity
	nodeSelectors     map[string]string
	shm               *resource.Quantity
}

func (s *developService) checkEndpoint(ctx context.Context, nb *model.Notebook) ([]*model.UserEndpoint, error) {
	if len(nb.TaskConfigs) == 0 {
		return nil, nil
	}

	endpoints := set.NewStrings()
	endpointsTb := make([]*model.UserEndpoint, 0)
	for _, taskConfigs := range nb.TaskConfigs {
		for _, endpoint := range taskConfigs.Endpoints {
			ue := buildUserEndpoint(endpoint.Endpoint)
			if endpoints.Contains(ue) {
				return nil, errors.Errorf(nil, errors.ErrorUserEndpointRepeat)
			}
			if endpoint.Port == servicePort {
				return nil, errors.Errorf(nil, errors.ErrorPortConflict)
			}
			endpoints.Add(ue)
			endpointsTb = append(endpointsTb, &model.UserEndpoint{
				Endpoint: ue,
			})
		}
	}

	notExist, err := s.data.UserEndpointDao.IsUserEndpointsNotExist(ctx, endpoints.Values())
	if err != nil {
		return nil, err
	}

	if !notExist {
		return nil, errors.Errorf(nil, errors.ErrorUserEndpointExisted)
	}

	return endpointsTb, nil
}

func (s *developService) deleteUserEndpoint(ctx context.Context, nb *model.Notebook) error {
	endpoints := make([]string, 0)
	for _, taskEndpoints := range nb.TaskConfigs {
		for _, endpoint := range taskEndpoints.Endpoints {
			endpoints = append(endpoints, buildUserEndpoint(endpoint.Endpoint))
		}
	}

	err := s.data.UserEndpointDao.DeleteUserEndpoints(ctx, endpoints)
	if err != nil {
		return err
	}

	return nil
}

func (s *developService) CreateNotebook(ctx context.Context, req *api.CreateNotebookRequest) (*api.CreateNotebookReply, error) {
	nbId := utils.GetUUIDStartWithAlphabetic() //k8s service首字母不允许数字 为方便 uuid处理一下
	_, size, err := s.data.DevelopDao.ListNotebook(ctx, &model.NotebookQuery{
		UserId:      req.UserId,
		WorkspaceId: req.WorkspaceId,
		Name:        req.Name,
	})
	if size > 0 {
		return nil, errors.Errorf(nil, errors.ErrorNotebookRepeat)
	}

	jobId := utils.GetUUIDStartWithAlphabetic()

	nb := &model.Notebook{}
	err = copier.Copy(nb, req)
	if err != nil {
		return nil, err
	}
	nb.Id = nbId
	nb.Status = constant.PREPARING
	nb.NotebookJobId = jobId
	nb.TaskNumber = int(req.TaskNumber)
	if req.AutoStopDuration == 0 {
		nb.AutoStopDuration = s.conf.Service.Develop.AutoStopIntervalSec
	}

	nbJob := &model.NotebookJob{
		Id:         jobId,
		NotebookId: nbId,
		Status:     constant.PREPARING,
		Detail:     "{}",
	}

	startJobInfo, err := s.checkPermAndAssign(ctx, nb, nbJob)
	if err != nil {
		return nil, err
	}

	endpointsTb, err := s.checkEndpoint(ctx, nb)
	if err != nil {
		return nil, err
	}

	closeFunc, err := s.startJob(ctx, nb, nbJob, startJobInfo)
	defer func() { //如果出错 重要的资源需要删除
		if err != nil && closeFunc != nil {
			err1 := closeFunc(ctx)
			if err1 != nil {
				s.log.Errorf(ctx, "err: %s", err1)
			}
		}
	}()
	if err != nil {
		return nil, err
	}

	err = s.data.UserEndpointDao.CreateUserEndpoints(ctx, endpointsTb)
	if err != nil {
		return nil, err
	}

	err = s.data.DevelopDao.CreateNotebook(ctx, nb)
	if err != nil {
		return nil, err
	}

	err = s.data.DevelopDao.CreateNotebookJob(ctx, nbJob)
	if err != nil {
		return nil, err
	}

	err1 := s.data.DevelopDao.CreateNotebookEventRecord(ctx, &model.NotebookEventRecord{
		Time:       time.Now(),
		NotebookId: nb.Id,
		Type:       commapi.NotebookEventRecordType_CREATE,
	})
	if err1 != nil { // 插入事件记录出错只打印
		s.log.Error(ctx, "create notebook event record error:", err)
	}

	return &api.CreateNotebookReply{Id: nbId}, nil
}

func (s *developService) StartNotebook(ctx context.Context, req *api.StartNotebookRequest) (*api.StartNotebookReply, error) {
	nb, err := s.data.DevelopDao.GetNotebook(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if !utils.IsCompletedState(nb.Status) {
		return nil, errors.Errorf(nil, errors.ErrorNotebookStatusForbidden)
	}

	jobId := utils.GetUUIDStartWithAlphabetic()
	nbJob := &model.NotebookJob{
		Id:         jobId,
		NotebookId: nb.Id,
		Status:     constant.PREPARING,
		Detail:     "{}",
	}

	startJobInfo, err := s.checkPermAndAssign(ctx, nb, nbJob)
	if err != nil {
		return nil, err
	}

	err = s.data.DevelopDao.CreateNotebookJob(ctx, nbJob)
	if err != nil {
		return nil, err
	}

	if req.AutoStopDuration != 0 {
		nb.AutoStopDuration = req.AutoStopDuration
	}

	err = s.data.DevelopDao.UpdateNotebookSelective(ctx, &model.Notebook{
		Id:               nb.Id,
		NotebookJobId:    jobId,
		Status:           constant.PREPARING,
		AutoStopDuration: nb.AutoStopDuration,
	})
	if err != nil {
		return nil, err
	}

	//数据库操作挪到前面，如果出错，直接不创建k8s vcjob，硬件资源有限的资源，出错需要及时释放掉
	closeFunc, err := s.startJob(ctx, nb, nbJob, startJobInfo)
	defer func() {
		if err != nil {
			err1 := closeFunc(ctx)
			if err1 != nil {
				s.log.Errorf(ctx, "err: %s", err1)
			}
		}
	}()

	err1 := s.data.DevelopDao.CreateNotebookEventRecord(ctx, &model.NotebookEventRecord{
		Time:       time.Now(),
		NotebookId: nb.Id,
		Type:       commapi.NotebookEventRecordType_START,
	})
	if err1 != nil { // 插入事件记录出错只打印
		s.log.Error(ctx, "create notebook event record error:", err)
	}

	return &api.StartNotebookReply{Id: req.Id}, nil
}

func (s *developService) submitJob(ctx context.Context, nb *model.Notebook, nbJob *model.NotebookJob, startJobInfo *startJobInfo) error {
	volume := "data"
	volumeMounts := []v1.VolumeMount{

		{
			Name:      "localtime",
			MountPath: "/etc/localtime",
		},
	}

	if !nb.DisableMountUserHome {
		volumeMounts = append(volumeMounts, v1.VolumeMount{
			Name:      volume,
			MountPath: s.conf.Service.DockerUserHomePath,
			SubPath:   common.GetUserHomePath(nb.UserId),
			ReadOnly:  false,
		})
	}

	if startJobInfo.algorithmPath != "" {
		volumeMounts = append(volumeMounts, v1.VolumeMount{
			Name:      volume,
			MountPath: s.conf.Service.DockerCodePath,
			SubPath:   startJobInfo.algorithmPath,
			ReadOnly:  false,
		})
	}

	if startJobInfo.datasetPath != "" {
		volumeMounts = append(volumeMounts, v1.VolumeMount{
			Name:      volume,
			MountPath: s.conf.Service.DockerDatasetPath,
			SubPath:   startJobInfo.datasetPath,
			ReadOnly:  true,
		})
	}

	if startJobInfo.preTrainModelPath != "" {
		volumeMounts = append(volumeMounts, v1.VolumeMount{
			Name:      volume,
			MountPath: common.DockerPreTrainModePath,
			SubPath:   startJobInfo.preTrainModelPath,
			ReadOnly:  true,
		})
	}

	volumes := []v1.Volume{
		{
			Name: volume,
			VolumeSource: v1.VolumeSource{
				PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
					ClaimName: common.GetStoragePersistentVolumeChaim(nb.UserId),
				},
			},
		},
		{
			Name: "localtime",
			VolumeSource: v1.VolumeSource{
				HostPath: &v1.HostPathVolumeSource{
					Path: "/etc/localtime",
				}},
		},
	}

	if startJobInfo.shm != nil {
		volumeMounts = append(volumeMounts, v1.VolumeMount{
			Name:      "cache-volume",
			MountPath: "/dev/shm",
		})
		volumes = append(volumes, v1.Volume{
			Name: "cache-volume",
			VolumeSource: v1.VolumeSource{
				EmptyDir: &v1.EmptyDirVolumeSource{
					Medium:    v1.StorageMediumMemory,
					SizeLimit: startJobInfo.shm,
				},
			},
		})
	}

	vs, vms := common.GetVolumes(nb.Mounts, volume)
	if len(vms) > 0 {
		volumeMounts = append(volumeMounts, vms...)
		volumes = append(volumes, vs...)
	}

	for k, _ := range startJobInfo.resources {
		if strings.HasPrefix(string(k), common.DCUResourceName) {
			volumeMounts = append(volumeMounts, v1.VolumeMount{
				Name:      "hyhal",
				MountPath: "/opt/hyhal",
			})
			volumes = append(volumes, v1.Volume{
				Name: "hyhal",
				VolumeSource: v1.VolumeSource{
					HostPath: &v1.HostPathVolumeSource{Path: "/opt/hyhal"},
				},
			})
		}
	}

	tasks := make([]typeJob.TaskSpec, 0)
	for i := 0; i < nb.TaskNumber; i++ {
		taskName := buildTaskName(i)
		task := typeJob.TaskSpec{}
		task.Name = taskName
		task.Replicas = 1
		envs := []v1.EnvVar{{
			Name:  envNotebookBaseUrl,
			Value: buildNotebookUrl(nb.Id, i),
		}, {
			Name:  envNotebookPort,
			Value: strconv.Itoa(servicePort),
		}}
		for k, v := range nb.Envs {
			envs = append(envs, v1.EnvVar{Name: k, Value: v})
		}
		task.Template = v1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{"volcano.sh/task-spec": buildTaskName(i)},
			},
			Spec: v1.PodSpec{
				RestartPolicy: v1.RestartPolicyNever,
				Containers: []v1.Container{
					{
						Name:    taskName,
						Image:   startJobInfo.imageAddr,
						Command: []string{"sh", "-c", startJobInfo.command},
						Resources: v1.ResourceRequirements{
							Requests: startJobInfo.resources,
							Limits:   startJobInfo.resources,
						},
						VolumeMounts: volumeMounts,
						Env:          envs,
					},
				},
				NodeSelector: startJobInfo.nodeSelectors,
				Volumes:      volumes,
			},
		}
		if s.conf.Service.IsUseMultusCNI {
			task.Template.Annotations =
				map[string]string{"k8s.v1.cni.cncf.io/networks": s.conf.Service.NetworksConf}
		}

		for k, _ := range startJobInfo.resources {
			if strings.HasPrefix(string(k), common.RdmaPrefix) {
				task.Template.Spec.Containers[0].SecurityContext = &v1.SecurityContext{
					Capabilities: &v1.Capabilities{
						Add: []v1.Capability{"IPC_LOCK"},
					},
				}
			}
		}
		tasks = append(tasks, task)
	}

	Job := &typeJob.Job{}
	Job.TypeMeta = metav1.TypeMeta{
		APIVersion: "batch.volcano.sh/v1alpha1",
		Kind:       "Job",
	}
	Job.ObjectMeta = metav1.ObjectMeta{
		Namespace: nb.UserId,
		Name:      nbJob.Id,
		Annotations: map[string]string{
			constant.JOB_TYPE: constant.NotebookJob,
		},
	}
	Job.Spec = typeJob.JobSpec{}
	Job.Spec.MinAvailable = int32(nb.TaskNumber)
	Job.Spec.Queue = startJobInfo.queue
	Job.Spec.SchedulerName = "volcano"
	//打开后在nginx的node上ping其他node的pod，网络不通导致jupyter打不开，先屏蔽
	Job.Spec.Plugins = map[string][]string{
		"env": {},
		"svc": {"--disable-network-policy=true"},
	}
	Job.Spec.Policies = []vcBatch.LifecyclePolicy{
		{Event: vcBus.PodEvictedEvent, Action: vcBus.RestartJobAction},
		{Event: vcBus.PodFailedEvent, Action: vcBus.RestartJobAction},
	}
	Job.Spec.Tasks = tasks
	common.AssignExtraHome(Job)

	err := s.data.Cluster.CreateJob(ctx, Job)
	if err != nil {
		return err
	}
	return nil
}

func (s *developService) StopNotebook(ctx context.Context, req *api.StopNotebookRequest) (*api.StopNotebookReply, error) {
	nb, err := s.data.DevelopDao.GetNotebook(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	nbJob, err := s.data.DevelopDao.GetNotebookJob(ctx, nb.NotebookJobId)
	if err != nil {
		return nil, err
	}

	if utils.IsCompletedState(nb.Status) {
		return nil, errors.Errorf(nil, errors.ErrorNotebookStatusForbidden)
	}

	err = s.deleteIngress(ctx, nb, nbJob)
	if err != nil {
		s.log.Errorw(ctx, "err", err)
	}

	err = s.deleteService(ctx, nb, nbJob)
	if err != nil {
		s.log.Errorw(ctx, "err", err)
	}

	err = s.data.Cluster.DeleteJob(ctx, nb.UserId, nbJob.Id)
	if err != nil {
		s.log.Errorw(ctx, "err", err)
	}

	err = s.data.DevelopDao.UpdateNotebookSelective(ctx, &model.Notebook{
		Id:     nb.Id,
		Status: constant.STOPPED,
	})
	if err != nil {
		return nil, err
	}

	s.sendEmail(nb.UserId, nb.Name, nb.Status, constant.STOPPED)

	now := time.Now()
	err = s.data.DevelopDao.UpdateNotebookJobSelective(ctx, &model.NotebookJob{
		Id:        nbJob.Id,
		Status:    constant.STOPPED,
		StoppedAt: &now,
		Operation: req.Operation,
	})
	if err != nil {
		return nil, err
	}

	err = s.data.DevelopDao.CreateNotebookEventRecord(ctx, &model.NotebookEventRecord{
		Time:       time.Now(),
		NotebookId: nb.Id,
		Type:       commapi.NotebookEventRecordType_STOP,
	})
	if err != nil { // 插入事件记录出错只打印
		s.log.Error(ctx, "create notebook event record error:", err)
	}

	return &api.StopNotebookReply{Id: req.Id}, nil
}

func (s *developService) createService(ctx context.Context, nb *model.Notebook, nbJob *model.NotebookJob) error {
	for i := 0; i < nb.TaskNumber; i++ {
		ports := []v1.ServicePort{{
			Port:       servicePort,
			TargetPort: intstr.FromInt(servicePort),
			Name:       fmt.Sprintf("port-%d", servicePort),
		}}
		if len(nb.TaskConfigs) > i {
			for _, endpoint := range nb.TaskConfigs[i].Endpoints {
				ports = append(ports, v1.ServicePort{
					Port:       int32(endpoint.Port),
					TargetPort: intstr.FromInt(int(endpoint.Port)),
					Name:       fmt.Sprintf("port-%d", int32(endpoint.Port)),
				})
			}
		}

		err := s.data.Cluster.CreateService(ctx, &v1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      buildServiceName(nbJob.Id, i),
				Namespace: nb.UserId,
			},
			Spec: v1.ServiceSpec{
				Selector: map[string]string{
					"volcano.sh/task-spec": buildTaskName(i),
					"volcano.sh/job-name":  nbJob.Id,
				},
				Ports: ports,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *developService) deleteService(ctx context.Context, nb *model.Notebook, nbJob *model.NotebookJob) error {
	for i := 0; i < nb.TaskNumber; i++ {
		err := s.data.Cluster.DeleteService(ctx, nb.UserId, buildServiceName(nbJob.Id, i))
		if err != nil {
			return err
		}

	}

	return nil
}

func (s *developService) createIngress(ctx context.Context, nb *model.Notebook, nbJob *model.NotebookJob) error {
	for i := 0; i < nb.TaskNumber; i++ {
		var upLoadFileSize string = ""
		if s.conf.Service.Develop.IsSetUploadFileSize {
			upLoadFileSize = "1000m" // 为空时jupyter文件上传大小不能超过1M，非空时不限制上传文件大小
		}
		rules := []v1beta1.IngressRule{
			{
				IngressRuleValue: v1beta1.IngressRuleValue{
					HTTP: &v1beta1.HTTPIngressRuleValue{
						Paths: []v1beta1.HTTPIngressPath{
							{
								Path: buildNotebookUrl(nb.Id, i),
								Backend: v1beta1.IngressBackend{
									ServiceName: buildServiceName(nbJob.Id, i),
									ServicePort: intstr.FromInt(servicePort),
								},
							},
						},
					},
				},
			},
		}
		if len(nb.TaskConfigs) > i {
			for _, endpoint := range nb.TaskConfigs[i].Endpoints {
				rules = append(rules, v1beta1.IngressRule{
					IngressRuleValue: v1beta1.IngressRuleValue{
						HTTP: &v1beta1.HTTPIngressRuleValue{
							Paths: []v1beta1.HTTPIngressPath{
								{
									Path: buildUserEndpoint(endpoint.Endpoint),
									Backend: v1beta1.IngressBackend{
										ServiceName: buildServiceName(nbJob.Id, i),
										ServicePort: intstr.FromInt(int(endpoint.Port)),
									},
								},
							},
						},
					},
				})
			}
		}

		err := s.data.Cluster.CreateIngress(ctx, &v1beta1.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      buildIngressName(nbJob.Id, i),
				Namespace: nb.UserId,
				Annotations: map[string]string{
					kubeAnnotationsProxyBodySize: upLoadFileSize,
				},
			},
			Spec: v1beta1.IngressSpec{
				Rules: rules,
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *developService) deleteIngress(ctx context.Context, nb *model.Notebook, nbJob *model.NotebookJob) error {
	for i := 0; i < nb.TaskNumber; i++ {
		err := s.data.Cluster.DeleteIngress(ctx, nb.UserId, buildIngressName(nbJob.Id, i))
		if err != nil {
			return err
		}

	}

	return nil
}

func (s *developService) sendEmail(userId string, jobName string, oldStatus string, newStatus string) {
	ctx := context.TODO()
	go utils.HandlePanic(ctx, func(i ...interface{}) {
		user, err := s.userService.FindUser(ctx, &api.FindUserRequest{Id: userId})
		if err != nil {
			log.Errorf(ctx, "FindUser err when sendEmail:"+userId, err)
			return
		}
		if *user.User.EmailNotify && !strings.EqualFold(oldStatus, newStatus) && utils.IsNotifyState(newStatus) {
			common.SendEmail(s.conf.Service.AdminEmail, user.User.Email, fmt.Sprintf("Notebook %s %s", jobName, newStatus))
		}
	})()
}

func (s *developService) DeleteNotebook(ctx context.Context, req *api.DeleteNotebookRequest) (*api.DeleteNotebookReply, error) {
	nb, err := s.data.DevelopDao.GetNotebook(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if !utils.IsCompletedState(nb.Status) {
		return nil, errors.Errorf(nil, errors.ErrorNotebookStatusForbidden)
	}

	err = s.data.DevelopDao.DeleteNotebook(ctx, nb.Id)
	if err != nil {
		return nil, err
	}

	err = s.data.DevelopDao.DeleteNotebookJobByNbId(ctx, nb.Id)
	if err != nil {
		return nil, err
	}

	err = s.deleteUserEndpoint(ctx, nb)
	if err != nil {
		return nil, err
	}
	return &api.DeleteNotebookReply{Id: req.Id}, nil
}

func (s *developService) ListNotebook(ctx context.Context, req *api.ListNotebookRequest) (*api.ListNotebookReply, error) {
	query := &model.NotebookQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	notebooksTbl, totalSize, err := s.data.DevelopDao.ListNotebook(ctx, query)
	if err != nil {
		return nil, err
	}

	notebooks, err := s.convertNotebook(ctx, notebooksTbl)
	if err != nil {
		return nil, err
	}

	return &api.ListNotebookReply{
		TotalSize: totalSize,
		Notebooks: notebooks,
	}, nil
}

func (s *developService) convertNotebook(ctx context.Context, notebooksTbl []*model.Notebook) ([]*api.Notebook, error) {
	jobIds := make([]string, 0)
	for _, n := range notebooksTbl {
		jobIds = append(jobIds, n.NotebookJobId)
	}
	notebookJobs, err := s.data.DevelopDao.ListNotebookJob(ctx, &model.NotebookJobQuery{Ids: jobIds})
	if err != nil {
		return nil, err
	}
	priceMap := make(map[string]float64)
	jobMap := make(map[string]*model.NotebookJob)
	for _, j := range notebookJobs {
		priceMap[j.Id] = j.ResourceSpecPrice
		jobMap[j.Id] = j
	}

	notebooks := make([]*api.Notebook, 0)
	for _, n := range notebooksTbl {
		notebook := &api.Notebook{}
		err := copier.Copy(notebook, n)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorStructCopy)
		}
		notebook.CreatedAt = n.CreatedAt.Unix()
		notebook.UpdatedAt = n.UpdatedAt.Unix()
		notebook.ResourceSpecPrice = priceMap[n.NotebookJobId]
		for i := 0; i < n.TaskNumber; i++ {
			notebook.Tasks = append(notebook.Tasks, &api.Notebook_Task{Name: buildTaskName(i), Url: buildNotebookUrl(n.Id, i)})
		}
		notebook.ExitMsg = s.getExitMsg(ctx, jobMap[n.NotebookJobId])
		notebook.Operation = jobMap[n.NotebookJobId].Operation
		if jobMap[n.NotebookJobId].StartedAt != nil {
			notebook.StartedAt = jobMap[n.NotebookJobId].StartedAt.Unix()
		} else {
			notebook.StartedAt = 0
		}
		if jobMap[n.NotebookJobId].StoppedAt != nil {
			notebook.StoppedAt = jobMap[n.NotebookJobId].StoppedAt.Unix()
		} else {
			notebook.StoppedAt = 0
		}
		notebooks = append(notebooks, notebook)
	}
	return notebooks, nil
}

func (s *developService) getExitMsg(ctx context.Context, job *model.NotebookJob) string {
	detail := &typeJob.JobStatusDetail{}
	err := json.Unmarshal([]byte(job.Detail), detail)
	if err != nil || detail.Job == nil {
		return ""
	}

	return detail.Job.ExitDiagnostics
}

func (s *developService) GetNotebook(ctx context.Context, req *api.GetNotebookRequest) (*api.GetNotebookReply, error) {
	notebookTbl, err := s.data.DevelopDao.GetNotebook(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	notebooks, err := s.convertNotebook(ctx, []*model.Notebook{notebookTbl})
	if err != nil {
		return nil, err
	}

	return &api.GetNotebookReply{Notebook: notebooks[0]}, nil
}

func (s *developService) GetNotebookEventList(ctx context.Context, req *api.NotebookEventListRequest) (*api.NotebookEventListReply, error) {

	query := &model.NotebookEventQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, err
	}

	if query.Id == "" {
		if req.NotebookId == "" {
			s.log.Errorf(ctx, "job id and notebook id empty")
			return nil, errors.Errorf(nil, errors.ErrorInvalidRequestParameter)
		}
		nbIds := make([]string, 0)
		nbIds = append(nbIds, req.NotebookId)
		nbIds = set.NewStrings(nbIds...).Values()
		nbs, _, err := s.data.DevelopDao.ListNotebook(ctx, &model.NotebookQuery{Ids: nbIds})
		if err != nil {
			s.log.Errorf(ctx, "ListNotebook err: %s", err)
			return nil, err
		}
		if len(nbs) > 0 {
			for _, nb := range nbs {
				query.Id = nb.NotebookJobId
			}
		} else {
			s.log.Errorf(ctx, "no notebook job found")
			return nil, fmt.Errorf("no notebook job found")
		}
	}

	events, totalSize, err := s.data.DevelopDao.GetNotebookEvents(query)
	if err != nil {
		return nil, err
	}

	notebookEvents := make([]*api.NotebookEvent, 0)

	for _, value := range events {
		event := &api.NotebookEvent{}
		event.Timestamp = value.Timestamp
		event.Name = value.Name
		event.Reason = value.Reason
		event.Message = value.Message
		notebookEvents = append(notebookEvents, event)
	}

	return &api.NotebookEventListReply{
		TotalSize:      totalSize,
		NotebookEvents: notebookEvents,
	}, nil
}

func (s *developService) SaveNotebook(ctx context.Context, req *api.SaveNotebookRequest) (*api.SaveNotebookReply, error) {
	notebook, err := s.data.DevelopDao.GetNotebook(ctx, req.NotebookId)
	if err != nil {
		return nil, err
	}
	if !utils.JobRunningState(notebook.Status) {
		return nil, errors.Errorf(nil, errors.ErrorNotebookStatusForbidden)
	}

	// check saveNotebook action existed
	podName := s.GetPodNameFromNoteBookTask(notebook, req.TaskName)
	nodeAction, err := s.data.Cluster.GetNodeAction(ctx, notebook.UserId, podName)
	if err != nil {
		return nil, err
	}
	if nodeAction != nil {
		return nil, errors.Errorf(nil, errors.ErrorNotebookRepeatedToSave)
	}

	nodeName, containerId, err := s.getNotebookTaskContainer(ctx, notebook, req.TaskName)
	if err != nil {
		return nil, err
	}

	imageReply, err := s.imageService.AddImage(ctx, &api.AddImageRequest{
		ImageName:    req.ImageName,
		ImageVersion: req.ImageVersion,
		UserId:       notebook.UserId,
		SpaceId:      notebook.WorkspaceId,
		IsPrefab:     api.ImageIsPrefab_IMAGE_IS_PREFAB_NO,
		SourceType:   api.ImageSourceType_IMAGE_SOURCE_TYPE_SAVED,
		ImageDesc:    req.LayerDescription,
	})
	if err != nil {
		return nil, err
	}
	nodeAction = &nav1.NodeAction{
		ObjectMeta: metav1.ObjectMeta{
			Name: podName,
			Labels: map[string]string{
				nodeActionLabelNotebookId: req.NotebookId,
				nodeActionLabelImageId:    imageReply.ImageId,
			},
		},
		Spec: nav1.NodeActionSpec{
			NodeName: nodeName,
			Actions: nav1.Action{
				Docker: &nav1.DockerAction{
					CommitAndPush: &nav1.DockerCommitCommand{
						Container:  containerId,
						Repository: fmt.Sprintf("%s/%s", s.conf.Data.Harbor.Host, imageReply.ImageAddr),
						Tag:        req.ImageVersion,
						Author:     notebook.UserId,
						Message:    req.LayerDescription,
						Changes:    []string{},
					},
				},
			},
		},
	}
	_, err = s.imageService.UpdateImage(ctx, &api.UpdateImageRequest{
		ImageId:     imageReply.ImageId,
		ImageStatus: api.ImageStatus_IMAGE_STATUS_MAKING,
	})
	if err != nil {
		s.log.Errorw(ctx, err)
		return nil, err
	}
	if _, err := s.data.Cluster.CreateNodeAction(ctx, notebook.UserId, nodeAction); err != nil {
		return nil, err
	}

	// acc node agent to commit image
	return &api.SaveNotebookReply{ImageId: imageReply.ImageId}, nil
}

func (s *developService) GetPodNameFromNoteBookTask(notebook *model.Notebook, taskName string) string {
	return fmt.Sprintf("%s-%s-0", notebook.NotebookJobId, taskName)
}

func (s *developService) GetPodNameFromNoteBookTaskByIndex(notebook *model.Notebook, taskIndex int32) string {
	return fmt.Sprintf("%s-task%d-0", notebook.NotebookJobId, taskIndex)
}

func (s *developService) getNotebookTaskContainer(ctx context.Context, notebook *model.Notebook, taskName string) (string, string, error) {
	pod, err := s.data.Cluster.GetPod(ctx, notebook.UserId, s.GetPodNameFromNoteBookTask(notebook, taskName))
	if err != nil {
		return "", "", err
	}
	if pod.Status.Phase != v1.PodRunning {
		return "", "", errors.Errorf(nil, errors.ErrorNotebookStatusForbidden)
	}
	for _, cs := range pod.Status.ContainerStatuses {
		if cs.Name == taskName {
			return pod.Spec.NodeName, strings.TrimPrefix(cs.ContainerID, "docker://"), nil
		}
	}
	return "", "", errors.Errorf(nil, errors.ErrorNotebookNoFoundRuntimeContainer)
}

func (s *developService) ListNotebookEventRecord(ctx context.Context, req *api.ListNotebookEventRecordRequest) (*api.ListNotebookEventRecordReply, error) {
	query := &model.NotebookEventRecordQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}

	recordsTbl, totalSize, err := s.data.DevelopDao.ListNotebookEventRecord(ctx, query)
	if err != nil {
		return nil, err
	}

	records := make([]*api.NotebookEventRecord, 0)
	for _, t := range recordsTbl {
		r := &api.NotebookEventRecord{}
		err := copier.Copy(r, t)
		if err != nil {
			return nil, err
		}
		r.Time = t.Time.Unix()
		records = append(records, r)
	}

	return &api.ListNotebookEventRecordReply{
		TotalSize: totalSize,
		Records:   records,
	}, nil
}

func (s *developService) getJobDetail(ctx context.Context, userID, jobID string) (*typeJob.JobStatusDetail, error) {

	nbJob, err := s.data.DevelopDao.GetNotebookJob(ctx, jobID)

	if err != nil {
		return nil, err
	}

	status := nbJob.Status

	if status == constant.FAILED || status == constant.STOPPED || status == constant.SUCCEEDED {
		if "" == nbJob.Detail || "{}" == nbJob.Detail {
			return defaultDetail(userID, nbJob), nil
		}
		detail := typeJob.JobStatusDetail{}
		json.Unmarshal([]byte(nbJob.Detail), &detail)
		return &detail, nil
	}

	var detail *typeJob.JobStatusDetail = nil
	job, err := s.data.Cluster.GetJob(ctx, userID, jobID)
	if nil == err && job != nil {
		detail = utils.Format(jobID, "notebookJob", job.Namespace, "", "", job)
	}
	if nil == detail {
		detail = defaultDetail(userID, nbJob)
		return detail, nil
	}
	return detail, nil
}

func defaultDetail(userID string, nbJob *model.NotebookJob) *typeJob.JobStatusDetail {

	status := constant.PREPARING

	if nbJob.Status == constant.STOPPED ||
		constant.SUSPENDED == nbJob.Status ||
		constant.FAILED == nbJob.Status {
		status = nbJob.Status
	}

	return &typeJob.JobStatusDetail{
		Version: "v1",
		Job: &typeJob.JobSummary{
			ID:     nbJob.Id,
			Name:   nbJob.Id,
			Type:   "notebookJob",
			UserID: userID,
			State:  status,
		},
	}
}

func (s *developService) GetNotebookMetric(ctx context.Context, req *api.GetNotebookMetricRequest) (*api.GetNotebookMetricReply, error) {
	notebook, err := s.data.DevelopDao.GetNotebook(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	resources, err := s.resourceService.ListResourceAll(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	resourceSpec, err := s.resourceSpecService.GetResourceSpecIgnore(ctx, &api.GetResourceSpecRequest{Id: notebook.ResourceSpecId})
	if err != nil {
		return nil, err
	}
	company, err := s.getCompany(ctx, resources, resourceSpec.ResourceSpec)
	if err != nil {
		return nil, err
	}
	podName := s.GetPodNameFromNoteBookTaskByIndex(notebook, req.TaskIndex)
	cpuUsage, err := s.data.Prometheus.QueryCpuUsage(ctx, podName, req.Start, int(req.Size), int(req.Step))
	if err != nil {
		return nil, err
	}
	memUsage, err := s.data.Prometheus.QueryMemUsage(ctx, podName, req.Start, int(req.Size), int(req.Step))
	if err != nil {
		return nil, err
	}
	gpuUtil, err := s.data.Prometheus.QueryGpuUtil(ctx, podName, req.Start, int(req.Size), int(req.Step))
	if err != nil {
		return nil, err
	}
	gpuMemUtil, err := s.data.Prometheus.QueryGpuMemUtil(ctx, podName, req.Start, int(req.Size), int(req.Step))
	if err != nil {
		return nil, err
	}
	accCardUtil, err := s.data.Prometheus.QueryAccCardUtil(ctx, podName, company, req.Start, int(req.Size), int(req.Step))
	if err != nil {
		return nil, err
	}
	accCardMemUtil, err := s.data.Prometheus.QueryAccCardMemUtil(ctx, podName, company, req.Start, int(req.Size), int(req.Step))
	if err != nil {
		return nil, err
	}
	networkReceiveBytes, err := s.data.Prometheus.QueryNetworkReceiveBytes(ctx, podName, req.Start, int(req.Size), int(req.Step))
	if err != nil {
		return nil, err
	}
	networkTransmitBytes, err := s.data.Prometheus.QueryNetworkTransmitBytes(ctx, podName, req.Start, int(req.Size), int(req.Step))
	if err != nil {
		return nil, err
	}
	fsUsageBytes, err := s.data.Prometheus.QueryFSUsageBytes(ctx, podName, req.Start, int(req.Size), int(req.Step))
	if err != nil {
		return nil, err
	}

	res := &api.GetNotebookMetricReply{
		MemUsage:             memUsage,
		GpuUtil:              gpuUtil,
		GpuMemUsage:          gpuMemUtil,
		AccCardUtil:          accCardUtil,
		AccCardMemUsage:      accCardMemUtil,
		NetworkReceiveBytes:  networkReceiveBytes,
		NetworkTransmitBytes: networkTransmitBytes,
		FsUsageBytes:         fsUsageBytes,
		Company:              company,
	}

	cpuAverageUsage, err := s.getCpuAverageUsage(ctx, resources, resourceSpec.ResourceSpec, cpuUsage)
	if err == nil { //忽略err
		res.CpuUsage = cpuAverageUsage
	} else {
		for range cpuUsage {
			res.CpuUsage = append(res.CpuUsage, -1)
		}
	}

	memUsagePercent, err := s.getMemUsagePercent(ctx, resources, resourceSpec.ResourceSpec, memUsage)
	if err == nil { //忽略err
		res.MemUsagePercent = memUsagePercent
	} else {
		for range memUsage {
			res.MemUsagePercent = append(res.MemUsagePercent, -1)
		}
	}

	return res, nil
}

func (s *developService) getCpuAverageUsage(
	ctx context.Context,
	resources *api.ResourceList,
	resourceSpec *api.ResourceSpec,
	cpuUsage []float64) ([]float64, error) {

	for _, r := range resources.Resources {
		for k, v := range resourceSpec.ResourceQuantity {
			if r.Name == k {
				if r.ResourceRef == cpuResource || r.Name == cpuResource {
					quantity, err := resource.ParseQuantity(v)
					if err != nil {
						return nil, err
					}
					res := make([]float64, 0)
					for _, v := range cpuUsage {
						if v == -1 {
							res = append(res, v)
						} else if quantity.Value() <= 0 {
							res = append(res, -1)
						} else {
							res = append(res, float64(v)/float64(quantity.Value()))
						}
					}
					return res, nil
				}
			}
		}
	}
	return nil, errors.Errorf(nil, errors.ErrorResourceNotExist)
}

func (s *developService) getMemUsagePercent(
	ctx context.Context,
	resources *api.ResourceList,
	resourceSpec *api.ResourceSpec,
	memUsage []float64) ([]float64, error) {

	for _, r := range resources.Resources {
		for k, v := range resourceSpec.ResourceQuantity {
			if r.Name == k {
				if r.ResourceRef == memResource || r.Name == memResource {
					quantity, err := resource.ParseQuantity(v)
					if err != nil {
						return nil, err
					}
					res := make([]float64, 0)
					for _, v := range memUsage {
						if v == -1 {
							res = append(res, v)
						} else if quantity.Value() <= 0 {
							res = append(res, -1)
						} else {
							res = append(res, float64(v)*100/float64(quantity.Value()))
						}
					}
					return res, nil
				}
			}
		}
	}
	return nil, errors.Errorf(nil, errors.ErrorResourceNotExist)
}

func (s *developService) getCompany(
	ctx context.Context,
	resources *api.ResourceList,
	resourceSpec *api.ResourceSpec) (string, error) {

	companyResource := []string{"nvidia", "huawei", "cambricon", "enflame", "iluvatar", "metax-tech"}
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
