package p202103

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{{}}
	counter := []int{1,0}
	n := len(nums)
	for i := 0; i < n; i++ {
		m := len(ans)
		start := 0
		if i > 0 && nums[i] == nums[i-1] {
			start = counter[(i+1)%2]
		}
		counter[(i+1)%2] = counter[i%2]
		for j := start; j < m; j++ {
			sub := append([]int{}, ans[j]...)
			sub = append(sub, nums[i])
			ans = append(ans, sub)
			counter[(i+1)%2] = counter[(i+1)%2] + 1
		}
	}
	return ans
}
