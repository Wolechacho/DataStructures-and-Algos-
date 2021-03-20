package DS

import "strings"

func wordPattern(pattern string, s string) bool {
	m := map[string]int{}
	words := strings.Split(s, " ")
	pt := strings.Split(pattern, "")

	if len(words) != len(pt) {
		return false
	}
	for i := 0; i < len(pt); i++ {
		_, found := m[string(words[i])]
		if !found {
			m[string(words[i])] = i
		}
		_, found1 := m[string(pt[i])]
		if !found1 {
			m[string(pt[i])] = i
		}

		if m[string(pt[i])] != m[string(words[i])] {
			return false
		}

	}
	return true
}
