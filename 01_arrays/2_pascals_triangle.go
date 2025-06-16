func generate(numRows int) [][]int {

	res := [][]int{{1}}
	if numRows == 1 {
		return res
	}
	for i := 2; i < numRows+1; i++ {
		res = append(res, next_row(res[len(res)-1]))
	}
	return res
}

func next_row(prev_row []int) []int {
	res := []int{1}
	for i := 1; i < len(prev_row); i++ {
		res = append(res, prev_row[i]+prev_row[i-1])
	}
	res = append(res, 1)
	return res
}
