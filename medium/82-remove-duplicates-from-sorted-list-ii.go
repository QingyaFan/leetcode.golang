package medium

// deleteDuplicates 给定一个排序链表，
// 删除所有含有重复数字的节点，只保留原始链表中没有重复出现的数字
// 双指针法 TODO: [1,1]不能通过测试，[1,2,2]不能通过测试
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var dummyHead ListNode
	dummyHead.Next = head

	slow := &dummyHead
	fast := head
	for fast != nil {
		if fast.Next != nil && fast.Val != fast.Next.Val {
			// 这里是难点，如果slow和fast相邻，slow移步fast
			// 如果不相邻，slow.Next = fast.Next
			if slow.Next == fast {
				slow = fast
			} else {
				slow.Next = fast.Next
			}
		}
		fast = fast.Next
	}

	return dummyHead.Next
}
