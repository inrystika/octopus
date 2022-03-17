package modeldeploy

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"server/common/constant"
	"server/common/errors"
	"server/common/log"
	"server/common/utils"
	"strings"
	"time"
	"unsafe"

	"github.com/jinzhu/copier"
	SeldonVersion "github.com/seldonio/seldon-core/operator/apis/machinelearning.seldon.io/v1"
	seldonv1 "github.com/seldonio/seldon-core/operator/apis/machinelearning.seldon.io/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

const (
	deployJobNum             = 1
	TensorFlowFrame          = "tensorflow"
	PytorchFrame             = "pytorch"
	ModelVolumePath          = "model_volume_path"
	SeldonDockerWorkDir      = "/app/models"
	PredictorSpecName        = "default"
	modelDeployContainerName = "model"
	SeldonInUrl              = "/seldon/"
	ServiceUrlSuffix         = "/api/v1.0/predictions"
	ModelUserId              = "model_user_Id"
	ModelId                  = "model_Id"
	ModelVersion             = "model_version"
	TFSignatureName          = "signature_name"
	TFModelName              = "model_name"
	PytorchServerVersion     = "192.168.202.110:5000/train/pytorchserver:2.0.0"
	STATE_PREPARING          = "Preparing"
	STATE_AVAILABLE          = "Available"
	STATE_CREATING           = "Creating"
	STATE_FAILED             = "Failed"
	STATE_STOPPED            = "Stopped"
)

type modelDeployService struct {
	api.UnimplementedModelDeployServiceServer
	conf                *conf.Bootstrap
	log                 *log.Helper
	data                *data.Data
	modelService        api.ModelServiceServer
	workspaceService    api.WorkspaceServiceServer
	resourceSpecService api.ResourceSpecServiceServer
	resourceService     api.ResourceServiceServer
	resourcePoolService api.ResourcePoolServiceServer
	billingService      api.BillingServiceServer
}

type ModelDeployService interface {
	api.ModelDeployServiceServer
}

type startJobInfoSpec struct {
	resources     map[v1.ResourceName]resource.Quantity
	nodeSelectors map[string]string
}

type startJobInfo struct {
	queue      string
	modelFrame string
	modelPath  string
	modelUrl   string
	specs      map[string]*startJobInfoSpec
}

func NewModelDeployService(conf *conf.Bootstrap, logger log.Logger, data *data.Data,
	workspaceService api.WorkspaceServiceServer, modelService api.ModelServiceServer,
	resourceSpecService api.ResourceSpecServiceServer, resourceService api.ResourceServiceServer,
	resourcePoolService api.ResourcePoolServiceServer, billingService api.BillingServiceServer) (ModelDeployService, error) {
	log := log.NewHelper("ModelDeployService", logger)

	s := &modelDeployService{
		conf:                conf,
		log:                 log,
		data:                data,
		workspaceService:    workspaceService,
		modelService:        modelService,
		resourceSpecService: resourceSpecService,
		resourceService:     resourceService,
		resourcePoolService: resourcePoolService,
		billingService:      billingService,
	}

	s.modelServiceBilling(context.Background())
	ctx := context.Background()
	s.data.Cluster.RegisterDeploymentInformerCallback(
		func(obj interface{}) {},
		func(old, obj interface{}) {
			s.updateDeployStatus(ctx, obj)
		},
		func(obj interface{}) {
			s.updateDeployStatus(ctx, obj)
		},
	)

	return s, nil
}

func (s *modelDeployService) updateDeployStatus(ctx context.Context, obj interface{}) {
	objSeldon := &SeldonVersion.SeldonDeployment{}
	err := runtime.DefaultUnstructuredConverter.
		FromUnstructured(obj.(*unstructured.Unstructured).UnstructuredContent(), objSeldon)
	if err != nil {
		s.log.Errorf(ctx, "could not convert obj to SeldonDeployment")
		s.log.Errorf(ctx, "err: %s", err)
		return
	}
	seldonNameArray := strings.Split(objSeldon.Name, "-sdep")
	seldonServiceId := seldonNameArray[0]
	deployService, err := s.data.ModelDeployDao.GetModelDeployService(ctx, seldonServiceId)
	if err != nil {
		return
	}
	if deployService.Status == STATE_STOPPED {
		return
	}
	newState := string(objSeldon.Status.State)
	if newState == deployService.Status {
		return
	}
	update := &model.ModelDeploy{
		Id:     seldonServiceId,
		Status: newState,
	}
	now := time.Now()
	if newState == STATE_AVAILABLE {
		update.StartedAt = &now
	} else if newState == STATE_FAILED {
		update.CompletedAt = &now
	}
	s.data.ModelDeployDao.UpdateModelDeployService(ctx, update)
}

