package algorithm

//IsHappy u
func IsHappy(n int) bool {
	nums := make(map[int]int)
	result := 0
	for n%10 != 0 {
		i := n / 10
		n = n / 10
		result += i * i
		n = result
		result = 0
		if _, ok := nums[n]; ok {
			return false
		}
	}
	return true
}
