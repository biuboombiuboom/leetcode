package algorithm

func MiddleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	if slow == head || fast == slow {
		return nil
	}
	return slow
}
