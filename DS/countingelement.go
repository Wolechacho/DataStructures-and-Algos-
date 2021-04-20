package DS

func CountingElement(arr []int) int {
	m := map[int]int{}
	for _, num := range arr {
		m[num] = 1
	}

	count := 0
	for _, num := range arr {
		if _, exist := m[num+1]; exist {
			count++
		}
	}

	return count

}
