package main

import "fmt"

type item struct {
	m map[string]bool
}

func main() {
	e := &item{}

	for _, val := range e.m {
		fmt.Println(val)
	}
	fmt.Println(e.m == nil)
}
