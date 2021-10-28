package platform

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"server/common/dao"
	"time"

	"gorm.io/plugin/soft_delete"
	"k8s.io/apimachinery/pkg/api/resource"
)

type Tasks []*Task
type Datasets []*Dataset

type PlatformTrainJob struct {
	Id           string     `gorm:"primaryKey;type:varchar(100);not null;comment:'Id'"`
	PlatformId   string     `gorm:"type:varchar(100);not null;index;uniqueIndex:platformId_name,priority:2;comment:'用户Id'"`
	Name         string     `gorm:"type:varchar(100);not null;default:'';uniqueIndex:platformId_name,priority:1;comment:'名称'"`
	Desc         string     `gorm:"type:varchar(1024);not null;default:'';comment:'描述'"`
	Datasets     Datasets   `gorm:"type:json;comment:'dataset信息'"`
	ImageName    string     `gorm:"type:varchar(200);comment:'镜像名称'"`
	ImageVersion string     `gorm:"type:varchar(100);comment:'镜像版本'"`
	Tasks        Tasks      `gorm:"type:json;comment:'task信息'"`
	Operation    string     `gorm:"type:varchar(100);not null;default:''"`
	Status       string     `gorm:"type:varchar(100);not null;comment:'preparing/pending/running/stopped/succeeded/failed'"`
	CompletedAt  *time.Time `gorm:"type:datetime(3);comment:'结束运行时间'"`
	StartedAt    *time.Time `gorm:"type:datetime(3);comment:'开始运行时间'"`
	dao.Model
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:platformId_name,priority:4"`
}

func (PlatformTrainJob) TableName() string {
	return "platform_train_job"
}

type Dataset struct {
	Name              string `json:"name"`
	Version           string `json:"version"`
	Addr              string `json:"addr"`
	Path              string `json:"path"`
	StorageConfigName string `json:"storageConfigName"`
}

type Task struct {
	Name                  string             `json:"name"`
	Command               string             `json:"command"`
	Parameters            []*Parameter       `json:"parameters"`
	Resources             []*Resource        `json:"resources"`
	TaskNumber            int                `json:"taskNumber"`
	MinFailedTaskCount    int                `json:"minFailedTaskCount"`
	MinSucceededTaskCount int                `json:"minSucceededTaskCount"`
	IsMainRole            bool               `json:"isMainRole"`
	ShareMemory           *resource.Quantity `json:"shareMemory"`
}

type Parameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Resource struct {
	Name string `json:"name"`
	Size string `json:"size"`
}

func (r Tasks) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Tasks) Scan(input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}

func (r Datasets) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Datasets) Scan(input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}

type PlatformTrainJobListQuery struct {
	PageIndex    int
	PageSize     int
	SortBy       string
	OrderBy      string
	CreatedAtGte int64
	CreatedAtLt  int64
	Status       string
	SearchKey    string
	PlatformId   string
	Ids          []string
}

type TrainJobStastics struct {
	CreatedAtGte int64
	CreatedAtLt  int64
}

type TrainJobStasticsReply struct {
	TotalSize     int64
	SucceededSize int64
	FailedSize    int64
	StoppedSize   int64
	RunningSize   int64
	WaitingSize   int64
}
