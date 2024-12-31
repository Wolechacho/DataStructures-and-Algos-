package Arrays

func PerformStringShifts(s string,shifts [][]int)string{
	totalShift := 0
	for _, shift := range shifts{
		direction := shift[0]
		amount := shift[1]

		if direction == 1{
			totalShift -= amount
		}else{
			totalShift += amount
		}
	}

	length := len(s)
	totalShift %= length

	if totalShift < 0{
		totalShift += length
	}

	return s[totalShift:] + s[0:totalShift]
}