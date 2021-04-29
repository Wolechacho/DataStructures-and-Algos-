package Arrays

import (
	"fmt"
)

func NumSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) == 0 || k <= 1 {
		return 0
	}
	product := 1
	count := 0
	start, end := 0, 0
	length := len(nums)
	for end < length {
		product *= nums[end]
		fmt.Println("product", product)

		for product >= k {
			product /= nums[start]
			start++
		}
		count += end - start + 1
		end++
		fmt.Println("start", start)
		fmt.Println("count", count)
		fmt.Println("end", end)

		fmt.Println()
		fmt.Println()

	}
	return count
}
