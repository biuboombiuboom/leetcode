package tree

import "fmt"

//BinaryTree b-tree
type BinaryTree struct {
}

//BTreeNode node
type BTreeNode struct {
	Left  *BTreeNode
	Right *BTreeNode
	Data  int
}

//Mirror M
func Mirror(root *BTreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	swapNode := root.Left
	root.Left = root.Right
	root.Right = swapNode
	Mirror(root.Left)
	Mirror(root.Right)
}

//IsBalanced b
func IsBalanced(root *BTreeNode) (bool, int) {
	dp := 0
	if &root == nil {
		return true, dp
	}
	lb, lDp := IsBalanced(root.Left)
	rb, rDp := IsBalanced(root.Right)
	diffDp := lDp - rDp
	if lDp > rDp {
		dp = 1 + lDp
	} else {
		dp = 1 + rDp
	}
	if lb && rb && diffDp >= -1 && diffDp <= 1 {
		return true, dp
	}
	return false, dp
}

//Print p
func Print(root *BTreeNode) {
	if root == nil {
		return
	}
	nodes := []*BTreeNode{root}
	for len(nodes) > 0 {
		l := len(nodes)
		for l > 0 {
			currentNode := nodes[0]
			if currentNode != nil {
				fmt.Printf("%d,", currentNode.Data)
				nodes = append(nodes, currentNode.Left, currentNode.Right)
			}
			nodes = nodes[1:]
			l--
		}
		fmt.Println()
	}
}
