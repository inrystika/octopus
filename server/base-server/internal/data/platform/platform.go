package platform

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"server/common/errors"
	"strconv"
	"strings"
	"time"

	"gopkg.in/resty.v1"
)

type Platform interface {
	UpdateJobStatus(ctx context.Context, url string, req *JobStatusInfo) error
}

type platform struct {
	httpClient *resty.Client
}

func NewPlatform() Platform {
	return &platform{
		httpClient: resty.New(),
	}
}

func getMd5String(s string) string {
	b := []byte(s)
	return fmt.Sprintf("%x", md5.Sum(b))
}

func getSign(t int64) string {
	newKey := fmt.Sprintf("%s|%s", secretkey, t)
	return getMd5String(newKey)
}

func (p *platform) UpdateJobStatus(ctx context.Context, url string, req *JobStatusInfo) error {
	r := &Reply{}
	paraBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	sign := getSign(now)

	_, err = p.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("time", strconv.FormatInt(now, 10)).
		SetHeader("sign", sign).
		SetBody(paraBytes).SetResult(r).
		Post(url)

	if err != nil {
		return err
	}
	if !strings.EqualFold(r.Code, success) {
		return errors.Errorf(nil, errors.ErrorPlatformRequestFail)
	}
	return nil
}
