package main

func singleNumber(nums []int) int {
	m := map[int]int{}
	sum := 0
	for _, num := range nums {
		if _, ok := m[num]; ok {
			sum -= num
		} else {
			m[num] = 1
			sum += num
		}
	}
	return sum
}
