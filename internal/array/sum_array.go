package arrays

func sumArray(arr []int) int {
	var sum int
	for _, v := range arr {
		sum += v
	}
	return sum
}
