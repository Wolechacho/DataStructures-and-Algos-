package DS


func lastWord(s string) int {
	lastWhiteSpaceIndex := 0

	for i, d := range s {
		if d == 32 {
			lastWhiteSpaceIndex = i
		}
	}
	l := len(s) - 1 - lastWhiteSpaceIndex
	if l == 0 {
		return 0
	}
	return l
}
