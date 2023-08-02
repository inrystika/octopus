package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	api "server/base-server/api/v1"
	"server/common/dao"
	"time"

	"server/base-server/internal/common"

	"gorm.io/plugin/soft_delete"
	"k8s.io/apimachinery/pkg/api/resource"
)

type Configs []*Config

type ResourceSpecPrices []*ResourceSpecPrice

type TrainJob struct {
	Id                   string                     `gorm:"primaryKey;type:varchar(100);not null;comment:'Id'"`
	UserId               string                     `gorm:"type:varchar(100);not null;index;uniqueIndex:name_userId_spaceId,priority:2;comment:'用户Id'"`
	WorkspaceId          string                     `gorm:"type:varchar(100);not null;default:'';uniqueIndex:name_userId_spaceId,priority:3;;comment:'群组Id'"`
	Name                 string                     `gorm:"type:varchar(100);not null;default:'';uniqueIndex:name_userId_spaceId,priority:1;comment:'名称'"`
	Desc                 string                     `gorm:"type:varchar(1024);not null;default:'';comment:'描述'"`
	AlgorithmId          string                     `gorm:"type:varchar(100);not null;default:'';comment:'算法Id'"`
	AlgorithmVersion     string                     `gorm:"type:varchar(100);not null;default:'';comment:'算法版本'"`
	AlgorithmName        string                     `gorm:"type:varchar(100);not null;default:'';comment: '算法名称''"`
	ImageId              string                     `gorm:"type:varchar(100);not null;default:'';comment:'镜像Id'"`
	ImageName            string                     `gorm:"type:varchar(100);not null;default:'';comment: '镜像名称''"`
	ImageVersion         string                     `gorm:"type:varchar(100);not null;default:'';comment:'镜像版本'"`
	ImageUrl             string                     `gorm:"type:varchar(300);not null;default:'';comment:'镜像Url'"`
	DataSetId            string                     `gorm:"type:varchar(100);not null;default:'';comment:'数据集Id'"`
	DataSetVersion       string                     `gorm:"type:varchar(100);not null;default:'';comment:'数据集版本'"`
	DatasetName          string                     `gorm:"type:varchar(100);not null;default:'';comment:'数据集名称''"`
	IsDistributed        bool                       `gorm:"default:false;comment:'是否是分布式'"`
	Config               Configs                    `gorm:"type:json;comment:'task信息'"`
	Operation            string                     `gorm:"type:varchar(100);not null;default:''"`
	Status               string                     `gorm:"type:varchar(100);not null;comment:'preparing/pending/running/stopped/succeeded/failed'"`
	CompletedAt          *time.Time                 `gorm:"type:datetime(3);comment:'结束运行时间'"`
	StartedAt            *time.Time                 `gorm:"type:datetime(3);comment:'开始运行时间'"`
	PayAmount            float64                    `gorm:"type:decimal(10,2);not null;default:0;comment:结算机时"`
	PayStartedAt         *time.Time                 `gorm:"type:datetime(3);comment:计费起始时间"`
	PayEndedAt           *time.Time                 `gorm:"type:datetime(3);comment:计费截止时间"`
	PayStatus            api.BillingPayRecordStatus `gorm:"type:tinyint;not null;default:1;comment:扣费状态 1扣费中 2扣费完成"`
	ResSpecPrice         ResourceSpecPrices         `gorm:"type:json;comment:资源规格价格,每个键值对代表第i个子任务的资源规格单价"`
	ResourcePool         string                     `gorm:"type:varchar(300);default:'';comment:资源池"`
	Detail               string                     `gorm:"column:detail;type:json" json:"detail"`
	Mounts               common.Mounts              `gorm:"type:json;comment:挂载外部存储"`
	DisableMountUserHome bool                       `gorm:"default:false;comment:是否不挂载userhome目录"`
	DisableMountModel    bool                       `gorm:"default:false;comment:是否不挂载model目录"`
	dao.Model

	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:name_userId_spaceId,priority:4"`
}

