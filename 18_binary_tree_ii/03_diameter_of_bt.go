var res int

func diameterOfBinaryTree(root *TreeNode) int {
	res = 0
	helper(root)
	return res
}

func helper(node *TreeNode) int {
	if node == nil {
		return 0
	}
	lh := helper(node.Left)
	rh := helper(node.Right)

	res = max(res, lh+rh)
	return 1 + max(lh, rh)
}