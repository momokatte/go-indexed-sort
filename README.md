
# go-indexed-sort

A Go package for using sequences compatible with sort.Interface to sort other sequences.


# Status

Current master branch should be considered a "release candidate" with interfaces and type signatures subject to change.

I will tag a 1.0 release when I am satisfied with usability and test coverage.


# Usage

Use a slice of ints to sort a slice of structs:

	var people []Person
	// ... populate with 4 people ...
	swapPeople := func(i, j int) {
		people[i], people[j] = people[j], people[i]
	}
	ageIndex := []int{32,56,19,24}
	idxsort.SortAsc(sort.IntSlice(ageIndex), swapPeople)

sort.IntSlice attaches the methods of sort.Interface to 'ageIndex'. Slice 'ageIndex' is then sorted by ascending value, with index swaps duplicated in slice 'people' by the 'swapPeople' closure.

If your linked sequence implements sort.Interface, you can use its Swap function:

	names := []string{"Amy","Bob","Carla","Dan"}
	ageIndex := []int{32,56,19,24}
	idxsort.SortAsc(sort.IntSlice(ageIndex), sort.StringSlice(names).Swap)


## Online GoDoc

https://godoc.org/github.com/momokatte/go-indexed-sort
