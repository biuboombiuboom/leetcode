package main

import "fmt"

func main() {
	head := biuldListNode([]int{1, 2})
	head = partition(head, 3)
	printListNode(head)
}

func biuldListNode(nums []int) *ListNode {
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
