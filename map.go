package idxsort

import (
	"sort"
)

func SortIntStringMap(m map[int]string, ascending bool) (is []int, ss []string) {
	for k, v := range m {
		is = append(is, k)
		ss = append(ss, v)
	}
	SortIndexedStrings(sort.IntSlice(is), ascending, ss)
	return
}

func SortStringIntMap(m map[string]int, ascending bool) (ss []string, is []int) {
	for k, v := range m {
		ss = append(ss, k)
		is = append(is, v)
	}
	SortIndexedInts(sort.StringSlice(ss), ascending, is)
	return
}

func SortStringInt64Map(m map[string]int64, ascending bool) (ss []string, is []int64) {
	for k, v := range m {
		ss = append(ss, k)
		is = append(is, v)
	}
	SortIndexedInt64s(sort.StringSlice(ss), ascending, is)
	return
}

func SortStringFloat64Map(m map[string]float64, ascending bool) (ss []string, fs []float64) {
	for k, v := range m {
		ss = append(ss, k)
		fs = append(fs, v)
	}
	SortIndexedFloat64s(sort.StringSlice(ss), ascending, fs)
	return
}
