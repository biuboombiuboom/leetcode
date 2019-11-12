package problem9

//IsPalindrome 1
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	temp := x
	rev := 0
	for x > 0 {
		rev = rev*10 + x%10
		x = x / 10
	}

	return rev-temp == 0
}
