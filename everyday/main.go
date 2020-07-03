package main

import "fmt"

func main() {
	//head := buildListNode([]int{1, 2})
	l1:=buildListNode([]int{8,9,9})
	l2:=buildListNode([]int{2})
	l1=addTwoNumbers(l2,l1)


	printListNode(l1)
}

func buildListNode(nums []int) *ListNode {
	begin := &ListNode{}
	pre := begin
	for i := 0; i < len(nums); i++ {
		pre.Next = &ListNode{
			Val: nums[i],
		}
		pre = pre.Next
	}
	return begin.Next
}

func printListNode(head *ListNode) {
	for head != nil {
		fmt.Printf("%d,", head.Val)
		head = head.Next
	}
}
