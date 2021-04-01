package problem

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ans := 0
	si := 0
	for i := 0; i < len(g); i++ {
		for si < len(s) && g[i] > s[si] {
			si++
		}
		if si < len(s) {
			ans++
			si++
		}
	}

	return ans
}
