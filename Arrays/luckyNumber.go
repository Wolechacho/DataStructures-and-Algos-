package Arrays

func GetLuckyNumber(nums []int) int {
	m := map[int]int{}
	max := -1
	for _, num := range nums {
		m[num]++
	}

	for k, v := range m {
		if k == v {
			if v > max {
				max = k
			}
		}
	}

	if max > -1 {
		return max
	}
	return -1
}
