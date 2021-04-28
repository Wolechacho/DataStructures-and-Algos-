package Arrays

import "fmt"

func NumOfSubarrays(nums []int) int {

	list := []int{}
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			list = append(list, sum)
		}

	}

	odds := []int{}
	for _, l := range list {
		if l%2 == 1 {
			odds = append(odds, l)
		}
	}
	fmt.Println(odds)
	return len(odds)
}
