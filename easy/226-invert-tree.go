package easy

// invertTree invert a tree
// golang语言的交换顺序非常简洁
// 使用递归和栈各有哪些优劣 ? TODO:
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}
