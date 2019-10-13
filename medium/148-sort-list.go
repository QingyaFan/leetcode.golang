package medium

// sortList 在 O(nlogn) 时间复杂度和常数级空间复杂度下，
// 对链表进行排序。
//
// 排序：主要的排序算法时间复杂度为O(nlogn)的有
// 	堆排序 O(nlogn),O(1)
// 	快排排序 O(nlogn),O(logn)
// 	归并排序 O(nlogn),O(1)
//
// 空间复杂度是常数，只剩下堆排序和归并排序
// 待排序是链表，只有归并排序了
func sortList(head *ListNode) *ListNode {
	var dummyHead ListNode
	dummyHead.Next = head

	listLen := 0
	tmp := head
	// 这里不要判断 tmp.Next != nil
	// 因为有可能tmp是nil，将会引起panic
	for tmp != nil {
		listLen++
		tmp = tmp.Next
	}

	for size := 1; size < listLen; size *= 2 {
		cur := dummyHead.Next
		tail := &dummyHead

		for cur != nil {
			left := cur
			right := cut(left, size)
			cur = cut(right, size)
			tail.Next = merge(left, right)
			for tail.Next != nil {
				tail = tail.Next
			}
		}
	}

	return dummyHead.Next
}

// 从链表中切出size个node
// 返回
func cut(head *ListNode, size int) *ListNode {
	p := head
	for size > 1 && p != nil {
		p = p.Next
		size--
	}

	if p == nil {
		return nil
	}

	// 断开链表
	next := p.Next
	p.Next = nil

	return next
}

// 拼接两个链表
func merge(l1, l2 *ListNode) *ListNode {
	var dummyHead ListNode
	p := &dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val >= l2.Val {
			p.Next = l2
			l2 = l2.Next
		} else {
			p.Next = l1
			l1 = l1.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	} else if l2 != nil {
		p.Next = l2
	}

	return dummyHead.Next
}

// 此题需要反复复盘
// 三个要点
// 链表的cut、merge与dummyHead
// 参考：https://leetcode-cn.com/problems/sort-list/solution/148-pai-xu-lian-biao-bottom-to-up-o1-kong-jian-by-/
