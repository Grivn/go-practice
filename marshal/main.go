package main

import (
	"encoding/json"
	"fmt"
)

type test struct {
	a string
	b string
}

func main() {
	t := &test{a: "a", b: "b"}
	payload, err := json.Marshal(t)

	if err != nil {
		panic(err)
	}
	fmt.Println(payload)

	var s interface{}
	err = json.Unmarshal(payload, &s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	t2, ok := s.(*test)
	if !ok {
		panic("failed")
	}
	fmt.Println(t2)
}
