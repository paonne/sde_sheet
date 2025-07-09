package main

import "fmt"

func knapsack01(wt, val []int, W int) int {

	type Key struct {
		ind  int
		curr int
	}
	dp := map[Key]int{}

	var helper func(ind, curr int) int

	helper = func(ind, curr int) int {

		if val, ok := dp[Key{ind, curr}]; ok {
			return val
		}
		if ind == len(wt) {
			return 0
		}

		pick := 0
		if wt[ind] <= curr {
			pick = val[ind] + helper(ind+1, curr-wt[ind])
		}

		not_pick := helper(ind+1, curr)
		dp[Key{ind, curr}] = max(pick, not_pick)
		return dp[Key{ind, curr}]
	}

	return helper(0, W)
}

func main() {
	fmt.Println(knapsack01([]int{10, 20, 30}, []int{60, 100, 120}, 50) == 220)
	fmt.Println(knapsack01([]int{5, 4, 6, 3}, []int{10, 40, 30, 50}, 10) == 90)
	fmt.Println(knapsack01([]int{1, 2, 3, 8, 7, 4}, []int{20, 5, 10, 40, 15, 25}, 10) == 60)
}
