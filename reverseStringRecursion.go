package main

func ReverseStringRecursion(letters string) string {
	if len(letters) == 1 {
		return letters
	}

	startIndex := len(letters) - 1
	endIndex := len(letters)
	character := letters[startIndex:endIndex]
	letters = letters[0 : len(letters)-1]
	return character + ReverseStringRecursion(letters)
}
