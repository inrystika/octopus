package develop

import (
	"context"
	"fmt"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"server/base-server/internal/data/pipeline"
	"server/base-server/internal/service/algorithm"
	"server/common/constant"
	"server/common/errors"
	"server/common/utils"
	"time"

	vcBus "volcano.sh/volcano/pkg/apis/bus/v1alpha1"

	"server/common/log"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	vcBatch "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
)

type developService struct {
	api.UnimplementedDevelopServer
	conf                *conf.Bootstrap
	log                 *log.Helper
	data                *data.Data
	workspaceService    api.WorkspaceServer
	algorithmService    api.AlgorithmServer
	imageService        api.ImageServer
	datasetService      api.DatasetServiceServer
	resourceSpecService api.ResourceSpecServiceServer
	resourceService     api.ResourceServiceServer
	resourcePoolService api.ResourcePoolServiceServer
	billingService      api.BillingServiceServer
}

type DevelopService interface {
	api.DevelopServer
	common.PipelineCallback
}

const (
	k8sTaskNamePrefix = "task"
	servicePort       = 8888
	shmResource       = "shm"
)

func buildTaskName(idx int) string {
	return fmt.Sprintf("%s%d", k8sTaskNamePrefix, idx)
}

func buildServiceName(jobId string, idx int) string {
	return fmt.Sprintf("%s-%s", jobId, buildTaskName(idx))
}

func buildIngressName(jobId string, idx int) string {
	return fmt.Sprintf("%s-%s", jobId, buildTaskName(idx))
}

func buildNotebookUrl(jobId string, idx int) string {
	return fmt.Sprintf("/notebook_%s_%s", jobId, buildTaskName(idx))
}

func NewDevelopService(conf *conf.Bootstrap, logger log.Logger, data *data.Data,
	workspaceService api.WorkspaceServer, algorithmService api.AlgorithmServer,
	imageService api.ImageServer, datasetService api.DatasetServiceServer, resourceSpecService api.ResourceSpecServiceServer,
	resourceService api.ResourceServiceServer, resourcePoolService api.ResourcePoolServiceServer,
	billingService api.BillingServiceServer) (DevelopService, error) {
	log := log.NewHelper("DevelopService", logger)

	err := upsertFeature(data, conf.Service.BaseServerAddr)
	if err != nil {
		if conf.App.IsDev {
			log.Error(context.Background(), err)
		} else {
			return nil, err
		}
	}

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
	}

	s.startNotebookTask()
	return s, nil
}

func upsertFeature(data *data.Data, baseServerAddr string) error {
	err := data.Pipeline.UpsertFeature(context.Background(), &pipeline.UpsertFeatureParam{
		FeatureName: "notebook",
		Author:      "octopus",
		Description: "notebook",
		Enabled:     true,
		JobSelector: &pipeline.JobSelector{
			Conditions: []*pipeline.Condition{{
				Name:   "type",
				Key:    "jobKind",
				Expect: "notebook",
			}},
			Expression: "type",
		},
		Plugins: []*pipeline.Plugin{
			{
				Key:         "bindlifehook",
				PluginType:  "LifeHook",
				CallAddress: baseServerAddr + "/v1/developmanage/pipelinecallback",
				Description: "notebook lifehook",
				JobSelector: &pipeline.JobSelector{
					States: []string{"*"},
				},
			}},
	})

	if err != nil {
		return err
	}

	return nil
}

type closeFunc func(ctx context.Context) error

