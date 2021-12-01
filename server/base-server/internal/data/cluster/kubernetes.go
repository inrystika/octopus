package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/seldonio/seldon-core/operator/apis/machinelearning.seldon.io/v1alpha2"
	"os"
	"path/filepath"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/common/errors"
	"sync"
	"time"

	"server/common/log"

	seldonclientset "github.com/seldonio/seldon-core/operator/client/machinelearning.seldon.io/v1alpha2/clientset/versioned"
	seldonfactory "github.com/seldonio/seldon-core/operator/client/machinelearning.seldon.io/v1alpha2/informers/externalversions"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	infov1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	schedulingv1beta1 "volcano.sh/volcano/pkg/apis/scheduling/v1beta1"
	vcclient "volcano.sh/volcano/pkg/client/clientset/versioned"
)

func buildConfigWithDefaultPath(kubeconfig string) (*rest.Config, error) {
	if kubeconfig == "" {
		homeDir, _ := os.UserHomeDir()
		kubeconfig = filepath.Join(homeDir, ".kube", "config")
	}

	cfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func buildConfigFromFlagsOrCluster(configPath string) (*rest.Config, error) {
	cfg, err1 := rest.InClusterConfig()
	if err1 == nil {
		return cfg, nil
	}

	cfg, err2 := buildConfigWithDefaultPath(configPath)
	if err2 == nil {
		return cfg, nil
	}
	return nil, fmt.Errorf("load kubernetes config failed %v %v", err1, err2)
}

func NewCluster(confData *conf.Data, logger log.Logger) Cluster {
	restConfig, err := buildConfigFromFlagsOrCluster(confData.Kubernetes.ConfigPath)
	if err != nil {
		panic(err)
	}
	return newKubernetesCluster(restConfig, logger)
}

func newKubernetesCluster(config *rest.Config, logger log.Logger) Cluster {

	seldonClientset, err := seldonclientset.NewForConfig(config)
	if err != nil {
		return nil
	}

	kc := &kubernetesCluster{
		nodes:           make(map[string]*v1.Node),
		kubeclient:      kubernetes.NewForConfigOrDie(config),
		vcClient:        vcclient.NewForConfigOrDie(config),
		seldonClientset: seldonClientset,
		log:             log.NewHelper("Cluster", logger),
		config:          config,
	}

	informerFactory := informers.NewSharedInformerFactory(kc.kubeclient, 0)

	// create informer for node information
	kc.nodeInformer = informerFactory.Core().V1().Nodes()
	kc.nodeInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.FilteringResourceEventHandler{
			FilterFunc: func(obj interface{}) bool {
				return true
			},
			Handler: cache.ResourceEventHandlerFuncs{
				AddFunc:    kc.addNode,
				UpdateFunc: kc.updateNode,
				DeleteFunc: kc.deleteNode,
			},
		},
		0,
	)

	ctx, cancel := context.WithCancel(context.TODO())
	kc.Run(ctx.Done())
	kc.WaitForCacheSync(ctx.Done())
	cancel()
	return kc
}

type kubernetesCluster struct {
	sync.Mutex
	log             *log.Helper
	kubeclient      *kubernetes.Clientset
	vcClient        *vcclient.Clientset
	seldonClientset *seldonclientset.Clientset

	nodeInformer infov1.NodeInformer

	nodes  map[string]*v1.Node
	config *rest.Config
}

func (kc *kubernetesCluster) Run(stopCh <-chan struct{}) {
	go kc.nodeInformer.Informer().Run(stopCh)
}

// WaitForCacheSync sync the cache with the api server
func (kc *kubernetesCluster) WaitForCacheSync(stopCh <-chan struct{}) bool {
	return cache.WaitForCacheSync(stopCh,
		func() []cache.InformerSynced {
			informerSynced := []cache.InformerSynced{
				kc.nodeInformer.Informer().HasSynced,
			}
			return informerSynced
		}()...,
	)
}

func (kc *kubernetesCluster) addNode(obj interface{}) {
	node := obj.(*v1.Node)

	kc.Mutex.Lock()
	defer kc.Mutex.Unlock()
	kc.nodes[node.Name] = node
}

func (kc *kubernetesCluster) updateNode(oldObj, newObj interface{}) {
	node := newObj.(*v1.Node)

	kc.Mutex.Lock()
	defer kc.Mutex.Unlock()
	kc.nodes[node.Name] = node
}

