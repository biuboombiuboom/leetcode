package problem

import "sort"

func maximumProduct(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	l := len(nums)
	return max_(nums[0]*nums[1]*nums[l-1], nums[l-1]*nums[l-2]*nums[l-3])
}

func max_(a, b int) int {
	if a > b {
		return a
	}
	return b
}
