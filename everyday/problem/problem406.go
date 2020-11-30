package problem

import "sort"

//[[7,0], [4,4], [7,1], [5,0], [6,1], [5,2]]
//[[5,0], [7,0], [5,2], [6,1], [4,4], [7,1]]
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		p1, p2 := people[i], people[j]
		return p1[0] < p2[0] || (p1[0] == p2[0] && p1[1] > p2[1])
	})
	ans := make([][]int, len(people))
	for _, person := range people {
		spaces := person[1] + 1
		for i := range ans {
			if ans[i] == nil {
				spaces--
				if spaces == 0 {
					ans[i] = person
					break
				}
			}
		}
	}

	return ans
}
