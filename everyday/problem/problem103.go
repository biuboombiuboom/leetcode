package problem

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	nodes := make([]*TreeNode, 0)
	if root != nil {
		nodes = append(nodes, root)
	}
	left := true
	for len(nodes) > 0 {
		nodeValues := make([]int, len(nodes))
		nextNodes := make([]*TreeNode, 0)
		for i := 0; i < len(nodes); i++ {
			if left {
				nodeValues[i] = nodes[i].Val
			} else {
				nodeValues[i] = nodes[len(nodes)-i-1].Val
			}
			if nodes[i].Left != nil {
				nextNodes = append(nextNodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nextNodes = append(nextNodes, nodes[i].Right)
			}
		}
		nodes = nextNodes
		ans = append(ans, nodeValues)
		left = !left
	}
	return ans
}
