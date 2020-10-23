package problem

func reorderList(head *ListNode)  {
	if head == nil {
		return
	}
	nodes := make([]*ListNode, 0)
	node := head
	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}
	i := 0
	j := len(nodes) - 1
	for i < j{
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil

}
