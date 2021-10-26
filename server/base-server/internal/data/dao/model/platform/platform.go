package platform

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"server/common/dao"

	"gorm.io/plugin/soft_delete"
)

type Platform struct {
	dao.Model
	Id           string                `gorm:"primaryKey;type:varchar(100);not null;comment:Id"`
	Name         string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:name_deleteAt,priority:1;comment:名称"`
	ClientSecret string                `gorm:"type:varchar(100);not null;default:'';comment:客户端Secret"`
	ContactName  string                `gorm:"type:varchar(100);not null;default:'';comment:联系人姓名"`
	ContactInfo  string                `gorm:"type:varchar(100);not null;default:'';comment:联系方式"`
	ResourcePool string                `gorm:"type:varchar(100);not null;default:'';comment:资源池"`
	DeletedAt    soft_delete.DeletedAt `gorm:"uniqueIndex:name_deleteAt,priority:2"`
}

func (Platform) TableName() string {
	return "platform"
}

type PlatformQuery struct {
	PageIndex    int
	PageSize     int
	SortBy       string
	OrderBy      string
	CreatedAtGte int64
	CreatedAtLt  int64
	SearchKey    string
	Ids          []string
	Name         string
}

type PlatformStorageConfig struct {
	dao.Model
	PlatformId string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:platformId_name_deleteAt,priority:1;comment:平台id"`
	Name       string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:platformId_name_deleteAt,priority:2;comment:名称"`
	Type       string                `gorm:"type:varchar(100);not null;default:'';comment:存储类型"`
	Options    *StorageOptions       `gorm:"type:json;comment:'存储配置'"`
	DeletedAt  soft_delete.DeletedAt `gorm:"uniqueIndex:platformId_name_deleteAt,priority:3"`
}

func (PlatformStorageConfig) TableName() string {
	return "platform_storage_config"
}

type Juicefs struct {
	Name    string `json:"name"`
	MetaUrl string `json:"metaUrl"`
}

type StorageOptions struct {
	Juicefs *Juicefs `json:"juicefs,omitempty"`
}

func (r StorageOptions) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *StorageOptions) Scan(input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}

type PlatformStorageConfigQuery struct {
	PageIndex    int
	PageSize     int
	SortBy       string
	OrderBy      string
	CreatedAtGte int64
	CreatedAtLt  int64
	SearchKey    string
	Ids          []string
	Name         string
	PlatformId   string
}

type PlatformConfig struct {
	dao.Model
	PlatformId string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:platformId_key_deleteAt,priority:1;comment:平台id"`
	Key        string                `gorm:"type:varchar(100);not null;default:'';uniqueIndex:platformId_key_deleteAt,priority:2;comment:key"`
	Value      string                `gorm:"type:text;comment:value"`
	DeletedAt  soft_delete.DeletedAt `gorm:"uniqueIndex:platformId_key_deleteAt,priority:3"`
}

func (PlatformConfig) TableName() string {
	return "platform_config"
}
