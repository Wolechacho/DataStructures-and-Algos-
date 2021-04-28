package Arrays

import (
	"sort"
)

func SortFrequency(nums []int) []int {
	m := map[int]int{}

	for _, num := range nums {
		m[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		count := m[nums[i]]
		count2 := m[nums[j]]

		if count == count2 {
			return nums[i] > nums[j]
		}
		return m[nums[i]] < m[nums[j]]
	})

	return nums
}
