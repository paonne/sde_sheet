func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	curr := root

	for curr != nil {
		if curr.Val < p.Val && curr.Val < q.Val {
			curr = curr.Right
		} else if curr.Val > p.Val && curr.Val > q.Val {
			curr = curr.Left
		} else {
			return curr
		}
	}
	return nil
}
