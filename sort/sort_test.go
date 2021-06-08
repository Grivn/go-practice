package sort

import (
	"fmt"
	"sort"
	"testing"
)

type Timestamps []int64

func (t Timestamps) Len() int { return len(t) }
func (t Timestamps) Less(i, j int) bool { return t[i] < t[j] }
func (t Timestamps) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

type Block struct {
	Value     string
	Timestamp int64
}

type SubBlock []Block

func (s SubBlock) Len() int { return len(s) }
func (s SubBlock) Less(i, j int) bool { return s[i].Timestamp < s[j].Timestamp }
func (s SubBlock) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func TestSortFunc(t *testing.T) {
	ts := Timestamps{8, 1, 2, 5, 6, 2, 1, 1, 10, 12, 25, 34}

	sort.Sort(ts)

	fmt.Println(ts)

	sub := SubBlock{Block{"value25", 25}, Block{"value2", 2}, Block{"value9", 9}}
	sort.Sort(sub)

	fmt.Println(sub)

	blank := SubBlock{}
	sort.Sort(blank)
	fmt.Println(blank)
}
