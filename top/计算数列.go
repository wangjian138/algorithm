package main

func determinant(matrix [][]int) int {
	size := len(matrix)
	if size == 1 {
		return matrix[0][0]
	} else if size == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	matrixRes := 0
	sign := 1
	for i := 0; i < size; i++ {
		minor := make([][]int, size-1)
		for j := 0; j < size-1; j++ {
			minor[j] = make([]int, size-1)
			for k := 0; k < size-1; k++ {
				if k < i {
					minor[j][k] = matrix[j+1][k]
				} else {
					minor[j][k] = matrix[j+1][k+1]
				}
			}
		}
		matrixRes += sign * matrix[0][i] * determinant(minor)
		sign *= -1
	}
	return matrixRes
}
