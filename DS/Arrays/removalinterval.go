package Arrays

func RemoveInterval(intervals [][]int, toBeRemoved []int) [][]int {
	res := [][]int{}
	for _, interval := range intervals {
		// no overlap
		if interval[1] <= toBeRemoved[0] || interval[0] >= toBeRemoved[1] {
			res = append(res, []int{interval[0], interval[1]})
		} else {
			// left end no overlap
			if interval[0] < toBeRemoved[0] {
				res = append(res, []int{interval[0], toBeRemoved[0]})
			}
			// right end no overlap
			if interval[1] > toBeRemoved[1] {
				res = append(res, []int{toBeRemoved[1], interval[1]})

			}
		}
	}
	return res
}

func RemoveInterval1(intervals [][]int, toBeRemoved []int) [][]int {
	result := [][]int{}
	for _, interval := range intervals {
		if interval[0] <= toBeRemoved[1] && toBeRemoved[0] <= interval[1]{
			if interval[0] < toBeRemoved[0] {
				result = append(result, []int{interval[0], toBeRemoved[0]})
			}

			if interval[1] > toBeRemoved[1] {
				result = append(result, []int{toBeRemoved[1], interval[1]})
			}
		} else {
			result = append(result, []int{interval[0], interval[1]})
		}
	}

	return result
}
