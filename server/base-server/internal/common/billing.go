package common

// 模型开发和训练生成extraInfo
func GetExtraInfo(userId string) map[string]string {
	return map[string]string{
		"userId": userId,
	}
}
