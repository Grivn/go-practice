package swissmap

import "github.com/dolthub/swiss"

func triggerSWISS() {
	m := swiss.NewMap[string, int](10)

	m.Put("foo", 1)
	m.Put("bar", 2)

	if x, ok := m.Get("foo"); ok {
		println(x)
	}
	if m.Has("bar") {
		x, _ := m.Get("bar")
		println(x)
	}

	m.Put("foo", -1)
	m.Delete("bar")

	if x, ok := m.Get("foo"); ok {
		println(x)
	}
	if m.Has("bar") {
		x, _ := m.Get("bar")
		println(x)
	}
}
