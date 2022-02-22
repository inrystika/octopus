package resources

import (
	"context"
	"fmt"
	api "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/common/errors"
	"server/common/log"
	"strings"

	"github.com/golang/protobuf/ptypes/empty"
	"k8s.io/apimachinery/pkg/api/resource"
)

type NodeService struct {
	api.UnimplementedNodeServiceServer
	conf *conf.Bootstrap
	log  *log.Helper
	data *data.Data
}

func NewNodeService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) api.NodeServiceServer {
	return &NodeService{
		conf: conf,
		log:  log.NewHelper("NodeService", logger),
		data: data,
	}
}

func (nsvc *NodeService) ListNode(ctx context.Context, req *empty.Empty) (*api.NodeList, error) {
	resNodeList := &api.NodeList{
		Nodes: []*api.Node{},
	}

	allNodeMap, err := nsvc.data.Cluster.GetAllNodes(ctx)

	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorListNode)
	}

	for nodename, node := range allNodeMap {
		resNode := &api.Node{
			Name:          nodename,
			Status:        "NotReady",
			Capacity:      make(map[string]string),
			Allocated:     make(map[string]string),
			ResourcePools: []string{},
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
				!strings.Contains(nsvc.conf.Service.Resource.IgnoreSystemResources, resname.String()) {
				resNode.Capacity[resname.String()] = quantityStr

			}
		}

		allocatedResourceMap := make(map[string]*resource.Quantity)
		pods, err := nsvc.data.Cluster.GetNodePods(ctx, nodename)
		if err != nil {
			return nil, errors.Errorf(err, errors.ErrorListNode)
		}

		for _, pod := range pods.Items {
			for _, container := range pod.Spec.Containers {
				for resname, quantity := range container.Resources.Requests {
					if _, ok := allocatedResourceMap[resname.String()]; !ok {
						newQ := quantity.DeepCopy()
						allocatedResourceMap[resname.String()] = &newQ
					} else {
						allocatedResourceMap[resname.String()].Add(quantity)
					}
				}
			}
		}

		for resname, quantity := range allocatedResourceMap {
			if !strings.Contains(nsvc.conf.Service.Resource.IgnoreSystemResources, resname) {
				resNode.Allocated[resname] = quantity.String()
			}
		}

		for resname := range resNode.Capacity {
			if _, ok := resNode.Allocated[resname]; !ok {
				resNode.Allocated[resname] = "0"
			}
		}

		for labelKey, labelValue := range node.ObjectMeta.Labels {
			rPoolBindingNodeLabelKeyFormat := nsvc.conf.Service.Resource.PoolBindingNodeLabelKeyFormat
			rPoolBindingNodeLabelValue := nsvc.conf.Service.Resource.PoolBindingNodeLabelValue
			rPoolBindingNodeLabelKeyPrefix := fmt.Sprintf(rPoolBindingNodeLabelKeyFormat, "")

			if strings.Contains(labelKey, rPoolBindingNodeLabelKeyPrefix) && labelValue == rPoolBindingNodeLabelValue {
				resourcePool := strings.ReplaceAll(labelKey, rPoolBindingNodeLabelKeyPrefix, "")
				resNode.ResourcePools = append(resNode.ResourcePools, resourcePool)
			}
		}

		resNodeList.Nodes = append(resNodeList.Nodes, resNode)
	}

	return resNodeList, nil
}
