package Arrays

import (
	"sort"
)

func GetStrongest(nums []int, k int) []int {
	sort.Ints(nums)
	median := nums[(len(nums)-1)/2]
	front := 0
	rear := len(nums) - 1

	ans := []int{}
	for len(ans) < k {
		if nums[rear]-median >= median-nums[front] {
			ans = append(ans, nums[rear])
			rear--
		} else {
			ans = append(ans, nums[front])
			front++
		}
	}

	return ans
}
