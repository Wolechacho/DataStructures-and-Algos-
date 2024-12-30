package SubArrays

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

func sumOddLengthSubarrays1(arr []int) int {
    oddsum := 0
    for i :=0;i < len(arr);i++{
        sum := 0
        counter := 0
        for j :=i;j < len(arr);j++{
            sum += arr[j]
            counter++   
            if counter % 2 != 0{
                oddsum += sum
            }        
        }
    }
    return oddsum
}