func (kc *kubernetesCluster) deleteNode(obj interface{}) {
	node := obj.(*v1.Node)

	kc.Mutex.Lock()
	defer kc.Mutex.Unlock()
	delete(kc.nodes, node.Name)
}

func (kc *kubernetesCluster) GetClusterConfig() *rest.Config {
	return kc.config
}

func (kc *kubernetesCluster) GetAllNodes(ctx context.Context) (map[string]v1.Node, error) {
	nodeList, err := kc.kubeclient.CoreV1().Nodes().List(ctx, metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	nodeMap := make(map[string]v1.Node)
	for _, node := range nodeList.Items {
		nodeMap[node.Name] = node
	}

	return nodeMap, err
}

func (kc *kubernetesCluster) GetRunningTasks(ctx context.Context) (*v1.PodList, error) {

	pods, err := kc.kubeclient.CoreV1().Pods("").List(ctx, metav1.ListOptions{
		FieldSelector: "status.phase=Running",
		LabelSelector: "volcano.sh/job-name",
	})

	if err != nil {
		kc.log.Debug(ctx, pods)
	}

	return pods, err
}

func (kc *kubernetesCluster) CreateService(ctx context.Context, service *v1.Service) error {
	_, err := kc.kubeclient.CoreV1().Services(service.Namespace).Create(ctx, service, metav1.CreateOptions{})
	if err != nil {
		return errors.Errorf(err, errors.ErrorK8sCreateServiceFailed)
	}

	return nil
}

func (kc *kubernetesCluster) DeleteService(ctx context.Context, namespace string, name string) error {
	err := kc.kubeclient.CoreV1().Services(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return errors.Errorf(err, errors.ErrorK8sDeleteServiceFailed)
	}

	return nil
}

func (kc *kubernetesCluster) CreateIngress(ctx context.Context, ingress *v1beta1.Ingress) error {
	_, err := kc.kubeclient.ExtensionsV1beta1().Ingresses(ingress.Namespace).Create(ctx, ingress, metav1.CreateOptions{})
	if err != nil {
		return errors.Errorf(err, errors.ErrorK8sCreateIngressFailed)
	}

	return nil
}

func (kc *kubernetesCluster) DeleteIngress(ctx context.Context, namespace string, name string) error {
	err := kc.kubeclient.ExtensionsV1beta1().Ingresses(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return errors.Errorf(err, errors.ErrorK8sDeleteIngressFailed)
	}

	return nil
}

func (kc *kubernetesCluster) CreateNamespace(ctx context.Context, namespace string) (*v1.Namespace, error) {
	ns, err := kc.kubeclient.CoreV1().Namespaces().Create(ctx, &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: namespace},
	}, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return ns, nil
}

func (kc *kubernetesCluster) DeleteNamespace(ctx context.Context, namespace string) error {
	err := kc.kubeclient.CoreV1().Namespaces().Delete(ctx, namespace, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (kc *kubernetesCluster) ListQueue(ctx context.Context, labelSelector string) ([]byte, error) {

	queues, err := kc.vcClient.SchedulingV1beta1().Queues().List(ctx, metav1.ListOptions{
		LabelSelector: labelSelector,
	})

	if err != nil {
		return nil, err
	}

	queueListBytes, err := json.Marshal(queues)

	return queueListBytes, err
}

func (kc *kubernetesCluster) GetQueue(ctx context.Context, name string) ([]byte, error) {

	queue, err := kc.vcClient.SchedulingV1beta1().Queues().Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		return nil, err
	}

	queueBytes, err := json.Marshal(queue)

	return queueBytes, err
}

func (kc *kubernetesCluster) CreateQueue(ctx context.Context, queueName string, queueSelectLabelKey string, queueSelectLabelValue string,
	nodeSelectorLabelKey string, nodeSelectorLabelValue string, meta map[string]string) error {

	resReclaimable := false

	_, err := kc.vcClient.SchedulingV1beta1().Queues().Create(ctx, &schedulingv1beta1.Queue{
		ObjectMeta: metav1.ObjectMeta{
			Name:        queueName,
			Labels:      map[string]string{queueSelectLabelKey: queueSelectLabelValue},
			Annotations: meta,
		},
		Spec: schedulingv1beta1.QueueSpec{
			Weight:       100,
			Reclaimable:  &resReclaimable,
			NodeSelector: map[string]string{nodeSelectorLabelKey: nodeSelectorLabelValue},
		},
	}, metav1.CreateOptions{})

	return err
}

func (kc *kubernetesCluster) UpdateQueue(ctx context.Context, name string, meta map[string]string) error {

	queue, err := kc.vcClient.SchedulingV1beta1().Queues().Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		return err
	}

	queue.ObjectMeta.Annotations = meta

	_, err = kc.vcClient.SchedulingV1beta1().Queues().Update(context.TODO(), queue, metav1.UpdateOptions{})

	return err
}

