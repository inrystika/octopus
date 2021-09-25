package influxdb

import (
	"server/base-server/internal/conf"
	"server/common/errors"
	"time"

	influxdbClient "github.com/influxdata/influxdb1-client/v2"
)

type Influxdb interface {
	//查询
	Query(cmd string) (res []influxdbClient.Result, err error)
	//写入
	Write(measurement string, tags map[string]string, fields map[string]interface{}) (err error)
}

type influxbd struct {
	conf   *conf.Data
	client influxdbClient.Client
}

func NewInfluxdb(conf *conf.Data) (db Influxdb, err error) {

	client, err := influxdbClient.NewHTTPClient(influxdbClient.HTTPConfig{
		Addr:     conf.Influxdb.Addr,
		Username: conf.Influxdb.Username,
		Password: conf.Influxdb.Password,
	})

	if err != nil {
		err = errors.Errorf(err, errors.ErroInfluxdbInitFailed)
		return nil, err
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

	batchPoints, err := influxdbClient.NewBatchPoints(influxdbClient.BatchPointsConfig{
		Database:  i.conf.Influxdb.Database,
		Precision: i.conf.Influxdb.Precision,
	})
	if err != nil {
		return err
	}

	point, err := influxdbClient.NewPoint(measurement, tags, fields, time.Now())
	if err != nil {
		return err
	}

	batchPoints.AddPoint(point)
	err = i.client.Write(batchPoints)
	if err != nil {
		return err
	}
	return nil
}
