package Arrays

import (
	"sort"
)

//CanMakeArithmeticProgression find if two consecutive
//element as thr same difference
//Return true if valid else return false,
func CanMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	difference := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != difference {
			return false
		}
	}
	return true
}
