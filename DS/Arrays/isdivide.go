package Arrays

import (
	"sort"
)

func IsPossibleDivide(nums []int, k int) bool {
	length := len(nums)
	if length%k != 0 {
		return false
	}

	sort.Ints(nums)
	list := []int{}
	list = append(list, nums...)

	for len(list) != 0 {
		min := list[0]
		list = list[1:]

		for i := min + 1; i <= min+k-1; i++ {
			found := false
			for j, l := range list {

				if l == i {
					b1 := list[:j]
					b2 := list[j+1:]
					list = b1
					list = append(list, b2...)
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}
