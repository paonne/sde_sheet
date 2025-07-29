package main

import "math"

func minPathSum(grid [][]int) int {
	type Key struct {
		r int
		c int
	}
	var helper func(r, c int) int
	dp := map[Key]int{}

	helper = func(r, c int) int {
		if val, ok := dp[Key{r, c}]; ok {
			return val
		}
		if r == len(grid)-1 && c == len(grid[0])-1 {
			return grid[r][c]
		}
		if r == len(grid) || c == len(grid[0]) {
			return math.MaxInt
		}
		p1 := helper(r+1, c)
		p2 := helper(r, c+1)
		dp[Key{r, c}] = grid[r][c] + min(p1, p2)
		return dp[Key{r, c}]
	}
	return helper(0, 0)
}
