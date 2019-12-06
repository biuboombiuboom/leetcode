package algorithm

//Merge 合并两个有序数组
func Merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		_ = append(nums1[0:0], nums2...)
		return
	}

	for i, j := 0, 0; i < m+n; i++ {
		n1 := nums1[i]
		for j < n {

			n2 := nums2[j]
			if n2 >= nums1[m+j-1] {
				_ = append(nums1[:m+j], nums2[j:]...)
				return
			}
			if n2 > n1 {

				break
			}
			left := nums1[:i]
			right := make([]int, m+j-i)
			copy(right, nums1[i:])
			left = append(left, n2)
			_ = append(left, right...)
			j++
			i++
		}
	}
}
