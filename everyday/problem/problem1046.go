package problem

import "sort"

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)
	l := len(stones)
	if l == 0 {
		return 0
	}
	y := stones[l-1]
	if l >= 2 {
		x := stones[l-2]
		newStones := stones[0 : l-2]
		if y > x {
			newStones = append(newStones, y-x)
		}

		y = lastStoneWeight(newStones)
	}
	return y
}
