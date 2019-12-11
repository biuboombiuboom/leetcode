package algorithm

//Merge 合并重叠空间
func Merge(intervals [][]int) [][]int {
	newIntervals := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		for j := i + 1; j < len(intervals); j++ {
			if overlap, nums1, nums2 := checkOverlap(intervals[i], intervals[j]); overlap {
				intervals[i] = nums2
				intervals[j] = nums1
				break
			}
		}
	}
	for i := 0; i < len(intervals); i++ {
		if intervals[i] != nil {
			newIntervals = append(newIntervals, intervals[i])
		}
	}
	return newIntervals
}

//检查两个空间是否重叠,
//重叠 则返回合并后的空间
//不重叠 则返回原始空间
func checkOverlap(nums1 []int, nums2 []int) (bool, []int, []int) {
	left1 := nums1[0]
	left2 := nums2[0]

	right1 := nums1[1]
	right2 := nums2[1]

	right := right1
	if right2 > right {
		right = right2
	}

	if left1 <= left2 && left2 <= right1 {
		return true, []int{left1, right}, nil
	}
	if left2 <= left1 && left1 <= right2 {
		return true, []int{left2, right}, nil
	}
	return false, nums1, nums2
}
