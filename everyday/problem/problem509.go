package problem

func fib(n int) int {
	if n == 0 {
		return 0
	}
	n1 := 0
	n2 := 1
	for i := 2; i <= n; i++ {
		n1, n2 = n2, n1+n2
	}
	return n2
}
