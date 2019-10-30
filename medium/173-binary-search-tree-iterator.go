package medium

/**
 * 实现一个二叉搜索树迭代器。你将使用二叉搜索树的根节点初始化迭代器。
 * 调用 next() 将返回二叉搜索树中的下一个最小的数。
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// BSTIterator bst iterator struct
type BSTIterator struct {
	root  *TreeNode
	stack []*TreeNode
}

// Constructor constructor
func Constructor(root *TreeNode) BSTIterator {
	bstIterator := BSTIterator{root: root}

	// 关键：BST的特性，最小值在左子树的最左叶子结点
	for root != nil {
		bstIterator.stack = append(bstIterator.stack, root)
		root = root.Left
	}

	return bstIterator
}

// Next 返回BST中最小值
func (iter *BSTIterator) Next() int {
	// 关键：取出最左节点，并将该节点的左子树push到栈中
	cur := iter.stack[len(iter.stack)-1]
	iter.stack = iter.stack[:len(iter.stack)-1]
	root := cur.Right
	for root != nil {
		iter.stack = append(iter.stack, root)
		root = root.Left
	}

	return cur.Val
}

// HasNext 是否有最小值
func (iter *BSTIterator) HasNext() bool {
	return len(iter.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
