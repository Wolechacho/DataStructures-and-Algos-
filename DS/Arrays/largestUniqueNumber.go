package Arrays

import "sort"

func LargestUniqueNumber(A []int) int {
	//arr ie empty
	if len(A) == 0 {
		return -1
	}

	//array contains 1 elements
	if len(A) == 1 {
		return A[0]
	}

	//sort the array
	sort.Ints(A)
	length := len(A)

	//array contains to 2 element
	if A[length-1] != A[length-2] {
		return A[length-1]
	}

	//start from the second to the last and compare curr,next,prev
	for i := length - 2; i > 0; i-- {
		prevNum := A[i-1]
		curNum := A[i]
		nextNum := A[i+1]
		//if not the same return curNum
		if curNum != prevNum && curNum != nextNum {
			return curNum
		}

	}
	//means there si duplicate ex : [1,2,2]
	if A[0] != A[1] {
		return A[0]
	}
	return -1
}
