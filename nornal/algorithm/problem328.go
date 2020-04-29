package algorithm

func OddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	even := head.Next
	for odd != nil && even != nil {
		swap := even.Next
		if swap == nil {
			break
		}
		even.Next = swap.Next
		swap.Next = odd.Next
		odd.Next = swap
		odd = odd.Next
		even = even.Next
	}
	return head
}
