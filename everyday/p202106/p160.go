package p202106

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//[4,1,8,4,5]
// [5,6,1,8,4,5]
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tail := headA
	for tail != nil && tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = headB

	slow := headA
	fast := headA

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = headA
			break
		}
	}
	for fast != nil && fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	tail.Next = nil

	return fast
}

type ListNode struct {
	Val  int
	Next *ListNode
}
