package resources

import (
	"context"
	"encoding/json"
	"fmt"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/common/errors"

	"server/common/log"

	"github.com/golang/protobuf/ptypes/empty"
	v1 "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	"volcano.sh/volcano/pkg/apis/scheduling/v1beta1"
)

type ResourcePoolService struct {
	api.UnimplementedResourcePoolServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewResourcePoolService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.ResourcePoolServiceServer {

	defaultPoolName := conf.Service.Resource.DefaultPoolName

	allNodes, err := data.Cluster.GetAllNodes(context.Background())

	if err != nil {
		panic(errors.Errorf(nil, errors.ErrorCreateResourcePool))
	}

	if defaultPoolName == "" {
		panic(errors.Errorf(nil, errors.ErrorCreateResourcePool))
	} else {
		err := initDefaultResourcePool(context.Background(), conf, data, allNodes, defaultPoolName)

		if err != nil {
			panic(errors.Errorf(nil, errors.ErrorCreateResourcePool))
		}
	}

	return &ResourcePoolService{
		conf: conf,
		log:  log.NewHelper("ResourcePoolService", logger),
		data: data,
	}
}

func initDefaultResourcePool(ctx context.Context, conf *conf.Bootstrap, data *data.Data, allNodes map[string]v1.Node, defaultPoolName string) error {
	_, err := data.Cluster.GetQueue(ctx, defaultPoolName)

	if err != nil {
		kerr := err.(*kerrors.StatusError)

		if kerr.ErrStatus.Code == 404 {
			//Create Default Queue
			selectResourcePoolLabelKey := conf.Service.Resource.PoolSelectLabelKey
			selectResourcePoolLabelValue := conf.Service.Resource.PoolSelectLabelValue

			rPoolBindingNodeLabelKeyFormat := conf.Service.Resource.PoolBindingNodeLabelKeyFormat
			rPoolBindingNodeLabelKey := fmt.Sprintf(rPoolBindingNodeLabelKeyFormat, defaultPoolName)
			rPoolBindingNodeLabelValue := conf.Service.Resource.PoolBindingNodeLabelValue

			err = data.Cluster.CreateQueue(ctx, defaultPoolName, selectResourcePoolLabelKey, selectResourcePoolLabelValue,
				rPoolBindingNodeLabelKey, rPoolBindingNodeLabelValue, map[string]string{})

			if err != nil {
				return err
			}

			err = initAllNodeLabelForDefaultResourcePool(ctx, conf, data, allNodes, defaultPoolName)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func initAllNodeLabelForDefaultResourcePool(ctx context.Context, conf *conf.Bootstrap, data *data.Data, allNodes map[string]v1.Node, defaultPoolName string) error {

	rPoolBindingNodeLabelKeyFormat := conf.Service.Resource.PoolBindingNodeLabelKeyFormat
	rPoolBindingNodeLabelKey := fmt.Sprintf(rPoolBindingNodeLabelKeyFormat, defaultPoolName)
	rPoolBindingNodeLabelValue := conf.Service.Resource.PoolBindingNodeLabelValue

	//add default pool node label to all cluster node
	for nodeName, node := range allNodes {
		if _, ok := node.Labels[rPoolBindingNodeLabelKey]; !ok {
			err := data.Cluster.AddNodeLabel(ctx, nodeName, rPoolBindingNodeLabelKey, rPoolBindingNodeLabelValue)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (rsps *ResourcePoolService) GetDefaultResourcePool(ctx context.Context, req *empty.Empty) (*api.GetResourcePoolReply, error) {
	return rsps.GetResourcePool(ctx, &api.GetResourcePoolRequest{Id: rsps.conf.Service.Resource.DefaultPoolName})
}
func (rsps *ResourcePoolService) GetResourcePool(ctx context.Context, req *api.GetResourcePoolRequest) (*api.GetResourcePoolReply, error) {

	defaultPoolName := rsps.conf.Service.Resource.DefaultPoolName

	if defaultPoolName == "" {
		return &api.GetResourcePoolReply{}, errors.Errorf(nil, errors.ErrorListResourcePool)
	}

	replyBytes, err := rsps.data.Cluster.GetQueue(ctx, req.Id)

	if err != nil {
		return &api.GetResourcePoolReply{}, errors.Errorf(err, errors.ErrorGetResourcePool)
	}

	queue := &v1beta1.Queue{}
	err = json.Unmarshal(replyBytes, queue)

	if err != nil {
		return &api.GetResourcePoolReply{}, errors.Errorf(err, errors.ErrorGetResourcePool)
	}

	realNodeList := make([]string, 0)

	rPoolBindingNodeLabelKeyFormat := rsps.conf.Service.Resource.PoolBindingNodeLabelKeyFormat
	rPoolBindingNodeLabelKey := fmt.Sprintf(rPoolBindingNodeLabelKeyFormat, queue.Name)

	nodeListBytes, err := rsps.data.Cluster.ListNode(ctx, rPoolBindingNodeLabelKey)

	if err != nil {
		return &api.GetResourcePoolReply{}, errors.Errorf(err, errors.ErrorGetResourcePool)
	}

	nodeList := &v1.NodeList{}
	err = json.Unmarshal(nodeListBytes, nodeList)

	if err != nil {
		return &api.GetResourcePoolReply{}, errors.Errorf(err, errors.ErrorGetResourcePool)
	}

	for _, node := range nodeList.Items {
		realNodeList = append(realNodeList, node.ObjectMeta.Name)
	}

	if queue.ObjectMeta.Annotations == nil {
		queue.ObjectMeta.Annotations = make(map[string]string)
	}

	resourcePoolInfoStoreKey := rsps.conf.Service.Resource.PoolInfoStoreKey
	resourcePoolInfoStr, ok := queue.ObjectMeta.Annotations[resourcePoolInfoStoreKey]
	var tempMapResourceSpecIdList map[string]*api.ResourceSpecIdList
	var tempResourcePoolDesc string
	tempResourcePoolInfo := &api.ResourcePool{}

	if ok {
		err = json.Unmarshal([]byte(resourcePoolInfoStr), &tempResourcePoolInfo)

		if err != nil {
			return &api.GetResourcePoolReply{}, errors.Errorf(err, errors.ErrorGetResourcePool)
		}

		tempResourcePoolDesc = tempResourcePoolInfo.Desc
		tempMapResourceSpecIdList = tempResourcePoolInfo.MapResourceSpecIdList

	} else {
		tempMapResourceSpecIdList = map[string]*api.ResourceSpecIdList{
			"debug":  {ResourceSpecIds: []string{}},
			"train":  {ResourceSpecIds: []string{}},
			"deploy": {ResourceSpecIds: []string{}},
		}
	}

	rp := &api.ResourcePool{
		Id:                    queue.Name,
		Name:                  queue.Name,
		Desc:                  tempResourcePoolDesc,
		Default:               queue.Name == defaultPoolName,
		MapResourceSpecIdList: tempMapResourceSpecIdList,
		BindingNodes:          realNodeList,
	}

	return &api.GetResourcePoolReply{ResourcePool: rp}, nil
}

func (rsps *ResourcePoolService) ListResourcePool(ctx context.Context, req *empty.Empty) (*api.ResourcePoolList, error) {

	defaultPoolName := rsps.conf.Service.Resource.DefaultPoolName

	if defaultPoolName == "" {
		return &api.ResourcePoolList{}, errors.Errorf(nil, errors.ErrorListResourcePool)
	}

	ResourcePoolList := &api.ResourcePoolList{
		ResourcePools: make([]*api.ResourcePool, 0),
	}

	rPoolSelectLabelKey := rsps.conf.Service.Resource.PoolSelectLabelKey
	rPoolSelectLabelValue := rsps.conf.Service.Resource.PoolSelectLabelValue
	rPoolLabelSelector := fmt.Sprintf("%s=%s", rPoolSelectLabelKey, rPoolSelectLabelValue)

	queueListBytes, err := rsps.data.Cluster.ListQueue(ctx, rPoolLabelSelector)

	if err != nil {
		return &api.ResourcePoolList{}, errors.Errorf(err, errors.ErrorListResourcePool)
	}

	queues := &v1beta1.QueueList{}

	err = json.Unmarshal(queueListBytes, queues)
	if err != nil {
		return &api.ResourcePoolList{}, errors.Errorf(err, errors.ErrorListResourcePool)
	}

	for _, queue := range queues.Items {

		if queue.Status.State != v1beta1.QueueStateOpen {
			continue
		}

		realNodeList := make([]string, 0)
		rPoolBindingNodeLabelKeyFormat := rsps.conf.Service.Resource.PoolBindingNodeLabelKeyFormat
		rPoolBindingNodeLabelKey := fmt.Sprintf(rPoolBindingNodeLabelKeyFormat, queue.Name)

		nodeListBytes, err := rsps.data.Cluster.ListNode(ctx, rPoolBindingNodeLabelKey)
		if err != nil {
			return &api.ResourcePoolList{}, errors.Errorf(err, errors.ErrorListResourcePool)
		}

		nodeList := &v1.NodeList{}
		err = json.Unmarshal(nodeListBytes, nodeList)

		if err != nil {
			return &api.ResourcePoolList{}, errors.Errorf(err, errors.ErrorListResourcePool)
		}

		for _, node := range nodeList.Items {
			realNodeList = append(realNodeList, node.ObjectMeta.Name)
		}

		if queue.ObjectMeta.Annotations == nil {
			queue.ObjectMeta.Annotations = make(map[string]string)
		}

		resourcePoolInfoStoreKey := rsps.conf.Service.Resource.PoolInfoStoreKey
		resourcePoolInfoStr, ok := queue.ObjectMeta.Annotations[resourcePoolInfoStoreKey]

		var tempMapResourceSpecIdList map[string]*api.ResourceSpecIdList
		var tempResourcePoolDesc string
		tempResourcePoolInfo := &api.ResourcePool{}

		if ok {
			err = json.Unmarshal([]byte(resourcePoolInfoStr), &tempResourcePoolInfo)

			if err != nil {
				return &api.ResourcePoolList{}, errors.Errorf(err, errors.ErrorListResourcePool)
			}

			tempResourcePoolDesc = tempResourcePoolInfo.Desc
			tempMapResourceSpecIdList = tempResourcePoolInfo.MapResourceSpecIdList

		} else {
			tempMapResourceSpecIdList = map[string]*api.ResourceSpecIdList{
				"debug":  {ResourceSpecIds: []string{}},
				"train":  {ResourceSpecIds: []string{}},
				"deploy": {ResourceSpecIds: []string{}},
			}
		}

		rspec := &api.ResourcePool{
			Id:                    queue.Name,
			Name:                  queue.Name,
			Desc:                  tempResourcePoolDesc,
			Default:               queue.Name == defaultPoolName,
			MapResourceSpecIdList: tempMapResourceSpecIdList,
			BindingNodes:          realNodeList,
		}

		ResourcePoolList.ResourcePools = append(ResourcePoolList.ResourcePools, rspec)
	}

	return ResourcePoolList, nil
}

func (rsps *ResourcePoolService) CreateResourcePool(ctx context.Context, req *api.CreateResourcePoolRequest) (*api.ResourcePoolReply, error) {

	//validate resourcespecid
	allResourceSpecList, err := rsps.data.ResourceSpecDao.ListResourceSpec(0, 0)
	if err != nil {
		return &api.ResourcePoolReply{}, errors.Errorf(nil, errors.ErrorCreateResourcePool)
	}

	allResourceSpecIdMap := make(map[string]bool)
	for _, resourceSpec := range allResourceSpecList {
		allResourceSpecIdMap[resourceSpec.Id] = true
	}
	for _, resourceSpecIdList := range req.MapResourceSpecIdList {
		for _, resourceSpecId := range resourceSpecIdList.ResourceSpecIds {
			if !allResourceSpecIdMap[resourceSpecId] {
				return &api.ResourcePoolReply{}, errors.Errorf(nil, errors.ErrorResourceSpecNotExist)
			}
		}
	}

	allNodeMap, err := rsps.data.Cluster.GetAllNodes(ctx)

	if err != nil {
		return &api.ResourcePoolReply{}, errors.Errorf(nil, errors.ErrorCreateResourcePool)
	}

	for _, nodeName := range req.BindingNodes {
		if _, ok := allNodeMap[nodeName]; !ok {
			return &api.ResourcePoolReply{}, errors.Errorf(nil, errors.ErrorCreateResourcePool)
		}
	}

	err = rsps.bindingNodeForResourcePool(ctx, req.Name, req.BindingNodes)

	if err != nil {
		return &api.ResourcePoolReply{}, errors.Errorf(err, errors.ErrorCreateResourcePool)
	}

	reqBytes, err := json.Marshal(req)

	if err != nil {
		return &api.ResourcePoolReply{}, errors.Errorf(err, errors.ErrorCreateResourcePool)
	}

	selectResourcePoolLabelKey := rsps.conf.Service.Resource.PoolSelectLabelKey
	selectResourcePoolLabelValue := rsps.conf.Service.Resource.PoolSelectLabelValue
	resourcePoolInfoStoreKey := rsps.conf.Service.Resource.PoolInfoStoreKey

	rPoolBindingNodeLabelKeyFormat := rsps.conf.Service.Resource.PoolBindingNodeLabelKeyFormat
	rPoolBindingNodeLabelKey := fmt.Sprintf(rPoolBindingNodeLabelKeyFormat, req.Name)
	rPoolBindingNodeLabelValue := rsps.conf.Service.Resource.PoolBindingNodeLabelValue

	err = rsps.data.Cluster.CreateQueue(ctx, req.Name, selectResourcePoolLabelKey, selectResourcePoolLabelValue,
		rPoolBindingNodeLabelKey, rPoolBindingNodeLabelValue, map[string]string{
			resourcePoolInfoStoreKey: string(reqBytes),
		})

	if err != nil {
		return &api.ResourcePoolReply{}, errors.Errorf(err, errors.ErrorCreateResourcePool)
	}

	return &api.ResourcePoolReply{Id: req.Name}, nil
}

func (rsps *ResourcePoolService) UpdateResourcePool(ctx context.Context, req *api.UpdateResourcePoolRequest) (*api.ResourcePoolReply, error) {
	//validate resourcespecid
	allResourceSpecList, err := rsps.data.ResourceSpecDao.ListResourceSpec(0, 0)
	if err != nil {
		return &api.ResourcePoolReply{}, errors.Errorf(nil, errors.ErrorCreateResourcePool)
	}

	allResourceSpecIdMap := make(map[string]bool)
	for _, resourceSpec := range allResourceSpecList {
		allResourceSpecIdMap[resourceSpec.Id] = true
	}
	for _, resourceSpecIdList := range req.MapResourceSpecIdList {
		for _, resourceSpecId := range resourceSpecIdList.ResourceSpecIds {
			if !allResourceSpecIdMap[resourceSpecId] {
				return &api.ResourcePoolReply{}, errors.Errorf(nil, errors.ErrorResourceSpecNotExist)
			}
		}
	}

	allNodeMap, err := rsps.data.Cluster.GetAllNodes(ctx)
	if err != nil {
		return &api.ResourcePoolReply{}, errors.Errorf(err, errors.ErrorUpdateResourcePool)
	}

	for _, nodeName := range req.BindingNodes {
		if _, ok := allNodeMap[nodeName]; !ok {
			return &api.ResourcePoolReply{}, errors.Errorf(nil, errors.ErrorUpdateResourcePool)
		}
	}

	err = rsps.bindingNodeForResourcePool(ctx, req.Id, req.BindingNodes)

	if err != nil {
		return &api.ResourcePoolReply{}, errors.Errorf(err, errors.ErrorUpdateResourcePool)
	}

	reqBytes, err := json.Marshal(req)

	if err != nil {
		return &api.ResourcePoolReply{}, errors.Errorf(err, errors.ErrorUpdateResourcePool)
	}

	resourcePoolInfoStoreKey := rsps.conf.Service.Resource.PoolInfoStoreKey
	err = rsps.data.Cluster.UpdateQueue(ctx, req.Id, map[string]string{
		resourcePoolInfoStoreKey: string(reqBytes),
	})

	if err != nil {
		return &api.ResourcePoolReply{}, errors.Errorf(err, errors.ErrorUpdateResourcePool)
	}

	return &api.ResourcePoolReply{Id: req.Id}, nil
}

func (rsps *ResourcePoolService) bindingNodeForResourcePool(ctx context.Context, resourcePoolName string, bindingNodes []string) error {

	bindingNodeMap := make(map[string]bool)

	for _, nodeName := range bindingNodes {
		bindingNodeMap[nodeName] = true
	}

	rPoolBindingNodeLabelKeyFormat := rsps.conf.Service.Resource.PoolBindingNodeLabelKeyFormat
	rPoolBindingNodeLabelKey := fmt.Sprintf(rPoolBindingNodeLabelKeyFormat, resourcePoolName)
	rPoolBindingNodeLabelValue := rsps.conf.Service.Resource.PoolBindingNodeLabelValue

	rPoolNodeListBytes, err := rsps.data.Cluster.ListNode(ctx, rPoolBindingNodeLabelKey)
	if err != nil {
		return errors.Errorf(err, errors.ErrorBindingNode)
	}

	clusterResourcePoolNodeList := &v1.NodeList{}
	err = json.Unmarshal(rPoolNodeListBytes, clusterResourcePoolNodeList)

	if err != nil {
		return errors.Errorf(err, errors.ErrorBindingNode)
	}

	clusterResourcePoolBindingNodeMap := make(map[string]string)

	for _, node := range clusterResourcePoolNodeList.Items {
		clusterResourcePoolBindingNodeMap[node.Name] = node.ObjectMeta.Labels[rPoolBindingNodeLabelKey]
	}

	//remove cluster node label which not exist in updateBindingNodes Info
	for clusterNodeName := range clusterResourcePoolBindingNodeMap {
		if _, ok := bindingNodeMap[clusterNodeName]; !ok || clusterResourcePoolBindingNodeMap[clusterNodeName] != rPoolBindingNodeLabelValue {
			err := rsps.data.Cluster.RemoveNodeLabel(ctx, clusterNodeName, rPoolBindingNodeLabelKey)
			if err != nil {
				return err
			}
		}
	}

	//add cluster node label which exist in updateBindingNodes Info & not exist in cluster ResourcePool NodeList
	for maybeNewNodeName := range bindingNodeMap {
		if _, ok := clusterResourcePoolBindingNodeMap[maybeNewNodeName]; !ok ||
			clusterResourcePoolBindingNodeMap[maybeNewNodeName] != rPoolBindingNodeLabelValue {
			err := rsps.data.Cluster.AddNodeLabel(ctx, maybeNewNodeName, rPoolBindingNodeLabelKey, rPoolBindingNodeLabelValue)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (rsps *ResourcePoolService) DeleteResourcePool(ctx context.Context, req *api.DeleteResourcePoolRequest) (*api.ResourcePoolReply, error) {
	err := rsps.removeNodeBindingForResourcePool(ctx, req.Id)
	if err != nil {
		return &api.ResourcePoolReply{}, errors.Errorf(nil, errors.ErrorDeleteResourcePool)
	}

	err = rsps.data.Cluster.DeleteQueue(ctx, req.Id)

	if err != nil {
		return &api.ResourcePoolReply{}, errors.Errorf(err, errors.ErrorDeleteResourcePool)
	}

	return &api.ResourcePoolReply{Id: req.Id}, nil
}

func (rsps *ResourcePoolService) removeNodeBindingForResourcePool(ctx context.Context, resourcePoolName string) error {

	rPoolBindingNodeLabelKeyFormat := rsps.conf.Service.Resource.PoolBindingNodeLabelKeyFormat
	rPoolBindingNodeLabelKey := fmt.Sprintf(rPoolBindingNodeLabelKeyFormat, resourcePoolName)
	rPoolBindingNodeLabelValue := rsps.conf.Service.Resource.PoolBindingNodeLabelValue

	rPoolNodeLabelSelector := fmt.Sprintf("%s=%s", rPoolBindingNodeLabelKey, rPoolBindingNodeLabelValue)

	resPoolNodeListBytes, err := rsps.data.Cluster.ListNode(ctx, rPoolNodeLabelSelector)
	if err != nil {
		return errors.Errorf(err, errors.ErrorBindingNode)
	}

	nodeList := &v1.NodeList{}
	err = json.Unmarshal(resPoolNodeListBytes, nodeList)

	if err != nil {
		return errors.Errorf(err, errors.ErrorBindingNode)
	}

	//remove expired node label binding from resource pool
	for _, node := range nodeList.Items {
		err := rsps.data.Cluster.RemoveNodeLabel(ctx, node.Name, rPoolBindingNodeLabelKey)
		if err != nil {
			return err
		}
	}

	return nil
}
