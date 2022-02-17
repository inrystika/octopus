package model

import (
	"server/common/dao"

	"gorm.io/plugin/soft_delete"
)

const (
	FILESTATUS_INIT       uint8 = 0
	FILESTATUS_UPLOGADING uint8 = 1
	FILESTATUS_FINISH     uint8 = 2
	FILESTATUS_FAILED     uint8 = 3
)

// 模型表
type Model struct {
	Id               string `gorm:"type:varchar(100);not null;default:'';comment:'模型Id';primaryKey"`
	SpaceId          string `gorm:"type:varchar(100);not null;default:'';comment:'归属群组Id;uniqueIndex:spaceId_userId_algorithmId_algorithmVersion_isPrefab_deletedAt"`
	UserId           string `gorm:"type:varchar(100);not null;default:'';comment:'归属用户';uniqueIndex:spaceId_userId_algorithmId_algorithmVersion_isPrefab_deletedAt"`
	AlgorithmId      string `gorm:"type:varchar(100);not null;default:'';comment:'算法Id';uniqueIndex:spaceId_userId_algorithmId_algorithmVersion_isPrefab_deletedAt"`
	AlgorithmVersion string `gorm:"type:varchar(100);not null;default:'';comment:'算法版本';uniqueIndex:spaceId_userId_algorithmId_algorithmVersion_isPrefab_deletedAt"`
	IsPrefab         bool   `gorm:"not null;default:false;comment:'模型是否为预置模型（false默认模型，true预置模型）';uniqueIndex:spaceId_userId_algorithmId_algorithmVersion_isPrefab_deletedAt"`
	ModelName        string `gorm:"type:varchar(255);not null;default:'';comment:'模型名称'"`
	ModelDescript    string `gorm:"type:varchar(1024);not null;default:'';comment:'模型描述'"`
	LatestVersion    string `gorm:"type:varchar(100);not null;default:'';comment:'最新版本Id'"`
	DataVersion      int64  `gorm:"not null;default:0;comment:'数据版本号，乐观锁'"`
	dao.Model
	DeletedAt     soft_delete.DeletedAt `gorm:"uniqueIndex:spaceId_userId_algorithmId_algorithmVersion_isPrefab_deletedAt"`
	ModelVersions []*ModelVersion
}

func (Model) TableName() string {
	return "model"
}

// 模型版本表
type ModelVersion struct {
	Id          string `gorm:"type:varchar(100);not null;default:'';comment:'模型版本Id';primaryKey"`
	ModelId     string `gorm:"type:varchar(100);not null;default:'';comment:'模型Id';uniqueIndex:modelId_version_deletedAt"`
	Version     string `gorm:"type:varchar(100);not null;default:'';comment:'模型版本';uniqueIndex:modelId_version_deletedAt"`
	Descript    string `gorm:"type:varchar(1024);not null;default:'';comment:'版本描述'"`
	FileStatus  uint8  `gorm:"not null;default:0;comment:'模型状态,0:初始态,1:文件上传中,2:文件已上传,3:文件上传失败'"`
	DataVersion int64  `gorm:"not null;default:0;comment:'数据版本号，乐观锁'"`
	dao.Model
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:modelId_version_deletedAt"`
}

func (ModelVersion) TableName() string {
	return "model_version"
}

// 模型可见表
type ModelAccess struct {
	Id                 string `gorm:"type:varchar(100);not null;default:'';comment:'可见模型Id';primaryKey"`
	SpaceId            string `gorm:"not null;default:0;comment:'可见群组Id';uniqueIndex:spaceId_modelId_deletedAt"`
	ModelId            string `gorm:"type:varchar(100);not null;default:'';comment:'模型Id';uniqueIndex:spaceId_modelId_deletedAt"`
	LatestModelVersion string `gorm:"type:varchar(100);not null;default:'';comment:'最新可见版本'"`
	DataVersion        int64  `gorm:"not null;default:0;comment:'数据版本号，乐观锁'"`
	dao.Model
	DeletedAt           soft_delete.DeletedAt `gorm:"uniqueIndex:spaceId_modelId_deletedAt"`
	ModelVersionAccesss []*ModelVersionAccess
}

func (ModelAccess) TableName() string {
	return "model_access"
}

// 模型可见版本表
type ModelVersionAccess struct {
	Id            string `gorm:"type:varchar(100);not null;default:'';comment:'可见模型版本Id';primaryKey"`
	ModelAccessId string `gorm:"type:varchar(100);not null;default:'';comment:'可见模型Id';uniqueIndex:modelAccessId_modelVersion_deletedAt"`
	ModelVersion  string `gorm:"type:varchar(100);not null;default:'';comment:'模型版本Id';uniqueIndex:modelAccessId_modelVersion_deletedAt;index:modelId_modelVersion"`
	ModelId       string `gorm:"type:varchar(100);not null;default:'';comment:'模型Id';index:modelId_modelVersion"`
	DataVersion   int64  `gorm:"not null;default:0;comment:'数据版本号，乐观锁'"`
	dao.Model
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:modelAccessId_modelVersion_deletedAt"`
}

func (ModelVersionAccess) TableName() string {
	return "model_version_access"
}

const (
	DESC string = "desc"
	ASC  string = "asc"
)

/* 访问接口 */
// Model
type ModelList struct {
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
	SearchKey             string
	PageIndex             int
	PageSize              int
	Ids                   []string
	CreatedAtGte          int64
	CreatedAtLt           int64
}

// ModelVersion
type ModelVersionList struct {
	ModelId        string
	VersionOrder   bool
	VersionSort    string
	CreatedAtOrder bool
	CreatedAtSort  string
	PageIndex      int
	PageSize       int
	Versions       []string
}

// ModelAccess
type ModelAccessList struct {
	SpaceIds       []string
	ModelIds       []string
	SpaceIdOrder   bool
	SpaceIdSort    string
	ModelIdOrder   bool
	ModelIdSort    string
	CreatedAtOrder bool
	CreatedAtSort  string
	PageIndex      int
	PageSize       int
	Ids            []string
	CreatedAtGte   int64
	CreatedAtLt    int64
}

// ModelAccessVersion
type ModelVersionAccessList struct {
	ModelAccessId      string
	ModelId            string
	ModelAccessIdOrder bool
	ModelAccessIdSort  string
	ModelIdOrder       bool
	ModelIdSort        string
	ModelVersionOrder  bool
	ModelVersionSort   string
	CreatedAtOrder     bool
	CreatedAtSort      string
	PageIndex          int
	PageSize           int
	ModelVersions      []string
}
