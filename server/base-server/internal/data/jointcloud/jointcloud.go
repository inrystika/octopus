package jointcloud

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"server/common/errors"
	"strings"

	"gopkg.in/resty.v1"
)

type JointCloud interface {
	ListDataSet(ctx context.Context, query *DataSetQuery) (*ListDataSetReply, error)
	ListDataSetVersion(ctx context.Context, query *DataSetVersionQuery) (*ListDataSetVersionReply, error)
}

func NewJointCloud(baseUrl, username, password string) JointCloud {
	j := &jointCloud{
		baseUrl:  baseUrl,
		username: username,
		password: password,
		client:   resty.New(),
	}
	return j
}

func parseBody(reply *Reply, body interface{}) error {
	if !strings.EqualFold(reply.Code, success) {
		return errors.Errorf(nil, errors.ErrorJointCloudRequestFailed)
	}

	if body != nil {
		err := json.Unmarshal(reply.Data, body)
		if err != nil {
			return errors.Errorf(nil, errors.ErrorJsonUnmarshal)
		}
	}

	return nil
}

func getPager(page, pageSize int) string {
	return fmt.Sprintf(`{"page":%v,"size":%v}`, page, pageSize)
}

func (j *jointCloud) ListDataSet(ctx context.Context, query *DataSetQuery) (*ListDataSetReply, error) {
	err := j.checkLogin()
	if err != nil {
		return nil, err
	}

	r := &Reply{}
	_, err = j.client.R().SetResult(r).SetQueryParams(map[string]string{"query": "{}", "pager": getPager(query.PageIndex, query.PageSize)}).Get(j.baseUrl + "/api/v1/dataSet")
	if err != nil {
		return nil, err
	}

	reply := &ListDataSetReply{}
	err = parseBody(r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (j *jointCloud) ListDataSetVersion(ctx context.Context, query *DataSetVersionQuery) (*ListDataSetVersionReply, error) {
	err := j.checkLogin()
	if err != nil {
		return nil, err
	}

	r := &Reply{}
	_, err = j.client.R().SetResult(r).
		SetQueryParams(map[string]string{"query": "{}", "pager": getPager(query.PageIndex, query.PageSize)}).
		Get(fmt.Sprintf("%s/api/v1/dataSet/%s/version", j.baseUrl, query.DataSetCode))
	if err != nil {
		return nil, err
	}

	reply := &ListDataSetVersionReply{}
	err = parseBody(r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (j *jointCloud) login() (*LoginReply, error) {
	r := &Reply{}
	_, err := j.client.R().SetResult(r).SetQueryParams(map[string]string{"username": j.username, "password": j.password}).Post(j.baseUrl + "/auth/login")
	if err != nil {
		return nil, err
	}

	reply := &LoginReply{}
	err = parseBody(r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (j *jointCloud) checkLogin() error {
	if len(j.client.Cookies) > 0 {
		return nil
	}

	reply, err := j.login()
	if err != nil {
		return err
	}

	j.client.SetCookie(&http.Cookie{Name: "SESSION", Value: base64.StdEncoding.EncodeToString([]byte(reply.SessionId))})
	return nil
}
