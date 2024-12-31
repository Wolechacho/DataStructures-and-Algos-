package main

import (
	"fmt"
)

func ValidAnagramMapping(A, B []int) []int {
	result := []int{}

	m := map[int]int{}
	for i := 0; i < len(B); i++ {
		fmt.Println("the key :", B[i])
		m[B[i]] = i
	}
	for i := 0; i < len(A); i++ {
		j, _ := m[A[i]]
		result = append(result, j)
	}
	return result
}
