func findTarget(root *TreeNode, k int) bool {
	return helper(root, map[int]bool{}, k)
}

func helper(node *TreeNode, seen map[int]bool, k int) bool {
	if node == nil {
		return false
	}
	if seen[k-node.Val] {
		return true
	}
	seen[node.Val] = true
	return helper(node.Left, seen, k) || helper(node.Right, seen, k)
}