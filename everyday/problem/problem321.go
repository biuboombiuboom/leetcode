package problem


/*
输入:
nums1 = [3, 4, 6, 5]
nums2 = [9, 1, 2, 5, 8, 3]
k = 5
输出:
[9, 8, 6, 5, 3]
*/
func maxNumber1(nums1 []int, nums2 []int, k int) []int {
	ans := make([]int, 0)
	for k > 0 {
		num, i1, i2 := findNext(nums1, nums2, k)
		ans = append(ans, num)
		nums1 = nums1[i1+1:]
		nums2 = nums2[i2+1:]
		k--

	}

	return ans
}

func findNext(nums1 []int, nums2 []int, k int) (int, int, int) {
	max1, index1 := findMaxWithIndex(nums1)
	max2, index2 := findMaxWithIndex(nums2)
	if index1 == -1 {
		return max2, index1, index2
	}
	if index2 == -1 {
		return max1, index1, -index2
	}
	if max1 > max2 {
		if len(nums1)+len(nums2)-index1 >= k {
			return max1, index1, -1
		} else {
			return findNext(nums1[:index1], nums2, k-len(nums1)+index1)
		}
	} else if max1 < max2 {
		if len(nums1)+len(nums2)-index2 >= k {
			return max2, -1, index2
		} else {
			return findNext(nums1, nums2[:index2], k-len(nums2)+index2)
		}
	} else {
		if len(nums1)+len(nums2)-index2 >= k && len(nums1)+len(nums2)-index1 >= k {
			_, i1, _ := findNext(nums1[index1+1:], nums2[index2+1:], k-2)
			if i1 > -1 {
				return max1, index1, -1
			} else {
				return max2, -1, index2
			}
		} else if len(nums1)+len(nums2)-index1 >= k {
			return max1, index1, -1
		} else if len(nums1)+len(nums2)-index2 >= k {
			return max2, -1, index2
		}
		return findNext(nums1[:index1], nums2[:index2], k-index2-index1-2)

	}
}

func findMaxWithIndex(nums []int) (int, int) {
	max := 0
	index := -1
	for i, num := range nums {
		if num > max {
			max = num
			index = i
		}
	}
	return max, index
}
