package Arrays

import (
	"fmt"
	"sort"
)

func Insert(intervals [][]int, newInterval []int) [][]int {
	var mergeIds []int
	for i, interval := range intervals {
		if isOverlapping(interval, newInterval) {
			newInterval = merge(interval, newInterval)
			intervals[i] = newInterval
			mergeIds = append(mergeIds, i)
		}
	}

	fmt.Println("mergeId", mergeIds)
	if len(mergeIds) > 0 {
		start := mergeIds[0]
		end := mergeIds[len(mergeIds)-1]
		if start == end {
			return intervals
		} else {
			fmt.Println("Intervals before : ", intervals)
			intervals = append(intervals[0:start], intervals[end:len(intervals)]...)
			fmt.Println("Intervals after : ", intervals)

			return intervals
		}
	} else {
		intervals = append(intervals, newInterval)
		sort.Slice(intervals[:], func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		})
		return intervals
	}
}

func isOverlapping(a []int, b []int) bool {
	booly := (a[1] < b[0] || b[1] < a[0])
	return !booly

	// if b[0] < a[1] {
	// 	//return for overlap
	// 	return true
	// }

	// if b[1] < a[0] {
	// 	return true
	// }
	//return false
}

func merge(a []int, b []int) []int {

	var start int
	var end int
	if a[0] < b[0] {
		start = a[0]
	} else {
		start = b[0]
	}

	if a[1] > b[1] {
		end = a[1]
	} else {
		end = b[1]
	}

	return []int{start, end}
}
