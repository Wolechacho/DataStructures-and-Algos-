package DS

func trailingZeroes(n int) int {
	c := 0
	for i := 5; n/i > 0; i *= 5 {
		c += n / i
	}

	return c
}
