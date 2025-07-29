package main

func canPartition(nums []int) bool {

	type Key struct {
		curr_sum int
		ind      int
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2

	dp := map[Key]bool{}
	var helper func(curr_sum, ind int) bool

	helper = func(curr_sum, ind int) bool {
		if val, ok := dp[Key{curr_sum, ind}]; ok {
			return val
		}
		if curr_sum == target {
			return true
		}
		if curr_sum > target {
			return false
		}
		if ind == len(nums) {
			return false
		}

		p1 := helper(curr_sum, ind+1)
		p2 := helper(curr_sum+nums[ind], ind+1)
		dp[Key{curr_sum, ind}] = p1 || p2
		return dp[Key{curr_sum, ind}]
	}
	return helper(0, 0)
}
