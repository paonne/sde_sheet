func detectCycle(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			curr := head
			for curr != slow {
				slow = slow.Next
				curr = curr.Next
			}
			return curr
		}
	}
	return nil
}