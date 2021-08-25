package snowflake

import (
	"time"

	"github.com/sony/sonyflake"
)

var defaultSonyflake *sonyflake.Sonyflake

func init() {
	startTime, _ := time.Parse("2006-01-02 15:04:05", "2021-04-09 00:00:00")
	setting := sonyflake.Settings{ //使用默认的内网ip转换为machineID
		StartTime: startTime,
		//CheckMachineID: checkMachineID,
	}

	defaultSonyflake = sonyflake.NewSonyflake(setting)
}

func NextUID() uint64 {
	id, _ := defaultSonyflake.NextID()
	return id
}
