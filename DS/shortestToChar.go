package DS

import (
	"fmt"
	"math"
)

func shortestToChar(s string, c byte) int {
	upperBoundary := -1
	lowerBoundary := upperBoundary
	m := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == 101 {
			lowerBoundary = upperBoundary
			upperBoundary = i

			if lowerBoundary == -1 {
				for j := 0; j < upperBoundary; j++ {
					a := math.Abs(float64(j - upperBoundary))
					m = append(m, int(a))
				}
				m = append(m, 0)
			} else {
				for j := lowerBoundary + 1; j < upperBoundary; j++ {
					v := math.Abs(float64(lowerBoundary - j))
					f := math.Abs(float64(upperBoundary - j))
					w := math.Min(v, f)
					m = append(m, int(w))
				}
				m = append(m, 0)
			}
		}
	}
	length := len(s) - 1
	if upperBoundary < length {
		for i := upperBoundary + 1; i <= length; i++ {
			m = append(m, int(math.Abs(float64(upperBoundary-i))))
		}
	}
	fmt.Println(m)
	return 0
}
