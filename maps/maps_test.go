package maps

import (
	"testing"
)

type testMap struct {
	id uint64
}

func TestMaps(t *testing.T) {
	m := make(map[uint64]*testMap)

	m[uint64(1)] = &testMap{id: uint64(1)}
	m[uint64(2)] = nil

	val1, ok1 := m[uint64(1)]
	val2, ok2 := m[uint64(2)]
	val3, ok3 := m[uint64(3)]

	println(val1)
	println(val2)
	println(val3)

	println(ok1)
	println(ok2)
	println(ok3)
}

