package oppositedirection

import "unicode"

func IsPalindrome(s string) bool {
	runes := []rune(s)
	q := []rune{}
	for i := 0; i < len(runes); i++ {
		if !((runes[i] >= 65 && runes[i] <= 90) || (runes[i] >= 97 && runes[i] <= 122)) {
			continue
		}

		q = append(q, unicode.ToLower(runes[i]))
	}

	count := 0
	if len(q)%2 == 0 {
		count = len(q) / 2
	} else {
		count = len(q)/2 + 1
	}

	startIndex := 0
	endIndex := len(q) - 1
	for i := 0; i < count; i++ {
		if q[startIndex] != q[endIndex] {
			return false
		}

		startIndex++
		endIndex--
	}

	return true
}
