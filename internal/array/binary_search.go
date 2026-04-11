package array

func binarySearch(arr []int, item, start, end int) bool {
	if start > end {
		return false
	}

	mid := (start + end) / 2

	if arr[mid] == item {
		return true
	} else if arr[mid] > item {
		return binarySearch(arr, item, start, end-1)
	} else {
		return binarySearch(arr, item, start+1, end)
	}
}
