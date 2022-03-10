package p202103

func deleteDuplicates(head *ListNode) *ListNode {
	begin := &ListNode{Next: head}
	pre := begin
	cur := head
	delete := false
	for cur != nil && cur.Next != nil {
		//遇到重复源数跳
		if cur.Val == cur.Next.Val {
			pre.Next = cur.Next
			delete = true
		} else if delete {
			pre.Next = cur.Next
			delete=false
		} else {
			pre = pre.Next
		}
		cur=cur.Next
	}
	if delete{
		pre.Next=pre.Next.Next
	}
	return begin.Next
}
