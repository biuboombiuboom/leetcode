package algorithm

//ReorderList R
func ReorderList(head *ListNode) {
	if head == nil {
		return
	}
	node := head
	nodes := make([]*ListNode, 0)
	for node != nil {
		next := node.Next
		nodes = append(nodes, node)
		node.Next = nil
		node = next
	}
	mid := len(nodes) / 2
	nodes1 := nodes[:mid]
	nodes2 := nodes[mid:]
	begin := &ListNode{}
	pre := begin
	for i := 0; i < mid; i++ {
		pre.Next = nodes1[i]
		pre = pre.Next
		pre.Next = nodes2[len(nodes2)-i-1]
		pre = pre.Next
	}
	if 2*mid < len(nodes) {
		pre.Next = nodes2[0]
	}
	head = begin.Next
}

func reverse(head *ListNode) *ListNode {
	pre := &ListNode{}
	tail := head
	curr := head
	for curr != nil {
		swap := pre.Next
		pre.Next = curr
		curr = curr.Next
		pre.Next.Next = swap
	}
	tail.Next = nil
	return pre.Next
}

//ReorderList1 R
func ReorderList1(head *ListNode) {
	curr := head
	for curr != nil && curr.Next != nil {
		tail := findTail(curr)
		if tail == curr.Next {
			break
		}
		next := curr.Next
		tail.Next = next
		curr.Next = tail
		curr = next
	}
}

func findTail(head *ListNode) *ListNode {
	pre := head
	for pre != nil && pre.Next != nil && pre.Next.Next != nil {
		pre = pre.Next
	}
	tail := pre.Next
	pre.Next = nil
	return tail
}
