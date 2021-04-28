package Arrays

func removeDuplicates(nums []int) int {
	uniqueArray := []int{}
	currentElem := -1

	for _, num := range nums {
		if currentElem != num {
			uniqueArray = append(uniqueArray, num)
			currentElem = num
		}
	}

	return len(uniqueArray)
}
