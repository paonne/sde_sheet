func getIntersectionNode(headA, headB *ListNode) *ListNode {

	l1 := get_length_of_ll(headA)
	l2 := get_length_of_ll(headB)
	curr1, curr2 := headA, headB
	if l1 > l2 {
		for range l1 - l2 {
			curr1 = curr1.Next
		}
	} else {
		for range l2 - l1 {
			curr2 = curr2.Next
		}
	}

	for curr1 != nil && curr2 != nil {
		if curr1 == curr2 {
			return curr1
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}
	return nil
}

func get_length_of_ll(node *ListNode) int {
	res := 0
	for node != nil {
		res += 1
		node = node.Next
	}
	return res
}