package model

import (
	"fmt"
	"gorm.io/plugin/soft_delete"
	"server/common/dao"

	"gorm.io/gorm"
)

type Workspace struct {
	dao.Model
	Id      string  `gorm:"type:varchar(100);not null;primaryKey;comment:'空间id'"`
	Name    string  `gorm:"type:varchar(100);not null;uniqueIndex:name_deletedAt,priority:1;comment:'空间名'"`
	RPoolId string  `gorm:"type:varchar(100);not null;index;comment:'资源池id'"`
	Users   []*User `gorm:"many2many:workspace_user;"`
	DeletedAt  soft_delete.DeletedAt `gorm:"uniqueIndex:name_deletedAt,priority:2"`

}

func (Workspace) TableName() string {
	return "workspace"
}

type WorkspaceUser struct {
	dao.Model
	UserId      string `gorm:"type:varchar(100);not null;uniqueIndex:userId_workspaceId"`
	WorkspaceId string `gorm:"type:varchar(100);not null;uniqueIndex:userId_workspaceId"`
}

func (WorkspaceUser) TableName() string {
	return "workspace_user"
}

// ******** params ***********

type WorkspaceList struct {
	SortBy    string
	OrderBy   string
	PageIndex uint32
	PageSize  uint32
	Name      string
	RPoolId   string
	SearchKey string
}

func (w WorkspaceList) Where(db *gorm.DB) *gorm.DB {
	listSql := "1 = 1"
	params := make([]interface{}, 0)
	if w.Name != "" {
		listSql += " and name  = ? "
		params = append(params, w.Name)
	}
	if w.RPoolId != "" {
		listSql += " and r_pool_id = ? "
		params = append(params, w.RPoolId)
	}
	return db.Where(listSql, params...)
}

func (w WorkspaceList) Or(db *gorm.DB) *gorm.DB {
	if w.SearchKey != "" {
		searchKeyLike := "%" + w.SearchKey + "%"
		db = db.Where("name like ?", searchKeyLike)
	}
	return db
}

func (w WorkspaceList) Order(db *gorm.DB) *gorm.DB {
	var orderBy, sortBy string
	if w.OrderBy != "" {
		orderBy = w.OrderBy
	} else {
		orderBy = "desc"
	}
	if w.SortBy != "" {
		sortBy = w.SortBy
	} else {
		sortBy = "created_at"
	}
	return db.Order(fmt.Sprintf("%v %v", sortBy, orderBy))
}

func (w WorkspaceList) Pagination(db *gorm.DB) *gorm.DB {
	var pageIndex, pageSize int
	if w.PageIndex <= 0 {
		pageIndex = 1
	} else {
		pageIndex = int(w.PageIndex)
	}
	if w.PageSize <= 0 {
		pageSize = 10
	} else {
		pageSize = int(w.PageSize)
	}
	db = db.Limit(pageSize).Offset((pageIndex - 1) * pageSize)
	return db
}

type WorkspaceQuery struct {
	Id string
}

type WorkspaceAdd struct {
	Id      string
	Name    string
	RPoolId string
}

type WorkspaceUpdate struct {
	Id      string
	Name    string
	RPoolId string
}

type WorkspaceDelete struct {
	Id string
}

type WorkspaceUserBatchAdd struct {
	UserIds     []string
	WorkspaceId string
}

type WorkspaceUserBatchDel struct {
	UserIds     []string
	WorkspaceId string
}

type WorkspaceUserList struct {
	WorkspaceId string
}

type UserWorkspaceList struct {
	UserId string
}

type WorkspaceListIn struct {
	Ids []string
}
