package algorithm

func RemoveZeroSumSublist(head *ListNode) *ListNode {
	begin := &ListNode{Next: head}

	m := make(map[int]*ListNode)
	cu := begin
	sum := 0
	for cu != nil {
		sum += cu.Val
		m[sum] = cu
		cu = cu.Next
	}
	sum = 0
	for c := begin; c != nil; c = c.Next {
		sum += c.Val
		c.Next = m[sum].Next
	}

	return begin.Next
}
