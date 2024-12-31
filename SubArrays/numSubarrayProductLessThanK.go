package SubArrays

import (
	"fmt"
)

func NumSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 || len(nums) == 0 {
		return 0
	}
	ans := 0
	product := 1
	left := 0
	for right := 0; right < len(nums); right++ {
		product *= nums[right]
		fmt.Println("product", product)

		for product >= k {
			product /= nums[left]
			fmt.Println("product in while loop", product)
			left++
		}
		ans += right - left + 1
		fmt.Println("ans", ans)
	}
	return ans
}
