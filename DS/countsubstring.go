package DS

func CountSubString(S string) int {
	sum := 0
	cur := S[0]
	c := 1
	for i := 1; i < len(S); i++ {
		if cur == S[i] {
			c++
			continue
		} else {
			sum += c * (c + 1) / 2
			c = 1
			cur = S[i]
		}
	}
	sum += c * (c + 1) / 2
	return sum

}
