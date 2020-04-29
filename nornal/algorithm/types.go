package algorithm

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// type Node struct {
// 	Val   int
// 	Pre   *Node
// 	Next  *Node
// 	Child *Node
// }

//cut 切断链表前n个元素 返回后部分的表头
func cut(head *ListNode, n int) *ListNode {
	if n == 0 || head == nil {
		return nil
	}
	pre := head
	for n > 1 && pre != nil {
		pre = pre.Next
		n--
	}
	if pre == nil {
		return nil
	}
	next := pre.Next
	pre.Next = nil
	return next
}

//TreeNode 二叉树
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

//Print p
func Print(root *TreeNode) {
	if root == nil {
		return
	}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		l := len(nodes)
		for l > 0 {
			currentNode := nodes[0]
			if currentNode != nil {
				fmt.Printf("%d,", currentNode.Val)
				nodes = append(nodes, currentNode.Left, currentNode.Right)
			}
			nodes = nodes[1:]
			l--
		}
		fmt.Println()
	}
}
