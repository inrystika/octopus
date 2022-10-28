package cluster

import (
	"context"
	"encoding/json"
	"fmt"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"

	seldonv1 "github.com/seldonio/seldon-core/operator/apis/machinelearning.seldon.io/v1"

	//seldonv2 "github.com/seldonio/seldon-core/operator/apis/machinelearning.seldon.io/v1alpha2"
	"os"
	"path/filepath"
	"server/base-server/internal/common"
	"server/base-server/internal/conf"
	"server/common/errors"
	"sync"
	"time"

	"server/common/log"

	nav1 "nodeagent/apis/agent/v1"
	naclient "nodeagent/clients/agent/clientset/versioned"
	nainformer "nodeagent/clients/agent/informers/externalversions"
	nainformerv1 "nodeagent/clients/agent/informers/externalversions/agent/v1"

	typeQueue "volcano.sh/apis/pkg/apis/scheduling/v1beta1"

	vcclient "volcano.sh/apis/pkg/client/clientset/versioned"
	libInformer "volcano.sh/apis/pkg/client/informers/externalversions"
	typejobInformer "volcano.sh/apis/pkg/client/informers/externalversions/batch/v1alpha1"
	typejobLister "volcano.sh/apis/pkg/client/listers/batch/v1alpha1"

	typejob "volcano.sh/apis/pkg/apis/batch/v1alpha1"

	fluidv1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
	seldonclient "github.com/seldonio/seldon-core/operator/client/machinelearning.seldon.io/v1/clientset/versioned"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	kubeerrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/informers"
	infov1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"sigs.k8s.io/controller-runtime/pkg/client"
	schedulingv1beta1 "volcano.sh/apis/pkg/apis/scheduling/v1beta1"
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

func NewCluster(confData *conf.Data, logger log.Logger) (Cluster, context.CancelFunc,error) {
	restConfig, err := buildConfigFromFlagsOrCluster(confData.Kubernetes.ConfigPath)
	if err != nil {
		panic(err)
	}
	return newKubernetesCluster(restConfig, logger)
}

func newKubernetesCluster(config *rest.Config, logger log.Logger) (Cluster, context.CancelFunc,error) {
	c, cancel := context.WithCancel(context.Background())

	kc := &kubernetesCluster{
		ctx:          c,
		nodes:        make(map[string]*v1.Node),
		kubeclient:   kubernetes.NewForConfigOrDie(config),
		vcClient:     vcclient.NewForConfigOrDie(config),
		naClient:     naclient.NewForConfigOrDie(config),
		seldonClient: seldonclient.NewForConfigOrDie(config),
		log:          log.NewHelper("Cluster", logger),
		config:       config,
	}
	scheme := runtime.NewScheme()
	err := fluidv1.AddToScheme(scheme)
	if err != nil {
		return nil, nil, errors.Errorf(err, errors.ErrorFluidInitFailed)
	}
	rtClient, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		return nil, nil, errors.Errorf(err, errors.ErrorFluidInitFailed)
	}
	kc.rtClient = rtClient
	informerFactory := informers.NewSharedInformerFactory(kc.kubeclient, 0)
	naInformerFactory := nainformer.NewSharedInformerFactory(kc.naClient, 0)

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

	kc.nodeActionInformer = naInformerFactory.Agent().V1().NodeActions()

	jobInformerFactory := libInformer.NewSharedInformerFactory(kc.vcClient, 0)
	vcjobInformer := jobInformerFactory.Batch().V1alpha1().Jobs()
	kc.vcjobInformer = vcjobInformer
	kc.vcjobLister = vcjobInformer.Lister()
	kc.Run()
	kc.WaitForCacheSync()
	return kc, cancel,nil
}

type kubernetesCluster struct {
	sync.Mutex
	ctx                context.Context
	log                *log.Helper
	kubeclient         *kubernetes.Clientset
	vcClient           *vcclient.Clientset
	vcjobInformer      typejobInformer.JobInformer
	vcjobLister        typejobLister.JobLister
	naClient           *naclient.Clientset
	seldonClient       *seldonclient.Clientset
	seldonInformer     informers.GenericInformer
	nodeInformer       infov1.NodeInformer
	nodeActionInformer nainformerv1.NodeActionInformer

	nodes  map[string]*v1.Node
	config *rest.Config
	rtClient           client.Client
}

func (kc *kubernetesCluster) Run() {
	go kc.nodeInformer.Informer().Run(kc.ctx.Done())
	go kc.nodeActionInformer.Informer().Run(kc.ctx.Done())
}

