package easy

// levelOrderBottom 给定一个二叉树，
// 返回其节点值自底向上的层次遍历
// 树，广度优先遍历
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int

	var cur []*TreeNode
	cur = append(cur, root)
	for len(cur) != 0 {
		var currentLayerVal []int
		var nextLayerNode []*TreeNode
		for _, node := range cur {
			if node != nil {
				currentLayerVal = append(currentLayerVal, node.Val)
				nextLayerNode = append(nextLayerNode, node.Left)
				nextLayerNode = append(nextLayerNode, node.Right)
			}
		}

		// 如何保证每层的数组是自底向上存储的？
		// 需要实现insertBeforeHead
		if len(currentLayerVal) != 0 {
			res = insertBeforeHead(res, currentLayerVal)
		}
		cur = nextLayerNode
	}

	return res
}

func insertBeforeHead(slice [][]int, val []int) [][]int {
	var cur [][]int
	cur = append(cur, val)
	cur = append(cur, slice...)

	return cur
}
