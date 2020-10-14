package offer

func getKthFromEnd1(head *ListNode, k int) *ListNode {
	// if head == nil {
	// 	return nil
	// }
	lastList := make([]*ListNode, k)
	for {
		for i := 0; i < k; i++ {
			if head == nil {
				return lastList[i]
			}
			lastList[i] = head
			head = head.Next
		}
	}
}

//双指针
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	slow := head
	for k > 0 && fast != nil {
		fast = fast.Next
		k--
	}
	if k > 0 {
		return nil
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow

}

type ListNode struct {
	Var  int
	Next *ListNode
}
