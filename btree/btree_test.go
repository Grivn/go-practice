package btree

import (
	"fmt"
	"github.com/google/btree"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)
type MyTree struct {
	Age  int
	Name string
	event interface{}
}

func (m *MyTree) Less(item btree.Item) bool {
	return m.Age < (item.(*MyTree)).Age
}

type MyTree1 struct {
	Age  int
	Name string
}

func (m *MyTree1) Less(item btree.Item) bool {
	return m.Age < (item.(*MyTree1)).Age
}

type MyTree2 struct {
	Age  int
	Point int
}

func (m *MyTree2) Less(item btree.Item) bool {
	return m.Age < (item.(*MyTree2)).Age
}

func TestDescendRange(t *testing.T) {
	tree := btree.New(2) //创建一个2-3-4 树
	for i := 0; i < 100; i++ {
		//插入数据
		tree.ReplaceOrInsert(&MyTree{Age: i, Name: "freedom" + strconv.Itoa(i)})
	}
	value := tree.Min()

	v := value.(*MyTree)

	fmt.Println(v.Age)
	fmt.Println(v.Name)

	tree.DescendRange(&MyTree{Age: 50}, &MyTree{Age: 48}, func(a btree.Item) bool {
		item := a.(*MyTree)
		fmt.Println(item)
		return true
	})

	tree.Ascend(func(a btree.Item) bool {
		item := a.(*MyTree)
		fmt.Println(item)
		return true
	})
}

func TestCon(t *testing.T) {
	tree := btree.New(2) //创建一个2-3-4 树
	for i := 0; i < 100; i++ {
		//插入数据
		tree.ReplaceOrInsert(&MyTree1{Age: i, Name: "freedom" + strconv.Itoa(i)})
	}
	value1 := tree.Min()

	v1 := value1.(*MyTree1)

	fmt.Println(v1)

	for i := 0; i < 100; i++ {
		//插入数据
		tree.ReplaceOrInsert(&MyTree2{Age: i, Point: i*10})
	}
	value2 := tree.Min()

	v2 := value2.(*MyTree2)

	fmt.Println(v2.Age)
}

func TestReplace(t *testing.T) {
	tree := btree.New(2) //创建一个2-3-4 树
	for i := 0; i < 100; i++ {
		//插入数据
		tree.ReplaceOrInsert(&MyTree{Age: i, Name: "freedom" + strconv.Itoa(i)})
	}
	value1 := tree.Min()

	v1 := value1.(*MyTree)

	fmt.Println(v1)

	for i := 0; i < 100; i++ {
		//插入数据
		tree.ReplaceOrInsert(&MyTree{Age: i, Name: "failed" + strconv.Itoa(i)})
	}
	value2 := tree.Min()

	v2 := value2.(*MyTree)

	fmt.Println(v2)
}

func TestHas(t *testing.T) {
	tree := btree.New(2) //创建一个2-3-4 树

	min := tree.Min()
	fmt.Println(min)

	v, ok := min.(*MyTree)
	fmt.Println(v)
	fmt.Println(ok)

	for i := 0; i < 100; i++ {
		//插入数据
		tree.ReplaceOrInsert(&MyTree{Age: i, Name: "freedom" + strconv.Itoa(i), event: &MyTree1{}})
	}
	value1 := tree.Min()

	v1 := value1.(*MyTree)

	fmt.Println(v1)

	for i := 0; i < 100; i++ {
		//插入数据
		assert.True(t, tree.Has(&MyTree{Age: i, Name: "failed" + strconv.Itoa(i), event: &MyTree2{}}))
	}
}
