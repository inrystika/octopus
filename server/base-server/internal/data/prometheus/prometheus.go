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

	QueryAccCardUtil(ctx context.Context, podName, company string, start int64, size int, step int) ([]float64, error)
	QueryAccCardMemUtil(ctx context.Context, podName, company string, start int64, size int, step int) ([]float64, error)
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

func (p *prometheus) QueryAccCardUtil(ctx context.Context, podName, company string, start int64, size int, step int) ([]float64, error) {
	items := map[string]string{
		"nvidia":     "dcgm_gpu_utilization",                                 // GPU utilization
		"huawei":     "container_npu_utilization",                            // NPU utilization
		"cambricon":  "mlu_utilization * on(uuid) group_right mlu_container", // MLU utilization
		"enflame":    "enflame_gcu_usage",                                    // GCU utilization
		"iluvatar":   "ix_gpu_utilization",                                   // iluvatar GPU utilization
		"metax-tech": "",                                                     //
	}
	query := fmt.Sprintf(`%s{pod_name="%s"}`, items[company], podName)
	return p.query(query, start, size, step)
}

func (p *prometheus) QueryAccCardMemUtil(ctx context.Context, podName, company string, start int64, size int, step int) ([]float64, error) {
	items := map[string]string{
		"nvidia": "dcgm_mem_copy_utilization", // GPU memory utilization
		// "huawei":     "container_npu_used_memory / container_npu_total_memory",      // NPU utilization
		"cambricon":  "mlu_memory_utilization * on(uuid) group_right mlu_container", // MLU memory utilization
		"enflame":    "100 * enflame_gcu_memory_usage",                              // GCU memory utilization
		"iluvatar":   "ix_mem_utilization",                                          // iluvatar GPU memory utilization
		"metax-tech": "",                                                            //
	}
	if company == "huawei" {
		query := fmt.Sprintf(`100 * container_npu_used_memory{pod_name="%s"} / container_npu_total_memory{pod_name="%s"}`, podName, podName)
	} else {
		query := fmt.Sprintf(`%s{pod_name="%s"}`, items[company], podName)
	}
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
