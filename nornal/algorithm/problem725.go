package algorithm

func SplitListToParts(root *ListNode, k int) []*ListNode {
	parts := make([]*ListNode, k)

	l := 0
	curr := root
	for curr != nil {
		l++
		curr = curr.Next
	}
	if l == 0 {
		return parts
	}
	l1 := l / k
	re := l % k
	head := root
	for i := 0; i < k; i++ {
		begin := &ListNode{}
		pre := begin
		for j := 0; j < l1; j++ {
			pre.Next = head
			head = head.Next
			pre = pre.Next
		}
		if re > 0 {
			pre.Next = head
			head = head.Next
			pre = pre.Next
			re--
		}
		if pre.Next != nil {
			pre.Next = nil
		}
		parts[i] = begin.Next

	}

	return parts
}
