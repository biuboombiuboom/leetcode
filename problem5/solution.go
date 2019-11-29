package problem5

//LongestPalindrome 暴力破解
func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	result := s[0:1] 2
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			s1 := s[i : j+1]

			flag := true
			left := len(s1)/2 - 1
			right := len(s1)/2 + 1
			if len(s1)%2 == 0 {
				right--
			}
			// if len(s1)%2 == 1 {
			// 	left++
			// }
			for left > -1 && right < len(s1) {
				if s1[left] != s1[right] {
					flag = false
					break
				}
				left--
				right++
			}
			if flag && len(s1) > len(result) {
				result = s1
			}
		}
	}

	return result
}
