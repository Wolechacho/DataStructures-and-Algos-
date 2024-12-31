package Sums

import (
	"math"
	"sort"
)

func TwosumlessathanK(nums []int, k int) int {
	sort.Ints(nums)

	front := 0
	end := len(nums) - 1
	var ans float64 = -1
	for front < end {
		if nums[front]+nums[end] < k {
			ans = math.Max(ans, float64(nums[front]+nums[end]))
			front++
		} else {
			end--
		}
	}

	return int(ans)
}
