package problem

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	pre := &ListNode{}
	curr := head
	for curr != nil {
		if curr.Val < x {
			pre.Next = curr.Next
			if curr != head {
				curr.Next = head
			} else {
				pre = head
			}
			head = curr
			curr = pre.Next
		} else {
			pre = curr
			curr = curr.Next
		}
	}
	return head
}

func partition1(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	leftBegin := &ListNode{}
	lPre := leftBegin
	rightBegin := &ListNode{}
	rPre := rightBegin
	curr := head
	for curr != nil {
		if curr.Val < x {
			lPre.Next = curr
			rPre.Next = curr.Next
			curr.Next = nil
			curr = rPre.Next
			lPre = lPre.Next
		} else {
			rPre.Next = curr
			rPre = rPre.Next
			curr = curr.Next
		}
	}
	lPre.Next = rightBegin.Next
	return leftBegin.Next
}
