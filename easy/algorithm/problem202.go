package algorithm

//IsHappy u
func IsHappy(n int) bool {
	if n == 1 {
		return true
	}
	nums := make(map[int]int)

	for n != 1 {
		if _, ok := nums[n]; ok {
			break
		}
		nums[n] = n
		newN := 0
		for n >= 10 {
			i := n % 10
			newN += i * i
			n = n / 10
		}
		n = newN + n*n

	}

	return false
}
