package algorithm

//SortList s
func SortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var left, right *ListNode
	for step := 1; ; step += step {
		begin := &ListNode{}
		pre := begin
		cCount := 0
		for head != nil {
			left = head
			right = cut(head, step)
			cCount++
			//一刀切到尾
			if cCount == 1 && right == nil {
				return head
			}
			head = cut(right, step)
			cCount++
			pre.Next = merge(left, right)
			for pre.Next != nil {
				pre = pre.Next
			}
		}
		head = begin.Next
	}

}

func merge(left *ListNode, right *ListNode) *ListNode {
	begin := &ListNode{}
	pre := begin
	for left != nil && right != nil {
		if left.Val < right.Val {
			pre.Next = left
			left = left.Next
		} else {
			pre.Next = right
			right = right.Next
		}
		pre = pre.Next
	}
	if left == nil {
		pre.Next = right
	} else {
		pre.Next = left
	}
	return begin.Next
}

//SortList1 s
func SortList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast := head.Next
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next

	slow.Next = nil
	left, right := SortList(head), SortList(mid)
	pre := &ListNode{}
	begin := pre
	for left != nil && right != nil {
		if left.Val < right.Val {
			pre.Next = left
			left = left.Next
		} else {
			pre.Next = right
			right = right.Next
		}
		pre = pre.Next
	}
	if left != nil {
		pre.Next = left
	} else {
		pre.Next = right
	}
	return begin.Next
}
