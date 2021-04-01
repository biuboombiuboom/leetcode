package p202103

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	cur:=head
	pre:=&ListNode{}
	pre.Next=head
	for i := 1; i <left  &&cur!=nil; i++ {
		pre=cur
		cur=cur.Next
	}
	for i := left; i < right&&cur!=nil&&cur.Next!=nil; i++ {
		next:=cur.Next
		cur.Next=next.Next
		next.Next=pre.Next
		pre.Next=next
	}
	if left==1{
		return  pre.Next
	}
	return head
}

