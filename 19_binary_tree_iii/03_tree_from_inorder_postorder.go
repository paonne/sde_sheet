func buildTree(inorder []int, postorder []int) *TreeNode {

	if len(postorder) == 0 {
		return nil
	}
	root_val := postorder[len(postorder)-1]
	postorder = postorder[:len(postorder)-1]
	root_inorder_ind := slices.Index(inorder, root_val)

	root := TreeNode{Val: root_val}
	root.Left = buildTree(inorder[:root_inorder_ind], postorder[:root_inorder_ind])
	root.Right = buildTree(inorder[root_inorder_ind+1:], postorder[root_inorder_ind:])
	return &root
}
