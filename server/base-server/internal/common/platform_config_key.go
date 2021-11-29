package common

import v1 "server/common/api/v1"

const (
	PlatformKeyJobStatusCallbackAddr = "jobStatusCallbackAddr"
)

var (
	PlatformConfigKeys = []*v1.ConfigKey{
		{Key: PlatformKeyJobStatusCallbackAddr, Title: "任务状态回调地址", Type: v1.ConfigKeyTypeInput},
	}
)
