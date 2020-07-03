package types

import (
	"fmt"
	"sync"
)

type Stack struct{
	mutex sync.Mutex
	height int
	items []int
}

func NewStack()*Stack{
	return &Stack{
		height: 0,
		items:[]int{},
	}
}

func (stack *Stack)Pop()(int,error){
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	if stack.height==0{
		return 0,fmt.Errorf("stack is empty")
	}
	stack.height--
	r:=stack.items[stack.height]
	stack.items=stack.items[0:stack.height]
	return r,nil
}

func (stack *Stack)Push(item int){
	stack.mutex.Lock()
	stack.items=append(stack.items,item)
	stack.height++
	stack.mutex.Unlock()
}

func (stack *Stack)Peek()(int,error){

	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	if stack.height==0{
		return 0,fmt.Errorf("stack is empty")
	}
	return stack.items[stack.height-1],nil
}

func (stack *Stack)Size()int{
	return  stack.height
}
