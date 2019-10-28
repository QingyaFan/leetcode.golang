package easy

import "math"

func isBalanced(root *TreeNode) bool {
	return getTreeDepth(root) != -1
}

// getTreeDepth 通过返回 -1 来阻断对已经判断
// 为不平衡对子树的父节点重复判断
func getTreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := getTreeDepth(root.Left)
	if leftDepth == -1 {
		return -1
	}

	rightDepth := getTreeDepth(root.Right)
	if rightDepth == -1 {
		return -1
	}

	if math.Abs(float64(leftDepth-rightDepth)) < 2 {
		if leftDepth > rightDepth {
			return leftDepth + 1
		} else {
			return rightDepth + 1
		}
	} else {
		return -1
	}
}
