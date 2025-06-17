func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	res := &ListNode{}
	curr := res
	remainder := 0

	for l1 != nil || l2 != nil || remainder != 0 {

		curr_sum := 0
		if l1 != nil {
			curr_sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curr_sum += l2.Val
			l2 = l2.Next
		}
		if remainder != 0 {
			curr_sum += remainder
		}

		curr.Next = &ListNode{curr_sum % 10, nil}
		curr = curr.Next
		remainder = curr_sum / 10
	}

	return res.Next
}
