package DS

import (
	"math"
)

func ConfusingNumber(num int) bool {
	m := map[int]int{
		6: 9,
		9: 6,
		8: 8,
		0: 0,
		1: 1,
	}
	newNum := 0
	x := num
	for x != 0 {
		remainder := x % 10

		if newNum > (math.MaxInt32 / 10) {
			return false
		}
		v, exist := m[remainder]
		if !exist {
			return false
		}
		newNum = newNum*10 + v
		x /= 10
	}
	return newNum == num
}
