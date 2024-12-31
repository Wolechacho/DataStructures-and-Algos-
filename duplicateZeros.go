package main

import "fmt"

func duplicateZeros(arr []int) {
	for i := 1; i <= len(arr)-1; i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			arr[i+1] = 0
			i = i + 1
		}
	}
	fmt.Println(arr)
}
