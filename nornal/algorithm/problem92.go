package algorithm

func ReverseBetween(head *ListNode, m int, n int) *ListNode {

	begin := &ListNode{}
	var tail *ListNode
	pre := begin
	k := 0
	for head != nil && k < n {
		k++
		if k < m {
			pre.Next = head
			pre = pre.Next
			head = head.Next
			continue
		}
		if k == m {
			tail = head
			head = head.Next
			tail.Next = nil
			pre.Next = tail
			continue
		}
		tmp := pre.Next
		pre.Next = head
		head = head.Next
		pre.Next.Next = tmp

	}
	tail.Next = head
	return begin.Next
}
