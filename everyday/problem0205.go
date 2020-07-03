package main

func addTwoNumbers(l1 *ListNode,l2 *ListNode)*ListNode{
	c1:=l1
	end:=c1
	c2:=l2
	for c1!=nil&&c2!=nil{
		c1.Val=c1.Val+c2.Val
		if c1.Val>9{
			c1.Val=c1.Val-10
			if c1.Next!=nil{
				c1.Next.Val++
			}else{
				c1.Next=&ListNode{
					Val: 1,
				}
			}
		}
		if c1.Next!=nil{
			end=c1.Next
		}
		c1=c1.Next
		c2=c2.Next
	}
	if c2!=nil{
		end.Next=c2
	}
	cur:=end
	for cur!=nil{
		if cur.Val>9 {
			cur.Val = cur.Val - 10
			if cur.Next != nil {
				cur.Next.Val++
			} else {
				cur.Next = &ListNode{
					Val: 1,
				}
			}
		}
		cur=cur.Next
	}
	return  l1
}
