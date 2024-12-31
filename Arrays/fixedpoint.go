package Arrays

func fixedPoint(A []int) int {
	res := -1
	for i, v := range A {
		if i == v {
			res = i
			break
		}
	}
	return res
}
