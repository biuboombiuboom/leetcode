package algorithm

type MyLinkedList struct {
	Val  int
	Next *MyLinkedList
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	node := this
	for index > 0 && node != nil {
		index--
		node = node.Next
	}
	if node != nil {
		return node.Val
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	next := this.Next
	this.Next = &MyLinkedList{
		Val:  this.Val,
		Next: next,
	}
	this.Val = val
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	tailPre := this
	for tailPre != nil && tailPre.Next != nil {
		tailPre = tailPre.Next
	}
	tailPre.Next = &MyLinkedList{
		Val: val,
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
	}
	pre := this
	for pre.Next != nil {
		index--
		pre = pre.Next
		if index == 1 {
			next := pre.Next
			pre.Next = &MyLinkedList{
				Val:  this.Val,
				Next: next,
			}
			break
		}
	}
	if index > 0 {
		return
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	pre := this
	var del *MyLinkedList
	for index > 0 && pre.Next != nil {
		pre = pre.Next
		del = pre.Next
		index--
	}
	if del != nil {
		pre.Next = del.Next
	}
}
