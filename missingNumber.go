package main

import "fmt"

func missingNumber(nums []int) int {
	actualSum := 0
	p := ((len(nums) + 1) / 2)
	fmt.Println(p)
	expectedSum := len(nums) * p
	for _, num := range nums {
		actualSum += num
	}
	return expectedSum - actualSum
}
