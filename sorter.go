/*
Package idxsort provides types and functions for sorting sequences via separate indexes.
*/
package idxsort

import (
	"sort"
)

/*
IndexedSorter wraps a sequence implementing the sort.Interface interface, supports ascending and descending sort, and duplicates index swaps to other sequences via provided swap functions.

IndexedSorter implements the sort.Interface interface.
*/
type IndexedSorter struct {
	index     sort.Interface
	swapFunc  func(i, j int)
	ascending bool
}

/*
NewIndexedSorter instantiates a new IndexedSorter with the provided index and swap functions.
*/
func NewIndexedSorter(index sort.Interface, swapFunc func(i, j int)) (s *IndexedSorter) {
	s = &IndexedSorter{
		index:     index,
		swapFunc:  swapFunc,
		ascending: true,
	}
	return
}

func (s *IndexedSorter) Sort(ascending bool) {
	s.ascending = ascending
	sort.Sort(s)
}

func (s *IndexedSorter) SortAsc() {
	s.Sort(true)
}

func (s *IndexedSorter) SortDesc() {
	s.Sort(false)
}

func (s *IndexedSorter) Len() int {
	return s.index.Len()
}

func (s *IndexedSorter) Swap(i, j int) {
	s.index.Swap(i, j)
	s.swapFunc(i, j)
}

func (s *IndexedSorter) Less(i, j int) bool {
	if s.ascending {
		return s.index.Less(i, j)
	} else {
		return s.index.Less(j, i)
	}
}

func Sort(index sort.Interface, ascending bool, swapFunc func(i, j int)) {
	sorter := NewIndexedSorter(index, swapFunc)
	sorter.Sort(ascending)
}

func SortAsc(index sort.Interface, swapFunc func(i, j int)) {
	Sort(index, true, swapFunc)
}

func SortDesc(index sort.Interface, swapFunc func(i, j int)) {
	Sort(index, false, swapFunc)
}
