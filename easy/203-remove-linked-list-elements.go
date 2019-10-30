package easy

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	var dummyHead ListNode
	dummyHead.Next = head

	cur := &dummyHead
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}
