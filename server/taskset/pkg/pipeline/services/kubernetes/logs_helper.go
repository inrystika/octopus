package kubernetes

import (
	"time"
	"strings"
    "k8s.io/client-go/tools/cache"
    "k8s.io/apimachinery/pkg/fields"
	corev1 "k8s.io/api/core/v1"
	"scheduler/pkg/pipeline/utils"
	typeJob "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
	lrucache "scheduler/pkg/pipeline/utils/lrucache"
)

type JobEvent struct {
	JobID 		string
	Namespace	string
	PodUID		map[string]string
	PodEvents 	map[string][][]*typeJob.PodEvent
}

var EventUIDCache = lrucache.NewMemCache(32768)
var PodStopChan = make(map[string](chan struct{}))
var EventStopChan = make(map[string](chan struct{}))


func getTaskRoleName(pod corev1.Pod) string {
	idx := strings.Index(pod.Name, "-")
	taskRoleName := pod.Name[idx+1:len(pod.Name)]
	return taskRoleName
}

func InitPodInfoWatch(s * Service, jobID, namespace string) {
	k8sClient := s.GetKubeClient()
    podWatchlist := cache.NewListWatchFromClient(k8sClient.CoreV1().RESTClient(), "pods", namespace, fields.Everything())
    _, podController := cache.NewInformer(
        podWatchlist,
        &corev1.Pod{},
        time.Second * 0,
        cache.ResourceEventHandlerFuncs{
            DeleteFunc: func(obj interface{}) {
				var pod *corev1.Pod = obj.(*corev1.Pod)
				if pod.Labels["volcano.sh/job-name"] == jobID {
					podEvent := &typeJob.PodEvent{
						UID:	utils.GetRandomString(12),
						Reason:	pod.Status.Reason,
						Message:pod.Status.Message,
					}
					taskRole := getTaskRoleName(*pod)
					jobInfo := &typeJob.JobInfo{}
					jobInfo.PodRoleName = make(map[string]string)
					jobInfo.PodRoleName[taskRole] = taskRole
					jobInfo.PodEvents = make(map[string][]*typeJob.PodEvent)
					jobInfo.PodEvents[taskRole] = append(jobInfo.PodEvents[taskRole], podEvent)
					s.app.Services().Job().UpdateJobSummary(jobID, jobInfo, nil, true)
				}
			},
        },
	)

	eventWatchlist := cache.NewListWatchFromClient(k8sClient.CoreV1().RESTClient(), "events", namespace, fields.Everything())
    _, eventController := cache.NewInformer(
        eventWatchlist,
        &corev1.Event{},
        time.Second * 0,
        cache.ResourceEventHandlerFuncs{
            AddFunc: func(e interface{}) {
				if event, ok := e.(*corev1.Event); ok {
					name := event.InvolvedObject.Name
					prefix := jobID + "-"
					if strings.Index(name, prefix) == 0 {
						uid := string(event.UID)
						_, ok := EventUIDCache.Get(uid)
						if ok {
							return
						}
						EventUIDCache.Set(uid, uid)
						idx := strings.Index(name, "-")
						taskRole := name[idx+1:len(name)]
						var jobInfo  = &typeJob.JobInfo{}
						jobInfo.PodRoleName = make(map[string]string)
						jobInfo.PodEvents = make(map[string][]*typeJob.PodEvent)
						jobInfo.PodRoleName[taskRole] = taskRole
						jobInfo.PodEvents[taskRole] = []*typeJob.PodEvent{}
						podEvent := &typeJob.PodEvent{
							UID:	uid,
							Reason:	event.Reason,
							Message:event.Message,
						}
						jobInfo.PodEvents[taskRole] = append(jobInfo.PodEvents[taskRole], podEvent)
						s.app.Services().Job().UpdateJobSummary(jobID, jobInfo, nil, true)
					}
				}
			},
        },
	)
	
	podStop := make(chan struct{})
	eventStop := make(chan struct{})
	PodStopChan[jobID] = podStop
	EventStopChan[jobID] = eventStop
	go podController.Run(podStop)
	go eventController.Run(eventStop)
}

func EndPodInfoWatch(jobID string) {
	podStop := PodStopChan[jobID]
	eventStop := EventStopChan[jobID]
	close(podStop)
	close(eventStop)
	delete(PodStopChan, jobID)
	delete(EventStopChan, jobID)
}

