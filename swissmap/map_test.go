package swissmap

import (
	"fmt"
	"github.com/dolthub/swiss"
	"github.com/spf13/cast"
	"testing"
)

var (
	swissMap  = swiss.NewMap[string, int](100)
	normalMap = make(map[string]int, 100)
)

func BenchmarkSwissMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := "map" + cast.ToString(n)
		swissMap.Put(key, n)
	}
}

func BenchmarkNormalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := "map" + cast.ToString(n)
		normalMap[key] = n
	}
}

func BenchmarkGetSwissMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := "map" + cast.ToString(n)
		swissMap.Get(key)
	}
}

func BenchmarkGetNormalMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		key := "map" + cast.ToString(n)
		_, _ = normalMap[key]
	}
}

func TestFunction(t *testing.T) {
	strs := []string{
		"a", "a", "b",
	}

	res := ArrRmRepeat(strs)
	fmt.Println(res)
}

func ArrRmRepeat(strs []string) []string {
	if len(strs) < 2 {
		return strs
	}

	res := make([]string, 0, len(strs))
	filter := make(map[string]bool, len(strs))
	for _, str := range strs {
		if filter[str] {
			continue
		}

		res = append(res, str)
		filter[str] = true
	}

	return res
}
