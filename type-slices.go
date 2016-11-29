package idxsort

import (
	"sort"
)

type Int64Slice []int64

func (s Int64Slice) Len() int {
	return s.Len()
}

func (s Int64Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Int64Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Int64Slice) Sort() {
	sort.Sort(s)
}

type Uint64Slice []uint64

func (s Uint64Slice) Len() int {
	return s.Len()
}

func (s Uint64Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Uint64Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Uint64Slice) Sort() {
	sort.Sort(s)
}
