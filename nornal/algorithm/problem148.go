package algorithm

import "github.com/biuboombiuboom/leetcode/nornal/types"

//SortList s
func SortList(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var left, right *types.ListNode
	for step := 1; ; step += step {
		begin := &types.ListNode{}
		pre := begin
		cCount := 0
		for head != nil {
			left = head
			right = cut(head, step)
			cCount++
			//一刀切到尾
			if cCount == 1 && right == nil {
				return head
			}
			head = cut(right, step)
			cCount++
			pre.Next = merge(left, right)
			for pre.Next != nil {
				pre = pre.Next
			}
		}
		head = begin.Next
	}

}

//cut 切断链表前n个元素 返回后部分的表头
func cut(head *types.ListNode, n int) *types.ListNode {
	if n == 0 || head == nil {
		return nil
	}
	pre := head
	for n > 1 && pre != nil {
		pre = pre.Next
		n--
	}
	if pre == nil {
		return nil
	}
	next := pre.Next
	pre.Next = nil
	return next
}

func merge(left *types.ListNode, right *types.ListNode) *types.ListNode {
	begin := &types.ListNode{}
	pre := begin
	for left != nil && right != nil {
		if left.Val < right.Val {
			pre.Next = left
			left = left.Next
		} else {
			pre.Next = right
			right = right.Next
		}
		pre = pre.Next
	}
	if left == nil {
		pre.Next = right
	} else {
		pre.Next = left
	}
	return begin.Next
}

//SortList1 s
func SortList1(head *types.ListNode) *types.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast := head.Next
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next

	slow.Next = nil
	left, right := SortList(head), SortList(mid)
	pre := &types.ListNode{}
	begin := pre
	for left != nil && right != nil {
		if left.Val < right.Val {
			pre.Next = left
			left = left.Next
		} else {
			pre.Next = right
			right = right.Next
		}
		pre = pre.Next
	}
	if left != nil {
		pre.Next = left
	} else {
		pre.Next = right
	}
	return begin.Next
}
