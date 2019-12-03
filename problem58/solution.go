package problem58

//LengthOfLastWord leetcode
func LengthOfLastWord(s string) int {
	count := 0
	flag := false
	for i := len(s) - 1; i > -1; i-- {
		s1 := s[i : i+1]
		if s1 != " " {
			count++
			flag = true
		}
		if s1 == " " {
			if flag {
				return count
			}
			continue
		}
	}
	return count
}
