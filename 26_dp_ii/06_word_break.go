package main

func wordBreak(s string, words []string) bool {

	wordDict := map[string]bool{}
	for _, word := range words {
		wordDict[word] = true
	}
	dp := map[int]bool{}
	var helper func(ind int) bool

	helper = func(ind int) bool {
		if val, ok := dp[ind]; ok {
			return val
		}
		if ind == len(s) {
			return true
		}
		for i := ind + 1; i <= len(s); i++ {
			_, ok := wordDict[s[ind:i]]
			if ok && helper(i) {
				dp[ind] = true
				return true
			}
		}
		dp[ind] = false
		return false
	}
	return helper(0)
}
