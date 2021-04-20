package DS

import (
	"fmt"
	"math"
)

func SumOfDigits(nums []int) int {

	min := math.MaxInt32

	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	fmt.Println(min)

	sum := 0
	for min != 0 {
		remainder := min % 10
		sum += remainder
		min = min / 10
	}

	if sum%2 == 0 {
		return 1

	}
	return 0
}
