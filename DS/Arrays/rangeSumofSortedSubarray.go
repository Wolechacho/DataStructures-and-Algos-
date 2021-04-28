package Arrays

import (
	"fmt"
	"math"
	"sort"
)

func RangeSum(nums []int, left, right int) float64 {
	mod := 1e9 + 7
	sumArray := []int{}
	for i := 0; i < len(nums); i++ {
		pre := 0
		for j := i; j < len(nums); j++ {
			pre += nums[j]
			sumArray = append(sumArray, pre)
		}
	}

	sort.Ints(sumArray)
	// result := 0
	// for _, s := range sumArray[left-1 : right] {
	// 	result += s
	// }

	var ans float64 = 0
	var b float64 = 0
	for i := left; i <= right; i++ {
		b = ans + float64(sumArray[i-1])
		fmt.Println("b", b)

		ans = math.Mod(b, mod)
		fmt.Println("ans", ans)
		
		fmt.Println()
	}

	return ans
}
