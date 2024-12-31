package Arrays

func FindDuplicates(nums []int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	arr := []int{}
	for k, v := range m {
		if v == 2 {
			arr = append(arr, k)
		}
	}

	return arr
}
