package Arrays

import (
	"math"
)

func MaxDistance(matrix [][]int) int {
	var maxDistnce float64 = 0

	firstArr := matrix[0]
	minHead := firstArr[0]
	maxHead := firstArr[len(firstArr)-1]

	for i := 1; i < len(matrix); i++ {
		currArr := matrix[i]
		minCurrVal := currArr[0]
		maxCurrVal := currArr[len(currArr)-1]

		vf := math.Abs(float64(minHead - maxCurrVal))
		gh := math.Abs(float64(maxHead - minCurrVal))

		mc := math.Max(vf, gh)
		maxDistnce = math.Max(maxDistnce, mc)

		if minCurrVal < minHead {
			minHead = minCurrVal
		}

		if maxCurrVal > maxHead {
			maxHead = maxCurrVal
		}

	}

	return int(maxDistnce)
}
