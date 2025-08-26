func uniquePaths(m int, n int) int {
	type Key struct {
		r, c int
	}
	var helper func(r, c int) int
	dp := map[Key]int{}

	helper = func(r, c int) int {
		if val, ok := dp[Key{r, c}]; ok {
			return val
		}
		if r == m || c == n {
			return 0
		} else if r == m-1 && c == n-1 {
			return 1
		}
		dp[Key{r, c}] = helper(r, c+1) + helper(r+1, c)
		return dp[Key{r, c}]
	}
	return helper(0, 0)
}