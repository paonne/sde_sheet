func removeDuplicates(nums []int) int {
	curr := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		nums[curr] = nums[i]
		curr += 1
	}
	return curr
}