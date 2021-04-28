package Arrays

import (
	"fmt"
	"sort"
)

func ThreeSumSmaller(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)
	fmt.Println("Sorted nums : ", nums)
	count := 0
	for i := 0; i < len(nums)-2; i++ {
		lo := i + 1
		hi := len(nums) - 1


		for lo < hi {
			jk := nums[i]
			jf := nums[lo]
			gh := nums[hi]

			fmt.Println(jk, jf, gh)
			if nums[i]+nums[lo]+nums[hi] < target {
				count += hi - lo // adding a smaller number must also be less than Target value
				lo++
			} else {
				hi--
			}
		}
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()
	return count
}
