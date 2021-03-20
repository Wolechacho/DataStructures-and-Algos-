package DS

import "fmt"

func findKthPositive(arr []int, k int) int {
	m := []int{}
	var l = 1
	for i := 0; i < len(arr); i++ {
		for j := l; j < arr[i]; j++ {
			m = append(m, j)
		}
		l = arr[i] + 1
	}

	if len(m) < k {
		diff := k - len(m)
		rear := len(arr) - 1
		r := arr[rear] + diff
		return r
	}
	fmt.Println(m)
	return m[k-1]
}
