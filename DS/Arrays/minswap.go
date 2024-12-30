package Arrays

import (
	"math"
)

func MinSwap(data []int) int {
	ones := 0
	for _, num := range data {
		if num == 1 {
			ones++
		}
	}

	if ones <= 1 {
		return 0
	}
	windowOnes := 0
	for i := 0; i < ones; i++ {
		if data[i] == 1 {
			windowOnes++
		}
	}
	maxOnes := windowOnes
	length := len(data)

	for i := ones; i < length; i++ {
		if data[i] == 1 {
			windowOnes++
		}
		if data[i-ones] == 1 {
			windowOnes--
		}

		maxOnes = int(math.Max(float64(maxOnes), float64(windowOnes)))
		if maxOnes == ones {
			break
		}
	}
	return ones - maxOnes

}
