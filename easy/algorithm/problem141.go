package algorithm

//HasCycle h
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
