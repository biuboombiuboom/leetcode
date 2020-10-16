package main

func connect(root *Node) *Node {
	cur := root
	for cur != nil {
		head := &Node{}
		pre := head
		for cur != nil && cur.Left != nil {
			pre.Next = cur.Left
			pre = pre.Next
			pre.Next = cur.Right
			pre = pre.Next
			cur = cur.Next
		}
	}
	return root
}
