package Arrays

import (
	"math"
	"sort"
)

func MergeInterval(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result = [][]int{}
	prev := intervals[0]

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]

		if prev[1] < current[0] || prev[0] > current[1] {
			result = append(result, prev)
			prev = current
		} else {
			prev[0] = int(math.Min(float64(prev[0]), float64(current[0])))
			prev[1] = int(math.Max(float64(prev[1]), float64(current[1])))
		}
	}

	result = append(result, prev)

	return result
}
