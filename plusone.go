package main

func plusOne(arr []int) {
	var carry = 1
	for i := len(arr) - 1; i >= 0; i-- {
		if carry == 0 {
			break
		} else {
			if arr[i] != 9 {
				arr[i] = arr[i] + carry
				carry = 0
			} else {
				arr[i] = 0
				carry = 1
			}
		}
	}
}
