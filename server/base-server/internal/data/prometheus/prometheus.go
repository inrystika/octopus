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

	QueryAccCardUtil(ctx context.Context, podName, company string, start int64, size int, step int) ([]float64, error)
	QueryAccCardMemUtil(ctx context.Context, podName, company string, start int64, size int, step int) ([]float64, error)

	QueryNetworkReceiveBytes(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error)
	QueryNetworkTransmitBytes(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error)
	QueryFSUsageBytes(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error)
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

func (p *prometheus) QueryAccCardUtil(
	ctx context.Context, podName, company string, start int64, size int, step int) ([]float64, error) {

	items := map[string]string{
		"huawei":     "container_npu_utilization", // NPU utilization
		"enflame":    "enflame_gcu_usage",         // GCU utilization
		"metax-tech": "",                          // "gopkg.in/resty.v1"
	}
	query := fmt.Sprintf(`%s{pod_name="%s"}`, items[company], podName)
	if company == "nvidia" {
		query = fmt.Sprint(`DCGM_FI_DEV_GPU_UTIL{pod="%s"}`, podName)
	}
	if company == "cambricon" {
		query = fmt.Sprintf(`mlu_utilization * on(uuid) group_right mlu_container{pod="%s"}`, podName) // MLU utilization
	}
	if company == "iluvatar" {
		query = fmt.Sprintf(`ix_gpu_utilization{pod="%s"}`, podName) // iluvatar GPU utilization
	}
	if company == "hygon" {
		query = fmt.Sprintf(`dcu_utilizationrate{dcu_pod_name="%s"}`, podName)
	}
	return p.query(query, start, size, step)
}

func (p *prometheus) QueryAccCardMemUtil(
	ctx context.Context, podName, company string, start int64, size int, step int) ([]float64, error) {

	items := map[string]string{
		"enflame":    "100 * enflame_gcu_memory_usage", // GCU memory utilization
		"metax-tech": "",                               //
	}
	query := fmt.Sprintf(`%s{pod_name="%s"}`, items[company], podName)
	if company == "nvidia" {
		query = fmt.Sprintf(`(100 * DCGM_FI_DEV_FB_USED{pod="%s"} / (DCGM_FI_DEV_FB_FREE{pod="%s"} + DCGM_FI_DEV_FB_USED{pod="%s"}))`, podName, podName, podName)
	}
	if company == "huawei" {
		query = fmt.Sprintf(
			`100 * container_npu_used_memory{pod_name="%s"} / container_npu_total_memory{pod_name="%s"}`, // NPU hbm memory utilization
			podName, podName)
	}
	if company == "cambricon" {
		query = fmt.Sprintf(
			`mlu_memory_utilization * on(uuid) group_right mlu_container{pod="%s"}`, podName) // MLU memory utilization
	}
	if company == "iluvatar" {
		query = fmt.Sprintf(`ix_mem_utilization{pod="%s"}`, podName) // iluvatar GPU memory utilization
	}
	if company == "hygon" {
		query = fmt.Sprintf(`100 * dcu_usedmemory_bytes{dcu_pod_name="%s"} / dcu_memorycap_bytes{dcu_pod_name="%s"}`, podName)
	}

	return p.query(query, start, size, step)
}

func (p *prometheus) QueryNetworkReceiveBytes(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error) {
	query := fmt.Sprintf(`sum(rate(container_network_receive_bytes_total{pod="%s"}[1m]))`, podName)
	return p.query(query, start, size, step)
}

func (p *prometheus) QueryNetworkTransmitBytes(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error) {
	query := fmt.Sprintf(`sum(rate(container_network_transmit_bytes_total{pod="%s"}[1m]))`, podName)
	return p.query(query, start, size, step)
}

func (p *prometheus) QueryFSUsageBytes(ctx context.Context, podName string, start int64, size int, step int) ([]float64, error) {
	query := fmt.Sprintf(`sum(container_fs_usage_bytes{device=~"^/dev/.*$",pod_name="%s"} or container_fs_usage_bytes{device=~"^/dev/.*$",pod="%s"})`, podName, podName)
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
