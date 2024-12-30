package Sums

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
		front := i + 1
		rear := len(nums) - 1


		for front < rear {
			sum := nums[i]+nums[front]+nums[rear]
			fmt.Println("sum",sum)
			if nums[i]+nums[front]+nums[rear] < target {
				count += rear - front // adding a smaller number must also be less than Target value
				front++
			} else {
				rear--
			}
		}
	}
	return count
}
