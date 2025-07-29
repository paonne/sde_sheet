package main

import "math"

func coinChange(coins []int, amount int) int {
	type Key struct {
		amt int
		ind int
	}
	dp := map[Key]int{}
	var helper func(amt, ind int) int

	helper = func(amt, ind int) int {
		if val, ok := dp[Key{amt, ind}]; ok {
			return val
		}
		if amt == 0 {
			return 0
		}
		if ind == len(coins) {
			return math.MaxInt32
		}

		p1 := helper(amt, ind+1)
		p2 := math.MaxInt32
		if amt >= coins[ind] {
			p2 = 1 + helper(amt-coins[ind], ind)
		}
		dp[Key{amt, ind}] = min(p1, p2)
		return dp[Key{amt, ind}]
	}
	res := helper(amount, 0)
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
