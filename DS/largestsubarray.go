package DS

func LargestSubArrayK(nums []int, k int) []int {
	end := len(nums) - k
	maxIndex := 0
	maxNum := -1
	for i := 0; i <= end; i++ {
		if nums[i] > maxNum {
			maxIndex = i
			maxNum = nums[i]
		}
	}

	subArray := nums[maxIndex : k+maxIndex]
	return subArray
}
