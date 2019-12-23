package main

import (
	"fmt"

	"github.com/biuboombiuboom/leetcode/nornal/algorithm"
)

func main() {

	head := biuldListNode([]int{1, 4, 3, 2, 5, 2})
	l := algorithm.Partition(head, 3)
	for l != nil {
		fmt.Printf("%d,", l.Val)
		l = l.Next
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
