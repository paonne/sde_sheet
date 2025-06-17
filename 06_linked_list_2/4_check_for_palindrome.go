func isPalindrome(head *ListNode) bool {

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var prev *ListNode
	curr := slow
	var nextnode *ListNode
	for curr != nil {
		nextnode = curr.Next
		curr.Next = prev
		prev = curr
		curr = nextnode
	}

	curr1, curr2 := head, prev
	for curr2 != nil {
		if curr1.Val != curr2.Val {
			return false
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}
	return true
}