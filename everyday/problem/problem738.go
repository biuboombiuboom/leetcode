package problem

func monotoneIncreasingDigits(N int) int {
	ns := make([]int, 0)
	for N > 0 {
		ns = append(ns, N%10)
		N = N / 10
	}

	for i := len(ns) - 1; i > 0; i-- {
		if ns[i] > ns[i-1] {
			ns[i] = ns[i] - 1
			j := i - 1
			for j >= 0 {
				ns[j] = 9
				j--
			}
			for i < len(ns)-1 {
				if ns[i] >= ns[i+1] {
					break
				} else {
					ns[i] = 9
					ns[i+1]--
				}
				i++
			}
			break
		}
	}
	ans := 0
	for i := len(ns) - 1; i >= 0; i-- {
		ans = 10*ans + ns[i]
	}
	return ans
}

func isIncrease(n int) bool {
	if n < 10 {
		return true
	}
	lastLe := n % 10
	n = n / 10
	for n > 10 {
		if n%10 >= lastLe {
			return false
		}
		lastLe = n % 10
		n = n / 10
	}
	return true
}
