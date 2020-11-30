package problem

func fourSumCount(A []int, B []int, C []int, D []int) int {
	l := len(A)
	bucket := make(map[int]int)
	count := 0
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			bucket[A[i]+B[j]]++
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			count += bucket[-C[i]-D[j]]
		}
	}
	return count
}
