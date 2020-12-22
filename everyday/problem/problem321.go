package problem

/*
输入:
nums1 = [3, 4, 6, 5]
nums2 = [9, 1, 2, 5, 8, 3]
k = 5
输出:
[9, 8, 6, 5, 3]
*/
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	ans := make([]int, k)
	for i := 0; i < k; i++ {
		max1 := findMax(nums1, i)
		max2 := findMax(nums2, k-i)
		maxNumber := mergeNums(max1, max2)
		if less(ans, maxNumber) {
			ans = maxNumber
		}
	}
	return ans
}

func findMax(nums []int, k int) []int {
	ans:=make([]int,0)
	for i, v := range nums {
		for len(ans) > 0 && len(ans)+len(nums)-1-i >= k && v > ans[len(ans)-1] {
			ans = ans[:len(ans)-1]
		}
		if len(ans) < k {
			ans = append(ans, v)
		}
	}
	return ans

}

func mergeNums(nums1, nums2 []int) []int {
	ans := make([]int, len(nums1)+len(nums2))
 	for i := 0; i < len(ans); i++ {
		if less(nums1, nums2) {
			ans[i] = nums2[0]
			nums2 = nums2[1:]
		} else {
			ans[i] = nums1[0]
			nums1 = nums1[1:]
		}
	}
	return ans
}

func less(nums1, nums2 []int) bool {
	for i := 0; i < len(nums1) && i < len(nums2); i++ {
		if nums1[i] != nums2[i] {
			return nums1[i] < nums2[i]
		}
	}
	return len(nums1)<len(nums2)
}