func (TrainJob) TableName() string {
	return "train_job"
}

type Config struct {
	Name                  string             `json:"name"`
	Command               string             `json:"command"`
	Parameters            []*Parameter       `json:"parameters"`
	TaskNumber            int                `json:"taskNumber"`
	MinFailedTaskCount    int                `json:"minFailedTaskCount"`
	MinSucceededTaskCount int                `json:"minSucceededTaskCount"`
	ResourceSpecId        string             `json:"resourceSpecId"`
	ResourceSpecName      string             `json:"resourceSpecName"`
	ResourceSpecPrice     float64            `json:"resourceSpecPrice"`
	IsMainRole            bool               `json:"isMainRole"`
	ShareMemory           *resource.Quantity `json:"shareMemory"`
	Envs                  map[string]string  `json:"envs"`
}

type ResourceSpecPrice struct {
	Task  int     `json:"task"`
	Price float64 `json:"price"`
}

type Parameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (r Configs) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Configs) Scan(input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}

func (r ResourceSpecPrices) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *ResourceSpecPrices) Scan(input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}

type TrainJobListQuery struct {
	PageIndex    int
	PageSize     int
	SortBy       string
	OrderBy      string
	CreatedAtGte int64
	CreatedAtLt  int64
	Status       string
	SearchKey    string
	UserNameLike string
	UserId       string
	WorkspaceId  string
	Ids          []string
	PayStatus    api.BillingPayRecordStatus
	Statuses     []string
}

//任务模板表
type TrainJobTemplate struct {
	Id               string  `gorm:"primaryKey;type:varchar(100);not null;comment:'Id'"`
	UserId           string  `gorm:"type:varchar(100);not null;index;uniqueIndex:name_userId_spaceId,priority:2;comment:'用户Id'"`
	WorkspaceId      string  `gorm:"type:varchar(100);not null;default:'';uniqueIndex:name_userId_spaceId,priority:3;comment:'群组Id'"`
	Name             string  `gorm:"type:varchar(100);not null;default:'';uniqueIndex:name_userId_spaceId,priority:1;comment:'名称'"`
	Desc             string  `gorm:"type:varchar(1024);not null;default:'';comment:'描述'"`
	AlgorithmId      string  `gorm:"type:varchar(100);not null;default:'';comment:'算法Id'"`
	AlgorithmVersion string  `gorm:"type:varchar(100);not null;default:'';comment:'算法版本'"`
	ImageId          string  `gorm:"type:varchar(100);not null;default:'';comment:'镜像Id'"`
	ImageVersion     string  `gorm:"type:varchar(100);not null;default:'';comment:'镜像版本'"`
	DataSetId        string  `gorm:"type:varchar(100);not null;default:'';comment:'数据集Id'"`
	DataSetVersion   string  `gorm:"type:varchar(100);not null;default:'';comment:'数据集版本'"`
	IsDistributed    bool    `gorm:"default:false;comment:'是否是分布式'"`
	Config           Configs `gorm:"type:json;comment:'task信息'"`
	ResourcePool     string  `gorm:"type:varchar(300);default:'';comment:资源池"`
	dao.Model
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:name_userId_spaceId,priority:4"`
}

func (TrainJobTemplate) TableName() string {
	return "train_job_template"
}

type TrainJobTemPlateListQuery struct {
	PageIndex    int
	PageSize     int
	SortBy       string
	OrderBy      string
	CreatedAtGte int64
	CreatedAtLt  int64
	Status       string
	SearchKey    string
	UserId       string
	WorkspaceId  string
	Ids          []string
}

type TrainJobEvent struct {
	Timestamp string
	Name      string
	Reason    string
	Message   string
}

type JobEventQuery struct {
	PageIndex    int
	PageSize     int
	Id           string
	TaskIndex    int
	ReplicaIndex int
}
