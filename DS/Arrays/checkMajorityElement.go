package Arrays

func CheckMajorityElement(nums []int, target int) bool {

	if len(nums) == 1 {
		return nums[0] == target
	}

	if len(nums) == 2 {
		return nums[0] == target && nums[1] == target
	}

	med := nums[len(nums)/2-1]
	med1 := nums[len(nums)/2+1]
	return med == target && med1 == target
}

