package Arrays

func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	n := len(arr)

	for i := 0; i < len(arr); i++ {
		end := i + 1
		start := n - i
		total := (start * end)

		odd := total / 2
		if (total % 2) == 1 {
			odd++
		}

		sum += (odd) * arr[i]
	}
	return sum
}
