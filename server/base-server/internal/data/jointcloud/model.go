package jointcloud

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"server/common/dao"

	"gopkg.in/resty.v1"
	"gorm.io/plugin/soft_delete"
)

type jointCloud struct {
	baseUrl  string
	username string
	password string
	client   *resty.Client
}

type Reply struct {
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

func (r Reply) String() string {
	return fmt.Sprintf("code:%v, message:%v, data:%v", r.Code, r.Message, string(r.Data))
}

const (
	success = "0"
)

type Pager struct {
	Page       int   `json:"page"`
	Size       int   `json:"size"`
	TotalPages int64 `json:"totalPages"`
	Total      int64 `json:"total"`
}

type ListDataSetReply struct {
	Pager *Pager `json:"pager"`
	List  []*struct {
		DataSetCode string `json:"dataSetCode"`
		Name        string `json:"name"`
		Remark      string `json:"remark"`
	} `json:"list"`
}

type DataSetQuery struct {
	PageIndex int
	PageSize  int
}

type LoginReply struct {
	SessionId string `json:"sessionId"`
}

type DataSetVersionQuery struct {
	PageIndex   int
	PageSize    int
	DataSetCode string
}

type StopJobRequest struct {
	TaskId string `json:"taskId"`
}

type ListDataSetVersionReply struct {
	Pager *Pager `json:"pager"`
	List  []*struct {
		Version string `json:"version"`
		Remark  string `json:"remark"`
	} `json:"list"`
}

type ListFrameworkReply struct {
	List []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"list"`
}

type ListFrameworkVersionReply struct {
	List []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"list"`
}

type ListInterpreterReply struct {
	List []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"list"`
}

type ListInterpreterVersionReply struct {
	List []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"list"`
}

type JointcloudJobParam struct {
	TaskName             string            `json:"taskName"`
	ExecCommand          string            `json:"execCommand"`
	Interpreter          string            `json:"interpreter"`
	Framework            string            `json:"framework"`
	Type                 string            `json:"type"`
	OutputPath           string            `json:"outputPath"`
	DataSetVersionVoList DataSetVersionVos `json:"dataSetVersionVoList"`
	Params               Parameters        `json:"params"`
	ResourceParams       Resources         `json:"resourceParams"`
	Remark               string            `json:"remark"`
}

type SubmitJobReply struct {
	TaskId string `json:"taskId"`
	JobId  string `json:"jobId"`
}

type DataSetVersionVo struct {
	DataSetName string `json:"dataSetName"`
	DataSetCode string `json:"dataSetCode"`
	Version     string `json:"version"`
	Path        string `json:"path"`
}

type Param struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ResourceParam struct {
	Name string  `json:"name"`
	Size float64 `json:"size"`
	Type string  `json:"type"`
	Unit string  `json:"unit"`
}

type DataSetVersionVos []*DataSetVersionVo
type Parameters []*Param
type Resources []*ResourceParam

type TrainJob struct {
	TaskId               string            `gorm:"primaryKey;type:varchar(100);not null;comment:'TaskId'"`
	JobId                string            `gorm:"primaryKey;type:varchar(100);not null;comment:'JobId'"`
	UserId               string            `gorm:"type:varchar(100);not null;index;uniqueIndex:taskName_userId_spaceId,priority:2;comment:'用户Id'"`
	WorkspaceId          string            `gorm:"type:varchar(100);not null;default:'';uniqueIndex:taskName_userId_spaceId,priority:3;;comment:'群组Id'"`
	TaskName             string            `gorm:"type:varchar(100);not null;default:'';uniqueIndex:taskName_userId_spaceId,priority:1;comment:'名称'"`
	Remark               string            `gorm:"type:varchar(1024);not null;default:'';comment:'描述'"`
	ExecCommand          string            `gorm:"type:varchar(1024);not null;default:'';comment:'执行命令'"`
	Interpreter          string            `gorm:"type:varchar(1024);not null;default:'';comment:'Interpreter'"`
	Framework            string            `gorm:"type:varchar(1024);not null;default:'';comment:'Framework'"`
	Type                 string            `gorm:"type:varchar(1024);not null;default:'';comment:'Type'"`
	OutputPath           string            `gorm:"type:varchar(1024);not null;default:'';comment:'容器输出路径'"`
	Operation            string            `gorm:"type:varchar(100);not null;default:''"`
	Status               string            `gorm:"type:varchar(100);not null;comment:'preparing/pending/running/stopped/succeeded/failed'"`
	DataSetVersionVoList DataSetVersionVos `gorm:"type:json;comment:'数据集版本信息'"`
	Params               Parameters        `gorm:"type:json;comment:'命令参数信息'"`
	ResourceParams       Resources         `gorm:"type:json;comment:'资源参数信息'"`
	CompletedAt          *time.Time        `gorm:"type:datetime(3);comment:'结束运行时间'"`
	StartedAt            *time.Time        `gorm:"type:datetime(3);comment:'开始运行时间'"`
	dao.Model
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:taskName_userId_spaceId,priority:4"`
}

type TrainJobListQuery struct {
	PageIndex    int
	PageSize     int
	SortBy       string
	OrderBy      string
	CreatedAtGte int64
	CreatedAtLt  int64
	Status       string
	SearchKey    string
	UserId       string
	WorkspaceId  string
	Ids          string
}

type ListJobReply struct {
	Pager *Pager      `json:"pager"`
	List  []*JobReply `json:"list"`
}

type JobReply struct {
	TaskId          string                   `json:"taskId"`
	TaskName        string                   `json:"taskName"`
	Interpreter     string                   `json:"interpreter"`
	Framework       string                   `json:"framework"`
	Type            string                   `json:"type"`
	CloudVendorId   int                      `json:"cloudVendorId"`
	CloudVendorName string                   `json:"cloudVendorName"`
	ImageUrl        string                   `json:"imageUrl"`
	ExecCommand     string                   `json:"execCommand"`
	OutputPath      string                   `json:"outputPath"`
	Status          int                      `json:"status"`
	Remark          string                   `json:"remark"`
	DataSetList     []*ReplyDataSetVersionVo `json:"dataSetList"`
	TaskParams      []*Param                 `json:"taskParams"`
	ResourceParams  []*ResourceParam         `json:"resourceParams"`
	CreateTime      string                   `json:"createTime"`
	Creator         int                      `json:"creator"`
	CreatorName     string                   `json:"creatorName"`
}

type JobQuery struct {
	PageIndex int
	PageSize  int
	Ids       string
}

type ReplyDataSetVersionVo struct {
	Category        string `json:"category"`
	DataSetCode     string `json:"dataSetCode"`
	Name            string `json:"name"`
	Remark          string `json:"remark"`
	CloudVendorId   int    `json:"cloudVendorId"`
	CloudVendorName string `json:"cloudVendorName"`
	VersionNumber   int    `json:"versionNumber"`
	Label           string `json:"label"`
	Creator         int    `json:"creator"`
	CreateTime      string `json:"createTime"`
}

func (TrainJob) TableName() string {
	return "jointcloud_train_job"
}

func (r DataSetVersionVos) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *DataSetVersionVos) Scan(input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}

func (r Parameters) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Parameters) Scan(input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}

func (r Resources) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Resources) Scan(input interface{}) error {
	switch v := input.(type) {
	case []byte:
		return json.Unmarshal(input.([]byte), r)
	default:
		return fmt.Errorf("cannot Scan() from: %#v", v)
	}
}
