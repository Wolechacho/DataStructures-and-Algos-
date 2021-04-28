package Arrays

func reverseFromPos(a []int, k int) []int {
	for i := 0; i < k/2; i++ {
		a[k-i-1], a[i] = a[i], a[k-i-1]
	}
	return a
}