// 部署模型服务
func (s *modelDeployService) DeployModel(ctx context.Context, req *api.DepRequest) (*api.DepReply, error) {
	modelServiceId := utils.GetUUIDStartWithAlphabetic() //k8s service首字母不允许数字 为方便 uuid处理一下
	//check 任务是否重名，联合索引。同名且未软删除，则报错。
	_, err := s.data.ModelDeployDao.GetModelDeployServiceByName(ctx, req.Name, req.UserId, req.WorkspaceId)
	if err != nil {
		return nil, err
	}

	deployJob := &model.ModelDeploy{}
	err = copier.Copy(deployJob, req)
	if err != nil {
		return nil, err
	}
	deployJob.Id = modelServiceId
	deployJob.Status = STATE_PREPARING
	//模型部署参数校验
	startJobInfo, err := s.checkParam(ctx, deployJob)
	if err != nil {
		return nil, err
	}
	// 获取模型路径
	reply, err := s.modelService.DownloadModelVersion(ctx, &api.DownloadModelVersionRequest{
		ModelId: req.ModelId,
		Version: req.ModelVersion,
		Domain:  req.Domain,
	})
	if err != nil {
		return nil, err
	}
	startJobInfo.modelUrl = reply.DownloadUrl
	startJobInfo.modelFrame = req.ModelFrame
	//submit deploy job
	closeFunc, modelInferServiceUrl, err := s.submitDeployJob(ctx, deployJob, startJobInfo)
	defer func() { //如果出错 重要的资源需要删除
		if err != nil {
			_ = closeFunc(ctx)
		}
	}()
	if err != nil {
		return nil, err
	}
	deployJob.ServiceUrl = modelInferServiceUrl
	//create recorde
	err = s.data.ModelDeployDao.CreateModelDeployService(ctx, deployJob)
	if err != nil {
		return nil, err
	}

	return &api.DepReply{
		ServiceId:  modelServiceId,
		ServiceUrl: modelInferServiceUrl,
		Message:    "deploy model infer service successfully!",
	}, nil
}

type closeFunc func(ctx context.Context) error

