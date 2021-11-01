package jointcloud

import (
	"fmt"
	"testing"
)

func newJointCloud() JointCloud {
	return NewJointCloud("http://192.168.207.141:8709", "test", "")
}
func TestJointCloud_ListDataSet(t *testing.T) {
	j := newJointCloud()
	reply, err := j.ListDataSet(&DataSetQuery{
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	reply, err = j.ListDataSet(&DataSetQuery{
		PageIndex: 1,
		PageSize:  10,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}

type Test struct {
	FirstName  string
	SecondName string
}