func (kc *kubernetesCluster) DeleteQueue(ctx context.Context, name string) error {

	queue, err := kc.vcClient.SchedulingV1beta1().Queues().Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		return err
	}

	queue.Status.State = schedulingv1beta1.QueueStateClosed
	_, err = kc.vcClient.SchedulingV1beta1().Queues().UpdateStatus(context.TODO(), queue, metav1.UpdateOptions{})

	if err != nil {
		return err
	}

	err = wait.Poll(2*time.Second, 10*time.Second, kc.queueClosed(ctx, name))

	if err != nil {
		return err
	}

	foreground := metav1.DeletePropagationForeground

	err = kc.vcClient.SchedulingV1beta1().Queues().Delete(ctx, name,
		metav1.DeleteOptions{
			PropagationPolicy: &foreground,
		})

	return err
}

func (kc *kubernetesCluster) queueClosed(ctx context.Context, name string) wait.ConditionFunc {
	return func() (bool, error) {
		queue, err := kc.vcClient.SchedulingV1beta1().Queues().Get(ctx, name, metav1.GetOptions{})
		if err != nil {
			return false, err
		}

		if queue.Status.State != schedulingv1beta1.QueueStateClosed {
			return false, nil
		}

		return true, nil
	}
}

func (kc *kubernetesCluster) ListNode(ctx context.Context, labelSelector string) ([]byte, error) {
	nodeList, err := kc.kubeclient.CoreV1().Nodes().List(ctx, metav1.ListOptions{LabelSelector: labelSelector})

	if err != nil {
		return nil, err
	}

	nodeListBytes, err := json.Marshal(nodeList)

	return nodeListBytes, err
}

func (kc *kubernetesCluster) AddNodeLabel(ctx context.Context, name string, labelKey string, labelValue string) error {
	node, err := kc.kubeclient.CoreV1().Nodes().Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		return err
	}

	node.ObjectMeta.Labels[labelKey] = labelValue

	_, err = kc.kubeclient.CoreV1().Nodes().Update(ctx, node, metav1.UpdateOptions{})

	return err
}

func (kc *kubernetesCluster) RemoveNodeLabel(ctx context.Context, name string, labelKey string) error {
	node, err := kc.kubeclient.CoreV1().Nodes().Get(ctx, name, metav1.GetOptions{})

	if err != nil {
		return err
	}

	delete(node.ObjectMeta.Labels, labelKey)

	_, err = kc.kubeclient.CoreV1().Nodes().Update(ctx, node, metav1.UpdateOptions{})

	return err
}

func (kc *kubernetesCluster) CreateAndListenJob(ctx context.Context, job *batchv1.Job, callback func(e error)) error {
	return nil
}

func (kc *kubernetesCluster) CreatePersistentVolume(ctx context.Context, pv *v1.PersistentVolume) (*v1.PersistentVolume, error) {
	p, err := kc.kubeclient.CoreV1().PersistentVolumes().Create(ctx, pv, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return p, nil
}
func (kc *kubernetesCluster) CreatePersistentVolumeClaim(ctx context.Context, pvc *v1.PersistentVolumeClaim) (*v1.PersistentVolumeClaim, error) {
	p, err := kc.kubeclient.CoreV1().PersistentVolumeClaims(pvc.Namespace).Create(ctx, pvc, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (kc *kubernetesCluster) CreateSeldonDeployment(ctx context.Context,namespace string, seldonDeployment *v1alpha2.SeldonDeployment) (*v1alpha2.SeldonDeployment, error ){
	p, err := kc.seldonClientset.MachinelearningV1alpha2().SeldonDeployments(namespace).Create(ctx, seldonDeployment,metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (kc *kubernetesCluster) RegisterDeploymentInformerCallback(ctx context.Context, onAdd common.OnDeploymentAdd, onUpdate common.OnDeploymentUpdate, onDelete common.OnDeploymentDelete) error {

	informerFactory := seldonfactory.NewSharedInformerFactory(kc.seldonClientset, 0)
	deploymentInformer := informerFactory.Machinelearning().V1alpha2().SeldonDeployments().Informer()
	deploymentInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		DeleteFunc: onDelete,
		UpdateFunc: onUpdate,
	})
	return nil
}
