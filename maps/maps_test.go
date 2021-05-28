package maps

import (
	"fmt"
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

type Value struct {
	data int
}

type MapStruct struct {
	m map[int]*Value
}

func NewMapStruct() *MapStruct {
	return &MapStruct{m: make(map[int]*Value)}
}

func (ms *MapStruct) write(key, data int) {
	ms.m[key] = &Value{data: data}
}

func (ms *MapStruct) readAndDelete(key int) *Value {
	value, ok := ms.m[key]
	if !ok {
		return nil
	}
	delete(ms.m, key)
	return value
}

func TestMapsClear(t *testing.T) {
	ms := NewMapStruct()
	ms.write(1, 1)
	v := ms.readAndDelete(1)
	fmt.Println(v)
	wrong := ms.readAndDelete(1)
	fmt.Println(wrong)
}
