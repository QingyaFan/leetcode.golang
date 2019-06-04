package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {

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
