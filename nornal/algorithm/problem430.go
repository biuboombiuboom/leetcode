package algorithm

func flatten(head *Node) *Node {
	curr := head
	for curr != nil && curr.Child != nil {
		swap := curr.Next
		next := flatten(curr.Child)
		curr.Next = next
		next.Pre = curr
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = swap
		curr = curr.Next
	}
	return head
}
