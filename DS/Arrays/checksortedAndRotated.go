package Arrays

import "fmt"

func Check(nums []int) bool {
	ans, n := 0, len(nums)
	for i := 0; i < n; i++ {
		fmt.Println("prev", nums[i])
		//fmt.Println("next", nums[i+1])
		fmt.Println("sol", nums[(i+1)%n])

		if nums[i] > nums[(i+1)%n] {
			ans++
			fmt.Println("ans", ans)
		}
		if ans > 1 {
			return false
		}

		fmt.Println()

	}
	return true
}
