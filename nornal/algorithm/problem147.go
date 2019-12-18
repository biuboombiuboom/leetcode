package algorithm

import "github.com/biuboombiuboom/leetcode/nornal/types"

//InsertionSortList s
func InsertionSortList(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	begin := &types.ListNode{
		// Val:0,
		Next: head,
	}
	var pre *types.ListNode
	for head != nil && head.Next != nil {
		//寻址未排序节点
		if head.Val <= head.Next.Val {
			head = head.Next
			continue
		}
		pre = begin
		//寻址前驱节点
		for pre.Next.Val < head.Next.Val {
			pre = pre.Next
		}

		current := head.Next
		head.Next = current.Next
		current.Next = pre.Next
		pre.Next = current
	}

	return begin.Next
}

//InsertionSortList1 s
func InsertionSortList1(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// mn:=head
	min := head.Val
	next := head.Next
	for next != nil {
		nn := next.Next
		if next.Val < head.Val {
			head.Val = next.Val
			next.Val = min
			min = head.Val
		}
		next = nn
	}
	head.Next = InsertionSortList(head.Next)
	return head
}
