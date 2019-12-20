package main

import (
	"fmt"

	nornal "github.com/biuboombiuboom/leetcode/nornal/algorithm"
	nTypes "github.com/biuboombiuboom/leetcode/nornal/types"
)

func main() {

	n5 := &nornal.ListNode{
		Val: 5,
	}
	n4 := &nornal.ListNode{
		Val:  4,
		Next: n5,
	}
	n3 := &nornal.ListNode{
		Val:  3,
		Next: n4,
	}
	n2 := &nornal.ListNode{
		Val:  2,
		Next: n3,
	}
	head := &nornal.ListNode{
		Val:  1,
		Next: n2,
	}

	l := nornal.SwapPairs(head)
	println(l.Val)
	n := l.Next
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
	return
	l1 := &nTypes.ListNode{
		Val: 2,
		Next: &nTypes.ListNode{
			Val:  4,
			Next: nil,
		},
	}

	println(l1.Val)
	next := l1.Next
	for next != nil {
		fmt.Println(next.Val)
		next = next.Next
	}

	return

	fmt.Println(nornal.Convert("LEETCODEISHIRING", 3))
}
