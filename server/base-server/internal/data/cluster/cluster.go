package cluster

import (
	"context"
	"server/base-server/internal/common"

	seldonv1 "github.com/seldonio/seldon-core/operator/apis/machinelearning.seldon.io/v1"

	nav1 "nodeagent/apis/agent/v1"
	nainformerv1 "nodeagent/clients/agent/informers/externalversions/agent/v1"

	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	infov1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/rest"
)

type ClusterClient interface {
	GetNodeInformer() infov1.NodeInformer
	GetPodInformer() infov1.PodInformer
	GetNodeActionInformer() nainformerv1.NodeActionInformer
}

type Cluster interface {
	ClusterClient
	GetClusterConfig() *rest.Config
	GetAllNodes(ctx context.Context) (map[string]v1.Node, error)
	GetNodeUnfinishedPods(ctx context.Context, nodeName string) (*v1.PodList, error)
	CreateService(ctx context.Context, service *v1.Service) error
	DeleteService(ctx context.Context, namespace string, name string) error
	CreateIngress(ctx context.Context, ingress *v1beta1.Ingress) error
	DeleteIngress(ctx context.Context, namespace string, name string) error
	CreateNamespace(ctx context.Context, namespace string) (*v1.Namespace, error)
	DeleteNamespace(ctx context.Context, namespace string) error
	ListQueue(ctx context.Context, labelSelector string) ([]byte, error)
	GetQueue(ctx context.Context, name string) ([]byte, error)
	CreateQueue(ctx context.Context, name string, queueSelectLabelKey string, queueSelectLabelValue string,
		nodeSelectorLabelKey string, nodeSelectorLabelValue string, meta map[string]string) error
	DeleteQueue(ctx context.Context, name string) error
	UpdateQueue(ctx context.Context, name string, meta map[string]string) error
	ListNode(ctx context.Context, labelSelector string) ([]byte, error)
	AddNodeLabel(ctx context.Context, name string, labelKey string, labelValue string) error
	RemoveNodeLabel(ctx context.Context, name string, labelKey string) error
	CreateAndListenJob(ctx context.Context, job *batchv1.Job, callback func(e error)) error
	CreatePersistentVolume(ctx context.Context, pv *v1.PersistentVolume) (*v1.PersistentVolume, error)
	CreatePersistentVolumeClaim(ctx context.Context, pvc *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error)
	CreateSecret(ctx context.Context, secret *v1.Secret) (*v1.Secret, error)
	DeletePersistentVolume(ctx context.Context, name string) error
	DeletePersistentVolumeClaim(ctx context.Context, namespace string, name string) error
	DeleteSecret(ctx context.Context, namespace string, name string) error
	CreateNodeAction(ctx context.Context, namespace string, nodeAction *nav1.NodeAction) (*nav1.NodeAction, error)
	GetNodeAction(ctx context.Context, namespace, name string) (*nav1.NodeAction, error)
	DeleteNodeAction(ctx context.Context, namespace string, name string) error
	GetPod(ctx context.Context, namespace string, name string) (*v1.Pod, error)
	CreateSeldonDeployment(ctx context.Context, namespace string, seldonDeployment *seldonv1.SeldonDeployment) (*seldonv1.SeldonDeployment, error)
	DeleteSeldonDeployment(ctx context.Context, namespace string, serviceName string) error
	RegisterDeploymentInformerCallback(onAdd common.OnDeploymentAdd, onUpdate common.OnDeploymentUpdate, onDelete common.OnDeploymentDelete)
}
