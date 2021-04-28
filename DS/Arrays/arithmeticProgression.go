package Arrays

import "sort"

func CanMakeArithmeticProgression(nums []int) bool {
	sort.Ints(nums)

	difference := nums[1] - nums[0]
	count := 1
	for i := 1; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] == difference {
			count++

			if count == 2 {
				return true
			}
		} else {
			count = 1
			difference = nums[i+1] - nums[i]
		}

	}
	return false
}
