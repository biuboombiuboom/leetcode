package problem

func longestMountain(A []int) int {
	l := len(A)

	if l < 3 {
		return 0
	}
	left := -1
	right := -1
	middle := -1
	maxLen := 0
	for i := 1; i < l; i++ {
		len := 0
		if A[i-1] == A[i] {
			if right != -1 {
				len = right - left + 1
				if len > maxLen {
					maxLen = len
				}
			}
			left = -1
			right = -1
			middle = -1
		} else if A[i-1] < A[i] {
			if left == -1 {
				left = i - 1
			} else if right != -1 {
				len = right - left + 1
				if len > maxLen {
					maxLen = len
				}
				left = i - 1
				right = -1
				middle = -1
			}
		} else {
			if left > -1 && middle == -1 {
				middle = i - 1
			}
			if middle != -1 {
				right = i
			}
		}

	}

	if right != -1 {
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}

	return maxLen
}
