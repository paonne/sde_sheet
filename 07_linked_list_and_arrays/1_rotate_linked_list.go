func rotateRight(head *ListNode, k int) *ListNode {
	curr := head
	n := 0
	for curr != nil {
		curr = curr.Next
		n += 1
	}
	if n == 0 {
		return head
	}
	k = k % n
	if k == 0 {
		return head
	}
	fast := head
	for range k {
		fast = fast.Next
	}
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	res := slow.Next
	fast.Next = head
	slow.Next = nil
	return res
}
