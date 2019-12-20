package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

//cut 切断链表前n个元素 返回后部分的表头
func cut(head *ListNode, n int) *ListNode {
	if n == 0 || head == nil {
		return nil
	}
	pre := head
	for n > 1 && pre != nil {
		pre = pre.Next
		n--
	}
	if pre == nil {
		return nil
	}
	next := pre.Next
	pre.Next = nil
	return next
}
