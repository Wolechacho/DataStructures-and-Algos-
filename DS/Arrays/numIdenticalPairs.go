package Arrays

func NumIdenticalPairs(nums []int) int {
	m := map[int]int{}
	ans := 0
	for _, num := range nums {
		ans += m[num]
		m[num]++
	}

	return ans
}
