package algorithm

func removeElements(head *ListNode, val int) *ListNode {
	begin := &ListNode{
		Next: head,
	}
	pre := begin
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}

	return begin.Next
}
