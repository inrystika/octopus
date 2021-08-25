package resources

import (
	"context"
	"encoding/json"
	"fmt"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/base-server/internal/data/dao/model/resources"
	"server/common/errors"
	"server/common/leaderleaselock"
	"server/common/utils"
	"sort"
	"strings"
	"time"

	"server/common/log"

	"github.com/golang/protobuf/ptypes/empty"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
)

type ResourceService struct {
	api.UnimplementedResourceServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewResourceService(ctx context.Context, conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.ResourceServiceServer {
	logHelper := log.NewHelper("ResourceService", logger)

	startResourceDiscoveryService(ctx, conf, data, logHelper)

	return &ResourceService{
		conf: conf,
		log:  logHelper,
		data: data,
	}
}

func startResourceDiscoveryService(ctx context.Context, conf *conf.Bootstrap, data *data.Data, log *log.Helper) {
	k8sns := utils.GetEnvOrDefault("K8S_NAMESPACE", "default")
	k8sResourceDiscoveryDurationStr := conf.Service.Resource.DiscoveryDuration
	k8sResourceDiscoveryDuration, err := time.ParseDuration(k8sResourceDiscoveryDurationStr)

	if err != nil {
		log.Error(ctx, err)
		return
	}

	resDiscoveryLeaderLeaseLockName := conf.Service.Resource.DiscoveryLeaderLeaseLockName
	svcLockName := fmt.Sprintf("%s-%s-%s", conf.App.Name, conf.App.Version, resDiscoveryLeaderLeaseLockName)
	rdlock := leaderleaselock.NewLeaderLeaselock(k8sns, svcLockName, data.Cluster.GetClusterConfig())
	rdlock.RunOrRetryLeaderElection(ctx, func(ctx context.Context) {
		startResourceDiscoveryTimer(ctx, conf, data, k8sResourceDiscoveryDuration, log)
	})
}

func startResourceDiscoveryTimer(ctx context.Context, conf *conf.Bootstrap, data *data.Data, period time.Duration, log *log.Helper) {
	go wait.Until(func() {

		allNodes, err := data.Cluster.GetAllNodes(ctx)

		if err != nil {
			log.Error(ctx, errors.Errorf(err, errors.ErrorResourceDiscovery))
			return
		}

		allResources, err := data.ResourceDao.ListResource()

		if err != nil {
			log.Error(ctx, errors.Errorf(err, errors.ErrorResourceDiscovery))
			return
		}

		dbSystemResourceInfoMap := make(map[string]*resources.Resource)
		for _, dbr := range allResources {
			if dbr.ResourceRef == "" {
				dbSystemResourceInfoMap[dbr.Name] = dbr
			}
		}

		systemResourceNodesMap := collectSystemResourceNodesInfo(conf.Service.Resource.IgnoreSystemResources, allNodes)

		//Create New System Resource In DB
		for systemResName := range systemResourceNodesMap {

			if _, ok := dbSystemResourceInfoMap[systemResName]; !ok {
				//not exist, create resource
				dbResourceReq := &resources.CreateResourceRequest{
					Name:        systemResName,
					Desc:        "",
					ResourceRef: "",
				}

				_, err := data.ResourceDao.CreateResource(dbResourceReq)
				if err != nil {
					log.Error(ctx, errors.Errorf(err, errors.ErrorResourceDiscovery))
					break
				}
			}

			//let 'shm' system resource create with 'memory'
			if systemResName == "memory" {
				if _, ok := dbSystemResourceInfoMap["shm"]; !ok {
					//not exist, create resource
					dbResourceReq := &resources.CreateResourceRequest{
						Name:        "shm",
						Desc:        "",
						ResourceRef: "",
					}

					_, err := data.ResourceDao.CreateResource(dbResourceReq)
					if err != nil {
						log.Error(ctx, errors.Errorf(err, errors.ErrorResourceDiscovery))
						break
					}
				}
			}
		}

		//Delete Cluster-Not-Exist System Resource In DB
		for dbResName := range dbSystemResourceInfoMap {
			if _, ok := systemResourceNodesMap[dbResName]; !ok {
				//shm not a clusterResource,but it also a kind of systemResource with memory
				//So do not sync to delete shm resource in db when it is not exist in systemResourceNodesMap
				if dbResName == "shm" {
					if _, memok := systemResourceNodesMap["memory"]; memok {
						break
					}
				}

				//resource not exist in cluster, delete resource in db
				_, err := data.ResourceDao.DeleteResourceByName(dbResName)
				if err != nil {
					log.Error(ctx, errors.Errorf(err, errors.ErrorResourceDiscovery))
					break
				}
			}
		}

	}, period, ctx.Done())
}

func collectSystemResourceNodesInfo(ignoreSystemResources string, nodes map[string]v1.Node) map[string][]string {

	systemResourceNodesMap := make(map[string][]string)
	for nodeName, node := range nodes {
		for resName := range node.Status.Capacity {
			//systemResource filter
			if strings.Contains(ignoreSystemResources, resName.String()) {
				continue
			}

			nodeSlice := systemResourceNodesMap[resName.String()]
			nodeSlice = append(nodeSlice, nodeName)
			systemResourceNodesMap[resName.String()] = nodeSlice
		}
	}

	return systemResourceNodesMap
}

func (rsvc *ResourceService) collectCustomizedResourceNodesInfo(allResources []*resources.Resource, allNodes map[string]v1.Node) map[string][]string {
	customizedResourceNodesMap := make(map[string][]string)

	cResourceBindingNodeLabelKeyFormat := rsvc.conf.Service.Resource.CustomizedResourceBindingNodeLabelKeyFormat
	cResourceBindingNodeLabelValue := rsvc.conf.Service.Resource.CustomizedResourceBindingNodeLabelValue

	for _, resource := range allResources {
		if resource.ResourceRef != "" { //customized resource

			customizedResourceNodeLabelKey := fmt.Sprintf(cResourceBindingNodeLabelKeyFormat, resource.Name)

			for nodeName, node := range allNodes {
				if value, ok := node.ObjectMeta.Labels[customizedResourceNodeLabelKey]; ok && value == cResourceBindingNodeLabelValue {
					nodeSlice := customizedResourceNodesMap[resource.Name]
					nodeSlice = append(nodeSlice, nodeName)
					customizedResourceNodesMap[resource.Name] = nodeSlice
				}
			}
		}
	}

	return customizedResourceNodesMap
}

func (rsvc *ResourceService) ListResource(ctx context.Context, req *empty.Empty) (*api.ResourceList, error) {
	resourceList := &api.ResourceList{
		Resources: make([]*api.Resource, 0),
	}

	allResources, err := rsvc.data.ResourceDao.ListResource()

	if err != nil {
		return &api.ResourceList{}, err
	}

	allNodes, err := rsvc.data.Cluster.GetAllNodes(ctx)

	if err != nil {
		return &api.ResourceList{}, err
	}

	systemResourceNodesMap := collectSystemResourceNodesInfo(rsvc.conf.Service.Resource.IgnoreSystemResources, allNodes)
	customizedResourceNodesMap := rsvc.collectCustomizedResourceNodesInfo(allResources, allNodes)

	cResourceBindingNodeLabelKeyFormat := rsvc.conf.Service.Resource.CustomizedResourceBindingNodeLabelKeyFormat

	for _, dbr := range allResources {

		var tempResourceNodes []string
		var cResourceBindingNodeLabelKey string
		var cResourceBindingNodeLabelValue string

		if dbr.ResourceRef == "" {

			if nodes, ok := systemResourceNodesMap[dbr.Name]; ok {
				tempResourceNodes = nodes
				sort.Strings(tempResourceNodes)
			}

			//let 'shm' resource bingdingNode as memory
			if dbr.Name == "shm" {
				if nodes, ok := systemResourceNodesMap["memory"]; ok {
					tempResourceNodes = nodes
					sort.Strings(tempResourceNodes)
				}
			}
		} else {
			if nodes, ok := customizedResourceNodesMap[dbr.Name]; ok {
				tempResourceNodes = nodes
				sort.Strings(tempResourceNodes)
			}

			cResourceBindingNodeLabelValue = rsvc.conf.Service.Resource.CustomizedResourceBindingNodeLabelValue
			cResourceBindingNodeLabelKey = fmt.Sprintf(cResourceBindingNodeLabelKeyFormat, dbr.Name)
		}

		rmr := &api.Resource{
			Id:                    dbr.Id,
			Name:                  dbr.Name,
			Desc:                  dbr.Desc,
			ResourceRef:           dbr.ResourceRef,
			BindingNodes:          tempResourceNodes,
			BindingNodeLabelKey:   cResourceBindingNodeLabelKey,
			BindingNodeLabelValue: cResourceBindingNodeLabelValue,
		}

		resourceList.Resources = append(resourceList.Resources, rmr)
	}

	return resourceList, err
}

//Do not save resource nodes info to db!!!
//Because when list resources, if delete resouce label in cluster node labels, it will occurs data consistence problem between db and k8s cluster
func (rsvc *ResourceService) CreateCustomizedResource(ctx context.Context, req *api.CreateCustomizedResourceRequest) (*api.CreateCustomizedResourceReply, error) {
	//validate exist resource
	allResources, err := rsvc.data.ResourceDao.ListResource()

	if err != nil {
		return &api.CreateCustomizedResourceReply{}, errors.Errorf(err, errors.ErrorCreateResource)
	}

	for _, existResource := range allResources {
		if existResource.Name == req.Name {
			return &api.CreateCustomizedResourceReply{}, errors.Errorf(err, errors.ErrorResourceExist)
		}
	}

	allNodeMap, err := rsvc.data.Cluster.GetAllNodes(ctx)
	if err != nil {
		return &api.CreateCustomizedResourceReply{}, errors.Errorf(err, errors.ErrorCreateResource)
	}

	for _, nodeName := range req.BindingNodes {
		if _, ok := allNodeMap[nodeName]; !ok {
			return &api.CreateCustomizedResourceReply{}, errors.Errorf(nil, errors.ErrorCreateResource)
		}
	}

	resDbReq := &resources.CreateResourceRequest{
		Name:        req.Name,
		Desc:        req.Desc,
		ResourceRef: req.ResourceRef,
	}

	id, err := rsvc.data.ResourceDao.CreateResource(resDbReq)

	if err != nil {
		return &api.CreateCustomizedResourceReply{}, errors.Errorf(err, errors.ErrorCreateResource)
	}

	//binding resource label to node after create customized resource in db successfully
	//because it can using Dulplicated Resource Key DB Error to filter right create customized resource request
	err = rsvc.bindingNodeForCustomizedResource(ctx, req.Name, req.BindingNodes)

	if err != nil {
		return &api.CreateCustomizedResourceReply{}, errors.Errorf(err, errors.ErrorCreateResource)
	}

	return &api.CreateCustomizedResourceReply{Id: id}, nil
}

func (rsvc *ResourceService) bindingNodeForCustomizedResource(ctx context.Context, resourceName string, bindingNodes []string) error {

	bindingNodeMap := make(map[string]bool)

	for _, nodeName := range bindingNodes {
		bindingNodeMap[nodeName] = true
	}

	cResourceBindingNodeLabelKeyFormat := rsvc.conf.Service.Resource.CustomizedResourceBindingNodeLabelKeyFormat
	cResourceBindingNodeLabelKey := fmt.Sprintf(cResourceBindingNodeLabelKeyFormat, resourceName)
	cResourceBindingNodeLabelValue := rsvc.conf.Service.Resource.CustomizedResourceBindingNodeLabelValue

	customizedResourceNodeListBytes, err := rsvc.data.Cluster.ListNode(ctx, cResourceBindingNodeLabelKey)
	if err != nil {
		return errors.Errorf(err, errors.ErrorBindingNode)
	}

	clusterResourceNodeList := &v1.NodeList{}
	err = json.Unmarshal(customizedResourceNodeListBytes, clusterResourceNodeList)

	clusterResourceBindingNodeMap := make(map[string]string)

	for _, node := range clusterResourceNodeList.Items {
		clusterResourceBindingNodeMap[node.Name] = node.ObjectMeta.Labels[cResourceBindingNodeLabelKey]
	}

	if err != nil {
		return errors.Errorf(err, errors.ErrorBindingNode)
	}

	//remove cluster node label which not exist in updateBindingNodes Info
	for clusterNodeName := range clusterResourceBindingNodeMap {
		if _, ok := bindingNodeMap[clusterNodeName]; !ok || clusterResourceBindingNodeMap[clusterNodeName] != cResourceBindingNodeLabelValue {
			err := rsvc.data.Cluster.RemoveNodeLabel(ctx, clusterNodeName, cResourceBindingNodeLabelKey)
			if err != nil {
				return err
			}
		}
	}

	//add cluster node label which exist in updateBindingNodes Info & not exist in cluster Resource NodeList
	for maybeNewNodeName := range bindingNodeMap {
		if _, ok := clusterResourceBindingNodeMap[maybeNewNodeName]; !ok ||
			clusterResourceBindingNodeMap[maybeNewNodeName] != cResourceBindingNodeLabelValue {
			err := rsvc.data.Cluster.AddNodeLabel(ctx, maybeNewNodeName, cResourceBindingNodeLabelKey, cResourceBindingNodeLabelValue)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

//System Resource Only can update Desc
//Customized Resource can update Desc,ResourceRef,BindingNodes
func (rsvc *ResourceService) UpdateResource(ctx context.Context, req *api.UpdateResourceRequest) (*api.UpdateResourceReply, error) {
	allNodeMap, err := rsvc.data.Cluster.GetAllNodes(ctx)

	if err != nil {
		return &api.UpdateResourceReply{}, errors.Errorf(err, errors.ErrorCreateResource)
	}

	for _, nodeName := range req.BindingNodes {
		if _, ok := allNodeMap[nodeName]; !ok {
			return &api.UpdateResourceReply{}, errors.Errorf(nil, errors.ErrorCreateResource)
		}
	}

	resource, err := rsvc.data.ResourceDao.GetResource(req.Id)

	if err != nil {
		return &api.UpdateResourceReply{}, errors.Errorf(err, errors.ErrorUpdateResource)
	}

	if resource.Id == "" {
		return &api.UpdateResourceReply{}, errors.Errorf(nil, errors.ErrorUpdateResource)
	}

	if resource.ResourceRef == "" {
		//System Resource
		resource.Desc = req.Desc
		id, err := rsvc.data.ResourceDao.UpdateResource(resource)

		if err != nil {
			return &api.UpdateResourceReply{}, errors.Errorf(err, errors.ErrorUpdateResource)
		}

		return &api.UpdateResourceReply{Id: id}, nil
	} else {
		//Customized Resource
		resource.Desc = req.Desc
		resource.ResourceRef = req.ResourceRef

		id, err := rsvc.data.ResourceDao.UpdateResource(resource)

		if err != nil {
			return &api.UpdateResourceReply{}, errors.Errorf(err, errors.ErrorUpdateResource)
		}

		//binding resource label to node after update customized resource in db successfully
		err = rsvc.bindingNodeForCustomizedResource(ctx, resource.Name, req.BindingNodes)

		if err != nil {
			return &api.UpdateResourceReply{}, errors.Errorf(err, errors.ErrorUpdateResource)
		}

		return &api.UpdateResourceReply{Id: id}, nil
	}
}

func (rsvc *ResourceService) DeleteCustomizedResource(ctx context.Context, req *api.DeleteCustomizedResourceRequest) (*api.DeleteCustomizedResourceReply, error) {

	resource, err := rsvc.data.ResourceDao.GetResource(req.Id)

	if err != nil {
		return &api.DeleteCustomizedResourceReply{}, errors.Errorf(err, errors.ErrorDeleteResource)
	}

	if resource.Id == "" {
		return &api.DeleteCustomizedResourceReply{}, errors.Errorf(nil, errors.ErrorDeleteResource)
	}

	if resource.ResourceRef == "" {
		return &api.DeleteCustomizedResourceReply{}, errors.Errorf(nil, errors.ErrorDeleteResource)
	}

	//Delete Resource before delete binding node label in cluster
	//Because cluster error do not effect list resource logic
	id, err := rsvc.data.ResourceDao.DeleteResource(req.Id)

	if err != nil {
		return &api.DeleteCustomizedResourceReply{}, errors.Errorf(err, errors.ErrorDeleteResource)
	}

	cResourceBindingNodeLabelKeyFormat := rsvc.conf.Service.Resource.CustomizedResourceBindingNodeLabelKeyFormat
	cResourceBindingNodeLabelKey := fmt.Sprintf(cResourceBindingNodeLabelKeyFormat, resource.Name)
	cResourceBindingNodeLabelValue := rsvc.conf.Service.Resource.CustomizedResourceBindingNodeLabelValue

	customizedResourceLabelSelector := fmt.Sprintf("%s=%s", cResourceBindingNodeLabelKey, cResourceBindingNodeLabelValue)
	customizedResourceNodeListBytes, err := rsvc.data.Cluster.ListNode(ctx, customizedResourceLabelSelector)
	if err != nil {
		return &api.DeleteCustomizedResourceReply{}, errors.Errorf(err, errors.ErrorDeleteResource)
	}

	nodeList := &v1.NodeList{}
	err = json.Unmarshal(customizedResourceNodeListBytes, nodeList)

	if err != nil {
		return &api.DeleteCustomizedResourceReply{}, errors.Errorf(err, errors.ErrorListResourcePool)
	}

	//remove binding node from resource
	for _, node := range nodeList.Items {
		err := rsvc.data.Cluster.RemoveNodeLabel(ctx, node.Name, cResourceBindingNodeLabelKey)
		if err != nil {
			return &api.DeleteCustomizedResourceReply{}, errors.Errorf(err, errors.ErrorDeleteResource)
		}
	}

	return &api.DeleteCustomizedResourceReply{Id: id}, nil
}
