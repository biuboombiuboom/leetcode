package algorithm

func NextLargerNodes(head *ListNode) []int {
	answer := make([]int, 0)
	stack := make([]t, 0)
	index := 0
	for current := head; current != nil; current = current.Next {
		answer = append(answer, 0)
		l := len(stack)
		for l > 0 && current.Val > stack[l-1].v {
			last := stack[l-1]
			answer[last.i] = current.Val
			stack = stack[0 : l-1]
			l = len(stack)
		}
		stack = append(stack, t{
			i: index,
			v: current.Val,
		})
		index++
	}
	return answer
}

type t struct {
	i int
	v int
}
