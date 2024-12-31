package main

import "math"

func reverseNum(num int) int {
	max := math.MaxInt32
	min := math.MinInt32
	val := 0

	for num != 0 {
		divisor := num / 10
		remainder := num - divisor*10
		num = divisor

		if val >= max || val <= min {
			return 0
		}
		val = val*10 + remainder
	}
	return val
}
