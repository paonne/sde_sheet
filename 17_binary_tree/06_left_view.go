func LeftView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	var curr *TreeNode
	for len(queue) > 0 {

		res = append(res, queue[0].Val)

		for range len(queue) {
			curr := queue[0]
			queue := queue[1:]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}
	return res
}
