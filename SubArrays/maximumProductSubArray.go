package SubArrays

import (
	"fmt"
	"math"
)

func MaxProduct(nums []int) int {
	n := len(nums)
	res := nums[0]
	front := 0
	rear := 0
	fmt.Println("nums", nums)
	fmt.Println()

	for i := 0; i < n; i++ {
		if front == 0 {
			front = nums[i] * 1
			fmt.Println("if zero for front", front)
		} else {
			front = nums[i] * front
			fmt.Println("if not zero for front", front)
		}

		if rear == 0 {
			rear = nums[n-1-i] * 1
			fmt.Println("if zero for rear", rear)

		} else {
			rear = nums[n-1-i] * rear
			fmt.Println("if not zero for rear", rear)
		}

		fmt.Println()

		s := math.Max(float64(front), float64(rear))
		res = int(math.Max(float64(res), s))
	}
	return res
}
