type BSTIterator struct {
	InorderArr []int
	CurrInd    int
}

func Constructor(root *TreeNode) BSTIterator {
	obj := BSTIterator{CurrInd: -1}
	inorder(root, &obj)
	return obj
}

func inorder(node *TreeNode, obj *BSTIterator) {
	if node == nil {
		return
	}
	inorder(node.Left, obj)
	obj.InorderArr = append(obj.InorderArr, node.Val)
	inorder(node.Right, obj)
}

func (this *BSTIterator) Next() int {
	this.CurrInd += 1
	return this.InorderArr[this.CurrInd]
}

func (this *BSTIterator) HasNext() bool {
	return this.CurrInd+1 < len(this.InorderArr)
}