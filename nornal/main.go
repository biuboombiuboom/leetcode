package main

import (
	"fmt"

	"github.com/biuboombiuboom/leetcode/nornal/algorithm"
)

func main() {

	head := biuldListNode([]int{1})
	algorithm.ReorderList(head)
	printListNode(head)

}

func printListNode(head *algorithm.ListNode) {
	for head != nil {
		fmt.Printf("%d,", head.Val)
		head = head.Next
	}
}

func biuldListNode(nums []int) *algorithm.ListNode {
	begin := &algorithm.ListNode{}
	pre := begin
	for i := 0; i < len(nums); i++ {
		pre.Next = &algorithm.ListNode{
			Val: nums[i],
		}
		pre = pre.Next
	}
	return begin.Next
}

func buildCycleListNode(nums []int, pos int) *algorithm.ListNode {
	begin := &algorithm.ListNode{}
	pre := begin
	var cycleBegin *algorithm.ListNode
	for i := 0; i < len(nums); i++ {
		pre.Next = &algorithm.ListNode{
			Val: nums[i],
		}
		pre = pre.Next
		if pos == i {
			cycleBegin = pre
		}
	}
	pre.Next = cycleBegin
	return begin.Next
}
