import "slices"

func trap(heights []int) int {
	n := len(heights)
	pre_max, suff_max := []int{}, []int{}
	curr_pre_max, curr_suff_max := 0, 0
	for i := range n {
		curr_pre_max = max(curr_pre_max, heights[i])
		pre_max = append(pre_max, curr_pre_max)

		curr_suff_max = max(curr_suff_max, heights[n-i-1])
		suff_max = append(suff_max, curr_suff_max)
	}
	slices.Reverse(suff_max)

	res := 0
	for i := range n {
		res += min(pre_max[i], suff_max[i]) - heights[i]
	}
	return res
}
