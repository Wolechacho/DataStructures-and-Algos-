package main

import "fmt"

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	r := []rune(A)
	for i := 0; i < len(A)-1; i++ {
		for j := 0; j < len(A)-1; j++ {
			r[j], r[j+1] = r[j+1], r[j]
		}
		fmt.Println(string(r), "  ", B)
		if string(r) == B {
			return true
		}
	}
	return false
}
