package problem4

//FindMedianSOrtedArrays f
func FindMedianSOrtedArrays(nums1 []int, nums2 []int) float64 {

	l1 := len(nums1)
	l2 := len(nums2)

	midSum := 0
	mid := (l1+l2)/2 + 1
	count := 0
	for count < l1+l2/2 {
		mid = mid/2 + 1
		k1, k2 := mid, mid
		if l1 < mid {
			k1 = l1
		}
		if l2 < mid {
			k2 = l2
		}
		num1 := nums1[k1]
		num2 := nums1[k2]
		if num1 <= num2 {
			nums1 = nums1[k1+1:]
			count += k1
			if mid == 1 {
				midSum += num1
			}
		} else {
			nums2 = nums1[k2+1:]
			count += k2
			if mid == 1 {
				midSum += num2
			}
		}
	}

	return float64(midSum) / 2
}

// 暴力
// func FindMedianSOrtedArrays(nums1 []int, nums2 []int) float64 {
// 	l1 := len(nums1)
// 	l2 := len(nums2)
// 	left := 0
// 	right := l2 - 1
// 	if l1 == 0 && l2 == 0 {
// 		return 0
// 	}
// 	for i := 0; i < l1; i++ {
// 		num := nums1[i]
// 		right = len(nums2) - 1
// 		for left <= right {
// 			mid := (left + right) / 2
// 			midNum := nums2[mid]
// 			if midNum >= num {
// 				right = mid - 1
// 			}
// 			if midNum == num {
// 				left = mid
// 			}
// 			if midNum < num {
// 				left = mid + 1
// 			}
// 		}
// 		leftNums := nums2[:left]
// 		rightNums := nums2[left:]
// 		t := append(append([]int{}, leftNums...), num)
// 		nums2 = append(t, rightNums...)
// 		left = len(leftNums) + 1

// 	}
// 	// }
// 	if (l1+l2)%2 == 1 {
// 		return float64(nums2[(l1+l2)/2])
// 	}
// 	mid := (l1 + l2) / 2

// 	return float64(nums2[mid]+nums2[mid-1]) / 2
// }
