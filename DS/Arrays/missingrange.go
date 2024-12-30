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

func FindMissingRange2(nums []int,lower,upper int) []string{
	result := []string{}
	if lower != math.MinInt32 {
		if lower != nums[0] {
			s := fmt.Sprintf("%d->%d", lower, nums[0]-1)
			result = append(result, s)
		}
	}

	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 != nums[i+1] {
			low := nums[i] + 1
			high := nums[i+1] - 1

			if low == high {
				s := fmt.Sprintf("%d", low)
				result = append(result, s)
			} else if low < high {
				s := fmt.Sprintf("%d->%d", low, high)
				result = append(result, s)
			}
		}
	}
	if upper != math.MaxInt32 {
		if nums[len(nums)-1] != upper {
			s := fmt.Sprintf("%d->%d", nums[len(nums)-1]+1, upper)
			result = append(result, s)
		}
	}

	return result

}
