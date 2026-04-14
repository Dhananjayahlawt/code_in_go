package main

import (
	arrays "code-in-go/internal/array"
	"fmt"
)

func main() {
	array := []int{1, -2, 3, 4, -4, 6, -14, 6, 2}
	result := arrays.LargestSubArraySum(array)
	fmt.Println("Welcome to the World of Coding in Golang :- Are You Ready for the twist", result)
}
