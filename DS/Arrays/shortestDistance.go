package Arrays

import "math"

func ShortestDistance(words []string, word1, word2 string) int {
	m := -1
	n := -1

	min := math.MaxFloat64
	for i := 0; i < len(words); i++ {
		s := words[i]
		if word1 == s {
			m = i
			if n != -1 {
				min = math.Min(float64(min), float64(m-n))
			}
		} else if word2 == s {
			n = i
			if m != -1 {
				min = math.Min(float64(min), float64(n-m))
			}
		}
	}

	return int(min)
}
