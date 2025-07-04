func sortedArrayToBST(nums []int) *TreeNode {

	n := len(nums)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{Val: nums[0]}
	}

	mid := n/2
	node := &TreeNode{Val: nums[mid], Left: sortedArrayToBST(nums[:mid]), Right: sortedArrayToBST(nums[mid+1:])}
	return node
}
