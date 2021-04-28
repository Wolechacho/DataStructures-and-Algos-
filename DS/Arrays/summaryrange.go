package Arrays

import "strconv"

func SummaryRanges(arr []int) []string {

	result := []string{}
	current := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i-1]+1 != arr[i] {
			s := strconv.Itoa(current)
			if arr[i-1] != current {
				s += "->" + strconv.Itoa(arr[i-1])

			}
			result = append(result, s)
			current = arr[i]
		}
	}
	s := strconv.Itoa(current)
	if arr[len(arr)-1] != current {
		s += "->" + strconv.Itoa(arr[len(arr)-1])
	}
	result = append(result, s)
	return result
}
