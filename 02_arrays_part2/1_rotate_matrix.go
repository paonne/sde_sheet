import "slices"

func rotate(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	for i := range n {
		for j := i + 1; j < m; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := range n {
		slices.Reverse(matrix[i])
	}
}