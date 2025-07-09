func minDistance(word1 string, word2 string) int {

	type Key struct {
		i int
		j int
	}
	dp := map[Key]int{}
	var helper func(ind1, ind2 int) int

	helper = func(ind1, ind2 int) int {
		if val, ok := dp[Key{ind1, ind2}]; ok {
			return val
		}
		if ind1 == len(word1) && ind2 == len(word2) {
			return 0
		}
		if ind1 == len(word1) {
			return len(word2) - ind2
		}
		if ind2 == len(word2) {
			return len(word1) - ind1
		}
		if word1[ind1] == word2[ind2] {
			dp[Key{ind1, ind2}] = helper(ind1+1, ind2+1)
			return dp[Key{ind1, ind2}]
		}

		insert := helper(ind1, ind2+1)
		delete := helper(ind1+1, ind2)
		replace := helper(ind1+1, ind2+1)
		dp[Key{ind1, ind2}] = min(insert, delete, replace) + 1
		return dp[Key{ind1, ind2}]
	}
	return helper(0, 0)
}