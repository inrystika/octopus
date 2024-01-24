package resources

import (
	"server/common/dao"
)

type ResourceSpec struct {
	dao.Model
	Id               string  `gorm:"primaryKey;type:varchar(100);not null"`
	Name             string  `gorm:"uniqueIndex;type:varchar(100);not null;default:''"`
	Price            float64 `gorm:"type:decimal(10,2)"`
	ResourceQuantity string  `gorm:"type:text"`
}

func (ResourceSpec) TableName() string {
	return "resourcespec"
}

type CreateResourceSpecRequest struct {
	Name             string
	Price            float64
	ResourceQuantity string
}

type UpdateResourceSpecRequest struct {
	Id               string
	Name             string
	Price            float64
	ResourceQuantity string
}
