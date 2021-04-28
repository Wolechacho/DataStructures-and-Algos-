package Arrays

func UniqueSumToZero(num int) []int {
	arr := make([]int, num)

	if num == 1 {
		arr[0] = 0
	} else {
		if num%2 == 0 {
			for i := 1; i < num; i += 2 {
				arr[i-1] = -i
				arr[i] = i
			}
		} else {
			for i := 1; i < num; i += 2 {
				arr[i] = i
				arr[i+1] = -i
			}
		}
	}

	return arr
}
