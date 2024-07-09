package resources

import (
	"gorm.io/plugin/soft_delete"
	"server/common/dao"
)

type Resource struct {
	dao.Model
	Id          string                `gorm:"primaryKey;type:varchar(100);not null"`
	Name        string                `gorm:"uniqueIndex:name,priority:1;type:varchar(100);not null;default:''"`
	Desc        string                `gorm:"type:text"`
	ResourceRef string                `gorm:"type:varchar(100);not null;default:''"`
	DeletedAt   soft_delete.DeletedAt `gorm:"uniqueIndex:name,priority:2"`
}

func (Resource) TableName() string {
	return "resource"
}

type CreateResourceRequest struct {
	Name        string
	Desc        string
	ResourceRef string
}
