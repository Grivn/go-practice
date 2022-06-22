package main

import "testing"

func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(10)
	}
}

func BenchmarkFib20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(20)
	}
}

func BenchmarkSum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sum(1, 2)
	}
}

func BenchmarkPreSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		preSet()
	}
}

func BenchmarkAppendSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		appendSlice()
	}
}
