package problem13

//RomanToInt r
func RomanToInt(s string) int {
	roman := make(map[string]int, 8)
	roman[""] = 0
	roman["I"] = 1
	roman["V"] = 5
	roman["X"] = 10
	roman["L"] = 50
	roman["C"] = 100
	roman["D"] = 500
	roman["M"] = 1000
	x := 0
	for i := 0; i < len(s); i++ {
		r1 := s[i : i+1]
		r2 := ""
		if i < len(s)-1 {
			r2 = s[i+1 : i+2]
		}
		if roman[r1] < roman[r2] {
			x -= roman[r1]
		} else {
			x += roman[r1]
		}
	}
	return x
}
