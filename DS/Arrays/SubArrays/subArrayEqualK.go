package SubArrays

import "fmt"

func SubarraySumEqualK(nums []int, k int) int {
	m := map[int]int{0: 1}

	sum := 0
	count := 0
	for _, num := range nums {
		sum += num
		if _, exist := m[sum-k]; exist {
			count += m[sum-k]
		}
		fmt.Println(m)
		m[sum]++
	}
	return count
}
