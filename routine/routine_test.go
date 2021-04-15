package routine

import (
	"sync"
	"testing"
)

type routine struct {
	chan1 chan bool
	chan2 chan bool
	chan3 chan bool
}

func (r *routine) listener() {
	for {
		select {
		case <-r.chan1:
		case <-r.chan2:
		case <-r.chan3:
		//default:
		//	continue
		}
	}
}

func TestListener(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	r := &routine{
		chan1: make(chan bool),
		chan2: make(chan bool),
		chan3: make(chan bool),
	}
	go r.listener()
	//wg.Wait()
}
