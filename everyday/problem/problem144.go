package problem

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	nums := []int{root.Val}
	leftNums := preorderTraversal(root.Left)
	rightNums := preorderTraversal(root.Right)
	if leftNums != nil {
		nums = append(nums, leftNums...)
	}
	if rightNums != nil {
		nums = append(nums, rightNums...)
	}
	return nums
}

//moriis
func preorderTraversalMorris(root *TreeNode) []int {
	nums := make([]int, 0)
	cur := root
	for cur != nil {
		pre := cur.Left
		if pre != nil {
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}
			if pre.Right == nil {
				nums = append(nums, cur.Val)
				pre.Right = cur
				cur = cur.Left
				continue
			}
			pre.Right = nil

		} else {
			nums = append(nums, cur.Val)
		}
		cur = cur.Right
	}

	return nums
}

func preorderTraversal1(root *TreeNode) (vals []int) {
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				vals = append(vals, p1.Val)
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
		} else {
			vals = append(vals, p1.Val)
		}
		p1 = p1.Right
	}
	return
}
