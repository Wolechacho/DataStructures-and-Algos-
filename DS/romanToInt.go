package DS

func romanToInt(s string) int {
	sum := 0
	numerals := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	count := len(s) - 1
	for i := 0; i < count; i++ {
		prev := s[i]
		next := s[i+1]

		prevVal := numerals[string(prev)]
		nextVal := numerals[string(next)]

		if prevVal < nextVal {
			sum -= prevVal
		} else {
			sum += prevVal
		}
	}
	last := s[len(s)-1]
	sum += numerals[string(last)]

	return sum
}
