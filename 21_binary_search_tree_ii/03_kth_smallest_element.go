var k_dup int
var res int

func kthSmallest(root *TreeNode, k int) int {
	k_dup = k
	inorder(root)
	return res
}

func inorder(node *TreeNode) {
	if node == nil {
		return
	}
	inorder(node.Left)
	k_dup -= 1
	if k_dup == 0 {
		res = node.Val
		return
	}
	inorder(node.Right)
}