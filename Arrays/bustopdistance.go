package Arrays

import (
	"math"
)

func DistanceBetweenBusStops(distance []int, start, destination int) int {
	totalDistance := 0
	for _, d := range distance {
		totalDistance += d
	}


	min := int(math.Min(float64(start), float64(destination)))
	max := int(math.Max(float64(start), float64(destination)))

	distance1 := 0
	for i := min; i < max; i++ {
		distance1 += distance[i]
	}

	distance2 := totalDistance - distance1
	return int(math.Min(float64(distance1), float64(distance2)))

}