func (s *developService) checkPermAndAssign(ctx context.Context, nb *model.Notebook, nbJob *model.NotebookJob) (*startJobInfo, error) {
	queue := ""
	if nb.WorkspaceId == constant.SYSTEM_WORKSPACE_DEFAULT {
		pool, err := s.resourcePoolService.GetDefaultResourcePool(ctx, &emptypb.Empty{})
		if err != nil {
			return nil, err
		}
		queue = pool.ResourcePool.Name
	} else {
		workspace, err := s.workspaceService.GetWorkspace(ctx, &api.GetWorkspaceRequest{WorkspaceId: nb.WorkspaceId})
		if err != nil {
			return nil, err
		}

		queue = workspace.Workspace.ResourcePoolId
	}

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
				if r.Name == shmResource {
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

	return &startJobInfo{
		queue:         queue,
		imageAddr:     image.ImageFullAddr,
		algorithmPath: algorithmVersion.Algorithm.Path,
		datasetPath:   datasetPath,
		resources:     k8sResources,
		nodeSelectors: nodeSelectors,
		shm:           shm,
		command:       s.conf.Service.Develop.JpyCommand,
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
		err1 := s.data.Pipeline.StopJob(ctx, &pipeline.UpdateJobParam{JobID: nbJob.Id, Reason: "stop job because error"})
		return err1
	})

	err = s.createService(ctx, nb, nbJob)
	if err != nil {
		return nil, err
	}
	closes = append(closes, func(ctx context.Context) error {
		return s.data.Cluster.DeleteService(ctx, nb.UserId, nbJob.Id)
	})

	err = s.createIngress(ctx, nb, nbJob)
	if err != nil {
		return nil, err
	}
	closes = append(closes, func(ctx context.Context) error {
		return s.data.Cluster.DeleteIngress(ctx, nb.UserId, nbJob.Id)
	})

	return resFunc, nil
}

type startJobInfo struct {
	queue         string
	imageAddr     string
	algorithmPath string
	datasetPath   string
	command       string
	resources     map[v1.ResourceName]resource.Quantity
	nodeSelectors map[string]string
	shm           *resource.Quantity
}

