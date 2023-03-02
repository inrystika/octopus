package prometheus

import (
	"context"
	"fmt"
	"testing"
)

func TestCpuUsage(t *testing.T) {
	ctx := context.Background()
	p := NewPrometheus("http://192.168.203.154:30003")
	podName := "e11274a8bb7148b5bf30d15b6c117c2e-task0-0"
	start := int64(1677567000)
	step := 30
	metricValues, err := p.QueryCpuUsage(ctx, podName, start, 100, step)
	if err != nil {
		panic(err)
	}
	fmt.Println(metricValues)
	metricValues, err = p.QueryMemUsage(ctx, podName, start, 100, step)
	if err != nil {
		panic(err)
	}
	fmt.Println(metricValues)
	metricValues, err = p.QueryGpuUtil(ctx, podName, start, 100, step)
	if err != nil {
		panic(err)
	}
	fmt.Println(metricValues)
	metricValues, err = p.QueryGpuMemUtil(ctx, podName, start, 100, step)
	if err != nil {
		panic(err)
	}
	fmt.Println(metricValues)
}
