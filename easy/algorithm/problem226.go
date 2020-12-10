package algorithm

func invertTree(root *TreeNode) *TreeNode {
	if root != nil && root.Left != nil {
		root.Right, root.Left = invertTree(root.Left), invertTree(root.Left)
	}
	return root
}
