package p202103

import "fmt"

func Print() {

	//newList:= reverseBetween(getListNodeFromArray([]int{1,2,3,4,5}),1,4)
	//printLN(newList)
	//fmt.Printf("%v\n", a)
	//nl := []*NestedInteger{
	//	//{
	//	//	Integers: []*NestedInteger{
	//	//		{
	//	//			Integer: 3,
	//	//		},
	//	//	},
	//	//},
	//	{
	//		Integers: []*NestedInteger{},
	//	},
	//
	//}
	//
	//ns := Constructor(nl)
	//for ns.HasNext() {
	//	fmt.Printf("%d,", ns.Next())
	//}

	//head:=getListNodeFromArray([]int{1,1})

	fmt.Printf("%v",clumsy(10))
	//10*9/8+7-6*5/4+3-2*1
	//6*5/4+3-2*1
	//30/4+3-2
	//7+3-2

}

func printTwoDArray(td [][]int) {
	for i := 0; i < len(td); i++ {
		for j := 0; j < len(td[i]); j++ {
			fmt.Printf("%d,", td[i][j])
		}
		fmt.Println()
	}
}

func printLN(head *ListNode) {
	cur := head
	for cur != nil {
		fmt.Printf("%d,", cur.Val)
		cur = cur.Next
	}
}

func getListNodeFromArray(arr []int) *ListNode {
	begin := &ListNode{}
	cur := begin
	for i := 0; i < len(arr); i++ {
		cur.Next = &ListNode{
			Val: arr[i],
		}
		cur = cur.Next
	}
	return begin.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
