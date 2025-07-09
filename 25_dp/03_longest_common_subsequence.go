func longestCommonSubsequence(text1 string, text2 string) int {

	type Key struct {
		i int
		j int
	}

	dp := map[Key]int{}
	var helper func(i, j int) int

	helper = func(i, j int) int {
		if val, ok := dp[Key{i, j}]; ok {
			return val
		}
		if i == len(text1) || j == len(text2) {
			return 0
		}
		res := 0
		if text1[i] == text2[j] {
			res = helper(i+1, j+1) + 1
		} else {
			res = max(helper(i+1, j), helper(i, j+1))
		}

		dp[Key{i, j}] = res
		return res
	}
	return helper(0, 0)
}