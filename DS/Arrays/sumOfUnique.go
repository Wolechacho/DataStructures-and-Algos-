package Arrays

func sumOfUnique(nums []int) int {
	m := map[int]int{}
	var sum = 0
	for _, num := range nums {
		if _, found := m[num]; !found {
			m[num] = 1
			sum += num
		} else {
			m[num]++
			if val, _ := m[num]; val == 2 {
				sum -= num
			}
		}
	}

	return sum
}
