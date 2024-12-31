package Arrays

import "sort"

func canBeEqualReverseArrary(nums, nums2 []int) bool {
	if len(nums) != len(nums2) {
		return false
	}
	sort.Ints(nums)
	sort.Ints(nums2)

	for i := 0; i < len(nums); i++ {
		if nums[i] != nums2[i] {
			return false
		}
	}

	return true

}
