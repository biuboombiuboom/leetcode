package problem21

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	n := &ListNode{}
	if l1.Val <= l2.Val {
		n.Val = l1.Val
		n.Next = MergeTwoLists(l1.Next, l2)
	} else {
		n.Val = l2.Val
		n.Next = MergeTwoLists(l1, l2.Next)
	}

	return n
}

//ListNode s
type ListNode struct {
	Val  int
	Next *ListNode
}
