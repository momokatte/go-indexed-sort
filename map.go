package idxsort

import (
	"sort"
)

func SortIntStringMap(m map[int]string, ascending bool) (is []int, ss []string) {
	for k, v := range m {
		is = append(is, k)
		ss = append(ss, v)
	}
	Sort(sort.IntSlice(is), ascending, sort.StringSlice(ss).Swap)
	return
}

func SortStringIntMap(m map[string]int, ascending bool) (ss []string, is []int) {
	for k, v := range m {
		ss = append(ss, k)
		is = append(is, v)
	}
	Sort(sort.StringSlice(ss), ascending, sort.IntSlice(is).Swap)
	return
}

func SortStringInt64Map(m map[string]int64, ascending bool) (ss []string, is []int64) {
	for k, v := range m {
		ss = append(ss, k)
		is = append(is, v)
	}
	Sort(sort.StringSlice(ss), ascending, Int64Slice(is).Swap)
	return
}

func SortStringUint64Map(m map[string]uint64, ascending bool) (ss []string, is []uint64) {
	for k, v := range m {
		ss = append(ss, k)
		is = append(is, v)
	}
	Sort(sort.StringSlice(ss), ascending, Uint64Slice(is).Swap)
	return
}

func SortStringFloat64Map(m map[string]float64, ascending bool) (ss []string, fs []float64) {
	for k, v := range m {
		ss = append(ss, k)
		fs = append(fs, v)
	}
	Sort(sort.StringSlice(ss), ascending, sort.Float64Slice(fs).Swap)
	return
}
