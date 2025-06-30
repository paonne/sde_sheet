import "slices"

func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	slices.Sort(nums)
	helper(0, []int{}, nums, &res)
	return res
}

func helper(ind int, curr []int, nums []int, res *[][]int) {
	if ind == len(nums) {
		curr_copy := make([]int, len(curr))
		copy(curr_copy, curr)
		*res = append(*res, curr_copy)
		return
	}

	curr = append(curr, nums[ind])
	helper(ind+1, curr, nums, res)

	curr = curr[:len(curr)-1]
	for ind+1 < len(nums) && nums[ind] == nums[ind+1] {
		ind += 1
	}
	helper(ind+1, curr, nums, res)
}
