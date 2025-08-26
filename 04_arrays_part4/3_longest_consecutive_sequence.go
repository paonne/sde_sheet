import "slices"

func longestConsecutive(nums []int) int {
	slices.Sort(nums)
	n := len(nums)
	if n <= 1 {
		return n
	}
	res, count, prev := 1, 1, nums[0]
	for _, num := range nums[1:] {
		if num == prev+1 {
			count += 1
			prev = num
		} else if num == prev {
			continue
		} else {
			prev = num
			res = max(res, count)
			count = 1
		}
	}
	return max(res, count)
}
