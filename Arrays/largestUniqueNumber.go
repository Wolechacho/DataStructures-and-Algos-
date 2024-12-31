package Arrays

import (
	"fmt"
	"math"
	"sort"
)

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
	fmt.Println(A)
	length := len(A)

	//check if the last element and the previous are not the same
	if A[length-1] != A[length-2] {
		return A[length-1]
	}

	//start from the second to the last and compare curr,next,prev
	for i := length - 2; i > 0; i-- {
		prevNum := A[i-1]
		curNum := A[i]
		nextNum := A[i+1]
		fmt.Println(prevNum)
		fmt.Println(curNum)
		fmt.Println(nextNum)

		//if not the same return curNum
		if curNum != prevNum && curNum != nextNum {
			return curNum
		}
	}
	//means there is duplicate ex : [1,2,2]
	if A[0] != A[1] {
		return A[0]
	}
	return -1
}

func LargestUniqueNumber1(A []int) int {
	m := map[int]int{}
	for _, num := range A {
		m[num]++
	}

	result := -1
	for k, v := range m {
		if v == 1 {
			result = int(math.Max(float64(result), float64(k)))
		}
	}
	return result
}
