import "slices"

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	pairs := map[[4]int]struct{}{}
	n := len(nums)
	for i := range n {
		for j := i + 1; j < n; j++ {
			left, right := j+1, n-1
			req := target - nums[i] - nums[j]
			for left < right {
				if nums[left]+nums[right] < req {
					left += 1
				} else if nums[left]+nums[right] > req {
					right -= 1
				} else {
					pairs[[4]int{nums[i], nums[j], nums[left], nums[right]}] = struct{}{}
					left += 1
					right -= 1
				}
			}
		}
	}
	res := [][]int{}
	for pair, _ := range pairs {
		res = append(res, pair[:])
	}
	return res
}