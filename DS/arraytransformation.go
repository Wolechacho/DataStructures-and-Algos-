package DS

func ArrayTransformation(nums []int) []int {
	flag := true

	//number of days is not given so we know when to stop using flag field
	for flag {
		//start loop excluding first and last index
		flag = false
		change := make([]int, len(nums))

		for i := 1; i < len(nums)-1; i++ {
			if nums[i] < nums[i-1] && nums[i] < nums[i+1] {
				flag = true
				change[i] = 1
			} else if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
				flag = true
				change[i] = -1
			}
		}

		if flag {
			for i := 0; i < len(nums); i++ {
				nums[i] += change[i]
			}
		}
	}
	return nums
}
