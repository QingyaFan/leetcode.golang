package easy

// sortedArrayToBST tree always use recursive
func sortedArrayToBST(nums []int) *TreeNode {
	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}
	if numsLen == 1 {
		return &TreeNode{nums[0], nil, nil}
	}

	var root TreeNode
	middle := numsLen / 2

	return &TreeNode{nums[middle], sortedArrayToBST(nums[:middle]), sortedArrayToBST(nums[middle+1:])}
}
