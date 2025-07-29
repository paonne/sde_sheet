package main

import "math"

func minCost(n int, cuts []int) int {

	type Key struct {
		l int
		r int
	}
	dp := map[Key]int{}
	var helper func(l, r int) int

	helper = func(l, r int) int {
		if val, ok := dp[Key{l, r}]; ok {
			return val
		}
		res := math.MaxInt32
		for _, cut := range cuts {
			if l < cut && cut < r {
				res = min(res, r-l+helper(l, cut)+helper(cut, r))
			}
		}
		if res == math.MaxInt32 {
			dp[Key{l, r}] = 0
			return 0
		}
		dp[Key{l, r}] = res
		return res
	}
	return helper(0, n)

}
