package main

import (
	"fmt"

	"github.com/biuboombiuboom/leetcode/hard/algorithm"
)

func main() {


	a:=algorithm.LargestRectangleArea([]int{2,1,5,6,2,3})
	fmt.Print(a)

	//node := biuldListNode([]int{1, 2, 3, 4, 5})
	//node = algorithm.ReverseKGroup(node, 2)
	//for node != nil {
	//	fmt.Printf("%d,", node.Val)
	//	node = node.Next
	//}
	//fmt.Print("")
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
