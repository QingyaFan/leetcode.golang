package easy

// deleteNode
// 只给定node，所以只能操作node
// 注意不能用将前node的Next指向下一个node的操作
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
