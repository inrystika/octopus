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
	UpdateJobStatus(ctx context.Context, url, clientSecret string, req *JobStatusInfo) error
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

func getSign(clientSecret, time string) string {
	newKey := fmt.Sprintf("%s|%s", clientSecret, time)
	return getMd5String(newKey)
}

func (p *platform) UpdateJobStatus(ctx context.Context, url, clientSecret string, req *JobStatusInfo) error {
	r := &Reply{}
	paraBytes, err := json.Marshal(req)
	if err != nil {
		return err
	}

	now := strconv.FormatInt(time.Now().Unix(), 10)
	sign := getSign(clientSecret, now)

	_, err = p.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("time", now).
		SetHeader("sign", sign).
		SetBody(paraBytes).SetResult(r).
		Put(url)
	if err != nil {
		return err
	}
	if !strings.EqualFold(r.Code, success) {
		return errors.Errorf(nil, errors.ErrorPlatformRequestFail)
	}
	return nil
}
