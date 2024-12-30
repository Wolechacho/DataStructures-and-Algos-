package SubArrays

import (
	"fmt"
	"math"
)

func LargestSubArrayK(nums []int, k int) []int {
	// Write your code here.
	left := 0
	right := 1
	l := 0

	for right+k-1 < len(nums) && right+l < len(nums) {
		fmt.Println("condition", right+k-1)
		fmt.Println("condition2", right+l)
		if nums[left+l] == nums[right+l] {
			fmt.Println("if condition", left+l)
			fmt.Println("if condition 2", right+l)
			l++
			continue
		}
		if nums[left+l] > nums[right+l] {
			right += l + 1
		} else {
			left = int(math.Max(float64(right), math.Min(float64(left+l+1), float64(len(nums)-k))))
			right = left + 1
		}
		l = 0
	}
	return nums[left : left+k]

}
