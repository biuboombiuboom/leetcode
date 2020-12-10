package problem

func lemonadeChange(bills []int) bool {

	five := 0
	ten := 0
	for _, bill := range bills {
		if bill == 5 {
			five++

		} else if bill == 10 {
			ten++
			five--
			if five < 0 {
				return false
			}
		} else {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five>2 {
				five -= 3
			}else{
				return false
			}
		}
	}
	return true
}
