package problem

func matrixScore(A [][]int) int {
	for i := 0; i < len(A); i++ {
		if A[i][0] == 0 {
			reversal(A[i])
		}
	}
	ans := len(A) * 1 << (len(A[0]) - 1)
	for i := 1; i < len(A[0]); i++ {
		ones := 0
		for j := 0; j < len(A); j++ {
			if A[j][i] == 1 {
				ones++
			}
		}
		if float64(ones) < float64(len(A))/2 {
			ones = len(A) - ones
		}
		ans += ones * (1 << (len(A[0])-i - 1))
	}
	return ans
}

func reversal(a []int) {
	for i := 0; i < len(a); i++ {
		if a[i] == 0 {
			a[i] = 1
		} else {
			a[i] = 0
		}
	}
}
