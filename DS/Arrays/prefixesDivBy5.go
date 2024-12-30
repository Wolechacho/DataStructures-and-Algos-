package Arrays

func PrefixesDivBy5(nums []int) []bool {
	booleans := []bool{}
	num := 0
	for i := 0; i < len(nums); i++ {
		num = (num*2 + nums[i]) % 5

		if num == 0 {
			booleans = append(booleans, true)
		} else {
			booleans = append(booleans, false)
		}
	}
	return booleans
}
