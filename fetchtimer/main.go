package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse("2006-01-02", "2022-09-22")
	fmt.Println(err)
	fmt.Println(t)
	fmt.Println(t.Unix())
	fmt.Println(time.Now().Unix())
}
