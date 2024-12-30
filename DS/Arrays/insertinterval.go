package Arrays

import (
	"math"
	"sort"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	var mergeIds []int
	for i := 0; i < len(intervals); i++ {
		current := intervals[i]

		if current[0] <= newInterval[1] && newInterval[0] <= current[1] {
			newInterval[0] = int(math.Min(float64(newInterval[0]), float64(current[0])))
			newInterval[1] = int(math.Max(float64(newInterval[1]), float64(current[1])))
			intervals[i] = newInterval
			mergeIds = append(mergeIds, i)
		}
	}

	if len(mergeIds) > 0 {
		start := mergeIds[0]
		end := mergeIds[len(mergeIds)-1]
		if start == end {
			return intervals
		} else {
			intervals = append(intervals[0:start], intervals[end:len(intervals)]...)
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