func (s *modelDeployService) submitDeployJob(ctx context.Context, modelDeploy *model.ModelDeploy, startJobInfo *startJobInfo) (closeFunc, string, error) {
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

	modelDeployContainerName := modelDeployContainerName
	//容器中的模型挂载路径
	mountPath := fmt.Sprintf("%s/%s/%s/%s", SeldonDockerWorkDir, modelDeploy.UserId, modelDeploy.ModelId, modelDeploy.ModelVersion)
	seldonTFDefaultMountPath := "/mnt/models"
	//挂载卷
	volumeMounts := make([]v1.VolumeMount, 0)
	if startJobInfo.modelFrame == PytorchFrame {
		volumeMounts = []v1.VolumeMount{
			{
				Name:      "modelfilepath",
				MountPath: mountPath,
				SubPath:   startJobInfo.modelPath,
				ReadOnly:  false,
			},
			{
				Name:      "localtime",
				MountPath: "/etc/localtime",
			},
		}
	} else {
		volumeMounts = []v1.VolumeMount{
			{
				Name:      "modelfilepath",
				MountPath: seldonTFDefaultMountPath,
				SubPath:   startJobInfo.modelPath,
				ReadOnly:  false,
			},
			{
				Name:      "localtime",
				MountPath: "/etc/localtime",
			},
		}
	}

	volumes := []v1.Volume{
		{
			Name: "modelfilepath",
			VolumeSource: v1.VolumeSource{
				PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
					ClaimName: common.GetStoragePersistentVolumeChaim(modelDeploy.UserId),
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
	//自定义pytorch框架的服务器
	parameters := []seldonv1.Parameter{
		{
			Name:  ModelVolumePath,
			Type:  "STRING",
			Value: SeldonDockerWorkDir,
		},
		{
			Name:  ModelUserId,
			Type:  "STRING",
			Value: modelDeploy.UserId,
		},
		{
			Name:  ModelId,
			Type:  "STRING",
			Value: modelDeploy.ModelId,
		},
		{
			Name:  ModelVersion,
			Type:  "STRING",
			Value: modelDeploy.ModelVersion,
		},
	}
	// tf 框架的服务器
	tfGraphParameters := []seldonv1.Parameter{
		{
			Name:  TFSignatureName,
			Type:  "STRING",
			Value: "predict_images",
		},
		{
			Name:  TFModelName,
			Type:  "STRING",
			Value: fmt.Sprintf("%s", modelDeploy.Name),
		},
	}

	seldonPodSpecs := make([]*seldonv1.SeldonPodSpec, 0)
	if startJobInfo.modelFrame == PytorchFrame {
		seldonPodSpec := &seldonv1.SeldonPodSpec{
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:         modelDeployContainerName,
						VolumeMounts: volumeMounts,
						Resources: v1.ResourceRequirements{
							Requests: startJobInfo.specs[modelDeploy.ResourceSpecId].resources,
							Limits:   startJobInfo.specs[modelDeploy.ResourceSpecId].resources,
						},
						Image: PytorchServerVersion,
					},
				},
				Volumes: volumes,
				// 使用火山调度器
				SchedulerName: "volcano",
			},
		}
		seldonPodSpecs = append(seldonPodSpecs, seldonPodSpec)
	} else if startJobInfo.modelFrame == TensorFlowFrame {
		seldonPodSpec := &seldonv1.SeldonPodSpec{
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:         modelDeployContainerName,
						VolumeMounts: volumeMounts,
						Resources: v1.ResourceRequirements{
							Requests: startJobInfo.specs[modelDeploy.ResourceSpecId].resources,
							Limits:   startJobInfo.specs[modelDeploy.ResourceSpecId].resources,
						},
						// 此处本可以不指定，不过外网中pull image的时间较长。将其放在章鱼镜像仓。
						//Image: "seldonio/tfserving-proxy:1.12.0",
					},
				},
				Volumes: volumes,
				// 使用火山调度器
				SchedulerName: "volcano",
			},
		}
		seldonPodSpecs = append(seldonPodSpecs, seldonPodSpec)
	}

	var modelServer seldonv1.PredictiveUnitImplementation
	var graphType seldonv1.PredictiveUnitType
	var graph seldonv1.PredictiveUnit
	graphType = "MODEL"
	if modelDeploy.ModelFrame == TensorFlowFrame {
		modelServer = "TENSORFLOW_SERVER"
		graph = seldonv1.PredictiveUnit{
			Name:           modelDeployContainerName,
			Children:       nil,
			Type:           &graphType,
			Implementation: &modelServer,
			//tf服务器直接从此处下载文件
			//ModelURI:   "gs://seldon-models/tfserving/mnist-model",
			ModelURI:   "gs:seldon-models/sklearn/iris",
			Parameters: tfGraphParameters,
		}
	} else if startJobInfo.modelFrame == PytorchFrame {
		modelServer = "PYTORCH_SERVER"
		graph = seldonv1.PredictiveUnit{
			Name:           modelDeployContainerName,
			Children:       nil,
			Type:           &graphType,
			Implementation: &modelServer,
			//pytorch服务器，此处做初始化下载作用。
			//ModelURI: startJobInfo.modelUrl, 这里用了这行代码，反而导致pod初始化失败;那就依然采用如下代码，从gs下载一个模型作为初始化辅助。
			//应该是seldon后台的模型文件解析问题。
			ModelURI:   "gs:seldon-models/sklearn/iris",
			Parameters: parameters,
		}
	}

	var replica int32 = deployJobNum
	predictors := make([]seldonv1.PredictorSpec, 0)
	predictor := seldonv1.PredictorSpec{
		Name:           PredictorSpecName,
		Replicas:       &replica,
		ComponentSpecs: seldonPodSpecs,
		Graph:          graph,
	}
	predictors = append(predictors, predictor)
	metaDataName := fmt.Sprintf("%s-sdep", modelDeploy.Id)
	//seldon deployment yaml
	modelSeldonDep := &seldonv1.SeldonDeployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "machinelearning.seldon.io/v1",
			Kind:       "SeldonDeployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      metaDataName,
			Namespace: modelDeploy.UserId,
		},
		Spec: seldonv1.SeldonDeploymentSpec{
			Predictors: predictors,
		},
	}

	_, error := s.data.Cluster.CreateSeldonDeployment(context.TODO(), modelDeploy.UserId, modelSeldonDep)

	if error != nil {
		return nil, "", errors.Errorf(err, errors.ErrorModelDeployFailed)
	}

	deploymentNameSpace := fmt.Sprintf("%s/", modelDeploy.UserId)
	//根据seldon-core官方格式，进行服务url路径拼接
	serviceUrl := s.conf.Data.Ambassador.Addr + SeldonInUrl + deploymentNameSpace + metaDataName + ServiceUrlSuffix

	return resFunc, serviceUrl, nil
}

