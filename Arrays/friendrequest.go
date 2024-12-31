package Arrays

import (
	"sort"
)

func numFriendRequests(ages []int) int {
	sort.Ints(ages)
	count := 0

	for i := len(ages) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if !(ages[j] > 100 && ages[i] < 100) {
				a := (0.5 * float64(ages[i])) + 7
				b := float64(ages[j])
				if !(b <= a) {
					if ages[i] == ages[j] {
						count += 2
					} else {
						count++
					}
				}
			}
		}
	}
	return count
}
