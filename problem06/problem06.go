package problem06

func Transpose(matrix [][]int) [][]int {
	if len(matrix) < 1 {
		return matrix
	}
	transposedMatrix := make([][]int, len(matrix[0]))
	for row := 0; row < len(transposedMatrix); row++ {
		transposedMatrix[row] = make([]int, len(matrix))
	}

	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			transposedMatrix[i][j] = matrix[j][i]
		}
	}
	return transposedMatrix
}
