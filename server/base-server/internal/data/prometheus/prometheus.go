package prometheus

import (
	"context"
	"fmt"
	"server/common/errors"
	"strconv"

	"gopkg.in/resty.v1"
)

type prometheus struct {
	client *resty.Client
	apiUrl string
}

type Reply struct {
	Status string `json:"status"`
	Data   struct {
		Result []struct {
			Values [][]interface{} `json:"values"`
		}
	}
}

type Prometheus interface {
	QueryCpuUsage(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error)
	QueryMemUsage(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error)
	QueryGpuUtil(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error)
	QueryGpuMemUtil(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error)
}

func NewPrometheus(baseUrl string) Prometheus {
	return &prometheus{
		apiUrl: baseUrl + "/api/v1",
		client: resty.New(),
	}
}

func (p *prometheus) QueryCpuUsage(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error) {
	query := fmt.Sprintf(`sum (rate (container_cpu_usage_seconds_total{pod="%s"}[1m]))`, podName)
	vals, err := p.query(query, start, size, step)
	for i, v := range vals {
		if v != -1 {
			vals[i] = v * 100
		}
	}
	return vals, err
}

func (p *prometheus) QueryMemUsage(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error) {
	query := fmt.Sprintf(`sum (container_memory_working_set_bytes{pod="%s"})`, podName)
	return p.query(query, start, size, step)
}

func (p *prometheus) QueryGpuUtil(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error) {
	query := fmt.Sprintf(`dcgm_gpu_utilization{pod_name="%s"}`, podName)
	return p.query(query, start, size, step)
}

func (p *prometheus) QueryGpuMemUtil(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error) {
	query := fmt.Sprintf(`dcgm_mem_copy_utilization{pod_name="%s"}`, podName)
	return p.query(query, start, size, step)
}

func (p *prometheus) query(query string, start int64, size int, step int) ([]float64, error) {
	r := &Reply{}
	params := map[string]string{
		"query": query,
		"start": strconv.FormatInt(start, 10),
		"end":   strconv.FormatInt(start+int64(step*size), 10),
		"step":  strconv.Itoa(step),
	}

	url := p.apiUrl + "/query_range"
	_, err := p.client.R().SetResult(r).SetQueryParams(params).Get(url)
	if err != nil || r.Status != "success" {
		return nil, errors.Errorf(err, errors.ErrorPrometheusQueryFailed)
	}

	m := make(map[int64]float64)
	if len(r.Data.Result) > 0 && len(r.Data.Result[0].Values) > 0 {
		vs := r.Data.Result[0].Values
		for _, v := range vs {
			val, err := strconv.ParseFloat(v[1].(string), 10)
			if err != nil {
				return nil, errors.Errorf(err, errors.ErrorInternal)
			}
			m[int64(v[0].(float64))] = val
		}
	}
	res := make([]float64, 0)
	for i := 0; i < size; i++ {
		val, exist := m[start+int64(step*i)]
		if exist {
			res = append(res, val)
		} else {
			res = append(res, -1)
		}
	}
	return res, nil
}
