package DS

func findTheDifference(s string, t string) string {
	var f byte = 0
	for i := 0; i < len(t); i++ {
		if i < len(s) {
			f -= s[i]
		}
		f += t[i]
	}
	return string(f)
}
