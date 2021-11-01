package jointcloud

import (
	"encoding/json"

	"gopkg.in/resty.v1"
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

const (
	success = "0"
)

type pager struct {
	Page string `json:"page"`
	Size string `json:"size"`
}

type ListDataSetReply struct {
	List []*struct {
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

type ListDataSetVersionReply struct {
	List []*struct {
		Version string `json:"version"`
		Remark  string `json:"remark"`
	} `json:"list"`
}
