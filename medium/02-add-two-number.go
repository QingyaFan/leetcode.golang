package medium

// AddTwoNumber 链表的遍历
// 两个链表长度不一样
// 某一位满十进位
func AddTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {

	l3Head := &ListNode{0, nil}
	l3 := l3Head
	over := 0

	for l1 != nil || l2 != nil {
		head := l1.Val + l2.Val + over
		over = head / 10

		l3.Val = head % 10

		l3.Next = &ListNode{0, nil}

		l1, l2, l3 = l1.Next, l2.Next, l3.Next
	}

	if over != 0 {
		l3.Next = &ListNode{over, nil}
	}

	return l3Head
}
