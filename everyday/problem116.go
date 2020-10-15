package main

func connect(root *Node) *Node {
	if root == nil{
		return  root
	}
	curr:=root
	for curr != nil {
		head := &Node{}
		pre:=head
		for curr!=nil&& curr.Left!=nil{
			pre.Next=curr.Left
			pre=pre.Next
			pre.Next=curr.Right
			pre=pre.Next
			curr=curr.Next
		}
		curr=head.Next
	}
	return root
}
