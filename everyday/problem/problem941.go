package problem

func validMountainArray(A []int) bool {
	left := -1
	top := 0
	for i := 1; i < len(A); i++ {
		if A[i-1] == A[i] {
			return false
		}
		if A[i-1] < A[i] {
			if top == 0 {
				if left == -1 {
					left = i
				}
				continue
			} else {
				return false
			}
		}
		if A[i-1] > A[i] {
			if top == 0 {
				top = i
			} else {
				continue
			}
		}

	}
	if top == 0 || left == -1 {
		return false
	}
	return true
}
