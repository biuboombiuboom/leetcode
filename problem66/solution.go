package problem66

//PlusOne p
func PlusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i > -1; i-- {
		num := digits[i]
		num++
		if num < 10 {
			digits[i] = num
			return digits
		}
		digits[i] = num % 10
	}

	return append([]int{1}, digits...)

}

//PlusOne1 p
func PlusOne1(digits []int) []int {
	l := len(digits)

	if l == 0 {
		return []int{1}
	}
	num := digits[l-1]
	num++
	if num >= 10 {
		num = num - 10
		digits = append(PlusOne(digits[0:l-1]), num)
	} else {
		digits[l-1] = num
	}

	return digits
}
