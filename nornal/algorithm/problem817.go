package algorithm

func NumComponents(head *ListNode, G []int) int {
	curr := head
	c := 0
	m := make(map[int]int)
	for i := 0; i < len(G); i++ {
		m[G[i]] = G[i]
	}
	flag := false
	for curr != nil {
		if _, ok := m[curr.Val]; !ok {
			if flag {
				c++
				flag = false
			}
		} else {
			flag = true
		}

		curr = curr.Next
	}
	if flag {
		c++
	}
	return c
}
