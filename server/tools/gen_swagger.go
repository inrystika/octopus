package main

import (
	"server/common/utils"
)

// 生成接口swagger文档
func main() {
	err := utils.GenSwagger()
	if err != nil {
		panic(err)
	}
}
