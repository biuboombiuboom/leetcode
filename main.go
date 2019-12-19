package main

import (
	"fmt"

	nornal "github.com/biuboombiuboom/leetcode/nornal/algorithm"
	nTypes "github.com/biuboombiuboom/leetcode/nornal/types"
)

func main() {

	// n5 := &nornal.ListNode{
	// 	Val: 5,
	// }
	// n4 := &nornal.ListNode{
	// 	Val:  4,
	// 	Next: n5,
	// }
	// n3 := &nornal.ListNode{
	// 	Val:  3,
	// 	Next: n4,
	// }
	n2 := &nornal.ListNode{
		Val:  2,
		Next: nil,
	}
	head := &nornal.ListNode{
		Val:  1,
		Next: n2,
	}

	fmt.Println(nornal.RemoveNthFromEnd(head, 2).Val)

	return
	l := &nTypes.ListNode{
		Val: 2,
		Next: &nTypes.ListNode{
			Val:  4,
			Next: nil,
		},
	}

	println(l.Val)
	next := l.Next
	for next != nil {
		fmt.Println(next.Val)
		next = next.Next
	}

	return

	fmt.Println(nornal.Convert("LEETCODEISHIRING", 3))
}
