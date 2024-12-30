package SubArrays

import "math"

func maxSubArray(nums []int) int {
	var maxsum = float64(nums[0])
	var currentsum = float64(nums[0])

	for i := 0; i < len(nums); i++ {
		currentsum = currentsum + float64(nums[i])
		currentsum = math.Max(currentsum, float64(nums[i]))
		maxsum = math.Max(maxsum, currentsum)
	}
	return int(maxsum)
}
