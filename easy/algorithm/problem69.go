package algorithm

//MySqrt m
func MySqrt(x int) int {
	left := 0
	right := x/2 + 1
	for left < right {
		mid := (left + right + 1) >> 1
		square := mid * mid
		if square > x {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return left
}
