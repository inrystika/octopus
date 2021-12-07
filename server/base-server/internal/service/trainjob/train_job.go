package trainjob

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
	"server/common/log"
	"server/common/utils"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	vcBatch "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
	vcBus "volcano.sh/volcano/pkg/apis/bus/v1alpha1"
)

const (
	k8sTaskNamePrefix   = "task"
	NoDistributedJobNum = 1
	shmResource         = "shm"
)

type trainJobService struct {
	api.UnimplementedTrainJobServiceServer
	conf                *conf.Bootstrap
	log                 *log.Helper
	data                *data.Data
	workspaceService    api.WorkspaceServiceServer
	algorithmService    api.AlgorithmServiceServer
	imageService        api.ImageServiceServer
	datasetService      api.DatasetServiceServer
	modelService        api.ModelServiceServer
	resourceSpecService api.ResourceSpecServiceServer
	resourceService     api.ResourceServiceServer
	resourcePoolService api.ResourcePoolServiceServer
	billingService      api.BillingServiceServer
}

type TrainJobService interface {
	api.TrainJobServiceServer
	common.PipelineCallback
}

func NewTrainJobService(conf *conf.Bootstrap, logger log.Logger, data *data.Data,
	workspaceService api.WorkspaceServiceServer, algorithmService api.AlgorithmServiceServer,
	imageService api.ImageServiceServer, datasetService api.DatasetServiceServer, modelService api.ModelServiceServer,
	resourceSpecService api.ResourceSpecServiceServer, resourceService api.ResourceServiceServer,
	resourcePoolService api.ResourcePoolServiceServer, billingService api.BillingServiceServer) (TrainJobService, error) {
	log := log.NewHelper("TrainJobService", logger)

	err := upsertFeature(data, conf.Service.BaseServerAddr)
	if err != nil {
		if conf.App.IsDev {
			log.Error(context.Background(), err) //todo 先只打印日志方便有些开发不启动taskset 后续修改为报错
		} else {
			return nil, err
		}
	}
	s := &trainJobService{
		conf:                conf,
		log:                 log,
		data:                data,
		workspaceService:    workspaceService,
		algorithmService:    algorithmService,
		imageService:        imageService,
		datasetService:      datasetService,
		modelService:        modelService,
		resourceSpecService: resourceSpecService,
		resourceService:     resourceService,
		resourcePoolService: resourcePoolService,
		billingService:      billingService,
	}

	s.trainJobBilling(context.Background())

	return s, nil
}

