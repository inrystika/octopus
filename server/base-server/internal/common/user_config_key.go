package common

import v1 "server/common/api/v1"

const (
	UserKeyJointCloudPermission = "jointCloudPermission"
)

var (
	UserConfigKeys = []*v1.ConfigKey{
		{Key: UserKeyJointCloudPermission, Title: "云际互联权限", Type: v1.ConfigKeyTypeRadio, Options: []string{"yes", "no"}},
	}
)
