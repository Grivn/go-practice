package benchtest

import "testing"

func InitMap() {
	_ = make(map[uint64]bool)
}

func WriteMap() {
	ma := make(map[uint64]bool)
	ma[uint64(1)] = true
}

func ReadMap() {
	ma := make(map[uint64]bool)

	ma[uint64(1)] = true

	_ = ma[uint64(1)]
}

func WriteMultiMap() {
	ma := make(map[uint64]bool)
	for i:=0; i<100000; i++ {
		ma[uint64(i)] = true
	}
}

func ReadMultiMap() {
	ma := make(map[uint64]bool)

	for i:=0; i<100000; i++ {
		ma[uint64(i)] = true
	}

	for i:=0; i<100000; i++ {
		_, _ = ma[uint64(i)]
	}
}

func BenchmarkInitMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		InitMap()
	}
}

func BenchmarkWriteMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WriteMap()
	}
}

func BenchmarkReadMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReadMap()
	}
}


func BenchmarkWriteMultiMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		WriteMultiMap()
	}
}

func BenchmarkReadMultiMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReadMultiMap()
	}
}
