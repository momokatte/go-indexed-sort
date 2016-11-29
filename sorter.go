package idxsort

import (
	"sort"
)

type IndexedSorter struct {
	index     sort.Interface
	swapFuncs []func(i, j int)
	ascending bool
}

func NewIndexedSorter(index sort.Interface, swapFuncs ...func(i, j int)) (s *IndexedSorter) {
	s = &IndexedSorter{
		index:     index,
		swapFuncs: swapFuncs,
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
	for _, swapFunc := range s.swapFuncs {
		swapFunc(i, j)
	}
}

func (s *IndexedSorter) Less(i, j int) bool {
	if s.ascending {
		return s.index.Less(i, j)
	} else {
		return s.index.Less(j, i)
	}
}

func SortWithIndex(index sort.Interface, ascending bool, swapFuncs ...func(i, j int)) {
	sis := NewIndexedSorter(index, swapFuncs...)
	sis.Sort(ascending)
}

func SortIndexed(index sort.Interface, ascending bool, linked sort.Interface) {
	SortWithIndex(index, ascending, linked.Swap)
}

func SortIndexedInts(index sort.Interface, ascending bool, is []int) {
	linked := sort.IntSlice(is)
	SortWithIndex(index, ascending, linked.Swap)
}

func SortIndexedInt64s(index sort.Interface, ascending bool, is []int64) {
	linked := Int64Slice(is)
	SortWithIndex(index, ascending, linked.Swap)
}

func SortIndexedUint64s(index sort.Interface, ascending bool, is []uint64) {
	linked := Uint64Slice(is)
	SortWithIndex(index, ascending, linked.Swap)
}

func SortIndexedFloat64s(index sort.Interface, ascending bool, fs []float64) {
	linked := sort.Float64Slice(fs)
	SortWithIndex(index, ascending, linked.Swap)
}

func SortIndexedStrings(index sort.Interface, ascending bool, ss []string) {
	linked := sort.StringSlice(ss)
	SortWithIndex(index, ascending, linked.Swap)
}

func SortWithIntIndex(index []int, ascending bool, swapFuncs ...func(i, j int)) {
	sis := NewIndexedSorter(sort.IntSlice(index), swapFuncs...)
	sis.Sort(ascending)
}

func SortWithInt64Index(index []int64, ascending bool, swapFuncs ...func(i, j int)) {
	sis := NewIndexedSorter(Int64Slice(index), swapFuncs...)
	sis.Sort(ascending)
}

func SortWithUint64Index(index []uint64, ascending bool, swapFuncs ...func(i, j int)) {
	sis := NewIndexedSorter(Uint64Slice(index), swapFuncs...)
	sis.Sort(ascending)
}

func SortWithFloat64Index(index []float64, ascending bool, swapFuncs ...func(i, j int)) {
	sis := NewIndexedSorter(sort.Float64Slice(index), swapFuncs...)
	sis.Sort(ascending)
}

func SortWithStringIndex(index []string, ascending bool, swapFuncs ...func(i, j int)) {
	sis := NewIndexedSorter(sort.StringSlice(index), swapFuncs...)
	sis.Sort(ascending)
}
