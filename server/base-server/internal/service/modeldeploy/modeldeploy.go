package modeldeploy

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	seldonv1 "github.com/seldonio/seldon-core/operator/apis/machinelearning.seldon.io/v1"
	"github.com/seldonio/seldon-core/operator/apis/machinelearning.seldon.io/v1alpha2"
	seldonv2 "github.com/seldonio/seldon-core/operator/apis/machinelearning.seldon.io/v1alpha2"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	api "server/base-server/api/v1"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model"
	"server/common/constant"
	"server/common/errors"
	"server/common/log"
	"server/common/utils"
	"time"
)

const (
	deployJobKind       = "deploy_job"
	deployJobNum        = 1
	MetaNamePrefix      = "deploy-"
	TensorFlowFrame     = "tensorflow"
	PytorchFrame        = "pytorch"
	ModelVolumePath     = "model_volume_path"
	SeldonDockerWorkDir = "/app/models"
	PredictorSpecName   = "default"
	SeldonInUrl         = "/seldon/"
	ServiceUrlSuffix    = "/api/v1.0/predictions"
	UserModelDir        = "user_model_dir"
	STATE_PREPARING     = "Preparing"
	STATE_AVAILABLE     = "Available"
	STATE_CREATING      = "Creating"
	STATE_FAILED        = "Failed"
	STATE_STOPPED       = "Stopped"
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

	//s.modelDepBilling(context.Background())
	ctx := context.Background()
	s.data.Cluster.RegisterDeploymentInformerCallback(ctx,
		func(obj interface{}) {},
		func(old, obj interface{}) {
			objSeldon, ok := obj.(*seldonv2.SeldonDeployment)
			if !ok {
				return
			}
			deployService, err := s.data.ModelDeployDao.GetModelDeployService(ctx, objSeldon.Name)
			if err != nil {
				return
			}
			newState := string(objSeldon.Status.State)
			if newState == deployService.Status {
				return
			}
			update := &model.ModelDeploy{
				Id:     objSeldon.Name,
				Status: newState,
			}
			now := time.Now()
			if newState == STATE_AVAILABLE {
				update.StartedAt = &now
			} else if newState == STATE_FAILED {
				update.CompletedAt = &now
			}
			s.data.ModelDeployDao.UpdateModelDeployService(ctx, update)
		},
		func(obj interface{}) {},
	)

	return s, nil
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

	modelDeployName := modelDeploy.Name

	//挂载卷
	volumeMounts := []v1.VolumeMount{
		{
			Name:      "modelFilePath",
			MountPath: s.conf.Service.DockerModelDeployPath,
			SubPath:   startJobInfo.modelPath,
			ReadOnly:  false,
		},
		{
			Name:      "localtime",
			MountPath: "/etc/localtime",
		},
	}

	volumes := []v1.Volume{
		{
			Name: "modelFilePath",
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

	parameters := []seldonv1.Parameter{
		{
			Name:  ModelVolumePath,
			Type:  "STRING",
			Value: SeldonDockerWorkDir,
		},
		{
			Name:  UserModelDir,
			Type:  "STRING",
			Value: startJobInfo.modelPath,
		},
	}

	seldonPodSpecs := make([]*seldonv1.SeldonPodSpec, 0)
	seldonPodSpec := &seldonv1.SeldonPodSpec{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:         modelDeployName,
					VolumeMounts: volumeMounts,
				},
			},
			Volumes: volumes,
		},
	}
	seldonPodSpecs = append(seldonPodSpecs, seldonPodSpec)

	var modelServer seldonv1.PredictiveUnitImplementation
	if startJobInfo.modelFrame == TensorFlowFrame {
		modelServer = "TENSORFLOW_SERVER"
	} else if startJobInfo.modelFrame == PytorchFrame {
		modelServer = "PYTORCH_SERVER"
	}

	graph := seldonv1.PredictiveUnit{
		Name:           modelDeployName,
		Children:       nil,
		Implementation: &modelServer,
		//todo 此处可以修改为pvc挂载
		ModelURI:   "gs:seldon-models/sklearn/iris",
		Parameters: parameters,
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

	//seldon deployment yaml
	modelSeldonDep := &v1alpha2.SeldonDeployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "machinelearning.seldon.io/v1alpha2",
			Kind:       "SeldonDeployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      modelDeploy.UserId,
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

	serviceUrl := s.conf.Data.Ambassador.Addr + SeldonInUrl + modelDeploy.UserId + "/" + modelDeployName + ServiceUrlSuffix

	return resFunc, serviceUrl, nil
}

