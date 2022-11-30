package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go readMap1()
	go readMap2()
	wg.Wait()
}

var m = map[string]map[string]bool{"cn1": {"dc1": true, "dc2": true}, "cn2": {"dc1": true, "dc2": true}}

func readMap1() {
	count := 100000000000
	for {
		count--
		for r, m2 := range m {
			fmt.Printf("read 1 %s", r)
			for k := range m2 {
				fmt.Printf(" %s", k)
			}
			fmt.Printf("  %d\n", count)
		}
		if count == 0 {
			break
		}
	}
	wg.Done()
}

func readMap2() {
	count := 100000000000
	for {
		count--
		for r, m2 := range m {
			fmt.Printf("read 2 %s", r)
			for k := range m2 {
				fmt.Printf(" %s", k)
			}
			fmt.Printf("  %d\n", count)
		}
		if count == 0 {
			break
		}
	}
	wg.Done()
}
