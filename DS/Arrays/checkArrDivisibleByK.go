package Arrays

import "fmt"

func CanArrange(nums []int, k int) {
	m := map[int]int{}
	for _, num := range nums {
		g := (num%k + k) % k
		m[g]++
	}

	fmt.Println(m)
}
