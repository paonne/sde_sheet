func bstFromPreorder(preorder []int) *TreeNode {
	inorder := slices.Clone(preorder)
	slices.Sort(inorder)
	return buildTree(preorder, inorder)
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}
	root_val := preorder[0]
	preorder = preorder[1:]
	root_inorder_ind := slices.Index(inorder, root_val)

	root := TreeNode{Val: root_val}
	root.Left = buildTree(preorder[:root_inorder_ind], inorder[:root_inorder_ind])
	root.Right = buildTree(preorder[root_inorder_ind:], inorder[root_inorder_ind+1:])
	return &root
}
