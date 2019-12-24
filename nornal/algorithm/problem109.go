package algorithm

var a *ListNode

//SortedListToBST s
func SortedListToBST(head *ListNode) *TreeNode {
	len := 0
	next := head
	for next != nil {
		len++
		next = next.Next
	}
	a = head
	root := convertListToBST(head, 0, len-1)
	return root
}

func convertListToBST(head *ListNode, l int, r int) *TreeNode {
	if l > r {
		return nil
	}
	mid := (r + l) / 2
	var root *TreeNode
	left := convertListToBST(head, l, mid-1)
	root = &TreeNode{
		Val: head.Val,
	}
	head = head.Next
	root.Left = left
	root.Right = convertListToBST(head, mid+1, r)
	return root
}

//SortedListToBST s
func SortedListToBST1(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	nodes := make([]int, 0)
	for head != nil {
		nodes = append(nodes, head.Val)
		head = head.Next
	}
	root := convertToTree1(nodes)
	return root
}

func convertToTree1(nodes []int) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	mid := len(nodes) / 2
	root := &TreeNode{
		Val: nodes[mid],
	}
	root.Left = convertToTree1(nodes[0:mid])
	root.Right = convertToTree1(nodes[mid+1:])
	return root
}
