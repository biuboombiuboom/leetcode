package algorithm

type stack interface {
	pop() (interface{}, error)
	push(interface{})
	top() (interface{}, error)
}
