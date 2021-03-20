package DS

func majorityElement(nums []int) int {
	m := map[int]int{}
	max := 0
	for _, num := range nums {
		if val, found := m[num]; found {
			m[num] = val + 1
			v := m[num]
			if v > max {
				max = num
			}
		} else {
			m[num] = 1
		}
	}
	return max
}
