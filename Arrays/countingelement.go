package Arrays

func CountingElement(arr []int) int {
	m := map[int]int{}
	count := 0
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}

	for _, num := range arr {
		r := num + 1
		if _, exist := m[r]; exist {
			count++
		}
	}

	return count
}
