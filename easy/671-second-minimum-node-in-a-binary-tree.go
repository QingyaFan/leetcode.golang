package easy

// findSecondMinimumValue 给定一个非空特殊的二叉树，每个节点都是正数，
// 并且每个节点的子节点数量只能为 2 或 0。
// 如果一个节点有两个子节点的话，那么这个节点的值不大于它的子节点的值。
// 给出这样的一个二叉树，你需要输出所有节点中的第二小的值。
// 如果第二小的值不存在的话，输出 -1 。
func findSecondMinimumValue(root *TreeNode) int {
	if root.Left == nil {
		return -1
	}
	var minimum, secondMinimum int
	// 非空，无需考虑root为nil的情况
	minimum = root.Val
	secondMinimum = root.Val
	var cur []*TreeNode
	cur = append(cur, root.Left, root.Right)
	for len(cur) != 0 {
		var currentLayerVal []int
		for _, node := range cur {
			if node.Val < secondMinimum
			if node != nil && node.Left != nil {
				currentLayerVal = append(currentLayerVal, node.Left)
				currentLayerVal = append(currentLayerVal, node.Right)
			}
		}
		cur = currentLayerVal
	}
}
