package easy

// mergeTwoLists 将两个有序链表合并为一个新的有序链表并返回。
// 新链表是通过拼接给定的两个链表的所有节点组成的。
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 这里是关键，记录头之前的位置
	// 最后返回resList.Next即可
	// 让resList的影分身tmpList去冲锋陷阵
	var resList ListNode
	tmpList := &resList

	// 合并l1和l2
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			tmpList.Next = l2
			l2 = l2.Next
		} else {
			tmpList.Next = l1
			l1 = l1.Next
		}
		tmpList = tmpList.Next
	}

	// 合并l1或l2剩余的node
	if l1 != nil {
		tmpList.Next = l1
	}

	if l2 != nil {
		tmpList.Next = l2
	}

	return resList.Next
}
