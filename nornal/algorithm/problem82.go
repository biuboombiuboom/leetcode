package algorithm

//DeleteDuplicates DeleteDuplicates
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	begin := &ListNode{
		Next: head,
	}
	pre := begin
	curr := pre.Next
	next := curr.Next
	//
	delFlag := false
	for curr != nil && next != nil {
		if curr.Val == next.Val {
			delFlag = true
			next = next.Next
			curr.Next = next
		} else {
			if delFlag {
				delFlag = !delFlag
				curr = next
				pre.Next = curr
				next = curr.Next
			} else {
				pre = curr
				curr = next
				next = next.Next
			}
		}
	}
	if delFlag {
		pre.Next = next
	}
	return begin.Next
}
