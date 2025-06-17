func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, res)
	*res = append(*res, node.Val)
	dfs(node.Right, res)
}