func (s *modelDeployService) checkParam(ctx context.Context, deployJob *model.ModelDeploy) (*startJobInfo, error) {
	//before commit job, check billing owner amount
	ownerId, ownerType := s.getOwner(deployJob)
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

	//资源队列
	queue := ""
	if deployJob.WorkspaceId == constant.SYSTEM_WORKSPACE_DEFAULT {
		pool, err := s.resourcePoolService.GetDefaultResourcePool(ctx, &emptypb.Empty{})
		if err != nil {
			return nil, err
		}
		queue = pool.ResourcePool.Name
	} else {
		workspace, err := s.workspaceService.GetWorkspace(ctx, &api.GetWorkspaceRequest{WorkspaceId: deployJob.WorkspaceId})
		if err != nil {
			return nil, err
		}

		queue = workspace.Workspace.ResourcePoolId
	}
	// 校验模型框架
	modelFrameType := deployJob.ModelFrame
	if modelFrameType != TensorFlowFrame && modelFrameType != PytorchFrame {
		return nil, errors.Errorf(err, errors.ErrorModelDeployForbidden)
	}

	//模型权限查询
	queryModelVersionReply, err := s.ModelAccessAuthCheck(ctx, deployJob.WorkspaceId, deployJob.UserId, deployJob.ModelId,
		deployJob.ModelVersion)
	if queryModelVersionReply == nil || err != nil {
		return nil, errors.Errorf(err, errors.ErrorModelAuthFailed)
	}
	//获取模型路径
	var modelFilePath string
	if queryModelVersionReply.Model.IsPrefab {
		modelFilePath = s.getPreFebModelSubPath(deployJob)
	} else {
		modelFilePath = s.getUserModelSubPath(deployJob)
	}
	//模型名称
	deployJob.ModelName = queryModelVersionReply.Model.ModelName
	//资源规格信息
	startJobSpecs := map[string]*startJobInfoSpec{}
	specs, err := s.resourceSpecService.ListResourceSpec(ctx, &api.ListResourceSpecRequest{})
	if err != nil {
		return nil, err
	}
	specMap := map[string]*api.ResourceSpec{}
	for _, i := range specs.ResourceSpecs {
		specMap[i.Id] = i
	}
	//资源
	resourcesReply, err := s.resourceService.ListResource(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	resourceMap := map[string]*api.Resource{}
	for _, i := range resourcesReply.Resources {
		resourceMap[i.Name] = i
	}

	// 获取推理资源规格的价格
	resSpec, err := s.resourceSpecService.GetResourceSpec(ctx, &api.GetResourceSpecRequest{Id: deployJob.ResourceSpecId})
	if err != nil {
		return nil, err
	}
	deployJob.ResSpecPrice = resSpec.ResourceSpec.Price

	//以下，获取提交任务所需的资源规格映射表及节点标签映射表
	resources := map[v1.ResourceName]resource.Quantity{}
	nodeSelectors := map[string]string{}
	//1st, 通过资源规格映射表，获取规格名称及价格
	spec, ok := specMap[deployJob.ResourceSpecId]
	if !ok {
		return nil, errors.Errorf(err, errors.ErrorInvalidRequestParameter)
	}
	//2nd, 解析资源规格包中的各项资源（cpu,gpu,memory等）的值
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
			if r.ResourceRef == "" {
				resources[v1.ResourceName(r.Name)] = quantity
			} else {
				resources[v1.ResourceName(r.ResourceRef)] = quantity
				nodeSelectors[r.BindingNodeLabelKey] = r.BindingNodeLabelValue
			}
		}
	}

	startJobSpecs[deployJob.ResourceSpecId] = &startJobInfoSpec{
		resources:     resources,
		nodeSelectors: nodeSelectors,
	}

	startModelDeployInfo := &startJobInfo{
		queue:     queue,
		modelPath: modelFilePath,
		specs:     startJobSpecs,
	}

	return startModelDeployInfo, nil
}

