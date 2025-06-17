func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	dfs(node.Left, res)
	dfs(node.Right, res)
}