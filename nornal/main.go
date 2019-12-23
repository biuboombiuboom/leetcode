package main

import (
	"fmt"

	"github.com/biuboombiuboom/leetcode/nornal/algorithm"
)

func main() {

	head := biuldListNode([]int{1, 2, 3, 4, 5})
	l := algorithm.ReverseBetween(head, 2, 5)
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