func (s *modelDeployService) ModelAccessAuthCheck(ctx context.Context, spaceId string, userId string,
	modelId string, modelVersion string) (*api.QueryModelVersionReply, error) {
	modelReq := &api.QueryModelVersionRequest{}
	modelReq.ModelId = modelId
	modelReq.Version = modelVersion
	queryModelVersionReply, err := s.modelService.QueryModelVersion(ctx, modelReq)
	if err != nil {
		return nil, err
	}
	if queryModelVersionReply.Model == nil {
		return nil, errors.Errorf(nil, errors.ErrorModelVersionFileNotFound)
	}
	//预置模型、分享模型以及自己的模型可以直接分布成服务
	//其他模型则无权限发布成服务
	if queryModelVersionReply.Model.IsPrefab || (queryModelVersionReply.Model.SpaceId == spaceId || queryModelVersionReply.Model.UserId == userId) {
		return queryModelVersionReply, nil
	} else {
		err := errors.Errorf(nil, errors.ErrorModelNoPermission)
		return nil, err
	}
}

//获取预置模型路径
//拼接后的路径形如：octopus/models/global/modelId/version
func (s *modelDeployService) getPreFebModelSubPath(model *model.ModelDeploy) string {
	return fmt.Sprintf("%s/%s", common.GetMinioBucket(), common.GetMinioPreModelObject(model.ModelId, model.ModelVersion))
}

//获取用户模型路径
//拼接后的路径形如：octopus/models/spaceId/userId/modelId/version
func (s *modelDeployService) getUserModelSubPath(model *model.ModelDeploy) string {
	return fmt.Sprintf("%s/%s", common.GetMinioBucket(), common.GetMinioModelObject(model.WorkspaceId, model.UserId, model.ModelId, model.ModelVersion))
}

