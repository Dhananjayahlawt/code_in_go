package array

import (
	"slices"
)

func sequentialSearch(arr []int, item int) bool {

	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

// Using Slice.Contains
func sequentialSearchUsingSlice(arr []int, item int) bool {
	return slices.Contains(arr, item)
}
