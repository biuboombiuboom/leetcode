package problem

func isPalindrome(head *ListNode) bool {

	nodes := make([]*ListNode, 0)
	curr := head
	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}
	for i := 0; i < len(nodes)/2; i++ {
		if nodes[i].Val != nodes[len(nodes)-i-1].Val {
			return false
		}
	}
	return true
}

func isPalindrome1(head *ListNode) bool {

	pre := &ListNode{}
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		s := slow
		slow = slow.Next
		s.Next = pre.Next
		pre.Next = s
	}
	if fast != nil {
		slow = slow.Next
	}
	cur := pre.Next
	for cur != nil {
		if cur.Val != slow.Val {
			return false
		}
		cur = cur.Next
		slow = slow.Next
	}

	return true
}
