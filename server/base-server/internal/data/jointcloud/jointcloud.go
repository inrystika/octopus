package jointcloud

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"server/common/errors"
	"server/common/log"
	"server/common/utils"
	"strings"
	"time"

	"gopkg.in/resty.v1"
)

type JointCloud interface {
	ListDataSet(ctx context.Context, query *DataSetQuery) (*ListDataSetReply, error)
	ListDataSetVersion(ctx context.Context, query *DataSetVersionQuery) (*ListDataSetVersionReply, error)
	SubmitJob(ctx context.Context, params *JointcloudJobParam) (*SubmitJobReply, error)
	ListFramework(ctx context.Context) (*ListFrameworkReply, error)
	ListFrameworkVersion(ctx context.Context, key string) (*ListFrameworkVersionReply, error)
	ListInterpreter(ctx context.Context) (*ListInterpreterReply, error)
	ListInterpreterVersion(ctx context.Context, key string) (*ListInterpreterVersionReply, error)
	ListJob(ctx context.Context, query *JobQuery) (*ListJobReply, error)
}

func NewJointCloud(baseUrl, username, password string, sessionExpirySec int32) JointCloud {
	j := &jointCloud{
		baseUrl:  baseUrl,
		username: username,
		password: password,
		client:   resty.New(),
	}

	ticker := time.NewTicker(time.Second * time.Duration(sessionExpirySec))
	go utils.HandlePanic(context.TODO(), func(i ...interface{}) {
		for range ticker.C {
			err := j.login(context.TODO())
			if err != nil {
				j.client.Cookies = nil
			}
		}
	})()
	return j
}

func parseBody(ctx context.Context, reply *Reply, body interface{}) error {
	log.Info(ctx, "reply:", reply)
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
	err := j.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	r := &Reply{}
	_, err = j.client.R().SetResult(r).SetQueryParams(map[string]string{"query": "{}", "pager": getPager(query.PageIndex, query.PageSize)}).Get(j.baseUrl + "/v1/dataSet")
	if err != nil {
		return nil, err
	}

	reply := &ListDataSetReply{}
	err = parseBody(ctx, r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (j *jointCloud) ListDataSetVersion(ctx context.Context, query *DataSetVersionQuery) (*ListDataSetVersionReply, error) {
	err := j.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	r := &Reply{}
	_, err = j.client.R().SetResult(r).
		SetQueryParams(map[string]string{"query": "{}", "pager": getPager(query.PageIndex, query.PageSize)}).
		Get(fmt.Sprintf("%s/v1/dataSet/%s/version", j.baseUrl, query.DataSetCode))
	if err != nil {
		return nil, err
	}

	reply := &ListDataSetVersionReply{}
	err = parseBody(ctx, r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (j *jointCloud) ListFramework(ctx context.Context) (*ListFrameworkReply, error) {
	err := j.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	r := &Reply{}
	_, err = j.client.R().SetResult(r).
		Get(fmt.Sprintf("%s/v1/jointcloud/framework/type", j.baseUrl))
	if err != nil {
		return nil, err
	}

	reply := &ListFrameworkReply{}
	err = parseBody(ctx, r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (j *jointCloud) ListFrameworkVersion(ctx context.Context, key string) (*ListFrameworkVersionReply, error) {
	err := j.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	r := &Reply{}
	_, err = j.client.R().SetResult(r).
		Get(fmt.Sprintf("%s/v1/jointcloud/framework/%s/version", j.baseUrl, key))
	if err != nil {
		return nil, err
	}

	reply := &ListFrameworkVersionReply{}
	err = parseBody(ctx, r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (j *jointCloud) ListInterpreter(ctx context.Context) (*ListInterpreterReply, error) {
	err := j.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	r := &Reply{}
	_, err = j.client.R().SetResult(r).
		Get(fmt.Sprintf("%s/v1/jointcloud/interpreter/type", j.baseUrl))
	if err != nil {
		return nil, err
	}

	reply := &ListInterpreterReply{}
	err = parseBody(ctx, r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (j *jointCloud) ListInterpreterVersion(ctx context.Context, key string) (*ListInterpreterVersionReply, error) {
	err := j.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	r := &Reply{}
	_, err = j.client.R().SetResult(r).
		Get(fmt.Sprintf("%s/v1/jointcloud/interpreter/%s/version", j.baseUrl, key))
	if err != nil {
		return nil, err
	}

	reply := &ListInterpreterVersionReply{}
	err = parseBody(ctx, r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (j *jointCloud) login(ctx context.Context) error {
	r := &Reply{}
	_, err := j.client.R().SetResult(r).SetQueryParams(map[string]string{"username": j.username, "password": j.password}).Post(j.baseUrl + "/auth/login")
	if err != nil {
		return err
	}

	reply := &LoginReply{}
	err = parseBody(ctx, r, reply)
	if err != nil {
		return err
	}

	j.client.SetCookie(&http.Cookie{Name: "SESSION", Value: base64.StdEncoding.EncodeToString([]byte(reply.SessionId))})
	return nil
}

func (j *jointCloud) checkLogin(ctx context.Context) error {
	if len(j.client.Cookies) > 0 {
		return nil
	}

	err := j.login(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (j *jointCloud) SubmitJob(ctx context.Context, params *JointcloudJobParam) (*SubmitJobReply, error) {
	err := j.checkLogin(ctx)
	if err != nil {
		return nil, err
	}
	r := &Reply{}
	paraBytes, err := json.Marshal(params)
	_, err = j.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(paraBytes).SetResult(r).
		Post(fmt.Sprintf("%s/v1/jointcloud/task/create", j.baseUrl))

	if err != nil {
		return nil, err
	}
	reply := &SubmitJobReply{}
	err = parseBody(ctx, r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (j *jointCloud) ListJob(ctx context.Context, query *JobQuery) (*ListJobReply, error) {

	err := j.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	IdsQuery := fmt.Sprintf(`{"query": {"jobIds":%v}}`, query.Ids)

	r := &Reply{}
	_, err = j.client.R().SetResult(r).SetQueryParams(map[string]string{"query": IdsQuery, "pager": getPager(query.PageIndex, query.PageSize)}).Get(j.baseUrl + "/v1/jointcloud/getTaskListByJobIds")
	if err != nil {
		return nil, err
	}

	reply := &ListJobReply{}
	err = parseBody(ctx, r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
