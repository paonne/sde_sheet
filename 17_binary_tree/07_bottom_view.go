func bottomView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	view := map[int]int{}
	queue := []Pair{{0, root}}

	for len(queue) > 0 {
		top_item := queue[len(queue)-1]
		view[top_item.Level] = top_item.Node.Val
		queue = queue[:len(queue)-1]

		if top_item.Node.Left != nil {
			queue = append(queue, Pair{top_item.Level - 1, top_item.Node.Left})
		}
		if top_item.Node.Right != nil {
			queue = append(queue, Pair{top_item.Level + 1, top_item.Node.Right})
		}
	}

	level := []int{}
	for k := range view {
		level = append(level, k)
	}
	slices.Sort(level)

	for _, i := range level {
		res = append(res, view[i])
	}
	return res
}