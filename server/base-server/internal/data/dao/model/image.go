package model

import (
	"fmt"
	"server/common/dao"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// 镜像表
type Image struct {
	Id             string         `gorm:"type:varchar(100);not null;default:'';comment:'镜像Id';primaryKey"`
	ImageName      string         `gorm:"type:varchar(100);not null;default:'';comment:'镜像名称';uniqueIndex:in_iv_it_sid_uid_da"`
	ImageType      int32          `gorm:"type:int;not null;default:0;comment:'镜像类型';uniqueIndex:in_iv_it_sid_uid_da"`
	ImageDesc      string         `gorm:"type:text;comment:'镜像描述'"`
	ImageAddr      string         `gorm:"type:varchar(255);comment:'镜像地址'"`
	ImageVersion   string         `gorm:"type:varchar(100);not null;default:'';comment:'镜像版本号';uniqueIndex:in_iv_it_sid_uid_da"`
	SourceType     int32          `gorm:"type:int;not null;default:0;comment:'来源类型（文件上传，远程镜像）'"`
	SourceFilePath string         `gorm:"type:varchar(255);not null;default:'';comment:'来源文件存储路径'"`
	SpaceId        string         `gorm:"type:varchar(100);not null;default:'';comment:'归属群组Id;index:spaceId_userId;uniqueIndex:in_iv_it_sid_uid_da"`
	UserId         string         `gorm:"type:varchar(100);not null;default:'';comment:'归属用户';index:spaceId_userId;uniqueIndex:in_iv_it_sid_uid_da"`
	IsPrefab       int32          `gorm:"not null;default:0;comment:'镜像是否为预置镜像（1:是预置镜像，2:不是预置镜像）'"`
	Status         int32          `gorm:"not null;default:0;comment:'镜像状态（1:未制作 2:制作中 3.制作完成）'"`
	Accesses       []*ImageAccess `gorm:"foreignKey:ImageId"`
	dao.Model
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:in_iv_it_sid_uid_da"`
}

func (Image) TableName() string {
	return "image"
}

// 镜像可见表
type ImageAccess struct {
	Id      string `gorm:"type:varchar(100);not null;default:'';comment:'可见镜像Id';primaryKey"`
	ImageId string `gorm:"type:varchar(100);not null;default:'';comment:'镜像Id';uniqueIndex:spaceId_userId_imageId_deletedAt"`
	SpaceId string `gorm:"not null;default:0;comment:'可见群组Id';uniqueIndex:spaceId_userId_imageId_deletedAt"`
	UserId  string `gorm:"type:varchar(100);not null;default:'';comment:'归属用户';index:spaceId_userId_imageId_deletedAt"`
	Image   *Image `gorm:"foreignKey:ImageId"`
	dao.Model
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:spaceId_userId_imageId_deletedAt"`
}

func (ImageAccess) TableName() string {
	return "image_access"
}

type ImageList struct {
	SortBy        string
	OrderBy       string
	PageIndex     uint32
	PageSize      uint32
	ImageNameLike string
	ImageName     string
	ImageVersion  string
	IsPrefab      int32
	Status        int32
	UserId        string
	SpaceId       string
	ImageType     int32
	SourceType    int32
	SearchKey     string
}

func (i ImageList) Where(db *gorm.DB) *gorm.DB {
	querySql := "1 = 1"
	params := make([]interface{}, 0)

	if i.ImageNameLike != "" {
		querySql += " and image_name like ?"
		params = append(params, i.ImageNameLike+"%")
	}

	if i.UserId != "" {
		querySql += " and user_id = ? "
		params = append(params, i.UserId)
	}

	if i.ImageName != "" {
		querySql += " and image_name = ? "
		params = append(params, i.ImageName)
	}

	if i.SpaceId != "" {
		querySql += " and space_id = ? "
		params = append(params, i.SpaceId)
	}

	if i.IsPrefab != 0 {
		querySql += " and is_prefab = ? "
		params = append(params, i.IsPrefab)
	}

	if i.ImageType != 0 {
		querySql += " and image_type = ? "
		params = append(params, i.ImageType)
	}

	if i.SourceType != 0 {
		querySql += " and source_type = ? "
		params = append(params, i.SourceType)
	}

	if i.Status != 0 {
		querySql += " and status = ? "
		params = append(params, i.Status)
	}

	if i.ImageVersion != "" {
		querySql += " and image_version = ? "
		params = append(params, i.ImageVersion)
	}

	return db.Where(querySql, params...)
}

func (i ImageList) Or(db *gorm.DB) *gorm.DB {
	if i.SearchKey != "" {
		searchKeyLike := "%" + i.SearchKey + "%"
		db = db.Where("image_name like ? or image_version like ? or image_desc like ?", searchKeyLike, searchKeyLike, searchKeyLike)
	}
	return db
}

func (i ImageList) Order(db *gorm.DB) *gorm.DB {
	var orderBy, sortBy string
	if i.OrderBy != "" {
		orderBy = i.OrderBy
	} else {
		orderBy = "desc"
	}
	if i.SortBy != "" {
		sortBy = i.SortBy
	} else {
		sortBy = "created_at"
	}

	db = db.Order(fmt.Sprintf("%v %v", sortBy, orderBy))
	return db
}

func (i ImageList) Pagination(db *gorm.DB) *gorm.DB {
	var pageIndex, pageSize int
	if i.PageIndex <= 0 {
		pageIndex = 1
	} else {
		pageIndex = int(i.PageIndex)
	}
	if i.PageSize <= 0 {
		pageSize = 10
	} else {
		pageSize = int(i.PageSize)
	}
	db = db.Limit(pageSize).Offset((pageIndex - 1) * pageSize)
	return db
}

type ImageQuery struct {
	Id string
}

type ImageDel struct {
	Id string
}

type ImageAdd struct {
	Id           string
	ImageName    string
	ImageVersion string
	ImageType    int32
	ImageDesc    string
	ImageAddr    string
	SourceType   int32
	SpaceId      string
	UserId       string
	IsPrefab     int32
	Status       int32
}

type ImageUpdate struct {
	ImageName      string
	ImageVersion   string
	ImageType      int32
	ImageDesc      string
	ImageAddr      string
	Status         int32
	SourceFilePath string
}

type ImageUpdateCond struct {
	Id string
}

type ImageAccessAdd struct {
	Id      string
	ImageId string
	SpaceId string
	UserId  string
}

type ImageAccessQuery struct {
	Id      string
	ImageId string
	SpaceId string
	UserId  string
}

type ImageAccessDel struct {
	Id      string
	ImageId string
	SpaceId string
	UserId  string
}

type ImageAccessList struct {
	SortBy        string
	OrderBy       string
	PageIndex     uint32
	PageSize      uint32
	ImageNameLike string
	ImageName     string
	ImageVersion  string
	IsPrefab      int32
	Status        int32
	UserId        string
	SpaceId       string
	ImageType     int32
	SourceType    int32
	SearchKey     string
}

func (i ImageAccessList) JoinImageAccess(db *gorm.DB) *gorm.DB {
	//joinStr := "LEFT JOIN image on image_access.image_id = image.id"
	//db = db.Joins(joinStr)

	joinSql := "INNER JOIN image_access ON image.id = image_access.image_id"
	querySql := "image_access.deleted_at = 0"
	params := make([]interface{}, 0)
	if i.ImageNameLike != "" {
		querySql += " and image.image_name like ?"
		params = append(params, i.ImageNameLike+"%")
	}

	if i.UserId != "" {
		querySql += " and image.user_id = ? "
		params = append(params, i.UserId)
	}

	if i.ImageName != "" {
		querySql += " and image.image_name = ? "
		params = append(params, i.ImageName)
	}

	if i.SpaceId != "" {
		querySql += " and image.space_id = ? "
		params = append(params, i.SpaceId)
	}

	if i.IsPrefab != 0 {
		querySql += " and image.is_prefab = ? "
		params = append(params, i.IsPrefab)
	}

	if i.ImageType != 0 {
		querySql += " and image.image_type = ? "
		params = append(params, i.ImageType)
	}

	if i.SourceType != 0 {
		querySql += " and image.source_type = ? "
		params = append(params, i.SourceType)
	}

	if i.Status != 0 {
		querySql += " and image.status = ? "
		params = append(params, i.Status)
	}

	if i.ImageVersion != "" {
		querySql += " and image.image_version = ? "
		params = append(params, i.ImageVersion)
	}

	if i.SearchKey != "" {
		querySql += " and image.image_name like ? or image.image_version like ? or image.image_desc like ?"
		searchKeyLike := "%" + i.SearchKey + "%"
		params = append(params, searchKeyLike, searchKeyLike, searchKeyLike)
	}

	return db.Joins(joinSql).Where(querySql, params...)
}

func (i ImageAccessList) JoinImage(db *gorm.DB) *gorm.DB {
	//joinStr := "LEFT JOIN image on image_access.image_id = image.id"
	//db = db.Joins(joinStr)

	joinSql := "INNER JOIN image ON image.id = image_access.image_id"
	querySql := "image.deleted_at = 0"
	params := make([]interface{}, 0)
	if i.ImageNameLike != "" {
		querySql += " and image.image_name like ?"
		params = append(params, i.ImageNameLike+"%")
	}

	if i.UserId != "" {
		querySql += " and image.user_id = ? "
		params = append(params, i.UserId)
	}

	if i.ImageName != "" {
		querySql += " and image.image_name = ? "
		params = append(params, i.ImageName)
	}

	if i.SpaceId != "" {
		querySql += " and image.space_id = ? "
		params = append(params, i.SpaceId)
	}

	if i.IsPrefab != 0 {
		querySql += " and image.is_prefab = ? "
		params = append(params, i.IsPrefab)
	}

	if i.ImageType != 0 {
		querySql += " and image.image_type = ? "
		params = append(params, i.ImageType)
	}

	if i.SourceType != 0 {
		querySql += " and image.source_type = ? "
		params = append(params, i.SourceType)
	}

	if i.Status != 0 {
		querySql += " and image.status = ? "
		params = append(params, i.Status)
	}

	if i.ImageVersion != "" {
		querySql += " and image.image_version = ? "
		params = append(params, i.ImageVersion)
	}

	if i.SearchKey != "" {
		querySql += " and image.image_name like ? or image.image_version like ? or image.image_desc like ?"
		searchKeyLike := "%" + i.SearchKey + "%"
		params = append(params, searchKeyLike, searchKeyLike, searchKeyLike)
	}

	return db.Joins(joinSql).Where(querySql, params...)
}

func (i ImageAccessList) Order(db *gorm.DB) *gorm.DB {
	var orderBy, sortBy string
	if i.OrderBy != "" {
		orderBy = i.OrderBy
	} else {
		orderBy = "desc"
	}
	if i.SortBy != "" {
		sortBy = i.SortBy
	} else {
		sortBy = "image_access.created_at"
	}

	db = db.Order(fmt.Sprintf("%v %v", sortBy, orderBy))
	return db
}

func (i ImageAccessList) Pagination(db *gorm.DB) *gorm.DB {
	var pageIndex, pageSize int
	if i.PageIndex <= 0 {
		pageIndex = 1
	} else {
		pageIndex = int(i.PageIndex)
	}
	if i.PageSize <= 0 {
		pageSize = 10
	} else {
		pageSize = int(i.PageSize)
	}
	db = db.Limit(pageSize).Offset((pageIndex - 1) * pageSize)
	return db
}

type ImageListIn struct {
	Ids []string
}

type ImageAccessListIn struct {
	Ids     []string
	SpaceId string
	UserId  string
}

func (i ImageAccessListIn) Where(db *gorm.DB) *gorm.DB {
	querySql := "1 = 1"
	params := make([]interface{}, 0)

	if len(i.Ids) > 0 {
		querySql += " AND image_id IN ? "
		params = append(params, i.Ids)
	}

	if i.UserId != "" {
		querySql += " AND user_id = ? "
		params = append(params, i.UserId)
	}

	if i.SpaceId != "" {
		querySql += " AND space_id = ? "
		params = append(params, i.SpaceId)
	}

	return db.Where(querySql, params...)
}
