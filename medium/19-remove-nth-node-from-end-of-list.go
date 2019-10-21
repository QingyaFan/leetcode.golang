package medium

// removeNthFromEnd 给定一个链表，
// 删除链表的倒数第 n 个节点，并且返回链表的头结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummyHead ListNode
	dummyHead.Next = head

	// 维护间隔为 n 的指针
	// 注意p1和p2都是指向dummyHead，而不是head
	// 可防止遍历开始时 p2 = nil 出现的问题
	p1, p2 := &dummyHead, &dummyHead
	for i := 0; i < n; i++ {
		p2 = p2.Next
	}

	// 遍历到链表尾部
	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	// 题中说明n保证是有效的，否则
	// 需要判断p1现在是否指向head

	// 删除p1指向的节点
	p1.Next = p1.Next.Next

	return dummyHead.Next
}
