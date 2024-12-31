package Arrays

import (
	"math"
	"sort"
)

func MinAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	minDifference := math.MaxFloat64
	for i := 1; i < len(arr); i++ {
		difference := arr[i] - arr[i-1]
		minDifference = math.Min(float64(minDifference), float64(difference))
	}

	pairs := [][]int{}

	for i := 1; i < len(arr); i++ {
		difference := arr[i] - arr[i-1]
		if int(minDifference) == difference {
			list := []int{}
			list = append(list, arr[i-1], arr[i])
			pairs = append(pairs, list)
		}
	}

	return pairs
}
