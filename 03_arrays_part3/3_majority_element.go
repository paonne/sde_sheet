func majorityElement(nums []int) int {
	res, count := nums[0], 1
	for _, num := range nums[1:] {
		if count == 0 {
			res = num
			count = 1
		} else if num == res {
			count += 1
		} else {
			count -= 1
		}
	}
	return res
}