package algorithm


func CopyRandomList(head *Node) *Node {
	copyMap:=make(map[*Node]*Node)
	return doCopy(head,copyMap)
}

func doCopy(head *Node,mapNode map[*Node]*Node) *Node {
	if mapNode[head]==nil {
		copy := &Node{
			Val: head.Val,
		}
		mapNode[head]=copy
		if head.Next != nil {
			copy.Next = doCopy(head.Next,mapNode)
		}
		if head.Random != nil {
			copy.Random = doCopy(head.Random,mapNode)
		}


	}
	return mapNode[head]

}