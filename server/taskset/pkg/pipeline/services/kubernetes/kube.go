package kubernetes

import (
	"context"

	typeJob "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
	vcclientset "volcano.sh/volcano/pkg/client/clientset/versioned"
	batchinformer "volcano.sh/volcano/pkg/client/informers/externalversions"

	config "scheduler/pkg/pipeline/config"

	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

func newkubeImp(config *rest.Config, evictConfig *config.EvictConfig, userCenterConfig *config.UserCenterConfig, chargeConfig *config.ChargeConfig) *kubeImp {
	kube := &kubeImp{
		config:           config,
		evictConfig:      evictConfig,
		userCenterConfig: userCenterConfig,
		chargeConfig:     chargeConfig,
		client:           vcclientset.NewForConfigOrDie(config),
		k8sClient:        kubernetes.NewForConfigOrDie(config),
	}
	informerFactory := batchinformer.NewSharedInformerFactory(kube.client, 0)

	jobInformer := informerFactory.Batch().V1alpha1().Jobs()

	kube.jobInformer = jobInformer

	kube.jobLister = jobInformer.Lister()

	return kube
}

func (k *kubeImp) Run(stopChan chan struct{}) {

	go k.jobInformer.Informer().Run(stopChan)

	synced := cache.WaitForCacheSync(stopChan, func() []cache.InformerSynced {
		informerSynced := []cache.InformerSynced{
			k.jobInformer.Informer().HasSynced,
		}

		return informerSynced
	}()...)

	if false == synced {
		panic("Failed to WaitForCacheSync")
	}
}

func (k *kubeImp) Shutdown() {
	privilegeSwitch = false
}

func (k *kubeImp) GetClient() *kubernetes.Clientset {
	return k.k8sClient
}

func (k *kubeImp) GetVcClient() *vcclientset.Clientset {
	return k.client
}

func (k *kubeImp) AddEventHandler(handlers cache.ResourceEventHandler) {
	k.jobInformer.Informer().AddEventHandler(handlers)
}

func (k *kubeImp) Create(tj *typeJob.Job) error {
	_, err := k.client.BatchV1alpha1().Jobs(tj.Namespace).Create(context.TODO(), tj, meta.CreateOptions{})
	return err
}

func (k *kubeImp) Get(namespace, jobID string) (*typeJob.Job, error) {
	job, err := k.jobLister.Jobs(namespace).Get(jobID)

	if err == nil && job != nil {
		return job, err
	}
	//try to get vcjob from remote server
	job, err = k.client.BatchV1alpha1().Jobs(namespace).Get(context.TODO(), jobID, meta.GetOptions{})

	if apiErrors.IsNotFound(err) {
		return nil, nil
	}

	return job, err
}

func (k *kubeImp) Delete(namespace, jobID string) error {
	err := k.client.BatchV1alpha1().Jobs(namespace).
		Delete(context.TODO(), jobID, meta.DeleteOptions{})

	if nil != err && apiErrors.IsNotFound(err) {
		return nil
	}

	return err
}
