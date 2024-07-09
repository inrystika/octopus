package model

import (
	"gorm.io/plugin/soft_delete"
	"server/common/dao"
)

type UserEndpoint struct {
	dao.Model
	Id        uint                  `gorm:"primarykey"`
	Endpoint  string                `gorm:"type:varchar(100);default:'';uniqueIndex:name,priority:1;comment:endpoint"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:name,priority:2"`
}

func (UserEndpoint) TableName() string {
	return "user_endpoint"
}
