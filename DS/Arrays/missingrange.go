package Arrays

import (
	"fmt"
	"math"
	"strconv"
)

var result = []string{}

//FindMissingRange - find missing range
func FindMissingRange(nums []int, lower, upper int) []string {
	if len(nums) == 0 {
		findRange(lower, upper)
		return result
	}
	if nums[0] != math.MinInt32 {
		findRange(lower, nums[0]-1)
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			continue
		}
		findRange(nums[i]+1, nums[i+1]-1)
	}
	if nums[len(nums)-1] != math.MaxInt32 {
		findRange(nums[len(nums)-1]+1, upper)
	}
	return result
}

func findRange(low, high int) {
	if low > high {
		return
	}
	if low == high {
		format := fmt.Sprint(strconv.Itoa(low), "")
		result = append(result, format)
		return
	}
	format := fmt.Sprint(strconv.Itoa(low), "->", high)
	result = append(result, format)

}
