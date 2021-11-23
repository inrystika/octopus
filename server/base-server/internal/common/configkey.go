package common

import v1 "server/common/api/v1"

const (
	PlatformKeyJobStatusCallbackAddr = "jobStatusCallbackAddr"

	UserKeyJointCloudPermission = "jointCloudPermission"
)

var (
	PlatformConfigKeys = []*v1.ConfigKey{
		{Key: PlatformKeyJobStatusCallbackAddr, Title: "任务状态回调地址", Type: v1.ConfigKeyTypeInput},
	}

	UserConfigKeys = []*v1.ConfigKey{
		{Key: UserKeyJointCloudPermission, Title: "云际互联权限", Type: v1.ConfigKeyTypeRadio, Options: []string{"yes", "no"}},
	}
)
