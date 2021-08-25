package resources

import (
	"server/common/dao"
)

type ResourceSpec struct {
	dao.Model
	Id               string `gorm:"primaryKey;type:varchar(100);not null"`
	Name             string `gorm:"uniqueIndex;type:varchar(100);not null;default:''"`
	Price            uint32 `gorm:"type:int unsigned"`
	ResourceQuantity string `gorm:"type:text"`
}

func (ResourceSpec) TableName() string {
	return "resourcespec"
}

type CreateResourceSpecRequest struct {
	Name             string
	Price            uint32
	ResourceQuantity string
}

type UpdateResourceSpecRequest struct {
	Price            uint32
	ResourceQuantity string
}
