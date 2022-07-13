package cache

import (
	"fmt"
	"sync"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"

	typeApi "server/volcano/pkg/scheduler/api"

	vcclient "volcano.sh/apis/pkg/client/clientset/versioned"
	vcinformer "volcano.sh/apis/pkg/client/informers/externalversions"
	vcinformerv1 "volcano.sh/apis/pkg/client/informers/externalversions/scheduling/v1beta1"
	schedulingapi "volcano.sh/volcano/pkg/scheduler/api"
)

// New returns a Cache implementation.
func New(config *rest.Config) Cache {
	return newSchedulerCache(config)
}

// SchedulerCache cache for the kube batch
type SchedulerCache struct {
	sync.Mutex
	vcClient             *vcclient.Clientset
	TypeQueues           map[schedulingapi.QueueID]*typeApi.QueueInfo
	vcInformerFactory    vcinformer.SharedInformerFactory
	queueInformerV1beta1 vcinformerv1.QueueInformer
}

func newSchedulerCache(config *rest.Config) *SchedulerCache {

	vcClient, err := vcclient.NewForConfig(config)
	if err != nil {
		panic(fmt.Sprintf("failed init vcClient, with err: %v", err))
	}

	sc := &SchedulerCache{
		TypeQueues: make(map[schedulingapi.QueueID]*typeApi.QueueInfo),
		vcClient:   vcClient,
	}

	vcinformers := vcinformer.NewSharedInformerFactory(sc.vcClient, 0)
	sc.vcInformerFactory = vcinformers
	// create informer(v1beta1) for Queue information
	sc.queueInformerV1beta1 = vcinformers.Scheduling().V1beta1().Queues()
	sc.queueInformerV1beta1.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    sc.AddQueueV1beta1,
		UpdateFunc: sc.UpdateQueueV1beta1,
		DeleteFunc: sc.DeleteQueueV1beta1,
	})

	return sc
}

// Snapshot returns the complete snapshot of the cluster from cache
func (sc *SchedulerCache) Snapshot() *typeApi.ClusterInfo {
	sc.Mutex.Lock()
	defer sc.Mutex.Unlock()

	snapshot := &typeApi.ClusterInfo{
		TypeQueues: make(map[schedulingapi.QueueID]*typeApi.QueueInfo),
	}

	for _, value := range sc.TypeQueues {
		snapshot.TypeQueues[schedulingapi.QueueID(value.UID)] = value.Clone()
	}

	return snapshot
}

// Run  starts the schedulerCache
func (sc *SchedulerCache) Run(stopCh <-chan struct{}) {
	sc.vcInformerFactory.Start(stopCh)
}

// WaitForCacheSync sync the cache with the api server
func (sc *SchedulerCache) WaitForCacheSync(stopCh <-chan struct{}) {
	sc.vcInformerFactory.WaitForCacheSync(stopCh)
}
