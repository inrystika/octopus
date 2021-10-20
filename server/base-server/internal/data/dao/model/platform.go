package model

import (
	"server/common/dao"

	"gorm.io/plugin/soft_delete"
)

type Platform struct {
	dao.Model
	Id           string                `gorm:"primaryKey;type:varchar(100);not null;comment:Id"`
	Name         string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:name_deleteAt,priority:1;comment:名称"`
	ClientSecret string                `gorm:"type:varchar(100);not null;default:'';comment:客户端Secret"`
	ContactName  string                `gorm:"type:varchar(100);not null;default:'';comment:联系人姓名"`
	ContactInfo  string                `gorm:"type:varchar(100);not null;default:'';comment:联系方式"`
	ResourcePool string                `gorm:"type:varchar(100);not null;default:'';comment:资源池"`
	DeletedAt    soft_delete.DeletedAt `gorm:"uniqueIndex:name_deleteAt,priority:2"`
}

func (Platform) TableName() string {
	return "platform"
}

type PlatformQuery struct {
	PageIndex    int
	PageSize     int
	SortBy       string
	OrderBy      string
	CreatedAtGte int64
	CreatedAtLt  int64
	SearchKey    string
	Ids          []string
	Name         string
}
