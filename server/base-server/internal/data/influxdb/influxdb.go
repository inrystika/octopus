package influxdb

import (
	"fmt"
	"server/base-server/internal/conf"
	"server/common/errors"
	"time"

	"net/url"

	influxdbClient "github.com/influxdata/influxdb/client"
)

type Influxdb interface {
	//查询
	Query(cmd string) (res []influxdbClient.Result, err error)
	//写入
	Write(measurement string, tags map[string]string, fields map[string]interface{}) (err error)
}

type influxbd struct {
	conf   *conf.Data
	client *influxdbClient.Client
}

func NewInfluxdb(conf *conf.Data) (db Influxdb, err error) {

	url := &url.URL{
		Scheme: "http",
		Host:   conf.Influxdb.Addr,
	}

	iConfig := &influxdbClient.Config{
		URL:      *url,
		Username: conf.Influxdb.Username,
		Password: conf.Influxdb.Password,
	}

	client, err := influxdbClient.NewClient(*iConfig)
	if err != nil {
		err = errors.Errorf(err, errors.ErroInfluxdbInitFailed)
		return nil, err
	}

	if _, _, err := client.Ping(); err != nil {
		err = fmt.Errorf("failed to ping influxDB server at %q - %v", conf.Influxdb.Addr, err)
		return nil, errors.Errorf(err, errors.ErroInfluxdbInitFailed)
	}

	if err != nil {
		return nil, errors.Errorf(err, errors.ErroInfluxdbInitFailed)
	}

	influxdb := &influxbd{
		conf:   conf,
		client: client,
	}

	return influxdb, nil
}

func (i *influxbd) Query(cmd string) (res []influxdbClient.Result, err error) {

	q := influxdbClient.Query{
		Command:  cmd,
		Database: i.conf.Influxdb.Database,
	}
	if response, err := i.client.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

func (i *influxbd) Write(measurement string, tags map[string]string, fields map[string]interface{}) (err error) {

	point := influxdbClient.Point{
		Measurement: measurement,
		Time:        time.Now().UTC(),
		Fields:      fields,
		Tags:        tags,
	}

	dataPoints := make([]influxdbClient.Point, 0, 10)
	dataPoints = append(dataPoints, point)

	batchPoints := influxdbClient.BatchPoints{
		Points:          dataPoints,
		Database:        i.conf.Influxdb.Database,
		RetentionPolicy: "default",
	}

	if _, err := i.client.Write(batchPoints); err != nil {
		return err
	}
	return nil
}
