package pipeline

import (
	"encoding/json"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typeJob "volcano.sh/volcano/pkg/apis/batch/v1alpha1"
)

type Reply struct {
	Code    string          `json:"code"`
	Msg     string          `json:"msg"`
	Payload json.RawMessage `json:"payload"`
}

type SubmitJobParam struct {
	UserID       string                 `json:"userID"`
	JobKind      string                 `json:"jobKind"`
	JobName      string                 `json:"jobName"`
	Header       map[string]interface{} `json:"header"`
	Job          *typeJob.Job           `json:"job"`
	JobNamespace string                 `json:"jobNameSpace"`
	Cluster      string                 `json:"cluster"`
}

type SubmitJobReply struct {
	JobId string `json:"jobID"`
}

type StopJobReq struct {
	JobId string `json:"JobId"`
}

type StopJobResp struct {
	Payload string `json:"payload"`
}

type BatchDeleteJobReq struct {
	JobIdList []string `json:"jobIdList"`
	Reason    string   `json:"reason"`
	UserID    string   `json:"userId"`
}

type DeleteJobResp struct {
}

type TrainJobDetailReq struct {
	JobId string `json:"jobId"`
}

type QueryParams struct {
	UserID     string
	JobName    string
	JobKind    string
	State      string
	Order      string
	Cluster    string
	PageNumber int64
	PageSize   int64
	StartTime  *time.Time
	EndTime    *time.Time
	IsDelete   bool
}

type UpdateJobParam struct {
	JobID  string `json:"jobID"`
	Reason string `json:"reason"`
}

type GetJobDetailParam struct {
	JobID string `json:"jobID"`
}

type UpsertFeatureParam struct {
	FeatureName string       `json:"name"`
	Author      string       `json:"author"`
	Description string       `json:"description"`
	Enabled     bool         `json:"enabled"`
	JobSelector *JobSelector `json:"jobSelector"`
	Plugins     []*Plugin    `json:"plugins"`
}

type Condition struct {
	Name   string `json:"name"`
	Key    string `json:"key"`
	Expect string `json:"expect"`
}

type JobSelector struct {
	Conditions []*Condition `json:"conditions"`
	Expression string       `json:"expression"`
	States     []string     `json:"states"`
}

type Plugin struct {
	Key               string       `json:"key"`
	PluginType        string       `json:"pluginType"`
	CallAddress       string       `json:"callAddress"`
	Description       string       `json:"description"`
	JobSelector       *JobSelector `json:"jobSelector"`
	ExecutionSequence int64        `json:"sequence"` //plugin被执行的顺序
}

type JobStatusDetail struct {
	Version              string                `json:"version"`
	Job                  *JobSummary           `json:"job"`
	Cluster              *ClusterInfo          `json:"cluster"`
	Tasks                []*TaskInfo           `json:"tasks"`
	PlatformSpecificInfo *PlatformSpecificInfo `json:"platformSpecificInfo"`
}

type BatchJobStatusDetail struct {
	Details []*JobStatusDetail `json:"details"`
}

type JobSummary struct {
	ID                string       `json:"id"`
	Name              string       `json:"name"`
	Type              string       `json:"type"`
	State             string       `json:"state"`
	UserID            string       `json:"userID"`
	StartAt           *metav1.Time `json:"startAt"`
	FinishedAt        *metav1.Time `json:"finishedAt"`
	TotalRetriedCount uint         `json:"totalRetriedCount"`
	ExitCode          int32        `json:"exitCode"`
	ExitDiagnostics   string       `json:"exitDiagnostics"`
}

type ClusterInfo struct {
	Identity string `json:"identity"`
}

type TaskInfo struct {
	Name                  string         `json:"name"`
	Image                 string         `json:"image"`
	State                 string         `json:"state"`
	Command               []string       `json:"command"`
	ReplicaAmount         uint           `json:"replicaAmount"`
	MaxFailedTaskCount    int32          `json:"maxFailedTaskCount"`
	MinSucceededTaskCount int32          `json:"minSucceededTaskCount"`
	Resource              string         `json:"resource"`
	Replicas              []*ReplicaInfo `json:"replicas"`
}

type ReplicaInfo struct {
	Index           uint         `json:"index"`
	State           string       `json:"state"`
	RetriedCount    uint         `json:"retriedCount"`
	StartAt         *metav1.Time `json:"startAt"`
	FinishedAt      *metav1.Time `json:"finishedAt"`
	ContainerID     string       `json:"containerID"`
	ContainerHostIP string       `json:"containerHostIP"`
	ExitCode        int32        `json:"exitCode"`
	ExitDiagnostics string       `json:"exitDiagnostics"`
}

type PlatformSpecificInfo struct {
	Platform        string             `json:"platform"`
	ApiVersion      string             `json:"apiVersion"`
	Namespace       string             `json:"namespace"`
	InstanceUID     string             `json:"instanceUID"`
	ConfigMapUID    string             `json:"configMapUID"`
	ConfigMapName   string             `json:"configMapName"`
	TaskRuntimeInfo []*TaskRuntimeInfo `json:"taskRuntimeInfo"`
}

type TaskRuntimeInfo struct {
	Name         string                    `json:"name"`
	NodeSelector map[string]string         `json:"nodeSelector"`
	Replicas     []*TaskRuntimeReplicaInfo `json:"replicas"`
	VolumeMounts []corev1.Volume           `json:"volumeMounts"`
}

type TaskRuntimeReplicaInfo struct {
	Index     uint   `json:"index"`
	PodIP     string `json:"podIP"`
	PodUID    string `json:"podUID"`
	PodName   string `json:"podName"`
	PodHostIP string `json:"podHostIP"`
	PodReason string `json:"podReason"`
}

const (
	PREPARING = "preparing"
	PENDING   = "pending"
	RUNNING   = "running"
	FAILED    = "failed"
	SUCCEEDED = "succeeded"
	STOPPED   = "stopped"
	SUSPENDED = "suspended"
	UNKNOWN   = "unknown"
)

func IsCompletedState(state string) bool {
	if strings.EqualFold(state, STOPPED) ||
		strings.EqualFold(state, SUCCEEDED) ||
		strings.EqualFold(state, FAILED) {
		return true
	}

	return false
}

func IsRunningOrCompletedState(state string) bool {
	if strings.EqualFold(state, STOPPED) ||
		strings.EqualFold(state, SUCCEEDED) ||
		strings.EqualFold(state, FAILED) ||
		strings.EqualFold(state, RUNNING) {
		return true
	}

	return false
}

func JobRunningState(state string) bool {
	return strings.EqualFold(state, RUNNING)
}

func NonCompletedStates() []string {
	return []string{PREPARING, PENDING, RUNNING}
}
