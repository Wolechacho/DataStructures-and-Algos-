package main

func ShiftString(S string, shift [][]int) string {

	totalShift := 0
	//calculate the net total shift for left and right
	for _, sf := range shift {
		direction := sf[0]
		amount := sf[1]

		//decrement the amount if right
		if direction == 1 {
			amount = -amount
		}
		totalShift += amount
	}

	length := len(S)
	totalShift %= length

	if totalShift < 0 {
		totalShift += length
	}

	shifted1 := S[totalShift:]
	shifted2 := S[0:totalShift]

	return shifted1 + shifted2
}
