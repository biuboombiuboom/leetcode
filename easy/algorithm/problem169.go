package algorithm

import "sort"

//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
// 
//
//示例 1：
//
//输入：[3,2,3]
//输出：3
//示例 2：
//
//输入：[2,2,1,1,1,2,2]
//输出：2

func majorityElement(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	j := 0
	num := nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			if i-j > n/2 {
				num = nums[i-1]
				break
			}
			j = i
		}
	}
	return num
}
