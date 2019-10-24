package medium

// rotateRight 给定一个链表，旋转链表，
// 将链表每个节点向右移动 k 个位置，其中 k 是非负数
func rotateRight(head *ListNode, k int) *ListNode {
	var dummyHead ListNode
	dummyHead.Next = head

	listLen := 0
	cur := &dummyHead
	for cur.Next != nil {
		cur = cur.Next
		listLen++
	}

	if listLen == 0 {
		return nil
	}

	tail := cur
	tail.Next = head

	shiftNum := k % listLen

	toTailStep := listLen - shiftNum
	cur = &dummyHead
	for toTailStep > 0 {
		cur = cur.Next
		toTailStep--
	}

	if cur.Next == nil {
		return dummyHead.Next
	}

	dummyHead.Next = cur.Next
	cur.Next = nil

	return dummyHead.Next
}
