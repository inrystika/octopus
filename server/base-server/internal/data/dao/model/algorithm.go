package model

import (
	"fmt"
	"server/common/dao"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type AlgorithmType struct {
	dao.Model
	Id         string                `gorm:"primaryKey;type:varchar(100);not null;default:'';comment:Id"`
	Desc       string                `gorm:"type:varchar(256);not null;default:'';uniqueIndex:desc_deletedAt;comment:类型描述"`
	ReferTImes int                   `gorm:"type:int;not null;default:0;comment:类型引用次数"`
	DeletedAt  soft_delete.DeletedAt `gorm:"uniqueIndex:desc_deletedAt"`
}

func (AlgorithmType) TableName() string {
	return "algorithm_type"
}

type AlgorithmFramework struct {
	dao.Model
	Id         string                `gorm:"primaryKey;type:varchar(100);not null;default:'';comment:Id"`
	Desc       string                `gorm:"type:varchar(256);not null;default:'';uniqueIndex:desc_deletedAt;comment:类型描述"`
	ReferTImes int                   `gorm:"type:int;not null;default:0;comment:类型引用次数"`
	DeletedAt  soft_delete.DeletedAt `gorm:"uniqueIndex:desc_deletedAt"`
}

func (AlgorithmFramework) TableName() string {
	return "algorithm_framework"
}

// 算法表
type Algorithm struct {
	AlgorithmId       string `gorm:"type:varchar(100);not null;default:'';comment:'算法Id';primaryKey"`
	AlgorithmName     string `gorm:"type:varchar(255);not null;default:'';comment:'算法名称';uniqueIndex:algorithmName_modelName_spaceId_userId_isPrefab_deletedAt"`
	ModelName         string `gorm:"type:varchar(255);not null;default:'';comment:'模型名称';uniqueIndex:algorithmName_modelName_spaceId_userId_isPrefab_deletedAt"`
	SpaceId           string `gorm:"type:varchar(100);not null;default:'';comment:'归属群组Id';uniqueIndex:algorithmName_modelName_spaceId_userId_isPrefab_deletedAt"`
	UserId            string `gorm:"type:varchar(100);not null;default:'';comment:'归属用户';uniqueIndex:algorithmName_modelName_spaceId_userId_isPrefab_deletedAt"`
	IsPrefab          bool   `gorm:"not null;default:false;comment:'算法是否为预置算法（false默认算法，true预置算法）';uniqueIndex:algorithmName_modelName_spaceId_userId_isPrefab_deletedAt"`
	LatestVersion     string `gorm:"type:varchar(100);not null;default:'';comment:'最新版本'"`
	DataVersion       int64  `gorm:"not null;default:0;comment:'数据版本号，乐观锁'"`
	AlgorithmDescript string `gorm:"type:varchar(1024);not null;default:'';comment:'算法描述'"`
	TypeId            string `gorm:"type:varchar(100);not null;default:'';comment:算法类型"`
	FrameworkId       string `gorm:"type:varchar(100);not null;default:'';comment:算法框架"`
	dao.Model
	DeletedAt         soft_delete.DeletedAt `gorm:"uniqueIndex:algorithmName_modelName_spaceId_userId_isPrefab_deletedAt"`
	AlgorithmVersions []*AlgorithmVersion
}

func (Algorithm) TableName() string {
	return "algorithm"
}

// 算法版本表
type AlgorithmVersion struct {
	Id                string `gorm:"type:varchar(100);not null;default:'';comment:'算法版本Id';primaryKey"`
	AlgorithmId       string `gorm:"type:varchar(100);not null;default:'';comment:'算法Id';uniqueIndex:AlgorithmId_Version_deletedAt"`
	Version           string `gorm:"type:varchar(100);not null;default:'';comment:'算法版本';uniqueIndex:AlgorithmId_Version_deletedAt"`
	AlgorithmName     string `gorm:"type:varchar(255);not null;default:'';comment:'算法名称'"`
	FileStatus        uint8  `gorm:"not null;default:0;comment:'模型状态,1:初始态,2:文件上传中,3:文件已上传,4:文件上传失败'"`
	LatestCompressed  int64  `gorm:"not null;default:0;comment:'最后文件夹压缩时间，用于下载'"`
	DataVersion       int64  `gorm:"not null;default:0;comment:'数据版本号，乐观锁'"`
	AlgorithmDescript string `gorm:"type:varchar(1024);not null;default:'';comment:'算法描述'"`
	dao.Model
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:AlgorithmId_Version_deletedAt"`
}

func (AlgorithmVersion) TableName() string {
	return "algorithm_version"
}

// 算法可见表
type AlgorithmAccess struct {
	Id                     string `gorm:"type:varchar(100);not null;default:'';comment:'可见算法Id';primaryKey"`
	AlgorithmId            string `gorm:"type:varchar(100);not null;default:'';comment:'算法Id';uniqueIndex:algorithmId_spaceId_deleteAt"`
	AlgorithmName          string `gorm:"type:varchar(255);not null;default:'';comment:'算法名称'"`
	SpaceId                string `gorm:"not null;default:0;comment:'可见群组Id';uniqueIndex:algorithmId_spaceId_deleteAt"`
	LatestAlgorithmVersion string `gorm:"type:varchar(100);not null;default:'';comment:'最新可见版本'"`
	DataVersion            int64  `gorm:"not null;default:0;comment:'数据版本号，乐观锁'"`
	AlgorithmDescript      string `gorm:"type:varchar(1024);not null;default:'';comment:'算法描述'"`
	dao.Model
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:algorithmId_spaceId_deleteAt"`
}

func (AlgorithmAccess) TableName() string {
	return "algorithm_access"
}

// 算法可见版本表
type AlgorithmAccessVersion struct {
	Id                string `gorm:"type:varchar(100);not null;default:'';comment:'可见算法版本Id';primaryKey"`
	SpaceId           string `gorm:"not null;default:0;comment:'可见群组Id';uniqueIndex:spaceId_algorithmAccessId_algorithmVersion_algorithmId_deleteAt"`
	AlgorithmAccessId string `gorm:"type:varchar(100);not null;default:'';comment:'可见算法Id';uniqueIndex:spaceId_algorithmAccessId_algorithmVersion_algorithmId_deleteAt"`
	AlgorithmVersion  string `gorm:"type:varchar(100);not null;default:'';comment:'算法版本Id';uniqueIndex:spaceId_algorithmAccessId_algorithmVersion_algorithmId_deleteAt"`
	AlgorithmId       string `gorm:"type:varchar(100);not null;default:'';comment:'算法Id';uniqueIndex:spaceId_algorithmAccessId_algorithmVersion_algorithmId_deleteAt"`
	AlgorithmName     string `gorm:"type:varchar(255);not null;default:'';comment:'算法名称'"`
	DataVersion       int64  `gorm:"not null;default:0;comment:'数据版本号，乐观锁'"`
	dao.Model
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:spaceId_algorithmAccessId_algorithmVersion_algorithmId_deleteAt"`
}

func (AlgorithmAccessVersion) TableName() string {
	return "algorithm_access_version"
}

/* 访问接口 */
// AlgorithmType
type AlgorithmTypeQuery struct {
	PageIndex int
	PageSize  int
}

// AlgorithmFramework
type AlgorithmFrameworkQuery struct {
	PageIndex int
	PageSize  int
}

// Algorithm
type AlgorithmList struct {
	IsPrefab              bool
	SpaceId               string
	UserId                string
	AlgorithmId           string
	AlgorithmVersion      string
	SpaceIdOrder          bool
	SpaceIdSort           string
	UserIdOrder           bool
	UserIdSort            string
	AlgorithmIdOrder      bool
	AlgorithmIdSort       string
	AlgorithmVersionOrder bool
	AlgorithmVersionSort  string
	CreatedAtOrder        bool
	CreatedAtSort         string
	PageIndex             int
	PageSize              int
	SearchKey             string
	NameLike              string
	SortBy                string
	OrderBy               string
	FileStatus            int
	CreatedAtGte          int64
	CreatedAtLt           int64
}
type AlgorithmQuery struct {
	AlgorithmId string
}

type AlgorithmBatchQuery struct {
	List []*AlgorithmQuery
}

type AlgorithmQueryByInfo struct {
	AlgorithmName string
	UserId        string
	SpaceId       string
	IsPrefab      bool
}

type AlgorithmDelete struct {
	AlgorithmId string
	IsPrefab    bool
}

// AlgorithmVersion
type AlgorithmVersionList struct {
	AlgorithmId    string
	AlgorithmOrder bool
	VersionOrder   bool
	VersionSort    string
	CreatedAtOrder bool
	CreatedAtSort  string
	PageIndex      int
	PageSize       int
	FileStatus     int
}

type AlgorithmVersionBatchQuery struct {
	AlgorithmVersionList []*AlgorithmVersionQuery
}

type AlgorithmVersionBatchQueryById struct {
	AlgorithmVersionIdList []string
}
type AlgorithmVersionQuery struct {
	AlgorithmId string
	Version     string
}
type AlgorithmVersionQueryById struct {
	AlgorithmVersionId string
}
type AlgorithmVersionDelete struct {
	AlgorithmId      string
	AlgorithmVersion string
}
type AlgorithmVersionDeleteById struct {
	AlgorithmVersionId string
}

type AlgorithmVersionBatchDelete struct {
	AlgorithmId      string
	AlgorithmVersion []string
}

// AlgorithmAccess
type AlgorithmAccessList struct {
	SpaceId          string
	AlgorithmId      string
	SpaceIdOrder     bool
	SpaceIdSort      string
	AlgorithmIdOrder bool
	AlgorithmIdSort  string
	CreatedAtOrder   bool
	CreatedAtSort    string
	PageIndex        int
	PageSize         int
	AlgorithmVersion string
	SearchKey        string
	NameLike         string
	SortBy           string
	OrderBy          string
	FileStatus       int
	CreatedAtGte     int64
	CreatedAtLt      int64
}

type AlgorithmAccessQuery struct {
	SpaceId     string
	AlgorithmId string
}
type AlgorithmAccessQueryById struct {
	AlgorithmAccessId string
}
type AlgorithmAccessBatchQuery struct {
	List []*AlgorithmAccessQuery
}
type AlgorithmAccessBatchQueryById struct {
	AlgorithmAccessIdList []string
}
type AlgorithmAccessDelete struct {
	SpaceId     string
	AlgorithmId string
}
type AlgorithmAccessDeleteById struct {
	AlgorithmAccessId string
}

// AlgorithmAccessVersion
type AlgorithmAccessVersionList struct {
	AlgorithmAccessId      string
	SpaceId                string
	AlgorithmId            string
	AlgorithmVersion       string
	AlgorithmAccessIdOrder bool
	AlgorithmAccessIdSort  string
	AlgorithmIdOrder       bool
	AlgorithmIdSort        string
	AlgorithmVersionOrder  bool
	AlgorithmVersionSort   string
	CreatedAtOrder         bool
	CreatedAtSort          string
	PageIndex              int
	PageSize               int
	FileStatus             int
}
type AlgorithmAccessVersionQuery struct {
	AlgorithmAccessId string
	AlgorithmVersion  string
	SpaceId           string
}
type AlgorithmAccessVersionQueryById struct {
	AlgorithmAccessVersionId string
}

type AlgorithmAccessVersionBatchQuery struct {
	List []*AlgorithmAccessVersionQuery
}
type AlgorithmAccessVersionBatchQueryById struct {
	AlgorithmlAccessVersionIdList []string
}

type AlgorithmAccessVersionDelete struct {
	AlgorithmAccessId string
	AlgorithmVersion  string
}

type AlgorithmAccessVersionBatchDelete struct {
	AlgorithmAccessId string
	AlgorithmVersion  []string
}

type AlgorithmAccessVersionDeleteById struct {
	AlgorithmAccessVersionId string
}

func (a AlgorithmList) Order(db *gorm.DB) *gorm.DB {
	var orderBy, sortBy string
	if a.OrderBy != "" {
		orderBy = a.OrderBy
	} else {
		orderBy = "desc"
	}
	if a.SortBy != "" {
		sortBy = a.SortBy
	} else {
		sortBy = "created_at"
	}

	db = db.Order(fmt.Sprintf("%v %v", sortBy, orderBy))
	return db
}

func (a AlgorithmAccessList) Order(db *gorm.DB) *gorm.DB {
	var orderBy, sortBy string
	if a.OrderBy != "" {
		orderBy = a.OrderBy
	} else {
		orderBy = "desc"
	}
	if a.SortBy != "" {
		sortBy = a.SortBy
	} else {
		sortBy = "created_at"
	}

	db = db.Order(fmt.Sprintf("%v %v", sortBy, orderBy))
	return db
}
