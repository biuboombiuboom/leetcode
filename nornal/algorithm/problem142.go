package algorithm

func DetectCycle(head *ListNode) *ListNode {

	slow := head
	fast := head
	hasCycle := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasCycle = true
			break
		}
	}
	if hasCycle {
		curr := head
		for {
			if curr == fast {
				return curr
			}
			fast = fast.Next
			curr = curr.Next
		}
	}
	return nil

}
