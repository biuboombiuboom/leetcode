package algorithm

//GetIntersectionNode g
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	var cross *ListNode

	if headA == nil || headB == nil {
		return cross
	}
	if headA == headB {
		return headA
	}

	// for endA != nil && endA.Next != nil {
	// 	endA = endA.Next
	// }
	endA := headA
	endB := headB
	lenA := 0
	lenB := 0
	maxLen := 0

	maxLN := headA
	minLN := headB
	for endA != nil && endA.Next != nil {
		endA = endA.Next
		lenA++
	}
	for endB != nil && endB.Next != nil {
		endB = endB.Next
		lenB++
	}
	if endA != endB {
		return nil
	}
	if lenA == lenB {
		return headA
	}
	diffLen := lenA - lenB
	maxLen = lenA
	if lenB > lenA {
		diffLen = lenB - lenA
		maxLN = headB
		minLN = headA
		maxLen = lenB
	}
	p1 := maxLN
	p2 := minLN
	for maxLen > 0 {
		p1 = p1.Next
		if diffLen == 0 {
			p2 = p2.Next
		}
		if p1 == p2 {
			return p1
		}
		if diffLen > 0 {
			diffLen--
		}
		maxLen--
	}
	return nil
}
