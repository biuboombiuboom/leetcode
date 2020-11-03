package problem

func sumNumbers(root *TreeNode) int {


	return sumNode(root, 0)

}

func sumNode(node *TreeNode, pSum int) int {
	if node == nil {
		return 0
	}
	sum:=pSum*10+node.Val
	if node.Right == nil&& node.Left == nil {
		return sum
	}
	return sumNode(node.Right,sum)+sumNode(node.Left,sum)
}
