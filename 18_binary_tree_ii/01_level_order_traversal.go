func levelOrder(root *TreeNode) [][]int {

	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		curr := []int{}
		n := len(queue)

		for range n {
			node := queue[0]
			queue = queue[1:]
			curr = append(curr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, curr)
	}
	return res
}