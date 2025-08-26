func searchMatrix(matrix [][]int, target int) bool {
	n, m := len(matrix), len(matrix[0])
	low, high := 0, n*m-1
	var mid, r, c int
	for low <= high {
		mid = (low + high) / 2
		r, c = mid/m, mid%m
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}