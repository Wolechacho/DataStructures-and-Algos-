package Arrays

import (
	"math"
)

func DominantIndex(nums []int) int {
	max := 0
	maxind := 0
	for i := 0; i < len(nums); i++ {
		max = int(math.Max(float64(nums[i]), float64(max)))
		if max == nums[i] {
			maxind = i
		}
	}

	for i := 0; i < len(nums); i++ {
		if i == maxind {
			continue
		}
		if max < 2*nums[i] {
			return -1
		}
	}
	return maxind
}