func upsertFeature(data *data.Data, baseServerAddr string) error {
	err := data.Pipeline.UpsertFeature(context.Background(), &pipeline.UpsertFeatureParam{
		FeatureName: "trainJob",
		Author:      "octopus",
		Description: "trainJob",
		Enabled:     true,
		JobSelector: &pipeline.JobSelector{
			Conditions: []*pipeline.Condition{{
				Name:   "type",
				Key:    "jobKind",
				Expect: "train_job",
			}},
			Expression: "type",
		},
		Plugins: []*pipeline.Plugin{
			{
				Key:         "bindlifehook",
				PluginType:  "LifeHook",
				CallAddress: baseServerAddr + "/v1/trainmanage/pipelinecallback",
				Description: "train_job life hook to update status and time",
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

func (s *trainJobService) TrainJob(ctx context.Context, req *api.TrainJobRequest) (*api.TrainJobReply, error) {
	trainJobId := utils.GetUUIDStartWithAlphabetic() //k8s service首字母不允许数字 为方便 uuid处理一下
	//check 任务是否重名，联合索引。同名且未软删除，则报错。
	_, err := s.data.TrainJobDao.GetTrainJobByName(ctx, req.Name, req.UserId, req.WorkspaceId)
	if err != nil {
		return nil, err
	}

	trainJob := &model.TrainJob{}
	err = copier.Copy(trainJob, req)
	if err != nil {
		return nil, err
	}
	trainJob.Id = trainJobId
	trainJob.Status = pipeline.PREPARING
	//各类参数校验
	startJobInfo, err := s.checkPermForJob(ctx, trainJob)
	if err != nil {
		return nil, err
	}
	//submit job
	closeFunc, err := s.submitJob(ctx, trainJob, startJobInfo)
	defer func() { //如果出错 重要的资源需要删除
		if err != nil {
			_ = closeFunc(ctx)
		}
	}()
	if err != nil {
		return nil, err
	}
	//create recorde
	err = s.data.TrainJobDao.CreateTrainJob(ctx, trainJob)
	if err != nil {
		return nil, err
	}

	return &api.TrainJobReply{JobId: trainJobId}, nil
}

func (s *trainJobService) buildCmd(config *model.Config) []string {
	cmd := fmt.Sprintf("cd %s;%s ", s.conf.Service.DockerCodePath, config.Command)
	if len(config.Parameters) == 0 {
		return []string{"sh", "-c", cmd}
	} else {
		for _, i := range config.Parameters {
			if i.Key == "" || i.Value == "" {
				continue
			} else {
				cmd += fmt.Sprintf("--%s=%s ", i.Key, i.Value)
			}
		}
	}
	return []string{"sh", "-c", cmd}
}

type closeFunc func(ctx context.Context) error

func (s *trainJobService) getModelSubPath(job *model.TrainJob) string {
	return fmt.Sprintf("%s/%s", common.GetMinioBucket(), common.GetMinioTrainJobObject(job.WorkspaceId, job.UserId, job.Id))
}

func (s *trainJobService) getImageAndCheckPerm(ctx context.Context, userId string, spaceId string, imageId string) (*api.FindImageReply, error) {
	reply, err := s.imageService.FindImage(ctx, &api.FindImageRequest{ImageId: imageId})
	if err != nil {
		return nil, err
	}

	if reply.Image == nil {
		return nil, errors.Errorf(nil, errors.ErrorImageNotExist)
	}

	if userId != reply.Image.UserId && reply.Image.IsPrefab == api.ImageIsPrefab_IMAGE_IS_PREFAB_NO {
		hasPerm := false
		for _, i := range reply.Accesses {
			if spaceId == i.SpaceId {
				hasPerm = true
				break
			}
		}

		if !hasPerm {
			return nil, errors.Errorf(err, errors.ErrorTrainImageForbidden)
		}
	}

	return reply, nil
}

func (s *trainJobService) getDatasetAndCheckPerm(ctx context.Context, userId string, spaceId string, datasetId string, datasetVersion string) (*api.GetDatasetVersionReply, error) {
	reply, err := s.datasetService.GetDatasetVersion(ctx, &api.GetDatasetVersionRequest{DatasetId: datasetId, Version: datasetVersion})
	if err != nil {
		return nil, err
	}
	if userId != reply.Dataset.UserId && reply.Dataset.SourceType == api.DatasetSourceType_DST_USER {
		hasPerm := false
		for _, i := range reply.VersionAccesses {
			if spaceId == i.SpaceId {
				hasPerm = true
				break
			}
		}

		if !hasPerm {
			return nil, errors.Errorf(err, errors.ErrorTrainDataSetForbidden)
		}
	}

	return reply, nil
}

func (s *trainJobService) getAlgorithmAndCheckPerm(ctx context.Context, userId string, spaceId string, algorithmId string, algorithmVersion string) (
	*api.QueryAlgorithmVersionReply, error) {
	reply, err := s.algorithmService.QueryAlgorithmVersion(ctx, &api.QueryAlgorithmVersionRequest{
		AlgorithmId: algorithmId,
		Version:     algorithmVersion,
	})
	if err != nil {
		return nil, err
	}
	if userId != reply.Algorithm.UserId && !reply.Algorithm.IsPrefab {
		hasPerm := false
		for _, i := range reply.VersionAccesses {
			if spaceId == i.SpaceId {
				hasPerm = true
				break
			}
		}

		if !hasPerm {
			return nil, errors.Errorf(err, errors.ErrorTrainAlgorithmForbidden)
		}
	}

	return reply, nil
}

type startJobInfoSpec struct {
	resources     map[v1.ResourceName]resource.Quantity
	nodeSelectors map[string]string
}

type startJobInfo struct {
	queue         string
	imageAddr     string
	algorithmPath string
	datasetPath   string
	specs         map[string]*startJobInfoSpec
}

func (s *trainJobService) checkPermForJob(ctx context.Context, job *model.TrainJob) (*startJobInfo, error) {
	//before commit job, check billing owner amount
	ownerId, ownerType := s.getOwner(job)
	owner, err := s.billingService.GetBillingOwner(ctx, &api.GetBillingOwnerRequest{
		OwnerId:   ownerId,
		OwnerType: ownerType,
	})
	if err != nil {
		return nil, err
	}

	if owner.BillingOwner.Amount <= 0 {
		return nil, errors.Errorf(nil, errors.ErrorTrainBalanceNotEnough)
	}

	queue := ""
	if job.WorkspaceId == constant.SYSTEM_WORKSPACE_DEFAULT {
		pool, err := s.resourcePoolService.GetDefaultResourcePool(ctx, &emptypb.Empty{})
		if err != nil {
			return nil, err
		}
		queue = pool.ResourcePool.Name
	} else {
		workspace, err := s.workspaceService.GetWorkspace(ctx, &api.GetWorkspaceRequest{WorkspaceId: job.WorkspaceId})
		if err != nil {
			return nil, err
		}

		queue = workspace.Workspace.ResourcePoolId
	}
	//image
	image, err := s.getImageAndCheckPerm(ctx, job.UserId, job.WorkspaceId, job.ImageId)
	if err != nil {
		return nil, err
	}

	if image.Image.ImageStatus != api.ImageStatus_IMAGE_STATUS_MADE {
		return nil, errors.Errorf(nil, errors.ErrorJobImageStatusForbidden)
	}
	job.ImageName = image.Image.ImageName
	job.ImageVersion = image.Image.ImageVersion
	//algorithm
	algorithmVersion, err := s.getAlgorithmAndCheckPerm(ctx, job.UserId, job.WorkspaceId, job.AlgorithmId, job.AlgorithmVersion)
	if err != nil {
		return nil, err
	}
	if algorithmVersion.Algorithm.FileStatus != int64(algorithm.FILESTATUS_FINISH) {
		return nil, errors.Errorf(err, errors.ErrorJobAlgorithmStatusForbidden)
	}
	job.AlgorithmName = algorithmVersion.Algorithm.AlgorithmName
	//dataSet
	dataSetVersion, err := s.getDatasetAndCheckPerm(ctx, job.UserId, job.WorkspaceId, job.DataSetId, job.DataSetVersion)
	if err != nil {
		return nil, err
	}
	if dataSetVersion.Version.Status != int32(api.DatasetVersionStatus_DVS_Unzipped) {
		return nil, errors.Errorf(err, errors.ErrorJobImageStatusForbidden)
	}
	job.DatasetName = dataSetVersion.Dataset.Name
	//resource spec info
	startJobSpecs := map[string]*startJobInfoSpec{}
	specs, err := s.resourceSpecService.ListResourceSpec(ctx, &api.ListResourceSpecRequest{})
	if err != nil {
		return nil, err
	}
	specMap := map[string]*api.ResourceSpec{}
	for _, i := range specs.ResourceSpecs {
		specMap[i.Id] = i
	}
	//resource
	resourcesReply, err := s.resourceService.ListResource(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	resourceMap := map[string]*api.Resource{}
	for _, i := range resourcesReply.Resources {
		resourceMap[i.Name] = i
	}

	//非分布式任务,config 个数不能超过1.
	if !job.IsDistributed && len(job.Config) > 1 {
		return nil, errors.Errorf(err, errors.ErrorInvalidRequestParameter)
	}

	//非分布式任务config中的副本总数、成功副本数、失败副本数，接口无需传参数; 若传，强制默认个数为1个。
	if !job.IsDistributed {
		for _, i := range job.Config {
			i.TaskNumber = NoDistributedJobNum
			i.MinFailedTaskCount = NoDistributedJobNum
			i.MinSucceededTaskCount = NoDistributedJobNum
		}
	}

	//job各config子任务对应的资源规格价格
	var resourceSpecPriceList []*model.ResourceSpecPrice
	for i, config := range job.Config {
		spec, err := s.resourceSpecService.GetResourceSpec(ctx, &api.GetResourceSpecRequest{Id: config.ResourceSpecId})
		if err != nil {
			return nil, err
		}
		resourceSpecPrice := &model.ResourceSpecPrice{}
		resourceSpecPrice.Task = i
		resourceSpecPrice.Price = spec.ResourceSpec.Price
		resourceSpecPriceList = append(resourceSpecPriceList, resourceSpecPrice)
	}
	job.ResSpecPrice = resourceSpecPriceList
	configNameMap := map[string]string{}
	for _, i := range job.Config {
		//检查子任务 task-number, minFailedTaskCount, minSucceededTaskCount 数量关系
		if i.MinFailedTaskCount > i.TaskNumber || i.MinSucceededTaskCount > i.MinSucceededTaskCount {
			return nil, errors.Errorf(err, errors.ErrorInvalidRequestParameter)
		}
		//检查子任务是否重名
		_, ok := configNameMap[i.Name]
		if ok {
			return nil, errors.Errorf(nil, errors.ErrorRepeatJobConfigName)
		} else {
			configNameMap[i.Name] = i.Name
		}
		//提交任务所需的资源规格映射表及节点标签映射表
		resources := map[v1.ResourceName]resource.Quantity{}
		nodeSelectors := map[string]string{}
		//通过资源规格映射表，获取规格名称及价格
		spec, ok := specMap[i.ResourceSpecId]
		if !ok {
			return nil, errors.Errorf(err, errors.ErrorInvalidRequestParameter)
		}
		i.ResourceSpecName = spec.Name
		i.ResourceSpecPrice = spec.Price
		var shm *resource.Quantity = nil
		//解析资源规格包中的各项资源（cpu,gpu,memory,shared-memory等）的值
		for k, v := range spec.ResourceQuantity {
			//获取资源规格包中的key及value
			r, ok := resourceMap[k]
			if !ok {
				return nil, errors.Errorf(err, errors.ErrorInvalidRequestParameter)
			}
			//解析资源规格value值
			quantity, err := resource.ParseQuantity(v)
			if err != nil {
				return nil, err
			}
			if r.Name == k {
				if r.Name == shmResource {
					shm = &quantity
					continue
				}
				if r.ResourceRef == "" {
					resources[v1.ResourceName(r.Name)] = quantity
				} else {
					resources[v1.ResourceName(r.ResourceRef)] = quantity
					nodeSelectors[r.BindingNodeLabelKey] = r.BindingNodeLabelValue
				}
			}
		}
		i.ShareMemory = shm
		startJobSpecs[i.ResourceSpecId] = &startJobInfoSpec{
			resources:     resources,
			nodeSelectors: nodeSelectors,
		}
	}

	return &startJobInfo{
		queue:         queue,
		imageAddr:     image.ImageFullAddr,
		algorithmPath: algorithmVersion.Algorithm.Path,
		datasetPath:   dataSetVersion.Version.Path,
		specs:         startJobSpecs,
	}, nil
}

//提交任务并将算法名称、数据集名称等字段赋值
func (s *trainJobService) submitJob(ctx context.Context, job *model.TrainJob, startJobInfo *startJobInfo) (closeFunc, error) {
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
			_ = resFunc(ctx)
		}
	}()

	param := &pipeline.SubmitJobParam{
		UserID:       job.UserId,
		JobKind:      "train_job",
		JobName:      job.Id,
		Header:       nil,
		JobNamespace: job.UserId,
		//JobNamespace: "default",
		Cluster: "",
	}

	minAvailable := 0
	tasks := make([]vcBatch.TaskSpec, 0)

	for idx, i := range job.Config {
		taskName := fmt.Sprintf("%s%d", k8sTaskNamePrefix, idx)
		minAvailable += i.TaskNumber
		//挂载卷
		volumeMounts := []v1.VolumeMount{
			{
				Name:      "data",
				MountPath: s.conf.Service.DockerCodePath,
				SubPath:   startJobInfo.algorithmPath,
				ReadOnly:  true,
			},
			{
				Name:      "data",
				MountPath: s.conf.Service.DockerDatasetPath,
				SubPath:   startJobInfo.datasetPath,
				ReadOnly:  true,
			},
			{
				Name:      "data",
				MountPath: s.conf.Service.DockerModelPath,
				SubPath:   s.getModelSubPath(job),
				ReadOnly:  false,
			},
			{
				Name:      "localtime",
				MountPath: "/etc/localtime",
			},
		}

		volumes := []v1.Volume{
			{
				Name: "data",
				VolumeSource: v1.VolumeSource{
					PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
						ClaimName: common.GetStoragePersistentVolumeChaim(job.UserId),
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

		//add shareMemory for each subTask
		if i.ShareMemory != nil {
			volumeMounts = append(volumeMounts, v1.VolumeMount{
				Name:      "cache-volume",
				MountPath: "/dev/shm",
			})
			volumes = append(volumes, v1.Volume{
				Name: "cache-volume",
				VolumeSource: v1.VolumeSource{
					EmptyDir: &v1.EmptyDirVolumeSource{
						Medium:    v1.StorageMediumMemory,
						SizeLimit: i.ShareMemory,
					},
				},
			})
		}
		//pod template
		task := vcBatch.TaskSpec{
			Name:     taskName,
			Replicas: int32(i.TaskNumber),
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{s.conf.Service.ResourceLabelKey: "train_job"},
				},
				Spec: v1.PodSpec{
					RestartPolicy: "Never",
					Containers: []v1.Container{
						{
							Name:  taskName,
							Image: startJobInfo.imageAddr,
							Resources: v1.ResourceRequirements{
								Requests: startJobInfo.specs[i.ResourceSpecId].resources,
								Limits:   startJobInfo.specs[i.ResourceSpecId].resources,
							},
							VolumeMounts: volumeMounts,
							Command:      s.buildCmd(i),
						},
					},
					NodeSelector: startJobInfo.specs[i.ResourceSpecId].nodeSelectors,
					Volumes:      volumes,
				},
			},
			CompletionPolicy: vcBatch.CompletionPolicy{
				MaxFailed:    int32(i.MinFailedTaskCount),
				MinSucceeded: int32(i.MinSucceededTaskCount),
			},
		}
		if i.IsMainRole {
			task.Policies = []vcBatch.LifecyclePolicy{
				{Event: vcBus.PodFailedEvent, Action: vcBus.AbortJobAction},
				{Event: vcBus.TaskCompletedEvent, Action: vcBus.CompleteJobAction},
			}
		}
		tasks = append(tasks, task)
	}

	param.Job = &vcBatch.Job{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "batch.volcano.sh/v1alpha1",
			Kind:       "Job",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: job.UserId,
			//Namespace: "default",
			Name: job.Id,
		},
		Spec: vcBatch.JobSpec{
			MinAvailable:  int32(minAvailable),
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
		Status: vcBatch.JobStatus{},
	}

	submitJobReply, err := s.data.Pipeline.SubmitJob(ctx, param)
	closes = append(closes, func(ctx context.Context) error {
		err1 := s.data.Pipeline.StopJob(ctx, &pipeline.UpdateJobParam{JobID: job.Id, Reason: "stop job because error"})
		return err1
	})
	if err != nil {
		return nil, err
	}
	if job.Id != submitJobReply.JobId {
		return nil, errors.Errorf(err, errors.ErrorPipelineDoRequest)
	}
	return resFunc, nil
}

func (s *trainJobService) StopJob(ctx context.Context, req *api.StopJobRequest) (*api.StopJobReply, error) {
	_, err := s.data.TrainJobDao.GetTrainJob(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	//pipeline删除任务成功后，任务从running转为terminate转态会触发callback机制,更新base-server中的任务状态信息。
	err = s.data.Pipeline.StopJob(ctx, &pipeline.UpdateJobParam{JobID: req.Id, Reason: req.Operation})
	if err != nil {
		return nil, err
	}

	now := time.Now()
	err = s.data.TrainJobDao.UpdateTrainJob(ctx, &model.TrainJob{
		Id:          req.Id,
		Operation:   req.Operation,
		Status:      pipeline.STOPPED,
		CompletedAt: &now,
	})
	if err != nil {
		return nil, err
	}

	return &api.StopJobReply{StoppedAt: time.Now().Unix()}, nil
}

func (s *trainJobService) DeleteJob(ctx context.Context, req *api.DeleteJobRequest) (*api.DeleteJobReply, error) {
	jobs, _, err := s.data.TrainJobDao.GetTrainJobList(ctx, &model.TrainJobListQuery{
		UserId: req.UserId,
		Ids:    req.JobIds,
	})
	if err != nil {
		return nil, err
	}

	for _, i := range jobs {
		//只有任务是终止状态，才可以删除
		if !pipeline.IsCompletedState(i.Status) {
			return nil, errors.Errorf(nil, errors.ErrorDeleteJobRequest)
		}

		//train_job软删除
		err = s.data.TrainJobDao.DeleteTrainJob(ctx, i.Id)
		if err != nil {
			return nil, err
		}
	}

	return &api.DeleteJobReply{DeletedAt: time.Now().Unix()}, nil
}

func (s *trainJobService) GetTrainJobInfo(ctx context.Context, req *api.TrainJobInfoRequest) (*api.TrainJobInfoReply, error) {
	// 网关层获取job基础信息
	trainJob, err := s.data.TrainJobDao.GetTrainJob(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	//pipeline获取job最新任务信息
	info, err := s.data.Pipeline.GetJobDetail(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	taskReplicaStatesMap := make(map[string]string)
	subTaskStateMap := make(map[int]string)

	replicaNum := 0
	for index, taskInfo := range info.Tasks {
		subTaskStateMap[index] = taskInfo.State
		for ri, replica := range taskInfo.Replicas {
			stateKey := "task" + strconv.Itoa(index) + "-replica-" + strconv.Itoa(ri)
			taskReplicaStatesMap[stateKey] = replica.State
			replicaNum++
		}
	}

	trainJobDetail, err := s.convertJobFromDb(trainJob)
	if err != nil {
		return nil, err
	}
	for index, config := range trainJobDetail.Config {
		replyStates := make([]*api.ReplicaState, 0)
		for ri := 0; ri < int(config.TaskNumber); ri++ {
			replicaState := new(api.ReplicaState)
			stateKey := "task" + strconv.Itoa(index) + "-replica-" + strconv.Itoa(ri)
			replicaState.Key = stateKey
			replicaState.State = taskReplicaStatesMap[stateKey]
			replicaState.Key = config.Name + "-replica-" + strconv.Itoa(ri)
			replyStates = append(replyStates, replicaState)
		}
		config.ReplicaStates = replyStates
		config.SubTaskState = subTaskStateMap[index]
	}

	return &api.TrainJobInfoReply{
		TrainJob: trainJobDetail,
	}, nil
}

func (s *trainJobService) convertJobFromDb(jobDb *model.TrainJob) (*api.TrainJob, error) {
	r := &api.TrainJob{}
	err := copier.CopyWithOption(r, jobDb, copier.Option{DeepCopy: true}) //model.Config实现了scan方法，这里需要DeepCopy设置为true，否则复制时转化为字符串
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	r.CreatedAt = jobDb.CreatedAt.Unix()
	r.UpdatedAt = jobDb.UpdatedAt.Unix()
	if jobDb.StartedAt != nil {
		r.StartedAt = jobDb.StartedAt.Unix()
	}
	if jobDb.CompletedAt != nil && jobDb.StartedAt != nil {
		//任务启动正常，终止正常：运行时间 = 终止时间-启动时间
		r.CompletedAt = jobDb.CompletedAt.Unix()
		r.RunSec = r.CompletedAt - r.StartedAt
	} else if jobDb.CompletedAt == nil && jobDb.StartedAt != nil {
		//任务启动正常，且尚未终止：运行时间 = 当前时间-启动时间
		r.RunSec = time.Now().Unix() - r.StartedAt
	} else {
		//其他情况，默认任务没有启动，不计算
		r.RunSec = 0
	}

	return r, nil
}

func (s *trainJobService) TrainJobList(ctx context.Context, req *api.TrainJobListRequest) (*api.TrainJobListReply, error) {
	query := &model.TrainJobListQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, err
	}

	trainJobsTbl, totalSize, err := s.data.TrainJobDao.GetTrainJobList(ctx, query)
	if err != nil {
		return nil, err
	}

	trainJobs := make([]*api.TrainJob, 0)
	for _, job := range trainJobsTbl {
		trainJob, err := s.convertJobFromDb(job)
		if err != nil {
			return nil, err
		}

		trainJobs = append(trainJobs, trainJob)
	}

	return &api.TrainJobListReply{
		TotalSize: totalSize,
		TrainJobs: trainJobs,
	}, nil
}

func (s *trainJobService) checkParamForTemplate(ctx context.Context, template *model.TrainJobTemplate) error {
	//镜像
	image, err := s.getImageAndCheckPerm(ctx, template.UserId, template.WorkspaceId, template.ImageId)
	if err != nil {
		return err
	}
	template.ImageVersion = image.Image.ImageVersion
	//算法
	_, err = s.getAlgorithmAndCheckPerm(ctx, template.UserId, template.WorkspaceId, template.AlgorithmId, template.AlgorithmVersion)
	if err != nil {
		return err
	}
	//数据集
	_, err = s.getDatasetAndCheckPerm(ctx, template.UserId, template.WorkspaceId, template.DataSetId, template.DataSetVersion)
	if err != nil {
		return err
	}
	//资源规格信息
	specs, err := s.resourceSpecService.ListResourceSpec(ctx, &api.ListResourceSpecRequest{})
	if err != nil {
		return err
	}
	specMap := map[string]*api.ResourceSpec{}
	for _, i := range specs.ResourceSpecs {
		specMap[i.Id] = i
	}

	configNameMap := map[string]string{}
	for _, i := range template.Config {
		//check task-number, minFailedTaskCount, minSucceededTaskCount
		if i.MinFailedTaskCount > i.TaskNumber || i.MinSucceededTaskCount > i.MinSucceededTaskCount {
			return errors.Errorf(err, errors.ErrorInvalidRequestParameter)
		}
		//check name repeat
		_, ok := configNameMap[i.Name]
		if ok {
			return errors.Errorf(nil, errors.ErrorRepeatJobConfigName)
		} else {
			configNameMap[i.Name] = i.Name
		}
		//get resource spec name and price
		spec, ok := specMap[i.ResourceSpecId]
		if !ok {
			return err
		}
		i.ResourceSpecName = spec.Name
		i.ResourceSpecPrice = spec.Price
	}
	return nil
}

func (s *trainJobService) CreateJobTemplate(ctx context.Context, req *api.TrainJobTemplateRequest) (*api.TrainJobTemplateReply, error) {
	jobTemplateId := utils.GetUUIDStartWithAlphabetic()
	//若模板已存在，前端提示模板已存在，不重复新建模板
	_, err := s.data.TrainJobDao.GetTrainJobTemplateByName(ctx, req.Name, req.UserId, req.WorkspaceId)
	if err != nil {
		return nil, err
	}

	trainJobTemplate := &model.TrainJobTemplate{}
	err = copier.Copy(trainJobTemplate, req)
	if err != nil {
		return nil, err
	}
	trainJobTemplate.Id = jobTemplateId

	err = s.checkParamForTemplate(ctx, trainJobTemplate)
	if err != nil {
		return nil, err
	}

	err = s.data.TrainJobDao.CreateTrainJobTemplate(ctx, trainJobTemplate)
	if err != nil {
		return nil, err
	}

	return &api.TrainJobTemplateReply{
		TemplateId: trainJobTemplate.Id,
	}, nil
}

func (s *trainJobService) CopyJobTemplate(ctx context.Context, req *api.CopyJobTemplateRequest) (*api.CopyJobTemplateReply, error) {
	tpl, err := s.data.TrainJobDao.GetTrainJobTemplate(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	newJobTemplateId := utils.GetUUIDStartWithAlphabetic()
	newTrainJobTemplate := &model.TrainJobTemplate{}
	err = copier.Copy(newTrainJobTemplate, tpl)
	if err != nil {
		return nil, err
	}
	newTrainJobTemplate.Id = newJobTemplateId
	newTrainJobTemplate.Name = fmt.Sprintf("copy-tpl-%v", time.Now().Unix())
	newTrainJobTemplate.DeletedAt = 0
	newTrainJobTemplate.CreatedAt = time.Time{}
	newTrainJobTemplate.UpdatedAt = time.Time{}

	//err = s.checkParamForTemplate(ctx, newTrainJobTemplate)
	//if err != nil {
	//	return nil, err
	//}

	err = s.data.TrainJobDao.CreateTrainJobTemplate(ctx, newTrainJobTemplate)
	if err != nil {
		return nil, err
	}
	return &api.CopyJobTemplateReply{
		TemplateId: newJobTemplateId,
	}, nil
}

func (s *trainJobService) convertTemplateFromDb(jobDb *model.TrainJobTemplate) (*api.TrainJobTemplate, error) {
	r := &api.TrainJobTemplate{}
	err := copier.CopyWithOption(r, jobDb, copier.Option{DeepCopy: true})
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorStructCopy)
	}
	r.CreatedAt = jobDb.CreatedAt.Unix()
	r.UpdatedAt = jobDb.UpdatedAt.Unix()

	return r, nil
}

func (s *trainJobService) GetJobTemplate(ctx context.Context, req *api.GetJobTemplateRequest) (*api.GetJobTemplateReply, error) {
	jobTemplateTbl, err := s.data.TrainJobDao.GetTrainJobTemplate(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	jobTemplate, err := s.convertTemplateFromDb(jobTemplateTbl)
	if err != nil {
		return nil, err
	}

	return &api.GetJobTemplateReply{
		JobTemplate: jobTemplate,
	}, nil
}

func (s *trainJobService) UpdateJobTemplate(ctx context.Context, req *api.TrainJobTemplateRequest) (*api.TrainJobTemplateReply, error) {
	_, err := s.data.TrainJobDao.GetTrainJobTemplate(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	trainJobTemplate := &model.TrainJobTemplate{}
	err = copier.Copy(trainJobTemplate, req)
	if err != nil {
		return nil, err
	}

	err = s.checkParamForTemplate(ctx, trainJobTemplate)
	if err != nil {
		return nil, err
	}

	err = s.data.TrainJobDao.UpdateTrainJobTemplate(ctx, trainJobTemplate)
	if err != nil {
		return nil, err
	}

	return &api.TrainJobTemplateReply{
		TemplateId: req.Id,
	}, nil
}

func (s *trainJobService) DeleteJobTemplate(ctx context.Context, req *api.DeleteJobTemplateRequest) (*api.DeleteJobTemplateReply, error) {
	err := s.data.TrainJobDao.DeleteTrainJobTemplate(req.UserId, req.TemplateIds)
	if err != nil {
		return nil, err
	}

	return &api.DeleteJobTemplateReply{DeletedAt: time.Now().Unix()}, nil
}

func (s *trainJobService) ListJobTemplate(ctx context.Context, req *api.TrainJobTemplateListRequest) (*api.TrainJobTemplateListReply, error) {
	query := &model.TrainJobTemPlateListQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, err
	}

	trainJobTemplatesTbl, totalSize, err := s.data.TrainJobDao.GetTrainJobTemplateList(ctx, query)
	if err != nil {
		return nil, err
	}

	trainJobTemplates := make([]*api.TrainJobTemplate, 0)
	for _, temp := range trainJobTemplatesTbl {
		trainJobTemplate, err := s.convertTemplateFromDb(temp)
		if err != nil {
			return nil, err
		}
		trainJobTemplates = append(trainJobTemplates, trainJobTemplate)
	}

	return &api.TrainJobTemplateListReply{
		TotalSize:    totalSize,
		JobTemplates: trainJobTemplates,
	}, nil
}

func (s *trainJobService) PipelineCallback(ctx context.Context, req *common.PipelineCallbackReq) string {
	s.log.Info(ctx, "pipeline callback for job :"+req.Id)
	trainJob, err := s.data.TrainJobDao.GetTrainJob(ctx, req.Id)
	if err != nil {
		return common.PipeLineCallbackRE
	}

	info, err := s.data.Pipeline.GetJobDetail(ctx, req.Id)
	if err != nil {
		return common.PipeLineCallbackRE
	}

	if pipeline.IsCompletedState(trainJob.Status) || strings.EqualFold(trainJob.Status, info.Job.State) {
		return common.PipeLineCallbackOK
	}

	update := &model.TrainJob{
		Id:     req.Id,
		Status: info.Job.State,
	}
	if strings.EqualFold(info.Job.State, pipeline.RUNNING) {
		update.StartedAt = &req.CurrentTime
	} else if strings.EqualFold(info.Job.State, pipeline.FAILED) ||
		strings.EqualFold(info.Job.State, pipeline.SUCCEEDED) {
		update.CompletedAt = &req.CurrentTime
	}

	err = s.data.TrainJobDao.UpdateTrainJob(ctx, update)
	if err != nil {
		return common.PipeLineCallbackRE
	}

	if strings.EqualFold(info.Job.State, pipeline.SUCCEEDED) ||
		strings.EqualFold(info.Job.State, pipeline.FAILED) {
		_, err = s.modelService.AddMyModel(ctx, &api.AddMyModelRequest{
			SpaceId:          trainJob.WorkspaceId,
			UserId:           trainJob.UserId,
			AlgorithmId:      trainJob.AlgorithmId,
			AlgorithmVersion: trainJob.AlgorithmVersion,
			FilePath:         fmt.Sprintf("%s/%s", s.conf.Data.Minio.Base.MountPath, s.getModelSubPath(trainJob)),
		})
		s.log.Error(ctx, err)
	}

	return common.PipeLineCallbackOK
}

func (s *trainJobService) GetJobEventList(ctx context.Context, req *api.JobEventListRequest) (*api.JobEventListReply, error) {

	query := &model.JobEventQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, err
	}

	events, totalSize, err := s.data.TrainJobDao.GetTrainJobEvents(query)
	if err != nil {
		return nil, err
	}

	jobEvents := make([]*api.JobEvent, 0)

	for _, value := range events {
		event := &api.JobEvent{}
		event.Timestamp = value.Timestamp
		event.Name = value.Name
		event.Reason = value.Reason
		event.Message = value.Message
		jobEvents = append(jobEvents, event)
	}

	return &api.JobEventListReply{
		TotalSize: totalSize,
		JobEvents: jobEvents,
	}, nil
}
