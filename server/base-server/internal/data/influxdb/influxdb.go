package influxdb


import (
    "fmt"
    "log"
    "time"
 
    influxbd "github.com/influxdata/influxdb1-client/v2"
)

type Influxdb interface {
	//查询
	Query(cmd string) (res []influxbd.Result, err error)
	//写入
	Write(measurement string, tags map[string]string, fields map[string]interface{}) (err error)
}

type influxbd struct {
	conf   *conf.Data
	client *influxbd.Client
}

func NewInfluxdb(conf *conf.Data) (db Influxdb, err error) {
	
	client, err := influxbd.NewHTTPClient(influxbd.HTTPConfig{
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
 
func (i *influxbd) Query(cmd string) (res []influxbd.Result, err error) {

    q := influxbd.Query{
        Command:  cmd,
        Database: i.conf.Database,
    }
    if response, err := cli.Query(q); err == nil {
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
    
	batchPoints, err := influxbd.NewBatchPoints(influxbd.BatchPointsConfig{
        Database:  i.conf.Database,
        Precision: i.conf.Precision,
    })
    if err != nil {
        return err
    }
 
    point, err := influxbd.NewPoint(measurement, tags, fields, time.Now())
	if err != nil {
        return err
    }

    batchPoints.AddPoint(point)
    err = cli.Write(bp)
    if err != nil {
        return err
    }
}