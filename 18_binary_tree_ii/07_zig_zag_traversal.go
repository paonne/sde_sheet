func zigzagLevelOrder(root *TreeNode) [][]int {

	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	zig_zag := 1

	for len(queue) > 0 {

		n := len(queue)
		curr := []int{}

		for range n {
			top := queue[0]
			queue = queue[1:]
			curr = append(curr, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
		if zig_zag == -1 {
			slices.Reverse(curr)
		}
		res = append(res, curr)
		zig_zag *= -1
	}
	return res
}
