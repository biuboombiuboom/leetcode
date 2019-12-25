package algorithm

func DetectCycle(head *ListNode) *ListNode {
	fr := 0
	r := 0
	slow := head
	fast := head
	hasCycle := false
	for fast != nil && fast.Next != nil {
		fr = fr + 2
		r++
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
			curr = curr.Next
		}
	}
	return nil

}
