package Arrays

import (
	"fmt"
	"strconv"
)

func SummaryRanges(arr []int) []string {

	result := []string{}
	current := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i-1]+1 != arr[i] {
			s := strconv.Itoa(current)
			if arr[i-1] != current {
				s += "->" + strconv.Itoa(arr[i-1])

			}
			result = append(result, s)
			current = arr[i]
		}
	}
	s := strconv.Itoa(current)
	if arr[len(arr)-1] != current {
		s += "->" + strconv.Itoa(arr[len(arr)-1])
	}
	result = append(result, s)
	return result
}

func summaryRanges2(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	current := nums[0]
	result := []string{}
	for i := 1; i < len(nums); i++ {
		if nums[i-1]+1 != nums[i] {
			if current != nums[i-1] {
				s := fmt.Sprintf("%d->%d", current, nums[i-1])
				result = append(result, s)
			} else {
				s := fmt.Sprintf("%d", current)
				result = append(result, s)
			}
			current = nums[i]
		}
	}

	if current == nums[len(nums)-1] {
		s := fmt.Sprintf("%d", nums[len(nums)-1])
		result = append(result, s)
	} else {
		f := nums[len(nums)-1]
		s := fmt.Sprintf("%d->%d", current, f)
		result = append(result, s)
	}
	return result
}
