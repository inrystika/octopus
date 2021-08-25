package model

import (
	"server/common/dao"
)

type AdminUser struct {
	dao.Model
	Id       string `gorm:"type:varchar(100);not null"`
	Username string `gorm:"type:varchar(100);not null;default:''"`
	Email    string `gorm:"type:varchar(100);not null;default:''"`
	Phone    string `gorm:"type:varchar(100);not null;default:''"`
	Password string `gorm:"type:varchar(255);not null;default:''"`
}

func (AdminUser) TableName() string {
	return "admin_user"
}

type AdminUserQuery struct {
	PageIndex int
	PageSize  int
	Username  string
}
