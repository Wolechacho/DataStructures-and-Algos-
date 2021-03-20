package DS

func reverseOnlyLetters(S string) string {
	runes := []rune(S)
	frontIndex := 0
	endIndex := len(runes) - 1

	for frontIndex < endIndex {
		isTrue := false
		isTrue2 := false

		if (runes[frontIndex] >= 65 && runes[frontIndex] <= 90) || (runes[frontIndex] >= 97 && runes[frontIndex] <= 120) {
			isTrue = true
		}

		if (runes[endIndex] >= 65 && runes[endIndex] <= 90) || (runes[endIndex] >= 97 && runes[endIndex] <= 120) {
			isTrue2 = true
		}

		if isTrue && isTrue2 {
			runes[frontIndex], runes[endIndex] = runes[endIndex], runes[frontIndex]
			frontIndex++
			endIndex--
		} else if !isTrue && !isTrue2 {
			frontIndex++
			endIndex--
		} else if !isTrue || isTrue2 {
			frontIndex++
		} else if isTrue || !isTrue2 {
			endIndex--
		}
	}
	return string(runes)
}