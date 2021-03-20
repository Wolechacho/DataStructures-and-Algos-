package DS

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, found := m[nums[i]]; found {
			return true
		} else {
			m[nums[i]] = 1
		}
	}
	return false
}
