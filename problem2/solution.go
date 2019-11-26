package problem2

//AddTwoNumbers  a
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
	l1.Val = l1.Val + l2.Val
	if l1.Val >= 10 {
		l1.Val = l1.Val - 10
		l1.Next = AddTwoNumbers(l1.Next, &ListNode{1, nil})
	}
	l1.Next = AddTwoNumbers(l1.Next, l2.Next)

	return l1
}

//ListNode s
type ListNode struct {
	Val  int
	Next *ListNode
}
