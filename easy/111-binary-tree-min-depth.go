package easy

func minDepth(root *treeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}

	if root.Left != nil && root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	leftDep := minDepth(root.Left)
	rightDep := minDepth(root.Right)
	if leftDep < rightDep {
		return 1 + leftDep
	}

	return 1 + rightDep
}
