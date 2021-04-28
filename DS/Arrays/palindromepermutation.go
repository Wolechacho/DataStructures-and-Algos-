package Arrays

func PermutatePalindrome(word string) bool {
	m := map[rune]int{}
	for _, s := range word {
		if _, found := m[s]; !found {
			m[s] = 1
		} else {
			delete(m, s)
		}
	}

	return len(m) <= 1
}
