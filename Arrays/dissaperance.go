package Arrays

func Dissaperance(nums []int) []int {
	list := []int{}

	counts := make([]int, len(nums)+1)
	for _, num := range nums {
		counts[num]++
	}

	for i := 1; i < len(counts); i++ {
		if counts[i] < 1 {
			list = append(list, i)
		}
	}
	return list
}
