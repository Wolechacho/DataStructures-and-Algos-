package DS

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	tags := map[string]string{"{": "}", "[": "]", "(": ")"}
	currentblock := -2
	nextblock := -1
	for i := 0; i < len(s)/2; i++ {

		//comparing extreme rear
		current1 := string(s[i])
		rear := string(s[len(s)-1-i])

		//comparing contigent
		current := string(s[currentblock+2])
		next := string(s[nextblock+2])

		//for contigent
		if tags[current] != next {
			//for rear
			if tags[current1] != rear {
				return false
			}
		}

	}
	return true
}
