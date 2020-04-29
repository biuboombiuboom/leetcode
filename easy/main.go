package main

import (
	"fmt"

	"github.com/biuboombiuboom/leetcode/easy/algorithm"
)

func main() {
	// cross := &algorithm.ListNode{
	// 	Val: 8,
	// 	Next: &algorithm.ListNode{
	// 		Val: 4,
	// 		Next: &algorithm.ListNode{
	// 			Val: 5,
	// 		},
	// 	},
	// }
	// headA := &algorithm.ListNode{
	// 	Val: 4,
	// 	Next: &algorithm.ListNode{
	// 		Val:  1,
	// 		Next: cross,
	// 	},
	// }
	// headB := &algorithm.ListNode{
	// 	Val: 5,
	// 	Next: &algorithm.ListNode{
	// 		Val: 0,
	// 		Next: &algorithm.ListNode{
	// 			Val:  1,
	// 			Next: cross,
	// 		},
	// 	},
	// }

	headA := &algorithm.ListNode{
		Val: 1,
	}
	headB := &algorithm.ListNode{
		Val: 2,
	}
	c := algorithm.GetIntersectionNode(headA, headB)
	if c != nil {
		fmt.Println(c.Val)

	} else {
		fmt.Println("null")

	}

}
