package main

import (
	"fmt"
	"math"
)

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	total := 0
	c := int(math.Sqrt(float64(num)))
	fmt.Println(c)
	for i := 1; i <= c; i++ {
		if num%i == 0 {
			s := (i + num/i)
			total += s
		}
	}
	total -= num
	return total == num
}
