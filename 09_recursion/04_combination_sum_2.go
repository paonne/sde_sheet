import "slices"

var res [][]int

func combinationSum2(candidates []int, target int) [][]int {
	res = [][]int{}
	slices.Sort(candidates)
	helper(0, 0, target, []int{}, candidates)
	return res
}

func helper(ind, curr_sum, target int, curr, candidates []int) {

	if curr_sum == target {
		var curr_copy = make([]int, len(curr))
		copy(curr_copy, curr)
		res = append(res, curr_copy)
		return
	}

	if ind == len(candidates) || curr_sum > target {
		return
	}

	curr = append(curr, candidates[ind])
	helper(ind+1, curr_sum+candidates[ind], target, curr, candidates)

	curr = curr[:len(curr)-1]
	for ind+1 < len(candidates) && candidates[ind+1] == candidates[ind] {
		ind += 1
	}
	helper(ind+1, curr_sum, target, curr, candidates)
}
