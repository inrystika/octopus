package platform

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	model "server/base-server/internal/data/dao/model/platform"
	"server/base-server/internal/data/pipeline"
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

type platformTrainJobService struct {
	api.UnimplementedPlatformTrainJobServiceServer
	conf                *conf.Bootstrap
	log                 *log.Helper
	data                *data.Data
	workspaceService    api.WorkspaceServer
	algorithmService    api.AlgorithmServer
	imageService        api.ImageServer
	datasetService      api.DatasetServiceServer
	modelService        api.ModelServer
	resourceSpecService api.ResourceSpecServiceServer
	resourceService     api.ResourceServiceServer
	resourcePoolService api.ResourcePoolServiceServer
	platformService     api.PlatformServiceServer
}

type PlatformTrainJobService interface {
	api.PlatformTrainJobServiceServer
	common.PipelineCallback
}

func NewPlatformTrainJobService(conf *conf.Bootstrap, logger log.Logger, data *data.Data,
	workspaceService api.WorkspaceServer, algorithmService api.AlgorithmServer,
	imageService api.ImageServer, datasetService api.DatasetServiceServer, modelService api.ModelServer,
	resourceSpecService api.ResourceSpecServiceServer, resourceService api.ResourceServiceServer,
	resourcePoolService api.ResourcePoolServiceServer, platformService api.PlatformServiceServer) (PlatformTrainJobService, error) {
	log := log.NewHelper("PlatformTrainJobService", logger)

	err := upsertFeature(data, conf.Service.BaseServerAddr)
	if err != nil {
		if conf.App.IsDev {
			log.Error(context.Background(), err) //todo 先只打印日志方便有些开发不启动taskset 后续修改为报错
		} else {
			return nil, err
		}
	}
	s := &platformTrainJobService{
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
		platformService:     platformService,
	}

	return s, nil
}

