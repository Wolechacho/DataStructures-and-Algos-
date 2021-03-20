package DS

func arrangeCoins(n int) int {
	sum := 0
	i := 1
	for i = 1; i <= n; i++ {
		if i < (n - sum) {
			sum += i
		} else {
			break
		}
	}
	return i - 1
}
