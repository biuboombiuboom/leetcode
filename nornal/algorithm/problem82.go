package algorithm

//DeleteDuplicates DeleteDuplicates
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nums := make(map[int]int)

	for head != nil {
		if n, ok := nums[head.Val]; ok {
			n++
			nums[head.Val] = n
		} else {
			nums[head.Val] = 1
		}
		head = head.Next
	}
	begin := &ListNode{}
	pre := begin
	for k, v := range nums {
		if v == 1 {
			pre.Next = &ListNode{
				Val: k,
			}
			pre = pre.Next
		}
	}
	return begin.Next
}
