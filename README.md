
# go-indexed-sort

A Go package for using sequences compatible with sort.Interface to sort other sequences.


# Status

Current master branch should be considered a "release candidate" with interfaces and type signatures subject to change.

I will tag a 1.0 release when I am satisfied with usability and test coverage.


# Usage

Use a slice of ints to sort a slice of strings:
	
	names := []string{"Amy","Bob","Carla","Dan"}
	ageIndex := []int{32,56,19,24}
	idxsort.SortIndexed(sort.IntSlice(ageIndex), true, sort.StringSlice(names))

sort.IntSlice and sort.StringSlice attach the methods of sort.Interface to slices of basic ints and strings.  Slice 'ageIndex' is then sorted by ascending value, with index swaps mirrored in slice 'names'.

The linked sequence does not have to implement all of the methods of sort.Interface; it is also possible to use a closure with the sort.Interface.Swap signature:

	var people []Person
	// ... populate with 4 people ...
	swapPeople := func(i, j int) {
		people[i], people[j] = people[j], people[i]
	}
	ageIndex := []int{32,56,19,24}
	idxsort.SortWithIndex(sort.IntSlice(ageIndex), true, swapPeople)

Again, 'ageIndex' is sorted by ascending value and index swaps are mirrored in slice 'people' via the 'swapPeople' closure.


## Online GoDoc

https://godoc.org/github.com/momokatte/go-indexed-sort
