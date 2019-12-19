package algorithm

//RemoveNthFromEnd r
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	slow := head
	fast := head

	k := 0
	for fast.Next != nil {
		if k >= n {
			slow = slow.Next
		}
		fast = fast.Next

		k++
	}

	if slow == fast {
		return nil
	}
	if k < n {
		return head.Next
	}
	slow.Next = slow.Next.Next

	return head
}

//RemoveNthFromEnd1 r
func RemoveNthFromEnd1(head *ListNode, n int) *ListNode {
	c := 0
	nodes := make([]*ListNode, n+1)
	curr := head
	for curr != nil {
		c++
		nodes[c%(n+1)] = curr
		curr = curr.Next
	}
	c++
	preIndex := c % (n + 1)
	c++
	delIndex := c % (n + 1)
	delNode := nodes[delIndex]
	preNode := nodes[preIndex]
	if preNode == nil {
		return delNode.Next
	}
	preNode.Next = delNode.Next

	return head
}
