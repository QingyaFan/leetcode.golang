package easy

type treeNode struct {
	Val int
	Left *treeNode
	Right *treeNode
}

func maxDepth(root *treeNode) int {

	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDep := maxDepth(root.Left)
	rightDep := maxDepth(root.Right)
	if leftDep > rightDep {
		return 1 + leftDep
	}

	return 1 + rightDep
}
