var res int

func maxPathSum(root *TreeNode) int {
	res = root.Val
	helper(root)
	return res
}

func helper(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := max(0, helper(node.Left))
	r := max(0, helper(node.Right))
	res = max(res, node.Val+l+r)

	return node.Val + max(l, r)
}