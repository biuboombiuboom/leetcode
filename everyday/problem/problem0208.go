package problem

func detectCycle(head *ListNode) *ListNode {
	if head==nil{
		return  head
	}
	slow:=head
	fast:=head
	hasCycle:=false
	for fast!=nil&&fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
		if slow==fast{
			hasCycle=true
			break
		}
	}
	if !hasCycle{
		return  nil
	}
	cycleHead:=head
	for cycleHead!=fast{
		cycleHead=cycleHead.Next
		fast=fast.Next
	}
	return cycleHead
}