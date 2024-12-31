package Arrays

import "sort"

func RankTransform(nums []int) []int {
	arr := make([]int, len(nums))
	n := copy(arr, nums)
	result := []int{}
	if n != 0 {
		sort.Ints(arr)
		m := map[int]int{}
		i := 0
		for _, elem := range arr {
			if _, exist := m[elem]; !exist {
				i++
				m[elem] = i
			}
		}

		for _, num := range nums {
			result = append(result, m[num])
		}
	}

	return result
}
