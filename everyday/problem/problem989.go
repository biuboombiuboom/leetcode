package problem

//对于非负整数 X 而言，X 的数组形式是每位数字按从左到右的顺序形成的数组。例如，如果 X = 1231，那么其数组形式为 [1,2,3,1]。
//
//给定非负整数 X 的数组形式 A，返回整数 X+K 的数组形式。

func addToArrayForm(A []int, K int) []int {
	n := len(A)
	for i := n - 1; i >= 0; i-- {
		l := K % 10
		K = K / 10
		A[i] += l
		if A[i] > 9 {
			A[i] = A[i] - 10
			K += 1
		}
		if K == 0 {
			break
		}
	}
	for K != 0 {
		l := K % 10
		K = K / 10
		A = append([]int{l}, A...)
	}
	return A
}
