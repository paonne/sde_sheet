var res bool

func isBalanced(root *TreeNode) bool {
	res = true
	helper(root)
	return res
}

func helper(node *TreeNode) int {
	if node == nil {
		return 0
	}
	lh := helper(node.Left)
	rh := helper(node.Right)

	if abs(lh-rh) > 1 {
		res = false
	}
	return 1 + max(lh, rh)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
