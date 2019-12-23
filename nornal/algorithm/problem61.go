package algorithm

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l := 0
	tail := head
	for l = 1; tail.Next != nil; l++ {
		tail = tail.Next
	}

	tail.Next = head

	new_tail := head
	for i := 0; i < l-k%l-1; i++ {
		new_tail = new_tail.Next
	}

	new_head := new_tail.Next

	new_tail.Next = nil
	return new_head
}
