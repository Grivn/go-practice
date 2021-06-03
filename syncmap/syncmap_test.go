package syncmap

import (
	"sync"
	"testing"
)

var iteTimes = 10000
var writePer = 10000
var mod = -1

func BenchmarkSyncMapGo(b *testing.B) {
	var mp sync.Map
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < iteTimes; j++ {
				if i%writePer == mod {
					mp.Store(0, 0)
				} else {
					_, _ = mp.Load(0)
				}
			}

		}(i)
	}
	wg.Wait()
}

func BenchmarkMapGo(b *testing.B) {
	var mp = make(map[int]int)
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			for j := 0; j < iteTimes; j++ {
				lock.Lock()
				if i%writePer == mod {
					mp[0] = 0
				} else {
					i = mp[0]
				}
				lock.Unlock()
			}
		}(i)
	}
	wg.Wait()
}

/*
sync.Map的性能高体现在读操作远多于写操作的时候。
极端情况下，只有读操作时，是普通map的性能的44.3倍。
反过来，如果是全写，没有读，那么sync.Map还不如加普通map+mutex锁。只有普通map性能的一半。
建议使用sync.Map时一定要考虑读定比例。
当写操作只占总操作的<=1/10的时候，使用sync.Map性能会明显高很多。
 */