// WaitForCacheSync sync the cache with the api server
func (kc *kubernetesCluster) WaitForCacheSync() bool {
	return cache.WaitForCacheSync(kc.ctx.Done(),
		func() []cache.InformerSynced {
			informerSynced := []cache.InformerSynced{
				kc.nodeInformer.Informer().HasSynced,
				kc.nodeActionInformer.Informer().HasSynced,
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

func (kc *kubernetesCluster) GetNodeUnfinishedPods(ctx context.Context, nodeName string) (*v1.PodList, error) {
	fieldSelector, err := fields.ParseSelector("spec.nodeName=" + nodeName +
		",status.phase!=" + string(v1.PodSucceeded) +
		",status.phase!=" + string(v1.PodFailed))

	if err != nil {
		return nil, err
	}
	pods, err := kc.kubeclient.CoreV1().Pods("").List(ctx, metav1.ListOptions{
		FieldSelector: fieldSelector.String(),
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

func (kc *kubernetesCluster) GetNamespace(ctx context.Context, namespace string) (*v1.Namespace, error) {
	ns, err := kc.kubeclient.CoreV1().Namespaces().Get(ctx, namespace, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return ns, nil
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

	spec := typeQueue.QueueSpec{
		NodeSelector: map[string]string{nodeSelectorLabelKey: nodeSelectorLabelValue},
	}
	spec.Weight = 100
	spec.Reclaimable = &resReclaimable

	queue := &typeQueue.Queue{
		Spec: spec,
	}
	queue.ObjectMeta = metav1.ObjectMeta{
		Name:        queueName,
		Labels:      map[string]string{queueSelectLabelKey: queueSelectLabelValue},
		Annotations: meta,
	}

	_, err := kc.vcClient.SchedulingV1beta1().Queues().Create(ctx, queue, metav1.CreateOptions{})

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

func (kc *kubernetesCluster) GetPersistentVolume(ctx context.Context, name string) (*v1.PersistentVolume, error) {
	p, err := kc.kubeclient.CoreV1().PersistentVolumes().Get(ctx, name, metav1.GetOptions{})
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

func (kc *kubernetesCluster) GetPersistentVolumeClaim(ctx context.Context, namespace string, name string) (*v1.PersistentVolumeClaim, error) {
	p, err := kc.kubeclient.CoreV1().PersistentVolumeClaims(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (kc *kubernetesCluster) CreateSecret(ctx context.Context, secret *v1.Secret) (*v1.Secret, error) {
	p, err := kc.kubeclient.CoreV1().Secrets(secret.Namespace).Create(ctx, secret, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (kc *kubernetesCluster) DeletePersistentVolume(ctx context.Context, name string) error {
	err := kc.kubeclient.CoreV1().PersistentVolumes().Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return errors.Errorf(err, errors.ErrorK8sDeletePVFailed)
	}
	return nil
}

func (kc *kubernetesCluster) DeletePersistentVolumeClaim(ctx context.Context, namespace string, name string) error {
	err := kc.kubeclient.CoreV1().PersistentVolumeClaims(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return errors.Errorf(err, errors.ErrorK8sDeletePVCFailed)
	}
	return nil
}

func (kc *kubernetesCluster) DeleteSecret(ctx context.Context, namespace string, name string) error {
	err := kc.kubeclient.CoreV1().Secrets(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		return errors.Errorf(err, errors.ErrorK8sDeleteSecretFailed)
	}
	return nil
}

func (kc *kubernetesCluster) GetNodeInformer() infov1.NodeInformer {
	return kc.nodeInformer
}

func (kc *kubernetesCluster) GetPodInformer() infov1.PodInformer {
	return nil
}

func (kc *kubernetesCluster) GetNodeActionInformer() nainformerv1.NodeActionInformer {
	return kc.nodeActionInformer
}

func (kc *kubernetesCluster) CreateNodeAction(ctx context.Context, namespace string, nodeAction *nav1.NodeAction) (*nav1.NodeAction, error) {
	return kc.naClient.AgentV1().NodeActions(namespace).Create(ctx, nodeAction, metav1.CreateOptions{})
}

func (kc *kubernetesCluster) GetNodeAction(ctx context.Context, namespace, name string) (*nav1.NodeAction, error) {
	na, err := kc.naClient.AgentV1().NodeActions(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		if kubeerrors.IsNotFound(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return na, nil
}

func (kc *kubernetesCluster) DeleteNodeAction(ctx context.Context, namespace string, name string) error {
	return kc.naClient.AgentV1().NodeActions(namespace).Delete(ctx, name, metav1.DeleteOptions{})
}

func (kc *kubernetesCluster) GetPod(ctx context.Context, namespace string, name string) (*v1.Pod, error) {
	pod, err := kc.kubeclient.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		if kubeerrors.IsNotFound(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return pod, nil
}

func (kc *kubernetesCluster) CreateSeldonDeployment(ctx context.Context, namespace string, seldonDeployment *seldonv1.SeldonDeployment) (*seldonv1.SeldonDeployment, error) {
	p, err := kc.seldonClient.MachinelearningV1().SeldonDeployments(namespace).Create(ctx, seldonDeployment, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (kc *kubernetesCluster) DeleteSeldonDeployment(ctx context.Context, namespace string, serviceName string) error {
	err := kc.seldonClient.MachinelearningV1().SeldonDeployments(namespace).Delete(ctx, serviceName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (kc *kubernetesCluster) GetSeldonDeployment(ctx context.Context, namespace string, serviceName string) (*seldonv1.SeldonDeployment, error) {
	obj, err := kc.seldonClient.MachinelearningV1().SeldonDeployments(namespace).Get(ctx, serviceName, metav1.GetOptions{})
	if err != nil {
		return obj, err
	}
	return nil, nil
}

func (kc *kubernetesCluster) GetDynamicInformer(resourceType string) (informers.GenericInformer, error) {
	cfg := kc.GetClusterConfig()
	dc, err := dynamic.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}
	factory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dc, 0, corev1.NamespaceAll, nil)
	gvr, _ := schema.ParseResourceArg(resourceType)
	informer := factory.ForResource(*gvr)
	return informer, nil
}

func (kc *kubernetesCluster) RegisterDeploymentInformerCallback(onAdd common.OnDeploymentAdd, onUpdate common.OnDeploymentUpdate, onDelete common.OnDeploymentDelete) {
	kc.seldonInformer, _ = kc.GetDynamicInformer("seldondeployments.v1.machinelearning.seldon.io")
	handlers := cache.ResourceEventHandlerFuncs{
		AddFunc:    onAdd,
		DeleteFunc: onDelete,
		UpdateFunc: onUpdate,
	}
	kc.seldonInformer.Informer().AddEventHandler(handlers)
	go kc.seldonInformer.Informer().Run(kc.ctx.Done())
	cache.WaitForCacheSync(kc.ctx.Done(), func() []cache.InformerSynced {
		informerSynced := []cache.InformerSynced{
			kc.seldonInformer.Informer().HasSynced,
		}
		return informerSynced
	}()...)
}

func (kc *kubernetesCluster) RegisterJobEventHandler(handlers cache.ResourceEventHandler) {
	kc.vcjobInformer.Informer().AddEventHandler(handlers)
	go kc.vcjobInformer.Informer().Run(kc.ctx.Done())
	cache.WaitForCacheSync(kc.ctx.Done(),
		func() []cache.InformerSynced {
			informerSynced := []cache.InformerSynced{
				kc.vcjobInformer.Informer().HasSynced,
			}
			return informerSynced
		}()...,
	)
}

func (kc *kubernetesCluster) GetJob(ctx context.Context, namespace, name string) (*typejob.Job, error) {

	job, err := kc.vcjobLister.Jobs(namespace).Get(name)

	if err == nil && job != nil {
		return job, err
	}
	job, err = kc.vcClient.BatchV1alpha1().Jobs(namespace).Get(ctx, name, metav1.GetOptions{})

	if kubeerrors.IsNotFound(err) {
		return nil, nil
	}

	return job, err
}

func (kc *kubernetesCluster) CreateJob(ctx context.Context, job *typejob.Job) error {
	kc.CreateNamespace(ctx, job.Namespace)
	_, err := kc.vcClient.BatchV1alpha1().Jobs(job.Namespace).Create(ctx, job, metav1.CreateOptions{})
	return err
}

func (kc *kubernetesCluster) DeleteJob(ctx context.Context, namespace, name string) error {
	err := kc.vcClient.BatchV1alpha1().Jobs(namespace).
		Delete(ctx, name, metav1.DeleteOptions{})

	if nil != err && kubeerrors.IsNotFound(err) {
		return nil
	}
	return err
}
func (kc *kubernetesCluster) CreateFluidDataset(ctx context.Context, dataset *fluidv1.Dataset) error {
	err := kc.rtClient.Create(ctx, dataset)
	if err != nil {
		return err
	}
	return nil
}

func (kc *kubernetesCluster) DeleteFluidDataset(ctx context.Context, namespace string, name string) error {
	err := kc.rtClient.Delete(ctx, &fluidv1.Dataset{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (kc *kubernetesCluster) CreateAlluxioRuntime(ctx context.Context, alluxio *fluidv1.AlluxioRuntime) error {
	err := kc.rtClient.Create(ctx, alluxio)
	if err != nil {
		return err
	}
	return nil
}

func (kc *kubernetesCluster) DeleteAlluxioRuntime(ctx context.Context, namespace string, name string) error {
	err := kc.rtClient.Delete(ctx, &fluidv1.AlluxioRuntime{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      name,
		},
	})
	if err != nil {
		return err
	}

	return nil
}

