func removeNthFromEnd(head *ListNode, n int) *ListNode {

	fast := head
	for range n {
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}
	
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return head
}
