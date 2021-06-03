package slice

import "testing"

type value struct {
	id uint64
}

func TestSliceNilAppend(t *testing.T) {
	var list []*value

	println(len(list))

	list = append(list, nil)

	println(len(list))
}