func (s *modelDeployService) checkParam(ctx context.Context, modelDeploy *model.ModelDeploy) (*startJobInfo, error) {
	//before commit job, check billing owner amount
	ownerId, ownerType := s.getOwner(modelDeploy)
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
	if modelDeploy.WorkspaceId == constant.SYSTEM_WORKSPACE_DEFAULT {
		pool, err := s.resourcePoolService.GetDefaultResourcePool(ctx, &emptypb.Empty{})
		if err != nil {
			return nil, err
		}
		queue = pool.ResourcePool.Name
	} else {
		workspace, err := s.workspaceService.GetWorkspace(ctx, &api.GetWorkspaceRequest{WorkspaceId: modelDeploy.WorkspaceId})
		if err != nil {
			return nil, err
		}

		queue = workspace.Workspace.ResourcePoolId
	}
	// 校验模型框架
	modelFrameType := modelDeploy.ModelFrame
	if modelFrameType != TensorFlowFrame && modelFrameType != PytorchFrame {
		return nil, errors.Errorf(err, errors.ErrorModelDeployForbidden)
	}

	//模型权限查询
	queryModelVersionReply, err := s.ModelAccessAuthCheck(ctx, modelDeploy.WorkspaceId, modelDeploy.UserId, modelDeploy.Id, modelDeploy.ModelVersion)
	if queryModelVersionReply == nil || err != nil {
		return nil, err
	}

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
	//提交任务所需的资源规格映射表及节点标签映射表
	resources := map[v1.ResourceName]resource.Quantity{}
	nodeSelectors := map[string]string{}
	//通过资源规格映射表，获取规格名称及价格
	spec, ok := specMap[modelDeploy.ResourceSpecId]
	if !ok {
		return nil, errors.Errorf(err, errors.ErrorInvalidRequestParameter)
	}
	//解析资源规格包中的各项资源（cpu,gpu,memory等）的值
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

	startJobSpecs[modelDeploy.ResourceSpecId] = &startJobInfoSpec{
		resources:     resources,
		nodeSelectors: nodeSelectors,
	}

	startModelDeployInfo := &startJobInfo{
		queue:     queue,
		modelPath: s.getModelSubPath(modelDeploy),
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
	//非自己模型，非分享模型，则无权限发布成服务
	if queryModelVersionReply.Model.SpaceId != spaceId || queryModelVersionReply.Model.UserId != userId || !queryModelVersionReply.Model.IsPrefab {
		err := errors.Errorf(nil, errors.ErrorModelNoPermission)
		return nil, err
	}

	return queryModelVersionReply, nil
}

//获取模型路径
func (s *modelDeployService) getModelSubPath(model *model.ModelDeploy) string {
	return fmt.Sprintf("%s/%s", common.GetMinioBucket(), common.GetMinioModelObject(model.WorkspaceId, model.UserId, model.ModelId, model.ModelVersion))
}

//停止模型服务
func (s *modelDeployService) StopDepModel(ctx context.Context, req *api.StopDepRequest) (*api.StopDepReply, error) {
	modelDep, err := s.data.ModelDeployDao.GetModelDeployService(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	//pipeline删除任务成功后，任务从running转为terminate转态会触发callback机制,更新base-server中的任务状态信息。
	serviceName := modelDep.ModelName
	seldonNameSpace := modelDep.UserId
	//停止任务
	err = s.data.Cluster.DeleteSeldonDeployment(context.TODO(), seldonNameSpace, serviceName)
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorModelDeployFailed)
	}

	now := time.Now()
	err = s.data.ModelDeployDao.UpdateModelDeployService(ctx, &model.ModelDeploy{
		Id:          req.Id,
		Operation:   req.Operation,
		Status:      STATE_STOPPED,
		CompletedAt: &now,
	})
	if err != nil {
		return nil, err
	}

	return &api.StopDepReply{StoppedAt: time.Now().Unix()}, nil
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
		serviceName := i.ModelName
		seldonNameSpace := i.UserId
		//删除服务前先停止服务
		err = s.data.Cluster.DeleteSeldonDeployment(context.TODO(), seldonNameSpace, serviceName)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorModelDeployFailed)
		}
		//软删除
		err = s.data.ModelDeployDao.DeleteModelDeployService(ctx, i.Id)
		if err != nil {
			return nil, err
		}
	}

	return &api.DeleteDepReply{DeletedAt: time.Now().Unix()}, nil
}

//获取模型服务详情
func (s *modelDeployService) GetModelDepInfo(ctx context.Context, req *api.DepInfoRequest) (*api.DepInfoReply, error) {

	deployService, err := s.data.ModelDeployDao.GetModelDeployService(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	depInfo := &api.DepInfo{}
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

	deployservices, totalSize, err := s.data.ModelDeployDao.GetModelDeployServiceList(ctx, query)
	if err != nil {
		return nil, err
	}

	deployInfos := make([]*api.DepInfo, 0)
	for _, svc := range deployservices {

		depInfo := &api.DepInfo{}
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
