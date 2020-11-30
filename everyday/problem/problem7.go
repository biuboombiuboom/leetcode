package problem

func reverse(x int) int {
	flag := false
	if x < 0 {
		flag = true
		x = x * -1
	}
	min := 2 << 30 * -1
	max := 2<<30 - 1
	ans := 0
	for x > 0 {
		ln := x % 10
		ans = ans*10 + ln
		x = x / 10
	}
	if flag {
		ans = ans * -1
	}
	if ans < min || ans > max {
		return 0
	}
	return ans
}
