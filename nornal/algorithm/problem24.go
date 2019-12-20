package algorithm

func SwapPairs(head *ListNode) *ListNode {

	begin := &ListNode{}
	pre := begin
	for head != nil {
		n1 := head
		head = head.Next
		var n2 *ListNode
		if head == nil {
			pre.Next = n1
			break
		}
		n1.Next = nil
		n2 = head
		head = head.Next
		n2.Next = nil
		pre.Next = n2
		pre.Next.Next = n1
		// n1.Next=nil
		pre = n1
	}

	return begin.Next
}
