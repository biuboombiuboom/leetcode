package algorithm

import (
	"github.com/biuboombiuboom/leetcode/types"
)

func LargestRectangleArea(heights []int)int{
	if len(heights)==0{
		return  0
	}

	stack:=types.NewStack()
	stack.Push(-1)
	maxarea:=0

	for i := 1; i <len(heights) ; i++ {
		top,_:=stack.Peek()
		for top!=-1&&heights[top]>heights[i]{
			top,_=stack.Pop()
			pt,_:=stack.Peek()
			area:=(i-pt-1)*heights[top]
			if area>maxarea{
				maxarea=area
			}
		}
		stack.Push(i)

	}
	for {
		sp,_:=stack.Peek();
		if sp==-1{
			break
		}
		top,_:=stack.Pop()
		pt,_:=stack.Peek()
		area:=heights[top]*(len(heights)-pt-1)
		if area>maxarea{
			maxarea=area
		}
	}
	return maxarea
}

