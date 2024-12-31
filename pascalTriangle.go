package main

func pascalTriangle(row int) [][]int {
	var matrix = [][]int{}
	if row <= 0 {
		return matrix
	}
	f := []int{1}
	matrix = append(matrix, f)
	for i := 1; i < row; i++ {
		arr := []int{}
		arr = append(arr, 1)
		for j := 1; j < i; j++ {
			val := matrix[i-1][j-1] + matrix[i-1][j]
			arr = append(arr, val)
		}
		arr = append(arr, 1)
		matrix = append(matrix, arr)
	}

	return matrix
}
