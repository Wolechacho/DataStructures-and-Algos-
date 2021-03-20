package DS

import "math"

func firstUniqChar(s string) int {
	var tt = map[string]struct {
		no           int
		currentIndex int
	}{}

	for i, p := range s {
		if _, found := tt[string(p)]; !found {
			tt[string(p)] = struct {
				no           int
				currentIndex int
			}{1, i}
		} else {
			df := tt[string(p)]
			df.no++
			df.currentIndex = i
			tt[string(p)] = df
		}

	}

	//fmt.Println(tt)
	min := math.MaxInt32
	for _, t := range tt {
		if t.no == 1 {
			if t.currentIndex < min {
				min = t.currentIndex
			}
		}
	}

	if min == math.MaxInt32 {
		return -1
	}
	return min
}
