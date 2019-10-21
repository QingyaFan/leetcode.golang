package medium

// swapPairs 给定一个链表，两两交换其中相邻的节点，
// 并返回交换后的链表
// 主要的考察点应该是交换链表的相邻节点
func swapPairs(head *ListNode) *ListNode {
	var dummyHead ListNode
	dummyHead.Next = head
	tmp := &dummyHead
	for tmp.Next != nil && tmp.Next.Next != nil {
		left, right := tmp.Next, tmp.Next.Next
		tmp.Next = right
		left.Next = right.Next
		right.Next = left
		tmp = left
	}

	return dummyHead.Next
}
