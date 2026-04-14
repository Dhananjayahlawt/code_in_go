package arrays

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{2, 4, 5, 6, 7, 12, 14, 15, 23, 45, 67, 89, 99}
	itemToSearch := 23
	result := binarySearch(arr, itemToSearch, 0, len(arr)-1)

	if result != true {
		t.Errorf("Expected True but got %v", result)
	}
}
