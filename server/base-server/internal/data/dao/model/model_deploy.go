package model

import (
	api "server/base-server/api/v1"
	"server/common/dao"
	"time"

	"gorm.io/plugin/soft_delete"
)

func (ModelDeploy) TableName() string {
	return "model_deploy"
}

type ModelDeploy struct {
	Id             string                     `gorm:"primaryKey;type:varchar(100);not null;comment:'Id'"`
	UserId         string                     `gorm:"type:varchar(100);not null;index;uniqueIndex:name_userId_spaceId,priority:2;comment:'用户Id'"`
	WorkspaceId    string                     `gorm:"type:varchar(100);not null;default:'';uniqueIndex:name_userId_spaceId,priority:3;;comment:'群组Id'"`
	Name           string                     `gorm:"type:varchar(100);not null;default:'';uniqueIndex:name_userId_spaceId,priority:1;comment:'名称'"`
	Desc           string                     `gorm:"type:varchar(1024);not null;default:'';comment:'描述'"`
	ModelId        string                     `gorm:"type:varchar(100);not null;default:'';comment:'模型Id'"`
	ModelName      string                     `gorm:"type:varchar(100);not null;default:'';comment:'模型名称'"`
	ModelVersion   string                     `gorm:"type:varchar(100);not null;default:'';comment:'模型版本'"`
	ModelFrame     string                     `gorm:"type:varchar(100);not null;default:'';comment:'模型框架名称''"`
	ServiceUrl     string                     `gorm:"type:varchar(256);not null;default:'';comment:'服务url路径''"`
	Operation      string                     `gorm:"type:varchar(100);not null;default:''"`
	Status         string                     `gorm:"type:varchar(100);not null;comment:'preparing/pending/running/stopped/succeeded/failed'"`
	CompletedAt    *time.Time                 `gorm:"type:datetime(3);comment:'结束运行时间'"`
	StartedAt      *time.Time                 `gorm:"type:datetime(3);comment:'开始运行时间'"`
	PayAmount      float64                    `gorm:"type:decimal(10,2);not null;default:0;comment:结算机时"`
	PayStartedAt   *time.Time                 `gorm:"type:datetime(3);comment:计费起始时间"`
	PayEndedAt     *time.Time                 `gorm:"type:datetime(3);comment:计费截止时间"`
	PayStatus      api.BillingPayRecordStatus `gorm:"type:tinyint;not null;default:1;comment:扣费状态 1扣费中 2扣费完成"`
	ResourceSpecId string                     `gorm:"type:varchar(100);not null;default:'';comment:'资源Id'"`
	ResSpecPrice   float64                    `gorm:"type:decimal(10,2);comment:资源规格单价"`
	ResourcePool   string                     `gorm:"type:varchar(300);default:'';comment:资源池"`
	dao.Model
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:name_userId_spaceId,priority:4"`
}

type ModelDeployListQuery struct {
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
}

type DeployEventQuery struct {
	PageIndex int
	PageSize  int
	Id        string
	IsMain    bool
}

type ModelDeployEvent struct {
	Timestamp string
	Name      string
	Reason    string
	Message   string
}
