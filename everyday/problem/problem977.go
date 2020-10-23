package problem

func sortedSquares(A []int) []int {
	index := -1
	for i := 0; i < len(A) && A[i] < 0; i++ {
		index = i
	}
	result := make([]int,0, len(A))
	for i, j := index, index+1; i >= 0 || j < len(A); {
		if i < 0 {
			result = append(result, A[j]*A[j])
			j++
		} else if j == len(A) {
			result = append(result, A[i]*A[i])
			i--
		} else if A[i]*A[i] < A[j]*A[j] {
			result = append(result, A[i]*A[i])
			i--
		} else {
			result = append(result, A[j]*A[j])
			j++
		}

	}

	return result
}
