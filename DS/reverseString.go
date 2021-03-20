package DS

func reverseString(word string) string {
	start := 0
	end := len(word) - 1
	r := []rune(word)
	for start <= end {
		r[start], r[end] = r[end], r[start]
		start++
		end--
	}
	return string(r)
}
