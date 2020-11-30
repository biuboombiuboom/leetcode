package problem

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	begin := &ListNode{}
	cur, cur1, cur2 := begin, l1, l2
	decade := 0
	for cur1 != nil || cur2 != nil {
		num1, num2 := 0, 0
		if cur1 != nil {
			num1 = cur1.Val
			cur1 = cur1.Next
		}
		if cur2 != nil {
			num2 = cur2.Val
			cur2 = cur2.Next
		}

		sum := num1 + num2 + decade
		decade = 0
		cur.Next = &ListNode{
			Val: sum,
		}
		cur = cur.Next
		if sum > 9 {
			cur.Val = cur.Val - 10
			decade = 1
		}
	}
	if decade==1{
		cur.Next=&ListNode{
			Val: 1,
		}
	}

	return begin.Next
}
