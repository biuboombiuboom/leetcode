package algorithm

import "math"

//AddTwoNumbers
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	len1 := 0
	len2 := 0
	var bigL *ListNode
	var littleL *ListNode
	h1 := l1
	h2 := l2
	for h1 != nil || h2 != nil {
		if h1 != nil {
			len1++
			h1 = h1.Next
		}
		if h2 != nil {
			len2++
			h2 = h2.Next
		}
	}
	if len1 > len2 {
		bigL, littleL = l1, l2
	} else {
		bigL, littleL = l2, l1
	}

	diffL := int(math.Abs(float64(len1 - len2)))
	for diffL > 0 {
		littleL = &ListNode{
			Val:  0,
			Next: littleL,
		}
		diffL--
	}
	begin := &ListNode{}
	pre := begin
	noC := true
	bHead := bigL
	for bHead != nil {
		bHead.Val = bHead.Val + littleL.Val
		if bHead.Val >= 10 {
			noC = false
			bHead.Val = bHead.Val - 10
			pre.Next = &ListNode{
				Val: 1,
			}
		} else {
			pre.Next = &ListNode{
				Val: 0,
			}
		}
		bHead = bHead.Next
		littleL = littleL.Next
		pre = pre.Next
	}
	if noC {
		return bigL
	} else {
		pre.Next = &ListNode{
			Val: 0,
		}
	}
	result := AddTwoNumbers(bigL, begin.Next)
	for result.Val == 0 {
		result = result.Next
	}
	return result
}
