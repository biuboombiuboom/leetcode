package algorithm

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	return dfs(root, head) || isSubPath(head, root.Left) || isSubPath(head, root.Right)

}

func dfs(root *TreeNode, head *ListNode) bool {
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	if head.Val != root.Val {
		return false
	}
	return dfs(root.Left, head.Next) || dfs(root.Right, head.Next)
}

func buildListNode(nums []int) *ListNode {
	begin := &ListNode{}
	pre := begin
	for i := 0; i < len(nums); i++ {
		pre.Next = &ListNode{
			Val: nums[i],
		}
		pre = pre.Next
	}
	return begin.Next
}

func buildTreeNode(nums []int) *ListNode {
	return nil
}
