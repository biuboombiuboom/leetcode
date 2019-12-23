package algorithm

//Partition p
func Partition(head *ListNode, x int) *ListNode {
	left := &ListNode{}
	right := &ListNode{}
	lPre := left
	rPre := right
	for head != nil {
		if head.Val >= x {
			rPre.Next = head
			rPre = rPre.Next
		} else {
			lPre.Next = head
			lPre = lPre.Next
		}
		head = head.Next
	}
	if rPre != nil {
		rPre.Next = nil
	}
	lPre.Next = right.Next

	return left.Next

}
