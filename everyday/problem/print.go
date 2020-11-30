package problem

import "fmt"

func Print() {

	nums1 := []int{ 0,0,0}
	nums2 := []int{ 0,1,0}
	nums3 := []int{ 0,1,0,3,12}
	moveZeroes(nums1)
	moveZeroes(nums2)

	moveZeroes(nums3)
	fmt.Printf("%v",reorganizeString("aaabbbcccdddeee") )
}
