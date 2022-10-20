package main

import (
	"fmt"
	"math"
)

func main() {
	l := []float64{}
	count := 0
	sum := float64(0)

	for _, v := range l {
		sum += v
		count++
	}

	avg := sum / float64(count)
	nan := math.IsNaN(avg)
	fmt.Println(avg)
	fmt.Println(nan)
}
