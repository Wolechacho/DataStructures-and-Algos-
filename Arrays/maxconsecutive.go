package Arrays

import "math"

func findMaxConsecutiveOnes(nums []int) int {
	maxcount := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			maxcount = int(math.Max(float64(maxcount), float64(count)))
			count = 0
		}
	}
	maxcount = int(math.Max(float64(maxcount), float64(count)))
	return maxcount
}