//停止模型服务
func (s *modelDeployService) StopDepModel(ctx context.Context, req *api.StopDepRequest) (*api.StopDepReply, error) {
	modelDep, err := s.data.ModelDeployDao.GetModelDeployService(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	serviceName := fmt.Sprintf("%s-sdep", modelDep.Id)
	seldonNameSpace := modelDep.UserId
	//停止任务前，要删除掉sdep服务
	err = s.data.Cluster.DeleteSeldonDeployment(context.TODO(), seldonNameSpace, serviceName)
	if err != nil {
		//如果停止动作的任务删除操作出错了，查询下sdep任务是否存在，不存在的话，强制修改任务状态
		obj, err := s.data.Cluster.GetSeldonDeployment(context.TODO(), seldonNameSpace, serviceName)
		if obj == nil {
			err = forceUpdateStatus(s, ctx, req)
			return &api.StopDepReply{StoppedAt: time.Now().Unix()}, nil
		} else {
			return nil, errors.Errorf(err, errors.ErrorModelDeployDeleteFailed)
		}
	}

	err = forceUpdateStatus(s, ctx, req)

	if err != nil {
		return nil, err
	}

	return &api.StopDepReply{StoppedAt: time.Now().Unix()}, nil
}

func forceUpdateStatus(s *modelDeployService, ctx context.Context, req *api.StopDepRequest) error {
	now := time.Now()
	//再执行状态更新
	err := s.data.ModelDeployDao.UpdateModelDeployService(ctx, &model.ModelDeploy{
		Id:          req.Id,
		Operation:   req.Operation,
		Status:      STATE_STOPPED,
		CompletedAt: &now,
	})
	if err != nil {
		return err
	}
	return nil
}

//删除模型服务
func (s *modelDeployService) DeleteDepModel(ctx context.Context, req *api.DeleteDepRequest) (*api.DeleteDepReply, error) {
	jobs, _, err := s.data.ModelDeployDao.GetModelDeployServiceList(ctx, &model.ModelDeployListQuery{
		UserId: req.UserId,
		Ids:    req.JobIds,
	})
	if err != nil {
		return nil, err
	}

	for _, i := range jobs {
		serviceName := fmt.Sprintf("%s-sdep", i.Id)
		seldonNameSpace := i.UserId
		//删除服务前，需要判断服务状态，如果是停止状态，就直接软删除记录。其他情况，先删除服务，再软删除。
		if i.Status == STATE_PREPARING || i.Status == STATE_AVAILABLE || i.Status == STATE_FAILED {
			err = s.data.Cluster.DeleteSeldonDeployment(context.TODO(), seldonNameSpace, serviceName)
			if err != nil {
				return nil, errors.Errorf(err, errors.ErrorModelDeployFailed)
			}
			//再对数据库进行软删除
			err = s.data.ModelDeployDao.DeleteModelDeployService(ctx, i.Id)
			if err != nil {
				return nil, err
			}
		} else {
			err = s.data.ModelDeployDao.DeleteModelDeployService(ctx, i.Id)
			if err != nil {
				return nil, err
			}
		}

	}

	return &api.DeleteDepReply{DeletedAt: time.Now().Unix()}, nil
}

//模型推理
func (s *modelDeployService) ModelServiceInfer(ctx context.Context, req *api.ServiceRequest) (*api.ServiceReply, error) {
	requestData := fmt.Sprintf("{ \"data\": { \"ndarray\": %s%s", req.Data.Ndarray, "}}")
	request, err := http.NewRequest("POST", req.ServiceUrl, strings.NewReader(requestData))
	if err != nil {
		resp := &api.ServiceReply{
			Response: "failed to post request",
		}
		return resp, errors.Errorf(err, errors.ErrorModelInferRequest)
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8") //添加请求头
	client := http.Client{}                                              //创建客户端
	resp, err := client.Do(request.WithContext(context.TODO()))          //发送请求
	if err != nil {
		resp := &api.ServiceReply{
			Response: "failed to post request",
		}
		return resp, errors.Errorf(err, errors.ErrorModelInferRequest)
	}
	defer resp.Body.Close() //程序在使用完回复后必须关闭回复的主体

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp := &api.ServiceReply{
			Response: "failed to get response",
		}
		return resp, errors.Errorf(err, errors.ErrorModelInferRequest)
	}
	respStr := (*string)(unsafe.Pointer(&respBytes))

	return &api.ServiceReply{Response: *respStr}, nil
}

//获取模型服务详情
func (s *modelDeployService) GetModelDepInfo(ctx context.Context, req *api.DepInfoRequest) (*api.DepInfoReply, error) {

	deployService, err := s.data.ModelDeployDao.GetModelDeployService(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	depInfo, err := s.convertJobFromDb(deployService)
	err = copier.Copy(depInfo, deployService)
	if err != nil {
		return nil, err
	}

	return &api.DepInfoReply{
		DepInfo: depInfo,
	}, nil
}

//获取模型服务列表
func (s *modelDeployService) ListDepModel(ctx context.Context, req *api.DepListRequest) (*api.DepListReply, error) {

	query := &model.ModelDeployListQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, err
	}

	deployServices, totalSize, err := s.data.ModelDeployDao.GetModelDeployServiceList(ctx, query)
	if err != nil {
		return nil, err
	}

	deployInfos := make([]*api.DepInfo, 0)
	for _, svc := range deployServices {
		depInfo, err := s.convertJobFromDb(svc)
		err = copier.Copy(depInfo, svc)
		if err != nil {
			return nil, err
		}
		deployInfos = append(deployInfos, depInfo)
	}

	return &api.DepListReply{
		TotalSize: totalSize,
		DepInfos:  deployInfos,
	}, nil
}

func (s *modelDeployService) convertJobFromDb(jobDb *model.ModelDeploy) (*api.DepInfo, error) {
	r := &api.DepInfo{}
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

//获取模型事件
func (s *modelDeployService) ListDepEvent(ctx context.Context, req *api.DepEventListRequest) (*api.DepEventListReply, error) {
	query := &model.DeployEventQuery{}
	err := copier.Copy(query, req)
	if err != nil {
		return nil, err
	}

	events, totalSize, err := s.data.ModelDeployDao.GetModelDeployEvents(query)
	if err != nil {
		return nil, err
	}

	depEvents := make([]*api.DepEvent, 0)

	for _, value := range events {
		event := &api.DepEvent{}
		event.Timestamp = value.Timestamp
		event.Name = value.Name
		event.Reason = value.Reason
		event.Message = value.Message
		depEvents = append(depEvents, event)
	}

	return &api.DepEventListReply{
		TotalSize: totalSize,
		DepEvents: depEvents,
	}, nil
}
