package algorithm

func MergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 1 {
		return lists[l-1]
	}
	if l == 0 {
		return nil
	}
	iteration := 1

	for iteration < l {

		for i := 0; i < (l+1)/(iteration*2); i++ {
			i1 := i * 2
			i2 := i*2 + 1
			l1 := lists[i1]
			if i2 < l {
				l2 := lists[i2]
				lists[i] = mergeList(l1, l2)
				lists[i2] = nil
			} else {
				lists[i] = mergeList(l1, nil)
			}
		}
		l = (l + 1) / 2
		iteration = iteration * 2
	}
	return lists[0]
}

func mergeList(left *ListNode, right *ListNode) *ListNode {
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
	if left != nil {
		pre.Next = left
	} else {
		pre.Next = right
	}
	return begin.Next
}
