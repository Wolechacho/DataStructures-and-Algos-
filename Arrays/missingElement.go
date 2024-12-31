package Arrays

func MissingElement(nums []int, k int) int {
	if len(nums) == 0 {
		return k
	}

	length := len(nums)
	if k >= nums[length-1] {
		return k + nums[0] + length - 1
	}
	num := nums[0]
	index := 1
	for index < length && k > 0 {
		num++
		if num == nums[index] {
			index++
		} else {
			k--
		}

	}
	return num + k
}
