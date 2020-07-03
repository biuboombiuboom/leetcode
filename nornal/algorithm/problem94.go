package algorithm

import "github.com/biuboombiuboom/leetcode/nornal/types"

func inorderTraversal(root *TreeNode) []int {
	r := []int{}
	curr := root
	stack := types.NewStack()
	for curr != nil || !stack.IsEmpty() {
		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}
		curr=stack.Pop().(*TreeNode)
		r=append(r,curr.Val)
		curr=curr.Right
	}
	return r
}
