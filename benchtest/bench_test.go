package benchtest

import (
	"sync"
	"testing"
)

// the performance of mutex is greater

var mutex = sync.Mutex{}
var ch = make(chan bool, 1)
var v bool

func UseMutex() {
	mutex.Lock()
	v = !v
	mutex.Unlock()
}
func UseChan() {
	ch <- true
	<-ch
}

func BenchmarkUseMutex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UseMutex()
	}
}

func BenchmarkUseChan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UseChan()
	}
}
