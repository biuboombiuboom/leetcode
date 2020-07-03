package types

type Stack interface {
	Push(interface{})
	Pop()interface{}
	IsEmpty()bool
}

type SStack struct{
	items []interface{}
	depth int
}

func NewStack()*SStack{
	return &SStack{
		items:[]interface{}{},
		depth: 0,
	}
}

func(stack *SStack)IsEmpty()bool{
	return stack.depth==0
}

func (stack *SStack)Push(val interface{}){
	stack.items=append(stack.items,val)
	stack.depth++
}

func (stack *SStack)Pop()(interface{}){
	if stack.depth>0{
		item:=stack.items[0]
		stack.items=stack.items[1:]
		stack.depth--
		return item
	}
	return nil

}