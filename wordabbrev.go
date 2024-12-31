package main

func ValidWordAbbreviation(word, abbr string) bool {
	count, i, j := 0, 0, 0
	for ; i < len(abbr) && j < len(word); i++ {

		if abbr[i] >= 48 && abbr[i] <= 57 {
			if count == 0 && abbr[i] == 48 {
				return false
			}
			count = count*10 + (int(abbr[i]) - 48)
		} else {
			if count != 0 {
				j += count
				count = 0
			}

			if j >= len(word) || abbr[i] != word[j] {
				return false
			}
			j++
		}

	}
	return i == len(abbr) && j+count == len(word)
}
