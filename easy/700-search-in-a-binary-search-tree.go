package easy

// searchBST 递归的空间复杂度 O(logn)
func searchBST(root *TreeNode, val int) *TreeNode {

	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			return searchBST(root.Left, val)
		} else if root.Val < val {
			return searchBST(root.Right, val)
		}
	}

	return nil
}

// searchBST2 空间复杂度好一些 O(1)
func searchBST2(root *TreeNode, val int) *TreeNode {

	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		}
	}

	return nil
}
