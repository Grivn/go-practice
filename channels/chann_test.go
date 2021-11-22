package channels

import (
	"fmt"
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

func TestDChannels(t *testing.T) {
	ch := make(chan bool)

	var wg sync.WaitGroup

	closeC := make(chan bool)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-closeC:
				fmt.Println("closed coroutine")
				return
			case <-ch:
				fmt.Println("test line")
			}
		}
	}()

	go func() {
		ch <- true
		fmt.Println("send finished")
		//wg.Done()
	}()

	drainChannel(ch)
	close(closeC)
	wg.Wait()
}

func drainChannel(ch chan bool) {
DrainLoop:
	for {
		select {
		case <-ch:
		default:
			break DrainLoop
		}
	}
}
