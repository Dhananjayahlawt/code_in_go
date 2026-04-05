package array

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := sumArray([]int{1, 2, 4, 5, 6, 7})
	if result != 25 {
		t.Errorf("Expected 25 but got %d", result)
	}
}
