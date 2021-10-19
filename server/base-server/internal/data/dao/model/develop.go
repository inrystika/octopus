package model

import (
	api "server/base-server/api/v1"
	"server/common/dao"
	"time"

	"gorm.io/plugin/soft_delete"
)

type Notebook struct {
	dao.Model
	Id               string                `gorm:"primaryKey;type:varchar(100);not null;comment:Id"`
	UserId           string                `gorm:"type:varchar(100);not null;index;uniqueIndex:name_userId_spaceId,priority:2;comment:用户Id"`
	WorkspaceId      string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:name_userId_spaceId,priority:3;comment:群组Id"`
	Name             string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:name_userId_spaceId,priority:1;comment:名称"`
	Desc             string                `gorm:"type:varchar(1024);not null;default:'';comment:描述"`
	ImageId          string                `gorm:"type:varchar(100);not null;default:'';comment:镜像Id"`
	ImageName        string                `gorm:"type:varchar(100);not null;default:'';comment:镜像名称"`
	ImageVersion     string                `gorm:"type:varchar(100);not null;default:'';comment:镜像版本"`
	AlgorithmId      string                `gorm:"type:varchar(100);not null;default:'';comment:算法Id"`
	AlgorithmVersion string                `gorm:"type:varchar(100);not null;default:'';comment:算法版本"`
	AlgorithmName    string                `gorm:"type:varchar(100);not null;default:'';comment:算法名称"`
	DatasetId        string                `gorm:"type:varchar(100);not null;default:'';comment:数据集Id"`
	DatasetVersion   string                `gorm:"type:varchar(100);not null;default:'';comment:数据集版本"`
	DatasetName      string                `gorm:"type:varchar(100);not null;default:'';comment:数据集名称"`
	ResourceSpecId   string                `gorm:"type:varchar(100);not null;default:'';comment:资源规格Id"`
	ResourceSpecName string                `gorm:"type:varchar(100);not null;default:'';comment:资源规格名称"`
	NotebookJobId    string                `gorm:"type:varchar(100);not null;index;comment:JobId"`
	Url              string                `gorm:"type:varchar(100);not null;default:'';comment:jpy uri"`
	Status           string                `gorm:"type:varchar(50);not null;default:'';comment:preparing/pending/running/stopped"`
	DeletedAt        soft_delete.DeletedAt `gorm:"uniqueIndex:name_userId_spaceId,priority:4"`
}

func (Notebook) TableName() string {
	return "notebook"
}

type NotebookJob struct {
	dao.Model
	Id                string                     `gorm:"type:varchar(100);not null;comment:Id"`
	NotebookId        string                     `gorm:"type:varchar(100);not null;index;comment:Notebook Id"`
	Status            string                     `gorm:"type:varchar(50);not null;default:'';comment:preparing/pending/running/stopped"`
	StartedAt         *time.Time                 `gorm:"type:datetime(3);comment:开始运行时间"`
	StoppedAt         *time.Time                 `gorm:"type:datetime(3);comment:结束运行时间"`
	PayAmount         float64                    `gorm:"type:decimal(10,2);not null;default:0;comment:计费机时"`
	PayStartedAt      *time.Time                 `gorm:"type:datetime(3);comment:计费起始时间"`
	PayEndedAt        *time.Time                 `gorm:"type:datetime(3);comment:计费截止时间"`
	PayStatus         api.BillingPayRecordStatus `gorm:"type:tinyint;not null;default:1;comment:计费状态 1计费中 2计费完成"`
	ResourceSpecPrice uint32                     `gorm:"type:int unsigned;not null;default:0;comment:资源规格价格"`
}

func (NotebookJob) TableName() string {
	return "notebook_job"
}

type NotebookQuery struct {
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
	Name         string
}

type NotebookJobQuery struct {
	PageIndex   int
	PageSize    int
	SortBy      string
	OrderBy     string
	StartedAtLt int64
	Status      string
	StatusList  []string
	PayStatus   api.BillingPayRecordStatus
	Ids         []string
}

type NotebookEvent struct {
	Timestamp string
	Name      string
	Reason    string
	Message   string
}

type NotebookEventQuery struct {
	PageIndex    int
	PageSize     int
	Id           string
	TaskIndex    int
	ReplicaIndex int
}