func upsertFeature(data *data.Data, baseServerAddr string) error {
	err := data.Pipeline.UpsertFeature(context.Background(), &pipeline.UpsertFeatureParam{
		FeatureName: "platformTrainJob",
		Author:      "octopus",
		Description: "platformTrainJob",
		Enabled:     true,
		JobSelector: &pipeline.JobSelector{
			Conditions: []*pipeline.Condition{{
				Name:   "type",
				Key:    "jobKind",
				Expect: "platform_train_job",
			}},
			Expression: "type",
		},
		Plugins: []*pipeline.Plugin{
			{
				Key:         "bindlifehook",
				PluginType:  "LifeHook",
				CallAddress: baseServerAddr + "/v1/trainmanage/pipelinecallback",
				Description: "platform_train_job life hook to update status and time",
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

func (s *platformTrainJobService) TrainJob(ctx context.Context, req *api.PlatformTrainJobRequest) (*api.PlatformTrainJobReply, error) {
	trainJobId := utils.GetUUIDStartWithAlphabetic() //k8s service首字母不允许数字 为方便 uuid处理一下
	//check 任务是否重名，联合索引。同名且未软删除，则报错。
	_, err := s.data.PlatformTrainJobDao.GetTrainJobByName(ctx, req.Name, req.PlatformId)
	if err != nil {
		return nil, err
	}

	trainJob := &model.PlatformTrainJob{}
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
	startJobInfo.queue = req.ResourcePool
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
	err = s.data.PlatformTrainJobDao.CreateTrainJob(ctx, trainJob)
	if err != nil {
		return nil, err
	}

	return &api.PlatformTrainJobReply{JobId: trainJobId}, nil
}

func (s *platformTrainJobService) buildCmd(task *model.Task) []string {
	cmd := fmt.Sprintf("cd %s;%s ", s.conf.Service.DockerCodePath, task.Command)
	if len(task.Parameters) == 0 {
		return []string{"sh", "-c", cmd}
	} else {
		for _, i := range task.Parameters {
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

//func (s *platformTrainJobService) getModelSubPath(job *model.TrainJob) string {
//	return fmt.Sprintf("%s/%s", common.GetMinioBucket(), common.GetMinioTrainJobObject(job.WorkspaceId, job.UserId, job.Id))
//}

type startJobInfoSpec struct {
	resources     map[v1.ResourceName]resource.Quantity
	nodeSelectors map[string]string
}

type startJobInfo struct {
	queue       string
	imageAddr   string
	datasetPath string
	specs       map[string]*startJobInfoSpec
}

func (s *platformTrainJobService) checkPermForJob(ctx context.Context, job *model.PlatformTrainJob) (*startJobInfo, error) {
	//image
	imageAddr := fmt.Sprintf("%s:%s", job.Image.Name, job.Image.Version)

	startJobSpecs := map[string]*startJobInfoSpec{}
	//resource
	resourcesReply, err := s.resourceService.ListResource(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	resourceMap := map[string]*api.Resource{}
	for _, i := range resourcesReply.Resources {
		resourceMap[i.Name] = i
	}

	if len(job.Tasks) < 1 {
		return nil, errors.Errorf(err, errors.ErrorInvalidRequestParameter)
	}

	//非分布式任务config中的副本总数、成功副本数、失败副本数，接口无需传参数; 若传，强制默认个数为1个。
	if len(job.Tasks) == 1 {
		for _, i := range job.Tasks {
			i.TaskNumber = NoDistributedJobNum
			i.MinFailedTaskCount = NoDistributedJobNum
			i.MinSucceededTaskCount = NoDistributedJobNum
		}
	}

	configNameMap := map[string]string{}
	for _, i := range job.Tasks {
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
		var shm *resource.Quantity = nil
		//解析资源规格包中的各项资源（cpu,gpu,memory,shared-memory等）的值
		for _, v := range i.Resources {
			name := v.Name
			r, ok := resourceMap[name]
			if !ok {
				return nil, errors.Errorf(err, errors.ErrorInvalidRequestParameter)
			}
			//解析资源规格value值
			quantity, err := resource.ParseQuantity(v.Size)
			if err != nil {
				return nil, err
			}
			if r.Name == name {
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
		startJobSpecs[i.Name] = &startJobInfoSpec{
			resources:     resources,
			nodeSelectors: nodeSelectors,
		}
	}

	return &startJobInfo{
		imageAddr:   imageAddr,
		datasetPath: job.Dataset.Addr,
		specs:       startJobSpecs,
	}, nil
}

//提交任务并将算法名称、数据集名称等字段赋值
func (s *platformTrainJobService) submitJob(ctx context.Context, job *model.PlatformTrainJob, startJobInfo *startJobInfo) (closeFunc, error) {
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
		UserID:       job.PlatformId,
		JobKind:      "platform_train_job",
		JobName:      job.Id,
		Header:       nil,
		JobNamespace: job.PlatformId,
		//JobNamespace: "default",
		Cluster: "",
	}

	err, pvcName := s.createDatasetStorageResource(ctx, job.PlatformId, job.Id, job.Dataset.StorageConfigName)
	if err != nil {
		return nil, err
	}

	minAvailable := 0
	tasks := make([]vcBatch.TaskSpec, 0)

	for idx, i := range job.Tasks {
		taskName := fmt.Sprintf("%s%d", k8sTaskNamePrefix, idx)
		minAvailable += i.TaskNumber
		//挂载卷
		volumeMounts := []v1.VolumeMount{
			{
				// 改为挂载dataset路径创建的pvc
				Name:      "data",
				MountPath: s.conf.Service.DockerDatasetPath,
				ReadOnly:  true,
			},
			//{
			//	Name:      "data",
			//	MountPath: s.conf.Service.DockerModelPath,
			//	SubPath:   s.getModelSubPath(job),
			//	ReadOnly:  false,
			//},
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
						ClaimName: pvcName,
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
					Labels: map[string]string{s.conf.Service.ResourceLabelKey: "platform_train_job"},
				},
				Spec: v1.PodSpec{
					RestartPolicy: "Never",
					Containers: []v1.Container{
						{
							Name:  taskName,
							Image: startJobInfo.imageAddr,
							Resources: v1.ResourceRequirements{
								Requests: startJobInfo.specs[i.Name].resources,
								Limits:   startJobInfo.specs[i.Name].resources,
							},
							VolumeMounts: volumeMounts,
							Command:      s.buildCmd(i),
						},
					},
					NodeSelector: startJobInfo.specs[i.Name].nodeSelectors,
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
			Namespace: job.PlatformId,
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

func (s *platformTrainJobService) StopJob(ctx context.Context, req *api.PlatformStopJobRequest) (*api.PlatformStopJobReply, error) {
	_, err := s.data.PlatformTrainJobDao.GetTrainJob(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	//pipeline删除任务成功后，任务从running转为terminate转态会触发callback机制,更新base-server中的任务状态信息。
	err = s.data.Pipeline.StopJob(ctx, &pipeline.UpdateJobParam{JobID: req.Id, Reason: req.Operation})
	if err != nil {
		return nil, err
	}

	now := time.Now()
	err = s.data.PlatformTrainJobDao.UpdateTrainJob(ctx, &model.PlatformTrainJob{
		Id:          req.Id,
		Operation:   req.Operation,
		Status:      pipeline.STOPPED,
		CompletedAt: &now,
	})
	if err != nil {
		return nil, err
	}

	return &api.PlatformStopJobReply{StoppedAt: time.Now().Unix()}, nil
}

func (s *platformTrainJobService) GetTrainJobInfo(ctx context.Context, req *api.PlatformTrainJobInfoRequest) (*api.PlatformTrainJobInfoReply, error) {
	// 网关层获取job基础信息
	trainJob, err := s.data.PlatformTrainJobDao.GetTrainJob(ctx, req.Id)
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

	return &api.PlatformTrainJobInfoReply{
		TrainJob: trainJobDetail,
	}, nil
}

func (s *platformTrainJobService) convertJobFromDb(jobDb *model.PlatformTrainJob) (*api.PlatformTrainJob, error) {
	r := &api.PlatformTrainJob{}
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

func (s *platformTrainJobService) TrainJobList(ctx context.Context, req *api.PlatformTrainJobListRequest) (*api.PlatformTrainJobListReply, error) {
	query := &model.PlatformTrainJobListQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, err
	}

	trainJobsTbl, totalSize, err := s.data.PlatformTrainJobDao.GetTrainJobList(ctx, query)
	if err != nil {
		return nil, err
	}

	trainJobs := make([]*api.PlatformTrainJob, 0)
	for _, job := range trainJobsTbl {
		trainJob, err := s.convertJobFromDb(job)
		if err != nil {
			return nil, err
		}

		trainJobs = append(trainJobs, trainJob)
	}

	return &api.PlatformTrainJobListReply{
		TotalSize: totalSize,
		TrainJobs: trainJobs,
	}, nil
}

func (s *platformTrainJobService) PipelineCallback(ctx context.Context, req *common.PipelineCallbackReq) string {
	s.log.Info(ctx, "pipeline callback for platformjob :"+req.Id)
	trainJob, err := s.data.PlatformTrainJobDao.GetTrainJob(ctx, req.Id)
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

	update := &model.PlatformTrainJob{
		Id:     req.Id,
		Status: info.Job.State,
	}
	if strings.EqualFold(info.Job.State, pipeline.RUNNING) {
		update.StartedAt = &info.Job.StartAt.Time
	} else if strings.EqualFold(info.Job.State, pipeline.FAILED) ||
		strings.EqualFold(info.Job.State, pipeline.SUCCEEDED) {
		update.CompletedAt = &info.Job.FinishedAt.Time
	}

	err = s.data.PlatformTrainJobDao.UpdateTrainJob(ctx, update)
	if err != nil {
		return common.PipeLineCallbackRE
	}

	/*if strings.EqualFold(info.Job.State, pipeline.SUCCEEDED) {
		_, err = s.modelService.AddMyModel(ctx, &api.AddMyModelRequest{
			SpaceId:          trainJob.WorkspaceId,
			UserId:           trainJob.UserId,
			AlgorithmId:      trainJob.AlgorithmId,
			AlgorithmVersion: trainJob.AlgorithmVersion,
			FilePath:         fmt.Sprintf("%s/%s", s.conf.Data.Minio.Base.MountPath, s.getModelSubPath(trainJob)),
		})
		s.log.Error(ctx, err)
	}*/

	return common.PipeLineCallbackOK
}

func (s *platformTrainJobService) PlatformResources(ctx context.Context, req *api.PlatformResourcesRequest) (*api.PlatformResourcesReply, error) {

	resNodeList := &api.PlatformResourcesReply{
		Resources: []*api.PlatformNode{},
	}

	resNodeAllcatedResourceMap := make(map[string]map[string]*resource.Quantity)

	allNodeMap, err := s.getNodesByResourcePool(ctx, req.ResourcePool)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorListNode)
	}

	tasks, err := s.data.Cluster.GetRunningTasks(ctx)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorListNode)
	}

	for nodename, node := range allNodeMap {
		resNodeAllcatedResourceMap[nodename] = make(map[string]*resource.Quantity)
		resNode := &api.PlatformNode{
			NodeName:  nodename,
			Status:    "NotReady",
			Capacity:  make(map[string]string),
			Allocated: make(map[string]string),
		}

		for _, addr := range node.Status.Addresses {
			if addr.Type == "InternalIP" {
				resNode.Ip = addr.Address
				break
			}
		}

		for _, cond := range node.Status.Conditions {
			if cond.Type == "Ready" && cond.Status == "True" {
				resNode.Status = "Ready"
				break
			}
		}

		for resname, quantity := range node.Status.Capacity {
			quantityStr := quantity.String()
			if quantityStr != "0" &&
				!strings.Contains(s.conf.Service.Resource.IgnoreSystemResources, resname.String()) {
				resNode.Capacity[resname.String()] = quantityStr

			}
		}

		resNodeList.Resources = append(resNodeList.Resources, resNode)
	}

	for _, task := range tasks.Items {
		taskNodeName := task.Spec.NodeName
		oneNodeAllcatedResourceMap := resNodeAllcatedResourceMap[taskNodeName]

		for _, container := range task.Spec.Containers {
			for resname, quantity := range container.Resources.Requests {
				if _, ok := oneNodeAllcatedResourceMap[resname.String()]; !ok {
					newQ := quantity.DeepCopy()
					oneNodeAllcatedResourceMap[resname.String()] = &newQ
				} else {
					oneNodeAllcatedResourceMap[resname.String()].Add(quantity)
				}
			}
		}
	}

	for _, node := range resNodeList.Resources {
		nodeAllcatedResourceMap := resNodeAllcatedResourceMap[node.NodeName]
		for resname, quantity := range nodeAllcatedResourceMap {
			if !strings.Contains(s.conf.Service.Resource.IgnoreSystemResources, resname) {
				node.Allocated[resname] = quantity.String()
			}
		}

		for resname := range node.Capacity {
			if _, ok := node.Allocated[resname]; !ok {
				node.Allocated[resname] = "0"
			}
		}
	}

	return resNodeList, nil
}

func (s *platformTrainJobService) getNodesByResourcePool(ctx context.Context, resourcePool string) (map[string]v1.Node, error) {

	nodeMap := make(map[string]v1.Node)
	rPoolBindingNodeLabelKeyFormat := s.conf.Service.Resource.PoolBindingNodeLabelKeyFormat
	rPoolBindingNodeLabelKey := fmt.Sprintf(rPoolBindingNodeLabelKeyFormat, resourcePool)

	nodeListBytes, err := s.data.Cluster.ListNode(ctx, rPoolBindingNodeLabelKey)
	if err != nil {
		return nodeMap, errors.Errorf(err, errors.ErrorListResourcePool)
	}

	nodeList := &v1.NodeList{}
	err = json.Unmarshal(nodeListBytes, nodeList)

	if err != nil {
		return nodeMap, errors.Errorf(err, errors.ErrorListResourcePool)
	}

	for _, node := range nodeList.Items {
		nodeMap[node.Name] = node
	}

	return nodeMap, err

}

func (s *platformTrainJobService) TrainJobStastics(ctx context.Context, req *api.TrainJobStasticsRequest) (*api.TrainJobStasticsReply, error) {
	reply, err := s.data.PlatformTrainJobDao.TrainJobStastics(ctx, &model.TrainJobStastics{
		CreatedAtGte: req.CreatedAtGte,
		CreatedAtLt:  req.CreatedAtLt,
	})
	if err != nil {
		return nil, err
	}
	return &api.TrainJobStasticsReply{
		TotalSize:     reply.TotalSize,
		SucceededSize: reply.SucceededSize,
		FailedSize:    reply.FailedSize,
		StoppedSize:   reply.StoppedSize,
		RunningSize:   reply.RunningSize,
		WaitingSize:   reply.WaitingSize,
	}, nil
}

func (s *platformTrainJobService) createDatasetStorageResource(ctx context.Context, platformId, jobId, storageConfigName string) (error, string) {
	pvName := fmt.Sprintf("octopus-pv-juicefs-%s", jobId)
	pvcName := fmt.Sprintf("octopus-pvc-juicefs-%s", jobId)
	capacity := "10Pi"
	volumeMode := v1.PersistentVolumeFilesystem
	pvLableKey := "octopus-pv-label-key"
	pvLableValue := fmt.Sprintf("octopus-pv-label-%s", jobId)

	reply, err := s.platformService.GetPlatformStorageConfigByName(ctx, &api.GetPlatformStorageConfigByNameRequest{
		PlatformId: platformId,
		Name:       storageConfigName,
	})
	if err != nil {
		return err, ""
	}

	juiceName := reply.PlatformStorageConfig.Options.Juicefs.Name
	metaUrl := reply.PlatformStorageConfig.Options.Juicefs.MetaUrl

	pv := &v1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:   pvName,
			Labels: map[string]string{pvLableKey: pvLableValue},
		},
		Spec: v1.PersistentVolumeSpec{
			VolumeMode:                    &volumeMode,
			PersistentVolumeReclaimPolicy: v1.PersistentVolumeReclaimRetain,
			AccessModes:                   []v1.PersistentVolumeAccessMode{v1.ReadWriteMany},
			Capacity:                      map[v1.ResourceName]resource.Quantity{v1.ResourceStorage: resource.MustParse(capacity)},
			PersistentVolumeSource: v1.PersistentVolumeSource{
				CSI: &v1.CSIPersistentVolumeSource{
					VolumeHandle: fmt.Sprintf("dataset-%s", jobId),
					Driver:       "csi.juicefs.com",
					FSType:       "juicefs",
					NodePublishSecretRef: &v1.SecretReference{
						Name:      "juicefs-secret",
						Namespace: "default",
					},
				},
			},
		},
	}
	pvc := &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:   pvcName,
			Labels: map[string]string{pvLableKey: pvLableValue},
		},
		Spec: v1.PersistentVolumeClaimSpec{
			VolumeMode:  &volumeMode,
			AccessModes: []v1.PersistentVolumeAccessMode{v1.ReadWriteMany},
			Resources: v1.ResourceRequirements{
				Requests: map[v1.ResourceName]resource.Quantity{v1.ResourceStorage: resource.MustParse(capacity)},
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{pvLableKey: pvLableValue},
			},
		},
	}
	sct := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: "juicefs-secret",
		},
		Type: v1.SecretTypeOpaque,
		Data: map[string][]byte{
			"Name":    []byte(base64.StdEncoding.EncodeToString([]byte(juiceName))),
			"Metaurl": []byte(base64.StdEncoding.EncodeToString([]byte(metaUrl))),
		},
	}

	_, err = s.data.Cluster.CreatePersistentVolume(ctx, pv)
	if err != nil {
		return err, ""
	}

	_, err = s.data.Cluster.CreatePersistentVolumeClaim(ctx, pvc)
	if err != nil {
		return err, ""
	}

	_, err = s.data.Cluster.CreateSecret(ctx, sct)
	if err != nil {
		return err, ""
	}

	return nil, pvcName
}