func (s *developService) CreateNotebook(ctx context.Context, req *api.CreateNotebookRequest) (*api.CreateNotebookReply, error) {
	nbId := utils.GetUUIDStartWithAlphabetic() //k8s service首字母不允许数字 为方便 uuid处理一下
	err := s.data.DevelopDao.Transaction(ctx, func(ctx context.Context) error {
		_, size, err := s.data.DevelopDao.ListNotebook(ctx, &model.NotebookQuery{
			UserId:      req.UserId,
			WorkspaceId: req.WorkspaceId,
			Name:        req.Name,
		})
		if size > 0 {
			return errors.Errorf(nil, errors.ErrorNotebookRepeat)
		}

		jobId := utils.GetUUIDStartWithAlphabetic()

		nb := &model.Notebook{}
		err = copier.Copy(nb, req)
		if err != nil {
			return err
		}
		nb.Id = nbId
		nb.Status = pipeline.PREPARING
		nb.NotebookJobId = jobId
		nb.TaskNumber = int(req.TaskNumber)

		nbJob := &model.NotebookJob{
			Id:         jobId,
			NotebookId: nbId,
			Status:     pipeline.PREPARING,
		}

		startJobInfo, err := s.checkPermAndAssign(ctx, nb, nbJob)
		if err != nil {
			return err
		}

		//startJobInfo := &startJobInfo{ //test
		//	queue:         "common-pool",
		//	imageAddr:     "nginx:latest",
		//	algorithmPath: "default-workspace/ddbe4b31-cc13-416f-aa80-97495abb80c2/codes/id1",
		//	resources:     map[v1.ResourceName]resource.Quantity{"cpu": resource.MustParse("1")},
		//	nodeSelectors: map[string]string{"resourceType": "debug_cpu"},
		//}

		err = s.data.DevelopDao.CreateNotebook(ctx, nb)
		if err != nil {
			return err
		}

		err = s.data.DevelopDao.CreateNotebookJob(ctx, nbJob)
		if err != nil {
			return err
		}

		//数据库操作挪到前面，如果出错，直接不创建k8s vcjob，硬件资源有限的资源，出错需要及时释放掉
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
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &api.CreateNotebookReply{Id: nbId}, nil
}

func (s *developService) StartNotebook(ctx context.Context, req *api.StartNotebookRequest) (*api.StartNotebookReply, error) {
	err := s.data.DevelopDao.Transaction(ctx, func(ctx context.Context) error {
		nb, err := s.data.DevelopDao.GetNotebook(ctx, req.Id)
		if err != nil {
			return err
		}

		if !pipeline.IsCompletedState(nb.Status) {
			return errors.Errorf(nil, errors.ErrorNotebookStatusForbidden)
		}

		jobId := utils.GetUUIDStartWithAlphabetic()
		nbJob := &model.NotebookJob{
			Id:         jobId,
			NotebookId: nb.Id,
			Status:     pipeline.PREPARING,
		}

		startJobInfo, err := s.checkPermAndAssign(ctx, nb, nbJob)
		if err != nil {
			return err
		}

		err = s.data.DevelopDao.CreateNotebookJob(ctx, nbJob)
		if err != nil {
			return err
		}

		err = s.data.DevelopDao.UpdateNotebookSelective(ctx, &model.Notebook{
			Id:            nb.Id,
			NotebookJobId: jobId,
			Status:        pipeline.PREPARING,
		})
		if err != nil {
			return err
		}

		//数据库操作挪到前面，如果出错，直接不创建k8s vcjob，硬件资源有限的资源，出错需要及时释放掉
		closeFunc, err := s.startJob(ctx, nb, nbJob, startJobInfo)
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				err1 := closeFunc(ctx)
				if err1 != nil {
					s.log.Errorf(ctx, "err: %s", err1)
				}
			}
		}()

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &api.StartNotebookReply{Id: req.Id}, nil
}

func (s *developService) submitJob(ctx context.Context, nb *model.Notebook, nbJob *model.NotebookJob, startJobInfo *startJobInfo) error {
	param := &pipeline.SubmitJobParam{
		UserID:       nb.UserId,
		JobKind:      "notebook",
		JobName:      nbJob.Id,
		Header:       nil,
		JobNamespace: nb.UserId,
		Cluster:      "",
	}

	VolumeMounts := []v1.VolumeMount{
		{
			Name:      "data",
			MountPath: s.conf.Service.DockerCodePath,
			SubPath:   startJobInfo.algorithmPath,
			ReadOnly:  false,
		},
		{
			Name:      "localtime",
			MountPath: "/etc/localtime",
		},
	}

	if startJobInfo.datasetPath != "" {
		VolumeMounts = append(VolumeMounts, v1.VolumeMount{
			Name:      "data",
			MountPath: s.conf.Service.DockerDatasetPath,
			SubPath:   startJobInfo.datasetPath,
			ReadOnly:  true,
		})
	}

	volumes := []v1.Volume{
		{
			Name: "data",
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
		VolumeMounts = append(VolumeMounts, v1.VolumeMount{
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

	tasks := make([]vcBatch.TaskSpec, 0)
	for i := 0; i < nb.TaskNumber; i++ {
		taskName := buildTaskName(i)
		task := vcBatch.TaskSpec{
			Name:     taskName,
			Replicas: 1,
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					RestartPolicy: v1.RestartPolicyNever,
					Containers: []v1.Container{
						{
							Name:  taskName,
							Image: startJobInfo.imageAddr,
							Command: []string{"sh" ,"-c" ,startJobInfo.command},
							Resources: v1.ResourceRequirements{
								Requests: startJobInfo.resources,
								Limits:   startJobInfo.resources,
							},
							VolumeMounts: VolumeMounts,
							Env: []v1.EnvVar{{
								Name:  s.conf.Service.Develop.JpyBaseUrlEnv,
								Value: nb.Url,
							}},
						},
					},
					NodeSelector: startJobInfo.nodeSelectors,
					Volumes:      volumes,
				},
			},
		}
		tasks = append(tasks, task)
	}

	param.Job = &vcBatch.Job{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "batch.volcano.sh/v1alpha1",
			Kind:       "Job",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: nb.UserId,
			Name:      nbJob.Id,
		},
		Spec: vcBatch.JobSpec{
			MinAvailable:  1,
			Queue:         startJobInfo.queue,
			SchedulerName: "volcano",
			Plugins: map[string][]string{
				"env": {},
				"svc": {},
			},
			Policies: []vcBatch.LifecyclePolicy{
				{Event: vcBus.PodEvictedEvent, Action: vcBus.RestartJobAction},
				{Event: vcBus.PodFailedEvent, Action: vcBus.RestartJobAction},
			},
			Tasks: tasks,
		},
	}

	submitJobReply, err := s.data.Pipeline.SubmitJob(ctx, param)
	if err != nil {
		return err
	}
	if nbJob.Id != submitJobReply.JobId {
		return errors.Errorf(err, errors.ErrorPipelineDoRequest)
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

	if pipeline.IsCompletedState(nb.Status) {
		return nil, errors.Errorf(nil, errors.ErrorNotebookStatusForbidden)
	}

	err = s.data.Pipeline.StopJob(ctx, &pipeline.UpdateJobParam{JobID: nbJob.Id, Reason: "stop job"})
	if err != nil {
		return nil, err
	}

	err = s.deleteIngress(ctx, nb, nbJob)
	if err != nil {
		s.log.Errorw(ctx, "err", err)
	}

	err = s.deleteService(ctx, nb, nbJob)
	if err != nil {
		s.log.Errorw(ctx, "err", err)
	}

	err = s.data.DevelopDao.UpdateNotebookSelective(ctx, &model.Notebook{
		Id:     nb.Id,
		Status: pipeline.STOPPED,
	})
	if err != nil {
		return nil, err
	}

	now := time.Now()
	err = s.data.DevelopDao.UpdateNotebookJobSelective(ctx, &model.NotebookJob{
		Id:        nbJob.Id,
		Status:    pipeline.STOPPED,
		StoppedAt: &now,
	})
	if err != nil {
		return nil, err
	}

	return &api.StopNotebookReply{Id: req.Id}, nil
}

func (s *developService) createService(ctx context.Context, nb *model.Notebook, nbJob *model.NotebookJob) error {
	for i := 0; i < nb.TaskNumber; i++ {
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
				Ports: []v1.ServicePort{{
					Port:       servicePort,
					TargetPort: intstr.FromInt(servicePort),
				}},
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
		err := s.data.Cluster.CreateIngress(ctx, &v1beta1.Ingress{
			ObjectMeta: metav1.ObjectMeta{
				Name:      buildIngressName(nbJob.Id, i),
				Namespace: nb.UserId,
			},
			Spec: v1beta1.IngressSpec{
				Rules: []v1beta1.IngressRule{
					{
						IngressRuleValue: v1beta1.IngressRuleValue{
							HTTP: &v1beta1.HTTPIngressRuleValue{
								Paths: []v1beta1.HTTPIngressPath{
									{
										Path: buildNotebookUrl(nbJob.Id, i),
										Backend: v1beta1.IngressBackend{
											ServiceName: buildServiceName(nbJob.Id, i),
											ServicePort: intstr.FromInt(servicePort),
										},
									},
								},
							},
						},
					},
				},
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

func (s *developService) DeleteNotebook(ctx context.Context, req *api.DeleteNotebookRequest) (*api.DeleteNotebookReply, error) {
	nb, err := s.data.DevelopDao.GetNotebook(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if !pipeline.IsCompletedState(nb.Status) {
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
	priceMap := make(map[string]uint32)
	for _, j := range notebookJobs {
		priceMap[j.Id] = j.ResourceSpecPrice
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
			notebook.Tasks = append(notebook.Tasks, &api.Notebook_Task{Url: buildNotebookUrl(n.NotebookJobId, i)})
		}
		notebooks = append(notebooks, notebook)
	}
	return notebooks, nil
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
