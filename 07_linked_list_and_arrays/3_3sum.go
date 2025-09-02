import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	pairs := map[[3]int]struct{}{}
	n := len(nums)
	for i := range n {
		left, right := i+1, n-1
		for left < right {
			curr_sum := nums[i] + nums[left] + nums[right]
			if curr_sum == 0 {
				pairs[[3]int{nums[i], nums[left], nums[right]}] = struct{}{}
				left += 1
				right -= 1
			} else if curr_sum < 0 {
				left += 1
			} else {
				right -= 1
			}
		}
	}
	res := [][]int{}
	for pair, _ := range pairs {
		res = append(res, pair[:])
	}
	return res
}
