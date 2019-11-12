package problem9

//IsPalindrome 1
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x = x / 10
	}

	return x-rev == 0 || x-rev/10 == 0
}
