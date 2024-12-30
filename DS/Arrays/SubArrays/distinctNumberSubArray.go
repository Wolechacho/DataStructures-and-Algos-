package SubArrays

import (
	"fmt"
)

func DistintNumberSubArray(nums []int, k int) []int {
	list := []int{}
	end := len(nums) - k
	for i := 0; i <= end; i++ {
		s := nums[i : i+k]
		fmt.Println(s)
		distinctCount := 0
		m := map[int]int{}
		for _, num := range s {
			if _, exist := m[num]; !exist {
				distinctCount++
			}
			m[num]++
		}
		list = append(list, distinctCount)
	}
	return list
}
