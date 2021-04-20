package DS

func MissingNumberProgression(nums []int) int {
	d := (nums[len(nums)-1] - nums[0]) / len(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+d {
			return nums[i-1] + d
		}
	}

	return 0
}
