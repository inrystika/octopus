package model

import (
	"server/common/dao"

	"gorm.io/plugin/soft_delete"
)

type DatasetType struct {
	dao.Model
	Id         string                `gorm:"primaryKey;type:varchar(100);not null;default:'';comment:Id"`
	Desc       string                `gorm:"type:varchar(256);not null;default:'';uniqueIndex:desc_deletedAt;comment:类型描述"`
	ReferTimes int                   `gorm:"type:int;not null;default:0;comment:类型引用次数"`
	DeletedAt  soft_delete.DeletedAt `gorm:"uniqueIndex:desc_deletedAt"`
}

func (DatasetType) TableName() string {
	return "dataset_type"
}

type Dataset struct {
	dao.Model
	Id         string                `gorm:"primaryKey;type:varchar(100);not null;comment:Id"`
	SpaceId    string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:name_userId_spaceId,priority:3;comment:归属群组Id"`
	UserId     string                `gorm:"type:varchar(100);not null;default:'';index;uniqueIndex:name_userId_spaceId,priority:2;comment:归属用户Id"`
	SourceType int                   `gorm:"type:tinyint;not null;default:0;comment:1预置数据集 2用户数据集"`
	Name       string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:name_userId_spaceId,priority:1;comment:名称"`
	TypeId     string                `gorm:"type:varchar(100);not null;default:'';comment:数据类型"`
	Desc       string                `gorm:"type:varchar(1024);not null;default:'';comment:描述"`
	DeletedAt  soft_delete.DeletedAt `gorm:"uniqueIndex:name_userId_spaceId,priority:4"`
}

func (Dataset) TableName() string {
	return "dataset"
}

type DatasetVersion struct {
	dao.Model
	DatasetId    string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:datasetId_version;comment:数据集Id"`
	Version      string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:datasetId_version;comment:版本"`
	VersionInt   int64                 `gorm:"type:int;not null;default:0;comment:版本int"`
	Desc         string                `gorm:"type:varchar(1024);not null;default:'';comment:描述"`
	Status       int                   `gorm:"type:tinyint;not null;default:0;comment:1初始 2解压中 3解压成功 4解压失败"`
	Path         string                `gorm:"type:varchar(200);not null;default:'';comment:存储路径"`
	OriginalPath string                `gorm:"type:varchar(200);not null;default:'';comment:原始文件路径"`
	DeletedAt    soft_delete.DeletedAt `gorm:"uniqueIndex:datasetId_version"`
}

func (DatasetVersion) TableName() string {
	return "dataset_version"
}

type DatasetAccess struct {
	dao.Model
	DatasetId string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:spaceId_datasetId;comment:数据集Id"`
	SpaceId   string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:spaceId_datasetId;comment:可见群组Id"`
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:spaceId_datasetId"`
}

func (DatasetAccess) TableName() string {
	return "dataset_access"
}

type DatasetVersionAccess struct {
	dao.Model
	DatasetId  string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:spaceId_datasetId_version;comment:数据集Id"`
	Version    string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:spaceId_datasetId_version;comment:版本"`
	VersionInt int64                 `gorm:"type:int;not null:default:0;comment:版本int"`
	SpaceId    string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:spaceId_datasetId_version;comment:可见群组Id"`
	DeletedAt  soft_delete.DeletedAt `gorm:"uniqueIndex:spaceId_datasetId_version"`
}

func (DatasetVersionAccess) TableName() string {
	return "dataset_version_access"
}

type DatasetTypeQuery struct {
	PageIndex int
	PageSize  int
}

type DatasetQuery struct {
	PageIndex    int
	PageSize     int
	SortBy       string
	OrderBy      string
	CreatedAtGte int64
	CreatedAtLt  int64
	SearchKey    string
	UserId       string
	SpaceId      string
	SourceType   int
	Ids          []string
	Name         string
	NameLike     string
}

type CommDatasetQuery struct {
	PageIndex    int
	PageSize     int
	SortBy       string
	OrderBy      string
	CreatedAtGte int64
	CreatedAtLt  int64
	SearchKey    string
	NameLike     string
	UserId       string
	SpaceId      string
	ShareSpaceId string
	SourceType   int
	Ids          []string
}

type DatasetVersionId struct {
	DatasetId string
	Version   string
}

type DatasetVersionQuery struct {
	PageIndex int
	PageSize  int
	SortBy    string
	OrderBy   string
	DatasetId string
	Ids       []DatasetVersionId
	Status    int
}

type CommDatasetVersionQuery struct {
	PageIndex    int
	PageSize     int
	SortBy       string
	OrderBy      string
	DatasetId    string
	ShareSpaceId string
	Ids          []DatasetVersionId
	Status       int
}

type DatasetVersionDelete struct {
	DatasetId string
	Version   string
}

type DatasetAccessQuery struct {
	DatasetId string
}

type DatasetAccessDelete struct {
	DatasetId string
	SpaceId   string
}

type DatasetVersionAccessQuery struct {
	DatasetId string
	Version   string
	SpaceId   string
}

type DatasetVersionAccessDelete struct {
	DatasetId string
	Version   string
	SpaceId   string
}

type DatasetAccessId struct {
	DatasetId string
	SpaceId   string
}
