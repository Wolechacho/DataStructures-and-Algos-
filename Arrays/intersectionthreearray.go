package Arrays

func IntersectionThreeSortedArray(arr1, arr2, arr3 []int) []int {
	if len(arr1) != len(arr2) {
		return []int{}
	}

	if len(arr2) != len(arr3) {
		return []int{}
	}

	if len(arr1) != len(arr3) {
		return []int{}
	}

	index1, index2, index3 := 0, 0, 0
	intersection := []int{}

	for index1 < len(arr1) && index2 < len(arr2) && index3 < len(arr3) {
		num1, num2, num3 := arr1[index1], arr2[index2], arr3[index3]

		if num1 == num3 && num2 == num3 {
			intersection = append(intersection, num1)
			index1++
			index2++
			index3++
		} else {
			increment1, increment2, increment3 := 0, 0, 0

			if num1 < num2 || num1 < num3 {
				increment1 = 1
			}
			if num2 < num1 || num2 < num3 {
				increment2 = 1
			}

			if num3 < num1 || num3 < num2 {
				increment3 = 1
			}
			index1 += increment1
			index2 += increment2
			index3 += increment3
		}
	}

	return intersection
}
