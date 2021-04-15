package benchtest

import (
	"testing"
)

func WriteSlice() {
	var slice []uint64
	slice = append(slice, uint64(1))
}

func ReadSlice() {
	var slice []uint64
	slice = append(slice, uint64(1))
	_ = slice[0]
}

func BenchmarkWriteSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WriteSlice()
	}
}

func BenchmarkReadSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReadSlice()
	}
}

func BenchmarkLenSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var slice []uint64
		_ = len(slice)
	}
}
