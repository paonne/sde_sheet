import "slices"

func nse(nums []int) []int {

	stack, res := []int{}, []int{}
	n := len(nums)

	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res = append(res, n)
		} else {
			res = append(res, stack[len(stack)-1])
		}
		stack = append(stack, i)
	}
	slices.Reverse(res)
	return res
}

func pse(nums []int) []int {

	stack, res := []int{}, []int{}
	n := len(nums)

	for i := range n {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res = append(res, -1)
		} else {
			res = append(res, stack[len(stack)-1])
		}
		stack = append(stack, i)
	}
	return res
}

func largestRectangleArea(heights []int) int {
	nse_slice := nse(heights)
	pse_slice := pse(heights)

	res := 0
	for i := range len(heights) {
		res = max(res, heights[i]*(nse_slice[i]-pse_slice[i]-1))
	}
	return res
}
