package DS

func threeConsecutiveOdds(arr []int) bool {
	n := 3
	for _, num := range arr {
		if num%2 != 0 {
			n--
		} else {
			n = 3
		}

		if n == 0 {
			return true
		}
	}
	return false
}
