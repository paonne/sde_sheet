func findMaxConsecutiveOnes(nums []int) int {
	res, curr := 0, 0
	for _, num := range nums {
		if num == 1 {
			curr += 1
		} else {
			res = max(res, curr)
			curr = 0
		}
	}
	return max(res, curr)
}