package main

import (
	"github.com/biuboombiuboom/leetcode/nornal/algorithm"
)

func main() {

	head := buildCycleListNode([]int{3, 2, 0, -4}, 1)
	algorithm.DetectCycle(head)

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
