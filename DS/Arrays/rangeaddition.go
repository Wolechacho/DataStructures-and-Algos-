package Arrays

func GetModifiedArray(updates [][]int, length int) []int {
	ans := make([]int, length)

	for _, update := range updates {
		start := update[0]
		end := update[1]
		inc := update[2]

		ans[start] += inc
		if end+1 < length {
			ans[end+1] += -inc
		}

	}

	for i := 1; i < length; i++ {
		ans[i] += ans[i-1]
	}

	return ans
}
