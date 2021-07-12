package slice

import (
	"fmt"
	"testing"
)

type value struct {
	id uint64
}

func TestSliceNilAppend(t *testing.T) {
	var list []*value

	println(len(list))

	list = append(list, nil)

	println(len(list))

	test := []int{1,2,3}

	set := append(test[:1], test[2:]...)
	fmt.Println(test)
	fmt.Println(set)
}

