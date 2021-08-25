package resources

import (
	"server/common/dao"
)

type Resource struct {
	dao.Model
	Id          string `gorm:"primaryKey;type:varchar(100);not null"`
	Name        string `gorm:"uniqueIndex;type:varchar(100);not null;default:''"`
	Desc        string `gorm:"type:text"`
	ResourceRef string `gorm:"type:varchar(100);not null;default:''"`
}

func (Resource) TableName() string {
	return "resource"
}

type CreateResourceRequest struct {
	Name        string
	Desc        string
	ResourceRef string
}
