package main

import (
	"fmt"
	"reflect"
)

func main() {
	replaceFunc := func(i interface{}) interface{} {
		fmt.Println("test")
		val, ok := i.(string)
		if ok {
			return "hello " + val
		}
		return i
	}

	r := NewReplacer(replaceFunc, WithKind(reflect.String))
	ret, err := r.Replacing([]byte(`{"name1":"Bob","name2":"Alice"}`))
	fmt.Println(string(ret)) // {"name1":"hello Bob","name2":"hello Alice"}
	fmt.Println(err)
}
