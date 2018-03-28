package idxsort

import (
	"sort"
)

/*
Int64Slice attaches the methods of Interface to []int64, sorting in increasing order.
*/
type Int64Slice []int64

func (s Int64Slice) Len() int {
	return len(s)
}

func (s Int64Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Int64Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

/*
Sort is a convenience method.
*/
func (s Int64Slice) Sort() {
	sort.Sort(s)
}

/*
Uint64Slice attaches the methods of Interface to []uint64, sorting in increasing order.
*/
type Uint64Slice []uint64

func (s Uint64Slice) Len() int {
	return len(s)
}

func (s Uint64Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Uint64Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

/*
Sort is a convenience method.
*/
func (s Uint64Slice) Sort() {
	sort.Sort(s)
}
