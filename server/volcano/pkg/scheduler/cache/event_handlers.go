package cache

import (
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog"

	scheduling "server/apis/pkg/apis/scheduling/v1beta1"

	schedulingv1beta1 "server/apis/pkg/apis/scheduling/v1beta1"

	schedulingapi "server/volcano/pkg/scheduler/api"

	"volcano.sh/apis/pkg/apis/scheduling/scheme"
	"volcano.sh/volcano/pkg/scheduler/api"
	"volcano.sh/volcano/pkg/scheduler/metrics"
)

// AddQueueV1beta1 add queue to scheduler cache
func (sc *SchedulerCache) AddQueueV1beta1(obj interface{}) {
	ss, ok := obj.(*schedulingv1beta1.Queue)
	if !ok {
		klog.Errorf("Cannot convert to *schedulingv1beta1.Queue: %v", obj)
		return
	}

	queue := &scheduling.Queue{}
	if err := scheme.Scheme.Convert(ss, queue, nil); err != nil {
		klog.Errorf("Failed to convert queue from %T to %T", ss, queue)
		return
	}

	sc.Mutex.Lock()
	defer sc.Mutex.Unlock()

	klog.V(4).Infof("Add Queue(%s) into cache, spec(%#v)", ss.Name, ss.Spec)
	sc.addQueue(queue)
}

// UpdateQueueV1beta1 update queue to scheduler cache
func (sc *SchedulerCache) UpdateQueueV1beta1(oldObj, newObj interface{}) {
	oldSS, ok := oldObj.(*schedulingv1beta1.Queue)
	if !ok {
		klog.Errorf("Cannot convert oldObj to *schedulingv1beta1.Queue: %v", oldObj)
		return
	}
	newSS, ok := newObj.(*schedulingv1beta1.Queue)
	if !ok {
		klog.Errorf("Cannot convert newObj to *schedulingv1beta1.Queue: %v", newObj)
		return
	}

	if oldSS.ResourceVersion == newSS.ResourceVersion {
		return
	}

	newQueue := &scheduling.Queue{}
	if err := scheme.Scheme.Convert(newSS, newQueue, nil); err != nil {
		klog.Errorf("Failed to convert queue from %T to %T", newSS, newQueue)
		return
	}

	sc.Mutex.Lock()
	defer sc.Mutex.Unlock()
	sc.updateQueue(newQueue)
}

// DeleteQueueV1beta1 delete queue from the scheduler cache
func (sc *SchedulerCache) DeleteQueueV1beta1(obj interface{}) {
	var ss *schedulingv1beta1.Queue
	switch t := obj.(type) {
	case *schedulingv1beta1.Queue:
		ss = t
	case cache.DeletedFinalStateUnknown:
		var ok bool
		ss, ok = t.Obj.(*schedulingv1beta1.Queue)
		if !ok {
			klog.Errorf("Cannot convert to *schedulingv1beta1.Queue: %v", t.Obj)
			return
		}
	default:
		klog.Errorf("Cannot convert to *schedulingv1beta1.Queue: %v", t)
		return
	}

	sc.Mutex.Lock()
	defer sc.Mutex.Unlock()
	sc.deleteQueue(api.QueueID(ss.Name))
}

func (sc *SchedulerCache) addQueue(queue *scheduling.Queue) {
	qi := schedulingapi.NewQueueInfo(queue)
	sc.TypeQueues[qi.UID] = qi
}

func (sc *SchedulerCache) updateQueue(queue *scheduling.Queue) {
	sc.addQueue(queue)
}

func (sc *SchedulerCache) deleteQueue(id api.QueueID) {
	if queue, ok := sc.TypeQueues[id]; ok {
		delete(sc.TypeQueues, id)
		metrics.DeleteQueueMetrics(queue.Name)
	}
}
