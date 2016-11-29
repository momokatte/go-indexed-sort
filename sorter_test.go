package idxsort

import (
	"sort"
	"testing"
)

func TestIntStringSort(t *testing.T) {

	m := make(map[int]string)
	m[5] = "five"
	m[4] = "four"
	m[3] = "three"
	m[2] = "two"
	m[1] = "one"
	m[0] = "zero"

	is, ss := SortIntStringMap(m, true)

	expectedInts := []int{0, 1, 2, 3, 4, 5}
	expectedStrings := []string{"zero", "one", "two", "three", "four", "five"}

	for i, expectedInt := range expectedInts {
		expectedString := expectedStrings[i]
		if actualInt := is[i]; actualInt != expectedInt {
			t.Fatalf("Expected %d, got %d", expectedInt, actualInt)
		}
		if actualString := ss[i]; actualString != expectedString {
			t.Fatalf("Expected '%s', got '%s'", expectedString, actualString)
		}
	}
}

func TestStringIntSort(t *testing.T) {

	m := make(map[string]int)
	m["five"] = 5
	m["four"] = 4
	m["three"] = 3
	m["two"] = 2
	m["one"] = 1
	m["zero"] = 0

	ss, is := SortStringIntMap(m, true)

	if actual := ss[0]; actual != "five" {
		t.Fatalf("Expected 'five', got '%s'", actual)
	}
	if actual := is[0]; actual != 5 {
		t.Fatalf("Expected 5, got %d", actual)
	}
	if actual := ss[5]; actual != "zero" {
		t.Fatalf("Expected 'zero', got '%s'", actual)
	}
	if actual := is[5]; actual != 0 {
		t.Fatalf("Expected 0, got %d", actual)
	}
}

type StringThing struct {
	Value string
}

func TestStringThingSort(t *testing.T) {

	ss := []string{"zero", "one", "two", "three", "four", "five"}
	var things []StringThing
	for _, v := range ss {
		things = append(things, StringThing{v})
	}

	swapThings := func(i, j int) {
		things[i], things[j] = things[j], things[i]
	}
	SortDesc(sort.StringSlice(ss), swapThings)

	expectedStrings := []string{"zero", "two", "three", "one", "four", "five"}

	for i, expectedString := range expectedStrings {
		if actual := ss[i]; actual != expectedString {
			t.Fatalf("Expected %d, got %d", expectedString, actual)
		}
		if actual := things[i].Value; actual != expectedString {
			t.Fatalf("Expected '%s', got '%s'", expectedString, actual)
		}
	}
}

type IntThing struct {
	Value int
}

func TestIntThingSort(t *testing.T) {

	m := make(map[int]IntThing)
	m[5] = IntThing{5}
	m[4] = IntThing{4}
	m[3] = IntThing{3}
	m[2] = IntThing{2}
	m[1] = IntThing{1}
	m[0] = IntThing{0}

	var is []int
	var things []IntThing
	for k, v := range m {
		is = append(is, k)
		things = append(things, v)
	}

	swapThings := func(i, j int) {
		things[i], things[j] = things[j], things[i]
	}
	SortDesc(sort.IntSlice(is), swapThings)

	expectedInts := []int{5, 4, 3, 2, 1, 0}

	for i, expectedInt := range expectedInts {
		if actual := is[i]; actual != expectedInt {
			t.Fatalf("Expected %d, got %d", expectedInt, actual)
		}
		if actual := things[i].Value; actual != expectedInt {
			t.Fatalf("Expected '%s', got '%s'", expectedInt, actual)
		}
	}
}
