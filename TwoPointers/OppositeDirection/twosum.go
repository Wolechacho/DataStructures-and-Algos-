package oppositedirection

import (
	"sort"
)

func TwoSum(nums []int, target int) []int {
	sort.Ints(nums)
	startIndex := 0
	endIndex := len(nums) - 1
	i := 0
	for i < len(nums) {
		if nums[startIndex]+nums[endIndex] < target {
			startIndex++
		} else if nums[startIndex]+nums[endIndex] > target {
			endIndex--
		} else {
			return []int{startIndex, endIndex}
		}
		i++
	}

	return []int{0, 0}
}

func TwoSumHashTable(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		difference := target - nums[i]
		if v, ok := m[difference]; !ok {
			m[nums[i]] = i
		} else {
			return []int{v, i}
		}
	}

	return []int{0, 0}
}
