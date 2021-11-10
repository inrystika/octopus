package jointcloud

import (
	"context"
	"fmt"
	"testing"
)

func newJointCloud() JointCloud {
	return NewJointCloud("http://192.168.207.141:8709", "test", "7ee15bc8fee766cad1bd70ccf5f4dc14", 30)
}
func TestJointCloud_ListDataSet(t *testing.T) {
	j := newJointCloud()
	ctx := context.TODO()
	reply, err := j.ListDataSet(ctx, &DataSetQuery{
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	reply, err = j.ListDataSet(ctx, &DataSetQuery{
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}
