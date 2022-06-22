package main

// 斐波那契数列
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func sum(a, b int) int {
	return a + b
}

func preSet() {
	result := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		result[i] = string(rune(i))
	}
}

func appendSlice() {
	result := make([]string, 0)
	for i := 0; i < 1000; i++ {
		result = append(result, string(rune(i)))
	}
}
