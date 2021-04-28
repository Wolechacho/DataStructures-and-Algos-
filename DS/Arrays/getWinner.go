package Arrays

import (
	"fmt"
)

func GetWinner(arr []int, k int) int {
	hashMap := map[int]int{}

	flag := false
	for !flag {
		prev := arr[0]
		next := arr[1]

		if prev > next {
			// for i := 1; i < len(arr)-1; i++ {
			// 	arr[i], arr[i+1] = arr[i+1], arr[i]
			// }
			hashMap[prev]++
			if val, _ := hashMap[prev]; val == k {
				flag = true
				return prev
			}
		} else {
			// for i := 0; i < len(arr)-1; i++ {
			// 	arr[i], arr[i+1] = arr[i+1], arr[i]
			// }
			hashMap[next]++

			if val, _ := hashMap[next]; val == k {
				flag = true
				return next
			}
		}
	}

	fmt.Println(hashMap)
	return -1
}

func GetWinner2(arr []int, k int) int {
	cur := arr[0]
	win := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > cur {
			cur = arr[i]
			win = 0
		}
		win++
		if win == k {
			break
		}
	}
	return cur
}
