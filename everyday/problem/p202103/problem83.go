package p202103

func deleteDuplicates1(head *ListNode) *ListNode {
	cur:=head
	for cur!=nil&&cur.Next!=nil{
		if cur.Val==cur.Next.Val{
			cur.Next=cur.Next.Next
		}else{
			cur=cur.Next
		}
	}
	return head
}
