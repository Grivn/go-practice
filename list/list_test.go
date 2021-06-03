package list

import (
	"container/list"
	"fmt"
	"strconv"
	"testing"
)

type commandList struct {
	list *list.List
}

type commandVale struct {
	key uint64
	value string
}

func TestList(t *testing.T) {
	cl := &commandList{
		list: list.New(),
	}

	for i:=0; i<100; i++ {
		cl.list.PushBack(&commandVale{key: uint64(i), value: "hello "+strconv.Itoa(i)})
	}

	item := cl.list.Front()
	value := item.Value.(*commandVale)
	fmt.Println(value)

	item = cl.list.Front()
	value = item.Value.(*commandVale)
	fmt.Println(value)
	cl.list.Remove(item)

	item = cl.list.Front()
	value = item.Value.(*commandVale)
	fmt.Println(value)
}
