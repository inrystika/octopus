package pipeline

import (
	"context"
	"encoding/json"
	"server/base-server/internal/conf"
	"server/common/errors"
	"strings"

	"gopkg.in/resty.v1"

	"server/common/log"
)

type Pipeline interface {
	SubmitJob(ctx context.Context, req *SubmitJobParam) (*SubmitJobReply, error)
	StopJob(ctx context.Context, req *UpdateJobParam) error
	UpsertFeature(ctx context.Context, req *UpsertFeatureParam) error
	//DeleteJob(ctx context.Context, req *BatchDeleteJobReq) (*DeleteJobResp, error)
	GetJobDetail(ctx context.Context, jobId string) (*JobStatusDetail, error)
	BatchGetJobDetail(ctx context.Context, jobIds []string) (*BatchJobStatusDetail, error)
}

type pipeline struct {
	log        *log.Helper
	baseUrl    string
	httpClient *resty.Client
	token      string
}

const (
	OPERATION_SUCCEEDED   = "OPERATION_SUCCEEDED"
	UNSUPPORTED_OPERATION = "UNSUPPORTED_OPERATION"
	headerToken           = "token"
)

func NewPipeline(confData *conf.Data, logger log.Logger) Pipeline {
	return &pipeline{
		log:        log.NewHelper("Pipeline", logger),
		baseUrl:    confData.Pipeline.BaseUrl,
		httpClient: resty.New(),
		token:      confData.Pipeline.Token,
	}
}

func parseBody(reply *Reply, body interface{}) error {

	if strings.EqualFold(reply.Code, UNSUPPORTED_OPERATION) {
		return nil
	}

	if !strings.EqualFold(reply.Code, OPERATION_SUCCEEDED) {
		return errors.Errorf(nil, errors.ErrorPipelineDoRequest)
	}

	if body != nil {
		err := json.Unmarshal(reply.Payload, body)
		if err != nil {
			return errors.Errorf(nil, errors.ErrorJsonUnmarshal)
		}
	}

	return nil
}

func (p *pipeline) SubmitJob(ctx context.Context, req *SubmitJobParam) (*SubmitJobReply, error) {
	r := &Reply{}
	_, err := p.httpClient.R().SetHeader(headerToken, p.token).SetBody(req).SetResult(r).Post(p.baseUrl + "/v1/job/")
	if err != nil {
		return nil, err
	}

	reply := &SubmitJobReply{}
	err = parseBody(r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (p *pipeline) StopJob(ctx context.Context, req *UpdateJobParam) error {
	r := &Reply{}
	_, err := p.httpClient.R().SetHeader(headerToken, p.token).SetBody(req).SetResult(r).Put(p.baseUrl + "/v1/job/stop/" + req.JobID)
	if err != nil {
		return err
	}

	err = parseBody(r, nil)
	if err != nil {
		p.log.Infof(ctx, "pipeline parseBody error:[%s, %v]", r.Code, r.Payload)
		return err
	}
	return nil
}

func (p *pipeline) UpsertFeature(ctx context.Context, req *UpsertFeatureParam) error {
	r := &Reply{}
	_, err := p.httpClient.R().SetHeader(headerToken, p.token).SetBody(req).SetResult(r).Post(p.baseUrl + "/v1/features/")
	if err != nil {
		return err
	}

	err = parseBody(r, nil)
	if err != nil {
		return err
	}
	return nil
}

func (p *pipeline) GetJobDetail(ctx context.Context, jobId string) (*JobStatusDetail, error) {
	r := &Reply{}
	_, err := p.httpClient.R().SetHeader(headerToken, p.token).SetResult(r).Get(p.baseUrl + "/v1/job/detail/" + jobId)
	if err != nil {
		return nil, err
	}

	reply := &JobStatusDetail{}
	err = parseBody(r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (p *pipeline) BatchGetJobDetail(ctx context.Context, jobIds []string) (*BatchJobStatusDetail, error) {
	r := &Reply{}
	_, err := p.httpClient.R().SetHeader(headerToken, p.token).SetResult(r).Get(p.baseUrl + "/v1/job/batchdetail?ids=" + strings.Join(jobIds, ","))
	if err != nil {
		return nil, err
	}

	reply := &BatchJobStatusDetail{}
	err = parseBody(r, reply)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
