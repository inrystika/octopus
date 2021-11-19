package model

import (
	"server/common/dao"

	"gorm.io/plugin/soft_delete"
)

type Lable struct {
	dao.Model
	Id             string                `gorm:"primaryKey;type:varchar(100);not null;default:'';comment:Id"`
	RelegationType int                   `gorm:"type:int;not null;default:0;comment:1数据集 2算法;uniqueIndex:relegationtype_labletype_desc_deletedAt;comment:标签所属模块"`
	SourceType     int                   `gorm:"type:int;not null;default:0;comment:1预置标签 2自定义标签"`
	LableType      int                   `gorm:"type:int;not null;default:0;uniqueIndex:relegationtype_labletype_desc_deletedAt;comment:标签类型"`
	LableDesc      string                `gorm:"type:varchar(256);not null;default:'';uniqueIndex:relegationtype_labletype_desc_deletedAt;comment:标签描述"`
	ReferTimes     int                   `gorm:"type:int;not null;default:0;comment:类型引用次数"`
	DeletedAt      soft_delete.DeletedAt `gorm:"uniqueIndex:relegationtype_labletype_desc_deletedAt"`
}

func (Lable) TableName() string {
	return "lable"
}

type LableListQuery struct {
	RelegationType int
	SourceType     int
	LableType      int
	PageIndex      int
	PageSize       int
}

type LableQuery struct {
	RelegationType int
	SourceType     int
	LableType      int
	LableDesc      string
}
