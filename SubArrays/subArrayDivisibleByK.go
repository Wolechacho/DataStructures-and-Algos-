package SubArrays

func SubarraysDivByK(nums []int, k int) int {
	hashTable := map[int]int{0: 1}
	accumulatedSum := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		accumulatedSum += nums[i]
		modulo := accumulatedSum % k
		if modulo < 0 {
			modulo += k
		}
		hashTable[modulo]++
		if _, exist := hashTable[modulo]; exist {
			count += hashTable[modulo] - 1
		}

	}

	return count
}
