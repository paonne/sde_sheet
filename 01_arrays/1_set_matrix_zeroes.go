func setZeroes(matrix [][]int) {

	m, n := len(matrix), len(matrix[0])
	r, c := map[int]bool{}, map[int]bool{}

	for i := range m {
		for j := range n {
			if matrix[i][j] == 0 {
				r[i] = true
				c[j] = true
			}
		}
	}

	for i := range m {
		for j := range n {
			if r[i] == true || c[j] == true {
				matrix[i][j] = 0
			}
		}
	}

}
