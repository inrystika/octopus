package model

import (
	"gorm.io/gorm"
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

type AdminUserList struct {
	PageIndex int
	PageSize  int
	Username  string
}

type AdminUserQuery struct {
	Username  string
}

func (u AdminUserQuery) Where(db *gorm.DB) *gorm.DB {
	querySql := "1 = 1"
	params := make([]interface{}, 0)

	if u.Username != "" {
		querySql += " and username = ? "
		params = append(params, u.Username)
	}

	return db.Where(querySql, params...)
}
