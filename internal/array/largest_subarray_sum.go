package arrays

import "fmt"

func LargestSubArraySum(elements []int) int {
	currentSum, maxSum := elements[0], elements[0]
	// pointer to track the largest sum array

	for j := 1; j <= len(elements)-1; j++ {
		currentSum = max(elements[j], currentSum+elements[j])
		maxSum = max(currentSum, maxSum)
	}

	fmt.Println("The largest Subaaray sum is ", maxSum)
	return maxSum

}
