package easy

// invertTree invert a tree
// golang语言的交换顺序非常简洁
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}
