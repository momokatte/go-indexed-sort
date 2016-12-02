/*
Package idxsort provides types and functions for sorting sequences via separate indexes.
*/
package idxsort

import (
	"sort"
)

/*
Sorter wraps a sequence implementing the sort.Interface interface, supports ascending and descending sort, and duplicates index swaps to other sequences via a swap function.

Methods of Sorter are not safe for concurrent use. The provided index and/or swap functions may also require explicit locking to prevent concurrent access.

Sorter implements the sort.Interface interface.
*/
type Sorter struct {
	index     sort.Interface
	swapFunc  func(i, j int)
	ascending bool
}

/*
NewSorter instantiates a new Sorter with the provided index and swap function.
*/
func NewSorter(index sort.Interface, swapFunc func(i, j int)) (s *Sorter) {
	s = &Sorter{
		index:     index,
		swapFunc:  swapFunc,
		ascending: true,
	}
	return
}

/*
Sorts the index in ascending or descending order.
*/
func (s *Sorter) Sort(ascending bool) {
	s.ascending = ascending
	sort.Sort(s)
}

/*
Sorts the index in ascending order.
*/
func (s *Sorter) SortAsc() {
	s.Sort(true)
}

/*
Sorts the index in descending order.
*/
func (s *Sorter) SortDesc() {
	s.Sort(false)
}

func (s *Sorter) Len() int {
	return s.index.Len()
}

func (s *Sorter) Swap(i, j int) {
	s.index.Swap(i, j)
	s.swapFunc(i, j)
}

func (s *Sorter) Less(i, j int) bool {
	if s.ascending {
		return s.index.Less(i, j)
	} else {
		return s.index.Less(j, i)
	}
}

/*
Sorts an index in the provided order, and duplicates index swaps to swapFunc.
*/
func Sort(index sort.Interface, ascending bool, swapFunc func(i, j int)) {
	sorter := NewSorter(index, swapFunc)
	sorter.Sort(ascending)
}

/*
Sorts an index in ascending order, and duplicates index swaps to swapFunc.
*/
func SortAsc(index sort.Interface, swapFunc func(i, j int)) {
	Sort(index, true, swapFunc)
}

/*
Sorts an index in descending order, and duplicates index swaps to swapFunc.
*/
func SortDesc(index sort.Interface, swapFunc func(i, j int)) {
	Sort(index, false, swapFunc)
}
