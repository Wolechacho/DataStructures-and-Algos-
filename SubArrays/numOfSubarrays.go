package SubArrays

import "math"

func NumOfSubarrays(nums []int) int {

	c1 := 0 //odd
	c2 := 1 //even
	n := len(nums)
	currSum := 0
	var ans float64 = 0
	mod := 1e9 + 7
	for i := 0; i < n; i++ {
		currSum += nums[i]
		if currSum%2 == 0 {
			ans = math.Mod(ans+float64(c1), mod)
			c2++
		} else {
			ans = math.Mod(ans + float64(c2),mod) 
			c1++
		}
	}
	return int(ans)
}
