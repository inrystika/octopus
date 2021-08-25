package kubernetes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	utils "scheduler/pkg/pipeline/utils"
	"strconv"
	"sync"
	"time"

	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//jsoniter "github.com/json-iterator/go"
)

var privilegeSwitch = true

type SyncMapJobInfo struct {
	info map[string]*JobInfo
	*sync.RWMutex
}

type SyncMapBool struct {
	info map[string]bool
	*sync.RWMutex
}

type SyncMapString struct {
	info map[string]string
	*sync.RWMutex
}

var toEvictJobs = SyncMapJobInfo{map[string]*JobInfo{}, new(sync.RWMutex)}
var evictedJobs = SyncMapBool{map[string]bool{}, new(sync.RWMutex)}
var emailedJobs = SyncMapString{map[string]string{}, new(sync.RWMutex)}

type JobInfo struct {
	UID        string
	JobID      string
	JobName    string
	Privileger string
}

func AddToEvictJob(jobID, jobName, uid, privileger string) {
	toEvictJobs.Lock()
	toEvictJobs.info[jobID] = &JobInfo{
		UID:        uid,
		JobID:      jobID,
		JobName:    jobName,
		Privileger: privileger,
	}
	toEvictJobs.Unlock()
}

func RemoveToEvictJob(jobID string) {
	toEvictJobs.Lock()
	delete(toEvictJobs.info, jobID)
	toEvictJobs.Unlock()
}

func ClearToEvictJob() {
	toEvictJobs.info = make(map[string]*JobInfo)
}

func AddEvictedJob(jobID string) {
	evictedJobs.Lock()
	evictedJobs.info[jobID] = true
	evictedJobs.Unlock()
}

func IsEvictedJob(jobID string) bool {
	evictedJobs.RLock()
	defer evictedJobs.RUnlock()
	if _, found := evictedJobs.info[jobID]; found {
		return true
	}
	return false
}

func RemoveEvictedJob(jobID string) {
	evictedJobs.Lock()
	delete(evictedJobs.info, jobID)
	evictedJobs.Unlock()
}

func AddEmailedJob(jobID, privileger string) {
	emailedJobs.Lock()
	emailedJobs.info[jobID] = privileger
	emailedJobs.Unlock()
}

func IsEmailedJob(jobID, privileger string) bool {
	emailedJobs.RLock()
	defer emailedJobs.RUnlock()
	if val, found := emailedJobs.info[jobID]; found {
		return privileger == val
	}
	return false
}

func EvictOneJob(s *Service, jobID, namespace string) error {

	err := s.app.Services().Core().StopJob(jobID, namespace, "Evited because of the the other job has a higher privilege.")

	if err != nil {
		return err
	}

	s.logger.Info("EvictOneJob success, jobID: " + jobID)

	return nil
}

func CompensateUser(s *Service, UID string, jobId string) error {

	ratio, err := strconv.ParseFloat(s.chargeConfig.Compensate, 64)

	if err != nil {
		return err
	}

	var addr = s.chargeConfig.Address

	deductReq := "/v1/deduction/taskChargeRunHours?userId=" + UID + "&jobId=" + jobId

	body := fmt.Sprintf(`{"userId":"%s", "jobId":%s}`, UID, jobId)

	bytes, code, err := utils.DoRequest("GET", addr+deductReq, body, nil)

	if err != nil {
		return err
	}

	if code != http.StatusOK {
		return fmt.Errorf(string(bytes))
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(bytes), &result)
	if err != nil {
		return err
	}
	totalRunHours := result["result"].(map[string]interface{})["totalRunHours"].(float64)

	totalRunHours *= ratio

	body = fmt.Sprintf(`{"adminUser":"charge_admin", "vcName":"default", "clusterName":"pcl-yunnao", "userId":"%s", "rechargeAmount":%f}`, UID, totalRunHours)

	result1, code1, err1 := utils.DoRequest("POST", addr+"/v1/deduction/recharge", body, nil)

	if err1 != nil {
		return err
	}

	if code1 != http.StatusOK {
		return fmt.Errorf(string(result1))
	}

	s.logger.Info("CompensateUser success.", zap.Any("UID", UID), zap.Any("jobId", jobId))

	return nil
}

// EvictUsers ...
func Evict(s *Service, jobInfoList []*JobInfo) error {

	waitMinutes, _ := strconv.Atoi(s.evictConfig.WaitMinutes)

	timer := time.NewTimer(time.Second * time.Duration(waitMinutes) * 60)
	timer1 := time.NewTimer(time.Second * time.Duration(waitMinutes+5) * 60)

	go func() {
		<-timer.C
		for _, jobInfo := range jobInfoList {
			ns := jobInfo.UID
			vcClient := s.GetVcClient()
			job, err := vcClient.BatchV1alpha1().Jobs(ns).Get(context.TODO(), jobInfo.JobID, metav1.GetOptions{})
			if err == nil {
				if job.Status.ToEvict {
					e := EvictOneJob(s, jobInfo.JobID, ns)
					if e != nil {
						s.logger.Error("EvictOneJob Failed: " + e.Error())
					} else {
						AddEvictedJob(jobInfo.JobID)
					}
				}
			} else {
				s.logger.Error("EvictOneJob Failed, cannot find job from k8s: " + err.Error())
			}

		}
	}()

	go func() {
		<-timer1.C
		for _, jobInfo := range jobInfoList {
			if IsEvictedJob(jobInfo.JobID) {
				RemoveEvictedJob(jobInfo.JobID)
				err := CompensateUser(s, jobInfo.UID, jobInfo.JobID)
				if err != nil {
					s.logger.Error("CompensateUser Failed: " + err.Error())
				}
			}
		}
	}()

	return nil
}

func EvictWoker(s *Service) {

	jobInfoList := make([]*JobInfo, 0)
	for _, v := range toEvictJobs.info {
		jobInfoList = append(jobInfoList, v)
	}
	ClearToEvictJob()
	Evict(s, jobInfoList)
}

func PrivilegeTimer(s *Service) {
	privilegeSwitch = true
	c := time.Tick(10 * time.Second)
	for {
		if !privilegeSwitch {
			break
		}
		select {
		case <-c:
			{
				EvictWoker(s)
			}
		}
	}
}
