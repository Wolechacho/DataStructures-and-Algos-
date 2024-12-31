package Arrays

import "math"

func FindTheDistanceValue(a, b []int, d int) int {
	ans := 0
	for i := 0; i < len(a); i++ {
		found := false
		for j := 0; j < len(b) && !found; j++ {
			fa := float64(a[i])
			fb := float64(b[j])
			if math.Abs(fa-fb) <= float64(d) {
				found = true
			}
		}
		if !found {
			ans++
		}
	}
	return ans
}
