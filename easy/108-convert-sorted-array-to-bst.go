package easy

// sortedArrayToBST 排序数组转换为二叉搜索树
// 看到树，就是递归
func sortedArrayToBST(nums []int) *TreeNode {
	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}

	if numsLen == 1 {
		return &TreeNode{nums[0], nil, nil}
	}

	middle := numsLen / 2

	return &TreeNode{nums[middle], sortedArrayToBST(nums[:middle]), sortedArrayToBST(nums[middle+1:])}
}
