func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(node *TreeNode, mini, maxi int) bool {

	if node == nil {
		return true
	}
	if !(mini < node.Val) || !(node.Val < maxi) {
		return false
	}
	return helper(node.Left, mini, node.Val) && helper(node.Right, node.Val, maxi)
}