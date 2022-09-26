package main

import "fmt"

func main() {
	run()
}

func run() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("test")
			panic(r)
		}
	}()

	panic("error log")
}
