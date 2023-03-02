package prometheus

import (
	"context"
	"encoding/json"
	"fmt"
	"path"
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

func Test2(t *testing.T) {
	fmt.Println(path.Join("http://aaa", "bbb"))
}

type S1 struct {
	Values []interface{} `json:"values"`
}

func Test3(t *testing.T) {
	//r := &Reply{}
	//j := `{"status":"success","data":{"resultType":"matrix","result":[{"metric":{},"values":[[1677567090,"0"],[1677567120,"0.01637470351884654"],[1677567150,"0.21790282004677583"],[1677567180,"0.08770905114814063"],[1677567210,"0.395199227323312"],[1677567240,"0.3379456741374604"],[1677567270,"0.00004642085314153978"],[1677567300,"1.1280422811411634"],[1677567330,"1.9997428162690427"],[1677567360,"1.959984365638428"],[1677567390,"0.39353920420726524"],[1677567420,"1.1254682080697571"],[1677567450,"1.9987111782561109"],[1677567480,"0.9439737209099524"],[1677567510,"0"],[1677567540,"0.0000032845062917319742"],[1677567570,"0"],[1677567600,"0"]]}]}}`
	r := &S1{}
	j := `{"values":[1234, "asdf"]}`
	err := json.Unmarshal([]byte(j), r)
	if err != nil {
		panic(err)
	}
}
