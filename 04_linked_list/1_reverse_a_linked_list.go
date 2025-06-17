func reverseList(head *ListNode) *ListNode {

	var prev *ListNode = nil
	var nextnode *ListNode
	curr := head

	for curr != nil {
		nextnode = curr.Next
		curr.Next = prev
		prev = curr
		curr = nextnode
	}
	return prev
}
