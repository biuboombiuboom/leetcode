package problem7

//Reverse R
func Reverse(x int) int {

	left := -2147483648
	right := 2147483647
	rev := 0
	for x != 0 {
		temp := x % 10
		x = x / 10
		rev = rev*10 + temp
	}
	if rev < left || rev > right {
		rev = 0
	}
	return rev
}
