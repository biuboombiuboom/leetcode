package problem4

//FindMedianSOrtedArrays f
func FindMedianSOrtedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	left := 0
	right := l2 - 1
	if l1 == 0 && l2 == 0 {
		return 0
	}
	for i := 0; i < l1; i++ {
		num := nums1[i]
		right = len(nums2) - 1
		for left <= right {
			mid := (left + right) / 2
			midNum := nums2[mid]
			if midNum >= num {
				right = mid - 1
			}
			if midNum == num {
				left = mid
			}
			if midNum < num {
				left = mid + 1
			}
		}
		leftNums := nums2[:left]
		rightNums := nums2[left:]
		t := append(append([]int{}, leftNums...), num)
		nums2 = append(t, rightNums...)
		left = len(leftNums) + 1

	}
	// }
	if (l1+l2)%2 == 1 {
		return float64(nums2[(l1+l2)/2])
	}
	mid := (l1 + l2) / 2

	return float64(nums2[mid]+nums2[mid-1]) / 2
}
