package algorithm

//DeleteDuplicates d
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre := head
	next := head.Next
	for next != nil {
		if next.Val == pre.Val {
			pre.Next = next.Next
		} else {
			pre = next
		}
		next = next.Next
	}
	return head
}
