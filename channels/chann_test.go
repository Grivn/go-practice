package channels

import (
	"sync"
	"testing"
)

type ChanTest struct {
	id uint64
	event interface{}
}

func TestChannels(t *testing.T) {
	//ct := &ChanTest{id: uint64(1)}

	ch := make(chan *ChanTest)

	//ch <- ct
	//_ = <-ch

	var wg sync.WaitGroup
	var val *ChanTest

	wg.Add(2)
	go func() {
		ch <- nil
		wg.Done()
	}()

	go func() {
		var active bool
		val, active = <-ch
		print(active)
		print(val)
		wg.Done()
	}()
	wg.Wait()
}
