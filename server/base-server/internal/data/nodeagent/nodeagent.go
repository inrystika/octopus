package nodeagent

import (
	v1 "k8s.io/api/core/v1"
	infov1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"
	"server/base-server/internal/common"
	"server/common/errors"
	"sync"
)

const NODE_AGENT_SERVER_ROLE = "nodeagent"

type NodeAgentManager interface {
	Switch(nodeName string) (NodeAgent, error)
}

type CommitImageOpt struct {
	author  string
	changes []string
	message string
}

type NodeAgent interface {
	CommitImage(CommitImageOpt) error
}

type nodeAgentManager struct {
	agentMap    sync.Map
	podInformer infov1.PodInformer
}

func NewNodeAgentManger(podInformer infov1.PodInformer) NodeAgentManager {
	manager := &nodeAgentManager{
		podInformer: podInformer,
	}
	manager.podInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.FilteringResourceEventHandler{
			FilterFunc: func(obj interface{}) bool {
				// match pod of nodeagent
				pod := obj.(*v1.Pod)
				for k, v := range pod.Labels {
					if k == common.OctopusServiceRoleLabel && v == NODE_AGENT_SERVER_ROLE {
						return true
					}
				}
				return false
			},
			Handler: cache.ResourceEventHandlerFuncs{
				AddFunc:    manager.addNodeAgent,
				UpdateFunc: manager.updateNodeAgent,
				DeleteFunc: manager.deleteNodeAgent,
			},
		},
		0,
	)
	return manager
}

func (n *nodeAgentManager) addNodeAgent(obj interface{}) {

}

func (n *nodeAgentManager) updateNodeAgent(oldObj interface{}, newObj interface{}) {

}

func (n *nodeAgentManager) deleteNodeAgent(obj interface{}) {

}

func (n *nodeAgentManager) Switch(nodeName string) (NodeAgent, error) {
	if obj, ok := n.agentMap.Load(nodeName); ok {
		return obj.(NodeAgent), nil
	}
	return nil, errors.Errorf(nil, errors.ErrorDirNotExisted)
}
