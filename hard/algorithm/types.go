package algorithm

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type myStack struct {
	items []int
	len   int
}

func (s *myStack) pop() (int, error) {
	if s == nil || s.len == 0 {
		return 0, fmt.Errorf("")
	}
	top := s.items[s.len-1]
	s.items = s.items[:s.len-1]
	return top, nil
}

func (s *myStack) top() (int, error) {
	if s == nil || s.len == 0 {
		return 0, fmt.Errorf("")
	}
	return s.items[s.len-1], nil
}

func (s *myStack) push(item int) {
	s.items = append(s.items, item)
}

//cut 切断链表前n个元素 返回后部分的表头
// func cut(head *ListNode, n int) *ListNode {
// 	if n == 0 || head == nil {
// 		return nil
// 	}
// 	pre := head
// 	for n > 1 && pre != nil {
// 		pre = pre.Next
// 		n--
// 	}
// 	if pre == nil {
// 		return nil
// 	}
// 	next := pre.Next
// 	pre.Next = nil
// 	return next
// }
