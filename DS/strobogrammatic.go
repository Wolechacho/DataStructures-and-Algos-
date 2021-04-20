package DS

func IsStrobogrammatic(num string) bool {
	// write your code here
	if len(num) == 0 {
		return true
	}

	m := map[string]string{"0": "0", "1": "1", "8": "8", "6": "9", "9": "6"}

	left := 0
	right := len(num) - 1
	for left <= right {
		d := string(num[right])
		if _, found := m[d]; !found {
			return false
		}
		if m[d] != string(num[left]) {
			return false
		}
		left++
		right--
	}

	return true
}
