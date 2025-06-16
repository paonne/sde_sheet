func maxSubArray(nums []int) int {

	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	curr_sum, maxi_sum := nums[0], nums[0]
	for i := 1; i < n; i++ {
		curr_sum = max(nums[i], curr_sum+nums[i])
		maxi_sum = max(maxi_sum, curr_sum)
	}

	return maxi_sum

}