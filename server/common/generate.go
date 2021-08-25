package main

import (
	"server/common/utils"
)

//go:generate go run generate.go
func main() {
	err := utils.Generate()
	if err != nil {
		panic(err)
	}
}
