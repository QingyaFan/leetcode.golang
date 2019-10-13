package easy

// deleteDuplicates 删除链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	var dummyHead ListNode
	dummyHead.Next = head
	p := head
	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return dummyHead.Next
}
