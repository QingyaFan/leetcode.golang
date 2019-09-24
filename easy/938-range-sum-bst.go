package easy

// rangeSumBST
// 树是一种自引用的数据结构，所以很自然想到递归
// 否则使用栈存储中间过程的值
// 为什么使用深度优先，而不是广度优先 ？ TODO:
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}
	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}

	return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
}
