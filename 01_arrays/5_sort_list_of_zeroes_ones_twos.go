func sortColors(nums []int) {

	l, r, c := 0, len(nums)-1, 0
	for c <= r {
		if nums[c] == 0 {
			nums[c], nums[l] = nums[l], nums[c]
			l += 1
			c += 1
		} else if nums[c] == 1 {
			c += 1
		} else {
			nums[c], nums[r] = nums[r], nums[c]
			r -= 1
		}
	}
}