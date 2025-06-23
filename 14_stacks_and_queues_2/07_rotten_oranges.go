import "slices"

func orangesRotting(grid [][]int) int {

	n, m := len(grid), len(grid[0])
	queue := [][2]int{}
	bad_oranges := 0

	for i := range n {
		for j := range m {
			if grid[i][j] == 2 {
				bad_oranges += 1
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	res := -1

	for len(queue) > 0 {
		for range len(queue) {
			curr_i, curr_j := queue[0][0], queue[0][1]
			queue = queue[1:]
			x_move, y_move := [4]int{-1, 0, 1, 0}, [4]int{0, -1, 0, 1}
			for i := range 4 {
				if is_valid(curr_i+x_move[i], curr_j+y_move[i], n, m) && grid[curr_i+x_move[i]][curr_j+y_move[i]] == 1 {
					grid[curr_i+x_move[i]][curr_j+y_move[i]] = 2
					queue = append(queue, [2]int{curr_i + x_move[i], curr_j + y_move[i]})
				}
			}
		}
		res += 1
	}

	for _, item := range grid {
		if slices.Contains(item, 1) {
			return -1
		}
	}

	if bad_oranges == 0 {
		return 0
	}

	return res
}

func is_valid(i, j, n, m int) bool {
	return (0 <= i) && (i < n) && (0 <= j) && (j < m)
}
