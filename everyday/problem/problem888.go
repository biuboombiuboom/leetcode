package problem

func fairCandySwap(A []int, B []int) []int {
	bMap := make(map[int]struct{})
	sumA, sumB := 0, 0
	for _, b := range B {
		bMap[b] = struct{}{}
		sumB += b
	}

	for _, a := range A {
		sumA += a
	}
	diff := sumA - sumB
	ans := make([]int, 0)
	for _, a := range A {
		b := a - diff/2
		if _, ok := bMap[b]; ok {
			ans = append(ans, a, b)
			break
		}
	}
	return ans
}
