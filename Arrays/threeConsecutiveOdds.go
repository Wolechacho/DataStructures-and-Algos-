package Arrays

func threeConsecutiveOdds(arr []int) bool {
	n := 3
	for _, num := range arr {
		if num%2 != 0 {
			n--
		} else {
			n = 3
		}

		if n == 0 {
			return true
		}
	}
	return false
}

func threeConsecutiveOdds1(arr []int) bool {
    oddCount := 0
    for i :=0; i< len(arr);i++{
        if arr[i] % 2 != 0{
            oddCount++
            
            if oddCount == 3{
                return true
            }
                  
        }else{
            oddCount = 0
        }
    }
    return false
}
