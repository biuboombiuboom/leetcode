package algorithm

//ReverseKGroup
func ReverseKGroup(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil || n == 1 {
		return head
	}
	begin := &ListNode{}
	pre := begin
	for {
		right, count := cut(head, n)
		if count < n {
			pre.Next = head
			break
		}
		if count == n {
			pre.Next = reverse(head, n)
		}
		for count > 0 {
			pre = pre.Next
			count--
		}
		head = right

	}
	return begin.Next
}

func reverse(head *ListNode, n int) *ListNode {
	begin := &ListNode{Next: head}
	head = head.Next
	begin.Next.Next = nil
	for head != nil {
		curr := head
		head = head.Next
		curr.Next = begin.Next
		begin.Next = curr
	}
	return begin.Next
}

func cut(head *ListNode, n int) (*ListNode, int) {
	pre := head
	count := 1
	for n > 1 && pre != nil {
		pre = pre.Next
		n--
		count++
	}
	if pre == nil {
		return nil, count - 1
	}
	next := pre.Next
	pre.Next = nil

	return next, count

}
